// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"testing"
	"time"
)

func TestEqualUnexportedImplementationDetails(t *testing.T) {
	t.Parallel()

	t.Run("samePointers", testSamePointers())
	t.Run("formatUnequalValue", testFormatUnequalValues())
	t.Run("isEmpty", testIsEmpty())
	t.Run("validateEqualArgs", testValidateEqualArgs())
}

func testSamePointers() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for tt := range equalSamePointersCases() {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				same, ok := samePointers(tt.args.first, tt.args.second)
				tt.same(t, same)
				tt.ok(t, ok)
			})
		}
	}
}

func testFormatUnequalValues() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for tt := range formatUnequalCases() {
			t.Run(tt.testName, func(t *testing.T) {
				t.Parallel()

				expected, actual := formatUnequalValues(tt.unequalExpected, tt.unequalActual)
				Equal(t, tt.expectedExpected, expected, tt.testName)
				Equal(t, tt.expectedActual, actual, tt.testName)
			})
		}
	}
}

func testIsEmpty() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		chWithValue := make(chan struct{}, 1)
		chWithValue <- struct{}{}

		True(t, isEmpty(""))
		True(t, isEmpty(nil))
		True(t, isEmpty(error(nil)))
		True(t, isEmpty((*int)(nil)))
		True(t, isEmpty((*string)(nil)))
		True(t, isEmpty(new(string)))
		True(t, isEmpty([]string{}))
		True(t, isEmpty([]string(nil)))
		True(t, isEmpty([]byte(nil)))
		True(t, isEmpty([]byte{}))
		True(t, isEmpty([]byte("")))
		True(t, isEmpty([]bool(nil)))
		True(t, isEmpty([]bool{}))
		True(t, isEmpty([]any(nil)))
		True(t, isEmpty([]any{}))
		True(t, isEmpty(struct{}{}))
		True(t, isEmpty(&struct{}{}))
		True(t, isEmpty(struct{ A int }{A: 0}))
		True(t, isEmpty(struct{ a int }{a: 0}))
		True(t, isEmpty(struct {
			a int
			B int
		}{a: 0, B: 0}))
		True(t, isEmpty(0))
		True(t, isEmpty(int(0)))
		True(t, isEmpty(int8(0)))
		True(t, isEmpty(int16(0)))
		True(t, isEmpty(uint16(0)))
		True(t, isEmpty(int32(0)))
		True(t, isEmpty(uint32(0)))
		True(t, isEmpty(int64(0)))
		True(t, isEmpty(uint64(0)))
		True(t, isEmpty('\u0000')) // rune => int32
		True(t, isEmpty(float32(0)))
		True(t, isEmpty(float64(0)))
		True(t, isEmpty(0i))   // complex
		True(t, isEmpty(0.0i)) // complex
		True(t, isEmpty(false))
		True(t, isEmpty(new(bool)))
		True(t, isEmpty(map[string]string{}))
		True(t, isEmpty(map[string]string(nil)))
		True(t, isEmpty(new(time.Time)))
		True(t, isEmpty(time.Time{}))
		True(t, isEmpty(make(chan struct{})))
		True(t, isEmpty(chan struct{}(nil)))
		True(t, isEmpty(chan<- struct{}(nil)))
		True(t, isEmpty(make(chan struct{})))
		True(t, isEmpty(make(chan<- struct{})))
		True(t, isEmpty(make(chan struct{}, 1)))
		True(t, isEmpty(make(chan<- struct{}, 1)))
		True(t, isEmpty([1]int{0}))
		True(t, isEmpty([2]int{0, 0}))
		True(t, isEmpty([8]int{}))
		True(t, isEmpty([...]int{7: 0}))
		True(t, isEmpty([...]bool{false, false}))
		True(t, isEmpty(errors.New(""))) // BEWARE
		True(t, isEmpty([]error{}))
		True(t, isEmpty([]error(nil)))
		True(t, isEmpty(&[1]int{0}))
		True(t, isEmpty(&[2]int{0, 0}))
		False(t, isEmpty("something"))
		False(t, isEmpty(errors.New("something")))
		False(t, isEmpty([]string{"something"}))
		False(t, isEmpty(1))
		False(t, isEmpty(int(1)))
		False(t, isEmpty(uint(1)))
		False(t, isEmpty(byte(1)))
		False(t, isEmpty(int8(1)))
		False(t, isEmpty(uint8(1)))
		False(t, isEmpty(int16(1)))
		False(t, isEmpty(uint16(1)))
		False(t, isEmpty(int32(1)))
		False(t, isEmpty(uint32(1)))
		False(t, isEmpty(int64(1)))
		False(t, isEmpty(uint64(1)))
		False(t, isEmpty('A')) // rune => int32
		False(t, isEmpty(true))
		False(t, isEmpty(1.0))
		False(t, isEmpty(1i))            // complex
		False(t, isEmpty([]byte{0}))     // elements values are ignored for slices
		False(t, isEmpty([]byte{0, 0}))  // elements values are ignored for slices
		False(t, isEmpty([]string{""}))  // elements values are ignored for slices
		False(t, isEmpty([]string{"a"})) // elements values are ignored for slices
		False(t, isEmpty([]bool{false})) // elements values are ignored for slices
		False(t, isEmpty([]bool{true}))  // elements values are ignored for slices
		False(t, isEmpty([]error{errors.New("xxx")}))
		False(t, isEmpty([]error{nil}))            // BEWARE
		False(t, isEmpty([]error{errors.New("")})) // BEWARE
		False(t, isEmpty(map[string]string{"Hello": "World"}))
		False(t, isEmpty(map[string]string{"": ""}))
		False(t, isEmpty(map[string]string{"foo": ""}))
		False(t, isEmpty(map[string]string{"": "foo"}))
		False(t, isEmpty(chWithValue))
		False(t, isEmpty([1]bool{true}))
		False(t, isEmpty([2]bool{false, true}))
		False(t, isEmpty([...]bool{10: true}))
		False(t, isEmpty([]int{0}))
		False(t, isEmpty([]int{42}))
		False(t, isEmpty([1]int{42}))
		False(t, isEmpty([2]int{0, 42}))
		False(t, isEmpty(&[1]int{42}))
		False(t, isEmpty(&[2]int{0, 42}))
		False(t, isEmpty([1]*int{new(int)})) // array elements must be the zero value, not any Empty value
		False(t, isEmpty(struct{ A int }{A: 42}))
		False(t, isEmpty(struct{ a int }{a: 42}))
		False(t, isEmpty(struct{ a *int }{a: new(int)})) // fields must be the zero value, not any Empty value
		False(t, isEmpty(struct{ a []int }{a: []int{}})) // fields must be the zero value, not any Empty value
		False(t, isEmpty(struct {
			a int
			B int
		}{a: 0, B: 42}))
		False(t, isEmpty(struct {
			a int
			B int
		}{a: 42, B: 0}))
	}
}

