// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"iter"
	"math"
	"slices"
	"strings"
	"testing"
	"time"
)

func TestObjectsAreEqual(t *testing.T) {
	t.Parallel()

	for c := range objectEqualCases() {
		t.Run(fmt.Sprintf("ObjectsAreEqual(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := ObjectsAreEqual(c.expected, c.actual)

			if res != c.result {
				t.Errorf("ObjectsAreEqual(%#v, %#v) should return %#v", c.expected, c.actual, c.result)
			}
		})
	}
}

func TestObjectsAreEqualValues(t *testing.T) {
	t.Parallel()

	for c := range objectEqualValuesCases() {
		t.Run(fmt.Sprintf("ObjectsAreEqualValues(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			res := ObjectsAreEqualValues(c.expected, c.actual)

			if res != c.result {
				t.Errorf("ObjectsAreEqualValues(%#v, %#v) should return %#v", c.expected, c.actual, c.result)
			}
		})
	}
}

func TestObjectsCopyExportedFields(t *testing.T) {
	t.Parallel()

	for c := range objectCopyExportedFieldsCases() {
		t.Run("", func(t *testing.T) {
			output := copyExportedFields(c.input)
			if !ObjectsAreEqualValues(c.expected, output) {
				t.Errorf("%#v, %#v should be equal", c.expected, output)
			}
		})
	}
}

func TestObjectsEqualExportedValues(t *testing.T) {
	t.Parallel()

	for c := range objectEqualExportedValuesCases() {
		t.Run("", func(t *testing.T) {
			mockT := new(mockT)

			actual := EqualExportedValues(mockT, c.value1, c.value2)
			if actual != c.expectedEqual {
				t.Errorf("Expected EqualExportedValues to be %t, but was %t", c.expectedEqual, actual)
			}

			actualFail := mockT.errorString()
			if !strings.Contains(actualFail, c.expectedFail) {
				t.Errorf("Contains failure should include %q but was %q", c.expectedFail, actualFail)
			}
		})
	}
}

type Nested struct {
	Exported    any
	notExported any
}

type S struct {
	Exported1    any
	Exported2    Nested
	notExported1 any
	notExported2 Nested
}

type S3 struct {
	Exported1 *Nested
	Exported2 *Nested
}

type S4 struct {
	Exported1 []*Nested
}

type S5 struct {
	Exported Nested
}

type S6 struct {
	Exported   string
	unexported string
}

type objectEqualCase struct {
	expected any
	actual   any
	result   bool
}

func objectEqualCases() iter.Seq[objectEqualCase] {
	return slices.Values([]objectEqualCase{
		// cases that are expected to be equal
		{"Hello World", "Hello World", true},
		{123, 123, true},
		{123.5, 123.5, true},
		{[]byte("Hello World"), []byte("Hello World"), true},
		{nil, nil, true},

		// cases that are expected not to be equal
		{map[int]int{5: 10}, map[int]int{10: 20}, false},
		{'x', "x", false},
		{"x", 'x', false},
		{0, 0.1, false},
		{0.1, 0, false},
		{time.Now, time.Now, false},
		{func() {}, func() {}, false},
		{uint32(10), int32(10), false},
	})
}

func objectEqualValuesCases() iter.Seq[objectEqualCase] {
	now := time.Now()

	return slices.Values([]objectEqualCase{
		{uint32(10), int32(10), true},
		{0, nil, false},
		{nil, 0, false},
		// should not be time zone independent
		{now, now.In(time.Local), false}, //nolint:gosmopolitan // ok in this context: this is precisely the goal of this test
		{int(270), int8(14), false},      // should handle overflow/underflow
		{int8(14), int(270), false},
		{[]int{270, 270}, []int8{14, 14}, false},
		{complex128(1e+100 + 1e+100i), complex64(complex(math.Inf(0), math.Inf(0))), false},
		{complex64(complex(math.Inf(0), math.Inf(0))), complex128(1e+100 + 1e+100i), false},
		{complex128(1e+100 + 1e+100i), 270, false},
		{270, complex128(1e+100 + 1e+100i), false},
		{complex128(1e+100 + 1e+100i), 3.14, false},
		{3.14, complex128(1e+100 + 1e+100i), false},
		{complex128(1e+10 + 1e+10i), complex64(1e+10 + 1e+10i), true},
		{complex64(1e+10 + 1e+10i), complex128(1e+10 + 1e+10i), true},
		{[]int{1, 2}, (*[3]int)(nil), false}, // panics should be caught and treated as inequality (https://github.com/stretchr/testify/issues/1699)
	})
}

