// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"time"
)

// Greater asserts that the first element is strictly greater than the second.
//
// Both elements must be of the same type in the [reflect.Kind] sense.
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.Greater(t, 2, 1)
//	assertions.Greater(t, float64(2), float64(1))
//	assertions.Greater(t, "b", "a")
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
func Greater(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not greater than \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareGreater}, failMessage, msgAndArgs...)
}

// GreaterT asserts that for two elements of the same type,
// the first element is strictly greater than the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [GreaterT] with [*time.Time].
//
// [GreaterT] ensures type safety at build time. If you need to compare values with a dynamically assigned type, use [Greater] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.GreaterT(t, 2, 1)
//	assertions.GreaterT(t, float64(2), float64(1))
//	assertions.GreaterT(t, "b", "a")
//	assertions.GreaterT(t, time.Date(2026,1,1,0,0,0,0,nil), time.Now())
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
func GreaterT[Orderable Ordered](t T, e1, e2 Orderable, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}

	result := compareOrderedWithAny(e1, e2)
	if result > 0 {
		return true
	}

	return Fail(t, fmt.Sprintf("\"%v\" is not greater than \"%v\"", e1, e2), msgAndArgs...)
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second.
//
// See also [Greater].
//
// # Usage
//
//	assertions.GreaterOrEqual(t, 2, 1)
//	assertions.GreaterOrEqual(t, 2, 2)
//	assertions.GreaterOrEqual(t, "b", "a")
//	assertions.GreaterOrEqual(t, "b", "b")
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
func GreaterOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not greater than or equal to \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareGreater, compareEqual}, failMessage, msgAndArgs...)
}

// GreaterOrEqualT asserts that for two elements of the same type,
// the first element is greater than or equal to the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [GreaterOrEqualT] with [*time.Time].
//
// [GreaterOrEqualT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
// use [GreaterOrEqual] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.GreaterOrEqualT(t, 2, 1)
//	assertions.GreaterOrEqualT(t, 2, 2)
//	assertions.GreaterOrEqualT(t, "b", "a")
//	assertions.GreaterOrEqualT(t, "b", "b")
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
func GreaterOrEqualT[Orderable Ordered](t T, e1, e2 Orderable, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}

	result := compareOrderedWithAny(e1, e2)
	if result >= 0 {
		return true
	}

	return Fail(t, fmt.Sprintf("\"%v\" is not greater than or equal to \"%v\"", e1, e2), msgAndArgs...)
}

// Less asserts that the first element is strictly less than the second.
//
// Both elements must be of the same type in the [reflect.Kind] sense.
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.Less(t, 1, 2)
//	assertions.Less(t, float64(1), float64(2))
//	assertions.Less(t, "a", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
func Less(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not less than \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareLess}, failMessage, msgAndArgs...)
}

// LessT asserts that for two elements of the same type, the first element is strictly less than the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [LessT] with [*time.Time].
//
// [LessT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
// use [Less] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.LessT(t, 1, 2)
//	assertions.LessT(t, float64(1), float64(2))
//	assertions.LessT(t, "a", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
func LessT[Orderable Ordered](t T, e1, e2 Orderable, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}

	result := compareOrderedWithAny(e1, e2)
	if result < 0 {
		return true
	}

	return Fail(t, fmt.Sprintf("\"%v\" is not less than \"%v\"", e1, e2), msgAndArgs...)
}

// LessOrEqual asserts that the first element is less than or equal to the second.
//
// # Usage
//
//	assertions.LessOrEqual(t, 1, 2)
//	assertions.LessOrEqual(t, 2, 2)
//	assertions.LessOrEqual(t, "a", "b")
//	assertions.LessOrEqual(t, "b", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
func LessOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not less than or equal to \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareLess, compareEqual}, failMessage, msgAndArgs...)
}

// LessOrEqualT asserts that for two elements of the same type, the first element is less than or equal to the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [LessOrEqualT] with [*time.Time].
//
// [LessOrEqualT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
// use [LessOrEqual] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you should use [LessOrEqual] instead.
//
// # Usage
//
//	assertions.LessOrEqualT(t, 1, 2)
//	assertions.LessOrEqualT(t, 2, 2)
//	assertions.LessOrEqualT(t, "a", "b")
//	assertions.LessOrEqualT(t, "b", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
func LessOrEqualT[Orderable Ordered](t T, e1, e2 Orderable, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}

	result := compareOrderedWithAny(e1, e2)
	if result <= 0 {
		return true
	}

	return Fail(t, fmt.Sprintf("\"%v\" is not less than or equal to \"%v\"", e1, e2), msgAndArgs...)
}

// Positive asserts that the specified element is strictly positive.
//
// # Usage
//
//	assertions.Positive(t, 1)
//	assertions.Positive(t, 1.23)
//
// # Examples
//
//	success: 1
//	failure: -1
func Positive(t T, e any, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}
	zero := reflect.Zero(reflect.TypeOf(e))
	failMessage := fmt.Sprintf("\"%v\" is not positive", e)
	return compareTwoValues(t, e, zero.Interface(), []compareResult{compareGreater}, failMessage, msgAndArgs...)
}

