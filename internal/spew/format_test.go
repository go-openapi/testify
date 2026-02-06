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

/*
Test Summary:
NOTE: For each test, a nil pointer, a single pointer and double pointer to the
base test element are also tested to ensure proper indirection across all types.

- Max int8, int16, int32, int64, int
- Max uint8, uint16, uint32, uint64, uint
- Boolean true and false
- Standard complex64 and complex128
- Array containing standard ints
- Array containing type with custom formatter on pointer receiver only
- Array containing interfaces
- Slice containing standard float32 values
- Slice containing type with custom formatter on pointer receiver only
- Slice containing interfaces
- Nil slice
- Standard string
- Nil interface
- Sub-interface
- Map with string keys and int vals
- Map with custom formatter type on pointer receiver only keys and vals
- Map with interface keys and values
- Map with nil interface value
- Struct with primitives
- Struct that contains another struct
- Struct that contains custom type with Stringer pointer interface via both
  exported and unexported fields
- Struct that contains embedded struct and field to same struct
- Uintptr to 0 (null pointer)
- Uintptr address of real variable
- Unsafe.Pointer to 0 (null pointer)
- Unsafe.Pointer to address of real variable
- Nil channel
- Standard int channel
- Function with no params and no returns
- Function with param and no returns
- Function with multiple params and multiple returns
- Struct that is circular through self referencing
- Structs that are circular through cross referencing
- Structs that are indirectly circular
- Type that panics in its Stringer interface
- Type that has a custom Error interface
- %x passthrough with uint
- %#x passthrough with uint
- %f passthrough with precision
- %f passthrough with width and precision
- %d passthrough with width
- %q passthrough with string
*/

package spew_test

import (
	"bytes"
	"fmt"
	"iter"
	"slices"
	"testing"
	"unsafe"

	"github.com/go-openapi/testify/v2/internal/spew"
)

// TestFormatter executes all of the tests described by formatterTestCases.
func TestFormatter(t *testing.T) {
	t.Parallel()

	i := 0
	for tc := range formatterTestCases() {
		buf := new(bytes.Buffer)
		spew.Fprintf(buf, tc.format, tc.in)
		s := buf.String()
		if testFailed(s, tc.wants) {
			t.Errorf("Formatter #%d format: %s got: %s %s", i, tc.format, s,
				stringizeWants(tc.wants))
		}
		i++
	}
}

func TestPrintSortedKeys(t *testing.T) {
	t.Parallel()

	cfg := spew.ConfigState{SortKeys: true}
	s := cfg.Sprint(map[int]string{1: "1", 3: "3", 2: "2"})
	expected := "map[1:1 2:2 3:3]"
	if s != expected {
		t.Errorf("Sorted keys mismatch 1:\n  %v %v", s, expected)
	}

	s = cfg.Sprint(map[stringer]int{"1": 1, "3": 3, "2": 2})
	expected = "map[stringer 1:1 stringer 2:2 stringer 3:3]"
	if s != expected {
		t.Errorf("Sorted keys mismatch 2:\n  %v %v", s, expected)
	}

	s = cfg.Sprint(map[pstringer]int{pstringer("1"): 1, pstringer("3"): 3, pstringer("2"): 2})
	expected = "map[stringer 1:1 stringer 2:2 stringer 3:3]"
	if spew.UnsafeDisabled {
		expected = "map[1:1 2:2 3:3]"
	}
	if s != expected {
		t.Errorf("Sorted keys mismatch 3:\n  %v %v", s, expected)
	}

	s = cfg.Sprint(map[testStruct]int{{1}: 1, {3}: 3, {2}: 2})
	expected = "map[ts.1:1 ts.2:2 ts.3:3]"
	if s != expected {
		t.Errorf("Sorted keys mismatch 4:\n  %v %v", s, expected)
	}

	if !spew.UnsafeDisabled {
		s = cfg.Sprint(map[testStructP]int{{1}: 1, {3}: 3, {2}: 2})
		expected = "map[ts.1:1 ts.2:2 ts.3:3]"
		if s != expected {
			t.Errorf("Sorted keys mismatch 5:\n  %v %v", s, expected)
		}
	}

	s = cfg.Sprint(map[customError]int{customError(1): 1, customError(3): 3, customError(2): 2})
	expected = "map[error: 1:1 error: 2:2 error: 3:3]"
	if s != expected {
		t.Errorf("Sorted keys mismatch 6:\n  %v %v", s, expected)
	}
}

// formatterTest is used to describe a test to be performed against NewFormatter.
type formatterTest struct {
	format string
	in     any
	wants  []string
}

// formatterTestCases returns an iterator over all formatter test cases.
func formatterTestCases() iter.Seq[formatterTest] {
	return slices.Values(slices.Concat(
		intFormatterTests(),
		uintFormatterTests(),
		boolFormatterTests(),
		floatFormatterTests(),
		complexFormatterTests(),
		arrayFormatterTests(),
		sliceFormatterTests(),
		stringFormatterTests(),
		interfaceFormatterTests(),
		mapFormatterTests(),
		structFormatterTests(),
		uintptrFormatterTests(),
		unsafePointerFormatterTests(),
		chanFormatterTests(),
		funcFormatterTests(),
		circularFormatterTests(),
		panicFormatterTests(),
		errorFormatterTests(),
		passthroughFormatterTests(),
	))
}

