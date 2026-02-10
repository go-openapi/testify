// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"math"
	"reflect"
	"time"
)

// InDelta asserts that the two numerals are within delta of each other.
//
// Delta must be greater than or equal to zero.
//
// Expected and actual values should convert to float64.
// To compare large integers that can't be represented accurately as float64 (e.g. uint64),
// prefer [InDeltaT] to preserve the original type.
//
// # Behavior with IEEE floating point arithmetic
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
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

	msg, skip, ok := checkDeltaEdgeCases(af, bf, delta)
	if !ok {
		return Fail(t, msg, msgAndArgs...)
	}
	if skip {
		return true
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
// # Behavior with IEEE floating point arithmetic
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

	msg, skip, ok := checkDeltaEdgeCases(expected, actual, delta)
	if !ok {
		return Fail(t, msg, msgAndArgs...)
	}
	if skip {
		return true
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
// # Behavior with IEEE floating point arithmetic
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
//
// Edge case: for very large integers that do not convert accurately to a float64 (e.g. uint64), prefer [InDeltaT].
//
// Formula:
//   - If expected == 0: fail if |actual - expected| > epsilon
//   - If expected != 0: fail if |actual - expected| > epsilon * |expected|
//
// This allows [InEpsilonT] to work naturally across the full numeric range including zero.
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
	af, aok := toFloat(expected)
	bf, bok := toFloat(actual)
	if !aok || !bok {
		return Fail(t, "Parameters must be numerical", msgAndArgs...)
	}

	msg, skip, ok := checkDeltaEdgeCases(af, bf, epsilon)
	if !ok {
		return Fail(t, msg, msgAndArgs...)
	}
	if skip {
		return true
	}

	msg, ok = compareRelativeError(af, bf, epsilon)
	if !ok {
		return Fail(t, msg, msgAndArgs...)
	}

	return true
}

// InEpsilonT asserts that expected and actual have a relative error less than epsilon.
//
// When expected is zero, epsilon is interpreted as an absolute error threshold,
// since relative error is mathematically undefined for zero values.
//
// Unlike [InDeltaT], which preserves the original type, [InEpsilonT] converts the expected and actual
// numbers to float64, since the relative error doesn't make sense as an integer.
//
// # Behavior with IEEE floating point arithmetic
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
//
// Edge case: for very large integers that do not convert accurately to a float64 (e.g. uint64), prefer [InDeltaT].
//
// Formula:
//   - If expected == 0: fail if |actual - expected| > epsilon
//   - If expected != 0: fail if |actual - expected| > epsilon * |expected|
//
// This allows [InEpsilonT] to work naturally across the full numeric range including zero.
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

	af := float64(expected)
	bf := float64(actual)
	msg, skip, ok := checkDeltaEdgeCases(af, bf, epsilon)
	if !ok {
		return Fail(t, msg, msgAndArgs...)
	}
	if skip {
		return true
	}

	msg, ok = compareRelativeError(af, bf, epsilon)
	if !ok {
		return Fail(t, msg, msgAndArgs...)
	}

	return true
}

// InDeltaSlice is the same as [InDelta], except it compares two slices.
//
// See [InDelta].
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

// InDeltaMapValues is the same as [InDelta], but it compares all values between two maps. Both maps must have exactly the same keys.
//
// See [InDelta].
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

		// from [reflect.MapIndex] contract, ev is always a valid [reflect.Value] here
		// because we know that the key has been found.
		// On the other hand, av may not be there.
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

// InEpsilonSlice is the same as [InEpsilon], except it compares each value from two slices.
//
// See [InEpsilon].
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

func checkDeltaEdgeCases[Number Measurable](expected, actual, delta Number) (msg string, skip bool, ok bool) {
	if delta < 0 {
		return "Delta must not be negative", true, false
	}

	// IEEE float edge cases: NaN, +Inf/-Inf
	if isNaN(delta) || isInf(delta, 0) {
		return "Delta must not be NaN or Inf", true, false
	}

	expectedInf := isInf(expected, 0)
	actualInf := isInf(actual, 0)
	if expectedInf {
		// expected -Inf/+Inf
		if !actualInf {
			return "Expected an Inf value", true, false
		}

		if isInf(expected, 1) && !isInf(actual, 1) {
			return "Expected a +Inf value but got -Inf", true, false
		}

		if isInf(expected, -1) && !isInf(actual, -1) {
			return "Expected a -Inf value but got +Inf", true, false
		}

		// Both are Inf and match - success
		return "", true, true
	}

	if actualInf {
		return "Actual is Inf", true, false
	}

	expectedNaN := isNaN(expected)
	actualNaN := isNaN(actual)

	if expectedNaN && actualNaN {
		// expected NaN
		return "", true, true
	}

	if expectedNaN {
		return "Expected a NaN value but actual is finite", true, false
	}

	if actualNaN {
		return fmt.Sprintf("Expected %v with delta %v, but was NaN", expected, delta), true, false
	}

	return "", false, true
}

func compareRelativeError(expected, actual, epsilon float64) (msg string, ok bool) {
	delta := math.Abs(expected - actual)
	if delta == 0 {
		return "", true
	}

	if expected == 0 {
		if delta > epsilon {
			return fmt.Sprintf(
				"Expected value is zero, using absolute error comparison.\n"+
					"Absolute difference is too high: %#v (expected)\n"+
					"        < %#v (actual)", epsilon, delta), false
		}

		return "", true
	}

	if delta > epsilon*math.Abs(expected) {
		return fmt.Sprintf("Relative error is too high: %#v (expected)\n"+
			"        < %#v (actual)", epsilon, delta/math.Abs(expected)), false
	}

	return "", true
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
		// try reflect conversion
		val := reflect.ValueOf(xn)
		typ := reflect.TypeFor[float64]()
		if val.IsValid() && val.CanConvert(typ) {
			rxf := val.Convert(typ)
			xf = rxf.Float()
			break
		}
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
