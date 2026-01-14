// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"iter"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"strings"
	"testing"
)

func TestCollectionLen(t *testing.T) {
	t.Parallel()

	t.Run("with invalid types", func(t *testing.T) {
		t.Parallel()
		mock := new(testing.T)

		False(t, Len(mock, nil, 0), "nil does not have length")
		False(t, Len(mock, 0, 0), "int does not have length")
		False(t, Len(mock, true, 0), "true does not have length")
		False(t, Len(mock, false, 0), "false does not have length")
		False(t, Len(mock, 'A', 0), "Rune does not have length")
		False(t, Len(mock, struct{}{}, 0), "Struct does not have length")
	})

	t.Run("with valid types", func(t *testing.T) {
		t.Parallel()
		mock := new(testing.T)

		for c := range collectionValidLenCases() {
			True(t, Len(mock, c.v, c.l), "%#v have %d items", c.v, c.l)
			False(t, Len(mock, c.v, c.l+1), "%#v have %d items", c.v, c.l)

			if c.expected1234567 != "" {
				msgMock := new(mockT)
				Len(msgMock, c.v, 1234567)
				Contains(t, msgMock.errorString(), c.expected1234567)
			}
		}
	})

	t.Run("with slice too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		Len(mock, longSlice, 1)
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	"[0 0 0`)
		Contains(t, mock.errorString(), `<... truncated>" should have 1 item(s), but has 1000000`)
	})
}

func TestCollectionContains(t *testing.T) {
	t.Parallel()

	t.Run("nil with slice too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		Nil(mock, &longSlice)
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Expected nil, but got: &[]int{0, 0, 0,`)
		Contains(t, mock.errorString(), `<... truncated>`)
	})

	t.Run("empty with slice too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		Empty(mock, longSlice)
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	Should be empty, but was [0 0 0`)
		Contains(t, mock.errorString(), `<... truncated>`)
	})

	t.Run("with slice too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		Contains(mock, longSlice, 1)
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	[]int{0, 0, 0,`)
		Contains(t, mock.errorString(), `<... truncated> does not contain 1`)
	})
}

func TestCollectionNotContains(t *testing.T) {
	t.Parallel()

	t.Run("with slice too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		NotContains(mock, longSlice, 0)
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	[]int{0, 0, 0,`)
		Contains(t, mock.errorString(), `<... truncated> should not contain 0`)
	})
}

func TestCollectionSubset(t *testing.T) {
	t.Parallel()

	t.Run("with slice too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		Subset(mock, longSlice, []int{1})
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	[]int{0, 0, 0,`)
		Contains(t, mock.errorString(), `<... truncated> does not contain 1`)
	})

	t.Run("with map too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		Subset(mock, map[bool][]int{true: longSlice}, map[bool][]int{false: longSlice})
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	map[bool][]int{true:[]int{0, 0, 0,`)
		Contains(t, mock.errorString(), `<... truncated> does not contain map[bool][]int{false:[]int{0, 0, 0,`)
	})
}

