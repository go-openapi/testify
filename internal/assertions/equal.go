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
// Function equality cannot be determined and will always fail.
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

// EqualValues asserts that two objects are equal or convertible to the larger
// type and equal.
//
// Function equality cannot be determined and will always fail.
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

	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
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

// NotEqualValues asserts that two objects are not equal even when converted to the same type.
//
// Function equality cannot be determined and will always fail.
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

	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
	}

	if ObjectsAreEqualValues(expected, actual) {
		return Fail(t, fmt.Sprintf("Should not be: %s\n", truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

// EqualExportedValues asserts that the types of two objects are equal and their public
// fields are also equal.
//
// This is useful for comparing structs that have private fields that could potentially differ.
//
// Function equality cannot be determined and will always fail.
//
// # Usage
//
//	type S struct {
//		Exported     	int
//		notExported   	int
//	}
//	assertions.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
//	assertions.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
//
// # Examples
//
//	success: &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}
//	failure: &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}
func EqualExportedValues(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if err := validateEqualArgs(expected, actual); err != nil {
		return Fail(t, fmt.Sprintf("Invalid operation: %#v == %#v (%s)",
			expected, actual, err), msgAndArgs...)
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

// validateEqualArgs checks whether provided arguments can be safely used in the
// Equal/NotEqual/EqualValues/NotEqualValues functions.
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
			if !index.CanInterface() {
				// this should not be possible with current reflect, since values are retrieved from an array or slice, not a struct
				panic(fmt.Errorf("internal error: can't resolve Interface() for value %v", index))
			}
			unexportedRemoved := copyExportedFields(index.Interface())
			result.Index(i).Set(reflect.ValueOf(unexportedRemoved))
		}
		return result.Interface()

	case reflect.Map:
		result := reflect.MakeMap(expectedType)
		for _, k := range expectedValue.MapKeys() {
			index := expectedValue.MapIndex(k)
			if !index.CanInterface() {
				// this should not be possible with current reflect, since values are retrieved from a map, not a struct
				panic(fmt.Errorf("internal error: can't resolve Interface() for value %v", index))
			}
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
