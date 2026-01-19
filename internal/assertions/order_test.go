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
			t.Parallel()

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

// Tests for generic ordering functions

type orderTestCase struct {
	name       string
	collection any
	shouldPass bool
}

func orderIncreasingCases() iter.Seq[orderTestCase] {
	return slices.Values([]orderTestCase{
		// Success cases - strictly increasing
		{name: "int/increasing", collection: []int{1, 2, 3}, shouldPass: true},
		{name: "int8/increasing", collection: []int8{1, 2, 3}, shouldPass: true},
		{name: "int16/increasing", collection: []int16{1, 2, 3}, shouldPass: true},
		{name: "int32/increasing", collection: []int32{1, 2, 3}, shouldPass: true},
		{name: "int64/increasing", collection: []int64{1, 2, 3}, shouldPass: true},
		{name: "uint/increasing", collection: []uint{1, 2, 3}, shouldPass: true},
		{name: "uint8/increasing", collection: []uint8{1, 2, 3}, shouldPass: true},
		{name: "uint16/increasing", collection: []uint16{1, 2, 3}, shouldPass: true},
		{name: "uint32/increasing", collection: []uint32{1, 2, 3}, shouldPass: true},
		{name: "uint64/increasing", collection: []uint64{1, 2, 3}, shouldPass: true},
		{name: "float32/increasing", collection: []float32{1.1, 2.2, 3.3}, shouldPass: true},
		{name: "float64/increasing", collection: []float64{1.1, 2.2, 3.3}, shouldPass: true},
		{name: "string/increasing", collection: []string{"a", "b", "c"}, shouldPass: true},

		// Failure cases - not strictly increasing (equal or decreasing)
		{name: "int/equal", collection: []int{1, 1, 2}, shouldPass: false},
		{name: "int/decreasing", collection: []int{3, 2, 1}, shouldPass: false},
		{name: "float64/equal", collection: []float64{1.1, 1.1, 2.2}, shouldPass: false},
		{name: "string/equal", collection: []string{"a", "a", "b"}, shouldPass: false},
	})
}

func orderDecreasingCases() iter.Seq[orderTestCase] {
	return slices.Values([]orderTestCase{
		// Success cases - strictly decreasing
		{name: "int/decreasing", collection: []int{3, 2, 1}, shouldPass: true},
		{name: "int8/decreasing", collection: []int8{3, 2, 1}, shouldPass: true},
		{name: "int16/decreasing", collection: []int16{3, 2, 1}, shouldPass: true},
		{name: "int32/decreasing", collection: []int32{3, 2, 1}, shouldPass: true},
		{name: "int64/decreasing", collection: []int64{3, 2, 1}, shouldPass: true},
		{name: "uint/decreasing", collection: []uint{3, 2, 1}, shouldPass: true},
		{name: "uint8/decreasing", collection: []uint8{3, 2, 1}, shouldPass: true},
		{name: "uint16/decreasing", collection: []uint16{3, 2, 1}, shouldPass: true},
		{name: "uint32/decreasing", collection: []uint32{3, 2, 1}, shouldPass: true},
		{name: "uint64/decreasing", collection: []uint64{3, 2, 1}, shouldPass: true},
		{name: "float32/decreasing", collection: []float32{3.3, 2.2, 1.1}, shouldPass: true},
		{name: "float64/decreasing", collection: []float64{3.3, 2.2, 1.1}, shouldPass: true},
		{name: "string/decreasing", collection: []string{"c", "b", "a"}, shouldPass: true},

		// Failure cases - not strictly decreasing (equal or increasing)
		{name: "int/equal", collection: []int{2, 1, 1}, shouldPass: false},
		{name: "int/increasing", collection: []int{1, 2, 3}, shouldPass: false},
		{name: "float64/equal", collection: []float64{2.2, 1.1, 1.1}, shouldPass: false},
		{name: "string/equal", collection: []string{"b", "a", "a"}, shouldPass: false},
	})
}

