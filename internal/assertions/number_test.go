// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"testing"
	"time"
)

func TestNumberInDelta(t *testing.T) {
	t.Parallel()

	t.Run("with simple input", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)

		True(t, InDelta(mock, 1.001, 1, 0.01), "|1.001 - 1| <= 0.01")
		True(t, InDelta(mock, 1, 1.001, 0.01), "|1 - 1.001| <= 0.01")
		True(t, InDelta(mock, 1, 2, 1), "|1 - 2| <= 1")
		False(t, InDelta(mock, 1, 2, 0.5), "Expected |1 - 2| <= 0.5 to fail")
		False(t, InDelta(mock, 2, 1, 0.5), "Expected |2 - 1| <= 0.5 to fail")
		False(t, InDelta(mock, "", nil, 1), "Expected non numerals to fail")
	})

	t.Run("with edge cases", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)

		False(t, InDelta(mock, 42, math.NaN(), 0.01), "Expected NaN for actual to fail")
		False(t, InDelta(mock, math.NaN(), 42, 0.01), "Expected NaN for expected to fail")
		True(t, InDelta(mock, math.NaN(), math.NaN(), 0.01), "Expected NaN for both to pass")
	})

	t.Run("should be within delta", func(t *testing.T) {
		for tc := range numberInDeltaCases() {
			t.Run(fmt.Sprintf("%f - %f", tc.a, tc.b), func(t *testing.T) {
				t.Parallel()

				mock := new(testing.T)

				True(t, InDelta(mock, tc.a, tc.b, tc.delta), "Expected |%V - %V| <= %v", tc.a, tc.b, tc.delta)
			})
		}
	})
}

func TestNumberInDeltaT(t *testing.T) {
	t.Parallel()

	t.Run("with simple input", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)

		True(t, InDeltaT(mock, 1.001, 1, 0.01), "|1.001 - 1| <= 0.01")
		True(t, InDeltaT(mock, 1, 1.001, 0.01), "|1 - 1.001| <= 0.01")
		True(t, InDeltaT(mock, 1, 2, 1), "|1 - 2| <= 1")
		False(t, InDeltaT[float32](mock, 1, 2, 0.5), "Expected |1 - 2| <= 0.5 to fail")
		False(t, InDeltaT(mock, 2, 1, 0.5), "Expected |2 - 1| <= 0.5 to fail")
	})

	t.Run("with edge cases", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)

		False(t, InDeltaT(mock, 42, math.NaN(), 0.01), "Expected NaN for actual to fail")
		False(t, InDeltaT(mock, math.NaN(), 42, 0.01), "Expected NaN for expected to fail")
		True(t, InDeltaT(mock, math.NaN(), math.NaN(), 0.01), "Expected NaN for both to pass")
	})

	for tc := range deltaTCases() {
		t.Run(tc.name, tc.test)
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

	for tc := range numberInEpsilonTrueCases() {
		t.Run("with InEpsilon true", func(t *testing.T) {
			t.Parallel()

			mock := new(testing.T)
			True(t,
				InEpsilon(mock, tc.a, tc.b, tc.epsilon,
					"Expected %V and %V to have a relative difference of %v",
					tc.a, tc.b, tc.epsilon,
				),
				"test: %q", tc,
			)
		})
	}

	for tc := range numberInEpsilonFalseCases() {
		t.Run("with InEpsilon false", func(t *testing.T) {
			t.Parallel()

			mock := new(testing.T)
			False(t,
				InEpsilon(mock, tc.a, tc.b, tc.epsilon,
					"Expected %V and %V to have a relative difference of %v",
					tc.a, tc.b, tc.epsilon,
				),
				"test: %q", tc,
			)
		})
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

func testDeltaT[Number Measurable](a, b, delta Number, success bool) func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		if success {
			True(t, InDeltaT(mock, a, b, delta))
			return
		}

		False(t, InDeltaT(mock, a, b, delta))
	}
}

func deltaTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// All integer types - basic success cases
		{"int/success", testDeltaT[int](2, 1, 1, true)},
		{"int8/success", testDeltaT[int8](2, 1, 1, true)},
		{"int16/success", testDeltaT[int16](2, 1, 1, true)},
		{"int32/success", testDeltaT[int32](2, 1, 1, true)},
		{"int64/success", testDeltaT[int64](2, 1, 1, true)},
		{"uint/success", testDeltaT[uint](2, 1, 1, true)},
		{"uint8/success", testDeltaT[uint8](2, 1, 1, true)},
		{"uint16/success", testDeltaT[uint16](2, 1, 1, true)},
		{"uint32/success", testDeltaT[uint32](2, 1, 1, true)},
		{"uint64/success", testDeltaT[uint64](2, 1, 1, true)},
		{"float32/success", testDeltaT[float32](2.0, 1.0, 1.0, true)},
		{"float64/success", testDeltaT[float64](2.0, 1.0, 1.0, true)},

		// Basic failure cases
		{"int/failure", testDeltaT[int](10, 1, 5, false)},
		{"uint/failure", testDeltaT[uint](10, 1, 5, false)},
		{"float64/failure", testDeltaT[float64](10.0, 1.0, 5.0, false)},

		// Exact match (zero delta)
		{"int/exact", testDeltaT[int](5, 5, 0, true)},
		{"uint/exact", testDeltaT[uint](5, 5, 0, true)},
		{"float64/exact", testDeltaT[float64](5.0, 5.0, 0.0, true)},

		// Zero values
		{"int/zero", testDeltaT[int](0, 0, 0, true)},
		{"uint/zero", testDeltaT[uint](0, 0, 0, true)},
		{"float64/zero", testDeltaT[float64](0.0, 0.0, 0.0, true)},
		{"int/near-zero", testDeltaT[int](1, 0, 1, true)},
		{"float64/near-zero", testDeltaT[float64](0.01, 0.0, 0.02, true)},

		// Negative numbers
		{"int/negative", testDeltaT[int](-5, -4, 2, true)},
		{"int/negative-fail", testDeltaT[int](-10, -1, 5, false)},
		{"float64/negative", testDeltaT[float64](-5.0, -4.0, 2.0, true)},

		// Mixed positive/negative
		{"int/mixed", testDeltaT[int](5, -5, 10, true)},
		{"int/mixed-fail", testDeltaT[int](5, -5, 9, false)},
		{"float64/mixed", testDeltaT[float64](5.0, -5.0, 10.0, true)},

		// Unsigned integer edge cases (overflow protection)
		{"uint/expected-greater", testDeltaT[uint](100, 50, 60, true)},
		{"uint/actual-greater", testDeltaT[uint](50, 100, 60, true)},
		{"uint8/max-value", testDeltaT[uint8](255, 250, 10, true)},
		{"uint8/max-value-fail", testDeltaT[uint8](255, 250, 4, false)},
		{"uint16/large-diff", testDeltaT[uint16](60000, 40000, 25000, true)},
		{"uint32/large-diff", testDeltaT[uint32](4000000000, 3000000000, 1000000001, true)},
		{"uint64/large-diff", testDeltaT[uint64](10000000000, 5000000000, 5000000001, true)},

		// Boundary testing for unsigned (expected > actual path)
		{"uint8/boundary-expected-gt-actual", testDeltaT[uint8](200, 100, 100, true)},
		{"uint8/boundary-expected-gt-actual-fail", testDeltaT[uint8](200, 100, 99, false)},
		// Boundary testing for unsigned (actual > expected path)
		{"uint8/boundary-actual-gt-expected", testDeltaT[uint8](100, 200, 100, true)},
		{"uint8/boundary-actual-gt-expected-fail", testDeltaT[uint8](100, 200, 99, false)},

		// Float32 NaN cases
		{"float32/both-nan", testDeltaT[float32](float32(math.NaN()), float32(math.NaN()), 1.0, true)},
		{"float32/expected-nan", testDeltaT[float32](float32(math.NaN()), 1.0, 1.0, false)},
		{"float32/actual-nan", testDeltaT[float32](1.0, float32(math.NaN()), 1.0, false)},

		// Float64 NaN cases
		{"float64/both-nan", testDeltaT[float64](math.NaN(), math.NaN(), 1.0, true)},
		{"float64/expected-nan", testDeltaT[float64](math.NaN(), 1.0, 1.0, false)},
		{"float64/actual-nan", testDeltaT[float64](1.0, math.NaN(), 1.0, false)},

		// Float32 +Inf cases
		{"float32/both-plus-inf", testDeltaT[float32](float32(math.Inf(1)), float32(math.Inf(1)), 1.0, true)},
		{"float32/expected-plus-inf-actual-minus-inf", testDeltaT[float32](float32(math.Inf(1)), float32(math.Inf(-1)), 1.0, false)},
		{"float32/expected-plus-inf-actual-finite", testDeltaT[float32](float32(math.Inf(1)), 1.0, 1.0, false)},
		{"float32/expected-finite-actual-plus-inf", testDeltaT[float32](1.0, float32(math.Inf(1)), 1.0, false)},

		// Float64 +Inf cases
		{"float64/both-plus-inf", testDeltaT[float64](math.Inf(1), math.Inf(1), 1.0, true)},
		{"float64/expected-plus-inf-actual-minus-inf", testDeltaT[float64](math.Inf(1), math.Inf(-1), 1.0, false)},
		{"float64/expected-plus-inf-actual-finite", testDeltaT[float64](math.Inf(1), 1.0, 1.0, false)},
		{"float64/expected-finite-actual-plus-inf", testDeltaT[float64](1.0, math.Inf(1), 1.0, false)},

		// Float32 -Inf cases
		{"float32/both-minus-inf", testDeltaT[float32](float32(math.Inf(-1)), float32(math.Inf(-1)), 1.0, true)},
		{"float32/expected-minus-inf-actual-plus-inf", testDeltaT[float32](float32(math.Inf(-1)), float32(math.Inf(1)), 1.0, false)},
		{"float32/expected-minus-inf-actual-finite", testDeltaT[float32](float32(math.Inf(-1)), 1.0, 1.0, false)},
		{"float32/expected-finite-actual-minus-inf", testDeltaT[float32](1.0, float32(math.Inf(-1)), 1.0, false)},

		// Float64 -Inf cases
		{"float64/both-minus-inf", testDeltaT[float64](math.Inf(-1), math.Inf(-1), 1.0, true)},
		{"float64/expected-minus-inf-actual-plus-inf", testDeltaT[float64](math.Inf(-1), math.Inf(1), 1.0, false)},
		{"float64/expected-minus-inf-actual-finite", testDeltaT[float64](math.Inf(-1), 1.0, 1.0, false)},
		{"float64/expected-finite-actual-minus-inf", testDeltaT[float64](1.0, math.Inf(-1), 1.0, false)},

		// Delta validation - NaN delta
		{"float64/delta-nan", testDeltaT[float64](1.0, 1.0, math.NaN(), false)},
		{"float32/delta-nan", testDeltaT[float32](1.0, 1.0, float32(math.NaN()), false)},

		// Delta validation - Inf delta
		{"float64/delta-plus-inf", testDeltaT[float64](1.0, 1.0, math.Inf(1), false)},
		{"float64/delta-minus-inf", testDeltaT[float64](1.0, 1.0, math.Inf(-1), false)},
		{"float32/delta-plus-inf", testDeltaT[float32](1.0, 1.0, float32(math.Inf(1)), false)},
		{"float32/delta-minus-inf", testDeltaT[float32](1.0, 1.0, float32(math.Inf(-1)), false)},

		// Very small deltas (precision testing)
		{"float64/small-delta", testDeltaT[float64](1.0, 1.0000001, 0.000001, true)},
		{"float64/small-delta-fail", testDeltaT[float64](1.0, 1.0000001, 0.00000001, false)},
		{"float32/small-delta", testDeltaT[float32](1.0, 1.00001, 0.0001, true)},
		{"float32/small-delta-fail", testDeltaT[float32](1.0, 1.00001, 0.00001, false)},

		// Large values
		{"int64/large-values", testDeltaT[int64](9223372036854775800, 9223372036854775700, 200, true)},
		{"uint64/large-values", testDeltaT[uint64](18446744073709551600, 18446744073709551500, 200, true)},
		{"float64/large-values", testDeltaT[float64](1e308, 1e308-1e307, 2e307, true)},

		// Edge case: delta is zero with different values
		{"int/zero-delta-different", testDeltaT[int](5, 6, 0, false)},
		{"float64/zero-delta-different", testDeltaT[float64](5.0, 5.00001, 0.0, false)},

		// Commutative property (order shouldn't matter)
		{"int/commutative-1", testDeltaT[int](10, 5, 6, true)},
		{"int/commutative-2", testDeltaT[int](5, 10, 6, true)},
		{"float64/commutative-1", testDeltaT[float64](10.0, 5.0, 6.0, true)},
		{"float64/commutative-2", testDeltaT[float64](5.0, 10.0, 6.0, true)},
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

