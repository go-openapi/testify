// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"reflect"
)

// Implements asserts that an object is implemented by the specified interface.
//
// # Usage
//
//	assertions.Implements(t, (*MyInterface)(nil), new(MyObject))
//
// # Examples
//
//	success: ptr(dummyInterface), new(testing.T)
//	failure: (*error)(nil), new(testing.T)
func Implements(t T, interfaceObject any, object any, msgAndArgs ...any) bool {
	// Domain: type
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
// # Usage
//
//	assertions.NotImplements(t, (*MyInterface)(nil), new(MyObject))
//
// # Examples
//
//	success: (*error)(nil), new(testing.T)
//	failure: ptr(dummyInterface), new(testing.T)
func NotImplements(t T, interfaceObject any, object any, msgAndArgs ...any) bool {
	// Domain: type
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

// IsType asserts that the specified objects are of the same type.
//
// # Usage
//
//	assertions.IsType(t, &MyStruct{}, &MyStruct{})
//
// # Examples
//
//	success: 123, 456
//	failure: int32(123), int64(456)
func IsType(t T, expectedType, object any, msgAndArgs ...any) bool {
	// Domain: type
	if isType(expectedType, object) {
		return true
	}
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return Fail(t, fmt.Sprintf("Object expected to be of type %T, but was %T", expectedType, object), msgAndArgs...)
}

// IsOfTypeT asserts that an object is of a given type.
//
// # Usage
//
//	assertions.IsOfTypeT[MyType](t,myVar)
//
// # Examples
//
//	success: myType(123.123)
//	failure: 123.123
func IsOfTypeT[EType any](t T, object any, msgAndArgs ...any) bool {
	// Domain: type
	if h, ok := t.(H); ok {
		h.Helper()
	}

	_, ok := object.(EType)
	if ok {
		return true
	}

	return Fail(t, fmt.Sprintf("Object expected to be of type %v, but was %T", reflect.TypeFor[EType](), object), msgAndArgs...)
}

// IsNotType asserts that the specified objects are not of the same type.
//
// # Usage
//
//	assertions.IsNotType(t, &NotMyStruct{}, &MyStruct{})
//
// # Examples
//
//	success: int32(123), int64(456)
//	failure: 123, 456
func IsNotType(t T, theType, object any, msgAndArgs ...any) bool {
	// Domain: type
	if !isType(theType, object) {
		return true
	}
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return Fail(t, fmt.Sprintf("Object type expected to be different than %T", theType), msgAndArgs...)
}

// IsNotOfTypeT asserts that an object is not of a given type.
//
// # Usage
//
//	assertions.IsOfType[MyType](t,myVar)
//
// # Examples
//
//	success: 123.123
//	failure: myType(123.123)
func IsNotOfTypeT[EType any](t T, object any, msgAndArgs ...any) bool {
	// Domain: type
	if h, ok := t.(H); ok {
		h.Helper()
	}

	_, ok := object.(EType)
	if !ok {
		return true
	}

	return Fail(t, fmt.Sprintf("Object type expected to be different than %T", reflect.TypeFor[EType]()), msgAndArgs...)
}

// Zero asserts that i is the zero value for its type.
//
// # Usage
//
//	assertions.Zero(t, obj)
//
// # Examples
//
//	success: 0
//	failure: 1
func Zero(t T, i any, msgAndArgs ...any) bool {
	// Domain: type
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
// # Usage
//
//	assertions.NotZero(t, obj)
//
// # Examples
//
//	success: 1
//	failure: 0
func NotZero(t T, i any, msgAndArgs ...any) bool {
	// Domain: type
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if i == nil || reflect.DeepEqual(i, reflect.Zero(reflect.TypeOf(i)).Interface()) {
		return Fail(t, fmt.Sprintf("Should not be zero, but was %v", i), msgAndArgs...)
	}
	return true
}

// Kind asserts that the [reflect.Kind] of a given object matches the expected [reflect.Kind].
//
// Kind reflects the concrete value stored in the object. The nil value (or interface with nil value)
// are comparable to [reflect.Invalid]. See also [reflect.Value.Kind].
//
// # Usage
//
//	assertions.Kind(t, reflect.String, "Hello World")
//
// # Examples
//
//	success: reflect.String, "hello"
//	failure: reflect.String, 0
func Kind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool {
	// Domain: type
	if h, ok := t.(H); ok {
		h.Helper()
	}

	val := reflect.ValueOf(object)
	kind := val.Kind()
	if kind != expectedKind {
		if kind == reflect.Invalid {
			// add some explanation when reflect.Invalid does not match the expectation (common gotcha with reflect)
			return Fail(t, "object has reflect.Invalid kind: this is nil or an interface with nil value", msgAndArgs...)
		}

		return Fail(t, fmt.Sprintf("object expected to be of kind %v, but was %v", expectedKind, kind), msgAndArgs...)
	}

	return true
}

// NotKind asserts that the [reflect.Kind] of a given object does not match the expected [reflect.Kind].
//
// Kind reflects the concrete value stored in the object. The nil value (or interface with nil value)
// are comparable to [reflect.Invalid]. See also [reflect.Value.Kind].
//
// # Usage
//
//	assertions.NotKind(t, reflect.Int, "Hello World")
//
// # Examples
//
//	success: reflect.String, 0
//	failure: reflect.String, "hello"
func NotKind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) bool {
	// Domain: type
	if h, ok := t.(H); ok {
		h.Helper()
	}

	val := reflect.ValueOf(object)
	kind := val.Kind()
	if kind != expectedKind {
		return true
	}

	return Fail(t, fmt.Sprintf("object expected not to be of kind %v, but was %v", expectedKind, kind), msgAndArgs...)
}

func isType(expectedType, object any) bool {
	return ObjectsAreEqual(reflect.TypeOf(object), reflect.TypeOf(expectedType))
}
