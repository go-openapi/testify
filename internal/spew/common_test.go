/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spew_test

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/internal/spew"
)

// TestSortValues ensures the sort functionality for relect.Value based sorting
// works as intended.
func TestSortValues(t *testing.T) {
	v := reflect.ValueOf

	a := v("a")
	b := v("b")
	c := v("c")
	embedA := v(embed{"a"})
	embedB := v(embed{"b"})
	embedC := v(embed{"c"})

	// test values for times
	t0, t1, t2 := testTimings()
	lt0, lt1, lt2 := testTimingsWithLocation()
	pt0, pt1, pt2 := ptr(t0), ptr(t1), ptr(t2)
	ppt2 := ptr(pt2)
	nilTimePtr := (*time.Time)(nil)
	nilTimePtrPtr := &nilTimePtr
	invalidTime := (**time.Time)(nil)

	tests := []sortTestCase{
		// No values.
		{
			[]reflect.Value{},
			[]reflect.Value{},
		},
		// Bools.
		{
			[]reflect.Value{v(false), v(true), v(false)},
			[]reflect.Value{v(false), v(false), v(true)},
		},
		// Ints.
		{
			[]reflect.Value{v(2), v(1), v(3)},
			[]reflect.Value{v(1), v(2), v(3)},
		},
		// Uints.
		{
			[]reflect.Value{v(uint8(2)), v(uint8(1)), v(uint8(3))},
			[]reflect.Value{v(uint8(1)), v(uint8(2)), v(uint8(3))},
		},
		// Floats.
		{
			[]reflect.Value{v(2.0), v(1.0), v(3.0)},
			[]reflect.Value{v(1.0), v(2.0), v(3.0)},
		},
		// Strings.
		{
			[]reflect.Value{b, a, c},
			[]reflect.Value{a, b, c},
		},
		// Array
		{
			[]reflect.Value{v([3]int{3, 2, 1}), v([3]int{1, 3, 2}), v([3]int{1, 2, 3})},
			[]reflect.Value{v([3]int{1, 2, 3}), v([3]int{1, 3, 2}), v([3]int{3, 2, 1})},
		},
		// Uintptrs.
		{
			[]reflect.Value{v(uintptr(2)), v(uintptr(1)), v(uintptr(3))},
			[]reflect.Value{v(uintptr(1)), v(uintptr(2)), v(uintptr(3))},
		},
		// SortableStructs.
		{
			// Note: not sorted - DisableMethods is set.
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
		},
		// UnsortableStructs.
		{
			// Note: not sorted - SpewKeys is false.
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
		},
		// time.Time
		{
			[]reflect.Value{v(t0), v(t2), v(t1)},
			[]reflect.Value{v(t0), v(t1), v(t2)},
		},
		// comparison with location
		{
			[]reflect.Value{v(lt2), v(lt1), v(lt0)},
			[]reflect.Value{v(lt0), v(lt1), v(lt2)},
		},
		// *time.Time and types that vert to it
		{
			[]reflect.Value{v(pt0), v(pt2), v(pt1)},
			[]reflect.Value{v(pt0), v(pt1), v(pt2)},
		},
		// hybrid pointers: *time.Time
		{
			[]reflect.Value{v(pt0), v(t2), v(pt1)},
			[]reflect.Value{v(pt0), v(pt1), v(t2)},
		},
		// nil pointers: *time.Time (vention: nil < any value)
		{
			[]reflect.Value{v(pt0), v(t2), v(pt1), v(nilTimePtr)},
			[]reflect.Value{v(nilTimePtr), v(pt0), v(pt1), v(t2)},
		},
		// indirection pointers: **time.Time
		{
			[]reflect.Value{v(pt0), v(ppt2), v((nilTimePtrPtr)), v(t1)},
			[]reflect.Value{v(nilTimePtrPtr), v(pt0), v(t1), v(ppt2)},
		},
		// invalid **time.Time (nil)
		{
			[]reflect.Value{v(pt0), v(invalidTime)},
			[]reflect.Value{v(invalidTime), v(pt0)},
		},
		// Invalid.
		{
			[]reflect.Value{embedB, embedA, embedC},
			[]reflect.Value{embedB, embedA, embedC},
		},
	}
	cs := spew.ConfigState{DisableMethods: true, SpewKeys: false}

	helpTestSortValues(t, tests, &cs)
}

