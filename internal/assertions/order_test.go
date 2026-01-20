// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"fmt"
	"iter"
	"slices"
	"strings"
	"testing"
	"time"
)

func TestOrderErrorMessages(t *testing.T) {
	// Test reflection-based assertions

	t.Run("with error messages", func(t *testing.T) {
		t.Parallel()

		const (
			format         = "format %s %x"
			arg1           = "this"
			arg2           = 0xc001
			expectedOutput = "format this c001\n"
		)

		msgAndArgs := []any{format, arg1, arg2}
		collection := []int{1, 2, 1} // is neither increasing nor decreasing
		increasingCollection := []int{1, 2, 3}
		decreasingCollection := []int{3, 2, 1}

		funcs := []struct {
			name       string
			fn         func(T, any, ...any) bool
			collection []int
		}{
			{"IsIncreasing", IsIncreasing, collection},
			{"IsNonIncreasing", IsNonIncreasing, increasingCollection},
			{"IsDecreasing", IsDecreasing, collection},
			{"IsNonDecreasing", IsNonDecreasing, decreasingCollection},
		}

		for _, fn := range funcs {
			t.Run(fn.name, func(t *testing.T) {
				t.Parallel()

				mock := &outputT{buf: bytes.NewBuffer(nil)}
				result := fn.fn(mock, fn.collection, msgAndArgs...)
				if result {
					t.Errorf("expected ordering assertion %q to fail on %v", fn.name, fn.collection)

					return
				}

				if !strings.Contains(mock.buf.String(), expectedOutput) {
					t.Errorf("expected error message to contain: %s but got %q", expectedOutput, mock.buf.String())
				}
			})
		}
	})

	t.Run("with detailed error messages", func(t *testing.T) {
		// Test specific error messages for reflection-based assertions
		t.Parallel()

		testCases := []struct {
			name       string
			fn         func(T, any, ...any) bool
			collection any
			expected   string
		}{
			// IsIncreasing errors
			{"IsIncreasing/string", IsIncreasing, []string{"b", "a"}, `"b" is not less than "a"`},
			{"IsIncreasing/int", IsIncreasing, []int{2, 1}, `"2" is not less than "1"`},
			{"IsIncreasing/int8", IsIncreasing, []int8{2, 1}, `"2" is not less than "1"`},
			{"IsIncreasing/float32", IsIncreasing, []float32{2.34, 1.23}, `"2.34" is not less than "1.23"`},
			{"IsIncreasing/invalid-type", IsIncreasing, struct{}{}, `object struct {} is not an ordered collection`},

			// IsNonIncreasing errors
			{"IsNonIncreasing/string", IsNonIncreasing, []string{"a", "b"}, `should not be increasing`},
			{"IsNonIncreasing/int", IsNonIncreasing, []int{1, 2}, `should not be increasing`},
			{"IsNonIncreasing/float64", IsNonIncreasing, []float64{1.23, 2.34}, `should not be increasing`},

			// IsDecreasing errors
			{"IsDecreasing/string", IsDecreasing, []string{"a", "b"}, `"a" is not greater than "b"`},
			{"IsDecreasing/int", IsDecreasing, []int{1, 2}, `"1" is not greater than "2"`},
			{"IsDecreasing/uint64", IsDecreasing, []uint64{1, 2}, `"1" is not greater than "2"`},

			// IsNonDecreasing errors
			{"IsNonDecreasing/string", IsNonDecreasing, []string{"b", "a"}, `should not be decreasing`},
			{"IsNonDecreasing/int", IsNonDecreasing, []int{2, 1}, `should not be decreasing`},
			{"IsNonDecreasing/float32", IsNonDecreasing, []float32{2.34, 1.23}, `should not be decreasing`},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				out := &outputT{buf: bytes.NewBuffer(nil)}
				result := tc.fn(out, tc.collection)
				if result {
					t.Errorf("expected ordering assertion %q to fail on %v", tc.name, tc.collection)

					return
				}
				if !strings.Contains(out.buf.String(), tc.expected) {
					t.Errorf("expected error message to contain: %s but got %q", tc.expected, out.buf.String())
				}
			})
		}
	})
}