func TestCollectionNotSubset(t *testing.T) {
	t.Parallel()

	t.Run("with slice too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		NotSubset(mock, longSlice, longSlice)
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	['\x00' '\x00' '\x00'`)
		Contains(t, mock.errorString(), `<... truncated> is a subset of ['\x00' '\x00' '\x00'`)
	})

	t.Run("with map too long to print", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		longSlice := make([]int, 1_000_000)
		NotSubset(mock, map[int][]int{1: longSlice}, map[int][]int{1: longSlice})
		Contains(t, mock.errorString(), `
	Error Trace:	
	Error:      	map['\x01':['\x00' '\x00' '\x00'`)
		Contains(t, mock.errorString(), `<... truncated> is a subset of map['\x01':['\x00' '\x00' '\x00'`)
	})
}

func TestCollectionContainsNotContains(t *testing.T) {
	t.Parallel()

	for c := range collectionContainsCases() {
		t.Run(fmt.Sprintf("Contains(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			t.Parallel()
			mock := new(testing.T)

			res := Contains(mock, c.expected, c.actual)

			if res != c.result {
				if res {
					t.Errorf("Contains(%#v, %#v) should return true:\n\t%#v contains %#v", c.expected, c.actual, c.expected, c.actual)
				} else {
					t.Errorf("Contains(%#v, %#v) should return false:\n\t%#v does not contain %#v", c.expected, c.actual, c.expected, c.actual)
				}
			}
		})
	}

	for c := range collectionContainsCases() {
		t.Run(fmt.Sprintf("NotContains(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			t.Parallel()
			mock := new(testing.T)

			res := NotContains(mock, c.expected, c.actual)

			// NotContains should be inverse of Contains. If it's not, something is wrong
			if res == Contains(mock, c.expected, c.actual) {
				if res {
					t.Errorf("NotContains(%#v, %#v) should return true:\n\t%#v does not contains %#v", c.expected, c.actual, c.expected, c.actual)
				} else {
					t.Errorf("NotContains(%#v, %#v) should return false:\n\t%#v contains %#v", c.expected, c.actual, c.expected, c.actual)
				}
			}
		})
	}
}

func TestCollectionContainsNotContainsFailMessage(t *testing.T) {
	t.Parallel()

	for c := range collectionContainsFailCases() {
		name := filepath.Base(runtime.FuncForPC(reflect.ValueOf(c.assertion).Pointer()).Name())
		t.Run(fmt.Sprintf("%v(%T, %T)", name, c.container, c.instance), func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)

			c.assertion(mock, c.container, c.instance)
			actualFail := mock.errorString()
			if !strings.Contains(actualFail, c.expected) {
				t.Errorf("Contains failure should include %q but was %q", c.expected, actualFail)
			}
		})
	}
}

func TestCollectionContainsNotContainsOnNilValue(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	Contains(mock, nil, "key")
	expectedFail := "<nil> could not be applied builtin len()"
	actualFail := mock.errorString()
	if !strings.Contains(actualFail, expectedFail) {
		t.Errorf("Contains failure should include %q but was %q", expectedFail, actualFail)
	}

	NotContains(mock, nil, "key")
	if !strings.Contains(actualFail, expectedFail) {
		t.Errorf("Contains failure should include %q but was %q", expectedFail, actualFail)
	}
}

func TestCollectionSubsetNotSubset(t *testing.T) {
	t.Parallel()

	for c := range collectionSubsetCases() {
		t.Run("SubSet: "+c.message, func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)

			res := Subset(mock, c.list, c.subset)

			if res != c.result {
				t.Errorf("Subset should return %t: %s", c.result, c.message)
			}

			if !c.result {
				expectedFail := c.message
				actualFail := mock.errorString()
				if !strings.Contains(actualFail, expectedFail) {
					t.Log(actualFail)
					t.Errorf("Subset failure should contain %q but was %q", expectedFail, actualFail)
				}
			}
		})
	}

	for c := range collectionSubsetCases() {
		t.Run("NotSubSet: "+c.message, func(t *testing.T) {
			t.Parallel()
			mock := new(mockT)

			res := NotSubset(mock, c.list, c.subset)

			// NotSubset should match the inverse of Subset. If it doesn't, something is wrong
			if res == Subset(mock, c.list, c.subset) {
				t.Errorf("NotSubset should return %t: %s", !c.result, c.message)
			}

			if c.result {
				expectedFail := c.message
				actualFail := mock.errorString()
				if !strings.Contains(actualFail, expectedFail) {
					t.Log(actualFail)
					t.Errorf("NotSubset failure should contain %q but was %q", expectedFail, actualFail)
				}
			}
		})
	}
}

func TestCollectionNotSubsetNil(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	NotSubset(mock, []string{"foo"}, nil)
	if !mock.Failed() {
		t.Error("NotSubset on nil set should have failed the test")
	}
}

func TestCollectionElementsMatch(t *testing.T) {
	t.Parallel()

	for c := range collectionElementsMatchCases() {
		t.Run(fmt.Sprintf("ElementsMatch(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			t.Parallel()
			mock := new(testing.T)

			res := ElementsMatch(mock, c.actual, c.expected)

			if res != c.result {
				t.Errorf("ElementsMatch(%#v, %#v) should return %v", c.actual, c.expected, c.result)
			}
		})
	}
}

func TestCollectionElementsMatchT(t *testing.T) {
	t.Parallel()

	for tc := range elementsMatchTCases() {
		t.Run(tc.name, tc.test)
	}
}

func TestCollectionNotElementsMatch(t *testing.T) {
	t.Parallel()

	for c := range collectionNotElementsMatchCases() {
		t.Run(fmt.Sprintf("NotElementsMatch(%#v, %#v)", c.expected, c.actual), func(t *testing.T) {
			t.Parallel()
			mock := new(testing.T)

			res := NotElementsMatch(mock, c.actual, c.expected)

			if res != c.result {
				t.Errorf("NotElementsMatch(%#v, %#v) should return %v", c.actual, c.expected, c.result)
			}
		})
	}
}

func TestCollectionNotElementsMatchT(t *testing.T) {
	t.Parallel()

	for tc := range notElementsMatchTCases() {
		t.Run(tc.name, tc.test)
	}
}

/* iterators for test cases */

type collectionValidLenCase struct {
	v               any
	l               int
	expected1234567 string // message when expecting 1234567 items
}

func collectionValidLenCases() iter.Seq[collectionValidLenCase] {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3

	return slices.Values([]collectionValidLenCase{
		{[]int{1, 2, 3}, 3, `"[1 2 3]" should have 1234567 item(s), but has 3`},
		{[...]int{1, 2, 3}, 3, `"[1 2 3]" should have 1234567 item(s), but has 3`},
		{"ABC", 3, `"ABC" should have 1234567 item(s), but has 3`},
		{map[int]int{1: 2, 2: 4, 3: 6}, 3, `"map[1:2 2:4 3:6]" should have 1234567 item(s), but has 3`},
		{ch, 3, ""},

		{[]int{}, 0, `"[]" should have 1234567 item(s), but has 0`},
		{map[int]int{}, 0, `"map[]" should have 1234567 item(s), but has 0`},
		{make(chan int), 0, ""},

		{[]int(nil), 0, `"[]" should have 1234567 item(s), but has 0`},
		{map[int]int(nil), 0, `"map[]" should have 1234567 item(s), but has 0`},
		{(chan int)(nil), 0, `"<nil>" should have 1234567 item(s), but has 0`},
	})
}

type collectionContainsCase = testCase

func collectionContainsCases() iter.Seq[collectionContainsCase] {
	type A struct {
		Name, Value string
	}

	list := []string{"Foo", "Bar"}
	complexList := []*A{
		{"b", "c"},
		{"d", "e"},
		{"g", "h"},
		{"j", "k"},
	}
	simpleMap := map[any]any{"Foo": "Bar"}
	var zeroMap map[any]any

	return slices.Values([]collectionContainsCase{
		{"Hello World", "Hello", true},
		{"Hello World", "Salut", false},
		{list, "Bar", true},
		{list, "Salut", false},
		{complexList, &A{"g", "h"}, true},
		{complexList, &A{"g", "e"}, false},
		{simpleMap, "Foo", true},
		{simpleMap, "Bar", false},
		{zeroMap, "Bar", false},
	})
}

type collectionContainsFailCase struct {
	assertion func(t T, s, contains any, msgAndArgs ...any) bool
	container any
	instance  any
	expected  string
}

func collectionContainsFailCases() iter.Seq[collectionContainsFailCase] {
	const pkg = "assertions"
	type nonContainer struct {
		Value string
	}

	return slices.Values([]collectionContainsFailCase{
		{
			assertion: Contains,
			container: "Hello World",
			instance:  errors.New("Hello"),
			expected:  "\"Hello World\" does not contain &errors.errorString{s:\"Hello\"}",
		},
		{
			assertion: Contains,
			container: map[string]int{"one": 1},
			instance:  "two",
			expected:  "map[string]int{\"one\":1} does not contain \"two\"\n",
		},
		{
			assertion: NotContains,
			container: map[string]int{"one": 1},
			instance:  "one",
			expected:  "map[string]int{\"one\":1} should not contain \"one\"",
		},
		{
			assertion: Contains,
			container: nonContainer{Value: "Hello"},
			instance:  "Hello",
			expected:  pkg + ".nonContainer{Value:\"Hello\"} could not be applied builtin len()\n",
		},
		{
			assertion: NotContains,
			container: nonContainer{Value: "Hello"},
			instance:  "Hello",
			expected:  pkg + ".nonContainer{Value:\"Hello\"} could not be applied builtin len()\n",
		},
	})
}

type collectionSubsetCase struct {
	list    any
	subset  any
	result  bool
	message string
}

func collectionSubsetCases() iter.Seq[collectionSubsetCase] {
	return slices.Values([]collectionSubsetCase{
		// cases that are expected to contain
		{[]int{1, 2, 3}, nil, true, `nil is the empty set which is a subset of every set`},
		{[]int{1, 2, 3}, []int{}, true, `[] is a subset of ['\x01' '\x02' '\x03']`},
		{[]int{1, 2, 3}, []int{1, 2}, true, `['\x01' '\x02'] is a subset of ['\x01' '\x02' '\x03']`},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true, `['\x01' '\x02' '\x03'] is a subset of ['\x01' '\x02' '\x03']`},
		{[]string{"hello", "world"}, []string{"hello"}, true, `["hello"] is a subset of ["hello" "world"]`},
		{map[string]string{
			"a": "x",
			"c": "z",
			"b": "y",
		}, map[string]string{
			"a": "x",
			"b": "y",
		}, true, `map["a":"x" "b":"y"] is a subset of map["a":"x" "b":"y" "c":"z"]`},
		{[]string{"a", "b", "c"}, map[string]int{"a": 1, "c": 3}, true, `map["a":'\x01' "c":'\x03'] is a subset of ["a" "b" "c"]`},

		// cases that are expected not to contain
		{[]string{"hello", "world"}, []string{"hello", "testify"}, false, `[]string{"hello", "world"} does not contain "testify"`},
		{[]int{1, 2, 3}, []int{4, 5}, false, `[]int{1, 2, 3} does not contain 4`},
		{[]int{1, 2, 3}, []int{1, 5}, false, `[]int{1, 2, 3} does not contain 5`},
		{map[string]string{
			"a": "x",
			"c": "z",
			"b": "y",
		}, map[string]string{
			"a": "x",
			"b": "z",
		}, false, `map[string]string{"a":"x", "b":"y", "c":"z"} does not contain map[string]string{"a":"x", "b":"z"}`},
		{map[string]string{
			"a": "x",
			"b": "y",
		}, map[string]string{
			"a": "x",
			"b": "y",
			"c": "z",
		}, false, `map[string]string{"a":"x", "b":"y"} does not contain map[string]string{"a":"x", "b":"y", "c":"z"}`},
		{[]string{"a", "b", "c"}, map[string]int{"c": 3, "d": 4}, false, `[]string{"a", "b", "c"} does not contain "d"`},
	})
}

type collectionElementsMatchCase = testCase

func collectionElementsMatchCases() iter.Seq[collectionElementsMatchCase] {
	return slices.Values([]collectionElementsMatchCase{
		// matching
		{nil, nil, true},

		{nil, nil, true},
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{1, 1}, []int{1, 1}, true},
		{[]int{1, 2}, []int{1, 2}, true},
		{[]int{1, 2}, []int{2, 1}, true},
		{[2]int{1, 2}, [2]int{2, 1}, true},
		{[]string{"hello", "world"}, []string{"world", "hello"}, true},
		{[]string{"hello", "hello"}, []string{"hello", "hello"}, true},
		{[]string{"hello", "hello", "world"}, []string{"hello", "world", "hello"}, true},
		{[3]string{"hello", "hello", "world"}, [3]string{"hello", "world", "hello"}, true},
		{[]int{}, nil, true},

		// not matching
		{[]int{1}, []int{1, 1}, false},
		{[]int{1, 2}, []int{2, 2}, false},
		{[]string{"hello", "hello"}, []string{"hello"}, false},
	})
}

type collectionNotElementsMatch = testCase

func collectionNotElementsMatchCases() iter.Seq[collectionNotElementsMatch] {
	return slices.Values([]collectionNotElementsMatch{
		// not matching
		{[]int{1}, []int{}, true},
		{[]int{}, []int{2}, true},
		{[]int{1}, []int{2}, true},
		{[]int{1}, []int{1, 1}, true},
		{[]int{1, 2}, []int{3, 4}, true},
		{[]int{3, 4}, []int{1, 2}, true},
		{[]int{1, 1, 2, 3}, []int{1, 2, 3}, true},
		{[]string{"hello"}, []string{"world"}, true},
		{[]string{"hello", "hello"}, []string{"world", "world"}, true},
		{[3]string{"hello", "hello", "hello"}, [3]string{"world", "world", "world"}, true},

		// matching
		{nil, nil, false},
		{[]int{}, nil, false},
		{[]int{}, []int{}, false},
		{[]int{1}, []int{1}, false},
		{[]int{1, 1}, []int{1, 1}, false},
		{[]int{1, 2}, []int{2, 1}, false},
		{[2]int{1, 2}, [2]int{2, 1}, false},
		{[]int{1, 1, 2}, []int{1, 2, 1}, false},
		{[]string{"hello", "world"}, []string{"world", "hello"}, false},
		{[]string{"hello", "hello"}, []string{"hello", "hello"}, false},
		{[]string{"hello", "hello", "world"}, []string{"hello", "world", "hello"}, false},
		{[3]string{"hello", "hello", "world"}, [3]string{"hello", "world", "hello"}, false},
	})
}

// elementsMatchTTestPairs builds test cases for both ElementsMatchT and NotElementsMatchT
// from shared test data, ensuring consistency between the inverse functions.
func elementsMatchTTestPairs() (matchCases, notMatchCases []genericTestCase) {
	// addPair adds corresponding test cases for both ElementsMatchT and NotElementsMatchT.
	addPair := func(name string, matchTest, notMatchTest func(*testing.T)) {
		matchCases = append(matchCases, genericTestCase{name, matchTest})
		notMatchCases = append(notMatchCases, genericTestCase{name, notMatchTest})
	}

	// Numeric types - test data defined once, used for both functions
	m, n := testElementsMatchTPair[int]([]int{1, 2, 3}, []int{3, 1, 2}, []int{1, 2, 3}, []int{1, 2, 4})
	addPair("int", m, n)
	m, n = testElementsMatchTPair[int8]([]int8{1, 2, 3}, []int8{3, 1, 2}, []int8{1, 2, 3}, []int8{1, 2, 4})
	addPair("int8", m, n)
	m, n = testElementsMatchTPair[int16]([]int16{1, 2, 3}, []int16{3, 1, 2}, []int16{1, 2, 3}, []int16{1, 2, 4})
	addPair("int16", m, n)
	m, n = testElementsMatchTPair[int32]([]int32{1, 2, 3}, []int32{3, 1, 2}, []int32{1, 2, 3}, []int32{1, 2, 4})
	addPair("int32", m, n)
	m, n = testElementsMatchTPair[int64]([]int64{1, 2, 3}, []int64{3, 1, 2}, []int64{1, 2, 3}, []int64{1, 2, 4})
	addPair("int64", m, n)
	m, n = testElementsMatchTPair[uint]([]uint{1, 2, 3}, []uint{3, 1, 2}, []uint{1, 2, 3}, []uint{1, 2, 4})
	addPair("uint", m, n)
	m, n = testElementsMatchTPair[uint8]([]uint8{1, 2, 3}, []uint8{3, 1, 2}, []uint8{1, 2, 3}, []uint8{1, 2, 4})
	addPair("uint8", m, n)
	m, n = testElementsMatchTPair[uint16]([]uint16{1, 2, 3}, []uint16{3, 1, 2}, []uint16{1, 2, 3}, []uint16{1, 2, 4})
	addPair("uint16", m, n)
	m, n = testElementsMatchTPair[uint32]([]uint32{1, 2, 3}, []uint32{3, 1, 2}, []uint32{1, 2, 3}, []uint32{1, 2, 4})
	addPair("uint32", m, n)
	m, n = testElementsMatchTPair[uint64]([]uint64{1, 2, 3}, []uint64{3, 1, 2}, []uint64{1, 2, 3}, []uint64{1, 2, 4})
	addPair("uint64", m, n)
	m, n = testElementsMatchTPair[float32]([]float32{1.5, 2.5, 3.5}, []float32{3.5, 1.5, 2.5}, []float32{1.5, 2.5, 3.5}, []float32{1.5, 2.5, 4.5})
	addPair("float32", m, n)
	m, n = testElementsMatchTPair[float64]([]float64{1.5, 2.5, 3.5}, []float64{3.5, 1.5, 2.5}, []float64{1.5, 2.5, 3.5}, []float64{1.5, 2.5, 4.5})
	addPair("float64", m, n)
	m, n = testElementsMatchTPair[string]([]string{"a", "b", "c"}, []string{"c", "a", "b"}, []string{"a", "b", "c"}, []string{"a", "b", "d"})
	addPair("string", m, n)
	m, n = testElementsMatchTPair[bool]([]bool{true, false}, []bool{false, true}, []bool{true, true}, []bool{true, false})
	addPair("bool", m, n)

	// Special cases
	m, n = testElementsMatchTEmptyPair()
	addPair("empty slices", m, n)
	m, n = testElementsMatchTDuplicatesPair()
	addPair("with duplicates", m, n)
	m, n = testElementsMatchTCustomTypePair()
	addPair("custom type", m, n)
	m, n = testElementsMatchTStructPair()
	addPair("struct type", m, n)

	return matchCases, notMatchCases
}

// elementsMatchTCases returns test cases for ElementsMatchT with various comparable types.
func elementsMatchTCases() iter.Seq[genericTestCase] {
	matchCases, _ := elementsMatchTTestPairs()
	return slices.Values(matchCases)
}

// notElementsMatchTCases returns test cases for NotElementsMatchT with various comparable types.
func notElementsMatchTCases() iter.Seq[genericTestCase] {
	_, notMatchCases := elementsMatchTTestPairs()
	return slices.Values(notMatchCases)
}

// testElementsMatchTPair creates test functions for both ElementsMatchT and NotElementsMatchT
// from the same test data, ensuring consistency between inverse functions.
// matchA/matchB are slices that should match; noMatchA/noMatchB are slices that should not match.
//
//nolint:thelper // linter false positive: these are not helpers
func testElementsMatchTPair[E comparable](matchA, matchB, noMatchA, noMatchB []E) (matchTest, notMatchTest func(*testing.T)) {
	matchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, ElementsMatchT(mock, matchA, matchB))
		False(t, ElementsMatchT(mock, noMatchA, noMatchB))
	}
	notMatchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, NotElementsMatchT(mock, noMatchA, noMatchB))
		False(t, NotElementsMatchT(mock, matchA, matchB))
	}
	return matchTest, notMatchTest
}

//nolint:thelper // linter false positive: these are not helpers
func testElementsMatchTEmptyPair() (matchTest, notMatchTest func(*testing.T)) {
	matchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, ElementsMatchT(mock, []int{}, []int{}))
		True(t, ElementsMatchT(mock, []string(nil), []string(nil)))
		True(t, ElementsMatchT(mock, []int(nil), []int{}))
	}
	notMatchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		// Empty slices match, so NotElementsMatchT should return false
		False(t, NotElementsMatchT(mock, []int{}, []int{}))
		False(t, NotElementsMatchT(mock, []string(nil), []string(nil)))
		// One empty, one not - they don't match
		True(t, NotElementsMatchT(mock, []int{1}, []int{}))
		True(t, NotElementsMatchT(mock, []int{}, []int{1}))
	}
	return matchTest, notMatchTest
}

//nolint:thelper // linter false positive: these are not helpers
func testElementsMatchTDuplicatesPair() (matchTest, notMatchTest func(*testing.T)) {
	matchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, ElementsMatchT(mock, []int{1, 1, 2}, []int{2, 1, 1}))
		False(t, ElementsMatchT(mock, []int{1, 1, 2}, []int{1, 2, 2}))
		False(t, ElementsMatchT(mock, []int{1, 1, 2}, []int{1, 2}))
	}
	notMatchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		// Different duplicate counts - should not match
		True(t, NotElementsMatchT(mock, []int{1, 1, 2}, []int{1, 2, 2}))
		True(t, NotElementsMatchT(mock, []int{1, 1, 2}, []int{1, 2}))
		// Same duplicates, different order - should match (NotElementsMatchT returns false)
		False(t, NotElementsMatchT(mock, []int{1, 1, 2}, []int{2, 1, 1}))
	}
	return matchTest, notMatchTest
}

//nolint:thelper // linter false positive: these are not helpers
func testElementsMatchTCustomTypePair() (matchTest, notMatchTest func(*testing.T)) {
	type myInt int
	matchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, ElementsMatchT(mock, []myInt{1, 2, 3}, []myInt{3, 2, 1}))
		False(t, ElementsMatchT(mock, []myInt{1, 2, 3}, []myInt{1, 2, 4}))
	}
	notMatchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, NotElementsMatchT(mock, []myInt{1, 2, 3}, []myInt{1, 2, 4}))
		False(t, NotElementsMatchT(mock, []myInt{1, 2, 3}, []myInt{3, 2, 1}))
	}
	return matchTest, notMatchTest
}

//nolint:thelper // linter false positive: these are not helpers
func testElementsMatchTStructPair() (matchTest, notMatchTest func(*testing.T)) {
	type point struct{ x, y int }
	p1, p2, p3 := point{1, 2}, point{3, 4}, point{5, 6}
	matchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, ElementsMatchT(mock, []point{p1, p2, p3}, []point{p3, p1, p2}))
		False(t, ElementsMatchT(mock, []point{p1, p2}, []point{p1, p3}))
	}
	notMatchTest = func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		True(t, NotElementsMatchT(mock, []point{p1, p2}, []point{p1, p3}))
		False(t, NotElementsMatchT(mock, []point{p1, p2, p3}, []point{p3, p1, p2}))
	}
	return matchTest, notMatchTest
}
