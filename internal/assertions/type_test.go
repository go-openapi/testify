// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"testing"
)

func TestTypeImplements(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Implements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject)) {
		t.Error("Implements method should return true: AssertionTesterConformingObject implements AssertionTesterInterface")
	}
	if Implements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject)) {
		t.Error("Implements method should return false: AssertionTesterNonConformingObject does not implements AssertionTesterInterface")
	}
	if Implements(mock, (*AssertionTesterInterface)(nil), nil) {
		t.Error("Implements method should return false: nil does not implement AssertionTesterInterface")
	}
}

func TestTypeNotImplements(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !NotImplements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject)) {
		t.Error("NotImplements method should return true: AssertionTesterNonConformingObject does not implement AssertionTesterInterface")
	}
	if NotImplements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject)) {
		t.Error("NotImplements method should return false: AssertionTesterConformingObject implements AssertionTesterInterface")
	}
	if NotImplements(mock, (*AssertionTesterInterface)(nil), nil) {
		t.Error("NotImplements method should return false: nil can't be checked to be implementing AssertionTesterInterface or not")
	}
}

func TestTypeIsType(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !IsType(mock, new(AssertionTesterConformingObject), new(AssertionTesterConformingObject)) {
		t.Error("IsType should return true: AssertionTesterConformingObject is the same type as AssertionTesterConformingObject")
	}
	if IsType(mock, new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject)) {
		t.Error("IsType should return false: AssertionTesterConformingObject is not the same type as AssertionTesterNonConformingObject")
	}
}

func TestTypeNotIsType(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !IsNotType(mock, new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject)) {
		t.Error("NotIsType should return true: AssertionTesterConformingObject is not the same type as AssertionTesterNonConformingObject")
	}
	if IsNotType(mock, new(AssertionTesterConformingObject), new(AssertionTesterConformingObject)) {
		t.Error("NotIsType should return false: AssertionTesterConformingObject is the same type as AssertionTesterConformingObject")
	}
}

func TestTypeZeroWithSliceTooLongToPrint(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	longSlice := make([]int, 1_000_000)
	Zero(mock, longSlice)
	Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Should be zero, but was [0 0 0`)
	Contains(t, mock.errorString(), `<... truncated>`)
}

func TestTypeZero(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	for test := range typeZeros() {
		True(t, Zero(mock, test, "%#v is not the %T zero value", test, test))
	}

	for test := range typeNonZeros() {
		False(t, Zero(mock, test, "%#v is not the %T zero value", test, test))
	}
}

func TestTypeNotZero(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	for test := range typeZeros() {
		False(t, NotZero(mock, test, "%#v is not the %T zero value", test, test))
	}

	for test := range typeNonZeros() {
		True(t, NotZero(mock, test, "%#v is not the %T zero value", test, test))
	}
}

func TestTypeDiffEmptyCases(t *testing.T) {
	t.Parallel()

	Equal(t, "", diff(nil, nil))
	Equal(t, "", diff(struct{ foo string }{}, nil))
	Equal(t, "", diff(nil, struct{ foo string }{}))
	Equal(t, "", diff(1, 2))
	Equal(t, "", diff(1, 2))
	Equal(t, "", diff([]int{1}, []bool{true}))
}

// Ensure there are no data races.
func TestTypeDiffRace(t *testing.T) {
	t.Parallel()

	expected := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	actual := map[string]string{
		"d": "D",
		"e": "E",
		"f": "F",
	}

	// run diffs in parallel simulating tests with t.Parallel()
	numRoutines := 10
	rChans := make([]chan string, numRoutines)
	for idx := range rChans {
		rChans[idx] = make(chan string)
		go func(ch chan string) {
			defer close(ch)
			ch <- diff(expected, actual)
		}(rChans[idx])
	}

	for _, ch := range rChans {
		for msg := range ch {
			NotZero(t, msg) // dummy assert
		}
	}
}

func typeZeros() iter.Seq[any] {
	return slices.Values([]any{
		false,
		byte(0),
		complex64(0),
		complex128(0),
		float32(0),
		float64(0),
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		rune(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		uintptr(0),
		"",
		[0]any{},
		[]any(nil),
		struct{ x int }{},
		(*any)(nil),
		(func())(nil),
		nil,
		any(nil),
		map[any]any(nil),
		(chan any)(nil),
		(<-chan any)(nil),
		(chan<- any)(nil),
	})
}

func typeNonZeros() iter.Seq[any] {
	var i int

	return slices.Values([]any{
		true,
		byte(1),
		complex64(1),
		complex128(1),
		float32(1),
		float64(1),
		int(1),
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		rune(1),
		uint(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),
		uintptr(1),
		"s",
		[1]any{1},
		[]any{},
		struct{ x int }{1},
		(&i),
		(func() {}),
		any(1),
		map[any]any{},
		(make(chan any)),
		(<-chan any)(make(chan any)),
		(chan<- any)(make(chan any)),
	})
}