// TestSortValuesWithMethods ensures the sort functionality for relect.Value
// based sorting works as intended when using string methods.
func TestSortValuesWithMethods(t *testing.T) {
	v := reflect.ValueOf

	a := v("a")
	b := v("b")
	c := v("c")
	t0, t1, t2 := testTimings()
	lt0, lt1, lt2 := testTimingsWithLocation()

	tests := []sortTestCase{
		// Ints.
		{
			[]reflect.Value{v(2), v(1), v(3)},
			[]reflect.Value{v(1), v(2), v(3)},
		},
		// Strings.
		{
			[]reflect.Value{b, a, c},
			[]reflect.Value{a, b, c},
		},
		// SortableStructs.
		{
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
			[]reflect.Value{v(sortableStruct{1}), v(sortableStruct{2}), v(sortableStruct{3})},
		},
		// time.Time and types that convvert to it
		{
			[]reflect.Value{v(t0), v(t2), v(t1)},
			[]reflect.Value{v(t0), v(t1), v(t2)},
		},
		// comparison with location
		{
			[]reflect.Value{v(lt2), v(lt1), v(lt0)},
			[]reflect.Value{v(lt0), v(lt1), v(lt2)},
		},
		// UnsortableStructs.
		{
			// Note: not sorted - SpewKeys is false.
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
		},
	}
	cs := spew.ConfigState{DisableMethods: false, SpewKeys: false}
	helpTestSortValues(t, tests, &cs)
}

// TestSortValuesWithSpew ensures the sort functionality for relect.Value
// based sorting works as intended when using spew to stringify keys.
func TestSortValuesWithSpew(t *testing.T) {
	v := reflect.ValueOf

	a := v("a")
	b := v("b")
	c := v("c")
	t0, t1, t2 := testTimings()

	tests := []sortTestCase{
		// Ints.
		{
			[]reflect.Value{v(2), v(1), v(3)},
			[]reflect.Value{v(1), v(2), v(3)},
		},
		// Strings.
		{
			[]reflect.Value{b, a, c},
			[]reflect.Value{a, b, c},
		},
		// SortableStructs.
		{
			[]reflect.Value{v(sortableStruct{2}), v(sortableStruct{1}), v(sortableStruct{3})},
			[]reflect.Value{v(sortableStruct{1}), v(sortableStruct{2}), v(sortableStruct{3})},
		},
		// time.Time and types that vert to it
		{
			[]reflect.Value{v(t0), v(t2), v(t1)},
			[]reflect.Value{v(t0), v(t1), v(t2)},
		},
		// UnsortableStructs.
		{
			[]reflect.Value{v(unsortableStruct{2}), v(unsortableStruct{1}), v(unsortableStruct{3})},
			[]reflect.Value{v(unsortableStruct{1}), v(unsortableStruct{2}), v(unsortableStruct{3})},
		},
	}
	cs := spew.ConfigState{DisableMethods: true, SpewKeys: true}
	helpTestSortValues(t, tests, &cs)
}

// TestSortValuesWithString ensures the sort functionality for relect.Value
// based string sorting for times works as intended.
func TestSortTimeValuesWithString(t *testing.T) {
	v := reflect.ValueOf
	t0, t1, t2 := testTimings()

	tests := []sortTestCase{
		// time.Time and types that vert to it
		{
			[]reflect.Value{v(t0), v(t2), v(t1)},
			[]reflect.Value{v(t0), v(t1), v(t2)},
		},
	}
	cs := spew.ConfigState{DisableMethods: true, EnableTimeStringer: true}
	helpTestSortValues(t, tests, &cs)
}

