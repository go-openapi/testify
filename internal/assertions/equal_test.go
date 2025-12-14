// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"iter"
	"os"
	"reflect"
	"regexp"
	"slices"
	"testing"
	"time"
)

const shortpkg = "assertions"

func TestEqualNotNil(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !NotNil(mock, new(AssertionTesterConformingObject)) {
		t.Error("NotNil should return true: object is not nil")
	}

	if NotNil(mock, nil) {
		t.Error("NotNil should return false: object is nil")
	}

	if NotNil(mock, (*struct{})(nil)) {
		t.Error("NotNil should return false: object is (*struct{})(nil)")
	}
}

func TestEqualNil(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Nil(mock, nil) {
		t.Error("Nil should return true: object is nil")
	}

	if !Nil(mock, (*struct{})(nil)) {
		t.Error("Nil should return true: object is (*struct{})(nil)")
	}

	if Nil(mock, new(AssertionTesterConformingObject)) {
		t.Error("Nil should return false: object is not nil")
	}
}

func TestEqualSameWithSliceTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	Same(mock, &[]int{}, &longSlice)
	Contains(t, mock.errorString(), `&[]int{0, 0, 0,`)
}

func TestEqualNotSameWithSliceTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	NotSame(mock, &longSlice, &longSlice)
	Contains(t, mock.errorString(), `&[]int{0, 0, 0,`)
}

func TestEqualNotEqualWithSliceTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	NotEqual(mock, longSlice, longSlice)
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Should not be: []int{0, 0, 0,`)
	Contains(t, mock.errorString(), `<... truncated>`)
}

func TestEqualNotEqualValuesWithSliceTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	NotEqualValues(mock, longSlice, longSlice)
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Should not be: []int{0, 0, 0,`)
	Contains(t, mock.errorString(), `<... truncated>`)
}

func TestEqual(t *testing.T) {
	t.Parallel()

	for c := range equalCases() {
		t.Run(fmt.Sprintf("Equal(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			t.Parallel()
			mock := new(testing.T)

			res := Equal(mock, c.expected, c.actual)
			if res != c.result {
				t.Errorf("Equal(%#v, %#v) should return %#v: %s", c.expected, c.actual, c.result, c.remark)
			}
		})
	}
}

func TestEqualSame(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if Same(mock, ptr(1), ptr(1)) {
		t.Error("Same should return false")
	}

	if Same(mock, 1, 1) {
		t.Error("Same should return false")
	}

	p := ptr(2)
	if Same(mock, p, *p) {
		t.Error("Same should return false")
	}

	if !Same(mock, p, p) {
		t.Error("Same should return true")
	}

	t.Run("same object, different type", func(t *testing.T) {
		type s struct {
			i int
		}
		type sPtr *s
		ps := &s{1}
		dps := sPtr(ps)
		if Same(mock, dps, ps) {
			t.Error("Same should return false")
		}
		expPat :=
			fmt.Sprintf(`expected: &%[1]s.s\{i:1\} \(%[1]s.sPtr\)\((0x[a-f0-9]+)\)\s*\n`, shortpkg) +
				fmt.Sprintf(`\s+actual  : &%[1]s.s\{i:1\} \(\*%[1]s.s\)\((0x[a-f0-9]+)\)`, shortpkg)
		Regexp(t, regexp.MustCompile(expPat), mock.errorString())
	})
}

func TestEqualNotSame(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !NotSame(mock, ptr(1), ptr(1)) {
		t.Error("NotSame should return true; different pointers")
	}

	if !NotSame(mock, 1, 1) {
		t.Error("NotSame should return true; constant inputs")
	}

	p := ptr(2)
	if !NotSame(mock, p, *p) {
		t.Error("NotSame should return true; mixed-type inputs")
	}

	if NotSame(mock, p, p) {
		t.Error("NotSame should return false")
	}
}

func TestEqualNotEqual(t *testing.T) {
	t.Parallel()

	for c := range equalNotEqualCases() {
		t.Run(fmt.Sprintf("NotEqual(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			t.Parallel()
			mock := new(testing.T)

			res := NotEqual(mock, c.expected, c.actual)

			if res != c.result {
				t.Errorf("NotEqual(%#v, %#v) should return %#v", c.expected, c.actual, c.result)
			}
		})
	}
}

