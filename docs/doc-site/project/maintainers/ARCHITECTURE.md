---
title: "Architecture"
description: Project structure
weight: 2
---

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
- n assertion functions × 8 variants = 608 generated functions (current: n=76)
- Package-level functions (`assert.Equal`, `require.Equal`)
- Format variants (`assert.Equalf`, `require.Equalf`)
- Forward methods (`a.Equal()`, `r.Equal()`)
- Tests for all variants
- Testable examples for godoc
- Documentation for documentation site, organized by domain

**Generated Packages: `assert/` and `require/`**

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
