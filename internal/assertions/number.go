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
// # Usage
//
// assertions.InDelta(t, math.Pi, 22/7.0, 0.01)
//
// # Examples
//
//	success: 1.0, 1.01, 0.02
//	failure: 1.0, 1.1, 0.05
func InDelta(t T, expected, actual any, delta float64, msgAndArgs ...any) bool {
	// Domain: number
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
		return Fail(t, "Expected must not be NaN", msgAndArgs...) // Proposal for enhancement: wrong message (this is accepted above)
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

// InDeltaT asserts that the two numerals of the same type numerical type are within delta of each other.
//
// [InDeltaT] accepts any go numeric type, including integer types.
//
// The main difference with [InDelta] is that the delta is expressed with the same type as the values, not necessarily a float64.
//
// Delta must be greater than or equal to zero.
//
// # Behavior with IEEE floating point arithmetics
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
//
// # Usage
//
// assertions.InDeltaT(t, math.Pi, 22/7.0, 0.01)
//
// # Examples
//
//	success: 1.0, 1.01, 0.02
//	failure: 1.0, 1.1, 0.05
func InDeltaT[Number Measurable](t T, expected, actual, delta Number, msgAndArgs ...any) bool {
	// Domain: number
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if delta < 0 {
		return Fail(t, "Delta must not be negative", msgAndArgs...) // TODO: add it to the original version
	}

	// IEEE float edge cases: NaN, +Inf/-Inf
	if isNaN(delta) || isInf(delta, 0) {
		return Fail(t, "Delta must not be NaN or Inf", msgAndArgs...) // TODO: add it to the original version
	}

	expectedInf := isInf(expected, 0)
	actualInf := isInf(actual, 0)
	if expectedInf {
		// expected -Inf/+Inf
		if !actualInf {
			return Fail(t, "Expected an Inf value", msgAndArgs...)
		}

		if isInf(expected, 1) && !isInf(actual, 1) {
			return Fail(t, "Expected a +Inf value but got -Inf", msgAndArgs...)
		}

		if isInf(expected, -1) && !isInf(actual, -1) {
			return Fail(t, "Expected a -Inf value but got +Inf", msgAndArgs...)
		}

		// Both are Inf and match - success
		return true
	}

	if actualInf {
		return Fail(t, "Actual is Inf", msgAndArgs...)
	}

	expectedNaN := isNaN(expected)
	actualNaN := isNaN(actual)

	if expectedNaN && actualNaN {
		// expected NaN
		return true
	}

	if expectedNaN {
		return Fail(t, "Expected a NaN value but actual is finite", msgAndArgs...)
	}

	if actualNaN {
		return Fail(t, fmt.Sprintf("Expected %v with delta %v, but was NaN", expected, delta), msgAndArgs...)
	}

	var (
		dt     Number
		failed bool
	)

	// check is a little slower than straight delta, but it can handle unsigned numbers without errors
	if expected > actual {
		dt = expected - actual
	} else {
		dt = actual - expected
	}
	failed = dt > delta
	if failed {
		return Fail(t, fmt.Sprintf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, dt), msgAndArgs...)
	}

	return true
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon.
//
// # Usage
//
//	assertions.InEpsilon(t, 100.0, 101.0, 0.02)
//
// # Examples
//
//	success: 100.0, 101.0, 0.02
//	failure: 100.0, 110.0, 0.05
func InEpsilon(t T, expected, actual any, epsilon float64, msgAndArgs ...any) bool {
	// Domain: number
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

// InEpsilonT asserts that expected and actual have a relative error less than epsilon.
//
// When expected is zero, epsilon is interpreted as an absolute error threshold,
// since relative error is mathematically undefined for zero values.
//
// Formula:
//   - If expected == 0: fail if |actual - expected| > epsilon
//   - If expected != 0: fail if |actual - expected| > epsilon * |expected|
//
// This allows InEpsilonT to work naturally across the full numeric range including zero.
//
// # Usage
//
//	assertions.InEpsilon(t, 100.0, 101.0, 0.02)
//
// # Examples
//
//	success: 100.0, 101.0, 0.02
//	failure: 100.0, 110.0, 0.05
func InEpsilonT[Number Measurable](t T, expected, actual Number, epsilon float64, msgAndArgs ...any) bool {
	// Domain: number
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if epsilon < 0 {
		return Fail(t, "Epsilon must not be negative", msgAndArgs...)
	}

	// IEEE float edge cases: NaN, +Inf/-Inf
	if isNaN(epsilon) || isInf(epsilon, 0) {
		return Fail(t, "Epsilon must not be NaN or Inf", msgAndArgs...)
	}

	expectedInf := isInf(expected, 0)
	actualInf := isInf(actual, 0)
	if expectedInf {
		// expected -Inf/+Inf
		if !actualInf {
			return Fail(t, "Expected an Inf value", msgAndArgs...)
		}

		if isInf(expected, 1) && !isInf(actual, 1) {
			return Fail(t, "Expected a +Inf value but got -Inf", msgAndArgs...)
		}

		if isInf(expected, -1) && !isInf(actual, -1) {
			return Fail(t, "Expected a -Inf value but got +Inf", msgAndArgs...)
		}

		// Both are Inf and match - success
		return true
	}

	if actualInf {
		return Fail(t, "Actual is Inf", msgAndArgs...)
	}

	expectedNaN := isNaN(expected)
	actualNaN := isNaN(actual)

	if expectedNaN && actualNaN {
		// expected NaN
		return true
	}

	if expectedNaN {
		return Fail(t, "Expected a NaN value but actual is finite", msgAndArgs...)
	}

	if actualNaN {
		return Fail(t, fmt.Sprintf("Expected %v with epsilon %v, but was NaN", expected, epsilon), msgAndArgs...)
	}

	af := float64(expected)
	bf := float64(actual)

	delta := math.Abs(af - bf)
	if delta == 0 {
		return true
	}
	if af == 0 {
		if delta > epsilon {
			return Fail(t, fmt.Sprintf(
				"Expected value is zero, using absolute error comparison.\n"+
					"Absolute difference is too high: %#v (expected)\n"+
					"        < %#v (actual)", epsilon, delta), msgAndArgs...)
		}
		return true
	}

	if delta > epsilon*math.Abs(af) {
		return Fail(t, fmt.Sprintf("Relative error is too high: %#v (expected)\n"+
			"        < %#v (actual)", epsilon, delta/math.Abs(af)), msgAndArgs...)
	}

	return true
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
//
// # Usage
//
//	assertions.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
//
// # Examples
//
//	success: []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02
//	failure: []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05
func InDeltaSlice(t T, expected, actual any, delta float64, msgAndArgs ...any) bool {
	// Domain: number
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
// # Usage
//
//	assertions.InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
//
// # Examples
//
//	success: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02
//	failure: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05
func InDeltaMapValues(t T, expected, actual any, delta float64, msgAndArgs ...any) bool {
	// Domain: number
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

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
//
// # Usage
//
//	assertions.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
//
// # Examples
//
//	success: []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02
//	failure: []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05
func InEpsilonSlice(t T, expected, actual any, epsilon float64, msgAndArgs ...any) bool {
	// Domain: number
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

func isNaN[Number Measurable](f Number) bool {
	return f != f //nolint:gocritic // yes this weird property is held by NaN only
}

func isInf[Number Measurable](f Number, sign int) bool {
	v := any(f)

	switch ff := v.(type) {
	case float32:
		return isInf32(ff, sign)
	case float64:
		return math.IsInf(ff, sign)
	default:
		return false
	}
}

func isInf32(f float32, sign int) bool {
	return (sign >= 0 && f > math.MaxFloat32) || (sign <= 0 && f < -math.MaxFloat32)
}