// PositiveT asserts that the specified element of a signed numeric type is strictly positive.
//
// # Usage
//
//	assertions.PositiveT(t, 1)
//	assertions.PositiveT(t, 1.23)
//
// # Examples
//
//	success: 1
//	failure: -1
func PositiveT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}

	result := cmp.Compare(e, SignedNumber(0))
	if result > 0 {
		return true
	}

	return Fail(t, fmt.Sprintf("\"%v\" is not positive", e), msgAndArgs...)
}

// Negative asserts that the specified element is strictly negative.
//
// # Usage
//
//	assertions.Negative(t, -1)
//	assertions.Negative(t, -1.23)
//
// # Examples
//
//	success: -1
//	failure: 1
func Negative(t T, e any, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}
	zero := reflect.Zero(reflect.TypeOf(e))
	failMessage := fmt.Sprintf("\"%v\" is not negative", e)
	return compareTwoValues(t, e, zero.Interface(), []compareResult{compareLess}, failMessage, msgAndArgs...)
}

// NegativeT asserts that the specified element of a signed numeric type is strictly negative.
//
// # Usage
//
//	assertions.NegativeT(t, -1)
//	assertions.NegativeT(t, -1.23)
//
// # Examples
//
//	success: -1
//	failure: 1
func NegativeT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) bool {
	// Domain: comparison
	if h, ok := t.(H); ok {
		h.Helper()
	}

	result := cmp.Compare(e, SignedNumber(0))
	if result < 0 {
		return true
	}

	return Fail(t, fmt.Sprintf("\"%v\" is not negative", e), msgAndArgs...)
}

type compareResult = int

const (
	compareLess compareResult = iota - 1
	compareEqual
	compareGreater
)

func compareTwoValues(t T, e1 any, e2 any, allowedComparesResults []compareResult, failMessage string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	result, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf(`Can not compare type "%T"`, e1), msgAndArgs...)
	}

	if !containsValue(allowedComparesResults, result) {
		return Fail(t, failMessage, msgAndArgs...)
	}

	return true
}

func containsValue(values []compareResult, value compareResult) bool {
	return slices.Contains(values, value)
}

func compare(obj1, obj2 any, kind reflect.Kind) (compareResult, bool) {
	obj1Value := reflect.ValueOf(obj1)
	obj2Value := reflect.ValueOf(obj2)

	switch kind {
	case reflect.Int:
		intobj1 := convertReflectValue[int](obj1, obj1Value)
		intobj2 := convertReflectValue[int](obj2, obj2Value)

		return cmp.Compare(intobj1, intobj2), true
	case reflect.Int8:
		int8obj1 := convertReflectValue[int8](obj1, obj1Value)
		int8obj2 := convertReflectValue[int8](obj2, obj2Value)

		return cmp.Compare(int8obj1, int8obj2), true
	case reflect.Int16:
		int16obj1 := convertReflectValue[int16](obj1, obj1Value)
		int16obj2 := convertReflectValue[int16](obj2, obj2Value)

		return cmp.Compare(int16obj1, int16obj2), true
	case reflect.Int32:
		int32obj1 := convertReflectValue[int32](obj1, obj1Value)
		int32obj2 := convertReflectValue[int32](obj2, obj2Value)

		return cmp.Compare(int32obj1, int32obj2), true
	case reflect.Int64:
		int64obj1 := convertReflectValue[int64](obj1, obj1Value)
		int64obj2 := convertReflectValue[int64](obj2, obj2Value)

		return cmp.Compare(int64obj1, int64obj2), true
	case reflect.Uint:
		uintobj1 := convertReflectValue[uint](obj1, obj1Value)
		uintobj2 := convertReflectValue[uint](obj2, obj2Value)

		return cmp.Compare(uintobj1, uintobj2), true
	case reflect.Uint8:
		uint8obj1 := convertReflectValue[uint8](obj1, obj1Value)
		uint8obj2 := convertReflectValue[uint8](obj2, obj2Value)

		return cmp.Compare(uint8obj1, uint8obj2), true
	case reflect.Uint16:
		uint16obj1 := convertReflectValue[uint16](obj1, obj1Value)
		uint16obj2 := convertReflectValue[uint16](obj2, obj2Value)

		return cmp.Compare(uint16obj1, uint16obj2), true
	case reflect.Uint32:
		uint32obj1 := convertReflectValue[uint32](obj1, obj1Value)
		uint32obj2 := convertReflectValue[uint32](obj2, obj2Value)

		return cmp.Compare(uint32obj1, uint32obj2), true
	case reflect.Uint64:
		uint64obj1 := convertReflectValue[uint64](obj1, obj1Value)
		uint64obj2 := convertReflectValue[uint64](obj2, obj2Value)

		return cmp.Compare(uint64obj1, uint64obj2), true
	case reflect.Float32:
		float32obj1 := convertReflectValue[float32](obj1, obj1Value)
		float32obj2 := convertReflectValue[float32](obj2, obj2Value)

		return cmp.Compare(float32obj1, float32obj2), true
	case reflect.Float64:
		float64obj1 := convertReflectValue[float64](obj1, obj1Value)
		float64obj2 := convertReflectValue[float64](obj2, obj2Value)

		return cmp.Compare(float64obj1, float64obj2), true
	case reflect.String:
		stringobj1 := convertReflectValue[string](obj1, obj1Value)
		stringobj2 := convertReflectValue[string](obj2, obj2Value)

		return cmp.Compare(stringobj1, stringobj2), true

	// Check for known struct types we can check for compare results.
	case reflect.Struct:
		return compareStruct(obj1, obj2, obj1Value, obj2Value)
	case reflect.Slice:
		return compareSlice(obj1, obj2, obj1Value, obj2Value)
	case reflect.Uintptr:
		uintptrobj1 := convertReflectValue[uintptr](obj1, obj1Value)
		uintptrobj2 := convertReflectValue[uintptr](obj2, obj2Value)

		return cmp.Compare(uintptrobj1, uintptrobj2), true
	default:
		return compareEqual, false
	}
}

