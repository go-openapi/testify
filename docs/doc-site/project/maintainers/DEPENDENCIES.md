---
title: Optional Dependencies
description: Zero-dependency architecture with opt-in external features.
weight: 4
---

{{% notice primary "TL;DR" "meteor" %}}
> The main module has zero external dependencies. Optional features (YAML, colorized output) are activated
> by importing separate `enable/` modules. Internal stubs panic with helpful messages when a feature is used
> without being enabled.
{{% /notice %}}

## The Problem

Testing libraries sit at the bottom of the dependency tree: every package in a project imports them.
Any dependency pulled in by the testing library propagates to all consumers. The original `stretchr/testify`
pulls in `gopkg.in/yaml.v3`, `github.com/davecgh/go-spew`, and `github.com/pmw/go-difflib` for *all*
users, even those who never call `YAMLEq`.

Our goal: **zero external dependencies in the main module**, with opt-in features for users who need them.

## Architecture Overview

Three layers collaborate to deliver optional features without coupling:

{{< mermaid align="center" zoom="true" >}}
graph TD
    user["User Code"]
    enablemod["enable/yaml  or  enable/colors<br/><i>separate Go modules</i>"]
    stubs["enable/stubs/yaml  or  enable/stubs/colors<br/><i>public delegation API</i>"]
    internal["internal/assertions/enable/yaml  or  .../colors<br/><i>internal stubs with function pointers</i>"]
    assertions["internal/assertions/yaml.go  or  equal.go, diff.go<br/><i>assertion implementations</i>"]

    user -- "blank import" --> enablemod
    enablemod -- "init(): wires real impl" --> stubs
    stubs -- "delegates to" --> internal
    assertions -- "calls" --> internal

    style enablemod fill:#4a9eff,color:#fff
    style stubs fill:#90ee90,color:#000
    style internal fill:#ffb6c1,color:#000
    style assertions fill:#f0f0f0,color:#000
{{< /mermaid >}}

| Layer | Location | Has external deps? | Purpose |
|-------|----------|-------------------|---------|
| **Feature module** | `enable/yaml/`, `enable/colors/` | Yes (own `go.mod`) | Imports the real library, wires it in via `init()` |
| **Public stubs** | `enable/stubs/yaml/`, `enable/stubs/colors/` | No | Stable public API that delegates to internal package |
| **Internal stubs** | `internal/assertions/enable/yaml/`, `.../colors/` | No | Holds function pointers, panics when unset |
| **Assertions** | `internal/assertions/*.go` | No | Calls internal stubs; unaware of external libraries |

---

## How It Works: YAML

### The Wiring Chain

```
User imports:  _ "github.com/go-openapi/testify/enable/yaml/v2"
                   │
                   ▼
enable/yaml/enable_yaml.go  init()
  ├─ calls yamlstub.EnableYAMLWithUnmarshal(yaml.Unmarshal)
  └─ calls yamlstub.EnableYAMLWithMarshal(yaml.Marshal)
                   │
                   ▼
enable/stubs/yaml/enable_yaml.go
  └─ delegates to internal/assertions/enable/yaml
                   │
                   ▼
internal/assertions/enable/yaml/enable_yaml.go
  └─ stores function pointers in package-level vars
                   │
                   ▼
internal/assertions/yaml.go
  └─ calls yaml.Unmarshal() / yaml.Marshal() via the stored pointers
```

### What Happens Without Enablement

If a user calls `assert.YAMLEq(t, a, b)` without the blank import, the internal stub panics:

```
panic: YAML is not enabled. To enable YAML support, add a blank import:

  import _ "github.com/go-openapi/testify/enable/yaml/v2"
```

This is intentional: fail fast with a clear fix, rather than silently returning wrong results.

### Internal Stub Pattern (YAML)

The internal stub stores function pointers that start as `nil`:

```go
// internal/assertions/enable/yaml/enable_yaml.go

var (
    enableYAMLUnmarshal func([]byte, any) error
    enableYAMLMarshal   func(any) ([]byte, error)
)

func Unmarshal(in []byte, out any) error {
    if enableYAMLUnmarshal == nil {
        panic("YAML is not enabled...")
    }
    return enableYAMLUnmarshal(in, out)
}
```