type objectCopyFieldsCase struct {
	input    any
	expected any
}

func objectCopyExportedFieldsCases() iter.Seq[objectCopyFieldsCase] {
	intValue := 1

	return slices.Values([]objectCopyFieldsCase{
		{
			input:    Nested{"a", "b"},
			expected: Nested{"a", nil},
		},
		{
			input:    Nested{&intValue, 2},
			expected: Nested{&intValue, nil},
		},
		{
			input:    Nested{nil, 3},
			expected: Nested{nil, nil},
		},
		{
			input:    S{1, Nested{2, 3}, 4, Nested{5, 6}},
			expected: S{1, Nested{2, nil}, nil, Nested{}},
		},
		{
			input:    S3{},
			expected: S3{},
		},
		{
			input:    S3{&Nested{1, 2}, &Nested{3, 4}},
			expected: S3{&Nested{1, nil}, &Nested{3, nil}},
		},
		{
			input:    S3{Exported1: &Nested{"a", "b"}},
			expected: S3{Exported1: &Nested{"a", nil}},
		},
		{
			input: S4{[]*Nested{
				nil,
				{1, 2},
			}},
			expected: S4{[]*Nested{
				nil,
				{1, nil},
			}},
		},
		{
			input: S4{
				[]*Nested{
					{1, 2},
				},
			},
			expected: S4{
				[]*Nested{
					{1, nil},
				},
			},
		},
		{
			input: S4{[]*Nested{
				{1, 2},
				{3, 4},
			}},
			expected: S4{[]*Nested{
				{1, nil},
				{3, nil},
			}},
		},
		{
			input:    S5{Exported: Nested{"a", "b"}},
			expected: S5{Exported: Nested{"a", nil}},
		},
		{
			input:    S6{"a", "b"},
			expected: S6{"a", ""},
		},
	})
}

type objectEqualExportedValuesCase struct {
	value1        any
	value2        any
	expectedEqual bool
	expectedFail  string
}

func objectEqualExportedValuesCases() iter.Seq[objectEqualExportedValuesCase] {
	return slices.Values([]objectEqualExportedValuesCase{
		{
			value1:        S{1, Nested{2, 3}, 4, Nested{5, 6}},
			value2:        S{1, Nested{2, nil}, nil, Nested{}},
			expectedEqual: true,
		},
		{
			value1:        S{1, Nested{2, 3}, 4, Nested{5, 6}},
			value2:        S{1, Nested{1, nil}, nil, Nested{}},
			expectedEqual: false,
			expectedFail: fmt.Sprintf(`
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
			value1:        S3{&Nested{1, 2}, &Nested{3, 4}},
			value2:        S3{&Nested{"a", 2}, &Nested{3, 4}},
			expectedEqual: false,
			expectedFail: fmt.Sprintf(`
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
			value1: S4{[]*Nested{
				{1, 2},
				{3, 4},
			}},
			value2: S4{[]*Nested{
				{1, "a"},
				{2, "b"},
			}},
			expectedEqual: false,
			expectedFail: fmt.Sprintf(`
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
			value1:        S{[2]int{1, 2}, Nested{2, 3}, 4, Nested{5, 6}},
			value2:        S{[2]int{1, 2}, Nested{2, nil}, nil, Nested{}},
			expectedEqual: true,
		},
		{
			value1:        &S{1, Nested{2, 3}, 4, Nested{5, 6}},
			value2:        &S{1, Nested{2, nil}, nil, Nested{}},
			expectedEqual: true,
		},
		{
			value1:        &S{1, Nested{2, 3}, 4, Nested{5, 6}},
			value2:        &S{1, Nested{1, nil}, nil, Nested{}},
			expectedEqual: false,
			expectedFail: fmt.Sprintf(`
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
			value1:        []int{1, 2},
			value2:        []int{1, 2},
			expectedEqual: true,
		},
		{
			value1:        []int{1, 2},
			value2:        []int{1, 3},
			expectedEqual: false,
			expectedFail: `
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
			value1: []*Nested{
				{1, 2},
				{3, 4},
			},
			value2: []*Nested{
				{1, "a"},
				{3, "b"},
			},
			expectedEqual: true,
		},
		{
			value1: []*Nested{
				{1, 2},
				{3, 4},
			},
			value2: []*Nested{
				{1, "a"},
				{2, "b"},
			},
			expectedEqual: false,
			expectedFail: fmt.Sprintf(`
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
