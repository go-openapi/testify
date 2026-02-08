// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"testing"
	"time"
)

func TestCompareGreaterAndLess(t *testing.T) {
	t.Parallel()

	// Unified tests with all comparison functions, reflection-based or generic
	for tc := range comparisonCases() {
		t.Run(tc.name+"/unified", testAllComparison(tc))
	}
}

func TestCompareGreaterAndLessT(t *testing.T) {
	t.Parallel()

	// Unified tests with all comparison functions
	for tc := range comparisonCases() {
		t.Run(tc.name, func(t *testing.T) {
			// Dispatch to type-specific test based on the type of tc.less
			switch tc.less.(type) {
			case string:
				testAllComparisonT[string](tc)(t)
			case int:
				testAllComparisonT[int](tc)(t)
			case int8:
				testAllComparisonT[int8](tc)(t)
			case int16:
				testAllComparisonT[int16](tc)(t)
			case int32:
				testAllComparisonT[int32](tc)(t)
			case int64:
				testAllComparisonT[int64](tc)(t)
			case uint:
				testAllComparisonT[uint](tc)(t)
			case uint8:
				testAllComparisonT[uint8](tc)(t)
			case uint16:
				testAllComparisonT[uint16](tc)(t)
			case uint32:
				testAllComparisonT[uint32](tc)(t)
			case uint64:
				testAllComparisonT[uint64](tc)(t)
			case float32:
				testAllComparisonT[float32](tc)(t)
			case float64:
				testAllComparisonT[float64](tc)(t)
			case uintptr:
				testAllComparisonT[uintptr](tc)(t)
			case time.Time:
				testAllComparisonT[time.Time](tc)(t)
			case []byte:
				testAllComparisonT[[]byte](tc)(t)
			default:
				// Custom types (like redefined uintptr) - skip, they're tested separately
				t.Logf("%s: custom types tested separately (got: %T)", t.Name(), tc.less)
			}
		})
	}

	// Additional type-specific tests
	t.Run("custom int type", testGreaterTCustomInt())
}

func TestComparePositiveT(t *testing.T) {
	t.Parallel()

	// Unified tests with both Positive and Negative functions
	for tc := range signCases() {
		t.Run(tc.name, func(t *testing.T) {
			// Dispatch to type-specific test based on the type of tc.positive
			switch tc.positive.(type) {
			case int:
				testAllSignT[int](tc)(t)
			case int8:
				testAllSignT[int8](tc)(t)
			case int16:
				testAllSignT[int16](tc)(t)
			case int32:
				testAllSignT[int32](tc)(t)
			case int64:
				testAllSignT[int64](tc)(t)
			case float32:
				testAllSignT[float32](tc)(t)
			case float64:
				testAllSignT[float64](tc)(t)
			}
		})
	}
}

func TestComparePositive(t *testing.T) {
	t.Parallel()

	// Unified tests with both Positive and Negative functions
	for tc := range signCases() {
		t.Run(tc.name+"/unified", testAllSign(tc))
	}
}

func TestCompareErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, compareFailCases())
}

// genericTestCase wraps a test function with its name for table-driven tests of generic functions.
// Kept for compatibility with existing special-case tests.
type genericTestCase struct {
	name string
	test func(*testing.T)
}

// Special-case test helpers for types that need custom handling

func testGreaterTCustomInt() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		type MyInt int
		if !GreaterT(mock, MyInt(2), MyInt(1)) {
			t.Error("expected GreaterT(2, 1) to pass")
		}
		if GreaterT(mock, MyInt(1), MyInt(2)) {
			t.Error("expected GreaterT(1, 2) to fail")
		}
	}
}

// Unified test helpers for comparison functions

// comparisonTestCase represents a test case for comparison functions.
type comparisonTestCase struct {
	name    string
	less    any
	greater any
	equal   bool // if true, less == greater (for testing equal values)
}