func TestNumberInEpsilonT(t *testing.T) {
	t.Parallel()

	t.Run("with simple input", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)

		// Basic relative error: 100 vs 101 has 1% error, within 2% epsilon
		True(t, InEpsilonT(mock, 100.0, 101.0, 0.02), "1% error should be within 2% epsilon")
		// 100 vs 105 has 5% error, exceeds 2% epsilon
		False(t, InEpsilonT(mock, 100.0, 105.0, 0.02), "5% error should exceed 2% epsilon")
		// Exact match always passes
		True(t, InEpsilonT(mock, 100, 100, 0.0), "Exact match should pass even with 0 epsilon")
	})

	t.Run("with edge cases", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)

		// NaN cases
		False(t, InEpsilonT(mock, 42.0, math.NaN(), 0.01), "Expected NaN for actual to fail")
		False(t, InEpsilonT(mock, math.NaN(), 42.0, 0.01), "Expected NaN for expected to fail")
		True(t, InEpsilonT(mock, math.NaN(), math.NaN(), 0.01), "Expected NaN for both to pass")

		// Zero expected value uses absolute error
		True(t, InEpsilonT(mock, 0.0, 0.009, 0.01), "Zero expected: |0.009| <= 0.01 should pass")
		False(t, InEpsilonT(mock, 0.0, 0.011, 0.01), "Zero expected: |0.011| > 0.01 should fail")
	})

	for tc := range epsilonTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestInDeltaTErrorMessage(t *testing.T) {
	t.Parallel()

	mock := new(mockT)

	// Test that error message shows correct difference
	InDeltaT(mock, 10, 1, 5)

	if !mock.Failed() {
		t.Error("Expected test to fail but it passed")
	}

	// Verify the error message contains the actual difference (9)
	errorMsg := mock.errorString()
	if !Contains(t, errorMsg, "difference was 9") {
		t.Errorf("Error message should contain 'difference was 9', got: %s", errorMsg)
	}
}

