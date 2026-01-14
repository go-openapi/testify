// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"iter"
	"runtime"
	"slices"
	"strings"
	"testing"
	"time"
)

const pkg = "github.com/go-openapi/testify/v2/internal/assertions"

//nolint:dupl // no this is not a duplicate: it just looks almost the same!
func TestCompareGreater(t *testing.T) {
	t.Parallel()

	t.Run("with basic input", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !Greater(mock, 2, 1) {
			t.Error("Greater should return true")
		}

		if Greater(mock, 1, 1) {
			t.Error("Greater should return false")
		}

		if Greater(mock, 1, 2) {
			t.Error("Greater should return false")
		}
	})

	for currCase := range compareStrictlyGreaterCases() {
		t.Run("should NOT be strictly greater, with expected error message", func(t *testing.T) {
			t.Parallel()

			mock := &outputT{buf: bytes.NewBuffer(nil)} // check error report
			False(t, Greater(mock, currCase.less, currCase.greater),
				"expected %v NOT to be strictly greater than %v",
				currCase.less, currCase.greater,
			)

			Contains(t, mock.buf.String(), currCase.msg)
			Contains(t, mock.helpers, pkg+".Greater")
		})

		t.Run("should be strictly greater", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT) // don't check output
			True(t, Greater(mock, currCase.greater, currCase.less),
				"expected %v to be strictly greater than %v",
				currCase.less, currCase.greater,
			)
		})
	}

	for currCase := range compareEqualCases() {
		t.Run("equal values should NOT be strictly greater", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			False(t, Greater(mock, currCase.less, currCase.greater),
				"expected (equal) %v NOT to be strictly greater than %v",
				currCase.less, currCase.greater,
			)
		})
	}
}

func TestCompareGreaterOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("with basic input", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !GreaterOrEqual(mock, 2, 1) {
			t.Error("GreaterOrEqual should return true")
		}

		if !GreaterOrEqual(mock, 1, 1) {
			t.Error("GreaterOrEqual should return true")
		}

		if GreaterOrEqual(mock, 1, 2) {
			t.Error("GreaterOrEqual should return false")
		}
	})

	for currCase := range compareStrictlyGreaterCases() {
		t.Run("should NOT be greater or equal, with expected error message", func(t *testing.T) {
			t.Parallel()

			mock := &outputT{buf: bytes.NewBuffer(nil)} // check error report

			False(t, GreaterOrEqual(mock, currCase.less, currCase.greater),
				"expected %v NOT to be greater than or equal to %v",
				currCase.less, currCase.greater,
			)

			Contains(t, mock.buf.String(), strings.ReplaceAll(currCase.msg, "than", "than or equal to"))
			Contains(t, mock.helpers, pkg+".GreaterOrEqual")
		})

		t.Run("should be greater or equal", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)

			True(t, GreaterOrEqual(mock, currCase.greater, currCase.less),
				"expected %v to be greater than or equal to %v",
				currCase.less, currCase.greater,
			)
		})
	}

	for currCase := range compareEqualCases() {
		t.Run("equal values should be greater or equal", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)

			True(t, GreaterOrEqual(mock, currCase.less, currCase.greater),
				"expected (equal) %v to be greater than or equal to %v",
				currCase.less, currCase.greater,
			)
		})
	}
}

//nolint:dupl // no this is not a duplicate: it just looks almost the same!
func TestCompareLess(t *testing.T) {
	t.Parallel()
	t.Run("with basic input", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		if !Less(mock, 1, 2) {
			t.Error("Less should return true")
		}

		if Less(mock, 1, 1) {
			t.Error("Less should return false")
		}

		if Less(mock, 2, 1) {
			t.Error("Less should return false")
		}
	})

	for currCase := range compareStrictlyLessCases() {
		t.Run("should NOT be stricly less, with expected error message", func(t *testing.T) {
			t.Parallel()

			mock := &outputT{buf: bytes.NewBuffer(nil)} // check error report
			False(t, Less(mock, currCase.greater, currCase.less),
				"expected %v NOT to be stricly less than %v",
				currCase.greater, currCase.less,
			)

			Contains(t, mock.buf.String(), currCase.msg)
			Contains(t, mock.helpers, pkg+".Less")
		})

		t.Run("should be stricly less", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)

			True(t, Less(mock, currCase.less, currCase.greater),
				"expected %v be stricly less than %v",
				currCase.less, currCase.greater,
			)
		})
	}

	for currCase := range compareEqualCases() {
		t.Run("equal values should NOT be strictly less", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)

			True(t, GreaterOrEqual(mock, currCase.less, currCase.greater),
				"expected (equal) %v NOT to be strictly less than %v",
				currCase.less, currCase.greater,
			)
		})
	}
}

func TestCompareLessOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("with basic input", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		if !LessOrEqual(mock, 1, 2) {
			t.Error("LessOrEqual should return true")
		}

		if !LessOrEqual(mock, 1, 1) {
			t.Error("LessOrEqual should return true")
		}

		if LessOrEqual(mock, 2, 1) {
			t.Error("LessOrEqual should return false")
		}
	})

	for currCase := range compareStrictlyLessCases() {
		t.Run("should NOT be less or equal, with expected error message", func(t *testing.T) {
			t.Parallel()

			mock := &outputT{buf: bytes.NewBuffer(nil)} // check error report

			False(t, LessOrEqual(mock, currCase.greater, currCase.less),
				"expected %v NOT to be less than or equal to %v",
				currCase.less, currCase.greater,
			)

			Contains(t, mock.buf.String(), strings.ReplaceAll(currCase.msg, "than", "than or equal to"))
			Contains(t, mock.helpers, pkg+".LessOrEqual")
		})

		t.Run("should be stricly less", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			True(t, LessOrEqual(mock, currCase.less, currCase.greater),
				"expected %v to be less than or equal to %v",
				currCase.less, currCase.greater,
			)
		})

		for currCase := range compareEqualCases() {
			t.Run("equal values should be less or equal", func(t *testing.T) {
				t.Parallel()

				mock := new(mockT)

				True(t, GreaterOrEqual(mock, currCase.less, currCase.greater),
					"expected (equal) %v to be less than or equal to %v",
					currCase.less, currCase.greater,
				)
			})
		}
	}
}

func TestCompareGreaterT(t *testing.T) {
	t.Parallel()

	for tc := range greaterTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestCompareGreaterOrEqualT(t *testing.T) {
	t.Parallel()

	for tc := range greaterOrEqualTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestCompareLessT(t *testing.T) {
	t.Parallel()

	for tc := range lessTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestCompareLessOrEqualT(t *testing.T) {
	t.Parallel()

	for tc := range lessOrEqualTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestComparePositiveT(t *testing.T) {
	t.Parallel()

	for tc := range positiveTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestCompareNegativeT(t *testing.T) {
	t.Parallel()

	for tc := range negativeTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestComparePositive(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Positive(mock, 1) {
		t.Error("Positive should return true")
	}

	if !Positive(mock, 1.23) {
		t.Error("Positive should return true")
	}

	if Positive(mock, -1) {
		t.Error("Positive should return false")
	}

	if Positive(mock, -1.23) {
		t.Error("Positive should return false")
	}

	// Check error report
	for currCase := range comparePositiveCases() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, Positive(out, currCase.e))
		Contains(t, out.buf.String(), currCase.msg)
		Contains(t, out.helpers, pkg+".Positive")
	}
}

func TestCompareNegative(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Negative(mock, -1) {
		t.Error("Negative should return true")
	}

	if !Negative(mock, -1.23) {
		t.Error("Negative should return true")
	}

	if Negative(mock, 1) {
		t.Error("Negative should return false")
	}

	if Negative(mock, 1.23) {
		t.Error("Negative should return false")
	}

	// Check error report
	for currCase := range compareNegativeCases() {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		False(t, Negative(out, currCase.e))
		Contains(t, out.buf.String(), currCase.msg)
		Contains(t, out.helpers, pkg+".Negative")
	}
}

func TestCompareMsgAndArgsForwarding(t *testing.T) {
	msgAndArgs := []any{"format %s %x", "this", 0xc001}
	const expectedOutput = "format this c001\n"

	funcs := []func(t T){
		func(t T) { Greater(t, 1, 2, msgAndArgs...) },
		func(t T) { GreaterOrEqual(t, 1, 2, msgAndArgs...) },
		func(t T) { Less(t, 2, 1, msgAndArgs...) },
		func(t T) { LessOrEqual(t, 2, 1, msgAndArgs...) },
		func(t T) { Positive(t, 0, msgAndArgs...) },
		func(t T) { Negative(t, 0, msgAndArgs...) },
	}

	for _, f := range funcs {
		out := &outputT{buf: bytes.NewBuffer(nil)}
		f(out)
		Contains(t, out.buf.String(), expectedOutput)
	}
}

// genericTestCase wraps a test function with its name for table-driven tests of generic functions.
type genericTestCase struct {
	name string
	test func(*testing.T)
}

// greaterTCases returns test cases for GreaterT with various Ordered types.
func greaterTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"int", testGreaterT[int](2, 1, 1, 2)},
		{"int8", testGreaterT[int8](2, 1, 1, 2)},
		{"int16", testGreaterT[int16](2, 1, 1, 2)},
		{"int32", testGreaterT[int32](2, 1, 1, 2)},
		{"int64", testGreaterT[int64](2, 1, 1, 2)},
		{"uint", testGreaterT[uint](2, 1, 1, 2)},
		{"uint8", testGreaterT[uint8](2, 1, 1, 2)},
		{"uint16", testGreaterT[uint16](2, 1, 1, 2)},
		{"uint32", testGreaterT[uint32](2, 1, 1, 2)},
		{"uint64", testGreaterT[uint64](2, 1, 1, 2)},
		{"float32", testGreaterT[float32](2.5, 1.5, 1.5, 2.5)},
		{"float64", testGreaterT[float64](2.5, 1.5, 1.5, 2.5)},
		{"string", testGreaterT[string]("b", "a", "a", "b")},
		{"uintptr", testGreaterT[uintptr](2, 1, 1, 2)},
		{"time.Time", testGreaterTTime()},
		{"[]byte", testGreaterTBytes()},
		{"custom int type", testGreaterTCustomInt()},
	})
}

