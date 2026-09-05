---
title: "Architecture"
description: Project structure
weight: 2
---

## Primer

### Goals

We want the maintenance of dozens of test assertions, times many variants, to remain reasonably low.

The maintenance flow is intended to require different activities and levels of understanding,
depending on the complexity of a planned evolution.

Not everything can be hassle-free, but the design should offer a relatively easy path to most common maintenance.

After 9 months of all sorts of maintenance tasks, the diagram below gives a rather faithful representation of what
it actually costs.

{{< mermaid align="center" zoom="true" >}}
quadrantChart
    title Change complexity vs Required knowledge
    x-axis Minimal Knowledge --> In-Depth Understanding
    y-axis Simple Change --> Complex Change

    quadrant-3 Hassle-free
    quadrant-2 Follow the tracks
    quadrant-4 Should stay empty
    quadrant-1 Generator work

    Bug fixes: [0.2, 0.2]
    Doc fixes: [0.25, 0.1]
    Minor enhancements: [0.3, 0.4]
    New assertions: [0.35, 0.23]
    New dependencies: [0.20, 0.70]
    New constructs: [0.78, 0.78]
    Guarded assertions: [0.40, 0.58]
    Package layout: [0.68, 0.62]

{{< /mermaid >}}

The bottom-right quadrant is empty on purpose: nothing here asks for in-depth knowledge of the
generator to make a small change. A point landing there means the code needs reshaping.

Most common maintenance tasks should not require much more than fixing/enhancing the code in `internal/assertions`.

Fixes and enhancements propagate naturally to the variants without the need to regenerate code.

API changes need an extra code generation, but no specific knowledge of the generator itself.

Dependency changes (adding new features that need extra dependencies) is a bit more involved, but still manageable:
the pattern is regular, follow the tracks.

The code & doc generator has now become a rather stable component. The maintenance of the generator itself remains
an operation that requires an extended understanding of the internals of the project.

Example of recent updates that required such in-depth maintenance: adding support for build guards.

### The maths with assertion variants

Each test assertion produces 2 base variants (assert, require).

Each of these variants produces another formatted variant, plus one "forward" variant and one
"forward formatted" variant (as methods of the `Assertions` type).

**For every assertion: 8 variants.**

**For every "helper" function (not an assertion): 2 variants.**

> Generic assertions reach 8 variants only from go1.27, the first release that accepts type parameters on methods.
>
> Their 4 method variants are generated into files guarded by `//go:build go1.27`
> (`assert/assert_forward_go127.go`, `require/require_forward_go127.go`), so a go1.26 build
> drops them and keeps the 4 package-level variants. The counts below assume go1.27.


All these variants make up several hundreds functions, which poses a challenge for maintenance and documentation.

We have adopted code and documentation generation as a means to mitigate this issue.

#### Current

 1. Generic assertions (with type parameters): {{% siteparam "metrics.generics" %}} functions
 2. Non-generic assertions (with t T parameter, no type parameters): {{% siteparam "metrics.nongeneric_assertions" %}} functions
 3. Helper functions (no t T parameter): {{% siteparam "metrics.helpers" %}} functions

 Total: {{% siteparam "metrics.functions" %}} functions to _maintain_

 **Generated Functions**

 1. Generated variants in each package (assert/require): {{% siteparam "metrics.package_variants" %}}
 2. Helpers:  {{% siteparam "metrics.helpers" %}}
 3. Constructors: 2 (1 in assert, 1 in require)

 Overall: {{% siteparam "metrics.total_functions" %}} generated functions

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
* ad hoc testable examples are not generated

**Optional Feature Packages: `enable/`**

The `enable/` package provides optional features that users can activate via blank imports:
- `enable/stubs/` - Public stub APIs for enabling features (yaml, colors)
- `enable/yaml/` - Activates YAML support via `import _ "github.com/go-openapi/testify/enable/yaml/v2"`
- `enable/colors/` - Activates colorized output via `import _ "github.com/go-openapi/testify/enable/colors/v2"`

These packages are not generated and allow optional dependencies to be isolated from the core library.

## See Also

- [Code generation](./CODEGEN.md) - Detailed view of our code and doc generator