func orderNonIncreasingCases() iter.Seq[orderTestCase] {
	return slices.Values([]orderTestCase{
		// Success cases - decreasing or equal (not increasing)
		{name: "int/decreasing", collection: []int{3, 2, 1}, shouldPass: true},
		{name: "int/with-equal", collection: []int{3, 2, 2, 1}, shouldPass: true},
		{name: "int/all-equal", collection: []int{2, 2, 2}, shouldPass: true},
		{name: "int8/decreasing", collection: []int8{3, 2, 1}, shouldPass: true},
		{name: "int16/decreasing", collection: []int16{3, 2, 1}, shouldPass: true},
		{name: "int32/decreasing", collection: []int32{3, 2, 1}, shouldPass: true},
		{name: "int64/decreasing", collection: []int64{3, 2, 1}, shouldPass: true},
		{name: "uint/decreasing", collection: []uint{3, 2, 1}, shouldPass: true},
		{name: "uint8/decreasing", collection: []uint8{3, 2, 1}, shouldPass: true},
		{name: "uint16/decreasing", collection: []uint16{3, 2, 1}, shouldPass: true},
		{name: "uint32/decreasing", collection: []uint32{3, 2, 1}, shouldPass: true},
		{name: "uint64/decreasing", collection: []uint64{3, 2, 1}, shouldPass: true},
		{name: "float32/decreasing", collection: []float32{3.3, 2.2, 1.1}, shouldPass: true},
		{name: "float64/decreasing", collection: []float64{3.3, 2.2, 1.1}, shouldPass: true},
		{name: "float64/with-equal", collection: []float64{3.3, 2.2, 2.2}, shouldPass: true},
		{name: "string/decreasing", collection: []string{"c", "b", "a"}, shouldPass: true},
		{name: "string/with-equal", collection: []string{"c", "b", "b"}, shouldPass: true},

		// Failure cases - increasing
		{name: "int/increasing", collection: []int{1, 2, 3}, shouldPass: false},
		{name: "float64/increasing", collection: []float64{1.1, 2.2, 3.3}, shouldPass: false},
		{name: "string/increasing", collection: []string{"a", "b", "c"}, shouldPass: false},
	})
}

func orderNonDecreasingCases() iter.Seq[orderTestCase] {
	return slices.Values([]orderTestCase{
		// Success cases - increasing or equal (not decreasing)
		{name: "int/increasing", collection: []int{1, 2, 3}, shouldPass: true},
		{name: "int/with-equal", collection: []int{1, 2, 2, 3}, shouldPass: true},
		{name: "int/all-equal", collection: []int{2, 2, 2}, shouldPass: true},
		{name: "int8/increasing", collection: []int8{1, 2, 3}, shouldPass: true},
		{name: "int16/increasing", collection: []int16{1, 2, 3}, shouldPass: true},
		{name: "int32/increasing", collection: []int32{1, 2, 3}, shouldPass: true},
		{name: "int64/increasing", collection: []int64{1, 2, 3}, shouldPass: true},
		{name: "uint/increasing", collection: []uint{1, 2, 3}, shouldPass: true},
		{name: "uint8/increasing", collection: []uint8{1, 2, 3}, shouldPass: true},
		{name: "uint16/increasing", collection: []uint16{1, 2, 3}, shouldPass: true},
		{name: "uint32/increasing", collection: []uint32{1, 2, 3}, shouldPass: true},
		{name: "uint64/increasing", collection: []uint64{1, 2, 3}, shouldPass: true},
		{name: "float32/increasing", collection: []float32{1.1, 2.2, 3.3}, shouldPass: true},
		{name: "float64/increasing", collection: []float64{1.1, 2.2, 3.3}, shouldPass: true},
		{name: "float64/with-equal", collection: []float64{1.1, 2.2, 2.2}, shouldPass: true},
		{name: "string/increasing", collection: []string{"a", "b", "c"}, shouldPass: true},
		{name: "string/with-equal", collection: []string{"a", "b", "b"}, shouldPass: true},

		// Failure cases - decreasing
		{name: "int/decreasing", collection: []int{3, 2, 1}, shouldPass: false},
		{name: "float64/decreasing", collection: []float64{3.3, 2.2, 1.1}, shouldPass: false},
		{name: "string/decreasing", collection: []string{"c", "b", "a"}, shouldPass: false},
	})
}