// greaterOrEqualTCases returns test cases for GreaterOrEqualT with various Ordered types.
func greaterOrEqualTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"int", testGreaterOrEqualT[int](2, 1, 1, 1, 0, 1)},
		{"int8", testGreaterOrEqualT[int8](2, 1, 1, 1, 0, 1)},
		{"int16", testGreaterOrEqualT[int16](2, 1, 1, 1, 0, 1)},
		{"int32", testGreaterOrEqualT[int32](2, 1, 1, 1, 0, 1)},
		{"int64", testGreaterOrEqualT[int64](2, 1, 1, 1, 0, 1)},
		{"uint", testGreaterOrEqualT[uint](2, 1, 1, 1, 0, 1)},
		{"uint8", testGreaterOrEqualT[uint8](2, 1, 1, 1, 0, 1)},
		{"uint16", testGreaterOrEqualT[uint16](2, 1, 1, 1, 0, 1)},
		{"uint32", testGreaterOrEqualT[uint32](2, 1, 1, 1, 0, 1)},
		{"uint64", testGreaterOrEqualT[uint64](2, 1, 1, 1, 0, 1)},
		{"float32", testGreaterOrEqualT[float32](2.5, 1.5, 1.5, 1.5, 0.5, 1.5)},
		{"float64", testGreaterOrEqualT[float64](2.5, 1.5, 1.5, 1.5, 0.5, 1.5)},
		{"string", testGreaterOrEqualT[string]("b", "a", "a", "a", "a", "b")},
		{"uintptr", testGreaterOrEqualT[uintptr](2, 1, 1, 1, 0, 1)},
		{"time.Time", testGreaterOrEqualTTime()},
		{"[]byte", testGreaterOrEqualTBytes()},
	})
}

// lessTCases returns test cases for LessT with various Ordered types.
func lessTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"int", testLessT[int](1, 2, 2, 1)},
		{"int8", testLessT[int8](1, 2, 2, 1)},
		{"int16", testLessT[int16](1, 2, 2, 1)},
		{"int32", testLessT[int32](1, 2, 2, 1)},
		{"int64", testLessT[int64](1, 2, 2, 1)},
		{"uint", testLessT[uint](1, 2, 2, 1)},
		{"uint8", testLessT[uint8](1, 2, 2, 1)},
		{"uint16", testLessT[uint16](1, 2, 2, 1)},
		{"uint32", testLessT[uint32](1, 2, 2, 1)},
		{"uint64", testLessT[uint64](1, 2, 2, 1)},
		{"float32", testLessT[float32](1.5, 2.5, 2.5, 1.5)},
		{"float64", testLessT[float64](1.5, 2.5, 2.5, 1.5)},
		{"string", testLessT[string]("a", "b", "b", "a")},
		{"uintptr", testLessT[uintptr](1, 2, 2, 1)},
		{"time.Time", testLessTTime()},
		{"[]byte", testLessTBytes()},
	})
}

