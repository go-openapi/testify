// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"math"
	"slices"
	"testing"
	"time"
)

func TestNumberInDelta(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	True(t, InDelta(mock, 1.001, 1, 0.01), "|1.001 - 1| <= 0.01")
	True(t, InDelta(mock, 1, 1.001, 0.01), "|1 - 1.001| <= 0.01")
	True(t, InDelta(mock, 1, 2, 1), "|1 - 2| <= 1")
	False(t, InDelta(mock, 1, 2, 0.5), "Expected |1 - 2| <= 0.5 to fail")
	False(t, InDelta(mock, 2, 1, 0.5), "Expected |2 - 1| <= 0.5 to fail")
	False(t, InDelta(mock, "", nil, 1), "Expected non numerals to fail")
	False(t, InDelta(mock, 42, math.NaN(), 0.01), "Expected NaN for actual to fail")
	False(t, InDelta(mock, math.NaN(), 42, 0.01), "Expected NaN for expected to fail")
	True(t, InDelta(mock, math.NaN(), math.NaN(), 0.01), "Expected NaN for both to pass")

	for tc := range numberInDeltaCases() {
		True(t, InDelta(mock, tc.a, tc.b, tc.delta), "Expected |%V - %V| <= %v", tc.a, tc.b, tc.delta)
	}
}

func TestNumberInDeltaSlice(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	True(t, InDeltaSlice(mock,
		[]float64{1.001, math.NaN(), 0.999},
		[]float64{1, math.NaN(), 1},
		0.1), "{1.001, NaN, 0.009} is element-wise close to {1, NaN, 1} in delta=0.1")

	True(t, InDeltaSlice(mock,
		[]float64{1, math.NaN(), 2},
		[]float64{0, math.NaN(), 3},
		1), "{1, NaN, 2} is element-wise close to {0, NaN, 3} in delta=1")

	False(t, InDeltaSlice(mock,
		[]float64{1, math.NaN(), 2},
		[]float64{0, math.NaN(), 3},
		0.1), "{1, NaN, 2} is not element-wise close to {0, NaN, 3} in delta=0.1")

	False(t, InDeltaSlice(mock, "", nil, 1), "Expected non numeral slices to fail")
}

func TestNumberInDeltaMapValues(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	for tc := range numberInDeltaMapCases() {
		tc.f(t, InDeltaMapValues(mock, tc.expect, tc.actual, tc.delta), tc.title+"\n"+diff(tc.expect, tc.actual))
	}
}

func TestNumberInEpsilon(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	for tc := range numberInEpsilonTrueCases() {
		True(t, InEpsilon(t, tc.a, tc.b, tc.epsilon, "Expected %V and %V to have a relative difference of %v", tc.a, tc.b, tc.epsilon), "test: %q", tc)
	}

	for tc := range numberInEpsilonFalseCases() {
		False(t, InEpsilon(mock, tc.a, tc.b, tc.epsilon, "Expected %V and %V to have a relative difference of %v", tc.a, tc.b, tc.epsilon))
	}
}

func TestNumberInEpsilonSlice(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	True(t, InEpsilonSlice(mock,
		[]float64{2.2, math.NaN(), 2.0},
		[]float64{2.1, math.NaN(), 2.1},
		0.06), "{2.2, NaN, 2.0} is element-wise close to {2.1, NaN, 2.1} in epsilon=0.06")

	False(t, InEpsilonSlice(mock,
		[]float64{2.2, 2.0},
		[]float64{2.1, 2.1},
		0.04), "{2.2, 2.0} is not element-wise close to {2.1, 2.1} in epsilon=0.04")

	False(t, InEpsilonSlice(mock, "", nil, 1), "Expected non numeral slices to fail")
}

type numberInDeltaCase struct {
	a, b  any
	delta float64
}

