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

> The generator scans source code, extracts metadata, builds a model, and applies templates to generate ~600+ functions, tests, and documentation from ~70-80+ source functions.

---

### Adding a New Assertion

**Complete workflow:**

1. **Add function to `internal/assertions/<domain>.go`:**

   The following example would like go to `string.go`, next to the `Regexp` assertion.

   ```go
   // StartsWith asserts that the string starts with the given prefix.
   //
   // Examples:
   //
   //   success: "hello world", "hello"
   //   failure: "hello world", "bye"
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

> **76 functions × 8 variants = 608 generated functions** (plus tests and documentation for each)

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

### Generator Flags

```bash
go run ./codegen/main.go \
  -work-dir=.. \
  -input-package=github.com/go-openapi/testify/v2/internal/assertions \
  -output-packages=assert,require \
  -target-root=.. \
  -include-format-funcs=true \
  -include-forward-funcs=true \
  -include-tests=true \
  -include-examples=true \
  -runnable-examples=true \
  -include-helpers=true \
  -include-generics=false
```

Current usage with `go generate` (see `doc.go`):

```go
//go:generate go run ./codegen/main.go -target-root . -work-dir .
```

**Note:** Generic functions are planned but not yet implemented.

### Verification

After generation, verify:
```bash
# All tests pass
go test ./...

# Coverage remains high
go test -cover ./internal/assertions  # Should be ~94%
go test -cover ./assert               # Should be ~99.5%
go test -cover ./require              # Should be ~99.5%

# Examples are valid
go test -run Example ./assert
go test -run Example ./require
```

The 0.5% coverage gap comes from helper functions (non-assertion functions) that don't have "Examples:" annotations.