// lessOrEqualTCases returns test cases for LessOrEqualT with various Ordered types.
func lessOrEqualTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"int", testLessOrEqualT[int](1, 2, 1, 1, 2, 1)},
		{"int8", testLessOrEqualT[int8](1, 2, 1, 1, 2, 1)},
		{"int16", testLessOrEqualT[int16](1, 2, 1, 1, 2, 1)},
		{"int32", testLessOrEqualT[int32](1, 2, 1, 1, 2, 1)},
		{"int64", testLessOrEqualT[int64](1, 2, 1, 1, 2, 1)},
		{"uint", testLessOrEqualT[uint](1, 2, 1, 1, 2, 1)},
		{"uint8", testLessOrEqualT[uint8](1, 2, 1, 1, 2, 1)},
		{"uint16", testLessOrEqualT[uint16](1, 2, 1, 1, 2, 1)},
		{"uint32", testLessOrEqualT[uint32](1, 2, 1, 1, 2, 1)},
		{"uint64", testLessOrEqualT[uint64](1, 2, 1, 1, 2, 1)},
		{"float32", testLessOrEqualT[float32](1.5, 2.5, 1.5, 1.5, 2.5, 1.5)},
		{"float64", testLessOrEqualT[float64](1.5, 2.5, 1.5, 1.5, 2.5, 1.5)},
		{"string", testLessOrEqualT[string]("a", "b", "a", "a", "b", "a")},
		{"uintptr", testLessOrEqualT[uintptr](1, 2, 1, 1, 2, 1)},
		{"time.Time", testLessOrEqualTTime()},
		{"[]byte", testLessOrEqualTBytes()},
	})
}

// positiveTCases returns test cases for PositiveT with various SignedNumeric types.
func positiveTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"int", testPositiveT[int](1, -1)},
		{"int8", testPositiveT[int8](1, -1)},
		{"int16", testPositiveT[int16](1, -1)},
		{"int32", testPositiveT[int32](1, -1)},
		{"int64", testPositiveT[int64](1, -1)},
		{"float32", testPositiveT[float32](1.5, -1.5)},
		{"float64", testPositiveT[float64](1.5, -1.5)},
		{"zero is not positive", testPositiveTZero()},
	})
}

// negativeTCases returns test cases for NegativeT with various SignedNumeric types.
func negativeTCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		{"int", testNegativeT[int](-1, 1)},
		{"int8", testNegativeT[int8](-1, 1)},
		{"int16", testNegativeT[int16](-1, 1)},
		{"int32", testNegativeT[int32](-1, 1)},
		{"int64", testNegativeT[int64](-1, 1)},
		{"float32", testNegativeT[float32](-1.5, 1.5)},
		{"float64", testNegativeT[float64](-1.5, 1.5)},
		{"zero is not negative", testNegativeTZero()},
	})
}

// Test helper functions for generic comparison assertions

func testGreaterT[V Ordered](successE1, successE2, failE1, failE2 V) func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, GreaterT(mock, successE1, successE2))
		False(t, GreaterT(mock, failE1, failE2))
		False(t, GreaterT(mock, successE1, successE1)) // equal values
	}
}

func testGreaterTTime() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		t0 := time.Now()
		t1 := t0.Add(-time.Second)

		True(t, GreaterT(mock, t0, t1))
		False(t, GreaterT(mock, t1, t0))
		False(t, GreaterT(mock, t0, t0))
	}
}

func testGreaterTBytes() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, GreaterT(mock, []byte{2}, []byte{1}))
		False(t, GreaterT(mock, []byte{1}, []byte{2}))
		False(t, GreaterT(mock, []byte{1}, []byte{1}))
	}
}

func testGreaterTCustomInt() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		type MyInt int
		True(t, GreaterT(mock, MyInt(2), MyInt(1)))
		False(t, GreaterT(mock, MyInt(1), MyInt(2)))
	}
}

func testGreaterOrEqualT[V Ordered](gtE1, gtE2, eqE1, eqE2, failE1, failE2 V) func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, GreaterOrEqualT(mock, gtE1, gtE2))      // greater
		True(t, GreaterOrEqualT(mock, eqE1, eqE2))      // equal
		False(t, GreaterOrEqualT(mock, failE1, failE2)) // less
	}
}

