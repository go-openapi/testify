---
title: "Architecture"
description: Project structure
weight: 2
---

## Primer

### Goals

We want the maintainance of dozens of test assertions, times many variants, to remain reasonably low.

The maintainance flow is intended to require different activities and levels of understanding,
dependending on the complexity of a planned evolution.

{{< mermaid align="center" zoom="true" >}}
journey
    section Fixes & minor enhancements
      internal/assertions:5: Knowledge of the functionality
    section New dependencies
      internal/assertions/enable/...:5: Understanding of the repo architecture
      enable/...:5:  Understanding of the repo architecture
    section API changes
      regenerate code:5: No specific knowledge
    section New constructs to support
      code & doc generator:5: Knowledge of internals
{{< /mermaid >}}

Most common maintenance tasks should not require much more than fixing/enhancing the code in `internal/assertions`.

API changes need an extra code generation.

Dependency changes (adding new features that need extra dependencies) is a bit more involved, but still manageable.

The code & doc generator should rapidly become a very stable component. The maintenance of the generator itself remains
an operation that requires an extended understanding of the internals of the project.

Fixes and enhancements propagate naturally to the variants without the need to regenerate code.

### The maths with assertion variants

Each test assertion produces 2 base variants (assert, require).

Each of these variants produces another formatted variant. Except for generic assertions, we produce
one "forward" variant and one "forward formatted" variant (as methods).

**For every non-generic assertion: 8 variants.**

**For every generic assertion: 4 variants.**

**For every "helper" function (not an assertion): 2 variants.**


All these variants make up several hundreds functions, which poses a challenge for maintainance and documentation.

We have adopted code and documentation generation as a mean to mitigate this issue.

#### Current (v2.3.0-unreleased)

 1. Generic assertions (with type parameters): 42 functions
 2. Non-generic assertions (with t T parameter, no type parameters): 82 functions
 3. Helper functions (no t T parameter): 4 functions

 Total: 128 functions to _maintain_

 **Generated Functions**

 1. Generic assertions: 168
 2. Non-generic assertions: 656
 3. Helper functions: 8
 4. Constructors: 2

 Total: 834 functions

## Architecture Overview

{{< mermaid align="center" zoom="true" >}}
graph LR;
  classDef event font-size:small,font-family:Monospace;
  trigger@{ shape: rounded, label: "go generate" }
  codegen[["🛠️ codegen"]]
  docs@{ shape: documents, label: "📚 API docs"}
  source[["📦 internal/assertions
    (source of truth)"]]

  trigger:::event -.-> codegen
  source --> codegen --> assert
  source --> codegen --> require
  codegen --> docs
{{< /mermaid >}}

**Single Source of Truth: `internal/assertions/`**

All assertion implementations live in `internal/assertions/`, organized by domain:
- Functions are implemented once with comprehensive tests
- Doc comments include "Examples:" sections that drive test generation (including testable examples)
- Both `assert/` and `require/` packages are 100% generated from this source

**Code Generator: `codegen/`**

The generator scans `internal/assertions/` and produces:
- Package-level functions (`assert.Equal`, `require.Equal`)
- Format variants (`assert.Equalf`, `require.Equalf`)
- Forward methods (`a.Equal()`, `r.Equal()`)
- Tests for all variants
- Testable examples for godoc
- Documentation for documentation site, organized by domain

**Generated Packages: `assert/` and `require/`**

The generated functions directly call the internal implementation: no code duplication or change in semantics.

**Generated Documentation: `docs/doc-site/api/`**

Everything in these packages is generated. Never edit generated files directly.

Exceptions:
* `doc.go` is not generated
* ad'hoc testable examples are not generated

**Optional Feature Packages: `enable/`**

The `enable/` package provides optional features that users can activate via blank imports:
- `enable/stubs/` - Public stub APIs for enabling features (yaml, colors)
- `enable/yaml/` - Activates YAML support via `import _ "github.com/go-openapi/testify/v2/enable/yaml"`
- `enable/colors/` - Activates colorized output via `import _ "github.com/go-openapi/testify/v2/enable/colors"`

These packages are not generated and allow optional dependencies to be isolated from the core library.

## See Also

- [Code generation](./CODEGEN.md) - Detailed view of our code and doc generator
