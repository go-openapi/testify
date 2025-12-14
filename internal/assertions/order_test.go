// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"fmt"
	"iter"
	"slices"
	"testing"
)

type orderedFixture struct {
	collection any
	msg        string
}

func TestOrderIsIncreasing(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !IsIncreasing(mock, []int{1, 2}) {
		t.Error("IsIncreasing should return true")
	}

	if !IsIncreasing(mock, []int{1, 2, 3, 4, 5}) {
		t.Error("IsIncreasing should return true")
	}

	if IsIncreasing(mock, []int{1, 1}) {
		t.Error("IsIncreasing should return false")
	}

	if IsIncreasing(mock, []int{2, 1}) {
		t.Error("IsIncreasing should return false")
	}

	// Check error report
	for currCase := range decreasingFixtures() {
		t.Run(fmt.Sprintf("%#v", currCase.collection), func(t *testing.T) {
			out := &outputT{buf: bytes.NewBuffer(nil)}
			False(t, IsIncreasing(out, currCase.collection))
			Contains(t, out.buf.String(), currCase.msg)
		})
	}
}

func TestOrderIsNonIncreasing(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !IsNonIncreasing(mock, []int{2, 1}) {
		t.Error("IsNonIncreasing should return true")
	}

	if !IsNonIncreasing(mock, []int{5, 4, 4, 3, 2, 1}) {
		t.Error("IsNonIncreasing should return true")
	}

	if !IsNonIncreasing(mock, []int{1, 1}) {
		t.Error("IsNonIncreasing should return true")
	}

	if IsNonIncreasing(mock, []int{1, 2}) {
		t.Error("IsNonIncreasing should return false")
	}

	// Check error report
	for currCase := range increasingFixtures() {
		t.Run(fmt.Sprintf("%#v", currCase.collection), func(t *testing.T) {
			out := &outputT{buf: bytes.NewBuffer(nil)}
			False(t, IsNonIncreasing(out, currCase.collection))
			Contains(t, out.buf.String(), currCase.msg)
		})
	}
}

func TestOrderIsDecreasing(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !IsDecreasing(mock, []int{2, 1}) {
		t.Error("IsDecreasing should return true")
	}

	if !IsDecreasing(mock, []int{5, 4, 3, 2, 1}) {
		t.Error("IsDecreasing should return true")
	}

	if IsDecreasing(mock, []int{1, 1}) {
		t.Error("IsDecreasing should return false")
	}

	if IsDecreasing(mock, []int{1, 2}) {
		t.Error("IsDecreasing should return false")
	}

	// Check error report
	for currCase := range increasingFixtures2() {
		t.Run(fmt.Sprintf("%#v", currCase.collection), func(t *testing.T) {
			out := &outputT{buf: bytes.NewBuffer(nil)}
			False(t, IsDecreasing(out, currCase.collection))
			Contains(t, out.buf.String(), currCase.msg)
		})
	}
}

func TestOrderIsNonDecreasing(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !IsNonDecreasing(mock, []int{1, 2}) {
		t.Error("IsNonDecreasing should return true")
	}

	if !IsNonDecreasing(mock, []int{1, 1, 2, 3, 4, 5}) {
		t.Error("IsNonDecreasing should return true")
	}

	if !IsNonDecreasing(mock, []int{1, 1}) {
		t.Error("IsNonDecreasing should return false")
	}

	if IsNonDecreasing(mock, []int{2, 1}) {
		t.Error("IsNonDecreasing should return false")
	}

	// Check error report
	for currCase := range decreasingFixtures2() {
		t.Run(fmt.Sprintf("%#v", currCase.collection), func(t *testing.T) {
			out := &outputT{buf: bytes.NewBuffer(nil)}
			False(t, IsNonDecreasing(out, currCase.collection))
			Contains(t, out.buf.String(), currCase.msg)
		})
	}
}

