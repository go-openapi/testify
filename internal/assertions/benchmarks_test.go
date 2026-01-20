// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"slices"
	"testing"
)

// Helper functions to reduce duplication in benchmarks

// benchmarkComparison runs a benchmark comparing reflection vs generic implementations.
func benchmarkComparison(
	b *testing.B,
	name string,
	reflectFn func(*mockT),
	genericFn func(*mockT),
) {
	b.Helper()
	mockT := &mockT{}

	b.Run("reflect/"+name, func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			reflectFn(mockT)
		}
	})

	b.Run("generic/"+name, func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			genericFn(mockT)
		}
	})
}

// benchmarkGenericOnly runs a benchmark for generic-only functions (no reflection equivalent).
func benchmarkGenericOnly(b *testing.B, genericFn func(*mockT)) {
	b.Helper()
	mockT := &mockT{}

	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			genericFn(mockT)
		}
	})
}

func Benchmark_isEmpty(b *testing.B) {
	b.ReportAllocs()

	v := new(int)

	for b.Loop() {
		isEmpty("")
		isEmpty(42)
		isEmpty(v)
	}
}

func BenchmarkNotNil(b *testing.B) {
	for b.Loop() {
		NotNil(b, b)
	}
}

func BenchmarkBytesEqual(b *testing.B) {
	const size = 1024 * 8
	s := make([]byte, size)
	for i := range s {
		s[i] = byte(i % 255)
	}
	s2 := make([]byte, size)
	copy(s2, s)

	mockT := &mockFailNowT{}

	for b.Loop() {
		Equal(mockT, s, s2)
	}
}

/*
 * Benchmarks comparing reflect-based vs generic implementations.
 *
 * These benchmarks measure the performance difference between:
 * - Reflect-based versions (e.g., Greater, ElementsMatch)
 * - Generic versions (e.g., GreaterT, ElementsMatchT)
 */

// BenchmarkGreater compares Greater (reflect) vs GreaterT (generic).
func BenchmarkGreater(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { Greater(t, 100, 50) },
		func(t *mockT) { GreaterT(t, 100, 50) },
	)

	benchmarkComparison(b, "float64",
		func(t *mockT) { Greater(t, 100.5, 50.5) },
		func(t *mockT) { GreaterT(t, 100.5, 50.5) },
	)

	benchmarkComparison(b, "string",
		func(t *mockT) { Greater(t, "beta", "alpha") },
		func(t *mockT) { GreaterT(t, "beta", "alpha") },
	)
}

// BenchmarkLess compares Less (reflect) vs LessT (generic).
func BenchmarkLess(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { Less(t, 50, 100) },
		func(t *mockT) { LessT(t, 50, 100) },
	)
}

// BenchmarkElementsMatch compares ElementsMatch (reflect) vs ElementsMatchT (generic).
func BenchmarkElementsMatch(b *testing.B) {
	// Small slices (10 elements)
	smallA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	smallB := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	benchmarkComparison(b, "small_10",
		func(t *mockT) { ElementsMatch(t, smallA, smallB) },
		func(t *mockT) { ElementsMatchT(t, smallA, smallB) },
	)

	// Medium slices (100 elements)
	mediumA := make([]int, 100)
	mediumB := make([]int, 100)
	for i := range mediumA {
		mediumA[i] = i
		mediumB[99-i] = i
	}
	benchmarkComparison(b, "medium_100",
		func(t *mockT) { ElementsMatch(t, mediumA, mediumB) },
		func(t *mockT) { ElementsMatchT(t, mediumA, mediumB) },
	)

	// Large slices (1000 elements)
	largeA := make([]int, 1000)
	largeB := make([]int, 1000)
	for i := range largeA {
		largeA[i] = i
		largeB[999-i] = i
	}
	benchmarkComparison(b, "large_1000",
		func(t *mockT) { ElementsMatch(t, largeA, largeB) },
		func(t *mockT) { ElementsMatchT(t, largeA, largeB) },
	)

	// String slices
	stringsA := []string{"apple", "banana", "cherry", "date", "elderberry"}
	stringsB := []string{"elderberry", "date", "cherry", "banana", "apple"}
	benchmarkComparison(b, "strings_5",
		func(t *mockT) { ElementsMatch(t, stringsA, stringsB) },
		func(t *mockT) { ElementsMatchT(t, stringsA, stringsB) },
	)
}

