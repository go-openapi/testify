// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions/enable/colors"
)

// Equal asserts that two objects are equal.
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
//
// Function equality cannot be determined and will always fail.
//
// # Usage
//
//	assertions.Equal(t, 123, 123)
//
// # Examples
//
//	success: 123, 123
//	failure: 123, 456
func Equal(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if !ObjectsAreEqual(expected, actual) {
		return failWithDiff(t, expected, actual, msgAndArgs...)
	}

	return true
}

// EqualT asserts that two objects of the same comparable type are equal.
//
// Pointer variable equality is determined based on the equality of the memory addresses (unlike [Equal], but like [Same]).
//
// Functions, slices and maps are not comparable. See also [ComparisonOperators].
//
// If you need to compare values of non-comparable types, or compare pointers by the value they point to,
// use [Equal] instead.
//
// # Usage
//
//	assertions.EqualT(t, 123, 123)
//
// # Examples
//
//	success: 123, 123
//	failure: 123, 456
//
// [ComparisonOperators]: https://go.dev/ref/spec#Comparison_operators.
func EqualT[V comparable](t T, expected, actual V, msgAndArgs ...any) bool {
	// Domain: equality
	if expected != actual {
		return failWithDiff(t, expected, actual, msgAndArgs...)
	}

	return true
}

