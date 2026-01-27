---
title: 'Benchmarks'
description: 'Performance Benchmarks: Generic vs Reflection'
weight: 10
---

**Last Updated**: 2026-01-20

## Overview

While the primary motivation for adding **generic assertion functions** to testify v2 was **compile-time type safety**
(see [Generics Guide](../../usage/GENERICS.md) for details), our benchmarking revealed an unexpected bonus:
**dramatic performance improvements** ranging from 1.2x to 81x faster,
with up to 99% reduction in memory allocations for collection operations.

This document focuses on the performance measurements and explains why these improvements occur.

## Type Safety First, Performance Second

Generic assertions catch type errors when writing tests, not when running them. For example:

```go
// Reflection: Compiles, fails at runtime
assert.ElementsMatch(t, []int{1, 2}, []string{"a", "b"})

// Generic: Compiler catches the error immediately
assert.ElementsMatchT(t, []int{1, 2}, []string{"a", "b"})  // ❌ Compile error!
```

See the [Generics Guide](../../usage/GENERICS.md) for comprehensive coverage of type safety benefits,
refactoring safety, and when to use generic vs reflection variants.

## Performance Results by Category

### 🏆 Collection Operations: Exceptional Gains

Collection operations see the most dramatic improvements due to elimination of per-element reflection overhead:

| Function | Speedup | Memory Impact | Why It Matters |
|----------|---------|---------------|----------------|
| **ElementsMatch (10 items)** | **21x faster** | 568 B → 320 B (44% reduction) | Common test operation |
| **ElementsMatch (100 items)** | **39x faster** | 41 KB → 3.6 KB (91% reduction) | Scales superlinearly |
| **ElementsMatch (1000 items)** | **81x faster** | 4 MB → 33 KB (99% reduction) | Large collection testing |
| **SliceContains** | **16x faster** | 4 allocs → 0 | Membership testing |
| **SeqContains (iter.Seq)** | **25x faster** | 55 allocs → 9 | Go 1.23+ iterators |
| **SliceSubset** | **43x faster** | 17 allocs → 0 | Subset verification |

**Key insight**: ElementsMatch's O(n²) complexity amplifies the benefits—the speedup **increases** with collection size (21x → 39x → 81x).

### ⚡ Comparison Operations: Zero-Allocation Wins

Direct operator usage (`>`, `<`, `==`) eliminates reflection overhead and boxing entirely:

| Function | Speedup | Allocations | Benchmark Data |
|----------|---------|-------------|----------------|
| **Greater/Less** | **10-15x faster** | 1 → 0 allocs | 139.1ns → 17.9ns |
| **Positive/Negative** | **16-22x faster** | 1 → 0 allocs | 121.5ns → 7.6ns |
| **GreaterOrEqual/LessOrEqual** | **10-11x faster** | 1 → 0 allocs | Similar pattern |
| **Equal** | **10-13x faster** | 0 allocs (both) | 44.8ns → 3.5ns |
| **NotEqual** | **11x faster** | 0 allocs (both) | Comparable to Equal |

**Key insight**: Comparison operations are frequently used in tests. 10-15x speedup on common assertions accumulates quickly across large test suites.

### 📊 Ordering Operations: Eliminating Per-Element Overhead

Ordering checks iterate over collections, so eliminating per-element reflection creates significant gains:

| Function | Speedup | Allocation Impact |
|----------|---------|-------------------|
| **IsIncreasing** | **7.4x faster** | 11 allocs → 0 |
| **IsDecreasing** | **9.5x faster** | 11 allocs → 0 |
| **IsNonDecreasing** | **8x faster** | 4 allocs → 0 |
| **IsNonIncreasing** | **6.5x faster** | 4 allocs → 0 |

### 🔍 Type Checks: Cleaner API, Better Performance

Generic type checks eliminate reflection and provide a cleaner API:

| Function | Speedup | Notes |
|----------|---------|-------|
| **IsOfType** | **9-11x faster** | No dummy value needed with generics |
| **IsNotOfType** | **Similar gains** | Type parameter makes intent explicit |

### ⚖️ Modest Gains: Where Processing Dominates

Some operations see smaller improvements because expensive processing dominates:

| Category | Speedup | Why Gains Are Limited |
|----------|---------|----------------------|
| **Same/NotSame** | **1.5-2x** | Pointer comparison already fast |
| **Boolean checks** | **~2x** | Simple bool comparison |
| **JSONEq** | **Marginal** | JSON parsing/unmarshaling dominates |
| **Regexp** | **Marginal** | Regex compilation dominates |

**Key insight**: Even modest performance gains come with the benefit of compile-time type safety.

## Understanding the Performance Gains

### Allocation Elimination

The most dramatic speedups come from eliminating allocations entirely:

- **ElementsMatch (1000 elements)**: 501,503 → 3 allocations (99.999% reduction)
- **All comparison operations**: 1 → 0 allocations
- **Ordering checks**: 4-11 → 0 allocations

Less allocation pressure means faster execution and reduced GC overhead, especially impactful in large test suites.

### Superlinear Scaling

For operations with O(n²) or O(n) complexity, eliminating per-element reflection overhead creates superlinear gains:

- **ElementsMatch**: 21x (10 items) → 39x (100 items) → 81x (1000 items)
- The speedup **increases** with collection size

### Cumulative Impact

Test suites typically run thousands of assertions:

- **Small test suite** (1,000 assertions): 10x average speedup = significantly faster CI runs
- **Large test suite** (10,000+ assertions): Cumulative savings become substantial
- **Particularly valuable** in CI/CD pipelines where test execution time directly affects deployment velocity

## Sample Benchmark Data

Representative results from `go test -bench=. ./internal/assertions`:

```
# Collection operations
BenchmarkElementsMatch/reflect/1000-16   25.5 ms/op     4.0 MB/op   501503 allocs/op
BenchmarkElementsMatch/generic/1000-16    316 µs/op      33 KB/op        3 allocs/op
                                          ↑ 81x faster   ↑ 99% less memory

# Comparison operations
BenchmarkGreater/reflect/int-16          139.1 ns/op      34 B/op        1 allocs/op
BenchmarkGreater/generic/int-16           17.9 ns/op       0 B/op        0 allocs/op
                                          ↑ 7.8x faster

# Equality checks
BenchmarkEqual/reflect/int-16             44.8 ns/op       0 B/op        0 allocs/op
BenchmarkEqual/generic/int-16              3.5 ns/op       0 B/op        0 allocs/op
                                          ↑ 13x faster
```

## Adopting Generic Assertions

See the [Migration Guide](../../usage/MIGRATION.md) for step-by-step instructions on migrating to generic assertions, and the [Generics Guide](../../usage/GENERICS.md) for comprehensive coverage of type safety benefits and usage patterns.

---

## Running Benchmarks

To run the benchmarks yourself:

```bash
go test -run=^$ -bench=. -benchmem ./internal/assertions
```

## Benchmark Coverage

**38 generic functions benchmarked across 10 domains:**
- Boolean (2): TrueT, FalseT
- Collection (12): StringContainsT, SliceContainsT, MapContainsT, SeqContainsT, ElementsMatchT, SliceSubsetT, and negative variants
- Comparison (6): GreaterT, LessT, GreaterOrEqualT, LessOrEqualT, PositiveT, NegativeT
- Equality (4): EqualT, NotEqualT, SameT, NotSameT
- JSON (1): JSONEqT
- Number (2): InDeltaT, InEpsilonT
- Ordering (6): IsIncreasingT, IsDecreasingT, IsNonIncreasingT, IsNonDecreasingT, SortedT, NotSortedT
- String (2): RegexpT, NotRegexpT
- Type (2): IsOfTypeT, IsNotOfTypeT
- YAML (1): YAMLEqT (benchmarked separately in enable/yaml module)

---