// BenchmarkNotElementsMatch compares NotElementsMatch (reflect) vs NotElementsMatchT (generic).
func BenchmarkNotElementsMatch(b *testing.B) {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{1, 2, 3, 4, 6}
	benchmarkComparison(b, "",
		func(t *mockT) { NotElementsMatch(t, sliceA, sliceB) },
		func(t *mockT) { NotElementsMatchT(t, sliceA, sliceB) },
	)
}

// BenchmarkPositive compares Positive (reflect) vs PositiveT (generic).
func BenchmarkPositive(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { Positive(t, 42) },
		func(t *mockT) { PositiveT(t, 42) },
	)
	benchmarkComparison(b, "float64",
		func(t *mockT) { Positive(t, 42.5) },
		func(t *mockT) { PositiveT(t, 42.5) },
	)
}

// BenchmarkNegative compares Negative (reflect) vs NegativeT (generic).
func BenchmarkNegative(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { Negative(t, -42) },
		func(t *mockT) { NegativeT(t, -42) },
	)
}

// BenchmarkEqual compares Equal (reflect) vs EqualT (generic).
func BenchmarkEqual(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { Equal(t, 42, 42) },
		func(t *mockT) { EqualT(t, 42, 42) },
	)
	benchmarkComparison(b, "string",
		func(t *mockT) { Equal(t, "hello", "hello") },
		func(t *mockT) { EqualT(t, "hello", "hello") },
	)
	benchmarkComparison(b, "float64",
		func(t *mockT) { Equal(t, 3.14, 3.14) },
		func(t *mockT) { EqualT(t, 3.14, 3.14) },
	)
}

// BenchmarkNotEqual compares NotEqual (reflect) vs NotEqualT (generic).
func BenchmarkNotEqual(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { NotEqual(t, 42, 43) },
		func(t *mockT) { NotEqualT(t, 42, 43) },
	)
}

// BenchmarkSame compares Same (reflect) vs SameT (generic).
func BenchmarkSame(b *testing.B) {
	v := 42
	p := &v
	benchmarkComparison(b, "",
		func(t *mockT) { Same(t, p, p) },
		func(t *mockT) { SameT(t, p, p) },
	)
}

// BenchmarkNotSame compares NotSame (reflect) vs NotSameT (generic).
func BenchmarkNotSame(b *testing.B) {
	v1, v2 := 42, 42
	p1, p2 := &v1, &v2
	benchmarkComparison(b, "",
		func(t *mockT) { NotSame(t, p1, p2) },
		func(t *mockT) { NotSameT(t, p1, p2) },
	)
}

// BenchmarkGreaterOrEqual compares GreaterOrEqual (reflect) vs GreaterOrEqualT (generic).
func BenchmarkGreaterOrEqual(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { GreaterOrEqual(t, 100, 50) },
		func(t *mockT) { GreaterOrEqualT(t, 100, 50) },
	)
}

// BenchmarkLessOrEqual compares LessOrEqual (reflect) vs LessOrEqualT (generic).
func BenchmarkLessOrEqual(b *testing.B) {
	benchmarkComparison(b, "int",
		func(t *mockT) { LessOrEqual(t, 50, 100) },
		func(t *mockT) { LessOrEqualT(t, 50, 100) },
	)
}

// BenchmarkIsIncreasing compares IsIncreasing (reflect) vs IsIncreasingT (generic).
func BenchmarkIsIncreasing(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	benchmarkComparison(b, "",
		func(t *mockT) { IsIncreasing(t, slice) },
		func(t *mockT) { IsIncreasingT(t, slice) },
	)
}