//nolint:dupl // int/uint test data follows the same pattern by design
func intFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Max int8.
	v := int8(127)
	nv := (*int8)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeInt8
	vs := "127"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Max int16.
	v2 := int16(32767)
	nv2 := (*int16)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "int16"
	v2s := "32767"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	// Max int32.
	v3 := int32(2147483647)
	nv3 := (*int32)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "int32"
	v3s := "2147483647"
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3s)
	add("%v", &pv3, "<**>"+v3s)
	add("%v", nv3, "<nil>")
	add("%+v", v3, v3s)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%+v", nv3, "<nil>")
	add("%#v", v3, "("+v3t+")"+v3s)
	add("%#v", pv3, "(*"+v3t+")"+v3s)
	add("%#v", &pv3, "(**"+v3t+")"+v3s)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")
	add("%#+v", v3, "("+v3t+")"+v3s)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")

	// Max int64.
	v4 := int64(9223372036854775807)
	nv4 := (*int64)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "int64"
	v4s := "9223372036854775807"
	add("%v", v4, v4s)
	add("%v", pv4, "<*>"+v4s)
	add("%v", &pv4, "<**>"+v4s)
	add("%v", nv4, "<nil>")
	add("%+v", v4, v4s)
	add("%+v", pv4, "<*>("+v4Addr+")"+v4s)
	add("%+v", &pv4, "<**>("+pv4Addr+"->"+v4Addr+")"+v4s)
	add("%+v", nv4, "<nil>")
	add("%#v", v4, "("+v4t+")"+v4s)
	add("%#v", pv4, "(*"+v4t+")"+v4s)
	add("%#v", &pv4, "(**"+v4t+")"+v4s)
	add("%#v", nv4, "(*"+v4t+")"+"<nil>")
	add("%#+v", v4, "("+v4t+")"+v4s)
	add("%#+v", pv4, "(*"+v4t+")("+v4Addr+")"+v4s)
	add("%#+v", &pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")"+v4s)
	add("%#+v", nv4, "(*"+v4t+")"+"<nil>")

	// Max int.
	v5 := int(2147483647)
	nv5 := (*int)(nil)
	pv5 := &v5
	v5Addr := fmt.Sprintf("%p", pv5)
	pv5Addr := fmt.Sprintf("%p", &pv5)
	v5t := typeInt
	v5s := "2147483647"
	add("%v", v5, v5s)
	add("%v", pv5, "<*>"+v5s)
	add("%v", &pv5, "<**>"+v5s)
	add("%v", nv5, "<nil>")
	add("%+v", v5, v5s)
	add("%+v", pv5, "<*>("+v5Addr+")"+v5s)
	add("%+v", &pv5, "<**>("+pv5Addr+"->"+v5Addr+")"+v5s)
	add("%+v", nv5, "<nil>")
	add("%#v", v5, "("+v5t+")"+v5s)
	add("%#v", pv5, "(*"+v5t+")"+v5s)
	add("%#v", &pv5, "(**"+v5t+")"+v5s)
	add("%#v", nv5, "(*"+v5t+")"+"<nil>")
	add("%#+v", v5, "("+v5t+")"+v5s)
	add("%#+v", pv5, "(*"+v5t+")("+v5Addr+")"+v5s)
	add("%#+v", &pv5, "(**"+v5t+")("+pv5Addr+"->"+v5Addr+")"+v5s)
	add("%#+v", nv5, "(*"+v5t+")"+"<nil>")

	return tests
}