func testGreaterOrEqualTTime() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		t0 := time.Now()
		t1 := t0.Add(-time.Second)

		True(t, GreaterOrEqualT(mock, t0, t1))  // greater
		True(t, GreaterOrEqualT(mock, t0, t0))  // equal
		False(t, GreaterOrEqualT(mock, t1, t0)) // less
	}
}

func testGreaterOrEqualTBytes() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, GreaterOrEqualT(mock, []byte{2}, []byte{1}))
		True(t, GreaterOrEqualT(mock, []byte{1}, []byte{1}))
		False(t, GreaterOrEqualT(mock, []byte{1}, []byte{2}))
	}
}

func testLessT[V Ordered](successE1, successE2, failE1, failE2 V) func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, LessT(mock, successE1, successE2))
		False(t, LessT(mock, failE1, failE2))
		False(t, LessT(mock, successE1, successE1)) // equal values
	}
}

func testLessTTime() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		t0 := time.Now()
		t1 := t0.Add(time.Second)

		True(t, LessT(mock, t0, t1))
		False(t, LessT(mock, t1, t0))
		False(t, LessT(mock, t0, t0))
	}
}

func testLessTBytes() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, LessT(mock, []byte{1}, []byte{2}))
		False(t, LessT(mock, []byte{2}, []byte{1}))
		False(t, LessT(mock, []byte{1}, []byte{1}))
	}
}

func testLessOrEqualT[V Ordered](ltE1, ltE2, eqE1, eqE2, failE1, failE2 V) func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, LessOrEqualT(mock, ltE1, ltE2))      // less
		True(t, LessOrEqualT(mock, eqE1, eqE2))      // equal
		False(t, LessOrEqualT(mock, failE1, failE2)) // greater
	}
}

func testLessOrEqualTTime() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		t0 := time.Now()
		t1 := t0.Add(time.Second)

		True(t, LessOrEqualT(mock, t0, t1))  // less
		True(t, LessOrEqualT(mock, t0, t0))  // equal
		False(t, LessOrEqualT(mock, t1, t0)) // greater
	}
}

func testLessOrEqualTBytes() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, LessOrEqualT(mock, []byte{1}, []byte{2}))
		True(t, LessOrEqualT(mock, []byte{1}, []byte{1}))
		False(t, LessOrEqualT(mock, []byte{2}, []byte{1}))
	}
}

func testPositiveT[V SignedNumeric](positive, negative V) func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, PositiveT(mock, positive))
		False(t, PositiveT(mock, negative))
	}
}

func testPositiveTZero() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		False(t, PositiveT(mock, 0))
		False(t, PositiveT(mock, 0.0))
	}
}

func testNegativeT[V SignedNumeric](negative, positive V) func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		True(t, NegativeT(mock, negative))
		False(t, NegativeT(mock, positive))
	}
}

func testNegativeTZero() func(*testing.T) {
	//nolint:thelper // linter false positive: this is not a helper
	return func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		False(t, NegativeT(mock, 0))
		False(t, NegativeT(mock, 0.0))
	}
}

// callerName gives the function name (qualified with a package path)
// for the caller after skip frames (where 0 means the current function).
func callerName(skip int) string {
	// Make room for the skip PC.
	var pc [1]uintptr
	n := runtime.Callers(skip+2, pc[:]) // skip + runtime.Callers + callerName
	if n == 0 {
		panic("testing: zero callers found")
	}
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return frame.Function
}

type compareFixture struct {
	less    any
	greater any
	msg     string
}

func compareStrictlyGreaterCases() iter.Seq[compareFixture] {
	const genMsg = `"1" is not greater than "2"`

	type redefinedUintptr uintptr

	return slices.Values(
		[]compareFixture{
			{less: "a", greater: "b", msg: `"a" is not greater than "b"`},
			{less: int(1), greater: int(2), msg: genMsg},
			{less: int8(1), greater: int8(2), msg: genMsg},
			{less: int16(1), greater: int16(2), msg: genMsg},
			{less: int32(1), greater: int32(2), msg: genMsg},
			{less: int64(1), greater: int64(2), msg: genMsg},
			{less: uint8(1), greater: uint8(2), msg: genMsg},
			{less: uint16(1), greater: uint16(2), msg: genMsg},
			{less: uint32(1), greater: uint32(2), msg: genMsg},
			{less: uint64(1), greater: uint64(2), msg: genMsg},
			{less: float32(1.23), greater: float32(2.34), msg: `"1.23" is not greater than "2.34"`},
			{less: float64(1.23), greater: float64(2.34), msg: `"1.23" is not greater than "2.34"`},
			{less: uintptr(1), greater: uintptr(2), msg: genMsg},
			{less: uintptr(9), greater: uintptr(10), msg: `"9" is not greater than "10"`},
			{less: redefinedUintptr(9), greater: redefinedUintptr(10), msg: `"9" is not greater than "10"`},
			{less: time.Time{}, greater: time.Time{}.Add(time.Hour), msg: `"0001-01-01 00:00:00 +0000 UTC" is not greater than "0001-01-01 01:00:00 +0000 UTC"`},
			{less: []byte{1, 1}, greater: []byte{1, 2}, msg: `"[1 1]" is not greater than "[1 2]"`},
		},
	)
}