func TestInEpsilonTErrorMessage(t *testing.T) {
	t.Parallel()

	t.Run("relative error message", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		// Test relative error: 100 vs 110 has 10% error, exceeds 5% epsilon
		InEpsilonT(mock, 100.0, 110.0, 0.05)

		if !mock.Failed() {
			t.Error("Expected test to fail but it passed")
		}

		// Verify the error message contains relative error
		errorMsg := mock.errorString()
		if !Contains(t, errorMsg, "Relative error is too high") {
			t.Errorf("Error message should contain 'Relative error is too high', got: %s", errorMsg)
		}
		// Should show actual relative error of 0.1 (10%)
		if !Contains(t, errorMsg, "0.1") {
			t.Errorf("Error message should contain '0.1' (10%% relative error), got: %s", errorMsg)
		}
	})

	t.Run("absolute error message for zero expected", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		// Test absolute error: expected=0, actual=0.5, epsilon=0.1
		InEpsilonT(mock, 0.0, 0.5, 0.1)

		if !mock.Failed() {
			t.Error("Expected test to fail but it passed")
		}

		// Verify the error message mentions absolute error
		errorMsg := mock.errorString()
		if !Contains(t, errorMsg, "Expected value is zero, using absolute error comparison") {
			t.Errorf("Error message should mention absolute error comparison, got: %s", errorMsg)
		}
		// Should show actual absolute difference of 0.5
		if !Contains(t, errorMsg, "0.5") {
			t.Errorf("Error message should contain '0.5' (absolute difference), got: %s", errorMsg)
		}
	})
}

func epsilonTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// All numeric types - basic success cases (12 cases)
		{"int/success", testEpsilonT[int](100, 101, 0.02, true)},             // 1% error < 2% epsilon
		{"int8/success", testEpsilonT[int8](100, 101, 0.02, true)},           // 1% error < 2% epsilon
		{"int16/success", testEpsilonT[int16](100, 101, 0.02, true)},         // 1% error < 2% epsilon
		{"int32/success", testEpsilonT[int32](100, 101, 0.02, true)},         // 1% error < 2% epsilon
		{"int64/success", testEpsilonT[int64](100, 101, 0.02, true)},         // 1% error < 2% epsilon
		{"uint/success", testEpsilonT[uint](100, 101, 0.02, true)},           // 1% error < 2% epsilon
		{"uint8/success", testEpsilonT[uint8](100, 101, 0.02, true)},         // 1% error < 2% epsilon
		{"uint16/success", testEpsilonT[uint16](100, 101, 0.02, true)},       // 1% error < 2% epsilon
		{"uint32/success", testEpsilonT[uint32](100, 101, 0.02, true)},       // 1% error < 2% epsilon
		{"uint64/success", testEpsilonT[uint64](100, 101, 0.02, true)},       // 1% error < 2% epsilon
		{"float32/success", testEpsilonT[float32](100.0, 101.0, 0.02, true)}, // 1% error < 2% epsilon
		{"float64/success", testEpsilonT[float64](100.0, 101.0, 0.02, true)}, // 1% error < 2% epsilon

		// Basic failure cases (3 cases)
		{"int/failure", testEpsilonT[int](100, 110, 0.05, false)},             // 10% error > 5% epsilon
		{"uint/failure", testEpsilonT[uint](100, 110, 0.05, false)},           // 10% error > 5% epsilon
		{"float64/failure", testEpsilonT[float64](100.0, 110.0, 0.05, false)}, // 10% error > 5% epsilon

		// Exact match (3 cases)
		{"int/exact", testEpsilonT[int](100, 100, 0.0, true)},             // Exact match
		{"uint/exact", testEpsilonT[uint](100, 100, 0.0, true)},           // Exact match
		{"float64/exact", testEpsilonT[float64](100.0, 100.0, 0.0, true)}, // Exact match

		// Zero expected value - uses absolute error (8 cases)
		{"int/both-zero", testEpsilonT[int](0, 0, 0.01, true)},                              // Both zero
		{"uint/both-zero", testEpsilonT[uint](0, 0, 0.01, true)},                            // Both zero
		{"float64/both-zero", testEpsilonT[float64](0.0, 0.0, 0.01, true)},                  // Both zero
		{"float64/zero-expected-within", testEpsilonT[float64](0.0, 0.009, 0.01, true)},     // |0.009| <= 0.01
		{"float64/zero-expected-at-boundary", testEpsilonT[float64](0.0, 0.01, 0.01, true)}, // |0.01| <= 0.01
		{"float64/zero-expected-exceed", testEpsilonT[float64](0.0, 0.011, 0.01, false)},    // |0.011| > 0.01
		{"float64/zero-expected-large", testEpsilonT[float64](0.0, 100.0, 0.01, false)},     // |100| > 0.01
		{"int/zero-expected-negative", testEpsilonT[int](0, -5, 10, true)},                  // |-5| <= 10 (absolute)

		// Near-zero values (2 cases)
		{"float64/near-zero-success", testEpsilonT[float64](0.001, 0.00101, 0.02, true)},  // 1% error < 2%
		{"float64/near-zero-failure", testEpsilonT[float64](0.001, 0.00110, 0.05, false)}, // 10% error > 5%

		// Negative numbers (3 cases)
		{"int/negative", testEpsilonT[int](-100, -101, 0.02, true)},             // 1% error < 2%
		{"int/negative-fail", testEpsilonT[int](-100, -110, 0.05, false)},       // 10% error > 5%
		{"float64/negative", testEpsilonT[float64](-100.0, -101.0, 0.02, true)}, // 1% error < 2%

		// Mixed positive/negative (3 cases)
		{"int/mixed-small", testEpsilonT[int](100, -100, 2.1, true)},       // 200% error <= 210% epsilon
		{"int/mixed-fail", testEpsilonT[int](100, -100, 1.9, false)},       // 200% error > 190% epsilon
		{"float64/mixed", testEpsilonT[float64](100.0, -100.0, 2.1, true)}, // 200% error <= 210% epsilon

		// Float32 NaN cases (3 cases)
		{"float32/both-nan", testEpsilonT[float32](float32(math.NaN()), float32(math.NaN()), 0.01, true)},
		{"float32/expected-nan", testEpsilonT[float32](float32(math.NaN()), 42.0, 0.01, false)},
		{"float32/actual-nan", testEpsilonT[float32](42.0, float32(math.NaN()), 0.01, false)},

		// Float64 NaN cases (3 cases)
		{"float64/both-nan", testEpsilonT[float64](math.NaN(), math.NaN(), 0.01, true)},
		{"float64/expected-nan", testEpsilonT[float64](math.NaN(), 42.0, 0.01, false)},
		{"float64/actual-nan", testEpsilonT[float64](42.0, math.NaN(), 0.01, false)},

		// Float32 +Inf cases (4 cases)
		{"float32/both-plus-inf", testEpsilonT[float32](float32(math.Inf(1)), float32(math.Inf(1)), 0.01, true)},
		{"float32/expected-plus-inf-actual-minus-inf", testEpsilonT[float32](float32(math.Inf(1)), float32(math.Inf(-1)), 0.01, false)},
		{"float32/expected-plus-inf-actual-finite", testEpsilonT[float32](float32(math.Inf(1)), 100.0, 0.01, false)},
		{"float32/expected-finite-actual-plus-inf", testEpsilonT[float32](100.0, float32(math.Inf(1)), 0.01, false)},

		// Float64 +Inf cases (4 cases)
		{"float64/both-plus-inf", testEpsilonT[float64](math.Inf(1), math.Inf(1), 0.01, true)},
		{"float64/expected-plus-inf-actual-minus-inf", testEpsilonT[float64](math.Inf(1), math.Inf(-1), 0.01, false)},
		{"float64/expected-plus-inf-actual-finite", testEpsilonT[float64](math.Inf(1), 100.0, 0.01, false)},
		{"float64/expected-finite-actual-plus-inf", testEpsilonT[float64](100.0, math.Inf(1), 0.01, false)},

		// Float32 -Inf cases (4 cases)
		{"float32/both-minus-inf", testEpsilonT[float32](float32(math.Inf(-1)), float32(math.Inf(-1)), 0.01, true)},
		{"float32/expected-minus-inf-actual-plus-inf", testEpsilonT[float32](float32(math.Inf(-1)), float32(math.Inf(1)), 0.01, false)},
		{"float32/expected-minus-inf-actual-finite", testEpsilonT[float32](float32(math.Inf(-1)), 100.0, 0.01, false)},
		{"float32/expected-finite-actual-minus-inf", testEpsilonT[float32](100.0, float32(math.Inf(-1)), 0.01, false)},

		// Float64 -Inf cases (4 cases)
		{"float64/both-minus-inf", testEpsilonT[float64](math.Inf(-1), math.Inf(-1), 0.01, true)},
		{"float64/expected-minus-inf-actual-plus-inf", testEpsilonT[float64](math.Inf(-1), math.Inf(1), 0.01, false)},
		{"float64/expected-minus-inf-actual-finite", testEpsilonT[float64](math.Inf(-1), 100.0, 0.01, false)},
		{"float64/expected-finite-actual-minus-inf", testEpsilonT[float64](100.0, math.Inf(-1), 0.01, false)},

		// Epsilon validation (6 cases)
		{"float64/epsilon-negative", testEpsilonT[float64](100.0, 100.0, -0.01, false)},         // Negative epsilon
		{"float64/epsilon-nan", testEpsilonT[float64](100.0, 100.0, math.NaN(), false)},         // NaN epsilon
		{"float32/epsilon-nan", testEpsilonT[float32](100.0, 100.0, math.NaN(), false)},         // NaN epsilon
		{"float64/epsilon-plus-inf", testEpsilonT[float64](100.0, 100.0, math.Inf(1), false)},   // +Inf epsilon
		{"float64/epsilon-minus-inf", testEpsilonT[float64](100.0, 100.0, math.Inf(-1), false)}, // -Inf epsilon
		{"float32/epsilon-plus-inf", testEpsilonT[float32](100.0, 100.0, math.Inf(1), false)},   // +Inf epsilon

		// Precision testing (4 cases)
		{"float64/small-epsilon-pass", testEpsilonT[float64](1.0, 1.000001, 0.00001, true)},  // Very small error
		{"float64/small-epsilon-fail", testEpsilonT[float64](1.0, 1.000011, 0.00001, false)}, // Exceeds small epsilon
		{"float32/small-epsilon-pass", testEpsilonT[float32](1.0, 1.000001, 0.00001, true)},  // Very small error
		{"float32/small-epsilon-fail", testEpsilonT[float32](1.0, 1.000011, 0.00001, false)}, // Exceeds small epsilon

		// Large values (3 cases)
		{"int64/large-values", testEpsilonT[int64](1000000000, 1010000000, 0.02, true)},   // 1% error < 2%
		{"uint64/large-values", testEpsilonT[uint64](1000000000, 1010000000, 0.02, true)}, // 1% error < 2%
		{"float64/large-values", testEpsilonT[float64](1e15, 1.01e15, 0.02, true)},        // 1% error < 2%

		// Edge cases (4 cases)
		{"int/zero-epsilon-same", testEpsilonT[int](100, 100, 0.0, true)},                   // Zero epsilon, exact match
		{"float64/zero-epsilon-different", testEpsilonT[float64](100.0, 100.1, 0.0, false)}, // Zero epsilon, different
		{"int/large-epsilon", testEpsilonT[int](100, 200, 1.5, true)},                       // 100% error < 150% epsilon
		{"float64/boundary", testEpsilonT[float64](100.0, 102.0, 0.02, true)},               // Exactly 2% error with 2% epsilon
	})
}

//nolint:thelper // linter false positive: this is not a helper
func testEpsilonT[Number Measurable](expected, actual Number, epsilon float64, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := &mockT{}
		result := InEpsilonT(mock, expected, actual, epsilon)

		if shouldPass {
			True(t, result, "Expected InEpsilonT(%v, %v, %v) to pass", expected, actual, epsilon)
			False(t, mock.Failed(), "Mock should not have failed")
		} else {
			False(t, result, "Expected InEpsilonT(%v, %v, %v) to fail", expected, actual, epsilon)
			True(t, mock.Failed(), "Mock should have failed")
		}
	}
}
