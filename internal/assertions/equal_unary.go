// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"reflect"
)

// Nil asserts that the specified object is nil.
//
// # Usage
//
//	assertions.Nil(t, err)
//
// # Examples
//
//	success: nil
//	failure: "not nil"
func Nil(t T, object any, msgAndArgs ...any) bool {
	// Domain: equality
	if isNil(object) {
		return true
	}
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return Fail(t, "Expected nil, but got: "+truncatingFormat("%#v", object), msgAndArgs...)
}

// NotNil asserts that the specified object is not nil.
//
// # Usage
//
// assertions.NotNil(t, err)
//
// # Examples
//
//	success: "not nil"
//	failure: nil
func NotNil(t T, object any, msgAndArgs ...any) bool {
	// Domain: equality
	if !isNil(object) {
		return true
	}
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return Fail(t, "Expected value not to be nil.", msgAndArgs...)
}

// Empty asserts that the given value is "empty".
//
// Zero values are "empty".
//
// Arrays are "empty" if every element is the zero value of the type (stricter than "empty").
//
// Slices, maps and channels with zero length are "empty".
//
// Pointer values are "empty" if the pointer is nil or if the pointed value is "empty".
//
// # Usage
//
//	assertions.Empty(t, obj)
//
// # Examples
//
//	success: ""
//	failure: "not empty"
//
// [Zero values]: https://go.dev/ref/spec#The_zero_value
func Empty(t T, object any, msgAndArgs ...any) bool {
	// Domain: equality
	pass := isEmpty(object)
	if !pass {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		Fail(t, "Should be empty, but was "+truncatingFormat("%v", object), msgAndArgs...)
	}

	return pass
}

// NotEmpty asserts that the specified object is NOT [Empty].
//
// # Usage
//
//	if assert.NotEmpty(t, obj) {
//		assertions.Equal(t, "two", obj[1])
//	}
//
// # Examples
//
//	success: "not empty"
//	failure: ""
func NotEmpty(t T, object any, msgAndArgs ...any) bool {
	// Domain: equality
	pass := !isEmpty(object)
	if !pass {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		Fail(t, fmt.Sprintf("Should NOT be empty, but was %v", object), msgAndArgs...)
	}

	return pass
}

// isNil checks if a specified object is nil or not, without Failing.
func isNil(object any) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	switch value.Kind() {
	case
		reflect.Chan, reflect.Func,
		reflect.Interface, reflect.Map,
		reflect.Ptr, reflect.Slice, reflect.UnsafePointer:

		return value.IsNil()
	default:
		return false
	}
}

// isEmpty gets whether the specified object is considered empty or not.
func isEmpty(object any) bool {
	// get nil case out of the way
	if object == nil {
		return true
	}

	return isEmptyValue(reflect.ValueOf(object))
}

// isEmptyValue gets whether the specified reflect.Value is considered empty or not.
func isEmptyValue(objValue reflect.Value) bool {
	if objValue.IsZero() {
		return true
	}
	// Special cases of non-zero values that we consider empty
	switch objValue.Kind() {
	// collection types are empty when they have no element
	// Note: array types are empty when they match their zero-initialized state.
	case reflect.Chan, reflect.Map, reflect.Slice:
		return objValue.Len() == 0
	// non-nil pointers are empty if the value they point to is empty
	case reflect.Ptr:
		return isEmptyValue(objValue.Elem())
	default:
		return false
	}
}