//nolint:dupl // int/uint test data follows the same pattern by design
func uintFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Max uint8.
	v := uint8(255)
	nv := (*uint8)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeUint8
	vs := "255"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Max uint16.
	v2 := uint16(65535)
	nv2 := (*uint16)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "uint16"
	v2s := "65535"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	// Max uint32.
	v3 := uint32(4294967295)
	nv3 := (*uint32)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "uint32"
	v3s := "4294967295"
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3s)
	add("%v", &pv3, "<**>"+v3s)
	add("%v", nv3, "<nil>")
	add("%+v", v3, v3s)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%+v", nv3, "<nil>")
	add("%#v", v3, "("+v3t+")"+v3s)
	add("%#v", pv3, "(*"+v3t+")"+v3s)
	add("%#v", &pv3, "(**"+v3t+")"+v3s)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")
	add("%#+v", v3, "("+v3t+")"+v3s)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")

	// Max uint64.
	v4 := uint64(18446744073709551615)
	nv4 := (*uint64)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "uint64"
	v4s := "18446744073709551615"
	add("%v", v4, v4s)
	add("%v", pv4, "<*>"+v4s)
	add("%v", &pv4, "<**>"+v4s)
	add("%v", nv4, "<nil>")
	add("%+v", v4, v4s)
	add("%+v", pv4, "<*>("+v4Addr+")"+v4s)
	add("%+v", &pv4, "<**>("+pv4Addr+"->"+v4Addr+")"+v4s)
	add("%+v", nv4, "<nil>")
	add("%#v", v4, "("+v4t+")"+v4s)
	add("%#v", pv4, "(*"+v4t+")"+v4s)
	add("%#v", &pv4, "(**"+v4t+")"+v4s)
	add("%#v", nv4, "(*"+v4t+")"+"<nil>")
	add("%#+v", v4, "("+v4t+")"+v4s)
	add("%#+v", pv4, "(*"+v4t+")("+v4Addr+")"+v4s)
	add("%#+v", &pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")"+v4s)
	add("%#+v", nv4, "(*"+v4t+")"+"<nil>")

	// Max uint.
	v5 := uint(4294967295)
	nv5 := (*uint)(nil)
	pv5 := &v5
	v5Addr := fmt.Sprintf("%p", pv5)
	pv5Addr := fmt.Sprintf("%p", &pv5)
	v5t := typeUint
	v5s := "4294967295"
	add("%v", v5, v5s)
	add("%v", pv5, "<*>"+v5s)
	add("%v", &pv5, "<**>"+v5s)
	add("%v", nv5, "<nil>")
	add("%+v", v5, v5s)
	add("%+v", pv5, "<*>("+v5Addr+")"+v5s)
	add("%+v", &pv5, "<**>("+pv5Addr+"->"+v5Addr+")"+v5s)
	add("%+v", nv5, "<nil>")
	add("%#v", v5, "("+v5t+")"+v5s)
	add("%#v", pv5, "(*"+v5t+")"+v5s)
	add("%#v", &pv5, "(**"+v5t+")"+v5s)
	add("%#v", nv5, "(*"+v5t+")"+"<nil>")
	add("%#+v", v5, "("+v5t+")"+v5s)
	add("%#+v", pv5, "(*"+v5t+")("+v5Addr+")"+v5s)
	add("%#+v", &pv5, "(**"+v5t+")("+pv5Addr+"->"+v5Addr+")"+v5s)
	add("%#v", nv5, "(*"+v5t+")"+"<nil>")

	return tests
}

func boolFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Boolean true.
	v := bool(true)
	nv := (*bool)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeBool
	vs := "true"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Boolean false.
	v2 := bool(false)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeBool
	v2s := "false"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)

	return tests
}

func floatFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Standard float32.
	v := float32(3.1415)
	nv := (*float32)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "float32"
	vs := "3.1415"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Standard float64.
	v2 := float64(3.1415926)
	nv2 := (*float64)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "float64"
	v2s := "3.1415926"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	return tests
}

func complexFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Standard complex64.
	v := complex(float32(6), -2)
	nv := (*complex64)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "complex64"
	vs := "(6-2i)"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Standard complex128.
	v2 := complex(float64(-6), 2)
	nv2 := (*complex128)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "complex128"
	v2s := "(-6+2i)"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	return tests
}

func arrayFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Array containing standard ints.
	v := [3]int{1, 2, 3}
	nv := (*[3]int)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "[3]int"
	vs := "[1 2 3]"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Array containing type with custom formatter on pointer receiver only.
	v2 := [3]pstringer{"1", "2", "3"}
	nv2 := (*[3]pstringer)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "[3]spew_test.pstringer"
	v2sp := "[stringer 1 stringer 2 stringer 3]"
	v2s := v2sp
	if spew.UnsafeDisabled {
		v2s = "[1 2 3]"
	}
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2sp)
	add("%v", &pv2, "<**>"+v2sp)
	add("%+v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2sp)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2sp)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2sp)
	add("%#v", &pv2, "(**"+v2t+")"+v2sp)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2sp)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2sp)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	// Array containing interfaces.
	v3 := [3]any{"one", int(2), uint(3)}
	nv3 := (*[3]any)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "[3]" + typeInterface
	v3t2 := typeString
	v3t3 := typeInt
	v3t4 := typeUint
	v3s := "[one 2 3]"
	v3s2 := "[(" + v3t2 + ")one (" + v3t3 + ")2 (" + v3t4 + ")3]"
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3s)
	add("%v", &pv3, "<**>"+v3s)
	add("%+v", nv3, "<nil>")
	add("%+v", v3, v3s)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%+v", nv3, "<nil>")
	add("%#v", v3, "("+v3t+")"+v3s2)
	add("%#v", pv3, "(*"+v3t+")"+v3s2)
	add("%#v", &pv3, "(**"+v3t+")"+v3s2)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")
	add("%#+v", v3, "("+v3t+")"+v3s2)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s2)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s2)
	add("%#+v", nv3, "(*"+v3t+")"+"<nil>")

	return tests
}

func sliceFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Slice containing standard float32 values.
	v := []float32{3.14, 6.28, 12.56}
	nv := (*[]float32)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "[]float32"
	vs := "[3.14 6.28 12.56]"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Slice containing type with custom formatter on pointer receiver only.
	v2 := []pstringer{"1", "2", "3"}
	nv2 := (*[]pstringer)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "[]spew_test.pstringer"
	v2s := "[stringer 1 stringer 2 stringer 3]"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	// Slice containing interfaces.
	v3 := []any{"one", int(2), uint(3), nil}
	nv3 := (*[]any)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "[]" + typeInterface
	v3t2 := typeString
	v3t3 := typeInt
	v3t4 := typeUint
	v3t5 := typeInterface
	v3s := "[one 2 3 <nil>]"
	v3s2 := "[(" + v3t2 + ")one (" + v3t3 + ")2 (" + v3t4 + ")3 (" + v3t5 +
		")<nil>]"
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3s)
	add("%v", &pv3, "<**>"+v3s)
	add("%+v", nv3, "<nil>")
	add("%+v", v3, v3s)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%+v", nv3, "<nil>")
	add("%#v", v3, "("+v3t+")"+v3s2)
	add("%#v", pv3, "(*"+v3t+")"+v3s2)
	add("%#v", &pv3, "(**"+v3t+")"+v3s2)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")
	add("%#+v", v3, "("+v3t+")"+v3s2)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s2)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s2)
	add("%#+v", nv3, "(*"+v3t+")"+"<nil>")

	// Nil slice.
	var v4 []int
	nv4 := (*[]int)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "[]int"
	v4s := typeNil
	add("%v", v4, v4s)
	add("%v", pv4, "<*>"+v4s)
	add("%v", &pv4, "<**>"+v4s)
	add("%+v", nv4, "<nil>")
	add("%+v", v4, v4s)
	add("%+v", pv4, "<*>("+v4Addr+")"+v4s)
	add("%+v", &pv4, "<**>("+pv4Addr+"->"+v4Addr+")"+v4s)
	add("%+v", nv4, "<nil>")
	add("%#v", v4, "("+v4t+")"+v4s)
	add("%#v", pv4, "(*"+v4t+")"+v4s)
	add("%#v", &pv4, "(**"+v4t+")"+v4s)
	add("%#v", nv4, "(*"+v4t+")"+"<nil>")
	add("%#+v", v4, "("+v4t+")"+v4s)
	add("%#+v", pv4, "(*"+v4t+")("+v4Addr+")"+v4s)
	add("%#+v", &pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")"+v4s)
	add("%#+v", nv4, "(*"+v4t+")"+"<nil>")

	return tests
}

func stringFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Standard string.
	v := "test"
	nv := (*string)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeString
	vs := "test"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	return tests
}

func interfaceFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Nil interface.
	var v any
	nv := (*any)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeInterface
	vs := typeNil
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Sub-interface.
	v2 := any(uint16(65535))
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "uint16"
	v2s := "65535"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)

	return tests
}

func mapFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Map with string keys and int vals.
	v := map[string]int{"one": 1, "two": 2}
	nilMap := map[string]int(nil)
	nv := (*map[string]int)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "map[string]int"
	vs := "map[one:1 two:2]"
	vs2 := "map[two:2 one:1]"
	add("%v", v, vs, vs2)
	add("%v", pv, "<*>"+vs, "<*>"+vs2)
	add("%v", &pv, "<**>"+vs, "<**>"+vs2)
	add("%+v", nilMap, "<nil>")
	add("%+v", nv, "<nil>")
	add("%+v", v, vs, vs2)
	add("%+v", pv, "<*>("+vAddr+")"+vs, "<*>("+vAddr+")"+vs2)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs,
		"<**>("+pvAddr+"->"+vAddr+")"+vs2)
	add("%+v", nilMap, "<nil>")
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs, "("+vt+")"+vs2)
	add("%#v", pv, "(*"+vt+")"+vs, "(*"+vt+")"+vs2)
	add("%#v", &pv, "(**"+vt+")"+vs, "(**"+vt+")"+vs2)
	add("%#v", nilMap, "("+vt+")"+"<nil>")
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs, "("+vt+")"+vs2)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs,
		"(*"+vt+")("+vAddr+")"+vs2)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs,
		"(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs2)
	add("%#+v", nilMap, "("+vt+")"+"<nil>")
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Map with custom formatter type on pointer receiver only keys and vals.
	v2 := map[pstringer]pstringer{"one": "1"}
	nv2 := (*map[pstringer]pstringer)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "map[spew_test.pstringer]spew_test.pstringer"
	v2s := "map[stringer one:stringer 1]"
	if spew.UnsafeDisabled {
		v2s = "map[one:1]"
	}
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	// Map with interface keys and values.
	v3 := map[any]any{"one": 1}
	nv3 := (*map[any]any)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "map[" + typeInterface + "]" + typeInterface
	v3t1 := typeString
	v3t2 := typeInt
	v3s := "map[one:1]"
	v3s2 := "map[(" + v3t1 + ")one:(" + v3t2 + ")1]"
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3s)
	add("%v", &pv3, "<**>"+v3s)
	add("%+v", nv3, "<nil>")
	add("%+v", v3, v3s)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%+v", nv3, "<nil>")
	add("%#v", v3, "("+v3t+")"+v3s2)
	add("%#v", pv3, "(*"+v3t+")"+v3s2)
	add("%#v", &pv3, "(**"+v3t+")"+v3s2)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")
	add("%#+v", v3, "("+v3t+")"+v3s2)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s2)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s2)
	add("%#+v", nv3, "(*"+v3t+")"+"<nil>")

	// Map with nil interface value
	v4 := map[string]any{"nil": nil}
	nv4 := (*map[string]any)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "map[string]" + typeInterface
	v4t1 := typeInterface
	v4s := "map[nil:<nil>]"
	v4s2 := "map[nil:(" + v4t1 + ")<nil>]"
	add("%v", v4, v4s)
	add("%v", pv4, "<*>"+v4s)
	add("%v", &pv4, "<**>"+v4s)
	add("%+v", nv4, "<nil>")
	add("%+v", v4, v4s)
	add("%+v", pv4, "<*>("+v4Addr+")"+v4s)
	add("%+v", &pv4, "<**>("+pv4Addr+"->"+v4Addr+")"+v4s)
	add("%+v", nv4, "<nil>")
	add("%#v", v4, "("+v4t+")"+v4s2)
	add("%#v", pv4, "(*"+v4t+")"+v4s2)
	add("%#v", &pv4, "(**"+v4t+")"+v4s2)
	add("%#v", nv4, "(*"+v4t+")"+"<nil>")
	add("%#+v", v4, "("+v4t+")"+v4s2)
	add("%#+v", pv4, "(*"+v4t+")("+v4Addr+")"+v4s2)
	add("%#+v", &pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")"+v4s2)
	add("%#+v", nv4, "(*"+v4t+")"+"<nil>")

	return tests
}

func structFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Struct with primitives.
	type s1 struct {
		a int8
		b uint8
	}
	v := s1{127, 255}
	nv := (*s1)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "spew_test.s1"
	vt2 := typeInt8
	vt3 := typeUint8
	vs := "{127 255}"
	vs2 := "{a:127 b:255}"
	vs3 := "{a:(" + vt2 + ")127 b:(" + vt3 + ")255}"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs2)
	add("%+v", pv, "<*>("+vAddr+")"+vs2)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs2)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs3)
	add("%#v", pv, "(*"+vt+")"+vs3)
	add("%#v", &pv, "(**"+vt+")"+vs3)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs3)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs3)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs3)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Struct that contains another struct.
	type s2 struct {
		s1 s1
		b  bool
	}
	v2 := s2{s1{127, 255}, true}
	nv2 := (*s2)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "spew_test.s2"
	v2t2 := "spew_test.s1"
	v2t3 := typeInt8
	v2t4 := typeUint8
	v2t5 := typeBool
	v2s := "{{127 255} true}"
	v2s2 := "{s1:{a:127 b:255} b:true}"
	v2s3 := "{s1:(" + v2t2 + "){a:(" + v2t3 + ")127 b:(" + v2t4 + ")255} b:(" +
		v2t5 + ")true}"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", nv2, "<nil>")
	add("%+v", v2, v2s2)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s2)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s2)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s3)
	add("%#v", pv2, "(*"+v2t+")"+v2s3)
	add("%#v", &pv2, "(**"+v2t+")"+v2s3)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s3)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s3)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s3)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	// Struct that contains custom type with Stringer pointer interface via both
	// exported and unexported fields.
	type s3 struct {
		s pstringer
		S pstringer
	}
	v3 := s3{"test", "test2"}
	nv3 := (*s3)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "spew_test.s3"
	v3t2 := "spew_test.pstringer"
	v3s := "{stringer test stringer test2}"
	v3sp := v3s
	v3s2 := "{s:stringer test S:stringer test2}"
	v3s2p := v3s2
	v3s3 := "{s:(" + v3t2 + ")stringer test S:(" + v3t2 + ")stringer test2}"
	v3s3p := v3s3
	if spew.UnsafeDisabled {
		v3s = "{test test2}"
		v3sp = "{test stringer test2}"
		v3s2 = "{s:test S:test2}"
		v3s2p = "{s:test S:stringer test2}"
		v3s3 = "{s:(" + v3t2 + ")test S:(" + v3t2 + ")test2}"
		v3s3p = "{s:(" + v3t2 + ")test S:(" + v3t2 + ")stringer test2}"
	}
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3sp)
	add("%v", &pv3, "<**>"+v3sp)
	add("%+v", nv3, "<nil>")
	add("%+v", v3, v3s2)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s2p)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s2p)
	add("%+v", nv3, "<nil>")
	add("%#v", v3, "("+v3t+")"+v3s3)
	add("%#v", pv3, "(*"+v3t+")"+v3s3p)
	add("%#v", &pv3, "(**"+v3t+")"+v3s3p)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")
	add("%#+v", v3, "("+v3t+")"+v3s3)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s3p)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s3p)
	add("%#+v", nv3, "(*"+v3t+")"+"<nil>")

	// Struct that contains embedded struct and field to same struct.
	e := embed{"embedstr"}
	v4 := embedwrap{embed: &e, e: &e}
	nv4 := (*embedwrap)(nil)
	pv4 := &v4
	eAddr := fmt.Sprintf("%p", &e)
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "spew_test.embedwrap"
	v4t2 := "spew_test.embed"
	v4t3 := typeString
	v4s := "{<*>{embedstr} <*>{embedstr}}"
	v4s2 := "{embed:<*>(" + eAddr + "){a:embedstr} e:<*>(" + eAddr +
		"){a:embedstr}}"
	v4s3 := "{embed:(*" + v4t2 + "){a:(" + v4t3 + ")embedstr} e:(*" + v4t2 +
		"){a:(" + v4t3 + ")embedstr}}"
	v4s4 := "{embed:(*" + v4t2 + ")(" + eAddr + "){a:(" + v4t3 +
		")embedstr} e:(*" + v4t2 + ")(" + eAddr + "){a:(" + v4t3 + ")embedstr}}"
	add("%v", v4, v4s)
	add("%v", pv4, "<*>"+v4s)
	add("%v", &pv4, "<**>"+v4s)
	add("%+v", nv4, "<nil>")
	add("%+v", v4, v4s2)
	add("%+v", pv4, "<*>("+v4Addr+")"+v4s2)
	add("%+v", &pv4, "<**>("+pv4Addr+"->"+v4Addr+")"+v4s2)
	add("%+v", nv4, "<nil>")
	add("%#v", v4, "("+v4t+")"+v4s3)
	add("%#v", pv4, "(*"+v4t+")"+v4s3)
	add("%#v", &pv4, "(**"+v4t+")"+v4s3)
	add("%#v", nv4, "(*"+v4t+")"+"<nil>")
	add("%#+v", v4, "("+v4t+")"+v4s4)
	add("%#+v", pv4, "(*"+v4t+")("+v4Addr+")"+v4s4)
	add("%#+v", &pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")"+v4s4)
	add("%#+v", nv4, "(*"+v4t+")"+"<nil>")

	return tests
}

func uintptrFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Null pointer.
	v := uintptr(0)
	nv := (*uintptr)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "uintptr"
	vs := typeNil
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Address of real variable.
	i := 1
	v2 := uintptr(unsafe.Pointer(&i))
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "uintptr"
	v2s := fmt.Sprintf("%p", &i)
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)

	return tests
}

func unsafePointerFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Null pointer.
	v := unsafe.Pointer(nil)
	nv := (*unsafe.Pointer)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "unsafe.Pointer"
	vs := typeNil
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Address of real variable.
	i := 1
	v2 := unsafe.Pointer(&i)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "unsafe.Pointer"
	v2s := fmt.Sprintf("%p", &i)
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)

	return tests
}

func chanFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Nil channel.
	var v chan int
	pv := &v
	nv := (*chan int)(nil)
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "chan int"
	vs := typeNil
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Real channel.
	v2 := make(chan int)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "chan int"
	v2s := fmt.Sprintf("%p", v2)
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)

	return tests
}

func funcFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Function with no params and no returns.
	v := intFormatterTests
	nv := (*func() []formatterTest)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "func() []spew_test.formatterTest"
	vs := fmt.Sprintf("%p", v)
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%+v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	// Function with param and no returns.
	v2 := TestFormatter
	nv2 := (*func(*testing.T))(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "func(*testing.T)"
	v2s := fmt.Sprintf("%p", v2)
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s)
	add("%v", &pv2, "<**>"+v2s)
	add("%+v", nv2, "<nil>")
	add("%+v", v2, v2s)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%+v", nv2, "<nil>")
	add("%#v", v2, "("+v2t+")"+v2s)
	add("%#v", pv2, "(*"+v2t+")"+v2s)
	add("%#v", &pv2, "(**"+v2t+")"+v2s)
	add("%#v", nv2, "(*"+v2t+")"+"<nil>")
	add("%#+v", v2, "("+v2t+")"+v2s)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s)
	add("%#+v", nv2, "(*"+v2t+")"+"<nil>")

	// Function with multiple params and multiple returns.
	v3 := func(_ int, _ string) (b bool, err error) {
		return true, nil
	}
	nv3 := (*func(int, string) (bool, error))(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "func(int, string) (bool, error)"
	v3s := fmt.Sprintf("%p", v3)
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3s)
	add("%v", &pv3, "<**>"+v3s)
	add("%+v", nv3, "<nil>")
	add("%+v", v3, v3s)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%+v", nv3, "<nil>")
	add("%#v", v3, "("+v3t+")"+v3s)
	add("%#v", pv3, "(*"+v3t+")"+v3s)
	add("%#v", &pv3, "(**"+v3t+")"+v3s)
	add("%#v", nv3, "(*"+v3t+")"+"<nil>")
	add("%#+v", v3, "("+v3t+")"+v3s)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s)
	add("%#+v", nv3, "(*"+v3t+")"+"<nil>")

	return tests
}

func circularFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Struct that is circular through self referencing.
	type circular struct {
		c *circular
	}
	v := circular{nil}
	v.c = &v
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "spew_test.circular"
	vs := "{<*>{<*><shown>}}"
	vs2 := "{<*><shown>}"
	vs3 := "{c:<*>(" + vAddr + "){c:<*>(" + vAddr + ")<shown>}}"
	vs4 := "{c:<*>(" + vAddr + ")<shown>}"
	vs5 := "{c:(*" + vt + "){c:(*" + vt + ")<shown>}}"
	vs6 := "{c:(*" + vt + ")<shown>}"
	vs7 := "{c:(*" + vt + ")(" + vAddr + "){c:(*" + vt + ")(" + vAddr +
		")<shown>}}"
	vs8 := "{c:(*" + vt + ")(" + vAddr + ")<shown>}"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs2)
	add("%v", &pv, "<**>"+vs2)
	add("%+v", v, vs3)
	add("%+v", pv, "<*>("+vAddr+")"+vs4)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs4)
	add("%#v", v, "("+vt+")"+vs5)
	add("%#v", pv, "(*"+vt+")"+vs6)
	add("%#v", &pv, "(**"+vt+")"+vs6)
	add("%#+v", v, "("+vt+")"+vs7)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs8)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs8)

	// Structs that are circular through cross referencing.
	v2 := xref1{nil}
	ts2 := xref2{&v2}
	v2.ps2 = &ts2
	pv2 := &v2
	ts2Addr := fmt.Sprintf("%p", &ts2)
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "spew_test.xref1"
	v2t2 := "spew_test.xref2"
	v2s := "{<*>{<*>{<*><shown>}}}"
	v2s2 := "{<*>{<*><shown>}}"
	v2s3 := "{ps2:<*>(" + ts2Addr + "){ps1:<*>(" + v2Addr + "){ps2:<*>(" +
		ts2Addr + ")<shown>}}}"
	v2s4 := "{ps2:<*>(" + ts2Addr + "){ps1:<*>(" + v2Addr + ")<shown>}}"
	v2s5 := "{ps2:(*" + v2t2 + "){ps1:(*" + v2t + "){ps2:(*" + v2t2 +
		")<shown>}}}"
	v2s6 := "{ps2:(*" + v2t2 + "){ps1:(*" + v2t + ")<shown>}}"
	v2s7 := "{ps2:(*" + v2t2 + ")(" + ts2Addr + "){ps1:(*" + v2t +
		")(" + v2Addr + "){ps2:(*" + v2t2 + ")(" + ts2Addr +
		")<shown>}}}"
	v2s8 := "{ps2:(*" + v2t2 + ")(" + ts2Addr + "){ps1:(*" + v2t +
		")(" + v2Addr + ")<shown>}}"
	add("%v", v2, v2s)
	add("%v", pv2, "<*>"+v2s2)
	add("%v", &pv2, "<**>"+v2s2)
	add("%+v", v2, v2s3)
	add("%+v", pv2, "<*>("+v2Addr+")"+v2s4)
	add("%+v", &pv2, "<**>("+pv2Addr+"->"+v2Addr+")"+v2s4)
	add("%#v", v2, "("+v2t+")"+v2s5)
	add("%#v", pv2, "(*"+v2t+")"+v2s6)
	add("%#v", &pv2, "(**"+v2t+")"+v2s6)
	add("%#+v", v2, "("+v2t+")"+v2s7)
	add("%#+v", pv2, "(*"+v2t+")("+v2Addr+")"+v2s8)
	add("%#+v", &pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")"+v2s8)

	// Structs that are indirectly circular.
	v3 := indirCir1{nil}
	tic2 := indirCir2{nil}
	tic3 := indirCir3{&v3}
	tic2.ps3 = &tic3
	v3.ps2 = &tic2
	pv3 := &v3
	tic2Addr := fmt.Sprintf("%p", &tic2)
	tic3Addr := fmt.Sprintf("%p", &tic3)
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "spew_test.indirCir1"
	v3t2 := "spew_test.indirCir2"
	v3t3 := "spew_test.indirCir3"
	v3s := "{<*>{<*>{<*>{<*><shown>}}}}"
	v3s2 := "{<*>{<*>{<*><shown>}}}"
	v3s3 := "{ps2:<*>(" + tic2Addr + "){ps3:<*>(" + tic3Addr + "){ps1:<*>(" +
		v3Addr + "){ps2:<*>(" + tic2Addr + ")<shown>}}}}"
	v3s4 := "{ps2:<*>(" + tic2Addr + "){ps3:<*>(" + tic3Addr + "){ps1:<*>(" +
		v3Addr + ")<shown>}}}"
	v3s5 := "{ps2:(*" + v3t2 + "){ps3:(*" + v3t3 + "){ps1:(*" + v3t +
		"){ps2:(*" + v3t2 + ")<shown>}}}}"
	v3s6 := "{ps2:(*" + v3t2 + "){ps3:(*" + v3t3 + "){ps1:(*" + v3t +
		")<shown>}}}"
	v3s7 := "{ps2:(*" + v3t2 + ")(" + tic2Addr + "){ps3:(*" + v3t3 + ")(" +
		tic3Addr + "){ps1:(*" + v3t + ")(" + v3Addr + "){ps2:(*" + v3t2 +
		")(" + tic2Addr + ")<shown>}}}}"
	v3s8 := "{ps2:(*" + v3t2 + ")(" + tic2Addr + "){ps3:(*" + v3t3 + ")(" +
		tic3Addr + "){ps1:(*" + v3t + ")(" + v3Addr + ")<shown>}}}"
	add("%v", v3, v3s)
	add("%v", pv3, "<*>"+v3s2)
	add("%v", &pv3, "<**>"+v3s2)
	add("%+v", v3, v3s3)
	add("%+v", pv3, "<*>("+v3Addr+")"+v3s4)
	add("%+v", &pv3, "<**>("+pv3Addr+"->"+v3Addr+")"+v3s4)
	add("%#v", v3, "("+v3t+")"+v3s5)
	add("%#v", pv3, "(*"+v3t+")"+v3s6)
	add("%#v", &pv3, "(**"+v3t+")"+v3s6)
	add("%#+v", v3, "("+v3t+")"+v3s7)
	add("%#+v", pv3, "(*"+v3t+")("+v3Addr+")"+v3s8)
	add("%#+v", &pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")"+v3s8)

	return tests
}

