// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"reflect"
)

// ObjectsAreEqual determines if two objects are considered equal.
//
// This function does no assertion of any kind.
func ObjectsAreEqual(expected, actual any) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}

	if exp == nil || act == nil {
		return exp == nil && act == nil
	}

	return bytes.Equal(exp, act)
}

// ObjectsAreEqualValues gets whether two objects are equal, or if their
// values are equal.
func ObjectsAreEqualValues(expected, actual any) bool {
	if ObjectsAreEqual(expected, actual) {
		return true
	}

	expectedValue := reflect.ValueOf(expected)
	actualValue := reflect.ValueOf(actual)
	if !expectedValue.IsValid() || !actualValue.IsValid() {
		return false
	}

	expectedType := expectedValue.Type()
	actualType := actualValue.Type()
	if !expectedType.ConvertibleTo(actualType) {
		return false
	}

	// Attempt conversion of expected to actual type.
	// This handles more cases than just the ConvertibleTo check above.
	if !expectedValue.CanConvert(actualType) {
		// Types are not convertible, so they cannot be equal
		// This prevents panics when calling [reflect.Value.Convert]
		return false
	}

	expectedConverted := expectedValue.Convert(actualType)
	if !expectedConverted.CanInterface() {
		// Cannot interface after conversion, so cannot be equal.
		// This prevents panics when calling [reflect.Value.Interface].
		return false
	}

	if !isNumericType(expectedType) || !isNumericType(actualType) {
		// Attempt comparison after type conversion.
		return reflect.DeepEqual(
			expectedConverted.Interface(), actual,
		)
	}

	// If BOTH values are numeric, there are chances of false positives due
	// to overflow or underflow. So, we need to make sure to always convert
	// the smaller type to a larger type before comparing.
	if expectedType.Size() >= actualType.Size() {
		if !actualValue.CanConvert(expectedType) {
			// Cannot convert actual to the expected type, so cannot be equal.
			// This is a hypothetical case to prevent panics when calling [reflect.Value.Convert].
			return false
		}

		actualConverted := actualValue.Convert(expectedType)
		if !actualConverted.CanInterface() {
			// Cannot interface after conversion, so cannot be equal.
			// This is a hypothetical case to prevent panics when calling [reflect.Value.Convert].
			return false
		}

		return actualConverted.Interface() == expected
	}

	return expectedConverted.Interface() == actual
}

// isNumericType returns true if the type is one of:
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
// float32, float64, complex64, complex128.
func isNumericType(t reflect.Type) bool {
	return t.Kind() >= reflect.Int && t.Kind() <= reflect.Complex128
}
