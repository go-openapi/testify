// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"iter"
	"regexp"
	"slices"
	"strings"
	"testing"
)

func TestEqualErrorMessages(t *testing.T) {
	t.Parallel()

	t.Run("should render when slice too long to print", testTooLongToPrint())
	t.Run("error message should match expression", func(t *testing.T) {
		// checking error messsages on Equal with a regexp. The object of the test is Equal, not Regexp
		for tc := range stringEqualFormattingCases() {
			t.Run(tc.name, func(t *testing.T) {
				mock := &bufferT{}

				isEqual := Equal(mock, tc.equalWant, tc.equalGot, tc.msgAndArgs...)
				if isEqual {
					t.Errorf("expected %q to be different than %q", tc.equalGot, tc.equalWant)

					return
				}

				rex := regexp.MustCompile(tc.want)
				match := rex.MatchString(mock.buf.String())
				if !match {
					t.Errorf("expected message to match %q, but got:\n%s", tc.want, mock.buf.String())
				}
			})
		}
	})
}

// Test NotEqualValues.
func TestEqualValuesAndNotEqualValues(t *testing.T) {
	t.Parallel()

	for tc := range equalValuesCases() {
		mock := new(testing.T)

		t.Run(tc.name, func(t *testing.T) {
			res := NotEqualValues(mock, tc.expected, tc.actual)
			if res != tc.notEqualValue {
				t.Errorf("NotEqualValues(%#v, %#v) should return %t", tc.expected, tc.actual, tc.notEqualValue)
			}
		})

		// Test EqualValues (inverse of NotEqualValues)
		t.Run(tc.name, func(t *testing.T) {
			res := EqualValues(mock, tc.expected, tc.actual)
			if res != tc.equalValue {
				t.Errorf("EqualValues(%#v, %#v) should return %t", tc.expected, tc.actual, tc.equalValue)
			}
		})
	}
}

// Test EqualExportedValues.
func TestEqualExportedValues(t *testing.T) {
	t.Parallel()

	for tc := range objectEqualExportedValuesCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			mockT := new(mockT)

			actual := EqualExportedValues(mockT, tc.expected, tc.actual)
			if actual != tc.expectedEqual {
				t.Errorf("Expected EqualExportedValues to be %t, but was %t", tc.expectedEqual, actual)
			}

			if tc.expectedFailMsg == "" {
				// skip error message check
				return
			}

			actualFail := mockT.errorString()
			if !strings.Contains(actualFail, tc.expectedFailMsg) {
				t.Errorf("Contains failure should include %q but was %q", tc.expectedFailMsg, actualFail)
			}
		})
	}
}

// Deep equality tests (Equal, EqualT, NotEqual, NotEqualT, Exactly).
func TestEqualDeepEqual(t *testing.T) {
	t.Parallel()

	for tc := range unifiedEqualityCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			expected, actual := tc.makeValues()

			if !tc.reflectionOnly {
				t.Run("with Equal", testEqualityAssertion(tc, equalKind, Equal, expected, actual))
				t.Run("with NotEqual", testEqualityAssertion(tc, notEqualKind, NotEqual, expected, actual))
				t.Run("with Exactly", testEqualityAssertion(tc, exactlyKind, Exactly, expected, actual))

				// Generic variants - type dispatch
				t.Run("with EqualT", testEqualityAssertionT(tc, equalTKind, expected, actual))
				t.Run("with NotEqualT", testEqualityAssertionT(tc, notEqualTKind, expected, actual))
			} else {
				// Reflection-only cases
				t.Run("with Equal (reflection)", testEqualityAssertion(tc, equalKind, Equal, expected, actual))
				t.Run("with NotEqual (reflection)", testEqualityAssertion(tc, notEqualKind, NotEqual, expected, actual))
				t.Run("with Exactly (reflection)", testEqualityAssertion(tc, exactlyKind, Exactly, expected, actual))
			}
		})
	}
}

type equalityRelationship int

const (
	eqBothNil equalityRelationship = iota
	eqOneNil
	eqSameIdentity
	eqEqualValueComparable
	eqEqualValueNonComparable
	eqDifferentValueSameType
	eqDifferentType
	eqFunction
)

