---
title: "Changes from v1"
description: "All changes from testify/v1"
weight: 15
---

## Summary

**Key Changes:**
- ✅ **Zero Dependencies**: Completely self-contained
- ✅ **New functions**: 58 additional assertions (43 generic + 15 reflection-based)
- ✅ **Performance**: ~10x for generic variants (from 1.2x to 81x, your mileage may vary)
- ✅ **Breaking changes**: Requires go1.24, removed suites, mocks, http tooling, and deprecated functions. YAMLEq becomes optional (panics by default).

---

**Testify v2 represents a comprehensive modernization**

- ✅ **Type Safety**: generic assertions catch errors at compile time
- ✅ **Documentation**: compelling documentation site to search the API by use-case domain
- ✅ **Maintainability**: 100% code generation from single source
- ✅ **Quality**: 96% test coverage, use unified test scenarios, extensive fuzzing & benchmarking

This fork maintains compatibility where possible while making bold improvements in architecture, safety, and performance.

The original philosophy of `testify` is preserved, and the new API is 90% compatible.

**Fork Information:**
- **Upstream repository**: [github.com/stretchr/testify](https://github.com/stretchr/testify)
- **Fork date**: 2025-01-09
- **Fork commit**: `feb1324bc3d000fed7b21dfe20bec72ecca27502`

See also a quick [migration guide](./MIGRATION.md).

## Cross-Domain Changes

{{% tabs %}}
{{% tab title="Additions" color=green %}}
### Major Additions

#### Usage

| Change | Origin | Description |
|--------|--------|-------------|
| **Generic assertions** | Multiple upstream proposals | Added 38 type-safe assertion functions with `T` suffix across 10 domains |
| **Zero dependencies** | Design goal | Internalized go-spew and difflib; removed all external dependencies |
| **Optional YAML support** | Design goal | YAML assertions are now enabled via opt-in `enable/yaml` module |
| **Colorized output** | [#1467], [#1480], [#1232], [#994] | Optional colorization via `enable/color` module with themes |
| **Enhanced diff output** | [#1829] | Improved time.Time rendering, deterministic map ordering |

[#1829]: https://github.com/stretchr/testify/issues/1829

#### Maintenability

| Change | Origin | Description |
|--------|--------|-------------|
| **Code generation** | Design goal | 100% generated assert/require packages (840 functions from 127 assertions) |
| **Code modernization** | Design goal | Relinted, refactored and modernized the code base, including internalized difflib and go-spew|
| **Refactored tests** | Design goal | Full refactoring of tests on assertion functions, with unified test scenarios for reflection-based/generic assertions |

[#1232]: https://github.com/stretchr/testify/pull/1232
[#1467]: https://github.com/stretchr/testify/pull/1467
[#1480]: https://github.com/stretchr/testify/pull/1480
[#1829]: https://github.com/stretchr/testify/issues/1829
[#994]: https://github.com/stretchr/testify/pull/994
{{% /tab %}}
{{% tab title="Removals" style=warning %}}

### Major Removals (Breaking Changes)

| Removed | Reason |
|---------|--------|
| **Suite package** | Complex interactions with dependencies; might re-introduce this feature later  |
| **Mock package** | Use specialized [mockery](https://github.com/vektra/mockery) tool instead |
| **HTTP package** | Simplified focus; may be reintroduced later |
| **Deprecated functions** | Clean slate for v2 |
| **Renaming** | `NoDirExists` renamed into `DirNotExists`. `NoFileExists` renamed into `FileNotExists`|

### Infrastructure Improvements

| Change | Description |
|--------|-------------|
| **Internalized dependencies** | go-spew and difflib internalized with modernized code |
| **Module structure** | Clean separation: core (zero deps), enable modules (optional) |
| **Documentation site** | Hugo-based site with domain-organized API reference |
| **Fuzz testing** | Fuzz test on spew.Sdump based on random data structures generation |
| **Comprehensive benchmarks** | 37 benchmarks comparing generic vs reflection performance |
| **Advanced CI** | Reuse go-openapi workflows with tests and coverage reporting, fuzz testing, release automation |
{{% /tab %}}
{{% /tabs %}}

## Bug Fixes and Safety Improvements

{{% tabs %}}
{{% tab title="Bug fixes" color=green %}}

### Critical Fixes reported upstream

| Issue/PR | Domain | Description |
|----------|--------|-------------|
| [#1223] | Display | Display uint values in decimal instead of hex |
| [#1611] | Condition | Fixed goroutine leak in Eventually/Never |
| [#1813] | Internal (spew) | Fixed panic with unexported fields (via #1828) |
| [#1818] | String | Fixed panic on invalid regex in Regexp/NotRegexp |
| [#1822] | Internal (spew) | Deterministic map ordering in diffs |
| [#1825] | Equality | Fixed panic when using EqualValues with uncomparable types |
| [#1828] | Internal (spew) | Fixed panic with unexported fields in maps |

[#1223]: https://github.com/stretchr/testify/pull/1223
[#1611]: https://github.com/stretchr/testify/issues/1611
[#1813]: https://github.com/stretchr/testify/issues/1813
[#1818]: https://github.com/stretchr/testify/pull/1818
[#1822]: https://github.com/stretchr/testify/issues/1822
[#1825]: https://github.com/stretchr/testify/pull/1825
[#1828]: https://github.com/stretchr/testify/pull/1828
{{% /tab %}}
{{% tab title="Safety Improvements" color=blue %}}

### Comprehensive Spew Testing

- Added property-based fuzzing for go-spew with random type generator
- Fixed circular reference edge cases (pointer wrapped in interface, circular map reference)
- Supersedes upstream [#1824]

[#1824]: https://github.com/stretchr/testify/pull/1824

### Reflection Safety

- More defensive guards re-reflect panic risk in `EqualExportedValues`
- Fixed 50 unchecked type assertions across test codebase
- Zero linting issues with `forcetypeassert` linter
{{% /tab %}}
{{% /tabs %}}

## Changes by Domain
### Boolean

{{% expand title="Generics" %}}

#### New Generic Functions (2)

| Function | Type | Origin | Description |
|----------|------|--------|-------------|
| `TrueT[B ~bool]` | Generic | Generics initiative | Type-safe boolean true assertion |
| `FalseT[B ~bool]` | Generic | Generics initiative | Type-safe boolean false assertion |

{{% /expand %}}

**Behavior changes**: None

### Collection

{{% expand title="Generics" %}}

#### New Generic Functions (12)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `StringContainsT[S Text]` | String or []byte | Type-safe string/bytes contains check |
| `StringNotContainsT[S Text]` | String or []byte | Type-safe string/bytes not-contains check |
| `SliceContainsT[E comparable]` | Comparable element | Type-safe slice membership check |
| `SliceNotContainsT[E comparable]` | Comparable element | Type-safe slice non-membership check |
| `MapContainsT[K comparable, V any]` | Key type | Type-safe map key check |
| `MapNotContainsT[K comparable, V any]` | Key type | Type-safe map key absence check |
| `SeqContainsT[E comparable]` | Iterator element | Type-safe iterator membership check (Go 1.23+) |
| `SeqNotContainsT[E comparable]` | Iterator element | Type-safe iterator non-membership check (Go 1.23+) |
| `ElementsMatchT[E comparable]` | Slice element | Type-safe slice equality (any order) |
| `NotElementsMatchT[E comparable]` | Slice element | Type-safe slice inequality |
| `SliceSubsetT[E comparable]` | Slice element | Type-safe subset relationship check |
| `SliceNotSubsetT[E comparable]` | Slice element | Type-safe non-subset check |

**Origin**: Generic initiative + [#1685] (partial - SeqContains variants only)

**Performance**: 16-81x faster than reflection-based variants (see [benchmarks](../project/maintainers/BENCHMARKS.md))

[#1685]: https://github.com/stretchr/testify/pull/1685
{{% /expand %}}


**Behavior changes**: None

### Comparison

{{% expand title="Generics" %}}

#### New Generic Functions (6)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `GreaterT[V Ordered]` | extended Ordered type[^1] | Type-safe greater-than comparison |
| `GreaterOrEqualT[V Ordered]` | Ordered type | Type-safe >= comparison |
| `LessT[V Ordered]` | Ordered type | Type-safe less-than comparison |
| `LessOrEqualT[V Ordered]` | Ordered type | Type-safe <= comparison |
| `PositiveT[V Ordered]` | Ordered type | Type-safe positive value check (> 0) |
| `NegativeT[V Ordered]` | Ordered type | Type-safe negative value check (< 0) |

[^1]: Ordered is defined as the union of standard go ordered types, plus `[]byte` and `time.Time`.

**Origin**: Generics initiative

**Performance**: 10-22x faster than reflection-based variants. See [Benchmarks](../project/maintainers/BENCHMARKS.md)
{{% /expand %}}

**Behavior changes**: None

### Condition

#### New Function (1)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `Consistently[C Conditioner]` | `func() bool` or `func(context.Context) error` | async assertion to express "always true" (adapted proposal [#1606], [#1087]) |

[#1087]: https://github.com/stretchr/testify/issues/1087
[#1606]: https://github.com/stretchr/testify/pulls/1606

#### ⚠️ Behavior Changes

| Change | Origin | Description |
|--------|--------|-------------|
| Fixed goroutine leak | [#1611] | Consolidated `Eventually`, `Never`, and `EventuallyWith` into single `pollCondition` function |
| Context-based polling | Internal refactoring | Reimplemented with context-based approach for better resource management |
| Unified implementation | Internal refactoring | Single implementation eliminates code duplication and prevents resource leaks |
| `func(context.Context) error` conditions | extensions to the async domain | control over context allows for more complex cases to be supported |
| Type parameter | Internal refactoring | `Eventually` now accepts several signatures for its condition and uses a type parameter (non-breaking) |

**Impact**: This fix eliminates goroutine leaks that could occur when using `Eventually` or `Never` assertions. The new implementation uses a context-based approach that properly manages resources and provides a cleaner shutdown mechanism. Callers should **NOT** assume that the call to `Eventually` or `Never` exits before the condition is evaluated. Callers should **NOT** assume that the call to `Eventually` or `Never` exits before the condition is evaluated.

**Supersedes**: This implementation also supersedes upstream proposals [#1819] (handle unexpected exits) and [#1830] (`CollectT.Halt`) with a more comprehensive solution.

[#1611]: https://github.com/stretchr/testify/issues/1611
[#1819]: https://github.com/stretchr/testify/pull/1819
[#1830]: https://github.com/stretchr/testify/pull/1830

### Breaking Changes

| Change | Origin | Description |
|--------|--------|-------------|
| **Renaming** | Internal refactorting | `EventuallyWithT` renamed into `EventuallyWith` (conflicted with the convention adopted for generics) |
| **Removal**  | API simplification    | `Comparison` type is removed as a mere alias to `func() bool` |

### Equality

{{% expand title="Generics" %}}

#### New Generic Functions (4)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `EqualT[V comparable]` | Comparable type | Type-safe equality check |
| `NotEqualT[V comparable]` | Comparable type | Type-safe inequality check |
| `SameT[V comparable]` | Comparable type | Type-safe pointer identity check |
| `NotSameT[V comparable]` | Comparable type | Type-safe pointer difference check |

**Origin**: Generics initiative

**Performance**: 10-13x faster for Equal/NotEqual, 1.5-2x for Same/NotSame
{{% /expand %}}

#### ⚠️ Behavior Changes

| Function | Change | Reason |
|----------|--------|--------|
| `EqualValues` | Now fails with function types (like `Equal`) | [#1825] - Consistency and safety |
| `Same`/`NotSame` | Two nil pointers of same type now correctly considered "same" | Edge case fix |

[#1825]: https://github.com/stretchr/testify/pull/1825

### Error

**New functions**: None

**Behavior changes**: None

### File

| Function | Type | Origin | Description |
|----------|------|--------|-------------|
| `FileEmpty` | Reflection | New addition | Assert file exists and is empty (0 bytes) |
| `FileNotEmpty` | Reflection | New addition | Assert file exists and is not empty |

**Note**: `DirExists` was already present in upstream, `NoDirExists` renamed into `DirNotExists`. `NoFileExists` renamed into `FileNotExists`.²

**Behavior changes**: None

### HTTP

**New functions**: None

**Behavior changes**: None

### JSON

{{% expand title="Generics" %}}

#### New Generic Functions (3)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `JSONEqT[S Text]` | String or []byte | Type-safe JSON semantic equality |
| `JSONMarshalAsT[EDoc Text]` | String or []byte | Type-safe JSON marshal and equality check |
| `JSONUnmarshalAsT[ADoc Text, Object any]` | String or []byte | Type-safe JSON unmarshal and equality check |

**Performance**: Comparable (JSON parsing dominates)
{{% /expand %}}

{{% expand title="Reflection-based" %}}

#### New Reflection Function (1)

| Function | Origin | Description |
|----------|--------|-------------|
| `JSONEqBytes` | [#1513] | JSON equality for byte slices |

[#1513]: https://github.com/stretchr/testify/pull/1513
{{% /expand %}}

**Behavior changes**: None

### Number

{{% expand title="Generics" %}}

#### New Generic Functions (2)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `InDeltaT[V Float\|Integer]` | Numeric type | Type-safe float comparison with absolute delta |
| `InEpsilonT[V Float]` | Float type | Type-safe float comparison with relative epsilon |

**Origin**: Generics initiative

**Performance**: 1.2-1.5x faster
{{% /expand %}}

#### ⚠️ Behavior Changes

- Fixed IEEE 754 edge case handling (NaN, Inf)
- Added support for zero expected value in `InEpsilon` (falls back to absolute error)
- Fixed invalid type conversion for `uintptr` in reflect-based compare

### Ordering

{{% expand title="Generics" %}}

#### New Generic Functions (6)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `IsIncreasingT[E Ordered]` | Ordered[^1] slice element | Type-safe strictly increasing check |
| `IsDecreasingT[E Ordered]` | Ordered slice element | Type-safe strictly decreasing check |
| `IsNonIncreasingT[E Ordered]` | Ordered slice element | Type-safe non-increasing check (allows equal) |
| `IsNonDecreasingT[E Ordered]` | Ordered slice element | Type-safe non-decreasing check (allows equal) |
| `SortedT[E cmp.Ordered]` | Ordered slice element | Type-safe sorted check (generic-only function) |
| `NotSortedT[E cmp.Ordered]` | Ordered slice element | Type-safe unsorted check (generic-only function) |

**Origin**: Generics initiative

**Performance**: 6.5-9.5x faster

**Note**: `SortedT` and `NotSortedT` are generic-only (no reflection equivalents)
{{% /expand %}}

#### ⚠️ Behavior Changes

| Function | Change | Reason |
|----------|--------|--------|
| `IsNonDecreasing` | Logic corrected to match documentation | Inverted logic fixed |
| `IsNonIncreasing` | Logic corrected to match documentation | Inverted logic fixed |

### Panic

**New functions**: None

#### ⚠️ Behavior Changes

Removed extraneous type declaration `PanicTestFunc` (`func()`).

### Safety

**New domain** for resource leak detection.

| Function | Type | Description |
|----------|------|-------------|
| `NoGoRoutineLeak` | Reflection | Assert that no goroutines leak from a tested function |
| `NoFileDescriptorLeak` | Reflection | Assert that no file descriptors leak from a tested function (Linux) |

#### Implementation

`NoGoRoutineLeak` uses **pprof labels** instead of stack-trace heuristics (like `go.uber.org/goleak`):
- Only goroutines spawned by the tested function are checked
- Pre-existing goroutines (runtime, pools, parallel tests) are ignored automatically
- No configuration or filter lists needed
- Works safely with `t.Parallel()`

`NoFileDescriptorLeak` compares open file descriptors before and after the tested function (Linux only, via `/proc/self/fd`).

See [Examples](./EXAMPLES.md#goroutine-leak-detection) for usage patterns.

### String

{{% expand title="Generics" %}}

#### New Generic Functions (2)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `RegexpT[S Text]` | String or []byte | Type-safe regex match check |
| `NotRegexpT[S Text]` | String or []byte | Type-safe regex non-match check |

**Origin**: Generics initiative

**Performance**: 1.2x faster (regex compilation dominates)
{{% /expand %}}

#### ⚠️ Behavior Changes

| Change | Origin | Description |
|--------|--------|-------------|
| Fix panic on invalid regex | [#1818] | Handle invalid regex patterns gracefully |
| Refactored regex handling | Internal | Fixed quirks with unexpected behavior on some input types |

[#1818]: https://github.com/stretchr/testify/pull/1818

### Testing

**New functions**: None

**Behavior changes**: None

### Time

**New functions**: None

#### ⚠️ Behavior Changes

| Change | Origin | Description |
|--------|--------|-------------|
| Fix time.Time rendering in diffs | [#1829] | Improved time display in failure messages |

[1829]: https://github.com/stretchr/testify/issues/1829

### Type

{{% expand title="Generics" %}}

#### New Generic Functions (2)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `IsOfTypeT[EType any]` | Expected type | Type assertion without dummy value |
| `IsNotOfTypeT[EType any]` | Expected type | Negative type assertion without dummy value |

**Origin**: [#1805]
**Performance**: 9-11x faster

[#1805]: https://github.com/stretchr/testify/issues/1805

{{% /expand %}}

{{% expand title="Reflection-based" %}}
#### New Reflection Functions (2)

| Function | Origin | Description |
|----------|--------|-------------|
| `Kind` | [#1803] | Assert value is of specific reflect.Kind |
| `NotKind` | [#1803] | Assert value is not of specific reflect.Kind |

[#1803]: https://github.com/stretchr/testify/pull/1803
{{% /expand %}}

**Behavior changes**: None

### YAML

{{% expand title="Generics" %}}

#### New Generic Functions (3)

| Function | Type Parameters | Description |
|----------|-----------------|-------------|
| `YAMLEqT[S Text]` | String or []byte | Type-safe YAML semantic equality |
| `YAMLMarshalAsT[EDoc Text]` | String or []byte | Type-safe YAML marshal and equality check |
| `YAMLUnmarshalAsT[ADoc Text, Object any]` | String or []byte | Type-safe YAML unmarshal and equality check |

**Performance**: Comparable (YAML parsing dominates)
{{% /expand %}}

{{% expand title="Reflection-based" %}}

#### New Reflection Function (1)

| Function | Origin | Description |
|----------|--------|-------------|
| `YAMLEqBytes` | Consistency | YAML equality for byte slices (matches JSONEqBytes) |
{{% /expand %}}


#### ⚠️ Behavior Changes

**Architecture change**: YAML support is now opt-in via `import _ "github.com/go-openapi/testify/v2/enable/yaml"`

**Behavior changes**: None

## Other changes

### Performance Improvements

See [Performance Benchmarks](../project/maintainers/BENCHMARKS.md) for a detailed presentation.

#### Generic vs Reflection Performance

| Domain | Function | Speedup | Key Benefit |
|--------|----------|---------|-------------|
| Collection | ElementsMatchT | **21-81x** | Scales with collection size |
| Equality | EqualT | **10-13x** | Zero allocations |
| Comparison | GreaterT/LessT | **10-22x** | Zero allocations |
| Collection | SliceContainsT | **16x** | Zero allocations |
| Collection | SeqContainsT | **25x** | Iterator optimization |
| Ordering | IsIncreasingT | **7-9x** | Zero allocations |
| Type | IsOfTypeT | **9-11x** | No reflection overhead |

**Memory savings**: Up to 99% reduction in allocations for large collections

### Architecture Changes

These affect the way the project is maintained, but not how it is used.

#### Code Generation

All assert and require packages are 100% generated from a single source:
- **Source**: `internal/assertions/` (~6,000 LOC)
- **Generated**: ~800+ functions across assert/require packages
- **Variants**: 8 variants per assertion (assert/require x standard/format/forward/forward+format),
  4 variants for generic assertions (assert/require x standard/format)

> NOTE: generic assertions obviously can't be propagated as a "forward variant", i.e
> as a method of the `Assertion` object.

#### Module Structure

The project adopts a mono-repo structure (with the appropriate changes made in CI).

This means that the github repo exposes several independant go modules.

```
github.com/go-openapi/testify/v2           # Core (zero deps) [go.mod]
├── assert/                                # Generated package
├── require/                               # Generated package
├── internal/                              # Internalized dependencies
│   ├── spew/                              # Internalized go-spew
│   ├── difflib/                           # Internalized go-difflib
│   └── assertions/                        # Single source of truth
├── enable/                                # Modules for optional features
│   ├── yaml/                              # Optional YAML support [go.mod]
│   └── color/                             # Optional colorization [go.mod]
│
└── codegen/                               # Code and documentation generator [go.mod]
```

### Documentation

- Hugo-based documentation site
- Domain-organized API reference (19 domains)
- Comprehensive examples and tutorials
- Performance benchmarks

## Project Metrics

| Metric | Value |
|--------|-------|
| **New functions** | 58 (43 generic + 15 reflection) |
| **Total assertions** | 127 base assertions |
| **Generated functions** | 840 (see [the maths](../project/maintainers/ARCHITECTURE.md#the-maths-with-assertion-variants)) |
| **Generic coverage** | 10 domains (10/19) |
| **Performance improvement** | 1.2x to 81x faster |
| **Dependencies** | 0 external (was 2 required) |
| **Test coverage** | 96% overall, 99% on public APIs |
| **Documentation domains** | 19 logical categories |

---

## See Also

- [Migration Guide](./MIGRATION.md) - Step-by-step guide to migrating from testify v1
- [Generics Guide](./GENERICS.md) - Detailed documentation of all 43 generic assertions
- [Performance Benchmarks](../project/maintainers/BENCHMARKS.md) - Comprehensive performance analysis
- [Examples](./EXAMPLES.md) - Practical usage examples showing new features
- [Tutorial](./TUTORIAL.md) - Best practices for writing tests with testify v2
- [API Reference](../api/_index.md) - Complete assertion catalog organized by domain

