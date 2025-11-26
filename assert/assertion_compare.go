package assert

import (
	"bytes"
	"cmp"
	"fmt"
	"reflect"
	"slices"
	"time"
)

// Deprecated: CompareType has only ever been for internal use and has accidentally been published since v1.6.0. Do not use it.
type CompareType = compareResult

type compareResult int

const (
	compareLess compareResult = iota - 1
	compareEqual
	compareGreater
)

func compare(obj1, obj2 any, kind reflect.Kind) (compareResult, bool) {
	obj1Value := reflect.ValueOf(obj1)
	obj2Value := reflect.ValueOf(obj2)

	switch kind {
	case reflect.Int:
		intobj1 := convertReflectValue[int](obj1, obj1Value)
		intobj2 := convertReflectValue[int](obj2, obj2Value)

		return compareOrdered(intobj1, intobj2)
	case reflect.Int8:
		int8obj1 := convertReflectValue[int8](obj1, obj1Value)
		int8obj2 := convertReflectValue[int8](obj2, obj2Value)

		return compareOrdered(int8obj1, int8obj2)
	case reflect.Int16:
		int16obj1 := convertReflectValue[int16](obj1, obj1Value)
		int16obj2 := convertReflectValue[int16](obj2, obj2Value)

		return compareOrdered(int16obj1, int16obj2)
	case reflect.Int32:
		int32obj1 := convertReflectValue[int32](obj1, obj1Value)
		int32obj2 := convertReflectValue[int32](obj2, obj2Value)

		return compareOrdered(int32obj1, int32obj2)
	case reflect.Int64:
		int64obj1 := convertReflectValue[int64](obj1, obj1Value)
		int64obj2 := convertReflectValue[int64](obj2, obj2Value)

		return compareOrdered(int64obj1, int64obj2)
	case reflect.Uint:
		uintobj1 := convertReflectValue[uint](obj1, obj1Value)
		uintobj2 := convertReflectValue[uint](obj2, obj2Value)

		return compareOrdered(uintobj1, uintobj2)
	case reflect.Uint8:
		uint8obj1 := convertReflectValue[uint8](obj1, obj1Value)
		uint8obj2 := convertReflectValue[uint8](obj2, obj2Value)

		return compareOrdered(uint8obj1, uint8obj2)
	case reflect.Uint16:
		uint16obj1 := convertReflectValue[uint16](obj1, obj1Value)
		uint16obj2 := convertReflectValue[uint16](obj2, obj2Value)

		return compareOrdered(uint16obj1, uint16obj2)
	case reflect.Uint32:
		uint32obj1 := convertReflectValue[uint32](obj1, obj1Value)
		uint32obj2 := convertReflectValue[uint32](obj2, obj2Value)

		return compareOrdered(uint32obj1, uint32obj2)
	case reflect.Uint64:
		uint64obj1 := convertReflectValue[uint64](obj1, obj1Value)
		uint64obj2 := convertReflectValue[uint64](obj2, obj2Value)

		return compareOrdered(uint64obj1, uint64obj2)
	case reflect.Float32:
		float32obj1 := convertReflectValue[float32](obj1, obj1Value)
		float32obj2 := convertReflectValue[float32](obj2, obj2Value)

		return compareOrdered(float32obj1, float32obj2)
	case reflect.Float64:
		float64obj1 := convertReflectValue[float64](obj1, obj1Value)
		float64obj2 := convertReflectValue[float64](obj2, obj2Value)

		return compareOrdered(float64obj1, float64obj2)
	case reflect.String:
		stringobj1 := convertReflectValue[string](obj1, obj1Value)
		stringobj2 := convertReflectValue[string](obj2, obj2Value)

		return compareOrdered(stringobj1, stringobj2)

	// Check for known struct types we can check for compare results.
	case reflect.Struct:
		return compareStruct(obj1, obj2, obj1Value, obj2Value)
	case reflect.Slice:
		return compareSlice(obj1, obj2, obj1Value, obj2Value)
	case reflect.Uintptr:
		uintptrobj1 := convertReflectValue[string](obj1, obj1Value)
		uintptrobj2 := convertReflectValue[string](obj2, obj2Value)

		return compareOrdered(uintptrobj1, uintptrobj2)
	default:
		return compareEqual, false
	}
}

func compareOrdered[T cmp.Ordered](obj1, obj2 T) (compareResult, bool) {
	return compareResult(cmp.Compare(obj1, obj2)), true
}

func compareStruct(obj1, obj2 any, obj1Value, obj2Value reflect.Value) (compareResult, bool) {
	// all structs enter here. We're not interested in most types.
	if !obj1Value.CanConvert(reflect.TypeFor[time.Time]()) {
		return compareEqual, false
	}

	// time.Time can be compared
	timeobj1 := convertReflectValue[time.Time](obj1, obj1Value)
	timeobj2 := convertReflectValue[time.Time](obj2, obj2Value)

	return compareTime(timeobj1, timeobj2)
}