func TestOrderIsIncreasingT(t *testing.T) {
	t.Parallel()

	for tc := range orderIncreasingCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Dispatch based on type
			switch coll := tc.collection.(type) {
			case []int:
				testOrderingT(IsIncreasingT[[]int, int], coll, tc.shouldPass)(t)
			case []int8:
				testOrderingT(IsIncreasingT[[]int8, int8], coll, tc.shouldPass)(t)
			case []int16:
				testOrderingT(IsIncreasingT[[]int16, int16], coll, tc.shouldPass)(t)
			case []int32:
				testOrderingT(IsIncreasingT[[]int32, int32], coll, tc.shouldPass)(t)
			case []int64:
				testOrderingT(IsIncreasingT[[]int64, int64], coll, tc.shouldPass)(t)
			case []uint:
				testOrderingT(IsIncreasingT[[]uint, uint], coll, tc.shouldPass)(t)
			case []uint8:
				testOrderingT(IsIncreasingT[[]uint8, uint8], coll, tc.shouldPass)(t)
			case []uint16:
				testOrderingT(IsIncreasingT[[]uint16, uint16], coll, tc.shouldPass)(t)
			case []uint32:
				testOrderingT(IsIncreasingT[[]uint32, uint32], coll, tc.shouldPass)(t)
			case []uint64:
				testOrderingT(IsIncreasingT[[]uint64, uint64], coll, tc.shouldPass)(t)
			case []float32:
				testOrderingT(IsIncreasingT[[]float32, float32], coll, tc.shouldPass)(t)
			case []float64:
				testOrderingT(IsIncreasingT[[]float64, float64], coll, tc.shouldPass)(t)
			case []string:
				testOrderingT(IsIncreasingT[[]string, string], coll, tc.shouldPass)(t)
			default:
				t.Fatalf("unexpected type: %T", coll)
			}
		})
	}
}

func TestOrderIsDecreasingT(t *testing.T) {
	t.Parallel()

	for tc := range orderDecreasingCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Dispatch based on type
			switch coll := tc.collection.(type) {
			case []int:
				testOrderingT(IsDecreasingT[[]int, int], coll, tc.shouldPass)(t)
			case []int8:
				testOrderingT(IsDecreasingT[[]int8, int8], coll, tc.shouldPass)(t)
			case []int16:
				testOrderingT(IsDecreasingT[[]int16, int16], coll, tc.shouldPass)(t)
			case []int32:
				testOrderingT(IsDecreasingT[[]int32, int32], coll, tc.shouldPass)(t)
			case []int64:
				testOrderingT(IsDecreasingT[[]int64, int64], coll, tc.shouldPass)(t)
			case []uint:
				testOrderingT(IsDecreasingT[[]uint, uint], coll, tc.shouldPass)(t)
			case []uint8:
				testOrderingT(IsDecreasingT[[]uint8, uint8], coll, tc.shouldPass)(t)
			case []uint16:
				testOrderingT(IsDecreasingT[[]uint16, uint16], coll, tc.shouldPass)(t)
			case []uint32:
				testOrderingT(IsDecreasingT[[]uint32, uint32], coll, tc.shouldPass)(t)
			case []uint64:
				testOrderingT(IsDecreasingT[[]uint64, uint64], coll, tc.shouldPass)(t)
			case []float32:
				testOrderingT(IsDecreasingT[[]float32, float32], coll, tc.shouldPass)(t)
			case []float64:
				testOrderingT(IsDecreasingT[[]float64, float64], coll, tc.shouldPass)(t)
			case []string:
				testOrderingT(IsDecreasingT[[]string, string], coll, tc.shouldPass)(t)
			default:
				t.Fatalf("unexpected type: %T", coll)
			}
		})
	}
}

func TestOrderIsNonIncreasingT(t *testing.T) {
	t.Parallel()

	for tc := range orderNonIncreasingCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Dispatch based on type
			switch coll := tc.collection.(type) {
			case []int:
				testOrderingT(IsNonIncreasingT[[]int, int], coll, tc.shouldPass)(t)
			case []int8:
				testOrderingT(IsNonIncreasingT[[]int8, int8], coll, tc.shouldPass)(t)
			case []int16:
				testOrderingT(IsNonIncreasingT[[]int16, int16], coll, tc.shouldPass)(t)
			case []int32:
				testOrderingT(IsNonIncreasingT[[]int32, int32], coll, tc.shouldPass)(t)
			case []int64:
				testOrderingT(IsNonIncreasingT[[]int64, int64], coll, tc.shouldPass)(t)
			case []uint:
				testOrderingT(IsNonIncreasingT[[]uint, uint], coll, tc.shouldPass)(t)
			case []uint8:
				testOrderingT(IsNonIncreasingT[[]uint8, uint8], coll, tc.shouldPass)(t)
			case []uint16:
				testOrderingT(IsNonIncreasingT[[]uint16, uint16], coll, tc.shouldPass)(t)
			case []uint32:
				testOrderingT(IsNonIncreasingT[[]uint32, uint32], coll, tc.shouldPass)(t)
			case []uint64:
				testOrderingT(IsNonIncreasingT[[]uint64, uint64], coll, tc.shouldPass)(t)
			case []float32:
				testOrderingT(IsNonIncreasingT[[]float32, float32], coll, tc.shouldPass)(t)
			case []float64:
				testOrderingT(IsNonIncreasingT[[]float64, float64], coll, tc.shouldPass)(t)
			case []string:
				testOrderingT(IsNonIncreasingT[[]string, string], coll, tc.shouldPass)(t)
			default:
				t.Fatalf("unexpected type: %T", coll)
			}
		})
	}
}