func TestEqualValuesAndNotEqualValues(t *testing.T) {
	t.Parallel()

	for c := range equalValuesCases() {
		mock := new(testing.T)

		// Test NotEqualValues
		t.Run(fmt.Sprintf("NotEqualValues(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := NotEqualValues(mock, c.expected, c.actual)

			if res != c.notEqualResult {
				t.Errorf("NotEqualValues(%#v, %#v) should return %#v", c.expected, c.actual, c.notEqualResult)
			}
		})

		// Test EqualValues (inverse of NotEqualValues)
		t.Run(fmt.Sprintf("EqualValues(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			expectedEqualResult := !c.notEqualResult // EqualValues should return opposite of NotEqualValues
			res := EqualValues(mock, c.expected, c.actual)

			if res != expectedEqualResult {
				t.Errorf("EqualValues(%#v, %#v) should return %#v", c.expected, c.actual, expectedEqualResult)
			}
		})
	}
}

func TestEqualEmpty(t *testing.T) {
	t.Parallel()

	// TODO(fredbi): redundant test context declaration
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	var tiP *time.Time
	var tiNP time.Time
	var s *string
	var f *os.File
	sP := &s
	x := 1
	xP := &x

	type TString string
	type TStruct struct {
		x int
	}

	t.Run("should be empty", func(t *testing.T) {
		mock := new(testing.T)

		True(t, Empty(mock, ""), "Empty string is empty")
		True(t, Empty(mock, nil), "Nil is empty")
		True(t, Empty(mock, []string{}), "Empty string array is empty")
		True(t, Empty(mock, 0), "Zero int value is empty")
		True(t, Empty(mock, false), "False value is empty")
		True(t, Empty(mock, make(chan struct{})), "Channel without values is empty")
		True(t, Empty(mock, s), "Nil string pointer is empty")
		True(t, Empty(mock, f), "Nil os.File pointer is empty")
		True(t, Empty(mock, tiP), "Nil time.Time pointer is empty")
		True(t, Empty(mock, tiNP), "time.Time is empty")
		True(t, Empty(mock, TStruct{}), "struct with zero values is empty")
		True(t, Empty(mock, TString("")), "empty aliased string is empty")
		True(t, Empty(mock, sP), "ptr to nil value is empty")
		True(t, Empty(mock, [1]int{}), "array is state")
	})

	t.Run("should not be empty", func(t *testing.T) {
		mock := new(testing.T)

		False(t, Empty(mock, "something"), "Non Empty string is not empty")
		False(t, Empty(mock, errors.New("something")), "Non nil object is not empty")
		False(t, Empty(mock, []string{"something"}), "Non empty string array is not empty")
		False(t, Empty(mock, 1), "Non-zero int value is not empty")
		False(t, Empty(mock, true), "True value is not empty")
		False(t, Empty(mock, chWithValue), "Channel with values is not empty")
		False(t, Empty(mock, TStruct{x: 1}), "struct with initialized values is empty")
		False(t, Empty(mock, TString("abc")), "non-empty aliased string is empty")
		False(t, Empty(mock, xP), "ptr to non-nil value is not empty")
		False(t, Empty(mock, [1]int{42}), "array is not state")
	})

	// error messages validation
	for tt := range equalEmptyCases() {
		t.Run(tt.name, func(t *testing.T) {
			mock := new(captureT)

			res := Empty(mock, tt.value)
			mock.checkResultAndErrMsg(t, res, tt.expectedResult, tt.expectedErrMsg)
		})
	}
}

func TestEqualNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("should not be empty", func(t *testing.T) {
		mock := new(testing.T)

		False(t, NotEmpty(mock, ""), "Empty string is empty")
		False(t, NotEmpty(mock, nil), "Nil is empty")
		False(t, NotEmpty(mock, []string{}), "Empty string array is empty")
		False(t, NotEmpty(mock, 0), "Zero int value is empty")
		False(t, NotEmpty(mock, false), "False value is empty")
		False(t, NotEmpty(mock, make(chan struct{})), "Channel without values is empty")
		False(t, NotEmpty(mock, [1]int{}), "array is state")
	})

	t.Run("should  be empty", func(t *testing.T) {
		mock := new(testing.T)

		chWithValue := make(chan struct{}, 1)
		chWithValue <- struct{}{}

		False(t, NotEmpty(mock, ""), "Empty string is empty")
		True(t, NotEmpty(mock, "something"), "Non Empty string is not empty")
		True(t, NotEmpty(mock, errors.New("something")), "Non nil object is not empty")
		True(t, NotEmpty(mock, []string{"something"}), "Non empty string array is not empty")
		True(t, NotEmpty(mock, 1), "Non-zero int value is not empty")
		True(t, NotEmpty(mock, true), "True value is not empty")
		True(t, NotEmpty(mock, chWithValue), "Channel with values is not empty")
		True(t, NotEmpty(mock, [1]int{42}), "array is not state")
	})

	// error messages validation
	for tt := range equalNotEmptyCases() {
		t.Run(tt.name, func(t *testing.T) {
			mock := new(captureT)

			res := NotEmpty(mock, tt.value)
			mock.checkResultAndErrMsg(t, tt.expectedResult, res, tt.expectedErrMsg)
		})
	}
}

func TestEqualExactly(t *testing.T) {
	t.Parallel()

	for c := range equalExactlyCases() {
		t.Run(fmt.Sprintf("Exactly(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			t.Parallel()
			mock := new(testing.T)

			res := Exactly(mock, c.expected, c.actual)
			if res != c.result {
				t.Errorf("Exactly(%#v, %#v) should return %#v", c.expected, c.actual, c.result)
			}
		})
	}
}

func TestEqualBytes(t *testing.T) {
	t.Parallel()

	i := 0
	for c := range equalBytesCases() {
		Equal(t, reflect.DeepEqual(c.a, c.b), ObjectsAreEqual(c.a, c.b), "case %d failed", i)
		i++
	}
}

type equalCase struct {
	expected any
	actual   any
	result   bool
	remark   string
}

func equalCases() iter.Seq[equalCase] {
	type myType string
	var m map[string]any
	return slices.Values([]equalCase{
		{"Hello World", "Hello World", true, ""},
		{123, 123, true, ""},
		{123.5, 123.5, true, ""},
		{[]byte("Hello World"), []byte("Hello World"), true, ""},
		{nil, nil, true, ""},
		{int32(123), int32(123), true, ""},
		{uint64(123), uint64(123), true, ""},
		{myType("1"), myType("1"), true, ""},
		{&struct{}{}, &struct{}{}, true, "pointer equality is based on equality of underlying value"},

		// Not expected to be equal
		{m["bar"], "something", false, ""},
		{myType("1"), myType("2"), false, ""},

		// A case that might be confusing, especially with numeric literals
		{10, uint(10), false, ""},
	})
}

type samePointersCase struct {
	name string
	args args
	same BoolAssertionFunc
	ok   BoolAssertionFunc
}

type args struct {
	first  any
	second any
}

func ptr(i int) *int {
	return &i
}

func equalSamePointersCases() iter.Seq[samePointersCase] {
	p := ptr(2)
	return slices.Values([]samePointersCase{
		{
			name: "1 != 2",
			args: args{first: 1, second: 2},
			same: False,
			ok:   False,
		},
		{
			name: "1 != 1 (not same ptr)",
			args: args{first: 1, second: 1},
			same: False,
			ok:   False,
		},
		{
			name: "ptr(1) == ptr(1)",
			args: args{first: p, second: p},
			same: True,
			ok:   True,
		},
		{
			name: "int(1) != float32(1)",
			args: args{first: int(1), second: float32(1)},
			same: False,
			ok:   False,
		},
		{
			name: "array != slice",
			args: args{first: [2]int{1, 2}, second: []int{1, 2}},
			same: False,
			ok:   False,
		},
		{
			name: "non-pointer vs pointer (1 != ptr(2))",
			args: args{first: 1, second: p},
			same: False,
			ok:   False,
		},
		{
			name: "pointer vs non-pointer (ptr(2) != 1)",
			args: args{first: p, second: 1},
			same: False,
			ok:   False,
		},
	})
}

