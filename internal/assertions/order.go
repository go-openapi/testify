// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"reflect"
)

// IsIncreasing asserts that the collection is increasing.
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

// IsDecreasing asserts that the collection is decreasing.
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

// IsNonDecreasing asserts that the collection is not decreasing.
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
//	failure: []int{2, 1, 1}
func IsNonDecreasing(t T, object any, msgAndArgs ...any) bool {
	// Domain: ordering
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return isOrdered(t, object, []compareResult{compareLess, compareEqual}, "\"%v\" is not less than or equal to \"%v\"", msgAndArgs...)
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
