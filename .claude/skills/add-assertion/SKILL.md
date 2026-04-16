# Adding a New Assertion

Step-by-step workflow for adding a new assertion function to testify.

## Workflow

1. Add function to the appropriate domain file in `internal/assertions/`
2. Add `// Domain: <domain>` as the first line inside the function body
3. Add `// Opposite: <Name>` on the next line if a logical opposite exists
4. Add `Examples:` section to the doc comment
5. Add tests to the corresponding `*_test.go` file
6. Run `go generate ./...` to produce all 8 variants + docs
7. Run `go test work ./...` to verify everything

## Function template

```go
// FuncName asserts that <what it does>.
//
// # Usage
//
//	assertions.FuncName(t, arg1, arg2)
//
// # Examples
//
//	success: arg1Value, arg2Value
//	failure: arg1Value, arg2Value
func FuncName(t T, arg1, arg2 any, msgAndArgs ...any) bool {
	// Domain: <domain>
	// Opposite: <OppositeName>  (omit if none)
	if h, ok := t.(H); ok {
		h.Helper()
	}

	// implementation
	if !condition {
		return Fail(t, "message", msgAndArgs...)
	}

	return true
}
```

## Doc comment annotations

### Domain tag (in doc comment header)

```go
// domain: equality
```

Assigns the function to a documentation domain. Add domain descriptions in
`internal/assertions/doc.go` if creating a new domain.

### Domain comment (inside function body)

```go
// Domain: equality
```

First line inside the function body. Used by the codegen scanner.

### Opposite annotation

```go
// Opposite: NotEqual
```

Second line inside the body (after Domain). Only on the affirmative form
(e.g., on `Equal`, not on `NotEqual`).

### Examples section

```go
// # Examples
//
//	success: 123, 123
//	failure: 123, 456
```

Drives generated smoke tests for all 8 variants. Three case types:
- `success: <args>` -- test should pass
- `failure: <args>` -- test should fail
- `panic: <args>` followed by `<expected panic message>` on next line

For complex values that can't be represented inline, use `// NOT IMPLEMENTED`:
```go
//	success: &customStruct{Field: "value"}, // NOT IMPLEMENTED
```

**Never use `// TODO`** -- it triggers false positives in code quality tools.

## What gets generated

From a single function, `go generate` produces:

| Variant | Package | Example |
|---------|---------|---------|
| Package-level | `assert` | `assert.FuncName(t, ...)` |
| Formatted | `assert` | `assert.FuncNamef(t, ..., "msg")` |
| Forward method | `assert` | `a.FuncName(...)` |
| Forward formatted | `assert` | `a.FuncNamef(..., "msg")` |
| Package-level | `require` | `require.FuncName(t, ...)` |
| Formatted | `require` | `require.FuncNamef(t, ..., "msg")` |
| Forward method | `require` | `r.FuncName(...)` |
| Forward formatted | `require` | `r.FuncNamef(..., "msg")` |

Plus: tests for all variants, documentation entry in `docs/doc-site/api/<domain>.md`.

Generic assertions (with type params) produce 4 variants (no forward methods --
Go limitation).

## Checklist

- [ ] Function in `internal/assertions/<domain>.go`
- [ ] `// Domain:` and optionally `// Opposite:` inside function body
- [ ] Doc comment with `# Usage`, `# Examples` sections
- [ ] Tests in `internal/assertions/<domain>_test.go`
- [ ] `go generate ./...` succeeds
- [ ] `go test work ./...` passes
- [ ] `golangci-lint run --new-from-rev master` clean