type equalNotEqualCase struct {
	expected any
	actual   any
	result   bool
}

func equalNotEqualCases() iter.Seq[equalNotEqualCase] {
	return slices.Values([]equalNotEqualCase{
		// cases that are expected not to match
		{"Hello World", "Hello World!", true},
		{123, 1234, true},
		{123.5, 123.55, true},
		{[]byte("Hello World"), []byte("Hello World!"), true},
		{nil, new(AssertionTesterConformingObject), true},

		// cases that are expected to match
		{nil, nil, false},
		{"Hello World", "Hello World", false},
		{123, 123, false},
		{123.5, 123.5, false},
		{[]byte("Hello World"), []byte("Hello World"), false},
		{new(AssertionTesterConformingObject), new(AssertionTesterConformingObject), false},
		{&struct{}{}, &struct{}{}, false},
		{func() int { return 23 }, func() int { return 24 }, false},
		// A case that might be confusing, especially with numeric literals
		{int(10), uint(10), true},
	})
}

type equalValuesCase struct {
	expected       any
	actual         any
	notEqualResult bool // result for NotEqualValues
}

func equalValuesCases() iter.Seq[equalValuesCase] {
	return slices.Values([]equalValuesCase{
		// cases that are expected not to match
		{"Hello World", "Hello World!", true},
		{123, 1234, true},
		{123.5, 123.55, true},
		{[]byte("Hello World"), []byte("Hello World!"), true},
		{nil, new(AssertionTesterConformingObject), true},

		// cases that are expected to match
		{nil, nil, false},
		{"Hello World", "Hello World", false},
		{123, 123, false},
		{123.5, 123.5, false},
		{[]byte("Hello World"), []byte("Hello World"), false},
		{new(AssertionTesterConformingObject), new(AssertionTesterConformingObject), false},
		{&struct{}{}, &struct{}{}, false},

		// Different behavior from NotEqual()
		{func() int { return 23 }, func() int { return 24 }, true},
		{int(10), int(11), true},
		{int(10), uint(10), false},

		{struct{}{}, struct{}{}, false},
	})
}

type equalEmptyCase struct {
	name           string
	value          any
	expectedResult bool
	expectedErrMsg string
}

