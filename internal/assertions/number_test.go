// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"math"
	"slices"
	"testing"
)

func TestNumberInDeltaEdgeCases(t *testing.T) {
	t.Parallel()

	t.Run("InDelta specific (type conversion)", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)
		result := InDelta(mock, "", nil, 1)
		if result {
			t.Errorf("Expected non numerals to fail")
		}
	})
}

func TestNumberInDelta(t *testing.T) {
	t.Parallel()

	// run all test cases with both InDelta and InDeltaT
	//
	// NOTE: testing pattern, focused on the expected result (true/false) and _NOT_ the content of the returned message.
	// - deltaCases: loop over generic test cases AND type combinations (reason: not all types are compatible, e.g. uint64 may overflow float64)
	//    - testAllDelta: dispatch over the assertion variants (reflection-based, generic, X vs NotX semantics)
	//      Single assertion test functions:
	//      - testDelta
	//      - testDeltaT
	for tc := range deltaCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestNumberInDeltaSlice(t *testing.T) {
	t.Parallel()

	// only have a reflection-based assertion here
	for tc := range deltaSliceCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestNumberInDeltaMapValues(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	// only have a reflection-based assertion here
	for tc := range numberInDeltaMapCases() {
		tc.f(t, InDeltaMapValues(mock, tc.expect, tc.actual, tc.delta), tc.name+"\n"+diff(tc.expect, tc.actual))
	}
}

func TestNumberInEpsilon(t *testing.T) {
	t.Parallel()

	// run all test cases with both InEpsilon and InEpsilonT
	for tc := range epsilonCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestNumberInEpsilonSlice(t *testing.T) {
	t.Parallel()

	for tc := range epsilonSliceCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestNumberErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, numberFailCases())
}

// =======================================
// Test NumberInDelta variants
// =======================================

func deltaCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// Simple input cases
		{"simple/within-delta-1.001-1-0.01", testAllDelta(1.001, 1.0, 0.01, true)},
		{"simple/within-delta-1-1.001-0.01", testAllDelta(1.0, 1.001, 0.01, true)},
		{"simple/within-delta-1-2-1", testAllDelta(1.0, 2.0, 1.0, true)},
		{"simple/exceeds-delta-1-2-0.5", testAllDelta(1.0, 2.0, 0.5, false)},
		{"simple/exceeds-delta-2-1-0.5", testAllDelta(2.0, 1.0, 0.5, false)},

		// Edge cases - NaN
		{"edge/nan-for-actual", testAllDelta(42.0, math.NaN(), 0.01, false)},
		{"edge/nan-for-expected", testAllDelta(math.NaN(), 42.0, 0.01, false)},
		{"edge/nan-for-both", testAllDelta(math.NaN(), math.NaN(), 0.01, true)},

		// All integer types - basic success cases
		{"int/success", testAllDelta(int(2), int(1), int(1), true)},
		{"int8/success", testAllDelta(int8(2), int8(1), int8(1), true)},
		{"int16/success", testAllDelta(int16(2), int16(1), int16(1), true)},
		{"int32/success", testAllDelta(int32(2), int32(1), int32(1), true)},
		{"int64/success", testAllDelta(int64(2), int64(1), int64(1), true)},
		{"uint/success", testAllDelta(uint(2), uint(1), uint(1), true)},
		{"uint8/success", testAllDelta(uint8(2), uint8(1), uint8(1), true)},
		{"uint16/success", testAllDelta(uint16(2), uint16(1), uint16(1), true)},
		{"uint32/success", testAllDelta(uint32(2), uint32(1), uint32(1), true)},
		{"uint64/success", testAllDelta(uint64(2), uint64(1), uint64(1), true)},
		{"float32/success", testAllDelta(float32(2.0), float32(1.0), float32(1.0), true)},
		{"float64/success", testAllDelta(2.0, 1.0, 1.0, true)},

		// Basic failure cases
		{"int/failure", testAllDelta(int(10), int(1), int(5), false)},
		{"uint/failure", testAllDelta(uint(10), uint(1), uint(5), false)},
		{"float64/failure", testAllDelta(10.0, 1.0, 5.0, false)},

		// Exact match (zero delta)
		{"int/exact", testAllDelta(int(5), int(5), int(0), true)},
		{"uint/exact", testAllDelta(uint(5), uint(5), uint(0), true)},
		{"float64/exact", testAllDelta(5.0, 5.0, 0.0, true)},

		// Zero values
		{"int/zero", testAllDelta(int(0), int(0), int(0), true)},
		{"uint/zero", testAllDelta(uint(0), uint(0), uint(0), true)},
		{"float64/zero", testAllDelta(0.0, 0.0, 0.0, true)},
		{"int/near-zero", testAllDelta(int(1), int(0), int(1), true)},
		{"float64/near-zero", testAllDelta(0.01, 0.0, 0.02, true)},

		// Negative numbers
		{"int/negative", testAllDelta(int(-5), int(-4), int(2), true)},
		{"int/negative-fail", testAllDelta(int(-10), int(-1), int(5), false)},
		{"float64/negative", testAllDelta(-5.0, -4.0, 2.0, true)},

		// Mixed positive/negative
		{"int/mixed", testAllDelta(int(5), int(-5), int(10), true)},
		{"int/mixed-fail", testAllDelta(int(5), int(-5), int(9), false)},
		{"float64/mixed", testAllDelta(5.0, -5.0, 10.0, true)},

		// Unsigned integer edge cases (overflow protection)
		{"uint/expected-greater", testAllDelta(uint(100), uint(50), uint(60), true)},
		{"uint/actual-greater", testAllDelta(uint(50), uint(100), uint(60), true)},
		{"uint8/max-value", testAllDelta(uint8(255), uint8(250), uint8(10), true)},
		{"uint8/max-value-fail", testAllDelta(uint8(255), uint8(250), uint8(4), false)},
		{"uint16/large-diff", testAllDelta(uint16(60000), uint16(40000), uint16(25000), true)},
		{"uint32/large-diff", testAllDelta(uint32(4000000000), uint32(3000000000), uint32(1000000001), true)},
		{"uint64/large-diff", testAllDelta(uint64(10000000000), uint64(5000000000), uint64(5000000001), true)},

		// Boundary testing for unsigned (expected > actual path)
		{"uint8/boundary-expected-gt-actual", testAllDelta(uint8(200), uint8(100), uint8(100), true)},
		{"uint8/boundary-expected-gt-actual-fail", testAllDelta(uint8(200), uint8(100), uint8(99), false)},
		// Boundary testing for unsigned (actual > expected path)
		{"uint8/boundary-actual-gt-expected", testAllDelta(uint8(100), uint8(200), uint8(100), true)},
		{"uint8/boundary-actual-gt-expected-fail", testAllDelta(uint8(100), uint8(200), uint8(99), false)},

		// Float32 NaN cases
		{"float32/both-nan", testAllDelta(float32(math.NaN()), float32(math.NaN()), float32(1.0), true)},
		{"float32/expected-nan", testAllDelta(float32(math.NaN()), float32(1.0), float32(1.0), false)},
		{"float32/actual-nan", testAllDelta(float32(1.0), float32(math.NaN()), float32(1.0), false)},

		// Float64 NaN cases
		{"float64/both-nan", testAllDelta(math.NaN(), math.NaN(), 1.0, true)},
		{"float64/expected-nan", testAllDelta(math.NaN(), 1.0, 1.0, false)},
		{"float64/actual-nan", testAllDelta(1.0, math.NaN(), 1.0, false)},

		// Float32 +Inf cases
		{"float32/both-plus-inf", testAllDelta(float32(math.Inf(1)), float32(math.Inf(1)), float32(1.0), true)},
		{"float32/expected-plus-inf-actual-minus-inf", testAllDelta(float32(math.Inf(1)), float32(math.Inf(-1)), float32(1.0), false)},
		{"float32/expected-plus-inf-actual-finite", testAllDelta(float32(math.Inf(1)), float32(1.0), float32(1.0), false)},
		{"float32/expected-finite-actual-plus-inf", testAllDelta(float32(1.0), float32(math.Inf(1)), float32(1.0), false)},

		// Float64 +Inf cases
		{"float64/both-plus-inf", testAllDelta(math.Inf(1), math.Inf(1), 1.0, true)},
		{"float64/expected-plus-inf-actual-minus-inf", testAllDelta(math.Inf(1), math.Inf(-1), 1.0, false)},
		{"float64/expected-plus-inf-actual-finite", testAllDelta(math.Inf(1), 1.0, 1.0, false)},
		{"float64/expected-finite-actual-plus-inf", testAllDelta(1.0, math.Inf(1), 1.0, false)},

		// Float32 -Inf cases
		{"float32/both-minus-inf", testAllDelta(float32(math.Inf(-1)), float32(math.Inf(-1)), float32(1.0), true)},
		{"float32/expected-minus-inf-actual-plus-inf", testAllDelta(float32(math.Inf(-1)), float32(math.Inf(1)), float32(1.0), false)},
		{"float32/expected-minus-inf-actual-finite", testAllDelta(float32(math.Inf(-1)), float32(1.0), float32(1.0), false)},
		{"float32/expected-finite-actual-minus-inf", testAllDelta(float32(1.0), float32(math.Inf(-1)), float32(1.0), false)},

		// Float64 -Inf cases
		{"float64/both-minus-inf", testAllDelta(math.Inf(-1), math.Inf(-1), 1.0, true)},
		{"float64/expected-minus-inf-actual-plus-inf", testAllDelta(math.Inf(-1), math.Inf(1), 1.0, false)},
		{"float64/expected-minus-inf-actual-finite", testAllDelta(math.Inf(-1), 1.0, 1.0, false)},
		{"float64/expected-finite-actual-minus-inf", testAllDelta(1.0, math.Inf(-1), 1.0, false)},

		// Delta validation - NaN delta
		{"float64/delta-nan", testAllDelta(1.0, 1.0, math.NaN(), false)},
		{"float32/delta-nan", testAllDelta(float32(1.0), float32(1.0), float32(math.NaN()), false)},

		// Delta validation - Inf delta
		{"float64/delta-plus-inf", testAllDelta(1.0, 1.0, math.Inf(1), false)},
		{"float64/delta-minus-inf", testAllDelta(1.0, 1.0, math.Inf(-1), false)},
		{"float32/delta-plus-inf", testAllDelta(float32(1.0), float32(1.0), float32(math.Inf(1)), false)},
		{"float32/delta-minus-inf", testAllDelta(float32(1.0), float32(1.0), float32(math.Inf(-1)), false)},

		// Very small deltas (precision testing)
		{"float64/small-delta", testAllDelta(1.0, 1.0000001, 0.000001, true)},
		{"float64/small-delta-fail", testAllDelta(1.0, 1.0000001, 0.00000001, false)},
		{"float32/small-delta", testAllDelta(float32(1.0), float32(1.00001), float32(0.0001), true)},
		{"float32/small-delta-fail", testAllDelta(float32(1.0), float32(1.00001), float32(0.00001), false)},

		// Large values
		{"int64/large-values", testAllDelta(int64(9223372036854775800), int64(9223372036854775700), int64(200), true)},
		{"uint64/large-values", testAllDelta(uint64(18446744073709551600), uint64(18446744073709551500), uint64(200), true)},
		{"float64/large-values", testAllDelta(1e308, 1e308-1e307, 2e307, true)},

		// Edge case: delta is zero with different values
		{"int/zero-delta-different", testAllDelta(int(5), int(6), int(0), false)},
		{"float64/zero-delta-different", testAllDelta(5.0, 5.00001, 0.0, false)},

		// Commutative property (order shouldn't matter)
		{"int/commutative-1", testAllDelta(int(10), int(5), int(6), true)},
		{"int/commutative-2", testAllDelta(int(5), int(10), int(6), true)},
		{"float64/commutative-1", testAllDelta(10.0, 5.0, 6.0, true)},
		{"float64/commutative-2", testAllDelta(5.0, 10.0, 6.0, true)},
	})
}