func TestOrderIsNonDecreasingT(t *testing.T) {
	t.Parallel()

	for tc := range orderNonDecreasingCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Dispatch based on type
			switch coll := tc.collection.(type) {
			case []int:
				testOrderingT(IsNonDecreasingT[[]int, int], coll, tc.shouldPass)(t)
			case []int8:
				testOrderingT(IsNonDecreasingT[[]int8, int8], coll, tc.shouldPass)(t)
			case []int16:
				testOrderingT(IsNonDecreasingT[[]int16, int16], coll, tc.shouldPass)(t)
			case []int32:
				testOrderingT(IsNonDecreasingT[[]int32, int32], coll, tc.shouldPass)(t)
			case []int64:
				testOrderingT(IsNonDecreasingT[[]int64, int64], coll, tc.shouldPass)(t)
			case []uint:
				testOrderingT(IsNonDecreasingT[[]uint, uint], coll, tc.shouldPass)(t)
			case []uint8:
				testOrderingT(IsNonDecreasingT[[]uint8, uint8], coll, tc.shouldPass)(t)
			case []uint16:
				testOrderingT(IsNonDecreasingT[[]uint16, uint16], coll, tc.shouldPass)(t)
			case []uint32:
				testOrderingT(IsNonDecreasingT[[]uint32, uint32], coll, tc.shouldPass)(t)
			case []uint64:
				testOrderingT(IsNonDecreasingT[[]uint64, uint64], coll, tc.shouldPass)(t)
			case []float32:
				testOrderingT(IsNonDecreasingT[[]float32, float32], coll, tc.shouldPass)(t)
			case []float64:
				testOrderingT(IsNonDecreasingT[[]float64, float64], coll, tc.shouldPass)(t)
			case []string:
				testOrderingT(IsNonDecreasingT[[]string, string], coll, tc.shouldPass)(t)
			default:
				t.Fatalf("unexpected type: %T", coll)
			}
		})
	}
}

func sortedTCases() iter.Seq[orderTestCase] {
	return slices.Values([]orderTestCase{
		// Success cases - sorted (non-strictly increasing, allows equal values)
		{name: "int/sorted-increasing", collection: []int{1, 2, 3}, shouldPass: true},
		{name: "int/sorted-with-equal", collection: []int{1, 1, 2, 3}, shouldPass: true},
		{name: "int/sorted-all-equal", collection: []int{2, 2, 2}, shouldPass: true},
		{name: "int8/sorted", collection: []int8{1, 2, 3}, shouldPass: true},
		{name: "int16/sorted", collection: []int16{1, 2, 3}, shouldPass: true},
		{name: "int32/sorted", collection: []int32{1, 2, 3}, shouldPass: true},
		{name: "int64/sorted", collection: []int64{1, 2, 3}, shouldPass: true},
		{name: "uint/sorted", collection: []uint{1, 2, 3}, shouldPass: true},
		{name: "uint8/sorted", collection: []uint8{1, 2, 3}, shouldPass: true},
		{name: "uint16/sorted", collection: []uint16{1, 2, 3}, shouldPass: true},
		{name: "uint32/sorted", collection: []uint32{1, 2, 3}, shouldPass: true},
		{name: "uint64/sorted", collection: []uint64{1, 2, 3}, shouldPass: true},
		{name: "float32/sorted", collection: []float32{1.1, 2.2, 3.3}, shouldPass: true},
		{name: "float64/sorted", collection: []float64{1.1, 2.2, 3.3}, shouldPass: true},
		{name: "string/sorted", collection: []string{"a", "b", "c"}, shouldPass: true},
		{name: "string/sorted-with-equal", collection: []string{"a", "a", "b"}, shouldPass: true},

		// Failure cases - not sorted
		{name: "int/unsorted", collection: []int{1, 4, 2}, shouldPass: false},
		{name: "int/decreasing", collection: []int{3, 2, 1}, shouldPass: false},
		{name: "float64/unsorted", collection: []float64{1.1, 3.3, 2.2}, shouldPass: false},
		{name: "string/unsorted", collection: []string{"b", "a", "c"}, shouldPass: false},
	})
}

