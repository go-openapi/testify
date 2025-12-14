// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"reflect"
)

// Implements asserts that an object is implemented by the specified interface.
//
//	assert.Implements(t, (*MyInterface)(nil), new(MyObject))
//
// Examples:
//
//	success: ptr(dummyInterface), new(testing.T)
//	failure: (*error)(nil), new(testing.T)
func Implements(t T, interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if object == nil {
		return Fail(t, fmt.Sprintf("Cannot check if nil implements %v", interfaceType), msgAndArgs...)
	}
	if !reflect.TypeOf(object).Implements(interfaceType) {
		return Fail(t, fmt.Sprintf("%T must implement %v", object, interfaceType), msgAndArgs...)
	}

	return true
}

// NotImplements asserts that an object does not implement the specified interface.
//
//	assert.NotImplements(t, (*MyInterface)(nil), new(MyObject))
//
// Examples:
//
//	success: (*error)(nil), new(testing.T)
//	failure: ptr(dummyInterface), new(testing.T)
func NotImplements(t T, interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	interfaceType := reflect.TypeOf(interfaceObject).Elem()

	if object == nil {
		return Fail(t, fmt.Sprintf("Cannot check if nil does not implement %v", interfaceType), msgAndArgs...)
	}
	if reflect.TypeOf(object).Implements(interfaceType) {
		return Fail(t, fmt.Sprintf("%T implements %v", object, interfaceType), msgAndArgs...)
	}

	return true
}

func isType(expectedType, object any) bool {
	return ObjectsAreEqual(reflect.TypeOf(object), reflect.TypeOf(expectedType))
}

// IsType asserts that the specified objects are of the same type.
//
//	assert.IsType(t, &MyStruct{}, &MyStruct{})
//
// Examples:
//
//	success: 123, 456
//	failure: int32(123), int64(456)
func IsType(t T, expectedType, object any, msgAndArgs ...any) bool {
	if isType(expectedType, object) {
		return true
	}
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return Fail(t, fmt.Sprintf("Object expected to be of type %T, but was %T", expectedType, object), msgAndArgs...)
}

// IsNotType asserts that the specified objects are not of the same type.
//
//	assert.IsNotType(t, &NotMyStruct{}, &MyStruct{})
//
// Examples:
//
//	success: int32(123), int64(456)
//	failure: 123, 456
func IsNotType(t T, theType, object any, msgAndArgs ...any) bool {
	if !isType(theType, object) {
		return true
	}
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return Fail(t, fmt.Sprintf("Object type expected to be different than %T", theType), msgAndArgs...)
}

// Zero asserts that i is the zero value for its type.
//
// Examples:
//
//	success: 0
//	failure: 1
func Zero(t T, i any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if i != nil && !reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface()) {
		return Fail(t, "Should be zero, but was "+truncatingFormat("%v", i), msgAndArgs...)
	}
	return true
}

// NotZero asserts that i is not the zero value for its type.
//
// Examples:
//
//	success: 1
//	failure: 0
func NotZero(t T, i any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if i == nil || reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface()) {
		return Fail(t, fmt.Sprintf("Should not be zero, but was %v", i), msgAndArgs...)
	}
	return true
}