// testAllDelta tests both InDelta and InDeltaT with the same input.
//
//nolint:thelper // linter false positive: this is not a helper
func testAllDelta[Number Measurable](expected, actual, delta Number, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if shouldPass {
			t.Run("should pass", func(t *testing.T) {
				t.Run("with InDelta", testDelta(expected, actual, delta, true))
				t.Run("with InDeltaT", testDeltaT(expected, actual, delta, true))
			})
		} else {
			t.Run("should fail", func(t *testing.T) {
				t.Run("with InDelta", testDelta(expected, actual, delta, false))
				t.Run("with InDeltaT", testDeltaT(expected, actual, delta, false))
			})
		}
	}
}

func testDelta[Number Measurable](expected, actual, delta Number, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		// InDelta requires delta as float64, so convert it
		result := InDelta(mock, expected, actual, float64(delta))

		if shouldPass {
			True(t, result)
			False(t, mock.Failed())
		} else {
			False(t, result)
			True(t, mock.Failed())
		}
	}
}

func testDeltaT[Number Measurable](expected, actual, delta Number, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaT(mock, expected, actual, delta)

		if shouldPass {
			True(t, result)
			False(t, mock.Failed())
		} else {
			False(t, result)
			True(t, mock.Failed())
		}
	}
}

// Helper functions and test data for InDeltaMapValues