// BenchmarkIsNonIncreasing compares IsNonIncreasing (reflect) vs IsNonIncreasingT (generic).
func BenchmarkIsNonIncreasing(b *testing.B) {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	benchmarkComparison(b, "",
		func(t *mockT) { IsNonIncreasing(t, slice) },
		func(t *mockT) { IsNonIncreasingT(t, slice) },
	)
}

// BenchmarkIsDecreasing compares IsDecreasing (reflect) vs IsDecreasingT (generic).
func BenchmarkIsDecreasing(b *testing.B) {
	slice := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	benchmarkComparison(b, "",
		func(t *mockT) { IsDecreasing(t, slice) },
		func(t *mockT) { IsDecreasingT(t, slice) },
	)
}

// BenchmarkIsNonDecreasing compares IsNonDecreasing (reflect) vs IsNonDecreasingT (generic).
func BenchmarkIsNonDecreasing(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	benchmarkComparison(b, "",
		func(t *mockT) { IsNonDecreasing(t, slice) },
		func(t *mockT) { IsNonDecreasingT(t, slice) },
	)
}

// BenchmarkContains compares Contains (reflect) vs various ContainsT generics.
func BenchmarkContains(b *testing.B) {
	benchmarkComparison(b, "string",
		func(t *mockT) { Contains(t, "hello world", "world") },
		func(t *mockT) { StringContainsT(t, "hello world", "world") },
	)

	slice := []int{1, 2, 3, 4, 5}
	benchmarkComparison(b, "slice",
		func(t *mockT) { Contains(t, slice, 3) },
		func(t *mockT) { SliceContainsT(t, slice, 3) },
	)

	m := map[string]int{"a": 1, "b": 2, "c": 3}
	benchmarkComparison(b, "map",
		func(t *mockT) { Contains(t, m, "b") },
		func(t *mockT) { MapContainsT(t, m, "b") },
	)
}

// BenchmarkNotContains compares NotContains (reflect) vs various NotContainsT generics.
func BenchmarkNotContains(b *testing.B) {
	benchmarkComparison(b, "string",
		func(t *mockT) { NotContains(t, "hello world", "xyz") },
		func(t *mockT) { StringNotContainsT(t, "hello world", "xyz") },
	)

	slice := []int{1, 2, 3, 4, 5}
	benchmarkComparison(b, "slice",
		func(t *mockT) { NotContains(t, slice, 99) },
		func(t *mockT) { SliceNotContainsT(t, slice, 99) },
	)
}

// BenchmarkSubset compares Subset (reflect) vs SubsetT (generic).
func BenchmarkSubset(b *testing.B) {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	subset := []int{2, 4, 6}
	benchmarkComparison(b, "",
		func(t *mockT) { Subset(t, list, subset) },
		func(t *mockT) { SliceSubsetT(t, list, subset) },
	)
}

// BenchmarkNotSubset compares NotSubset (reflect) vs NotSubsetT (generic).
func BenchmarkNotSubset(b *testing.B) {
	list := []int{1, 2, 3, 4, 5}
	subset := []int{1, 2, 99}
	benchmarkComparison(b, "",
		func(t *mockT) { NotSubset(t, list, subset) },
		func(t *mockT) { SliceNotSubsetT(t, list, subset) },
	)
}

// BenchmarkInDelta compares InDelta (reflect) vs InDeltaT (generic).
func BenchmarkInDelta(b *testing.B) {
	benchmarkComparison(b, "float64",
		func(t *mockT) { InDelta(t, 3.14, 3.15, 0.02) },
		func(t *mockT) { InDeltaT(t, 3.14, 3.15, 0.02) },
	)
	benchmarkComparison(b, "int",
		func(t *mockT) { InDelta(t, 100, 102, 5) },
		func(t *mockT) { InDeltaT(t, 100, 102, 5) },
	)
}

// BenchmarkInEpsilon compares InEpsilon (reflect) vs InEpsilonT (generic).
func BenchmarkInEpsilon(b *testing.B) {
	benchmarkComparison(b, "float64",
		func(t *mockT) { InEpsilon(t, 100.0, 101.0, 0.02) },
		func(t *mockT) { InEpsilonT(t, 100.0, 101.0, 0.02) },
	)
}

