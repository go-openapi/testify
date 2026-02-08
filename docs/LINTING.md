# Linting Directives Audit

This document catalogs all `//nolint` directives in the codebase and their justifications.

Last reviewed: 2026-02-07

## Summary

| Category | Count | Actionable |
|---|---|---|
| Internalized third-party (`internal/spew/`) | 17 | No (inherited) |
| Enable pattern (`init()` + globals) | 8 | No (architectural) |
| Generated code (shared test pointer) | 4+1 template | No (generated) |
| Linter false positives | 4 | No (linter limitation) |
| Intentional code patterns | 7 | No (by design) |
| Intentional test environment | 4 | No (test-specific) |
| Scanner test infrastructure | 3 | No (test caching) |
| Inherent type-switch complexity | 2 | No (irreducible) |

All directives are justified. No actionable items remain.

## Internalized Third-Party: `internal/spew/`

The `internal/spew/` package is an internalized copy of `go-spew`. Its 17 nolint directives
are inherited and appropriate for a low-level reflection library:

- `bypass.go:42,71,84` - `gochecknoglobals` - reflect internals, set once during init
- `bypass.go:120` - `gochecknoinits` - validates reflect internals at startup
- `spew.go:80,92,104` - `forbidigo` - public API wrapping `fmt.Print/Printf/Println`
- `config.go:111` - `gochecknoglobals` - global configuration (backward-compatible)
- `config.go:184,196,208` - `forbidigo` - public API wrapping `fmt.Print/Printf/Println`
- `common.go:43` - `gochecknoglobals` - immutable byte literals
- `common.go:82` - `gochecknoglobals` - immutable lookup table
- `common.go:221` - `unparam` - `base` param kept for consistency with `FormatInt`
- `dump.go:31` - `gochecknoglobals` - immutable reflect types and compiled regexps
- `dump.go:216` - `gosec` - uint8 conversion is safe (original type is uint8)
- `format_test.go:175,310` - `dupl` - int/uint test data follows same pattern by design
- `format_test.go:1564,1599` - `dupl` - panic/error test data follows same pattern
- `dump_test.go:215,290` - `dupl` - int/uint test data follows same pattern
- `dump_test.go:1235` - `unconvert` - intentional: proves aliased types don't matter

## Enable Pattern: `init()` + Globals

The "enable" pattern activates optional features via blank imports
(`import _ "github.com/go-openapi/testify/v2/enable/yaml"`).
By design, this requires `init()` functions and package-level globals.

| File | Directive | Reason |
|---|---|---|
| `enable/colors/enable.go:24` | `gochecknoglobals` | CLI flags state |
| `enable/colors/enable.go:32` | `gochecknoinits` | Blank-import activation |
| `enable/yaml/enable_yaml.go:12` | `gochecknoinits` | Blank-import activation |
| `internal/assertions/enable/colors/colors.go:55` | `gochecknoglobals` | Package-level colorizers |
| `internal/assertions/enable/colors/colors.go:68` | `gochecknoglobals` | Package-level printer builders |
| `internal/assertions/enable/colors/enable_colors.go:12` | `gochecknoglobals` | Cross-module feature flag |
| `internal/assertions/enable/yaml/enable_yaml.go:10` | `gochecknoglobals` | Cross-module feature flag |
| `internal/leak/leak.go:22` | `gochecknoinits` | Volatile internal API check |

## Generated Code: Shared Test Pointer

A single template generates a package-level `var` used to share a common pointer across
generated test cases. The 4 output files all originate from one template:

| File | Directive |
|---|---|
| `codegen/.../assertion_test_shared.gotmpl:23` | Template source |
| `assert/assert_assertions_test.go:3210` | `gochecknoglobals` |
| `assert/assert_examples_test.go:1030` | `gochecknoglobals` |
| `require/require_assertions_test.go:2738` | `gochecknoglobals` |
| `require/require_examples_test.go:1031` | `gochecknoglobals` |

## Linter False Positives