func equalEmptyCases() iter.Seq[equalEmptyCase] {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	// var tiP *time.Time
	// var tiNP time.Time
	// var s *string
	// var f *os.File
	// sP := &s
	x := 1
	xP := &x

	type TString string
	type TStruct struct {
		x int
	}

	return slices.Values([]equalEmptyCase{
		{
			name:           "Non Empty string is not empty",
			value:          "something",
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was something\n",
		},
		{
			name:           "Non nil object is not empty",
			value:          errors.New("something"),
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was something\n",
		},
		{
			name:           "Non empty string array is not empty",
			value:          []string{"something"},
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was [something]\n",
		},
		{
			name:           "Non-zero int value is not empty",
			value:          1,
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was 1\n",
		},
		{
			name:           "True value is not empty",
			value:          true,
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was true\n",
		},
		{
			name:           "Channel with values is not empty",
			value:          chWithValue,
			expectedResult: false,
			expectedErrMsg: fmt.Sprintf("Should be empty, but was %v\n", chWithValue),
		},
		{
			name:           "struct with initialized values is empty",
			value:          TStruct{x: 1},
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was {1}\n",
		},
		{
			name:           "non-empty aliased string is empty",
			value:          TString("abc"),
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was abc\n",
		},
		{
			name:           "ptr to non-nil value is not empty",
			value:          xP,
			expectedResult: false,
			expectedErrMsg: fmt.Sprintf("Should be empty, but was %p\n", xP),
		},
		{
			name:           "array is not state",
			value:          [1]int{42},
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was [42]\n",
		},

		// Here are some edge cases
		{
			name:           "string with only spaces is not empty",
			value:          "   ",
			expectedResult: false,
			expectedErrMsg: "Should be empty, but was    \n", // TODO FIX THIS strange error message
		},
		{
			name:           "string with a line feed is not empty",
			value:          "\n",
			expectedResult: false,
			// TODO This is the exact same error message as for an empty string
			expectedErrMsg: "Should be empty, but was \n", // TODO FIX THIS strange error message
		},
		{
			name:           "string with only tabulation and lines feed is not empty",
			value:          "\n\t\n",
			expectedResult: false,
			// TODO The line feeds and tab are not helping to spot what is expected
			expectedErrMsg: "" + // this syntax is used to show how errors are reported.
				"Should be empty, but was \n" +
				"\t\n",
		},
		{
			name:           "string with trailing lines feed is not empty",
			value:          "foo\n\n",
			expectedResult: false,
			// TODO it's not clear if one or two lines feed are expected
			expectedErrMsg: "Should be empty, but was foo\n\n",
		},
		{
			name:           "string with leading and trailing tabulation and lines feed is not empty",
			value:          "\n\nfoo\t\n\t\n",
			expectedResult: false,
			// TODO The line feeds and tab are not helping to figure what is expected
			expectedErrMsg: "" +
				"Should be empty, but was \n" +
				"\n" +
				"foo\t\n" +
				"\t\n",
		},

		{
			name:           "non-printable character is not empty",
			value:          "\u00a0", // NO-BREAK SPACE UNICODE CHARACTER
			expectedResult: false,
			// TODO here you cannot figure out what is expected
			expectedErrMsg: "Should be empty, but was \u00a0\n",
		},

		// Here we are testing there is no error message on success
		{
			name:           "Empty string is empty",
			value:          "",
			expectedResult: true,
			expectedErrMsg: "",
		},
	})
}

type equalNotEmptyCase struct {
	name           string
	value          any
	expectedResult bool
	expectedErrMsg string
}

func equalNotEmptyCases() iter.Seq[equalNotEmptyCase] {
	return slices.Values([]equalNotEmptyCase{
		{
			name:           "Empty string is empty",
			value:          "",
			expectedResult: false,
			expectedErrMsg: `Should NOT be empty, but was ` + "\n", // TODO FIX THIS strange error message
		},
		{
			name:           "Nil is empty",
			value:          nil,
			expectedResult: false,
			expectedErrMsg: "Should NOT be empty, but was <nil>\n",
		},
		{
			name:           "Empty string array is empty",
			value:          []string{},
			expectedResult: false,
			expectedErrMsg: "Should NOT be empty, but was []\n",
		},
		{
			name:           "Zero int value is empty",
			value:          0,
			expectedResult: false,
			expectedErrMsg: "Should NOT be empty, but was 0\n",
		},
		{
			name:           "False value is empty",
			value:          false,
			expectedResult: false,
			expectedErrMsg: "Should NOT be empty, but was false\n",
		},
		{
			name:           "array is state",
			value:          [1]int{},
			expectedResult: false,
			expectedErrMsg: "Should NOT be empty, but was [0]\n",
		},

		// Here we are testing there is no error message on success
		{
			name:           "Non Empty string is not empty",
			value:          "something",
			expectedResult: true,
			expectedErrMsg: "",
		},
	})
}

type diffTestingStruct struct {
	A string
	B int
}

func (d *diffTestingStruct) String() string {
	return d.A
}

type equalExactlyCase struct {
	expected any
	actual   any
	result   bool
}

func equalExactlyCases() iter.Seq[equalExactlyCase] {
	a := float32(1)
	b := float64(1)
	c := float32(1)
	d := float32(2)

	return slices.Values([]equalExactlyCase{
		{a, b, false},
		{a, d, false},
		{a, c, true},
		{nil, a, false},
		{a, nil, false},
	})
}

type equalBytesCase struct {
	a, b []byte
}

func equalBytesCases() iter.Seq[equalBytesCase] {
	return slices.Values([]equalBytesCase{
		{make([]byte, 2), make([]byte, 2)},
		{make([]byte, 2), make([]byte, 2, 3)},
		{nil, make([]byte, 0)},
	})
}
