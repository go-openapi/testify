// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"reflect"
	"slices"
	"testing"
	"time"
)

func TestCompareUnexportedImplementationDetails(t *testing.T) {
	t.Parallel()

	t.Run("compare", testCompare())
	t.Run("containsValue", testContainsValue())
	t.Run("compareTwoValues - different types", testCompareTwoValuesDifferentTypes())
	t.Run("compareTwoValues - not comparable", testCompareTwoValuesNotComparable())
	t.Run("compareTwoValues - compare result", testCompareTwoValuesCorrectCompareResult())
}

func testCompare() func(*testing.T) {
	return func(t *testing.T) {
		for currCase := range compareImplCompareCases() {
			t.Run("compare", func(t *testing.T) {
				t.Parallel()

				resLess, isComparable := compare(currCase.less, currCase.greater, reflect.ValueOf(currCase.less).Kind())
				if !isComparable {
					t.Error("object should be comparable for type " + currCase.cType)
				}

				if resLess != compareLess {
					t.Errorf("object less (%v) should be less than greater (%v) for type "+currCase.cType,
						currCase.less, currCase.greater)
				}

				resGreater, isComparable := compare(currCase.greater, currCase.less, reflect.ValueOf(currCase.less).Kind())
				if !isComparable {
					t.Error("object are comparable for type " + currCase.cType)
				}

				if resGreater != compareGreater {
					t.Errorf("object greater should be greater than less for type %s", currCase.cType)
				}

				resEqual, isComparable := compare(currCase.less, currCase.less, reflect.ValueOf(currCase.less).Kind())
				if !isComparable {
					t.Errorf("object are comparable for type %s", currCase.cType)
				}

				if resEqual != 0 {
					t.Errorf("objects should be equal for type %s", currCase.cType)
				}
			})
		}
	}
}

func testContainsValue() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for currCase := range compareContainsValueCases() {
			result := containsValue(currCase.values, currCase.value)
			Equal(t, currCase.result, result)
		}
	}
}

func testCompareTwoValuesDifferentTypes() func(*testing.T) {
	return func(t *testing.T) {
		for currCase := range compareTwoValuesDifferentTypesCases() {
			t.Run("different types should not be comparable", func(t *testing.T) {
				t.Parallel()
				mock := new(testing.T)

				result := compareTwoValues(mock, currCase.v1, currCase.v2, []compareResult{compareLess, compareEqual, compareGreater}, "testFailMessage")
				False(t, result)
			})
		}
	}
}

func testCompareTwoValuesNotComparable() func(*testing.T) {
	return func(t *testing.T) {
		for currCase := range compareTwoValuesNotComparableCases() {
			t.Run("should not be comparable", func(t *testing.T) {
				t.Parallel()
				mock := new(testing.T)

				result := compareTwoValues(mock, currCase.v1, currCase.v2, []compareResult{compareLess, compareEqual, compareGreater}, "testFailMessage")
				False(t, result)
			})
		}
	}
}

func testCompareTwoValuesCorrectCompareResult() func(*testing.T) {
	return func(t *testing.T) {
		for currCase := range compareTwoValuesCorrectResultCases() {
			t.Run("should be comparable", func(t *testing.T) {
				t.Parallel()
				mock := new(testing.T)

				result := compareTwoValues(mock, currCase.v1, currCase.v2, currCase.allowedResults, "testFailMessage")
				True(t, result)
			})
		}
	}
}

type compareImplCompareCase struct {
	less    any
	greater any
	cType   string
}