| File | Directive | Reason |
|---|---|---|
| `internal/assertions/error.go:273` | `errorlint` | Type switch checks for interfaces, not unwrapping errors |
| `internal/assertions/compare.go:496` | `ireturn` | Generic function returning `V`; linter doesn't understand type parameter |
| `codegen/.../examples-parser/parser_test.go:15` | `gochecknoglobals` | `sync.Once` cache for parsed testdata |
| `codegen/.../comments/extractor_test.go:16` | `gochecknoglobals` | `sync.Once` cache for parsed testdata |

Note: the `thelper` linter was disabled entirely (see History) — it previously accounted
for 12 false positives on test case factories returning `func(*testing.T)`.

## Intentional Code Patterns

These directives suppress warnings on code that is correct and intentional.

| File | Directive | Reason |
|---|---|---|
| `internal/assertions/number.go:482` | `gocritic` | `f != f` is the standard NaN check idiom |
| `internal/assertions/compare.go:516` | `forcetypeassert` | Type guaranteed by prior reflection |
| `internal/assertions/json.go:4` | `dupl` | JSON/YAML implementations intentionally parallel |
| `internal/assertions/yaml.go:4` | `dupl` | JSON/YAML implementations intentionally parallel |
| `internal/assertions/panic_test.go:43` | `nilness` | Deliberate `panic(nil)` edge case test |
| `internal/assertions/mock_test.go:64` | `containedctx` | Context injection for test mocks |
| `internal/assertions/file_test.go:126` | `errcheck` | Best-effort `os.Chmod` restore in `t.Cleanup` |

## Intentional Test Environment

These directives are required because the tests intentionally exercise
locale-specific or escape-sequence behavior.

| File | Directive | Reason |
|---|---|---|
| `enable/colors/assertions_test.go:84` | `staticcheck` | Testing ANSI escape sequences |
| `internal/difflib/options_test.go:63` | `staticcheck` | Testing escape sequences |
| `internal/assertions/object_test.go:127` | `gosmopolitan` | `time.Local` is precisely what's being tested |
| `internal/assertions/compare_impl_test.go:178,179` | `gosmopolitan` | `time.Local` is precisely what's being tested |

## Scanner Test Infrastructure

| File | Directive | Reason |
|---|---|---|
| `codegen/.../examples-parser/parser_test.go:15` | `gochecknoglobals` | `sync.Once` cache for parsed testdata |
| `codegen/.../comments/extractor_test.go:16` | `gochecknoglobals` | `sync.Once` cache for parsed testdata |
| `codegen/.../testdata/examplespkg/examplespkg.go:21` | `unused` | Fixture: verifies unexported symbols are skipped |

## Inherent Type-Switch Complexity

These two test functions contain exhaustive type switches over all Go types
(`int`, `int8`, `int16`, `int32`, `int64`, `float32`, `float64`, `string`, etc.).
Unlike the map-based dispatcher pattern (see `.claude/skills/refactor-tests.md` pattern 15),
there is no dispatch key to route over -- just a flat enumeration of types.
Extracting per-type helpers would create dozens of trivial one-liner functions
with no readability gain.

| File | Directive | Reason |
|---|---|---|
| `internal/assertions/equal_test.go:130` | `gocognit,gocyclo,cyclop` | Big type switch over all Go types |
| `internal/assertions/collection_test.go:272` | `gocognit,gocyclo,cyclop` | Type dispatch over all Go collection types |

## History

- **2026-02-07**: Disabled `thelper` linter entirely. Removed 12 `//nolint:thelper` directives
  across 6 test files (`yaml_test.go`, `json_test.go`, `number_test.go`, `string_test.go`,
  `compare_test.go`). The linter produces persistent false positives on test case factories
  returning `func(*testing.T)` and enforces unwanted naming conventions. Issue reported upstream.
- **2026-02-07**: Removed `//nolint:gocognit,gocyclo,cyclop` from
  `codegen/internal/generator/domains/domains_test.go` by extracting inline check closures
  into named helper functions with a map-based dispatcher
  (see `.claude/skills/refactor-tests.md` pattern 15).