type numberInDeltaMapCase struct {
	name   string
	expect any
	actual any
	f      func(T, bool, ...any) bool
	delta  float64
}

func numberInDeltaMapCases() iter.Seq[numberInDeltaMapCase] {
	keyA := "a"
	var iface any

	return slices.Values([]numberInDeltaMapCase{
		{
			name: "Within delta",
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
			name: "Within delta",
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
			name: "Different number of keys",
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
			name: "Within delta with zero value",
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
			name: "With missing key with zero value",
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
		{
			name:   "With nil maps",
			expect: map[string]float64(nil),
			actual: map[string]float64(nil),
			f:      True,
		},
		{
			name:   "With nil values (not a map)",
			expect: map[string]float64(nil),
			actual: []float64(nil),
			f:      False,
		},
		{
			name:   "With nil values (not a map)",
			expect: []float64(nil),
			actual: map[string]float64(nil),
			f:      False,
		},
		{
			name: "With expected nil keys",
			expect: map[*string]float64{
				&keyA:          1.00,
				(*string)(nil): 2.00,
			},
			actual: map[*string]float64{
				&keyA:          1.00,
				(*string)(nil): 2.00,
			},
			f: True,
		},
		{
			name: "With expected invalid value",
			expect: map[string]any{
				keyA: &iface,
			},
			actual: map[string]any{
				keyA: &iface,
			},
			f: False,
		},
	})
}

// =======================================
// Test NumberInEpsilon variants
// =======================================

func epsilonCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// Simple input cases
		{"simple/1pct-error-within-2pct-epsilon", testAllEpsilon(100.0, 101.0, 0.02, true)},
		{"simple/5pct-error-exceeds-2pct-epsilon", testAllEpsilon(100.0, 105.0, 0.02, false)},
		{"simple/exact-match-zero-epsilon", testAllEpsilon(100.0, 100.0, 0.0, true)},

		// Edge cases - NaN
		{"edge/nan-for-actual", testAllEpsilon(42.0, math.NaN(), 0.01, false)},
		{"edge/nan-for-expected", testAllEpsilon(math.NaN(), 42.0, 0.01, false)},
		{"edge/nan-for-both", testAllEpsilon(math.NaN(), math.NaN(), 0.01, true)},

		// Edge cases - zero expected (uses absolute error)
		{"edge/zero-expected-within", testAllEpsilon(0.0, 0.009, 0.01, true)},
		{"edge/zero-expected-exceeds", testAllEpsilon(0.0, 0.011, 0.01, false)},

		// All numeric types - basic success cases (12 cases)
		{"int/success", testAllEpsilon(int(100), int(101), 0.02, true)},                 // 1% error < 2% epsilon
		{"int8/success", testAllEpsilon(int8(100), int8(101), 0.02, true)},              // 1% error < 2% epsilon
		{"int16/success", testAllEpsilon(int16(100), int16(101), 0.02, true)},           // 1% error < 2% epsilon
		{"int32/success", testAllEpsilon(int32(100), int32(101), 0.02, true)},           // 1% error < 2% epsilon
		{"int64/success", testAllEpsilon(int64(100), int64(101), 0.02, true)},           // 1% error < 2% epsilon
		{"uint/success", testAllEpsilon(uint(100), uint(101), 0.02, true)},              // 1% error < 2% epsilon
		{"uint8/success", testAllEpsilon(uint8(100), uint8(101), 0.02, true)},           // 1% error < 2% epsilon
		{"uint16/success", testAllEpsilon(uint16(100), uint16(101), 0.02, true)},        // 1% error < 2% epsilon
		{"uint32/success", testAllEpsilon(uint32(100), uint32(101), 0.02, true)},        // 1% error < 2% epsilon
		{"uint64/success", testAllEpsilon(uint64(100), uint64(101), 0.02, true)},        // 1% error < 2% epsilon
		{"float32/success", testAllEpsilon(float32(100.0), float32(101.0), 0.02, true)}, // 1% error < 2% epsilon
		{"float64/success", testAllEpsilon(100.0, 101.0, 0.02, true)},                   // 1% error < 2% epsilon

		// Basic failure cases (3 cases)
		{"int/failure", testAllEpsilon(int(100), int(110), 0.05, false)},    // 10% error > 5% epsilon
		{"uint/failure", testAllEpsilon(uint(100), uint(110), 0.05, false)}, // 10% error > 5% epsilon
		{"float64/failure", testAllEpsilon(100.0, 110.0, 0.05, false)},      // 10% error > 5% epsilon

		// Exact match (3 cases)
		{"int/exact", testAllEpsilon(int(100), int(100), 0.0, true)},    // Exact match
		{"uint/exact", testAllEpsilon(uint(100), uint(100), 0.0, true)}, // Exact match
		{"float64/exact", testAllEpsilon(100.0, 100.0, 0.0, true)},      // Exact match

		// Zero expected value - uses absolute error (8 cases)
		{"int/both-zero", testAllEpsilon(int(0), int(0), 0.01, true)},                // Both zero
		{"uint/both-zero", testAllEpsilon(uint(0), uint(0), 0.01, true)},             // Both zero
		{"float64/both-zero", testAllEpsilon(0.0, 0.0, 0.01, true)},                  // Both zero
		{"float64/zero-expected-within", testAllEpsilon(0.0, 0.009, 0.01, true)},     // |0.009| <= 0.01
		{"float64/zero-expected-at-boundary", testAllEpsilon(0.0, 0.01, 0.01, true)}, // |0.01| <= 0.01
		{"float64/zero-expected-exceed", testAllEpsilon(0.0, 0.011, 0.01, false)},    // |0.011| > 0.01
		{"float64/zero-expected-large", testAllEpsilon(0.0, 100.0, 0.01, false)},     // |100| > 0.01
		{"int/zero-expected-negative", testAllEpsilon(int(0), int(-5), 10.0, true)},  // |-5| <= 10 (absolute)

		// Near-zero values (2 cases)
		{"float64/near-zero-success", testAllEpsilon(0.001, 0.00101, 0.02, true)},  // 1% error < 2%
		{"float64/near-zero-failure", testAllEpsilon(0.001, 0.00110, 0.05, false)}, // 10% error > 5%

		// Negative numbers (3 cases)
		{"int/negative", testAllEpsilon(int(-100), int(-101), 0.02, true)},       // 1% error < 2%
		{"int/negative-fail", testAllEpsilon(int(-100), int(-110), 0.05, false)}, // 10% error > 5%
		{"float64/negative", testAllEpsilon(-100.0, -101.0, 0.02, true)},         // 1% error < 2%

		// Mixed positive/negative (3 cases)
		{"int/mixed-small", testAllEpsilon(int(100), int(-100), 2.1, true)}, // 200% error <= 210% epsilon
		{"int/mixed-fail", testAllEpsilon(int(100), int(-100), 1.9, false)}, // 200% error > 190% epsilon
		{"float64/mixed", testAllEpsilon(100.0, -100.0, 2.1, true)},         // 200% error <= 210% epsilon

		// Float32 NaN cases (3 cases)
		{"float32/both-nan", testAllEpsilon(float32(math.NaN()), float32(math.NaN()), 0.01, true)},
		{"float32/expected-nan", testAllEpsilon(float32(math.NaN()), float32(42.0), 0.01, false)},
		{"float32/actual-nan", testAllEpsilon(float32(42.0), float32(math.NaN()), 0.01, false)},

		// Float64 NaN cases (3 cases)
		{"float64/both-nan", testAllEpsilon(math.NaN(), math.NaN(), 0.01, true)},
		{"float64/expected-nan", testAllEpsilon(math.NaN(), 42.0, 0.01, false)},
		{"float64/actual-nan", testAllEpsilon(42.0, math.NaN(), 0.01, false)},

		// Float32 +Inf cases (4 cases)
		{"float32/both-plus-inf", testAllEpsilon(float32(math.Inf(1)), float32(math.Inf(1)), 0.01, true)},
		{"float32/expected-plus-inf-actual-minus-inf", testAllEpsilon(float32(math.Inf(1)), float32(math.Inf(-1)), 0.01, false)},
		{"float32/expected-plus-inf-actual-finite", testAllEpsilon(float32(math.Inf(1)), float32(100.0), 0.01, false)},
		{"float32/expected-finite-actual-plus-inf", testAllEpsilon(float32(100.0), float32(math.Inf(1)), 0.01, false)},

		// Float64 +Inf cases (4 cases)
		{"float64/both-plus-inf", testAllEpsilon(math.Inf(1), math.Inf(1), 0.01, true)},
		{"float64/expected-plus-inf-actual-minus-inf", testAllEpsilon(math.Inf(1), math.Inf(-1), 0.01, false)},
		{"float64/expected-plus-inf-actual-finite", testAllEpsilon(math.Inf(1), 100.0, 0.01, false)},
		{"float64/expected-finite-actual-plus-inf", testAllEpsilon(100.0, math.Inf(1), 0.01, false)},

		// Float32 -Inf cases (4 cases)
		{"float32/both-minus-inf", testAllEpsilon(float32(math.Inf(-1)), float32(math.Inf(-1)), 0.01, true)},
		{"float32/expected-minus-inf-actual-plus-inf", testAllEpsilon(float32(math.Inf(-1)), float32(math.Inf(1)), 0.01, false)},
		{"float32/expected-minus-inf-actual-finite", testAllEpsilon(float32(math.Inf(-1)), float32(100.0), 0.01, false)},
		{"float32/expected-finite-actual-minus-inf", testAllEpsilon(float32(100.0), float32(math.Inf(-1)), 0.01, false)},

		// Float64 -Inf cases (4 cases)
		{"float64/both-minus-inf", testAllEpsilon(math.Inf(-1), math.Inf(-1), 0.01, true)},
		{"float64/expected-minus-inf-actual-plus-inf", testAllEpsilon(math.Inf(-1), math.Inf(1), 0.01, false)},
		{"float64/expected-minus-inf-actual-finite", testAllEpsilon(math.Inf(-1), 100.0, 0.01, false)},
		{"float64/expected-finite-actual-minus-inf", testAllEpsilon(100.0, math.Inf(-1), 0.01, false)},

		// Epsilon validation (6 cases)
		{"float64/epsilon-negative", testAllEpsilon(100.0, 100.0, -0.01, false)},                         // Negative epsilon
		{"float64/epsilon-nan", testAllEpsilon(100.0, 100.0, math.NaN(), false)},                         // NaN epsilon
		{"float32/epsilon-nan", testAllEpsilon(float32(100.0), float32(100.0), math.NaN(), false)},       // NaN epsilon
		{"float64/epsilon-plus-inf", testAllEpsilon(100.0, 100.0, math.Inf(1), false)},                   // +Inf epsilon
		{"float64/epsilon-minus-inf", testAllEpsilon(100.0, 100.0, math.Inf(-1), false)},                 // -Inf epsilon
		{"float32/epsilon-plus-inf", testAllEpsilon(float32(100.0), float32(100.0), math.Inf(1), false)}, // +Inf epsilon

		// Precision testing (4 cases)
		{"float64/small-epsilon-pass", testAllEpsilon(1.0, 1.000001, 0.00001, true)},                    // Very small error
		{"float64/small-epsilon-fail", testAllEpsilon(1.0, 1.000011, 0.00001, false)},                   // Exceeds small epsilon
		{"float32/small-epsilon-pass", testAllEpsilon(float32(1.0), float32(1.000001), 0.00001, true)},  // Very small error
		{"float32/small-epsilon-fail", testAllEpsilon(float32(1.0), float32(1.000011), 0.00001, false)}, // Exceeds small epsilon

		// Large values (3 cases)
		{"int64/large-values", testAllEpsilon(int64(1000000000), int64(1010000000), 0.02, true)},    // 1% error < 2%
		{"uint64/large-values", testAllEpsilon(uint64(1000000000), uint64(1010000000), 0.02, true)}, // 1% error < 2%
		{"float64/large-values", testAllEpsilon(1e15, 1.01e15, 0.02, true)},                         // 1% error < 2%

		// Edge cases (4 cases)
		{"int/zero-epsilon-same", testAllEpsilon(int(100), int(100), 0.0, true)},     // Zero epsilon, exact match
		{"float64/zero-epsilon-different", testAllEpsilon(100.0, 100.1, 0.0, false)}, // Zero epsilon, different
		{"int/large-epsilon", testAllEpsilon(int(100), int(200), 1.5, true)},         // 100% error < 150% epsilon
		{"float64/boundary", testAllEpsilon(100.0, 102.0, 0.02, true)},               // Exactly 2% error with 2% epsilon
	})
}

