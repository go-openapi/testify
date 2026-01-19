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
func IsIncreasing(t T, object any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return isOrdered(t, object, []compareResult{compareLess}, "\"%v\" is not less than \"%v\"", msgAndArgs...)
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
func IsNonIncreasing(t T, object any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return isOrdered(t, object, []compareResult{compareEqual, compareGreater}, "\"%v\" is not greater than or equal to \"%v\"", msgAndArgs...)
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
func IsDecreasing(t T, object any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return isOrdered(t, object, []compareResult{compareGreater}, "\"%v\" is not greater than \"%v\"", msgAndArgs...)
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
func IsNonDecreasing(t T, object any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return isOrdered(t, object, []compareResult{compareLess, compareEqual}, "\"%v\" is not less than or equal to \"%v\"", msgAndArgs...)
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

// isOrdered checks that collection contains orderable elements.
func isOrdered(t T, object any, allowedComparesResults []compareResult, failMessage string, msgAndArgs ...any) bool {
	objKind := reflect.TypeOf(object).Kind()
	if objKind != reflect.Slice && objKind != reflect.Array {
		return Fail(t, fmt.Sprintf("object %T is not an ordered collection", object), msgAndArgs...)
	}

	objValue := reflect.ValueOf(object)
	objLen := objValue.Len()

	if objLen <= 1 {
		return true
	}

	value := objValue.Index(0)
	valueInterface := value.Interface()
	firstValueKind := value.Kind()

	for i := 1; i < objLen; i++ {
		prevValue := value
		prevValueInterface := valueInterface

		value = objValue.Index(i)
		valueInterface = value.Interface()

		compareResult, isComparable := compare(prevValueInterface, valueInterface, firstValueKind)

		if !isComparable {
			return Fail(t, fmt.Sprintf(`Can not compare type "%T" and "%T"`, value, prevValue), msgAndArgs...)
		}

		if !containsValue(allowedComparesResults, compareResult) {
			return Fail(t, fmt.Sprintf(failMessage, prevValue, value), msgAndArgs...)
		}
	}

	return true
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