func testEqualityAssertion(tc equalityTestCase, kind equalityAssertionKind, equalityAssertion func(T, any, any, ...any) bool, expected, actual any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := equalityAssertion(mock, expected, actual)
		shouldPass := expectedStatusForEqualityAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

//nolint:gocognit,gocyclo,cyclop // no other way here than a big type switch
func testEqualityAssertionT(tc equalityTestCase, kind equalityAssertionKind, expected, actual any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		stop := func(expected, actual any) {
			t.Fatalf("test case error: expected=%s, actual=%T", expected, actual)
		}

		// Type switch with safety check
		//
		// Add more (comparable) types when new test cases require it.
		var result bool
		switch exp := expected.(type) {
		case int:
			act, ok := actual.(int)
			if !ok {
				stop("int", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case int8:
			act, ok := actual.(int8)
			if !ok {
				stop("int8", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case int16:
			act, ok := actual.(int16)
			if !ok {
				stop("int16", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case int32:
			act, ok := actual.(int32)
			if !ok {
				stop("int32", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case int64:
			act, ok := actual.(int64)
			if !ok {
				stop("int32", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case uint:
			act, ok := actual.(uint)
			if !ok {
				stop("uint", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case uint8:
			act, ok := actual.(uint8)
			if !ok {
				stop("uint8", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case uint16:
			act, ok := actual.(uint16)
			if !ok {
				stop("uint16", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case uint32:
			act, ok := actual.(uint32)
			if !ok {
				stop("uint32", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case uint64:
			act, ok := actual.(uint64)
			if !ok {
				stop("uint64", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case string:
			act, ok := actual.(string)
			if !ok {
				stop("string", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case bool:
			act, ok := actual.(bool)
			if !ok {
				stop("bool", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case float32:
			act, ok := actual.(float32)
			if !ok {
				stop("float32", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case float64:
			act, ok := actual.(float64)
			if !ok {
				stop("float64", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case *int:
			act, ok := actual.(*int)
			if !ok {
				stop("*int", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case *string:
			act, ok := actual.(*string)
			if !ok {
				stop("*string", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case myType:
			act, ok := actual.(myType)
			if !ok {
				stop("myType", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case struct{}:
			act, ok := actual.(struct{})
			if !ok {
				stop("struct{}", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		case *struct{}:
			act, ok := actual.(*struct{})
			if !ok {
				stop("*struct{}", actual)
			}
			result = testEqualityGenericAssertion(mock, kind, exp, act)
		default:
			t.Fatalf("unsupported type: %T", expected)
		}

		shouldPass := expectedStatusForEqualityAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

type equalityTestCase struct {
	name           string
	makeValues     func() (expected, actual any)
	relationship   equalityRelationship
	reflectionOnly bool
}

type (
	structWithUnexportedMapWithArrayKey struct {
		m any
	}
	s struct {
		f map[[1]byte]int
	}

	myType string
	myMap  map[string]any
)

func unifiedEqualityCases() iter.Seq[equalityTestCase] {
	const hello = "hello"
	s1 := struct{}{}
	p1 := &s1
	p2 := &s1
	f1 := func() bool { return true }

	return slices.Values([]equalityTestCase{
		// Both nil
		{"both-nil/ptr", func() (any, any) { return (*int)(nil), (*int)(nil) }, eqBothNil, false},
		{"both-nil/interface", func() (any, any) { return (any)(nil), (any)(nil) }, eqBothNil, true},

		// One nil (reflection only - type mismatch)
		{"one-nil/first", func() (any, any) { v := 42; return nil, &v }, eqOneNil, true},
		{"one-nil/second", func() (any, any) { v := 42; return &v, nil }, eqOneNil, true},
		{"one-nil/bytes", func() (any, any) { return nil, make([]byte, 0) }, eqOneNil, true},
		{"one-nil/struct", func() (any, any) { return nil, new(AssertionTesterConformingObject) }, eqOneNil, true},

		// Same identity (pointers to same object)
		{"same-identity/int-ptr", func() (any, any) { v := 42; return &v, &v }, eqSameIdentity, false},
		{"same-identity/string-ptr", func() (any, any) { s := hello; return &s, &s }, eqSameIdentity, false},
		{
			"same-identity/pointer-to-struct", func() (any, any) {
				return p1, p2
			},
			eqSameIdentity, false,
		},

		// Equal value, comparable (core types - start with 5)
		{"equal-comparable/int", func() (any, any) { return 42, 42 }, eqEqualValueComparable, false},
		{"equal-comparable/string", func() (any, any) { return hello, hello }, eqEqualValueComparable, false},
		{"equal-comparable/bool-true", func() (any, any) { return true, true }, eqEqualValueComparable, false},
		{"equal-comparable/bool-false", func() (any, any) { return false, false }, eqEqualValueComparable, false},
		{"equal-comparable/float64", func() (any, any) { return 3.14, 3.14 }, eqEqualValueComparable, false},
		{"equal-comparable/float32", func() (any, any) { return float32(3.14), float32(3.14) }, eqEqualValueComparable, false},
		{"equal-comparable/int8", func() (any, any) { return int8(10), int8(10) }, eqEqualValueComparable, false},
		{"equal-comparable/int16", func() (any, any) { return int16(100), int16(100) }, eqEqualValueComparable, false},
		{"equal-comparable/int32", func() (any, any) { return int32(1000), int32(1000) }, eqEqualValueComparable, false},
		{"equal-comparable/int64", func() (any, any) { return int64(10000), int64(10000) }, eqEqualValueComparable, false},
		{"equal-comparable/uint", func() (any, any) { return uint(42), uint(42) }, eqEqualValueComparable, false},
		{"equal-comparable/uint8", func() (any, any) { return uint8(10), uint8(10) }, eqEqualValueComparable, false},
		{"equal-comparable/uint16", func() (any, any) { return uint16(100), uint16(100) }, eqEqualValueComparable, false},
		{"equal-comparable/uint32", func() (any, any) { return uint32(1000), uint32(1000) }, eqEqualValueComparable, false},
		{"equal-comparable/uint64", func() (any, any) { return uint64(10000), uint64(10000) }, eqEqualValueComparable, false},
		{"equal-comparable/~string", func() (any, any) { return myType("1"), myType("1") }, eqEqualValueComparable, false},
		{
			"equal-comparable/anonymous-struct", func() (any, any) {
				return struct{}{}, struct{}{}
			}, eqEqualValueComparable, false,
		},
		{
			"equal-comparable/pointer-to-anonymous-struct", func() (any, any) {
				return &struct{}{}, &struct{}{} // this a special case in go, as the pointer to this empty type is not allocated: values are equal
			}, eqEqualValueComparable, false,
		},

		// Equal value, non-comparable (reflection only)
		{"equal-non-comparable/slice", func() (any, any) { return []int{1, 2, 3}, []int{1, 2, 3} }, eqEqualValueNonComparable, true},
		{"equal-non-comparable/struct-ptr", func() (any, any) { return &struct{}{}, &struct{}{} }, eqEqualValueNonComparable, true},
		{"equal-non-comparable/bytes", func() (any, any) { return []byte(hello), []byte(hello) }, eqEqualValueNonComparable, true},
		{"equal-non-comparable/map", func() (any, any) { return myMap{"bar": 1}, myMap{"bar": 1} }, eqEqualValueNonComparable, true},
		{
			"equal-non-comparable/bytes-zero-same-len", func() (any, any) {
				return make([]byte, 2), make([]byte, 2)
			}, eqEqualValueNonComparable, true,
		},
		{
			"equal-non-comparable/bytes-zero-same-len-diff-cap", func() (any, any) {
				return make([]byte, 2), make([]byte, 2, 3)
			}, eqEqualValueNonComparable, true,
		},
		{
			"equal-non-comparable/bytes-zero-same-len-diff-cap", func() (any, any) {
				return new(AssertionTesterConformingObject), new(AssertionTesterConformingObject)
			}, eqEqualValueNonComparable, true,
		},
		{
			"equal-non-comparable/map-unexported-struct", func() (any, any) {
				return structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{{1}: nil, {2}: nil}},
					structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{{2}: nil, {1}: nil}}
			},
			eqEqualValueNonComparable, true,
		},
		{
			"equal-non-comparable/map-unexported-struct-non-nil", func() (any, any) {
				return structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{{1}: {}, {2}: nil}},
					structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{{1}: {}, {2}: nil}}
			},
			eqEqualValueNonComparable, true,
		},

		// Different value, same type
		{"diff-value/int", func() (any, any) { return 42, 43 }, eqDifferentValueSameType, false},
		{"diff-value/string", func() (any, any) { return hello, "world" }, eqDifferentValueSameType, false},
		{"diff-value/bool", func() (any, any) { return true, false }, eqDifferentValueSameType, false},
		{"diff-value/float64", func() (any, any) { return 3.14, 2.71 }, eqDifferentValueSameType, false},
		{"diff-value/float32", func() (any, any) { return float32(3.15), float32(3.14) }, eqDifferentValueSameType, false},
		{"diff-value/int8", func() (any, any) { return int8(10), int8(11) }, eqDifferentValueSameType, false},
		{"diff-value/int16", func() (any, any) { return int16(110), int16(100) }, eqDifferentValueSameType, false},
		{"diff-value/int32", func() (any, any) { return int32(1003), int32(1000) }, eqDifferentValueSameType, false},
		{"diff-value/int64", func() (any, any) { return int64(10400), int64(10000) }, eqDifferentValueSameType, false},
		{"diff-value/uint", func() (any, any) { return uint(43), uint(42) }, eqDifferentValueSameType, false},
		{"diff-value/uint8", func() (any, any) { return uint8(10), uint8(11) }, eqDifferentValueSameType, false},
		{"diff-value/uint16", func() (any, any) { return uint16(101), uint16(100) }, eqDifferentValueSameType, false},
		{"diff-value/uint32", func() (any, any) { return uint32(1040), uint32(1000) }, eqDifferentValueSameType, false},
		{"diff-value/uint64", func() (any, any) { return uint64(10000), uint64(14000) }, eqDifferentValueSameType, false},
		{"diff-value/~string", func() (any, any) { return myType("1"), myType("2") }, eqDifferentValueSameType, false},
		{"diff-value/slice", func() (any, any) { return []int{1, 2}, []int{1, 3} }, eqDifferentValueSameType, true},
		{"diff-value/map", func() (any, any) { return myMap{"bar": 1}, myMap{"bar": 2} }, eqDifferentValueSameType, true},
		{
			"diff-value/map-unexported-struct", func() (any, any) {
				return structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{{1}: nil, {2}: nil}},
					structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{}}
			},
			eqDifferentValueSameType, true,
		},
		{
			"diff-value/map-unexported-struct-non-nil", func() (any, any) {
				return structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{{1}: {}, {2}: nil}},
					structWithUnexportedMapWithArrayKey{map[[1]byte]*struct{}{{1}: nil, {2}: {}}}
			},
			eqDifferentValueSameType, true,
		},
		{
			"diff-value/func", func() (any, any) {
				return func() int { return 23 }, func() int { return 24 }
			},
			eqFunction, true,
		},
		{
			"same-value/func", func() (any, any) {
				return f1, f1
			},
			eqFunction, true,
		},

		// Different type (reflection only - can't use with generics)
		{"diff-type/int-uint", func() (any, any) { return 42, uint(42) }, eqDifferentType, true},
		{"diff-type/int-int64", func() (any, any) { return 42, int64(42) }, eqDifferentType, true},
		{"diff-type/int-float64", func() (any, any) { return 10, 10.0 }, eqDifferentType, true},
		{"diff-type/float32-float64", func() (any, any) { return float32(10), float64(10) }, eqDifferentType, true},
		{"diff-value/edge-case-map", func() (any, any) { // this case used to panic
			return s{
					f: map[[1]byte]int{
						{0x1}: 0,
						{0x2}: 0,
					},
				},
				s{}
		}, eqDifferentValueSameType, true},
	})
}

type equalityAssertionKind int

const (
	equalKind equalityAssertionKind = iota
	equalTKind
	notEqualKind
	notEqualTKind
	exactlyKind
)

func expectedStatusForEqualityAssertion(kind equalityAssertionKind, relationship equalityRelationship) bool {
	positive := kind == equalKind || kind == equalTKind || kind == exactlyKind

	switch relationship {
	case eqFunction:
		// A special validation is carried out to reject function types with an error
		return false
	case eqBothNil, eqSameIdentity, eqEqualValueComparable, eqEqualValueNonComparable:
		return positive
	case eqOneNil, eqDifferentValueSameType:
		return !positive
	case eqDifferentType:
		// Exactly requires exact type match (fails on different types)
		// Equal uses reflection (can handle different types)
		// EqualT requires same type (won't compile with different types)
		if kind == exactlyKind {
			return false
		}
		// Equal might handle some type coercion, but generally fails
		return !positive
	default:
		panic(fmt.Errorf("test case configuration error: invalid equalityRelationship: %d", relationship))
	}
}

func testEqualityGenericAssertion[V comparable](mock T, kind equalityAssertionKind, expected, actual V) bool {
	switch kind {
	case equalTKind:
		return EqualT(mock, expected, actual)
	case notEqualTKind:
		return NotEqualT(mock, expected, actual)
	default:
		panic(fmt.Errorf("test case configuration error: invalid equalityAssertionKind for generic: %d", kind))
	}
}

func testTooLongToPrint() func(*testing.T) {
	const (
		expected = `&[]int{0, 0, 0,`
		message  = `
	Error Trace:	
	Error:      	Should not be: []int{0, 0, 0,`
		trailer = `<... truncated>`
	)

	return func(t *testing.T) {
		t.Run("with Same", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)

			longSlice := make([]int, 1_000_000)
			result := Same(mock, &[]int{}, &longSlice)
			if result {
				t.Errorf("expected Same to fail")
				return
			}

			if !strings.Contains(mock.errorString(), expected) {
				t.Errorf("expected message to contain %q but got: %q", expected, mock.errorString())
			}
		})

		t.Run("with NotSame", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)

			longSlice := make([]int, 1_000_000)
			result := NotSame(mock, &longSlice, &longSlice)
			if result {
				t.Errorf("expected NotSame to fail")
				return
			}

			if !strings.Contains(mock.errorString(), expected) {
				t.Errorf("expected message to contain %q but got: %q", expected, mock.errorString())
			}
		})

		t.Run("with NotEqual", func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)

			longSlice := make([]int, 1_000_000)
			result := NotEqual(mock, longSlice, longSlice)
			if result {
				t.Errorf("expected NotEqual to fail")
				return
			}

			if !strings.Contains(mock.errorString(), message) {
				t.Errorf("expected message to contain %q but got: %q", message, mock.errorString())
			}

			if !strings.Contains(mock.errorString(), trailer) {
				t.Errorf("expected message to contain %q but got: %q", trailer, mock.errorString())
			}
		})

		t.Run("with NotEqualValues", func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)

			longSlice := make([]int, 1_000_000)
			result := NotEqualValues(mock, longSlice, longSlice)
			if result {
				t.Errorf("expected NotEqualValues to fail")
				return
			}

			if !strings.Contains(mock.errorString(), message) {
				t.Errorf("expected message to contain %q but got: %q", message, mock.errorString())
			}
			const trailer = `<... truncated>`
			if !strings.Contains(mock.errorString(), trailer) {
				t.Errorf("expected message to contain %q but got: %q", trailer, mock.errorString())
			}
		})
	}
}

type equalStringCase struct {
	name       string
	equalWant  string
	equalGot   string
	msgAndArgs []any
	want       string
}

func stringEqualFormattingCases() iter.Seq[equalStringCase] {
	return slices.Values([]equalStringCase{
		{
			name:      "multiline diff message",
			equalWant: "hi, \nmy name is",
			equalGot:  "what,\nmy name is",
			want: "\t[a-z]+.go:\\d+: \n" +
				"\t+Error Trace:\t\n+" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"hi, \\\\nmy name is\"\n" +
				"\\s+actual\\s+: " + "\"what,\\\\nmy name is\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n\\s+\\++ " +
				"Actual\n" +
				"\\s+@@ -1,2 \\+1,2 @@\n" +
				"\\s+-hi, \n\\s+\\+what,\n" +
				"\\s+my name is",
		},
		{
			name:      "single line diff message",
			equalWant: "want",
			equalGot:  "got",
			want: "\t[a-z]+.go:\\d+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n\\s+-+ Expected\n\\s+\\++ " +
				"Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n",
		},
		{
			name:       "diff message with args",
			equalWant:  "want",
			equalGot:   "got",
			msgAndArgs: []any{"hello, %v!", "world"},
			want: "\t[a-z]+.go:[0-9]+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n" +
				"\\s+\\++ Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n" +
				"\\s+Messages:\\s+hello, world!\n",
		},
		{
			name:       "diff message with integer arg",
			equalWant:  "want",
			equalGot:   "got",
			msgAndArgs: []any{123},
			want: "\t[a-z]+.go:[0-9]+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n" +
				"\\s+\\++ Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n" +
				"\\s+Messages:\\s+123\n",
		},
		{
			name:       "diff message with struct arg",
			equalWant:  "want",
			equalGot:   "got",
			msgAndArgs: []any{struct{ a string }{"hello"}},
			want: "\t[a-z]+.go:[0-9]+: \n" +
				"\t+Error Trace:\t\n" +
				"\t+Error:\\s+Not equal:\\s+\n" +
				"\\s+expected: \"want\"\n" +
				"\\s+actual\\s+: \"got\"\n" +
				"\\s+Diff:\n" +
				"\\s+-+ Expected\n" +
				"\\s+\\++ Actual\n" +
				"\\s+@@ -1 \\+1 @@\n" +
				"\\s+-want\n" +
				"\\s+\\+got\n" +
				"\\s+Messages:\\s+{a:hello}\n",
		},
	})
}

type equalValuesCase struct {
	name          string
	expected      any
	actual        any
	equalValue    bool
	notEqualValue bool // notEqualValue = !equalValue, except for invalid types (e.g. functions)
}

func equalValuesCases() iter.Seq[equalValuesCase] {
	return slices.Values([]equalValuesCase{
		// cases that are expected not to match
		{"not-equal/string", "Hello World", "Hello World!", false, true},
		{"not-equal/int", 123, 1234, false, true},
		{"not-equal/float64", 123.5, 123.55, false, true},
		{"not-equal/[]byte", []byte("Hello World"), []byte("Hello World!"), false, true},
		{"not-equal/nil-not-nil", nil, new(AssertionTesterConformingObject), false, true},
		{"not-equal/converted", uint(10), int(11), false, true},

		// cases that are expected to match
		{"equal/nil-nil", nil, nil, true, false},
		{"equal/string", "Hello World", "Hello World", true, false},
		{"equal/int", 123, 123, true, false},
		{"equal/float64", 123.5, 123.5, true, false},
		{"equal/[]byte", []byte("Hello World"), []byte("Hello World"), true, false},
		{"equal/pointer-to-struct", new(AssertionTesterConformingObject), new(AssertionTesterConformingObject), true, false},
		{"equal/pointer-to-anonymous-struct", &struct{}{}, &struct{}{}, true, false},
		{"equal/converted", int(10), uint(10), true, false},
		{"equal/anonymous-struct", struct{}{}, struct{}{}, true, false},

		// always fail
		{"always-fail/func", func() int { return 23 }, func() int { return 24 }, false, false},
	})
}

type objectEqualExportedValuesCase struct {
	name            string
	expected        any
	actual          any
	expectedEqual   bool
	expectedFailMsg string
}

func objectEqualExportedValuesCases() iter.Seq[objectEqualExportedValuesCase] {
	type specialKey struct {
		a string
	}

	return slices.Values([]objectEqualExportedValuesCase{
		{
			name:          "edge-case/func",
			expected:      func() {},
			actual:        func() {},
			expectedEqual: false,
		},
		{
			name:          "edge-case/expect-nil",
			expected:      nil,
			actual:        nil,
			expectedEqual: true,
		},
		{
			name:          "edge-case/expect-nil-actual-not-nil",
			expected:      nil,
			actual:        1,
			expectedEqual: false,
		},
		{
			name:          "edge-case/map-with-struct-key",
			expected:      map[specialKey]S{{a: "a"}: {}},
			actual:        map[specialKey]S{{a: "a"}: {}},
			expectedEqual: true,
		},
		{
			name: "equal-values/map",
			expected: map[string]S{
				"key": {1, Nested{2, 3}, 4, Nested{5, 6}},
			},
			actual: map[string]S{
				"key": {1, Nested{2, nil}, nil, Nested{}},
			},
			expectedEqual: true,
		},
		{
			name: "diff-values/map",
			expected: map[string]S{
				"key": {1, Nested{2, 3}, 4, Nested{5, 6}},
			},
			actual: map[string]S{
				"x": {1, Nested{2, nil}, nil, Nested{}},
			},
			expectedEqual: false,
		},
		{
			name:          "equal-values/nested-struct",
			expected:      S{1, Nested{2, 3}, 4, Nested{5, 6}},
			actual:        S{1, Nested{2, nil}, nil, Nested{}},
			expectedEqual: true,
		},
		{
			name:          "diff-values/nested-struct(1)",
			expected:      S{1, Nested{2, 3}, 4, Nested{5, 6}},
			actual:        S{1, Nested{1, nil}, nil, Nested{}},
			expectedEqual: false,
			expectedFailMsg: fmt.Sprintf(`
	            	Diff:
	            	--- Expected
	            	+++ Actual
	            	@@ -3,3 +3,3 @@
	            	  Exported2: (%s.Nested) {
	            	-  Exported: (int) 2,
	            	+  Exported: (int) 1,
	            	   notExported: (interface {}) <nil>`,
				shortpkg),
		},
		{
			name:          "diff-values/nested-struct(2)",
			expected:      S3{&Nested{1, 2}, &Nested{3, 4}},
			actual:        S3{&Nested{"a", 2}, &Nested{3, 4}},
			expectedEqual: false,
			expectedFailMsg: fmt.Sprintf(`
	            	Diff:
	            	--- Expected
	            	+++ Actual
	            	@@ -2,3 +2,3 @@
	            	  Exported1: (*%s.Nested)({
	            	-  Exported: (int) 1,
	            	+  Exported: (string) (len=1) "a",
	            	   notExported: (interface {}) <nil>`,
				shortpkg),
		},
		{
			name: "diff-values/inner-slice",
			expected: S4{[]*Nested{
				{1, 2},
				{3, 4},
			}},
			actual: S4{[]*Nested{
				{1, "a"},
				{2, "b"},
			}},
			expectedEqual: false,
			expectedFailMsg: fmt.Sprintf(`
	            	Diff:
	            	--- Expected
	            	+++ Actual
	            	@@ -7,3 +7,3 @@
	            	   (*%s.Nested)({
	            	-   Exported: (int) 3,
	            	+   Exported: (int) 2,
	            	    notExported: (interface {}) <nil>`,
				shortpkg),
		},
		{
			name:          "equal-values/inner-array-unexported-diff",
			expected:      S{[2]int{1, 2}, Nested{2, 3}, 4, Nested{5, 6}},
			actual:        S{[2]int{1, 2}, Nested{2, nil}, nil, Nested{}},
			expectedEqual: true,
		},
		{
			name:          "equal-values/inner-array",
			expected:      &S{1, Nested{2, 3}, 4, Nested{5, 6}},
			actual:        &S{1, Nested{2, nil}, nil, Nested{}},
			expectedEqual: true,
		},
		{
			name:          "diff-values/inner-slice-exported-diff",
			expected:      &S{1, Nested{2, 3}, 4, Nested{5, 6}},
			actual:        &S{1, Nested{1, nil}, nil, Nested{}},
			expectedEqual: false,
			expectedFailMsg: fmt.Sprintf(`
	            	Diff:
	            	--- Expected
	            	+++ Actual
	            	@@ -3,3 +3,3 @@
	            	  Exported2: (%s.Nested) {
	            	-  Exported: (int) 2,
	            	+  Exported: (int) 1,
	            	   notExported: (interface {}) <nil>`,
				shortpkg),
		},
		{
			name:          "equal-values/slice",
			expected:      []int{1, 2},
			actual:        []int{1, 2},
			expectedEqual: true,
		},
		{
			name:          "diff-values/slice",
			expected:      []int{1, 2},
			actual:        []int{1, 3},
			expectedEqual: false,
			expectedFailMsg: `
	            	Diff:
	            	--- Expected
	            	+++ Actual
	            	@@ -2,3 +2,3 @@
	            	  (int) 1,
	            	- (int) 2
	            	+ (int) 3
	            	 }`,
		},
		{
			name:          "equal-values/slice-of-pointers",
			expected:      []*int{ptr(1), nil, ptr(2)},
			actual:        []*int{ptr(1), nil, ptr(2)},
			expectedEqual: true,
		},
		{
			name: "equal-values/slice-of-struct",
			expected: []*Nested{
				{1, 2},
				{3, 4},
			},
			actual: []*Nested{
				{1, "a"},
				{3, "b"},
			},
			expectedEqual: true,
		},
		{
			name: "diff-values/slice-of-struct",
			expected: []*Nested{
				{1, 2},
				{3, 4},
			},
			actual: []*Nested{
				{1, "a"},
				{2, "b"},
			},
			expectedEqual: false,
			expectedFailMsg: fmt.Sprintf(`
	            	Diff:
	            	--- Expected
	            	+++ Actual
	            	@@ -6,3 +6,3 @@
	            	  (*%s.Nested)({
	            	-  Exported: (int) 3,
	            	+  Exported: (int) 2,
	            	   notExported: (interface {}) <nil>`,
				shortpkg),
		},
	})
}