// testAllEpsilon tests both InEpsilon and InEpsilonT with the same input.
//
//nolint:thelper // linter false positive: this is not a helper
func testAllEpsilon[Number Measurable](expected, actual Number, epsilon float64, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		if shouldPass {
			t.Run("should pass", func(t *testing.T) {
				t.Run("with InEpsilon", testEpsilon(expected, actual, epsilon, true))
				t.Run("with InEpsilonT", testEpsilonT(expected, actual, epsilon, true))
			})
		} else {
			t.Run("should fail", func(t *testing.T) {
				t.Run("with InEpsilon", testEpsilon(expected, actual, epsilon, false))
				t.Run("with InEpsilonT", testEpsilonT(expected, actual, epsilon, false))
			})
		}
	}
}

func testEpsilon[Number Measurable](expected, actual Number, epsilon float64, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilon(mock, expected, actual, epsilon)

		if shouldPass {
			True(t, result)
			False(t, mock.Failed())
		} else {
			False(t, result)
			True(t, mock.Failed())
		}
	}
}

func testEpsilonT[Number Measurable](expected, actual Number, epsilon float64, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilonT(mock, expected, actual, epsilon)

		if shouldPass {
			True(t, result)
			False(t, mock.Failed())
		} else {
			False(t, result)
			True(t, mock.Failed())
		}
	}
}

