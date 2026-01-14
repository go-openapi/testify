// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

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
	mockT := &mockT{}

	b.Run("reflect/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			Greater(mockT, 100, 50)
		}
	})

	b.Run("generic/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			GreaterT(mockT, 100, 50)
		}
	})

	b.Run("reflect/float64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			Greater(mockT, 100.5, 50.5)
		}
	})

	b.Run("generic/float64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			GreaterT(mockT, 100.5, 50.5)
		}
	})

	b.Run("reflect/string", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			Greater(mockT, "beta", "alpha")
		}
	})

	b.Run("generic/string", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			GreaterT(mockT, "beta", "alpha")
		}
	})
}

// BenchmarkLess compares Less (reflect) vs LessT (generic).
func BenchmarkLess(b *testing.B) {
	mockT := &mockT{}

	b.Run("reflect/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			Less(mockT, 50, 100)
		}
	})

	b.Run("generic/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			LessT(mockT, 50, 100)
		}
	})
}

// BenchmarkElementsMatch compares ElementsMatch (reflect) vs ElementsMatchT (generic).
func BenchmarkElementsMatch(b *testing.B) {
	mockT := &mockT{}

	// Small slices (10 elements)
	smallA := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	smallB := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	b.Run("reflect/small_10", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatch(mockT, smallA, smallB)
		}
	})

	b.Run("generic/small_10", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatchT(mockT, smallA, smallB)
		}
	})

	// Medium slices (100 elements)
	mediumA := make([]int, 100)
	mediumB := make([]int, 100)
	for i := range mediumA {
		mediumA[i] = i
		mediumB[99-i] = i
	}

	b.Run("reflect/medium_100", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatch(mockT, mediumA, mediumB)
		}
	})

	b.Run("generic/medium_100", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatchT(mockT, mediumA, mediumB)
		}
	})

	// Large slices (1000 elements)
	largeA := make([]int, 1000)
	largeB := make([]int, 1000)
	for i := range largeA {
		largeA[i] = i
		largeB[999-i] = i
	}

	b.Run("reflect/large_1000", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatch(mockT, largeA, largeB)
		}
	})

	b.Run("generic/large_1000", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatchT(mockT, largeA, largeB)
		}
	})

	// String slices
	stringsA := []string{"apple", "banana", "cherry", "date", "elderberry"}
	stringsB := []string{"elderberry", "date", "cherry", "banana", "apple"}

	b.Run("reflect/strings_5", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatch(mockT, stringsA, stringsB)
		}
	})

	b.Run("generic/strings_5", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			ElementsMatchT(mockT, stringsA, stringsB)
		}
	})
}

// BenchmarkNotElementsMatch compares NotElementsMatch (reflect) vs NotElementsMatchT (generic).
func BenchmarkNotElementsMatch(b *testing.B) {
	mockT := &mockT{}

	// Slices that don't match
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{1, 2, 3, 4, 6}

	b.Run("reflect", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			NotElementsMatch(mockT, sliceA, sliceB)
		}
	})

	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			NotElementsMatchT(mockT, sliceA, sliceB)
		}
	})
}

// BenchmarkPositive compares Positive (reflect) vs PositiveT (generic).
func BenchmarkPositive(b *testing.B) {
	mockT := &mockT{}

	b.Run("reflect/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			Positive(mockT, 42)
		}
	})

	b.Run("generic/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			PositiveT(mockT, 42)
		}
	})

	b.Run("reflect/float64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			Positive(mockT, 42.5)
		}
	})

	b.Run("generic/float64", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			PositiveT(mockT, 42.5)
		}
	})
}

// BenchmarkNegative compares Negative (reflect) vs NegativeT (generic).
func BenchmarkNegative(b *testing.B) {
	mockT := &mockT{}

	b.Run("reflect/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			Negative(mockT, -42)
		}
	})

	b.Run("generic/int", func(b *testing.B) {
		b.ReportAllocs()
		for b.Loop() {
			NegativeT(mockT, -42)
		}
	})
}