func compareImplCompareCases() iter.Seq[compareImplCompareCase] {
	type customString string
	type customInt int
	type customInt8 int8
	type customInt16 int16
	type customInt32 int32
	type customInt64 int64
	type customUInt uint
	type customUInt8 uint8
	type customUInt16 uint16
	type customUInt32 uint32
	type customUInt64 uint64
	type customFloat32 float32
	type customFloat64 float64
	type customUintptr uintptr
	type customTime time.Time
	type customBytes []byte

	return slices.Values([]compareImplCompareCase{
		{less: customString("a"), greater: customString("b"), cType: "string"},
		{less: "a", greater: "b", cType: "string"},
		{less: customInt(1), greater: customInt(2), cType: "int"},
		{less: int(1), greater: int(2), cType: "int"},
		{less: customInt8(1), greater: customInt8(2), cType: "int8"},
		{less: int8(1), greater: int8(2), cType: "int8"},
		{less: customInt16(1), greater: customInt16(2), cType: "int16"},
		{less: int16(1), greater: int16(2), cType: "int16"},
		{less: customInt32(1), greater: customInt32(2), cType: "int32"},
		{less: int32(1), greater: int32(2), cType: "int32"},
		{less: customInt64(1), greater: customInt64(2), cType: "int64"},
		{less: int64(1), greater: int64(2), cType: "int64"},
		{less: customUInt(1), greater: customUInt(2), cType: "uint"},
		{less: uint8(1), greater: uint8(2), cType: "uint8"},
		{less: customUInt8(1), greater: customUInt8(2), cType: "uint8"},
		{less: uint16(1), greater: uint16(2), cType: "uint16"},
		{less: customUInt16(1), greater: customUInt16(2), cType: "uint16"},
		{less: uint32(1), greater: uint32(2), cType: "uint32"},
		{less: customUInt32(1), greater: customUInt32(2), cType: "uint32"},
		{less: uint64(1), greater: uint64(2), cType: "uint64"},
		{less: customUInt64(1), greater: customUInt64(2), cType: "uint64"},
		{less: float32(1.23), greater: float32(2.34), cType: "float32"},
		{less: customFloat32(1.23), greater: customFloat32(2.23), cType: "float32"},
		{less: float64(1.23), greater: float64(2.34), cType: "float64"},
		{less: customFloat64(1.23), greater: customFloat64(2.34), cType: "float64"},
		{less: uintptr(1), greater: uintptr(2), cType: "uintptr"},
		{less: customUintptr(1), greater: customUintptr(2), cType: "uint64"},
		{less: time.Now(), greater: time.Now().Add(time.Hour), cType: "time.Time"},
		{
			// using time.Local is ok in this context: this is precisely the goal of this test
			less:    time.Date(2024, 0, 0, 0, 0, 0, 0, time.Local), //nolint:gosmopolitan // ok. See above
			greater: time.Date(2263, 0, 0, 0, 0, 0, 0, time.Local), //nolint:gosmopolitan // ok. See above
			cType:   "time.Time",
		},
		{less: customTime(time.Now()), greater: customTime(time.Now().Add(time.Hour)), cType: "time.Time"},
		{less: []byte{1, 1}, greater: []byte{1, 2}, cType: "[]byte"},
		{less: customBytes([]byte{1, 1}), greater: customBytes([]byte{1, 2}), cType: "[]byte"},
	})
}

type compareTwoValuesCase struct {
	v1 any
	v2 any
	// compareResult  bool
	allowedResults []compareResult
}

func compareTwoValuesDifferentTypesCases() iter.Seq[compareTwoValuesCase] {
	return slices.Values([]compareTwoValuesCase{
		{v1: 123, v2: "abc"},
		{v1: "abc", v2: 123456},
		{v1: float64(12), v2: "123"},
		{v1: "float(12)", v2: float64(1)},
	})
}

func compareTwoValuesNotComparableCases() iter.Seq[compareTwoValuesCase] {
	type CompareStruct struct{}
	return slices.Values([]compareTwoValuesCase{
		{v1: CompareStruct{}, v2: CompareStruct{}},
		{v1: map[string]int{}, v2: map[string]int{}},
		{v1: make([]int, 5), v2: make([]int, 5)},
	})
}

func compareTwoValuesCorrectResultCases() iter.Seq[compareTwoValuesCase] {
	return slices.Values([]compareTwoValuesCase{
		{v1: 1, v2: 2, allowedResults: []compareResult{compareLess}},
		{v1: 1, v2: 2, allowedResults: []compareResult{compareLess, compareEqual}},
		{v1: 2, v2: 2, allowedResults: []compareResult{compareGreater, compareEqual}},
		{v1: 2, v2: 2, allowedResults: []compareResult{compareEqual}},
		{v1: 2, v2: 1, allowedResults: []compareResult{compareEqual, compareGreater}},
		{v1: 2, v2: 1, allowedResults: []compareResult{compareGreater}},
	})
}

type compareContainsValueCase struct {
	values []compareResult
	value  compareResult
	result bool
}

func compareContainsValueCases() iter.Seq[compareContainsValueCase] {
	return slices.Values([]compareContainsValueCase{
		{values: []compareResult{compareGreater}, value: compareGreater, result: true},
		{values: []compareResult{compareGreater, compareLess}, value: compareGreater, result: true},
		{values: []compareResult{compareGreater, compareLess}, value: compareLess, result: true},
		{values: []compareResult{compareGreater, compareLess}, value: compareEqual, result: false},
	})
}