func compareStruct(obj1, obj2 any, obj1Value, obj2Value reflect.Value) (compareResult, bool) {
	// all structs enter here. We're not interested in most types.
	if !obj1Value.CanConvert(reflect.TypeFor[time.Time]()) {
		return compareEqual, false
	}

	// time.Time can be compared
	timeobj1 := convertReflectValue[time.Time](obj1, obj1Value)
	timeobj2 := convertReflectValue[time.Time](obj2, obj2Value)

	return timeobj1.Compare(timeobj2), true
}

func compareSlice(obj1, obj2 any, obj1Value, obj2Value reflect.Value) (compareResult, bool) {
	// we only care about the []byte type.
	if !obj1Value.CanConvert(reflect.TypeFor[[]byte]()) {
		return compareEqual, false
	}

	// []byte can be compared
	bytesobj1 := convertReflectValue[[]byte](obj1, obj1Value)
	bytesobj2 := convertReflectValue[[]byte](obj2, obj2Value)

	return bytes.Compare(bytesobj1, bytesobj2), true
}

func convertReflectValue[V any](obj any, value reflect.Value) V { //nolint:ireturn // false positive
	// we try and avoid calling [reflect.Value.Convert()] whenever possible,
	// as this has a pretty big performance impact
	converted, ok := obj.(V)
	if !ok {
		converted, ok = value.Convert(reflect.TypeFor[V]()).Interface().(V)
		if !ok {
			// should never get there
			panic("internal error: expected that reflect.Value.Convert yields its target type")
		}
	}

	return converted
}

// compareOrderedWithAny compares two [Ordered] values.
//
// This is an internal function that should only be called with actually [Ordered] types,
// even though it doesn't enforce this a build time.
//
//nolint:forcetypeassert // e2 is guaranteed to be of the same type as e1
func compareOrderedWithAny[Orderable any](e1, e2 Orderable) int {
	v := any(e1)
	o := any(e2)
	switch value := v.(type) {
	case time.Time:
		other := o.(time.Time)
		return value.Compare(other)
	case []byte:
		other := o.([]byte)
		return bytes.Compare(value, other)
	case string:
		other := o.(string)
		return cmp.Compare(value, other)
	case int:
		other := o.(int)
		return cmp.Compare(value, other)
	case int8:
		other := o.(int8)
		return cmp.Compare(value, other)
	case int16:
		other := o.(int16)
		return cmp.Compare(value, other)
	case int32:
		other := o.(int32)
		return cmp.Compare(value, other)
	case int64:
		other := o.(int64)
		return cmp.Compare(value, other)
	case uint:
		other := o.(uint)
		return cmp.Compare(value, other)
	case uint8:
		other := o.(uint8)
		return cmp.Compare(value, other)
	case uint16:
		other := o.(uint16)
		return cmp.Compare(value, other)
	case uint32:
		other := o.(uint32)
		return cmp.Compare(value, other)
	case uint64:
		other := o.(uint64)
		return cmp.Compare(value, other)
	case float32:
		other := o.(float32)
		return cmp.Compare(value, other)
	case float64:
		other := o.(float64)
		return cmp.Compare(value, other)
	case uintptr:
		other := o.(uintptr)
		return cmp.Compare(value, other)
	default:
		// we have a custom type: convert with reflection.
		// We have less edge cases to guard than when comparing with the purely reflection-based call.
		e1Kind := reflect.ValueOf(e1).Kind()
		result, ok := compare(e1, e2, e1Kind)
		if !ok {
			// should never get there
			panic(fmt.Errorf("internal error: expected that reflect.Value.Convert yields its target type for %T", e1))
		}

		return result
	}
}