// NotEqual asserts that the specified values are NOT equal.
//
// # Usage
//
//	assertions.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
//
// # Examples
//
//	success: 123, 456
//	failure: 123, 123
func NotEqual(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v != %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if ObjectsAreEqual(expected, actual) {
		return Fail(t, fmt.Sprintf("Should not be: %s\n", truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

// NotEqualT asserts that the specified values of the same comparable type are NOT equal.
//
// See [EqualT].
//
// # Usage
//
//	assertions.NotEqualT(t, obj1, obj2)
//
// # Examples
//
//	success: 123, 456
//	failure: 123, 123
func NotEqualT[V comparable](t T, expected, actual V, msgAndArgs ...any) bool {
	// Domain: equality
	if expected == actual {
		return Fail(t, fmt.Sprintf("Should not be: %s\n", truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

// Same asserts that two pointers reference the same object.
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
//
// # Usage
//
//	assertions.Same(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, staticVarPtr
//	failure: &staticVar, ptr("static string")
func Same(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	same, ok := samePointers(expected, actual)
	if !ok {
		return Fail(t, "Both arguments must be pointers", msgAndArgs...)
	}

	if !same {
		// both are pointers but not the same type & pointing to the same address
		return Fail(t, fmt.Sprintf("Not same: \n"+
			"expected: %[2]s (%[1]T)(%[1]p)\n"+
			"actual  : %[4]s (%[3]T)(%[3]p)", expected, truncatingFormat("%#v", expected), actual, truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

// SameT asserts that two pointers of the same type reference the same object.
//
// # Usage
//
//	assertions.SameT(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, staticVarPtr
//	failure: &staticVar, ptr("static string")
func SameT[P any](t T, expected, actual *P, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if expected != actual {
		return Fail(t, fmt.Sprintf("Not same: \n"+
			"expected: %[2]s (%[1]T)(%[1]p)\n"+
			"actual  : %[4]s (%[3]T)(%[3]p)", expected, truncatingFormat("%#v", expected), actual, truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

// NotSame asserts that two pointers do not reference the same object.
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
//
// # Usage
//
//	assertions.NotSame(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, ptr("static string")
//	failure: &staticVar, staticVarPtr
func NotSame(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	same, ok := samePointers(expected, actual)
	if !ok {
		// fails when the arguments are not pointers
		return !(Fail(t, "Both arguments must be pointers", msgAndArgs...))
	}

	if same {
		return Fail(t, fmt.Sprintf(
			"Expected and actual point to the same object: %p %s",
			expected, truncatingFormat("%#v", expected)), msgAndArgs...)
	}
	return true
}

// NotSameT asserts that two pointers do not reference the same object.
//
// See [SameT]
//
// # Usage
//
//	assertions.NotSameT(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, ptr("static string")
//	failure: &staticVar, staticVarPtr
func NotSameT[P any](t T, expected, actual *P, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if expected == actual {
		return Fail(t, fmt.Sprintf(
			"Expected and actual point to the same object: %p %s",
			expected, truncatingFormat("%#v", expected)), msgAndArgs...)
	}

	return true
}

// EqualValues asserts that two objects are equal or convertible to the larger
// type and equal.
//
// # Usage
//
//	assertions.EqualValues(t, uint32(123), int32(123))
//
// # Examples
//
//	success: uint32(123), int32(123)
//	failure: uint32(123), int32(456)
func EqualValues(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if !ObjectsAreEqualValues(expected, actual) {
		diff := diff(expected, actual)
		expected, actual = formatUnequalValues(expected, actual)
		return Fail(t, fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s%s", expected, actual, diff), msgAndArgs...)
	}

	return true
}

// EqualExportedValues asserts that the types of two objects are equal and their public
// fields are also equal. This is useful for comparing structs that have private fields
// that could potentially differ.
//
// # Usage
//
//	 type S struct {
//		Exported     	int
//		notExported   	int
//	 }
//	assertions.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
//	assertions.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
//
// # Examples
//
//	success: &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}
//	failure:  &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}
func EqualExportedValues(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	aType := reflect.TypeOf(expected)
	bType := reflect.TypeOf(actual)

	if aType != bType {
		return Fail(t, fmt.Sprintf("Types expected to match exactly\n\t%v != %v", aType, bType), msgAndArgs...)
	}

	expected = copyExportedFields(expected)
	actual = copyExportedFields(actual)

	if !ObjectsAreEqualValues(expected, actual) {
		diff := diff(expected, actual)
		expected, actual = formatUnequalValues(expected, actual)
		return Fail(t, fmt.Sprintf("Not equal (comparing only exported fields): \n"+
			"expected: %s\n"+
			"actual  : %s%s", expected, actual, diff), msgAndArgs...)
	}

	return true
}

// Exactly asserts that two objects are equal in value and type.
//
// # Usage
//
//	assertions.Exactly(t, int32(123), int64(123))
//
// # Examples
//
//	success: int32(123), int32(123)
//	failure: int32(123), int64(123)
func Exactly(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	aType := reflect.TypeOf(expected)
	bType := reflect.TypeOf(actual)

	if aType != bType {
		return Fail(t, fmt.Sprintf("Types expected to match exactly\n\t%v != %v", aType, bType), msgAndArgs...)
	}

	return Equal(t, expected, actual, msgAndArgs...)
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

// NotEqualValues asserts that two objects are not equal even when converted to the same type.
//
// # Usage
//
//	assertions.NotEqualValues(t, obj1, obj2)
//
// # Examples
//
//	success: uint32(123), int32(456)
//	failure: uint32(123), int32(123)
func NotEqualValues(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if ObjectsAreEqualValues(expected, actual) {
		return Fail(t, fmt.Sprintf("Should not be: %s\n", truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

func failWithDiff(t T, expected, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	diff := diff(expected, actual)
	expectedStr, actualStr := formatUnequalValues(expected, actual)

	if colors.Enabled() {
		expectedStr = colors.ExpectedColorizer()(expectedStr)
		actualStr = colors.ActualColorizer()(actualStr)
	}

	return Fail(t,
		fmt.Sprintf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s%s",
			expectedStr,
			actualStr, diff),
		msgAndArgs...,
	)
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

// validateEqualArgs checks whether provided arguments can be safely used in the
// Equal/NotEqual functions.
func validateEqualArgs(expected, actual any) error {
	if expected == nil && actual == nil {
		return nil
	}

	if isFunction(expected) || isFunction(actual) {
		return errors.New("cannot take func type as argument")
	}
	return nil
}

// samePointers checks if two arbitrary interface objects are pointers of the same
// type pointing to the same object.
//
// It returns two values: same indicating if they are the same type and point to the same object,
// and ok indicating that both inputs are pointers.
func samePointers(first, second any) (same bool, ok bool) {
	firstPtr, secondPtr := reflect.ValueOf(first), reflect.ValueOf(second)
	if firstPtr.Kind() != reflect.Pointer || secondPtr.Kind() != reflect.Pointer {
		return false, false // not both are pointers
	}

	firstType, secondType := reflect.TypeOf(first), reflect.TypeOf(second)
	if firstType != secondType {
		return false, true // both are pointers, but of different types
	}

	// compare pointer addresses
	return first == second, true
}

// formatUnequalValues takes two values of arbitrary types and returns string
// representations appropriate to be presented to the user.
//
// If the values are not of like type, the returned strings will be prefixed
// with the type name, and the value will be enclosed in parentheses similar
// to a type conversion in the Go grammar.
func formatUnequalValues(expected, actual any) (e string, a string) {
	if reflect.TypeOf(expected) != reflect.TypeOf(actual) {
		return fmt.Sprintf("%T(%s)", expected, truncatingFormat("%#v", expected)),
			fmt.Sprintf("%T(%s)", actual, truncatingFormat("%#v", actual))
	}
	switch expected.(type) {
	case time.Duration, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprint(expected), fmt.Sprint(actual)
	default:
		return truncatingFormat("%#v", expected), truncatingFormat("%#v", actual)
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

// copyExportedFields iterates downward through nested data structures and creates a copy
// that only contains the exported struct fields.
func copyExportedFields(expected any) any {
	if isNil(expected) {
		return expected
	}

	expectedType := reflect.TypeOf(expected)
	expectedKind := expectedType.Kind()
	expectedValue := reflect.ValueOf(expected)

	switch expectedKind {
	case reflect.Struct:
		result := reflect.New(expectedType).Elem()
		for i := range expectedType.NumField() {
			field := expectedType.Field(i)
			isExported := field.IsExported()
			if isExported {
				fieldValue := expectedValue.Field(i)
				if isNil(fieldValue) || isNil(fieldValue.Interface()) {
					continue
				}
				newValue := copyExportedFields(fieldValue.Interface())
				result.Field(i).Set(reflect.ValueOf(newValue))
			}
		}
		return result.Interface()

	case reflect.Pointer:
		result := reflect.New(expectedType.Elem())
		unexportedRemoved := copyExportedFields(expectedValue.Elem().Interface())
		result.Elem().Set(reflect.ValueOf(unexportedRemoved))
		return result.Interface()

	case reflect.Array, reflect.Slice:
		var result reflect.Value
		if expectedKind == reflect.Array {
			result = reflect.New(reflect.ArrayOf(expectedValue.Len(), expectedType.Elem())).Elem()
		} else {
			result = reflect.MakeSlice(expectedType, expectedValue.Len(), expectedValue.Len())
		}
		for i := range expectedValue.Len() {
			index := expectedValue.Index(i)
			if isNil(index) {
				continue
			}
			unexportedRemoved := copyExportedFields(index.Interface())
			result.Index(i).Set(reflect.ValueOf(unexportedRemoved))
		}
		return result.Interface()

	case reflect.Map:
		result := reflect.MakeMap(expectedType)
		for _, k := range expectedValue.MapKeys() {
			index := expectedValue.MapIndex(k)
			unexportedRemoved := copyExportedFields(index.Interface())
			result.SetMapIndex(k, reflect.ValueOf(unexportedRemoved))
		}
		return result.Interface()

	default:
		return expected
	}
}

func isFunction(arg any) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}
