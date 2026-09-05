---
title: Code Generation
description: Code generation workflow and maintenance.
weight: 3
---

{{% notice primary "TL;DR" "meteor" %}}
> The entire assert/require API (600+ functions) is generated from 76 source assertions in `internal/assertions/`.
> Run `go generate ./...` to regenerate everything. Add new assertions by editing source files and adding examples.
{{% /notice %}}

## Maintaining Generated Code

This repository uses code generation extensively to maintain consistency across assertion packages.

### Code Generation Pipeline

{{< mermaid align="center" zoom="true" >}}
  graph TD
    source["📦 internal/assertions/*.go"]
    scanner["🔍 Scanner
    go/packages + go/types"]
    model["fa:fa-database
    Model data structures"]
    templates["📝 Templates
    Go text/template"]
    outputs["📤 Generated Code"]

    source --> scanner
    scanner --> extract_metadata
    extract_metadata --> model
    model --> templates
    templates --> outputs

    subgraph extract_metadata["Extract Metadata"]
      direction BT
      extract["Extractor"]
      comments["godoc comments"] --o extract
      examples["examples: values comments"] --o extract
      domains["domain tags"] --o extract
      sigs["Function signatures"] --o extract
      sigs["Other internal annotations comments"] --o extract
    end

    outputs -.-> assert_package
    outputs -.-> require_package
    outputs -.-> docs@{shape: documents, label: "docs/doc-site/**/*.md"}
    
    subgraph assert_package
      direction BT
      assert@{shape: documents, label: "assert/*.go"}
      tests_assert["*_test.go files"] --o assert
      example_tests_assert["*_examples_test.go files"] --o assert

      subgraph not_generated_assert["*not generated*"]
        direction LR
        docgo_assert@{ shape: document, label: "doc.go" }
        adhoc_assert@{ shape: document, label: "*_adhoc*_test.go" }
      end
    end

    subgraph require_package
      direction BT

      require@{shape: documents, label: "require/*.go"}
      tests_require["*_test.go files"] --o require
      example_tests_require["*_examples_test.go files"] --o require

      subgraph not_generated_require["*not generated*"]
        direction LR
        docgo_require@{ shape: document, label: "doc.go" }
        adhoc_require@{ shape: document, label: "*_adhoc*_test.go" }
      end
    end

    style not_generated_assert fill:#4a9eff,color:#fff
    style not_generated_require fill:#4a9eff,color:#fff
{{< /mermaid >}}

> The generator scans source code, extracts metadata, builds a model, and applies templates to generate ~800+ functions, tests, and documentation from ~100+ source functions.

---

### Adding a New Assertion

**Complete workflow:**

1. **Add function to `internal/assertions/<domain>.go`:**

   The following example would like go to `string.go`, next to the `Regexp` assertion.

   ```go
   import (
   	"fmt"
   	"strings"
   )
   
   // StartsWith asserts that the string starts with the given prefix.
   //
   // Examples:
   //
   //	success: "hello world", "hello"
   //	failure: "hello world", "bye"
   func StartsWith(t T, str, prefix string, msgAndArgs ...any) bool {
   	if h, ok := t.(H); ok {
   		h.Helper()
   	}
   	if !strings.HasPrefix(str, prefix) {
   		return Fail(t, fmt.Sprintf("Expected %q to start with %q", str, prefix), msgAndArgs...)
   	}
   	return true
   }
   ```

2. **Add tests to `internal/assertions/<domain>_test.go`:**
   Write comprehensive table-driven tests covering edge cases.

3. **Run code generation:**
   ```bash
   go generate ./...
   ```

4. **Done!** All 8 variants are generated with tests and examples:
   - `assert.StartsWith(t, str, prefix)`
   - `assert.StartsWithf(t, str, prefix, "msg")`
   - `a.StartsWith(str, prefix)` (forward method)
   - `a.StartsWithf(str, prefix, "msg")`
   - `require.StartsWith(t, str, prefix)`
   - `require.StartsWithf(t, str, prefix, "msg")`
   - `r.StartsWith(str, prefix)` (forward method)
   - `r.StartsWithf(str, prefix, "msg")`

#### How One Function Becomes Eight

{{< mermaid align="center" zoom="true" >}}
graph TD
    source["1 Source Function
    internal/assertions/Equal()"]

    source --> assert_group["assert Package"]
    source --> require_group["require Package"]

    assert_group --> assert_pkg["assert.Equal(t, a, b)
    package-level"]
    assert_group --> assert_fmt["assert.Equalf(t, a, b, msg)
    formatted variant"]
    assert_group --> assert_fwd["a.Equal(a, b)
    forward method"]
    assert_group --> assert_fwdfmt["a.Equalf(a, b, msg)
    forward + formatted"]

    require_group --> require_pkg["require.Equal(t, a, b)
    package-level (fatal)"]
    require_group --> require_fmt["require.Equalf(t, a, b, msg)
    formatted variant (fatal)"]
    require_group --> require_fwd["r.Equal(a, b)
    forward method (fatal)"]
    require_group --> require_fwdfmt["r.Equalf(a, b, msg)
    forward + formatted (fatal)"]

    style source fill:#4a9eff,color:#fff
    style assert_group fill:#90ee90,color:#000
    style require_group fill:#ffb6c1,color:#000
{{< /mermaid >}}

> **reflection-based assertions become 8, generic assertions become 4 + 4 more available on go1.27)**
>
> (plus tests and documentation for each).

---

### Example Annotations Format

The "Examples:" section in doc comments drives test and example generation:

```go
// Examples:
//
//   success: <test arguments that should succeed>
//   failure: <test arguments that should fail>
//   panic: <test arguments that cause panic>
//          <expected panic message>
```

**Rules:**
- Use valid Go expressions that can be directly inserted into test code
- `success:` and `failure:` are required for most assertions
- `panic:` is optional (used for assertions like Panics, YAMLEq)
- Multiple examples of the same type are allowed (e.g., multiple `success:` lines)
- Examples are extracted by the scanner and used to generate:
  - Unit tests (success + failure cases)
  - Testable examples (success cases only for simplicity)

**Example with multiple success cases:**
```go
// Examples:
//
//   success: []string{"a", "b"}, 2        // slice
//   success: map[string]int{"a": 1}, 1    // map
//   success: "hello", 5                    // string
//   failure: []string{"a"}, 5
```

#### Example-Driven Test Generation

{{< mermaid align="center" zoom="true" >}}
  graph LR
    doccomment["Doc Comment
    with Examples:"]
    parser["📖 Example Parser"]
    cases["Test Cases
    success/failure/panic"]
    multiplier["Multiply × 8"]
    tests["Generated Tests"]

    doccomment --> parser
    parser --> cases
    cases --> multiplier

    multiplier --> pkg_assert["assert package test"]
    multiplier --> fmt_assert["assert format test"]
    multiplier --> fwd_assert["assert forward test"]
    multiplier --> fwdfmt_assert["assert fwd+fmt test"]
    multiplier --> pkg_require["require package test"]
    multiplier --> fmt_require["require format test"]
    multiplier --> fwd_require["require forward test"]
    multiplier --> fwdfmt_require["require fwd+fmt test"]

    pkg_assert & fmt_assert & fwd_assert & fwdfmt_assert & pkg_require & fmt_require & fwd_require & fwdfmt_require --> tests

    style cases fill:orange,color:black;
    style multiplier fill:yellow,color:black;
    style tests fill:lightgreen,color:black;
{{< /mermaid >}}

> Each example in doc comments generates 8 test functions (one per variant), ensuring 100% test coverage of generated code.
> In addition, the generator produces testable examples (somewhat redundant with "passed" tests) so every function gets
> a testable example displayed on pkg.go.dev.

---

### Special Cases in Generated Tests

For complex assertions requiring special setup, the test templates support conditional logic. See `codegen/internal/generator/templates/assertion_assertions_test.gotmpl` for examples of:
- Custom mock selection based on function behavior (mockT vs mockFailNowT)
- Package-specific test helpers (testDataPath, httpOK, etc.)
- Handling functions without test examples (generates `t.Skip()`)

Some go expressions won't fit nicely for examples (examples use an external package, e.g. `assert_test`).
To cover these edge cases, a `relocate` function map currently rewrites the example values to be used
from an external package. The relocation uses go parsing capabilities. The only hard-coded exception if for `PanicFunc`.
(see `codegen/internal/generator/funcmap.go`).

---

### Adding an assertion with go version gate

Some assertions need a standard library function that only exists from a given Go release on.
`ErrorAsType` was the first: it wraps `errors.AsType`, added in go1.26, at a time when the
library still supported go1.25.

Build constraints in Go apply to a whole file, so such an assertion gets its own source file,
and the generator replicates that file's `//go:build` line onto a parallel set of generated
files. Users on an older toolchain never see the file: it drops out before the compiler reads it.

The walkthrough below follows what we did for `ErrorAsType` in v2.6. The gate is gone since
v2.7 raised the floor to go1.26 (PR #164), so read the recipe here rather than in the tree.

**1. Write the assertion in its own guarded file.**

Name it `internal/assertions/<domain>_go1NN.go` — `error_go126.go` for `ErrorAsType` — and put
the constraint between the SPDX header and the package clause, with a blank line on each side:

```go
// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build go1.26

package assertions
```

Write `go1.26`, not `1.26`; the latter parses as a plain tag name and gates nothing.

Everything else stays as in [Adding a New Assertion](#adding-a-new-assertion): the same
`// Domain: error` tag, the same `Examples:` block. The domain tag decides which page the
assertion documents on, so a guarded assertion lands next to its unguarded siblings.
Say in the doc comment that the assertion needs the newer toolchain — the generated variants
copy the comment verbatim to `pkg.go.dev`, where nothing else marks the constraint.

**2. Regenerate on a toolchain that can see the guard.**

The generator never reads the guarded file itself. `packages.Load` shells out to `go list` in
the work-dir, and that subprocess decides which files exist. Run it on a toolchain older than
the guard and the file is simply absent: the assertion disappears from `assert/`, from
`require/` and from the doc site, with no error and nothing in the diff to catch the eye.

Two things that look like they would protect you, and do not:

- **The toolchain you built the generator with.** The binary carries no toolchain. Build it
  with go1.27 and run it with go1.26 on `PATH` and the scan still sees a go1.26 view.
- **The `toolchain` directive in `codegen/go.mod`.** A module's directive is ignored in
  workspace mode, and the scan's `go list` runs in the repository root with the workspace
  active. That directive covers standalone use — `go run github.com/go-openapi/testify/codegen/v2@latest`,
  or `GOWORK=off` inside `codegen/` — and nothing else.

So name the toolchain on the regeneration command itself:

```bash
cd codegen
go build -o codegen .
GOTOOLCHAIN=go1.26.0 ./codegen
```

**Do not raise the `go.work` toolchain line to cover a guard.** That line applies to every
command run in the workspace, and CI runs the whole `oldstable`/`stable` matrix through
`go test work ./...`. Pin the workspace and both matrix entries build the same toolchain, so
the older one stops exercising the exclusion path — which is the one property the guard exists
to protect.

{{% notice tip "The generator refuses an incomplete scan" "shield" %}}
> You do not have to remember this. `verifyGuardedFilesLoaded` (`codegen/internal/scanner/buildtags.go`)
> reads the package directory textually, compares the `//go:build go1.N` files it finds there against
> the files the load actually returned, and stops the run when one is missing:
>
> ```
> guarded source file missing from the scan: error_go127.go (//go:build go1.27). The go command
> running the scan is older than these guards ... Rerun with GOTOOLCHAIN=go1.27.0
> ```
>
> The scan on disk has to be textual: a file guarded above the running toolchain never appears in the
> typed package, so the typed view cannot report what it is missing. Nothing is generated before the
> check passes. Only a plain `go1.N` constraint counts — a file selected on an OS or an architecture is
> absent on some machines by design.
{{% /notice %}}

**3. Add the hand-written tests that examples cannot express.**

`Examples:` still drives the generated tests for all variants. Anything they cannot reach goes
in `internal/assertions/<domain>_go1NN_test.go`, carrying the same `//go:build` line. For
`ErrorAsType`, `error_go126_test.go` covered the nil-target form, where `E` cannot be inferred
from the argument and has to be written out: `ErrorAsType[*customError](mock, err, nil)`.

**4. Regenerate.**

```bash
go generate ./...
```

No change to the generator itself. Each category of generated file gains a guarded twin,
suffixed with the constraint (`model.GoBuildTag` turns `go1.26` into `go126`):

| Default file | Guarded twin |
|--------------|--------------|
| `assert/assert_assertions.go` | `assert/assert_assertions_go126.go` |
| `assert/assert_format.go` | `assert/assert_format_go126.go` |
| `assert/assert_assertions_test.go` | `assert/assert_assertions_go126_test.go` |
| `assert/assert_format_test.go` | `assert/assert_format_go126_test.go` |
| `assert/assert_examples_test.go` | `assert/assert_examples_go126_test.go` |

`require/` gets the same five. Package boilerplate is not duplicated: the `Assertions` type,
`New`, `forwardArgs`, the mock types and the shared test fixtures stay in the default files,
gated in the templates behind `{{ if not .BuildConstraint }}`, and the guarded files refer to
them. A category with nothing to declare produces no file at all.

**The forward methods are keyed on a second constraint.** A generic assertion becomes a method
only from go1.27 on, whatever guard its source carries, so `ErrorAsType`'s methods went to
`assert_forward_go127.go` — `Function.ForwardGoBuild` returns the higher of the source guard
and go1.27 for a generic function, and `Generate` partitions the forward files on that instead
of on `GoBuild`. This is why the go1.26 batch above has no `assert_forward_go126.go`.

On the doc site the badges are automatic. The domain page prints
{{% goversion "go1.26" %}} next to the assertion heading and adds a line to the legend, and
`metrics.md` badges the row. Nothing to hand-edit.

**5. Removing the gate, later.**

When the supported floor catches up, move the functions into the plain domain file, delete the
two `//go:build` lines, fold the hand-written tests back into `<domain>_test.go`, and
regenerate. `sweepOrphanVariants` deletes the ten `*_go126*.go` files this run no longer
produces, logging each removal; it only touches files matching `<pkg>_*_go<N>[_test].go` that
carry our "DO NOT EDIT" marker, so a hand-authored file of the same shape survives.

> `go generate ./...` snapshots each package's file list before running the directives, so the
> run that removes an orphan can end on a harmless "no such file". Rerun it, or call
> `go run ./codegen/main.go` directly.

Also re-enable `TestBuildConstraintDetection` in `codegen/internal/scanner/buildtags_test.go`
when a guard comes back, and skip it again when the last one goes: it asserts that the scanner
attaches the constraint to the guarded function, and needs a guarded assertion to look at.

---

### Regenerating

Build the generator and run it from `codegen/`. Every flag defaults to what this repository
needs — `-work-dir ..` to scan, `-target-root ..` to write — so the bare command is the whole
procedure:

```bash
cd codegen
go build -o codegen .
./codegen
```

`go generate ./...` from the repository root does the same through the directive in `doc.go`.
It is the shorter form, and the one to avoid when a run deletes a generated file: `go generate`
snapshots each package's file list before it runs the directives, so it can end on a spurious
"no such file" for the file just removed (see [Removing the gate](#adding-an-assertion-with-go-version-gate)).

**Which Go runs the scan matters.** The generator reads `internal/assertions` through
`packages.Load`, which shells out to `go list`; that subprocess uses the `go` on `PATH`, in the
repository root, with the workspace active. The toolchain that compiled the generator has no
say in it. This only bites when `internal/assertions` carries a `//go:build go1.N` file — see
[Adding an assertion with go version gate](#adding-an-assertion-with-go-version-gate) — and the
fix is to name the toolchain on the command: `GOTOOLCHAIN=go1.N ./codegen`.

Then check the result:

```bash
git diff --stat        # generated files only; nothing under internal/assertions
go test ./...
```

---

### Generator Flags

> Defaults are defined so running the command is essentially hassle-free to maintain this project: no arguments needed.

```bash
go run ./codegen/main.go \
	-work-dir=.. \
	-input-package=github.com/go-openapi/testify/v2/internal/assertions \
	-output-packages=assert,require \
	-target-root=.. \
	-target-doc=docs/doc-site/api \
	-include-format-funcs=true \
	-include-forward-funcs=true \
	-include-tests=true \
	-include-generics=true \
	-include-helpers=true \
	-include-examples=true \
	-runnable-examples=true \
	-include-doc=true
```

Current usage with `go generate` (see `doc.go`):

```go
//go:generate go run ./codegen/main.go -target-root . -work-dir .
```

Pass `-include-doc=false` to regenerate the code alone and leave `docs/doc-site/api` untouched.

### Verification

After generation, verify:
```bash
# All tests pass
go test ./...

# Coverage remains high
go test -cover ./internal/assertions # Should be ~94%
go test -cover ./assert              # Should be ~99.5%
go test -cover ./require             # Should be ~99.5%

# Examples are valid
go test -run Example ./assert
go test -run Example ./require
```

The 0.5% coverage gap comes from helper functions (non-assertion functions) that don't have "Examples:" annotations.