// Helper functions and test data for InDeltaSlice

func deltaSliceCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// Success cases - slices are element-wise within delta
		{
			"within-delta-with-nan",
			testDeltaSlice(
				[]float64{1.001, math.NaN(), 0.999},
				[]float64{1, math.NaN(), 1},
				0.1,
				true,
			),
		},
		{
			"within-delta-1.0",
			testDeltaSlice(
				[]float64{1, math.NaN(), 2},
				[]float64{0, math.NaN(), 3},
				1,
				true,
			),
		},

		// Failure cases - slices are not element-wise within delta
		{
			"not-within-delta",
			testDeltaSlice(
				[]float64{1, math.NaN(), 2},
				[]float64{0, math.NaN(), 3},
				0.1,
				false,
			),
		},

		// Edge cases - invalid inputs
		{
			"invalid-non-slice-inputs",
			testDeltaSlice("", nil, 1, false),
		},
	})
}

//nolint:thelper // linter false positive: this is not a helper
func testDeltaSlice(expected, actual any, delta float64, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaSlice(mock, expected, actual, delta)

		if shouldPass {
			True(t, result)
			False(t, mock.Failed())
		} else {
			False(t, result)
			True(t, mock.Failed())
		}
	}
}

// Helper functions and test data for InEpsilonSlice

func epsilonSliceCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// Success cases - slices are element-wise within epsilon
		{
			"within-epsilon-with-nan",
			testEpsilonSlice(
				[]float64{2.2, math.NaN(), 2.0},
				[]float64{2.1, math.NaN(), 2.1},
				0.06,
				true,
			),
		},

		// Failure cases - slices are not element-wise within epsilon
		{
			"not-within-epsilon",
			testEpsilonSlice(
				[]float64{2.2, 2.0},
				[]float64{2.1, 2.1},
				0.04,
				false,
			),
		},

		// Edge cases - invalid inputs
		{
			"invalid-expected-nil",
			testEpsilonSlice("", nil, 1, false),
		},
		{
			"invalid-actual-nil",
			testEpsilonSlice(nil, "", 1, false),
		},
		{
			"invalid-expected-not-slice",
			testEpsilonSlice(1, []int{}, 1, false),
		},
		{
			"invalid-actual-not-slice",
			testEpsilonSlice([]int{}, 1, 1, false),
		},
		{
			"invalid-expected-non-numeric-slice",
			testEpsilonSlice([]string{}, []int{}, 1, false),
		},
		{
			"invalid-actual-non-numeric-slice",
			testEpsilonSlice([]int{}, []string{}, 1, false),
		},
	})
}

//nolint:thelper // linter false positive: this is not a helper
func testEpsilonSlice(expected, actual any, epsilon float64, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilonSlice(mock, expected, actual, epsilon)

		if shouldPass {
			True(t, result)
			False(t, mock.Failed())
		} else {
			False(t, result)
			True(t, mock.Failed())
		}
	}
}

// =======================================
// Test NumberErrorMessages
// =======================================

func numberFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "InDeltaT/shows-difference",
			assertion:    func(t T) bool { return InDeltaT(t, 10, 1, 5) },
			wantContains: []string{"difference was 9"},
		},
		{
			name:         "InEpsilonT/relative-error",
			assertion:    func(t T) bool { return InEpsilonT(t, 100.0, 110.0, 0.05) },
			wantContains: []string{"Relative error is too high", "0.1"},
		},
		{
			name:         "InEpsilonT/absolute-error-for-zero",
			assertion:    func(t T) bool { return InEpsilonT(t, 0.0, 0.5, 0.1) },
			wantContains: []string{"Expected value is zero, using absolute error comparison", "0.5"},
		},
	})
}