### Feature Module (YAML)

The feature module is a separate Go module with its own `go.mod`:

```
// enable/yaml/go.mod
module github.com/go-openapi/testify/enable/yaml/v2

require go.yaml.in/yaml/v3 v3.0.4
```

Its `init()` wires in the real implementation:

```go
// enable/yaml/enable_yaml.go

func init() {
    yamlstub.EnableYAMLWithUnmarshal(yaml.Unmarshal)
    yamlstub.EnableYAMLWithMarshal(yaml.Marshal)
}
```

### Custom YAML Library

Users can bypass the default `enable/yaml` module and inject their own YAML library:

```go
import (
    yaml "github.com/goccy/go-yaml"
    yamlstub "github.com/go-openapi/testify/v2/enable/stubs/yaml"
)

func init() {
    yamlstub.EnableYAMLWithUnmarshal(yaml.Unmarshal)
}
```

This works because the public stubs API accepts any function matching the expected signature.

---

## How It Works: Colorized Output

### The Wiring Chain

```
User imports:  _ "github.com/go-openapi/testify/enable/colors/v2"
                   │
                   ▼
enable/colors/enable.go  init()
  ├─ registers CLI flags: -testify.colorized, -testify.theme, -testify.colorized.notty
  ├─ reads env vars: TESTIFY_COLORIZED, TESTIFY_THEME, TESTIFY_COLORIZED_NOTTY
  ├─ detects TTY via golang.org/x/term
  └─ calls colorstub.Enable(func() []Option { ... })
                   │
                   ▼
enable/stubs/colors/enable_colors.go
  └─ delegates Enable() to internal/assertions/enable/colors
                   │
                   ▼
internal/assertions/enable/colors/enable_colors.go
  └─ stores enabler function, resolves lazily via sync.Once
                   │
                   ▼
internal/assertions/equal.go, diff.go
  └─ calls colors.ExpectedColorizer(), colors.ActualColorizer(), colors.Options()
```

### Lazy Initialization

Colors use a different pattern than YAML: lazy initialization with `sync.Once`.

```go
// internal/assertions/enable/colors/enable_colors.go

var (
    resolveOptionsOnce sync.Once
    optionsEnabler     func() []Option
)

func Enable(enabler func() []Option) {
    optionsEnabler = enabler
}

func resolveOptions() {
    resolveOptionsOnce.Do(func() {
        if optionsEnabler == nil {
            // Not enabled: use no-op colorizers
            return
        }
        // Enabled: build ANSI colorizers from options
        o := optionsWithDefaults(optionsEnabler())
        colorOptions = makeDiffOptions(o)
        stringColorizers = setColorizers(o)
    })
}
```

**Why lazy?** The `init()` function in `enable/colors` registers CLI flags, but flag values
are only available after `flag.Parse()` runs (which happens when the test binary starts).
The `sync.Once` defers resolution until the first assertion call, when flags are ready.

**Why not panic?** Colors are purely cosmetic. If not enabled, assertions work identically
with plain text output. No-op colorizers are used by default.

### Feature Module (Colors)

```
// enable/colors/go.mod
module github.com/go-openapi/testify/enable/colors/v2

require golang.org/x/term v0.39.0
```

The `golang.org/x/term` dependency is used solely for TTY detection (`term.IsTerminal`).

---

## Design Decisions

### Why Three Layers?

A simpler two-layer design (feature module calls internal directly) would work, but three layers provide:

1. **Stable public API** -- `enable/stubs/` is the contract users depend on for custom wiring.
   Internal package paths can change without breaking user code.
2. **Testability** -- stubs can be tested independently of feature modules.
3. **Substitutability** -- users can wire alternative implementations through the stubs API
   without importing the default feature module.

### Panic vs. Silent Degradation

| Feature | Strategy | Reason |
|---------|----------|--------|
| YAML | Panic | Incorrect results are worse than a crash. If `YAMLEq` silently fails, tests pass when they shouldn't. |
| Colors | No-op | Missing colors don't affect correctness. Assertions work fine without ANSI codes. |