// Test functions for reflection-based and generic assertions

// Unified test for all order assertions, with different input types
//
// NOTE: Unified testing pattern for ordering assertions.
// Test cases are defined with their intrinsic ordering property (kind).
// Expected pass/fail is determined from the kind + assertion semantics.
//
// Unlike the pattern used in string_test.go, we can't easily use a conversion for slices.
// Therefore, we have to resort to a type switch with a fixed known list of tested slice types.
//
// The matrix of expected assertion semantics is defined by [expectedStatusForAssertion].
func TestOrder(t *testing.T) {
	t.Parallel()

	for tc := range unifiedOrderCases() {
		t.Run(tc.name, testAllOrdersWithTypes(tc))
	}
}

func testAllOrdersWithTypes(tc orderTestCase) func(*testing.T) {
	return func(t *testing.T) {
		t.Run("with IsIncreasing", func(t *testing.T) {
			t.Parallel()

			shouldPass := expectedStatusForAssertion(increasingKind, tc.kind)
			t.Run("with reflection", testOrderReflectBased(IsIncreasing, tc.collection, shouldPass))
			if !tc.reflectionOnly {
				t.Run("with generic", testOrderGeneric(increasingKind, tc.collection, shouldPass))
			}
		})

		t.Run("with IsNonIncreasing", func(t *testing.T) {
			t.Parallel()

			shouldPass := expectedStatusForAssertion(notIncreasingKind, tc.kind)
			t.Run("with reflection", testOrderReflectBased(IsNonIncreasing, tc.collection, shouldPass))
			if !tc.reflectionOnly {
				t.Run("with generic", testOrderGeneric(notIncreasingKind, tc.collection, shouldPass))
			}
		})

		t.Run("with IsDecreasing", func(t *testing.T) {
			t.Parallel()

			shouldPass := expectedStatusForAssertion(decreasingKind, tc.kind)
			t.Run("with reflection", testOrderReflectBased(IsDecreasing, tc.collection, shouldPass))
			if !tc.reflectionOnly {
				t.Run("with generic", testOrderGeneric(decreasingKind, tc.collection, shouldPass))
			}
		})

		t.Run("with IsNonDecreasing", func(t *testing.T) {
			t.Parallel()

			shouldPass := expectedStatusForAssertion(notDecreasingKind, tc.kind)
			t.Run("with reflection", testOrderReflectBased(IsNonDecreasing, tc.collection, shouldPass))
			if !tc.reflectionOnly {
				t.Run("with generic", testOrderGeneric(notDecreasingKind, tc.collection, shouldPass))
			}
		})

		if tc.reflectionOnly {
			return
		}

		t.Run("with SortedT", func(t *testing.T) {
			t.Parallel()

			shouldPass := expectedStatusForAssertion(sortedKind, tc.kind)
			t.Run("with generic only", testOrderGeneric(sortedKind, tc.collection, shouldPass))
		})

		t.Run("with NotSortedT", func(t *testing.T) {
			t.Parallel()

			shouldPass := expectedStatusForAssertion(notSortedKind, tc.kind)
			t.Run("with generic only", testOrderGeneric(notSortedKind, tc.collection, shouldPass))
		})
	}
}

// collectionKind represents the ordering property of a collection.
type collectionKind int

const (
	allEqual        collectionKind = iota // all values equal (sorted but not strictly)
	strictlyAsc                           // strictly ascending (each < next)
	strictlyDesc                          // strictly descending (each > next)
	nonStrictlyAsc                        // non-strictly ascending (each <= next, some equal)
	nonStrictlyDesc                       // non-strictly descending (each >= next, some equal)
	unsorted                              // no ordering
	passAll                               // empty or single element collection
	errorCase                             // should fail with error (not panic)
)

type orderAssertionKind int

const (
	increasingKind orderAssertionKind = iota
	notIncreasingKind
	decreasingKind
	notDecreasingKind
	sortedKind
	notSortedKind
)

// orderTestCase represents a test case that can be used for all ordering assertions.
type orderTestCase struct {
	name           string
	collection     any
	kind           collectionKind
	reflectionOnly bool
}