func TestOrderMsgAndArgsForwarding(t *testing.T) {
	t.Parallel()

	msgAndArgs := []any{"format %s %x", "this", 0xc001}
	expectedOutput := "format this c001\n"
	collection := []int{1, 2, 1}
	funcs := []func(t T){
		func(t T) { IsIncreasing(t, collection, msgAndArgs...) },
		func(t T) { IsNonIncreasing(t, collection, msgAndArgs...) },
		func(t T) { IsDecreasing(t, collection, msgAndArgs...) },
		func(t T) { IsNonDecreasing(t, collection, msgAndArgs...) },
	}
	for _, f := range funcs {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		f(out)
		Contains(t, out.buf.String(), expectedOutput)
	}
}

func decreasingFixtures() iter.Seq[orderedFixture] {
	return slices.Values(
		[]orderedFixture{
			{collection: []string{"b", "a"}, msg: `"b" is not less than "a"`},
			{collection: []int{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []int{2, 1, 3, 4, 5, 6, 7}, msg: `"2" is not less than "1"`},
			{collection: []int{-1, 0, 2, 1}, msg: `"2" is not less than "1"`},
			{collection: []int8{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []int16{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []int32{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []int64{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []uint8{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []uint16{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []uint32{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []uint64{2, 1}, msg: `"2" is not less than "1"`},
			{collection: []float32{2.34, 1.23}, msg: `"2.34" is not less than "1.23"`},
			{collection: []float64{2.34, 1.23}, msg: `"2.34" is not less than "1.23"`},
			{collection: struct{}{}, msg: `object struct {} is not an ordered collection`},
		},
	)
}

func increasingFixtures() iter.Seq[orderedFixture] {
	return slices.Values(
		[]orderedFixture{
			{collection: []string{"a", "b"}, msg: `"a" is not greater than or equal to "b"`},
			{collection: []int{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []int{1, 2, 7, 6, 5, 4, 3}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []int{5, 4, 3, 1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []int8{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []int16{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []int32{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []int64{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []uint8{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []uint16{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []uint32{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []uint64{1, 2}, msg: `"1" is not greater than or equal to "2"`},
			{collection: []float32{1.23, 2.34}, msg: `"1.23" is not greater than or equal to "2.34"`},
			{collection: []float64{1.23, 2.34}, msg: `"1.23" is not greater than or equal to "2.34"`},
			{collection: struct{}{}, msg: `object struct {} is not an ordered collection`},
		},
	)
}

func increasingFixtures2() iter.Seq[orderedFixture] {
	return slices.Values(
		[]orderedFixture{
			{collection: []string{"a", "b"}, msg: `"a" is not greater than "b"`},
			{collection: []int{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []int{1, 2, 7, 6, 5, 4, 3}, msg: `"1" is not greater than "2"`},
			{collection: []int{5, 4, 3, 1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []int8{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []int16{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []int32{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []int64{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []uint8{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []uint16{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []uint32{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []uint64{1, 2}, msg: `"1" is not greater than "2"`},
			{collection: []float32{1.23, 2.34}, msg: `"1.23" is not greater than "2.34"`},
			{collection: []float64{1.23, 2.34}, msg: `"1.23" is not greater than "2.34"`},
			{collection: struct{}{}, msg: `object struct {} is not an ordered collection`},
		},
	)
}

func decreasingFixtures2() iter.Seq[orderedFixture] {
	return slices.Values(
		[]orderedFixture{
			{collection: []string{"b", "a"}, msg: `"b" is not less than or equal to "a"`},
			{collection: []int{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []int{2, 1, 3, 4, 5, 6, 7}, msg: `"2" is not less than or equal to "1"`},
			{collection: []int{-1, 0, 2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []int8{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []int16{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []int32{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []int64{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []uint8{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []uint16{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []uint32{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []uint64{2, 1}, msg: `"2" is not less than or equal to "1"`},
			{collection: []float32{2.34, 1.23}, msg: `"2.34" is not less than or equal to "1.23"`},
			{collection: []float64{2.34, 1.23}, msg: `"2.34" is not less than or equal to "1.23"`},
			{collection: struct{}{}, msg: `object struct {} is not an ordered collection`},
		},
	)
}