### Why Separate Go Modules?

Each `enable/` feature is its own Go module (own `go.mod`). This means:

- `go mod tidy` in the main module never pulls in `go.yaml.in/yaml/v3` or `golang.org/x/term`
- Users who `go get github.com/go-openapi/testify/v2` get zero transitive dependencies
- Feature dependencies are resolved only when the feature module is imported

---

## Adding a New Optional Feature

Follow these steps to add a new opt-in feature (e.g., a hypothetical JSON schema validator):

### 1. Create internal stubs

```go
// internal/assertions/enable/jsonschema/enable.go
package jsonschema

var validateFunc func(schema, document []byte) error

func Validate(schema, document []byte) error {
    if validateFunc == nil {
        panic(`JSON Schema validation is not enabled. Import:
  _ "github.com/go-openapi/testify/enable/jsonschema/v2"`)
    }
    return validateFunc(schema, document)
}

func EnableWithValidate(fn func([]byte, []byte) error) {
    validateFunc = fn
}
```

### 2. Create public stubs

```go
// enable/stubs/jsonschema/enable.go
package jsonschema

import internal "github.com/go-openapi/testify/v2/internal/assertions/enable/jsonschema"

func EnableWithValidate(fn func([]byte, []byte) error) {
    internal.EnableWithValidate(fn)
}
```

### 3. Create feature module

```
// enable/jsonschema/go.mod
module github.com/go-openapi/testify/enable/jsonschema/v2

require github.com/santhosh-tekuri/jsonschema/v6 v6.0.1
```

```go
// enable/jsonschema/enable.go
package jsonschema

import (
    "github.com/santhosh-tekuri/jsonschema/v6"
    stub "github.com/go-openapi/testify/v2/enable/stubs/jsonschema"
)

func init() {
    stub.EnableWithValidate(func(schema, doc []byte) error {
        // wire in the real validator
    })
}
```

### 4. Use in assertions

```go
// internal/assertions/jsonschema.go
package assertions

import jsonschema "github.com/go-openapi/testify/v2/internal/assertions/enable/jsonschema"

func JSONSchemaValid(t T, schema, document []byte, msgAndArgs ...any) bool {
    if h, ok := t.(H); ok {
        h.Helper()
    }
    if err := jsonschema.Validate(schema, document); err != nil {
        return Fail(t, fmt.Sprintf("JSON Schema validation failed: %s", err), msgAndArgs...)
    }
    return true
}
```

### 5. Add integration tests

Create `internal/testintegration/jsonschema/` with tests that import the feature module
and exercise the assertion end-to-end. Update `internal/testintegration/go.mod` to add
the new dependency.

---

## Integration Testing

Optional features require a separate Go module for end-to-end testing, since the main module
cannot import external dependencies.

The `internal/testintegration/` module serves this purpose:

```
internal/testintegration/
├── go.mod           # Imports: yaml, colors, rapid
├── yaml/            # Tests YAMLEq with real YAML parser
├── colors/          # Tests colorized output with ANSI detection
└── spew/            # Property-based testing (unrelated to opt-in features)
```

See `internal/testintegration/README.md` for details on running these tests.

---

## Module Map

{{< mermaid align="center" zoom="true" >}}
graph LR
    main["github.com/go-openapi/testify/v2<br/><b>main module</b><br/><i>zero dependencies</i>"]
    yaml_mod["enable/yaml/v2<br/><i>go.yaml.in/yaml/v3</i>"]
    colors_mod["enable/colors/v2<br/><i>golang.org/x/term</i>"]
    testint["internal/testintegration/v2<br/><i>yaml + colors + rapid</i>"]

    yaml_mod -- "replace =>" --> main
    colors_mod -- "replace =>" --> main
    testint -- "replace =>" --> main

    style main fill:#4a9eff,color:#fff
    style yaml_mod fill:#90ee90,color:#000
    style colors_mod fill:#90ee90,color:#000
    style testint fill:#ffb6c1,color:#000
{{< /mermaid >}}

All feature modules use `replace` directives to point to the local main module during development.
In production, Go module resolution handles versioning automatically.
