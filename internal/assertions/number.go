// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"time"
)

// InDelta asserts that the two numerals are within delta of each other.
//
//	assert.InDelta(t, math.Pi, 22/7.0, 0.01)
//
// Examples:
//
//	success: 1.0, 1.01, 0.02
//	failure: 1.0, 1.1, 0.05
func InDelta(t T, expected, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	af, aok := toFloat(expected)
	bf, bok := toFloat(actual)

	if !aok || !bok {
		return Fail(t, "Parameters must be numerical", msgAndArgs...)
	}

	if math.IsNaN(af) && math.IsNaN(bf) {
		return true
	}

	if math.IsNaN(af) {
		return Fail(t, "Expected must not be NaN", msgAndArgs...)
	}

	if math.IsNaN(bf) {
		return Fail(t, fmt.Sprintf("Expected %v with delta %v, but was NaN", expected, delta), msgAndArgs...)
	}

	dt := af - bf
	if dt < -delta || dt > delta {
		return Fail(t, fmt.Sprintf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, dt), msgAndArgs...)
	}

	return true
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
//
// Examples:
//
//	success: []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02
//	failure: []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05
func InDeltaSlice(t T, expected, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if expected == nil || actual == nil ||
		reflect.TypeOf(actual).Kind() != reflect.Slice ||
		reflect.TypeOf(expected).Kind() != reflect.Slice {
		return Fail(t, "Parameters must be slice", msgAndArgs...)
	}

	actualSlice := reflect.ValueOf(actual)
	expectedSlice := reflect.ValueOf(expected)

	for i := range actualSlice.Len() {
		result := InDelta(t, actualSlice.Index(i).Interface(), expectedSlice.Index(i).Interface(), delta, msgAndArgs...)
		if !result {
			return result
		}
	}

	return true
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
//
// Examples:
//
//	success: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02
//	failure: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05
func InDeltaMapValues(t T, expected, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if expected == nil || actual == nil ||
		reflect.TypeOf(actual).Kind() != reflect.Map ||
		reflect.TypeOf(expected).Kind() != reflect.Map {
		return Fail(t, "Arguments must be maps", msgAndArgs...)
	}

	expectedMap := reflect.ValueOf(expected)
	actualMap := reflect.ValueOf(actual)

	if expectedMap.Len() != actualMap.Len() {
		return Fail(t, "Arguments must have the same number of keys", msgAndArgs...)
	}

	for _, k := range expectedMap.MapKeys() {
		ev := expectedMap.MapIndex(k)
		av := actualMap.MapIndex(k)

		if !ev.IsValid() {
			return Fail(t, fmt.Sprintf("missing key %q in expected map", k), msgAndArgs...)
		}

		if !av.IsValid() {
			return Fail(t, fmt.Sprintf("missing key %q in actual map", k), msgAndArgs...)
		}

		if !InDelta(
			t,
			ev.Interface(),
			av.Interface(),
			delta,
			msgAndArgs...,
		) {
			return false
		}
	}

	return true
}

func calcRelativeError(expected, actual any) (float64, error) {
	af, aok := toFloat(expected)
	bf, bok := toFloat(actual)
	if !aok || !bok {
		return 0, errors.New("parameters must be numerical")
	}
	if math.IsNaN(af) && math.IsNaN(bf) {
		return 0, nil
	}
	if math.IsNaN(af) {
		return 0, errors.New("expected value must not be NaN")
	}
	if af == 0 {
		return 0, errors.New("expected value must have a value other than zero to calculate the relative error")
	}
	if math.IsNaN(bf) {
		return 0, errors.New("actual value must not be NaN")
	}

	return math.Abs(af-bf) / math.Abs(af), nil
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon.
//
// Examples:
//
//	success: 100.0, 101.0, 0.02
//	failure: 100.0, 110.0, 0.05
func InEpsilon(t T, expected, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if math.IsNaN(epsilon) {
		return Fail(t, "epsilon must not be NaN", msgAndArgs...)
	}
	actualEpsilon, err := calcRelativeError(expected, actual)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}
	if math.IsNaN(actualEpsilon) {
		return Fail(t, "relative error is NaN", msgAndArgs...)
	}
	if actualEpsilon > epsilon {
		return Fail(t, fmt.Sprintf("Relative error is too high: %#v (expected)\n"+
			"        < %#v (actual)", epsilon, actualEpsilon), msgAndArgs...)
	}

	return true
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
//
// Examples:
//
//	success: []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02
//	failure: []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05
func InEpsilonSlice(t T, expected, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if expected == nil || actual == nil {
		return Fail(t, "Parameters must be slice", msgAndArgs...)
	}

	expectedSlice := reflect.ValueOf(expected)
	actualSlice := reflect.ValueOf(actual)

	if expectedSlice.Type().Kind() != reflect.Slice {
		return Fail(t, "Expected value must be slice", msgAndArgs...)
	}

	expectedLen := expectedSlice.Len()
	if !IsType(t, expected, actual) || !Len(t, actual, expectedLen) {
		return false
	}

	for i := range expectedLen {
		if !InEpsilon(t, expectedSlice.Index(i).Interface(), actualSlice.Index(i).Interface(), epsilon, "at index %d", i) {
			return false
		}
	}

	return true
}

func toFloat(x any) (float64, bool) {
	var xf float64
	xok := true

	switch xn := x.(type) {
	case uint:
		xf = float64(xn)
	case uint8:
		xf = float64(xn)
	case uint16:
		xf = float64(xn)
	case uint32:
		xf = float64(xn)
	case uint64:
		xf = float64(xn)
	case int:
		xf = float64(xn)
	case int8:
		xf = float64(xn)
	case int16:
		xf = float64(xn)
	case int32:
		xf = float64(xn)
	case int64:
		xf = float64(xn)
	case float32:
		xf = float64(xn)
	case float64:
		xf = xn
	case time.Duration:
		xf = float64(xn)
	default:
		xok = false
	}

	return xf, xok
}