// comparisonCases returns unified test data for all comparison functions.
func comparisonCases() iter.Seq[comparisonTestCase] {
	type redefinedUintptr uintptr

	return slices.Values([]comparisonTestCase{
		// Strict inequality cases
		{name: "string", less: "a", greater: "b", equal: false},
		{name: "int", less: int(1), greater: int(2), equal: false},
		{name: "int8", less: int8(1), greater: int8(2), equal: false},
		{name: "int16", less: int16(1), greater: int16(2), equal: false},
		{name: "int32", less: int32(1), greater: int32(2), equal: false},
		{name: "int64", less: int64(1), greater: int64(2), equal: false},
		{name: "uint", less: uint(1), greater: uint(2), equal: false},
		{name: "uint8", less: uint8(1), greater: uint8(2), equal: false},
		{name: "uint16", less: uint16(1), greater: uint16(2), equal: false},
		{name: "uint32", less: uint32(1), greater: uint32(2), equal: false},
		{name: "uint64", less: uint64(1), greater: uint64(2), equal: false},
		{name: "float32", less: float32(1.23), greater: float32(2.34), equal: false},
		{name: "float64", less: float64(1.23), greater: float64(2.34), equal: false},
		{name: "uintptr", less: uintptr(1), greater: uintptr(2), equal: false},
		{name: "uintptr/9-10", less: uintptr(9), greater: uintptr(10), equal: false},
		{name: "redefined-uintptr", less: redefinedUintptr(9), greater: redefinedUintptr(10), equal: false},
		{name: "time.Time", less: time.Time{}, greater: time.Time{}.Add(time.Hour), equal: false},
		{name: "[]byte", less: []byte{1, 1}, greater: []byte{1, 2}, equal: false},

		// Equality cases
		{name: "string/equal", less: "a", greater: "a", equal: true},
		{name: "int/equal", less: int(1), greater: int(1), equal: true},
		{name: "int8/equal", less: int8(1), greater: int8(1), equal: true},
		{name: "int16/equal", less: int16(1), greater: int16(1), equal: true},
		{name: "int32/equal", less: int32(1), greater: int32(1), equal: true},
		{name: "int64/equal", less: int64(1), greater: int64(1), equal: true},
		{name: "uint/equal", less: uint(1), greater: uint(1), equal: true},
		{name: "uint8/equal", less: uint8(1), greater: uint8(1), equal: true},
		{name: "uint16/equal", less: uint16(1), greater: uint16(1), equal: true},
		{name: "uint32/equal", less: uint32(1), greater: uint32(1), equal: true},
		{name: "uint64/equal", less: uint64(1), greater: uint64(1), equal: true},
		{name: "float32/equal", less: float32(1.23), greater: float32(1.23), equal: true},
		{name: "float64/equal", less: float64(1.23), greater: float64(1.23), equal: true},
		{name: "uintptr/equal", less: uintptr(1), greater: uintptr(1), equal: true},
		{name: "time.Time/equal", less: time.Time{}, greater: time.Time{}, equal: true},
		{name: "[]byte/equal", less: []byte{1, 1}, greater: []byte{1, 1}, equal: true},
	})
}

// testAllComparison tests all four comparison functions with the same test data.
func testAllComparison(tc comparisonTestCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if tc.equal {
			// For equal values:
			// - Greater and Less should return false
			// - GreaterOrEqual and LessOrEqual should return true
			t.Run("with equal values", func(t *testing.T) {
				t.Run("Greater should fail", testComparison(Greater, tc.less, tc.greater, false))
				t.Run("GreaterOrEqual should pass", testComparison(GreaterOrEqual, tc.less, tc.greater, true))
				t.Run("Less should fail", testComparison(Less, tc.less, tc.greater, false))
				t.Run("LessOrEqual should pass", testComparison(LessOrEqual, tc.less, tc.greater, true))
			})

			return
		}

		// For strict inequality:
		// Test both directions to verify the inverse relationships
		t.Run("with strict inequality", func(t *testing.T) {
			t.Run("Greater", func(t *testing.T) {
				t.Run("should pass (greater > less)", testComparison(Greater, tc.greater, tc.less, true))
				t.Run("should fail (less > greater)", testComparison(Greater, tc.less, tc.greater, false))
			})
			t.Run("GreaterOrEqual", func(t *testing.T) {
				t.Run("should pass (greater >= less)", testComparison(GreaterOrEqual, tc.greater, tc.less, true))
				t.Run("should fail (less >= greater)", testComparison(GreaterOrEqual, tc.less, tc.greater, false))
			})
			t.Run("Less", func(t *testing.T) {
				t.Run("should pass (less < greater)", testComparison(Less, tc.less, tc.greater, true))
				t.Run("should fail (greater < less)", testComparison(Less, tc.greater, tc.less, false))
			})
			t.Run("LessOrEqual", func(t *testing.T) {
				t.Run("should pass (less <= greater)", testComparison(LessOrEqual, tc.less, tc.greater, true))
				t.Run("should fail (greater <= less)", testComparison(LessOrEqual, tc.greater, tc.less, false))
			})
		})
	}
}