//nolint:dupl // panic/error test data follows the same pattern by design
func panicFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Type that panics in its Stringer interface.
	v := panicer(127)
	nv := (*panicer)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "spew_test.panicer"
	vs := "(PANIC=test panic)127"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")

	return tests
}

//nolint:dupl // panic/error test data follows the same pattern by design
func errorFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// Type that has a custom Error interface.
	v := customError(127)
	nv := (*customError)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "spew_test.customError"
	vs := "error: 127"
	add("%v", v, vs)
	add("%v", pv, "<*>"+vs)
	add("%v", &pv, "<**>"+vs)
	add("%v", nv, "<nil>")
	add("%+v", v, vs)
	add("%+v", pv, "<*>("+vAddr+")"+vs)
	add("%+v", &pv, "<**>("+pvAddr+"->"+vAddr+")"+vs)
	add("%+v", nv, "<nil>")
	add("%#v", v, "("+vt+")"+vs)
	add("%#v", pv, "(*"+vt+")"+vs)
	add("%#v", &pv, "(**"+vt+")"+vs)
	add("%#v", nv, "(*"+vt+")"+"<nil>")
	add("%#+v", v, "("+vt+")"+vs)
	add("%#+v", pv, "(*"+vt+")("+vAddr+")"+vs)
	add("%#+v", &pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")"+vs)
	add("%#+v", nv, "(*"+vt+")"+"<nil>")
	return tests
}

func passthroughFormatterTests() []formatterTest {
	var tests []formatterTest
	add := func(format string, in any, wants ...string) {
		tests = append(tests, formatterTest{format, in, wants})
	}

	// %x passthrough with uint.
	v := uint(4294967295)
	pv := &v
	vAddr := fmt.Sprintf("%x", pv)
	pvAddr := fmt.Sprintf("%x", &pv)
	vs := "ffffffff"
	add("%x", v, vs)
	add("%x", pv, vAddr)
	add("%x", &pv, pvAddr)

	// %#x passthrough with uint.
	v2 := int(2147483647)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%#x", pv2)
	pv2Addr := fmt.Sprintf("%#x", &pv2)
	v2s := "0x7fffffff"
	add("%#x", v2, v2s)
	add("%#x", pv2, v2Addr)
	add("%#x", &pv2, pv2Addr)

	// %f passthrough with precision.
	add("%.2f", 3.1415, "3.14")
	add("%.3f", 3.1415, "3.142")
	add("%.4f", 3.1415, "3.1415")

	// %f passthrough with width and precision.
	add("%5.2f", 3.1415, " 3.14")
	add("%6.3f", 3.1415, " 3.142")
	add("%7.4f", 3.1415, " 3.1415")

	// %d passthrough with width.
	add("%3d", 127, "127")
	add("%4d", 127, " 127")
	add("%5d", 127, "  127")

	// %q passthrough with string.
	add("%q", "test", "\"test\"")

	return tests
}

// =================================
// Types for TestPrintSortedKeys
// =================================

type testStruct struct {
	x int
}

func (ts testStruct) String() string {
	return fmt.Sprintf("ts.%d", ts.x)
}

type testStructP struct {
	x int
}

func (ts *testStructP) String() string {
	return fmt.Sprintf("ts.%d", ts.x)
}