// BenchmarkTrue compares True (reflect) vs TrueT (generic).
func BenchmarkTrue(b *testing.B) {
	benchmarkComparison(b, "",
		func(t *mockT) { True(t, true) },
		func(t *mockT) { TrueT(t, true) },
	)
}

// BenchmarkFalse compares False (reflect) vs FalseT (generic).
func BenchmarkFalse(b *testing.B) {
	benchmarkComparison(b, "",
		func(t *mockT) { False(t, false) },
		func(t *mockT) { FalseT(t, false) },
	)
}

// BenchmarkRegexp compares Regexp (reflect) vs RegexpT (generic).
func BenchmarkRegexp(b *testing.B) {
	benchmarkComparison(b, "string",
		func(t *mockT) { Regexp(t, `^\d{3}-\d{4}$`, "123-4567") },
		func(t *mockT) { RegexpT(t, `^\d{3}-\d{4}$`, "123-4567") },
	)
}

// BenchmarkNotRegexp compares NotRegexp (reflect) vs NotRegexpT (generic).
func BenchmarkNotRegexp(b *testing.B) {
	benchmarkComparison(b, "string",
		func(t *mockT) { NotRegexp(t, `^\d{3}-\d{4}$`, "hello") },
		func(t *mockT) { NotRegexpT(t, `^\d{3}-\d{4}$`, "hello") },
	)
}

// BenchmarkSorted benchmarks SortedT (generic-only, no reflection version).
func BenchmarkSorted(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	benchmarkGenericOnly(b, func(t *mockT) { SortedT(t, slice) })
}

// BenchmarkNotSorted benchmarks NotSortedT (generic-only, no reflection version).
func BenchmarkNotSorted(b *testing.B) {
	slice := []int{1, 3, 2, 4, 5}
	benchmarkGenericOnly(b, func(t *mockT) { NotSortedT(t, slice) })
}

// BenchmarkSeqContains compares Contains (reflect) vs SeqContainsT (generic) for iterators.
func BenchmarkSeqContains(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	benchmarkComparison(b, "",
		func(t *mockT) { Contains(t, slices.Values(slice), 5) },
		func(t *mockT) { SeqContainsT(t, slices.Values(slice), 5) },
	)
}

// BenchmarkSeqNotContains compares NotContains (reflect) vs SeqNotContainsT (generic) for iterators.
func BenchmarkSeqNotContains(b *testing.B) {
	slice := []int{1, 2, 3, 4, 5}
	benchmarkComparison(b, "",
		func(t *mockT) { NotContains(t, slices.Values(slice), 99) },
		func(t *mockT) { SeqNotContainsT(t, slices.Values(slice), 99) },
	)
}

// BenchmarkJSONEq compares JSONEq (reflect) vs JSONEqT (generic).
func BenchmarkJSONEq(b *testing.B) {
	expected := `{"name":"John","age":30}`
	actual := `{"age":30,"name":"John"}`
	benchmarkComparison(b, "",
		func(t *mockT) { JSONEq(t, expected, actual) },
		func(t *mockT) { JSONEqT(t, expected, actual) },
	)
}

// BenchmarkIsOfType compares IsType (reflect) vs IsOfTypeT (generic).
func BenchmarkIsOfType(b *testing.B) {
	var expected int
	actual := 42
	benchmarkComparison(b, "",
		func(t *mockT) { IsType(t, expected, actual) },
		func(t *mockT) { IsOfTypeT[int](t, actual) },
	)
}

// BenchmarkIsNotOfType compares IsNotType (reflect) vs IsNotOfTypeT (generic).
func BenchmarkIsNotOfType(b *testing.B) {
	var expected int
	actual := "string"
	benchmarkComparison(b, "",
		func(t *mockT) { IsNotType(t, expected, actual) },
		func(t *mockT) { IsNotOfTypeT[int](t, actual) },
	)
}
