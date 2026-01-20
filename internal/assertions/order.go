// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"reflect"
	"slices"
)

// IsIncreasing asserts that the collection is strictly increasing.
//
// # Usage
//
//	assertions.IsIncreasing(t, []int{1, 2, 3})
//	assertions.IsIncreasing(t, []float{1, 2})
//	assertions.IsIncreasing(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 2, 3}
//	failure: []int{1, 1, 2}
func IsIncreasing(t T, collection any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	values, ok, err := isStrictlyOrdered(collection, false)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}
	if !ok {
		return Fail(t, fmt.Sprintf("\"%v\" is not less than \"%v\"", values...), msgAndArgs...)
	}

	return true
}

// IsIncreasingT asserts that a slice of [Ordered] is strictly increasing.
//
// # Usage
//
//	assertions.IsIncreasingT(t, []int{1, 2, 3})
//	assertions.IsIncreasingT(t, []float{1, 2})
//	assertions.IsIncreasingT(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 2, 3}
//	failure: []int{1, 1, 2}
func IsIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	isIncreasing := slices.IsSortedFunc(collection, compareStrictOrdered)
	if !isIncreasing {
		return Fail(t, "should be increasing", msgAndArgs...)
	}

	return true
}

// SortedT asserts that the slice of [Ordered] is sorted (i.e. non-strictly increasing).
//
// Unlike [IsIncreasingT], it accepts elements to be equal.
//
// # Usage
//
//	assertions.SortedT(t, []int{1, 2, 3})
//	assertions.SortedT(t, []float{1, 2})
//	assertions.SortedT(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 1, 3}
//	failure: []int{1, 4, 2}
func SortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	isSorted := slices.IsSortedFunc(collection, compareOrdered)
	if !isSorted {
		return Fail(t, "should be sorted", msgAndArgs...)
	}

	return true
}

// NotSortedT asserts that the slice of [Ordered] is NOT sorted (i.e. non-strictly increasing).
//
// Unlike [IsDecreasingT], it accepts slices that are neither increasing nor decreasing.
//
// # Usage
//
//	assertions.NotSortedT(t, []int{3, 2, 3})
//	assertions.NotSortedT(t, []float{2, 1})
//	assertions.NotSortedT(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{3, 1, 3}
//	failure: []int{1, 4, 8}
func NotSortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	isSorted := slices.IsSortedFunc(collection, compareOrdered)
	if isSorted {
		return Fail(t, "should not be sorted", msgAndArgs...)
	}

	return true
}

// IsNonIncreasing asserts that the collection is not increasing.
//
// # Usage
//
//	assertions.IsNonIncreasing(t, []int{2, 1, 1})
//	assertions.IsNonIncreasing(t, []float{2, 1})
//	assertions.IsNonIncreasing(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{2, 1, 1}
//	failure: []int{1, 2, 3}
func IsNonIncreasing(t T, collection any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	_, ok, err := isStrictlyOrdered(collection, false)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}
	if !ok {
		return true
	}

	return Fail(t, "should not be increasing", msgAndArgs...)
}

// IsNonIncreasingT asserts that a slice of [Ordered] is NOT strictly increasing.
//
// # Usage
//
//	assertions.IsNonIncreasing(t, []int{2, 1, 1})
//	assertions.IsNonIncreasing(t, []float{2, 1})
//	assertions.IsNonIncreasing(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{2, 1, 1}
//	failure: []int{1, 2, 3}
func IsNonIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	isIncreasing := slices.IsSortedFunc(collection, compareStrictOrdered)
	if isIncreasing {
		return Fail(t, "should not be increasing", msgAndArgs...)
	}

	return true
}

// IsDecreasing asserts that the collection is strictly decreasing.
//
// # Usage
//
//	assertions.IsDecreasing(t, []int{2, 1, 0})
//	assertions.IsDecreasing(t, []float{2, 1})
//	assertions.IsDecreasing(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{3, 2, 1}
//	failure: []int{1, 2, 3}
func IsDecreasing(t T, collection any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	values, ok, err := isStrictlyOrdered(collection, true)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}
	if !ok {
		values = append(values, msgAndArgs...)
		return Fail(t, fmt.Sprintf("\"%v\" is not greater than \"%v\"", values...), msgAndArgs...)
	}

	return true
}