// custom type to test Stinger interface on non-pointer receiver.
type stringer string

// String implements the Stringer interface for testing invocation of custom
// stringers on types with non-pointer receivers.
func (s stringer) String() string {
	return "stringer " + string(s)
}

// custom type to test Stinger interface on pointer receiver.
type pstringer string

// String implements the Stringer interface for testing invocation of custom
// stringers on types with only pointer receivers.
func (s *pstringer) String() string {
	return "stringer " + string(*s)
}

// xref1 and xref2 are cross referencing structs for testing circular reference
// detection.
type xref1 struct {
	ps2 *xref2
}
type xref2 struct {
	ps1 *xref1
}

// indirCir1, indirCir2, and indirCir3 are used to generate an indirect circular
// reference for testing detection.
type indirCir1 struct {
	ps2 *indirCir2
}
type indirCir2 struct {
	ps3 *indirCir3
}
type indirCir3 struct {
	ps1 *indirCir1
}

// embed is used to test embedded structures.
type embed struct {
	a string
}

// embedwrap is used to test embedded structures.
type embedwrap struct {
	*embed

	e *embed
}

// panicer is used to intentionally cause a panic for testing spew properly
// handles them.
type panicer int

func (p panicer) String() string {
	panic("test panic")
}

// customError is used to test custom error interface invocation.
type customError int

func (e customError) Error() string {
	return fmt.Sprintf("error: %d", int(e))
}

// stringizeWants converts a slice of wanted test output into a format suitable
// for a test error message.
func stringizeWants(wants []string) string {
	var b strings.Builder

	for i, want := range wants {
		if i > 0 {
			fmt.Fprintf(&b, "want%d: %s", i+1, want)

			continue
		}

		b.WriteString("want: " + want)
	}

	return b.String()
}

// testFailed returns whether or not a test failed by checking if the result
// of the test is in the slice of wanted strings.
func testFailed(result string, wants []string) bool {
	return !slices.Contains(wants, result)
}

type sortableStruct struct {
	x int
}

func (ss sortableStruct) String() string {
	return fmt.Sprintf("ss.%d", ss.x)
}

type unsortableStruct struct {
	x int
}

type sortTestCase struct {
	input    []reflect.Value
	expected []reflect.Value
}

func helpTestSortValues(t *testing.T, tests []sortTestCase, cs *spew.ConfigState) {
	t.Helper()

	getInterfaces := func(values []reflect.Value) []any {
		interfaces := make([]any, 0, len(values))
		for _, v := range values {
			interfaces = append(interfaces, v.Interface())
		}
		return interfaces
	}

	for _, test := range tests {
		spew.SortValues(test.input, cs)
		// reflect.DeepEqual cannot really make sense of reflect.Value,
		// probably because of all the pointer tricks. For instance,
		// v(2.0) != v(2.0) on a 32-bits system. Turn them into interface{}
		// instead.
		input := getInterfaces(test.input)
		expected := getInterfaces(test.expected)
		if !reflect.DeepEqual(input, expected) {
			t.Errorf("Sort mismatch:\n %v != %v\n\n%#v != %#v", input, expected, input, expected)
		}
	}
}

func testTimings() (t0, t1, t2 time.Time) {
	t0 = time.Now()
	t1 = t0.Add(time.Hour)
	t2 = t1.Add(time.Hour)
	return t0, t1, t2
}

func testTimingsWithLocation() (lt0, lt1, lt2 time.Time) {
	t0, t1, t2 := testTimings()
	lt0 = t0.In(time.FixedZone("UTC+5", 5))
	lt1 = t1.In(time.FixedZone("UTC-4", -4))
	lt2 = t2.In(time.FixedZone("UTC-6", -6))

	return lt0, lt1, lt2
}