func compareStrictlyLessCases() iter.Seq[compareFixture] {
	const genMsg = `"2" is not less than "1"`

	return slices.Values(
		[]compareFixture{
			{less: "a", greater: "b", msg: `"b" is not less than "a"`},
			{less: int(1), greater: int(2), msg: genMsg},
			{less: int8(1), greater: int8(2), msg: genMsg},
			{less: int16(1), greater: int16(2), msg: genMsg},
			{less: int32(1), greater: int32(2), msg: genMsg},
			{less: int64(1), greater: int64(2), msg: genMsg},
			{less: uint8(1), greater: uint8(2), msg: genMsg},
			{less: uint16(1), greater: uint16(2), msg: genMsg},
			{less: uint32(1), greater: uint32(2), msg: genMsg},
			{less: uint64(1), greater: uint64(2), msg: genMsg},
			{less: float32(1.23), greater: float32(2.34), msg: `"2.34" is not less than "1.23"`},
			{less: float64(1.23), greater: float64(2.34), msg: `"2.34" is not less than "1.23"`},
			{less: uintptr(1), greater: uintptr(2), msg: genMsg},
			{less: time.Time{}, greater: time.Time{}.Add(time.Hour), msg: `"0001-01-01 01:00:00 +0000 UTC" is not less than "0001-01-01 00:00:00 +0000 UTC"`},
			{less: []byte{1, 1}, greater: []byte{1, 2}, msg: `"[1 2]" is not less than "[1 1]"`},
		},
	)
}

func compareEqualCases() iter.Seq[compareFixture] {
	// This iterator produces equal-values to check edge cases with strict comparisons.
	// The message cannot be used for error message checks.
	return func(yield func(compareFixture) bool) {
		for greater := range compareStrictlyGreaterCases() {
			greater.msg = ""
			equal1 := greater
			equal1.less = equal1.greater
			if !yield(equal1) {
				return
			}

			equal2 := greater
			equal2.greater = equal2.less
			if !yield(equal2) {
				return
			}
		}

		for less := range compareStrictlyLessCases() {
			less.msg = ""
			equal1 := less
			equal1.less = equal1.greater
			if !yield(equal1) {
				return
			}

			equal2 := less
			equal2.greater = equal2.less
			if !yield(equal2) {
				return
			}
		}
	}
}

type compareTestCase struct {
	e   any
	msg string
}

type comparePositiveCase = compareTestCase

type compareNegativeCase = compareTestCase

func comparePositiveCases() iter.Seq[comparePositiveCase] {
	const genMsg = `"-1" is not positive`

	return slices.Values([]comparePositiveCase{
		{e: int(-1), msg: genMsg},
		{e: int8(-1), msg: genMsg},
		{e: int16(-1), msg: genMsg},
		{e: int32(-1), msg: genMsg},
		{e: int64(-1), msg: genMsg},
		{e: float32(-1.23), msg: `"-1.23" is not positive`},
		{e: float64(-1.23), msg: `"-1.23" is not positive`},
	})
}

func compareNegativeCases() iter.Seq[compareNegativeCase] {
	const genMsg = `"1" is not negative`

	return slices.Values([]compareNegativeCase{
		{e: int(1), msg: genMsg},
		{e: int8(1), msg: genMsg},
		{e: int16(1), msg: genMsg},
		{e: int32(1), msg: genMsg},
		{e: int64(1), msg: genMsg},
		{e: float32(1.23), msg: `"1.23" is not negative`},
		{e: float64(1.23), msg: `"1.23" is not negative`},
	})
}