func TestSortedT(t *testing.T) {
	t.Parallel()

	for tc := range sortedTCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Dispatch based on type
			switch coll := tc.collection.(type) {
			case []int:
				testOrderingT(SortedT[[]int, int], coll, tc.shouldPass)(t)
			case []int8:
				testOrderingT(SortedT[[]int8, int8], coll, tc.shouldPass)(t)
			case []int16:
				testOrderingT(SortedT[[]int16, int16], coll, tc.shouldPass)(t)
			case []int32:
				testOrderingT(SortedT[[]int32, int32], coll, tc.shouldPass)(t)
			case []int64:
				testOrderingT(SortedT[[]int64, int64], coll, tc.shouldPass)(t)
			case []uint:
				testOrderingT(SortedT[[]uint, uint], coll, tc.shouldPass)(t)
			case []uint8:
				testOrderingT(SortedT[[]uint8, uint8], coll, tc.shouldPass)(t)
			case []uint16:
				testOrderingT(SortedT[[]uint16, uint16], coll, tc.shouldPass)(t)
			case []uint32:
				testOrderingT(SortedT[[]uint32, uint32], coll, tc.shouldPass)(t)
			case []uint64:
				testOrderingT(SortedT[[]uint64, uint64], coll, tc.shouldPass)(t)
			case []float32:
				testOrderingT(SortedT[[]float32, float32], coll, tc.shouldPass)(t)
			case []float64:
				testOrderingT(SortedT[[]float64, float64], coll, tc.shouldPass)(t)
			case []string:
				testOrderingT(SortedT[[]string, string], coll, tc.shouldPass)(t)
			default:
				t.Fatalf("unexpected type: %T", coll)
			}
		})
	}
}

func TestNotSortedT(t *testing.T) {
	t.Parallel()

	for tc := range sortedTCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Invert shouldPass for NotSorted
			shouldPass := !tc.shouldPass

			// Dispatch based on type
			switch coll := tc.collection.(type) {
			case []int:
				testOrderingT(NotSortedT[[]int, int], coll, shouldPass)(t)
			case []int8:
				testOrderingT(NotSortedT[[]int8, int8], coll, shouldPass)(t)
			case []int16:
				testOrderingT(NotSortedT[[]int16, int16], coll, shouldPass)(t)
			case []int32:
				testOrderingT(NotSortedT[[]int32, int32], coll, shouldPass)(t)
			case []int64:
				testOrderingT(NotSortedT[[]int64, int64], coll, shouldPass)(t)
			case []uint:
				testOrderingT(NotSortedT[[]uint, uint], coll, shouldPass)(t)
			case []uint8:
				testOrderingT(NotSortedT[[]uint8, uint8], coll, shouldPass)(t)
			case []uint16:
				testOrderingT(NotSortedT[[]uint16, uint16], coll, shouldPass)(t)
			case []uint32:
				testOrderingT(NotSortedT[[]uint32, uint32], coll, shouldPass)(t)
			case []uint64:
				testOrderingT(NotSortedT[[]uint64, uint64], coll, shouldPass)(t)
			case []float32:
				testOrderingT(NotSortedT[[]float32, float32], coll, shouldPass)(t)
			case []float64:
				testOrderingT(NotSortedT[[]float64, float64], coll, shouldPass)(t)
			case []string:
				testOrderingT(NotSortedT[[]string, string], coll, shouldPass)(t)
			default:
				t.Fatalf("unexpected type: %T", coll)
			}
		})
	}
}

//nolint:thelper // linter false positive: these are not helpers
func testOrderingT[OrderedSlice ~[]E, E Ordered](
	fn func(T, OrderedSlice, ...any) bool,
	collection OrderedSlice,
	shouldPass bool,
) func(*testing.T) {
	return func(t *testing.T) {
		mock := new(mockT)
		result := fn(mock, collection)

		if shouldPass {
			True(t, result)
			False(t, mock.Failed())
			return
		}

		False(t, result)
		True(t, mock.Failed())
	}
}