// testComparison is a helper that tests a comparison function.
func testComparison(cmp func(T, any, any, ...any) bool, e1, e2 any, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := cmp(mock, e1, e2)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

// testAllComparisonT tests all four generic comparison functions with the same test data.
func testAllComparisonT[V Ordered](tc comparisonTestCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		// Type assert the values
		less, ok1 := tc.less.(V)
		greater, ok2 := tc.greater.(V)
		if !ok1 || !ok2 {
			t.Fatalf("type mismatch in testcase: expected %T, got less=%T, greater=%T", *new(V), tc.less, tc.greater)
		}

		if tc.equal {
			// For equal values:
			// - GreaterT and LessT should return false
			// - GreaterOrEqualT and LessOrEqualT should return true
			t.Run("with equal values", func(t *testing.T) {
				t.Run("GreaterT should fail", testComparisonT(GreaterT[V], less, greater, false))
				t.Run("GreaterOrEqualT should pass", testComparisonT(GreaterOrEqualT[V], less, greater, true))
				t.Run("LessT should fail", testComparisonT(LessT[V], less, greater, false))
				t.Run("LessOrEqualT should pass", testComparisonT(LessOrEqualT[V], less, greater, true))
			})

			return
		}

		// For strict inequality:
		// Test both directions to verify the inverse relationships
		t.Run("with strict inequality", func(t *testing.T) {
			t.Run("GreaterT", func(t *testing.T) {
				t.Run("should pass (greater > less)", testComparisonT(GreaterT[V], greater, less, true))
				t.Run("should fail (less > greater)", testComparisonT(GreaterT[V], less, greater, false))
			})
			t.Run("GreaterOrEqualT", func(t *testing.T) {
				t.Run("should pass (greater >= less)", testComparisonT(GreaterOrEqualT[V], greater, less, true))
				t.Run("should fail (less >= greater)", testComparisonT(GreaterOrEqualT[V], less, greater, false))
			})
			t.Run("LessT", func(t *testing.T) {
				t.Run("should pass (less < greater)", testComparisonT(LessT[V], less, greater, true))
				t.Run("should fail (greater < less)", testComparisonT(LessT[V], greater, less, false))
			})
			t.Run("LessOrEqualT", func(t *testing.T) {
				t.Run("should pass (less <= greater)", testComparisonT(LessOrEqualT[V], less, greater, true))
				t.Run("should fail (greater <= less)", testComparisonT(LessOrEqualT[V], greater, less, false))
			})
		})
	}
}