func testValidateEqualArgs() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if validateEqualArgs(func() {}, func() {}) == nil {
			t.Error("non-nil functions should error")
		}

		if validateEqualArgs(func() {}, func() {}) == nil {
			t.Error("non-nil functions should error")
		}

		if validateEqualArgs(nil, nil) != nil {
			t.Error("nil functions are equal")
		}
	}
}

type formatUnequalCase struct {
	unequalExpected  any
	unequalActual    any
	expectedExpected string
	expectedActual   string
	testName         string
}

func formatUnequalCases() iter.Seq[formatUnequalCase] {
	type testStructType struct {
		Val string
	}

	return slices.Values([]formatUnequalCase{
		{"foo", "bar", `"foo"`, `"bar"`, "value should not include type"},
		{123, 123, `123`, `123`, "value should not include type"},
		{int64(123), int32(123), `int64(123)`, `int32(123)`, "value should include type"},
		{int64(123), nil, `int64(123)`, `<nil>(<nil>)`, "value should include type"},
		{
			unequalExpected:  &testStructType{Val: "test"},
			unequalActual:    &testStructType{Val: "test"},
			expectedExpected: fmt.Sprintf(`&%s.testStructType{Val:"test"}`, shortpkg),
			expectedActual:   fmt.Sprintf(`&%s.testStructType{Val:"test"}`, shortpkg),
			testName:         "value should not include type annotation",
		},
		{uint(123), uint(124), `123`, `124`, "uint should print clean"},
		{uint8(123), uint8(124), `123`, `124`, "uint8 should print clean"},
		{uint16(123), uint16(124), `123`, `124`, "uint16 should print clean"},
		{uint32(123), uint32(124), `123`, `124`, "uint32 should print clean"},
		{uint64(123), uint64(124), `123`, `124`, "uint64 should print clean"},
	})
}
