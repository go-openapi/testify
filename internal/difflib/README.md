# internal/difflib

This is an internalized and modernized copy of [github.com/pmezard/go-difflib](https://github.com/pmezard/go-difflib),
a partial port of Python's difflib module for generating textual diffs.

## Original Source

**Source repository:** github.com/pmezard/go-difflib

**Original license:** BSD 3-Clause License (see [NOTICE](./NOTICE))

**Copyright:** 2013 Patrick Mezard

**Maintenance status:** ⚠️ No longer maintained (archived by author)

`difflib` provides tools to compare sequences of strings and generate textual diffs in unified or context format.

It implements Python's `SequenceMatcher` class and `unified_diff()`/`context_diff()` functions.

## Why Internalized

This fork of testify maintains **zero external dependencies** for the core assertion packages.
By internalizing `difflib`, we:

1. Eliminate the external dependency on an **unmaintained package** (last updated 2014)
2. Gain full control to apply modernizations aligned with our go1.24 target
3. Can apply targeted fixes and optimizations specific to testify's use cases

## Modernizations Applied

This internalized copy has been significantly modernized and refactored from the original go-difflib codebase:

### Go Language Features (Go 1.21+)

- **Built-in functions:** Removed custom `min()` and `max()` functions in favor of Go 1.21+ built-ins
- **Modern operators:** Used `--` instead of `-= 1` for decrement operations
- **Efficient conversions:** Used `strconv.Itoa()` instead of `fmt.Sprintf("%d", ...)` for integer-to-string conversion
- **Buffer handling:** Used `bytes.Buffer.String()` instead of `string(bytes.Buffer.Bytes())`
- Used `new(bytes.Buffer)` instead of `&bytes.Buffer{}`

### Code Organization & Complexity Reduction

- **Function extraction:** Refactored complex functions by extracting helper functions:
  - `writeGroup()` - Handles writing diff groups
  - `writeEqual()` - Writes unchanged lines
  - `writeReplaceOrDelete()` - Writes deleted/replaced lines (prefix `-`)
  - `writeReplaceOrInsert()` - Writes inserted/replaced lines (prefix `+`)
- **Method reorganization:** Reordered methods for better logical flow (public methods first, then helpers)
- **Named constants:** Added named constants for magic numbers (e.g., `hundred = 100`, `maxDisplayElements = 200`)

### Code Quality

- **Godoc compliance:** Updated all function comments to start with the function name for proper godoc generation
  - `// Set two sequences` → `// SetSeqs sets two sequences`
  - `// Return list of triples` → `// GetMatchingBlocks return the list of triples`
- **Linting compliance:** Removed blank identifiers in range loops where value is unused
  - `for s, _ := range` → `for s := range`
- **Modern control flow:** Replaced if-else chains with switch statements for better readability
- **Simplified logic:** Improved boolean expressions using De Morgan's laws
  - `!(len(group) == 1 && group[0].Tag == 'e')` → `(len(group) != 1 || group[0].Tag != 'e')`
- **Struct literals:** Simplified composite literals where types are inferred
  - `OpCode{'e', 0, 1, 0, 1}` instead of `OpCode{Tag: 'e', I1: 0, I2: 1, J1: 0, J2: 1}`

### Documentation

- **Comment punctuation:** Standardized comment formatting and added proper punctuation
- **Code clarity:** Added inline comments for switch cases explaining tag meanings (`'r'`, `'d'`, `'i'`, `'e'`)

## API Compatibility

The internalized copy maintains full API compatibility with the original go-difflib while incorporating the modernizations above. All public functions and types work identically to the upstream version.

Key exports:
- `SequenceMatcher` - Compares sequences of strings using the Ratcliff-Obershelp algorithm
- `UnifiedDiff` / `WriteUnifiedDiff()` / `GetUnifiedDiffString()` - Generate unified diff format
- `ContextDiff` / `WriteContextDiff()` / `GetContextDiffString()` - Generate context diff format
- `SplitLines()` - Split strings on newlines while preserving them

## Use in Testify

This package is used by testify's assertion functions to generate human-readable diffs when assertions fail, particularly for comparing strings, slices, and complex data structures.

The diff output helps developers quickly identify what changed between expected and actual values during test failures.

## Enhancements
- [x] Colorized output support

## Future Enhancements

As an internalized dependency, this copy can receive targeted improvements:

- **Potential:** Performance optimizations for large diffs
- **Potential:** Enhanced diff algorithms for specific data types

## Maintenance

This internalized copy is maintained as part of github.com/go-openapi/testify/v2 and follows the same Go version requirements (currently go1.24). It does not track upstream go-difflib releases, as the original repository is no longer maintained and this copy has diverged through modernization and refactoring.

For issues or improvements specific to this internalized version, please file issues at:
https://github.com/go-openapi/testify/issues

## License

This library ships under the [SPDX-License-Identifier: Apache-2.0](./LICENSE) and acknowledges the original licensing
terms (BSD-3-Clause-License) from the original source it was copied from: [NOTICE](./NOTICE).

The original copyright and license terms are preserved in accordance with the BSD License requirements.