// IsDecreasingT asserts that a slice of [Ordered] is strictly decreasing.
//
// # Usage
//
//	assertions.IsDecreasingT(t, []int{2, 1, 0})
//	assertions.IsDecreasingT(t, []float{2, 1})
//	assertions.IsDecreasingT(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{3, 2, 1}
//	failure: []int{1, 2, 3}
func IsDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	isDecreasing := slices.IsSortedFunc(collection, reverseCompareStrictOrdered)
	if !isDecreasing {
		return Fail(t, "should be decreasing", msgAndArgs...)
	}

	return true
}

// IsNonDecreasing asserts that the collection is not strictly decreasing.
//
// # Usage
//
//	assertions.IsNonDecreasing(t, []int{1, 1, 2})
//	assertions.IsNonDecreasing(t, []float{1, 2})
//	assertions.IsNonDecreasing(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 1, 2}
//	failure: []int{2, 1, 0}
func IsNonDecreasing(t T, collection any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	_, ok, err := isStrictlyOrdered(collection, true)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}
	if !ok {
		return true
	}

	return Fail(t, "should not be decreasing", msgAndArgs...)
}

// IsNonDecreasingT asserts that a slice of [Ordered] is not decreasing.
//
// # Usage
//
//	assertions.IsNonDecreasingT(t, []int{1, 1, 2})
//	assertions.IsNonDecreasingT(t, []float{1, 2})
//	assertions.IsNonDecreasingT(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 1, 2}
//	failure: []int{2, 1, 0}
func IsNonDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}

	isDecreasing := slices.IsSortedFunc(collection, reverseCompareStrictOrdered)
	if isDecreasing {
		return Fail(t, "should not be decreasing", msgAndArgs...)
	}

	return true
}

// isStrictlyOrdered checks that collection contains orderable elements, which are strictly ordered.
//
// It returns an error if the object can't be ordered.
// When not strictly ordered, it returns the first 2 offending values found.
func isStrictlyOrdered(object any, reverseOrder bool) ([]any, bool, error) {
	objKind := reflect.TypeOf(object).Kind()
	if objKind != reflect.Slice && objKind != reflect.Array {
		return nil, false, fmt.Errorf("object %T is not an ordered collection", object)
	}

	objValue := reflect.ValueOf(object)
	objLen := objValue.Len()

	if objLen <= 1 {
		return nil, true, nil
	}

	value := objValue.Index(0)
	if !value.CanInterface() {
		// this should not be possible with current relect, since values are retrieved from an array or slice, not a struct
		panic(fmt.Errorf("internal error: can't resolve Interface() for value %v", value))
	}
	valueInterface := value.Interface()
	firstValueKind := value.Kind()

	for i := 1; i < objLen; i++ {
		prevValue := value
		prevValueInterface := valueInterface

		value = objValue.Index(i)
		if !value.CanInterface() {
			panic(fmt.Errorf("internal error: can't resolve Interface() for value %v", value))
		}
		valueInterface = value.Interface()

		compareResult, isComparable := compare(prevValueInterface, valueInterface, firstValueKind)

		if !isComparable {
			return nil, false, fmt.Errorf(`cannot compare type "%T" and "%T"`, value, prevValue)
		}

		if (!reverseOrder && compareResult != -1) || (reverseOrder && compareResult != 1) {
			return []any{prevValueInterface, valueInterface}, false, nil
		}
	}

	return nil, true, nil
}

func compareStrictOrdered[E Ordered](a, b E) int {
	v := compareOrderedWithAny[E](a, b)
	if v == 0 {
		return -1
	}

	return v
}

func compareOrdered[E Ordered](a, b E) int {
	v := compareOrderedWithAny[E](a, b)

	return v
}

func reverseCompareStrictOrdered[E Ordered](a, b E) int {
	v := compareOrderedWithAny[E](b, a)
	if v == 0 {
		return -1
	}

	return v
}
