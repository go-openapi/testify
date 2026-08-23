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

	// A signed and an unsigned integer cannot be compared through a conversion: a negative
	// value wraps to a huge positive one and matches it. Ordering the conversion from the
	// smaller type to the larger one does not help either, since int8(-1) wraps to
	// uint64(math.MaxUint64) exactly as int64(-1) does.
	if isMixedSignIntegerPair(expectedType, actualType) {
		return equalMixedSignIntegers(expectedValue, actualValue)
	}

	expectedConverted := expectedValue.Convert(actualType)
	if !expectedConverted.CanInterface() {
		// Unreachable with current Go reflection: values from reflect.ValueOf()
		// are always interfaceable, and Convert() preserves that property.
		// CanInterface() is only false for values obtained via unexported struct fields.
		panic("reflect: converted value is not interfaceable")
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
			// Unreachable with current Go reflection: all numeric kinds (Int through Complex128)
			// are convertible to all other numeric kinds.
			panic("reflect: numeric value is not convertible to numeric type")
		}

		actualConverted := actualValue.Convert(expectedType)
		if !actualConverted.CanInterface() {
			// Unreachable with current Go reflection: values from reflect.ValueOf()
			// are always interfaceable, and Convert() preserves that property.
			// CanInterface() is only false for values obtained via unexported struct fields.
			panic("reflect: converted value is not interfaceable")
		}

		return actualConverted.Interface() == expected
	}

	return expectedConverted.Interface() == actual
}

// isSignedInteger reports whether the type is one of int, int8, int16, int32, int64.
func isSignedInteger(t reflect.Type) bool {
	return t.Kind() >= reflect.Int && t.Kind() <= reflect.Int64
}

// isUnsignedInteger reports whether the type is one of uint, uint8, uint16, uint32, uint64,
// uintptr.
func isUnsignedInteger(t reflect.Type) bool {
	return t.Kind() >= reflect.Uint && t.Kind() <= reflect.Uintptr
}

// isMixedSignIntegerPair reports whether one type is a signed integer and the other an
// unsigned one, in either order.
func isMixedSignIntegerPair(a, b reflect.Type) bool {
	return (isSignedInteger(a) && isUnsignedInteger(b)) || (isUnsignedInteger(a) && isSignedInteger(b))
}

// equalMixedSignIntegers compares a signed and an unsigned integer by value rather than by
// conversion. A negative signed value equals no unsigned value at all; otherwise both fit in
// a uint64 and compare directly.
func equalMixedSignIntegers(a, b reflect.Value) bool {
	signed, unsigned := a, b
	if isSignedInteger(b.Type()) {
		signed, unsigned = b, a
	}

	value := signed.Int()
	if value < 0 {
		return false
	}

	return uint64(value) == unsigned.Uint()
}

// isNumericType returns true if the type is one of:
// int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64,
// float32, float64, complex64, complex128.
func isNumericType(t reflect.Type) bool {
	return t.Kind() >= reflect.Int && t.Kind() <= reflect.Complex128
}