// testComparisonT is a helper that tests a generic comparison function.
func testComparisonT[V Ordered](cmp func(T, V, V, ...any) bool, e1, e2 V, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := cmp(mock, e1, e2)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

// Unified test helpers for sign functions (Positive/Negative)

// signTestCase represents a test case for Positive/Negative functions.
type signTestCase struct {
	name     string
	positive any
	negative any
	isZero   bool // if true, positive is zero (both Positive and Negative should fail)
}

// signCases returns unified test data for Positive/Negative functions.
func signCases() iter.Seq[signTestCase] {
	return slices.Values([]signTestCase{
		{name: "int", positive: int(1), negative: int(-1), isZero: false},
		{name: "int8", positive: int8(1), negative: int8(-1), isZero: false},
		{name: "int16", positive: int16(1), negative: int16(-1), isZero: false},
		{name: "int32", positive: int32(1), negative: int32(-1), isZero: false},
		{name: "int64", positive: int64(1), negative: int64(-1), isZero: false},
		{name: "float32", positive: float32(1.5), negative: float32(-1.5), isZero: false},
		{name: "float64", positive: float64(1.5), negative: float64(-1.5), isZero: false},

		// Zero cases - both Positive and Negative should fail
		{name: "int/zero", positive: int(0), negative: int(0), isZero: true},
		{name: "float64/zero", positive: float64(0.0), negative: float64(0.0), isZero: true},
	})
}

// testAllSign tests both Positive and Negative functions with the same test data.
func testAllSign(tc signTestCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if tc.isZero {
			// Zero should fail both Positive and Negative
			t.Run("zero should fail both", func(t *testing.T) {
				t.Run("Positive should fail", testSign(Positive, tc.positive, false))
				t.Run("Negative should fail", testSign(Negative, tc.positive, false))
			})

			return
		}

		// Test positive and negative values
		t.Run("with positive/negative values", func(t *testing.T) {
			t.Run("Positive", func(t *testing.T) {
				t.Run("should pass (positive value)", testSign(Positive, tc.positive, true))
				t.Run("should fail (negative value)", testSign(Positive, tc.negative, false))
			})
			t.Run("Negative", func(t *testing.T) {
				t.Run("should pass (negative value)", testSign(Negative, tc.negative, true))
				t.Run("should fail (positive value)", testSign(Negative, tc.positive, false))
			})
		})
	}
}

// testSign is a helper that tests a sign function.
func testSign(sign func(T, any, ...any) bool, e any, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := sign(mock, e)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

// testAllSignT tests both PositiveT and NegativeT functions with the same test data.
func testAllSignT[V SignedNumeric](tc signTestCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		// Type assert the values
		positive, ok1 := tc.positive.(V)
		negative, ok2 := tc.negative.(V)
		if !ok1 || !ok2 {
			t.Fatalf("type mismatch: expected %T, got positive=%T, negative=%T", *new(V), tc.positive, tc.negative)
		}

		if tc.isZero {
			// Zero should fail both PositiveT and NegativeT
			t.Run("zero should fail both", func(t *testing.T) {
				t.Run("PositiveT should fail", testSignT(PositiveT[V], positive, false))
				t.Run("NegativeT should fail", testSignT(NegativeT[V], positive, false))
			})

			return
		}

		// Test positive and negative values
		t.Run("with positive/negative values", func(t *testing.T) {
			t.Run("PositiveT", func(t *testing.T) {
				t.Run("should pass (positive value)", testSignT(PositiveT[V], positive, true))
				t.Run("should fail (negative value)", testSignT(PositiveT[V], negative, false))
			})
			t.Run("NegativeT", func(t *testing.T) {
				t.Run("should pass (negative value)", testSignT(NegativeT[V], negative, true))
				t.Run("should fail (positive value)", testSignT(NegativeT[V], positive, false))
			})
		})
	}
}

// testSignT is a helper that tests a generic sign function.
func testSignT[V SignedNumeric](sign func(T, V, ...any) bool, e V, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := sign(mock, e)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

// ============================================================================
// TestCompareErrorMessages
// ============================================================================

func compareFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "Greater/int",
			assertion:    func(t T) bool { return Greater(t, 1, 2) },
			wantContains: []string{`"1" is not greater than "2"`},
		},
		{
			name:         "GreaterOrEqual/int",
			assertion:    func(t T) bool { return GreaterOrEqual(t, 1, 2) },
			wantContains: []string{`"1" is not greater than or equal to "2"`},
		},
		{
			name:         "Less/int",
			assertion:    func(t T) bool { return Less(t, 2, 1) },
			wantContains: []string{`"2" is not less than "1"`},
		},
		{
			name:         "LessOrEqual/int",
			assertion:    func(t T) bool { return LessOrEqual(t, 2, 1) },
			wantContains: []string{`"2" is not less than or equal to "1"`},
		},
		{
			name:         "Positive/negative-value",
			assertion:    func(t T) bool { return Positive(t, -1) },
			wantContains: []string{`"-1" is not positive`},
		},
		{
			name:         "Negative/positive-value",
			assertion:    func(t T) bool { return Negative(t, 1) },
			wantContains: []string{`"1" is not negative`},
		},
	})
}