func compareSlice(obj1, obj2 any, obj1Value, obj2Value reflect.Value) (compareResult, bool) {
	// we only care about the []byte type.
	if !obj1Value.CanConvert(reflect.TypeFor[[]byte]()) {
		return compareEqual, false
	}

	// []byte can be compared
	bytesobj1 := convertReflectValue[[]byte](obj1, obj1Value)
	bytesobj2 := convertReflectValue[[]byte](obj2, obj2Value)

	return compareBytes(bytesobj1, bytesobj2)
}

func compareTime(obj1, obj2 time.Time) (compareResult, bool) {
	switch {
	case obj1.Before(obj2):
		return compareLess, true
	case obj1.Equal(obj2):
		return compareEqual, true
	default:
		return compareGreater, true
	}
}

func compareBytes(obj1, obj2 []byte) (compareResult, bool) {
	return compareResult(bytes.Compare(obj1, obj2)), true
}

func convertReflectValue[T any](obj any, value reflect.Value) T { //nolint:ireturn // false positive
	// we try and avoid calling [reflect.Value.Convert()] whenever possible,
	// as this has a pretty big performance impact
	converted, ok := obj.(T)
	if !ok {
		converted, ok = value.Convert(reflect.TypeFor[T]()).Interface().(T)
		if !ok {
			panic("internal error: expected that reflect.Value.Convert yields its target type")
		}
	}

	return converted
}

// Greater asserts that the first element is greater than the second
//
//	assert.Greater(t, 2, 1)
//	assert.Greater(t, float64(2), float64(1))
//	assert.Greater(t, "b", "a")
func Greater(t TestingT, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not greater than \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareGreater}, failMessage, msgAndArgs...)
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second
//
//	assert.GreaterOrEqual(t, 2, 1)
//	assert.GreaterOrEqual(t, 2, 2)
//	assert.GreaterOrEqual(t, "b", "a")
//	assert.GreaterOrEqual(t, "b", "b")
func GreaterOrEqual(t TestingT, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not greater than or equal to \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareGreater, compareEqual}, failMessage, msgAndArgs...)
}

// Less asserts that the first element is less than the second
//
//	assert.Less(t, 1, 2)
//	assert.Less(t, float64(1), float64(2))
//	assert.Less(t, "a", "b")
func Less(t TestingT, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not less than \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareLess}, failMessage, msgAndArgs...)
}

// LessOrEqual asserts that the first element is less than or equal to the second
//
//	assert.LessOrEqual(t, 1, 2)
//	assert.LessOrEqual(t, 2, 2)
//	assert.LessOrEqual(t, "a", "b")
//	assert.LessOrEqual(t, "b", "b")
func LessOrEqual(t TestingT, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	failMessage := fmt.Sprintf("\"%v\" is not less than or equal to \"%v\"", e1, e2)
	return compareTwoValues(t, e1, e2, []compareResult{compareLess, compareEqual}, failMessage, msgAndArgs...)
}

// Positive asserts that the specified element is positive
//
//	assert.Positive(t, 1)
//	assert.Positive(t, 1.23)
func Positive(t TestingT, e any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	zero := reflect.Zero(reflect.TypeOf(e))
	failMessage := fmt.Sprintf("\"%v\" is not positive", e)
	return compareTwoValues(t, e, zero.Interface(), []compareResult{compareGreater}, failMessage, msgAndArgs...)
}

// Negative asserts that the specified element is negative
//
//	assert.Negative(t, -1)
//	assert.Negative(t, -1.23)
func Negative(t TestingT, e any, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}
	zero := reflect.Zero(reflect.TypeOf(e))
	failMessage := fmt.Sprintf("\"%v\" is not negative", e)
	return compareTwoValues(t, e, zero.Interface(), []compareResult{compareLess}, failMessage, msgAndArgs...)
}

func compareTwoValues(t TestingT, e1 any, e2 any, allowedComparesResults []compareResult, failMessage string, msgAndArgs ...any) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	e1Kind := reflect.ValueOf(e1).Kind()
	e2Kind := reflect.ValueOf(e2).Kind()
	if e1Kind != e2Kind {
		return Fail(t, "Elements should be the same type", msgAndArgs...)
	}

	compareResult, isComparable := compare(e1, e2, e1Kind)
	if !isComparable {
		return Fail(t, fmt.Sprintf(`Can not compare type "%T"`, e1), msgAndArgs...)
	}

	if !containsValue(allowedComparesResults, compareResult) {
		return Fail(t, failMessage, msgAndArgs...)
	}

	return true
}

func containsValue(values []compareResult, value compareResult) bool {
	return slices.Contains(values, value)
}