// Unified test cases for all ordering assertions.
type (
	myFloat      float64
	myCollection []myFloat
)

func unifiedOrderCases() iter.Seq[orderTestCase] {
	t0 := time.Now()
	t1 := t0.Add(time.Second)
	t2 := t1.Add(time.Second)

	// Test types for reflection-only edge cases.
	type nonComparableStruct struct {
		Value int
		Data  []int // slices make structs non-comparable
	}

	type structWithUnexportedField struct {
		unexported int
	}

	return slices.Values([]orderTestCase{
		// Edge cases: nil, empty, single element collections
		{"empty/int", []int{}, passAll, false},
		{"nil/int", []int(nil), passAll, false},
		{"single/int", []int{1}, passAll, false},

		// All equal - both non-strict pass, both strict fail, sorted
		{"all-equal/int", []int{2, 2, 2}, allEqual, false},
		{"all-equal/float64", []float64{1.5, 1.5, 1.5}, allEqual, false},
		{"all-equal/~float64", []myFloat{1.5, 1.5, 1.5}, allEqual, false},
		{"all-equal/~[]~float64", myCollection{1.5, 1.5, 1.5}, allEqual, false},
		{"all-equal/string", []string{"a", "a", "a"}, allEqual, false},
		{"all-equal/time.Time", []time.Time{t0, t0, t0}, allEqual, false},
		{"all-equal/[]byte", [][]byte{[]byte("a"), []byte("a"), []byte("a")}, allEqual, false},

		// Strictly ascending - IsIncreasing passes, IsNonDecreasing passes, sorted
		{"strictly-asc/int-short", []int{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/int-long", []int{1, 2, 3, 4, 5}, strictlyAsc, false},
		{"strictly-asc/int8", []int8{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/int16", []int16{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/int32", []int32{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/int64", []int64{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/uint", []uint{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/uint8", []uint8{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/uint16", []uint16{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/uint32", []uint32{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/uint64", []uint64{1, 2, 3}, strictlyAsc, false},
		{"strictly-asc/float32", []float32{1.1, 2.2, 3.3}, strictlyAsc, false},
		{"strictly-asc/float64", []float64{1.1, 2.2, 3.3}, strictlyAsc, false},
		{"strictly-asc/~float64", []myFloat{1.1, 2.2, 3.3}, strictlyAsc, false},
		{"strictly-asc/~[]~float64", myCollection{1.1, 2.2, 3.3}, strictlyAsc, false},
		{"strictly-asc/string", []string{"a", "b", "c"}, strictlyAsc, false},
		{"strictly-asc/time.Time", []time.Time{t0, t1, t2}, strictlyAsc, false},
		{"strictly-asc/[]byte", [][]byte{[]byte("a"), []byte("b"), []byte("c")}, strictlyAsc, false},

		// Strictly descending - IsDecreasing passes, IsNonIncreasing passes, not sorted
		{"strictly-desc/int-short", []int{3, 2, 1}, strictlyDesc, false},
		{"strictly-desc/int-long", []int{5, 4, 3, 2, 1}, strictlyDesc, false},
		{"strictly-desc/float64", []float64{3.3, 2.2, 1.1}, strictlyDesc, false},
		{"strictly-desc/~float64", []myFloat{3.3, 2.2, 1.1}, strictlyDesc, false},
		{"strictly-desc/~[]~float64", myCollection{3.3, 2.2, 1.1}, strictlyDesc, false},
		{"strictly-desc/string", []string{"c", "b", "a"}, strictlyDesc, false},
		{"strictly-desc/time.Time", []time.Time{t2, t1, t0}, strictlyDesc, false},
		{"strictly-desc/[]byte", [][]byte{[]byte("c"), []byte("b"), []byte("a")}, strictlyDesc, false},

		// Non-strictly ascending - sorted, but not strictly (has equal adjacent)
		{"non-strictly-asc/int-with-equal", []int{1, 1, 2, 3}, nonStrictlyAsc, false},
		{"non-strictly-asc/int-with-equal-middle", []int{1, 2, 2, 3}, nonStrictlyAsc, false},
		{"non-strictly-asc/float64", []float64{1.1, 2.2, 2.2, 3.3}, nonStrictlyAsc, false},
		{"non-strictly-asc/~float64", []myFloat{1.1, 1.1, 2.2, 3.3}, nonStrictlyAsc, false},
		{"non-strictly-asc/~[]~float64", myCollection{1.1, 1.1, 2.2, 3.3}, nonStrictlyAsc, false},
		{"non-strictly-asc/string", []string{"a", "a", "b", "c"}, nonStrictlyAsc, false},
		{"non-strictly-asc/time.Time", []time.Time{t0, t0, t1, t2}, nonStrictlyAsc, false},
		{"non-strictly-asc/[]byte", [][]byte{[]byte("a"), []byte("a"), []byte("b"), []byte("c")}, nonStrictlyAsc, false},

		// Non-strictly descending - not sorted, but consistently >= (has equal adjacent)
		{"non-strictly-desc/int-with-equal", []int{3, 2, 2, 1}, nonStrictlyDesc, false},
		{"non-strictly-desc/int-with-equal-start", []int{3, 3, 2, 1}, nonStrictlyDesc, false},
		{"non-strictly-desc/float64", []float64{3.3, 2.2, 2.2, 1.1}, nonStrictlyDesc, false},
		{"non-strictly-desc/~float64", []myFloat{3.3, 2.2, 2.2, 1.1}, nonStrictlyDesc, false},
		{"non-strictly-desc/~[]~float64", myCollection{3.3, 2.2, 2.2, 1.1}, nonStrictlyDesc, false},
		{"non-strictly-desc/string", []string{"c", "b", "b", "a"}, nonStrictlyDesc, false},
		{"non-strictly-desc/time.Time", []time.Time{t2, t1, t1, t0}, nonStrictlyDesc, false},
		{"non-strictly-desc/[]byte", [][]byte{[]byte("c"), []byte("b"), []byte("b"), []byte("a")}, nonStrictlyDesc, false},

		// Unsorted - no ordering pattern
		{"unsorted/int-mixed", []int{1, 4, 2}, unsorted, false},
		{"unsorted/int-up-down-up", []int{1, 3, 2, 4}, unsorted, false},
		{"unsorted/float64", []float64{1.1, 3.3, 2.2}, unsorted, false},
		{"unsorted/~float64", []myFloat{1.1, 3.3, 2.2}, unsorted, false},
		{"unsorted/~[]~float64", myCollection{1.1, 3.3, 2.2}, unsorted, false},
		{"unsorted/string", []string{"b", "a", "c"}, unsorted, false},
		{"unsorted/time.Time", []time.Time{t1, t0, t2}, unsorted, false},
		{"unsorted/[]byte", [][]byte{[]byte("b"), []byte("a"), []byte("c")}, unsorted, false},

		// Reflection-only edge cases

		// Case 1: Object is not a slice or array (should error)
		{"error/not-a-collection", nonComparableStruct{Value: 1, Data: []int{1}}, errorCase, true},

		// Case 2: Slice of non-comparable elements (should error)
		{"error/non-comparable-elements", []nonComparableStruct{
			{Value: 1, Data: []int{1}},
			{Value: 2, Data: []int{2}},
		}, errorCase, true},

		// Case 3: Slice with unexported fields (triggers panic bug in isStrictlyOrdered)
		{"panic/unexported-fields", []structWithUnexportedField{
			{unexported: 1},
			{unexported: 2},
		}, errorCase, true},
	})
}

// Determine the expected pass/fail status for each assertion based on ordering kind.
func expectedStatusForAssertion(assertionKind orderAssertionKind, kind collectionKind) bool {
	// Error cases always fail (return false)
	if kind == errorCase {
		return false
	}

	switch assertionKind {
	case increasingKind:
		// IsIncreasing: strictly ascending only
		return kind == strictlyAsc || kind == passAll
	case notIncreasingKind:
		return kind != strictlyAsc && kind != passAll
	case decreasingKind:
		// IsDecreasing: strictly descending only
		return kind == strictlyDesc || kind == passAll
	case notDecreasingKind:
		return kind != strictlyDesc && kind != passAll
	case sortedKind:
		// SortedT: passes for sorted (non-strictly ascending, allows equal)
		return kind == allEqual || kind == strictlyAsc || kind == nonStrictlyAsc || kind == passAll
	case notSortedKind:
		// NotSortedT: inverse of SortedT
		return kind != allEqual && kind != strictlyAsc && kind != nonStrictlyAsc && kind != passAll
	default:
		panic(fmt.Errorf("test case configuration error: invalid orderAssertionKind: %d", assertionKind))
	}
}

func testOrderReflectBased(orderAssertion func(T, any, ...any) bool, collection any, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := orderAssertion(mock, collection)

		if shouldPass {
			t.Run("should pass", func(t *testing.T) {
				if !result || mock.Failed() {
					t.Errorf("expected to pass")
				}
			})

			return
		}

		t.Run("should fail", func(t *testing.T) {
			if result || !mock.Failed() {
				t.Errorf("expected to fail")
			}
		})
	}
}

func testOrderGeneric(assertionKind orderAssertionKind, collection any, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := testOrderAssertionResult(mock, assertionKind, collection)

		if shouldPass {
			t.Run("should pass", func(t *testing.T) {
				if !result || mock.Failed() {
					t.Errorf("expected to pass")
				}
			})

			return
		}

		t.Run("should fail", func(t *testing.T) {
			if result || !mock.Failed() {
				t.Errorf("expected to fail")
			}
		})
	}
}

func testOrderAssertionResult(mock T, assertionKind orderAssertionKind, collection any) bool {
	// Type switch to call the appropriate generic function.
	//
	// This switch doesn't cover ALL variants of ~[]Ordered but is deemed sufficient for the purpose
	// of testing ordering.
	switch coll := collection.(type) {
	case []int:
		return testGenericAssertion(mock, assertionKind, coll)
	case []int8:
		return testGenericAssertion(mock, assertionKind, coll)
	case []int16:
		return testGenericAssertion(mock, assertionKind, coll)
	case []int32:
		return testGenericAssertion(mock, assertionKind, coll)
	case []int64:
		return testGenericAssertion(mock, assertionKind, coll)
	case []uint:
		return testGenericAssertion(mock, assertionKind, coll)
	case []uint8:
		return testGenericAssertion(mock, assertionKind, coll)
	case []uint16:
		return testGenericAssertion(mock, assertionKind, coll)
	case []uint32:
		return testGenericAssertion(mock, assertionKind, coll)
	case []uint64:
		return testGenericAssertion(mock, assertionKind, coll)
	case []uintptr:
		return testGenericAssertion(mock, assertionKind, coll)
	case []float32:
		return testGenericAssertion(mock, assertionKind, coll)
	case []float64:
		return testGenericAssertion(mock, assertionKind, coll)
	case []myFloat:
		return testGenericAssertion(mock, assertionKind, coll)
	case myCollection:
		return testGenericAssertion(mock, assertionKind, coll)
	case [][]byte:
		return testGenericAssertion(mock, assertionKind, coll)
	case []string:
		return testGenericAssertion(mock, assertionKind, coll)
	case []time.Time:
		return testGenericAssertion(mock, assertionKind, coll)
	default:
		panic(fmt.Errorf("internal test error: unsupported collection type in test suite: %T", coll))
	}
}

func testGenericAssertion[Collection ~[]E, E Ordered](mock T, assertionKind orderAssertionKind, collection Collection) bool {
	switch assertionKind {
	case increasingKind:
		return IsIncreasingT(mock, collection)
	case notIncreasingKind:
		return IsNonIncreasingT(mock, collection)
	case decreasingKind:
		return IsDecreasingT(mock, collection)
	case notDecreasingKind:
		return IsNonDecreasingT(mock, collection)
	case sortedKind:
		return SortedT(mock, collection)
	case notSortedKind:
		return NotSortedT(mock, collection)
	default:
		panic(fmt.Errorf("test case configuration error: invalid orderAssertionKind: %d", assertionKind))
	}
}
