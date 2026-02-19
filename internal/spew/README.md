# internal/spew

This is an internalized and modernized copy of [github.com/davecgh/go-spew](https://github.com/davecgh/go-spew),
a deep pretty printer for Go data structures.

## Original Source

**Source repository:** github.com/davecgh/go-spew

**Original license:** ISC License (see [LICENSE](./LICENSE))

**Copyright:** 2012-2016 Dave Collins

---

go-spew implements a deep pretty printer for Go data structures to aid in debugging,
providing features like pointer dereferencing, circular reference detection,
custom Stringer/error interface handling, and hexdump-style byte array output.

## Why Internalized

This fork of testify maintains **zero external dependencies** for its core assertion packages.
By internalizing go-spew, we eliminate the external dependency while gaining full control over the code
to apply modernizations aligned with our go1.24 target.

## Modernizations Applied

This internalized copy has been modernized from the original go-spew codebase with the following changes:

### Go Language Features (Go 1.18+)

- **Type aliases:** Replaced `interface{}` with `any` throughout
- **Modern build tags:** Updated from `// +build` to `//go:build` format
- **Range iteration:** Used `for i := range n` instead of `for i := 0; i < n; i++`
- **reflect.TypeFor:** Replaced `reflect.TypeOf(T(0))` with `reflect.TypeFor[T]()`
- **Standard library improvements:** Used `strings.ReplaceAll()` instead of `strings.Replace(..., -1)`

### Code Quality

- **Linting compliance:** Added linting directives (`//nolint`) with explanations where appropriate
- **Test improvements:** Added `t.Helper()` calls in test helper functions
- **Modern idioms:** Used `slices.Contains()` instead of manual loops
- **String building:** Used `strings.Builder` for efficient string concatenation
- **Code organization:** Improved struct field ordering and switch statement structure with explicit `fallthrough` comments

### Documentation

- **Markdown formatting:** Updated godoc documentation comments to use modern markdown headings (`#`) and list formats (`-`)
- **Comment punctuation:** Standardized comment punctuation and formatting
- **Clarity improvements:** Fixed typos and improved readability

### File Organization

- **Test data relocation:** Moved `testdata/` to `testsrc/` with proper documentation

## Notable Functional Enhancements

The internalized copy maintains API compatibility with the original while incorporating targeted improvements:

- **Deterministic map sorting:** The `SpewKeys` configuration option enables sorted map key output for consistent diffs (relevant for testify's assertion output)
  - [x] Additional optimizations for deterministic diff generation (`stretchr/testify#1822`)
- **time.Time rendering:** Enhanced handling of `time.Time` values in nested structures (applied from `stretchr/testify#1829`)
- **panic/hang** Reinforced input checking to avoid edge cases
  - [x] Proper fix for panic on unexported struct keys in maps (`stretchr/testify#1816`)
  - [x] Fix for circular map references (runtime stack overflow)

## Future Enhancements

The linting effort is still going on...

## Maintenance

This internalized copy is maintained as part of github.com/go-openapi/testify/v2 and follows
the same Go version requirements (currently go1.24).
It does not track upstream go-spew releases, as it has diverged through modernization.

For issues or improvements specific to this internalized version, please file issues at:
https://github.com/go-openapi/testify/issues

## License

This code retains its original ISC License. See [LICENSE](./LICENSE) for the full license text.

The original copyright and license terms are preserved in accordance with the ISC License requirements.