func numberInDeltaCases() iter.Seq[numberInDeltaCase] {
	return slices.Values([]numberInDeltaCase{
		{uint(2), uint(1), 1},
		{uint8(2), uint8(1), 1},
		{uint16(2), uint16(1), 1},
		{uint32(2), uint32(1), 1},
		{uint64(2), uint64(1), 1},
		{int(2), int(1), 1},
		{int8(2), int8(1), 1},
		{int16(2), int16(1), 1},
		{int32(2), int32(1), 1},
		{int64(2), int64(1), 1},
		{float32(2), float32(1), 1},
		{float64(2), float64(1), 1},
	})
}

type numberInDeltaMapCase struct {
	title  string
	expect any
	actual any
	f      func(T, bool, ...any) bool
	delta  float64
}

func numberInDeltaMapCases() iter.Seq[numberInDeltaMapCase] {
	return slices.Values([]numberInDeltaMapCase{
		{
			title: "Within delta",
			expect: map[string]float64{
				"foo": 1.0,
				"bar": 2.0,
				"baz": math.NaN(),
			},
			actual: map[string]float64{
				"foo": 1.01,
				"bar": 1.99,
				"baz": math.NaN(),
			},
			delta: 0.1,
			f:     True,
		},
		{
			title: "Within delta",
			expect: map[int]float64{
				1: 1.0,
				2: 2.0,
			},
			actual: map[int]float64{
				1: 1.0,
				2: 1.99,
			},
			delta: 0.1,
			f:     True,
		},
		{
			title: "Different number of keys",
			expect: map[int]float64{
				1: 1.0,
				2: 2.0,
			},
			actual: map[int]float64{
				1: 1.0,
			},
			delta: 0.1,
			f:     False,
		},
		{
			title: "Within delta with zero value",
			expect: map[string]float64{
				"zero": 0,
			},
			actual: map[string]float64{
				"zero": 0,
			},
			delta: 0.1,
			f:     True,
		},
		{
			title: "With missing key with zero value",
			expect: map[string]float64{
				"zero": 0,
				"foo":  0,
			},
			actual: map[string]float64{
				"zero": 0,
				"bar":  0,
			},
			f: False,
		},
	})
}

type numberInEpsilonCase struct {
	a, b    any
	epsilon float64
}

func numberInEpsilonTrueCases() iter.Seq[numberInEpsilonCase] {
	return slices.Values([]numberInEpsilonCase{
		{uint8(2), uint16(2), .001},
		{2.1, 2.2, 0.1},
		{2.2, 2.1, 0.1},
		{-2.1, -2.2, 0.1},
		{-2.2, -2.1, 0.1},
		{uint64(100), uint8(101), 0.01},
		{0.1, -0.1, 2},
		{0.1, 0, 2},
		{math.NaN(), math.NaN(), 1},
		{time.Second, time.Second + time.Millisecond, 0.002},
	})
}

func numberInEpsilonFalseCases() iter.Seq[numberInEpsilonCase] {
	return slices.Values([]numberInEpsilonCase{
		{uint8(2), int16(-2), .001},
		{uint64(100), uint8(102), 0.01},
		{2.1, 2.2, 0.001},
		{2.2, 2.1, 0.001},
		{2.1, -2.2, 1},
		{2.1, "bla-bla", 0},
		{0.1, -0.1, 1.99},
		{0, 0.1, 2}, // expected must be different to zero
		{time.Second, time.Second + 10*time.Millisecond, 0.002},
		{math.NaN(), 0, 1},
		{0, math.NaN(), 1},
		{0, 0, math.NaN()},
		{math.Inf(1), 1, 1},
		{math.Inf(-1), 1, 1},
		{1, math.Inf(1), 1},
		{1, math.Inf(-1), 1},
		{math.Inf(1), math.Inf(1), 1},
		{math.Inf(1), math.Inf(-1), 1},
		{math.Inf(-1), math.Inf(1), 1},
		{math.Inf(-1), math.Inf(-1), 1},
	})
}
