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
- Array containing bytes
- Slice containing standard float32 values
- Slice containing type with custom formatter on pointer receiver only
- Slice containing interfaces
- Slice containing bytes
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
*/

package spew_test

import (
	"bytes"
	"fmt"
	"iter"
	"slices"
	"strconv"
	"testing"
	"time"
	"unsafe"

	"github.com/go-openapi/testify/v2/internal/spew"
)

// Type name constants used in test expectations.
const (
	typeInt8          = "int8"
	typeInt16         = "int16"
	typeInt32         = "int32"
	typeInt64         = "int64"
	typeInt           = "int"
	typeUint8         = "uint8"
	typeUint16        = "uint16"
	typeUint32        = "uint32"
	typeUint64        = "uint64"
	typeUint          = "uint"
	typeBool          = "bool"
	typeFloat32       = "float32"
	typeFloat64       = "float64"
	typeComplex64     = "complex64"
	typeComplex128    = "complex128"
	typeString        = "string"
	typeInterface     = "interface {}"
	typeUintptr       = "uintptr"
	typeUnsafePointer = "unsafe.Pointer"
	typeChanInt       = "chan int"
	typePstringer     = "spew_test.pstringer"
	typeS1            = "spew_test.s1"
	typeS2            = "spew_test.s2"
	typeNil           = "<nil>"
	valOne            = "one"
	valTest           = "test"
)

// TestDump executes all of the tests described by dumpTestCases.
func TestDump(t *testing.T) {
	t.Parallel()

	i := 0
	for tc := range dumpTestCases() {
		buf := new(bytes.Buffer)
		spew.Fdump(buf, tc.in)
		s := buf.String()
		if testFailed(s, tc.wants) {
			t.Errorf("Dump #%d\n  got: %s %s", i, s, stringizeWants(tc.wants))
		}
		i++
	}
}

func TestDumpSortedKeys(t *testing.T) {
	t.Parallel()

	cfg := spew.ConfigState{SortKeys: true}
	s := cfg.Sdump(map[int]string{1: "1", 3: "3", 2: "2"})
	expected := "(map[int]string) (len=3) {\n(int) 1: (string) (len=1) " +
		"\"1\",\n(int) 2: (string) (len=1) \"2\",\n(int) 3: (string) " +
		"(len=1) \"3\"\n" +
		"}\n"
	if s != expected {
		t.Errorf("Sorted keys mismatch:\n  %v %v", s, expected)
	}

	s = cfg.Sdump(map[stringer]int{"1": 1, "3": 3, "2": 2})
	expected = "(map[spew_test.stringer]int) (len=3) {\n" +
		"(spew_test.stringer) (len=1) stringer 1: (int) 1,\n" +
		"(spew_test.stringer) (len=1) stringer 2: (int) 2,\n" +
		"(spew_test.stringer) (len=1) stringer 3: (int) 3\n" +
		"}\n"
	if s != expected {
		t.Errorf("Sorted keys mismatch:\n  %v %v", s, expected)
	}

	s = cfg.Sdump(map[pstringer]int{pstringer("1"): 1, pstringer("3"): 3, pstringer("2"): 2})
	expected = "(map[spew_test.pstringer]int) (len=3) {\n" +
		"(spew_test.pstringer) (len=1) stringer 1: (int) 1,\n" +
		"(spew_test.pstringer) (len=1) stringer 2: (int) 2,\n" +
		"(spew_test.pstringer) (len=1) stringer 3: (int) 3\n" +
		"}\n"
	if spew.UnsafeDisabled {
		expected = "(map[spew_test.pstringer]int) (len=3) {\n" +
			"(spew_test.pstringer) (len=1) \"1\": (int) 1,\n" +
			"(spew_test.pstringer) (len=1) \"2\": (int) 2,\n" +
			"(spew_test.pstringer) (len=1) \"3\": (int) 3\n" +
			"}\n"
	}
	if s != expected {
		t.Errorf("Sorted keys mismatch:\n  %v %v", s, expected)
	}

	s = cfg.Sdump(map[customError]int{customError(1): 1, customError(3): 3, customError(2): 2})
	expected = "(map[spew_test.customError]int) (len=3) {\n" +
		"(spew_test.customError) error: 1: (int) 1,\n" +
		"(spew_test.customError) error: 2: (int) 2,\n" +
		"(spew_test.customError) error: 3: (int) 3\n" +
		"}\n"
	if s != expected {
		t.Errorf("Sorted keys mismatch:\n  %v %v", s, expected)
	}
}

// =======================
// TestDump test cases
// =======================

// dumpTest is used to describe a test to be performed against the Dump method.
type dumpTest struct {
	in    any
	wants []string
}

// dumpTestCases returns an iterator over all dump test cases.
func dumpTestCases() iter.Seq[dumpTest] {
	return slices.Values(slices.Concat(
		intDumpTests(),
		uintDumpTests(),
		boolDumpTests(),
		floatDumpTests(),
		complexDumpTests(),
		arrayDumpTests(),
		sliceDumpTests(),
		stringDumpTests(),
		interfaceDumpTests(),
		mapDumpTests(),
		structDumpTests(),
		uintptrDumpTests(),
		unsafePointerDumpTests(),
		chanDumpTests(),
		funcDumpTests(),
		circularDumpTests(),
		panicDumpTests(),
		errorDumpTests(),
		cgoDumpTests(),
		timeDumpTests(),
	))
}

func ptr[T any](value T) *T {
	v := value
	return &v
}

//nolint:dupl // int/uint test data follows the same pattern by design
func intDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Max int8.
	v := int8(127)
	nv := (*int8)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeInt8
	vs := "127"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Max int16.
	v2 := int16(32767)
	nv2 := (*int16)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeInt16
	v2s := "32767"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*"+v2t+")(<nil>)\n")

	// Max int32.
	v3 := int32(2147483647)
	nv3 := (*int32)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := typeInt32
	v3s := "2147483647"
	add(v3, "("+v3t+") "+v3s+"\n")
	add(pv3, "(*"+v3t+")("+v3Addr+")("+v3s+")\n")
	add(&pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")("+v3s+")\n")
	add(nv3, "(*"+v3t+")(<nil>)\n")

	// Max int64.
	v4 := int64(9223372036854775807)
	nv4 := (*int64)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := typeInt64
	v4s := "9223372036854775807"
	add(v4, "("+v4t+") "+v4s+"\n")
	add(pv4, "(*"+v4t+")("+v4Addr+")("+v4s+")\n")
	add(&pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")("+v4s+")\n")
	add(nv4, "(*"+v4t+")(<nil>)\n")

	// Max int.
	v5 := int(2147483647)
	nv5 := (*int)(nil)
	pv5 := &v5
	v5Addr := fmt.Sprintf("%p", pv5)
	pv5Addr := fmt.Sprintf("%p", &pv5)
	v5t := typeInt
	v5s := "2147483647"
	add(v5, "("+v5t+") "+v5s+"\n")
	add(pv5, "(*"+v5t+")("+v5Addr+")("+v5s+")\n")
	add(&pv5, "(**"+v5t+")("+pv5Addr+"->"+v5Addr+")("+v5s+")\n")
	add(nv5, "(*"+v5t+")(<nil>)\n")

	return tests
}

//nolint:dupl // int/uint test data follows the same pattern by design
func uintDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Max uint8.
	v := uint8(255)
	nv := (*uint8)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeUint8
	vs := "255"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Max uint16.
	v2 := uint16(65535)
	nv2 := (*uint16)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeUint16
	v2s := "65535"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*"+v2t+")(<nil>)\n")

	// Max uint32.
	v3 := uint32(4294967295)
	nv3 := (*uint32)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := typeUint32
	v3s := "4294967295"
	add(v3, "("+v3t+") "+v3s+"\n")
	add(pv3, "(*"+v3t+")("+v3Addr+")("+v3s+")\n")
	add(&pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")("+v3s+")\n")
	add(nv3, "(*"+v3t+")(<nil>)\n")

	// Max uint64.
	v4 := uint64(18446744073709551615)
	nv4 := (*uint64)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := typeUint64
	v4s := "18446744073709551615"
	add(v4, "("+v4t+") "+v4s+"\n")
	add(pv4, "(*"+v4t+")("+v4Addr+")("+v4s+")\n")
	add(&pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")("+v4s+")\n")
	add(nv4, "(*"+v4t+")(<nil>)\n")

	// Max uint.
	v5 := uint(4294967295)
	nv5 := (*uint)(nil)
	pv5 := &v5
	v5Addr := fmt.Sprintf("%p", pv5)
	pv5Addr := fmt.Sprintf("%p", &pv5)
	v5t := typeUint
	v5s := "4294967295"
	add(v5, "("+v5t+") "+v5s+"\n")
	add(pv5, "(*"+v5t+")("+v5Addr+")("+v5s+")\n")
	add(&pv5, "(**"+v5t+")("+pv5Addr+"->"+v5Addr+")("+v5s+")\n")
	add(nv5, "(*"+v5t+")(<nil>)\n")

	return tests
}

func boolDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Boolean true.
	v := bool(true)
	nv := (*bool)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeBool
	vs := "true"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Boolean false.
	v2 := bool(false)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeBool
	v2s := "false"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")

	return tests
}

func floatDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Standard float32.
	v := float32(3.1415)
	nv := (*float32)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeFloat32
	vs := "3.1415"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Standard float64.
	v2 := float64(3.1415926)
	nv2 := (*float64)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeFloat64
	v2s := "3.1415926"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*"+v2t+")(<nil>)\n")

	return tests
}

func complexDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Standard complex64.
	v := complex(float32(6), -2)
	nv := (*complex64)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeComplex64
	vs := "(6-2i)"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Standard complex128.
	v2 := complex(float64(-6), 2)
	nv2 := (*complex128)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeComplex128
	v2s := "(-6+2i)"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*"+v2t+")(<nil>)\n")

	return tests
}

func arrayDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Array containing standard ints.
	v := [3]int{1, 2, 3}
	vLen := strconv.Itoa(len(v))
	vCap := strconv.Itoa(cap(v))
	nv := (*[3]int)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeInt
	vs := "(len=" + vLen + " cap=" + vCap + ") {\n (" + vt + ") 1,\n (" +
		vt + ") 2,\n (" + vt + ") 3\n}"
	add(v, "([3]"+vt+") "+vs+"\n")
	add(pv, "(*[3]"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**[3]"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*[3]"+vt+")(<nil>)\n")

	// Array containing type with custom formatter on pointer receiver only.
	v2i0 := pstringer("1")
	v2i1 := pstringer("2")
	v2i2 := pstringer("3")
	v2 := [3]pstringer{v2i0, v2i1, v2i2}
	v2i0Len := strconv.Itoa(len(v2i0))
	v2i1Len := strconv.Itoa(len(v2i1))
	v2i2Len := strconv.Itoa(len(v2i2))
	v2Len := strconv.Itoa(len(v2))
	v2Cap := strconv.Itoa(cap(v2))
	nv2 := (*[3]pstringer)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typePstringer
	v2sp := "(len=" + v2Len + " cap=" + v2Cap + ") {\n (" + v2t +
		") (len=" + v2i0Len + ") stringer 1,\n (" + v2t +
		") (len=" + v2i1Len + ") stringer 2,\n (" + v2t +
		") (len=" + v2i2Len + ") " + "stringer 3\n}"
	v2s := v2sp
	if spew.UnsafeDisabled {
		v2s = "(len=" + v2Len + " cap=" + v2Cap + ") {\n (" + v2t +
			") (len=" + v2i0Len + ") \"1\",\n (" + v2t + ") (len=" +
			v2i1Len + ") \"2\",\n (" + v2t + ") (len=" + v2i2Len +
			") " + "\"3\"\n}"
	}
	add(v2, "([3]"+v2t+") "+v2s+"\n")
	add(pv2, "(*[3]"+v2t+")("+v2Addr+")("+v2sp+")\n")
	add(&pv2, "(**[3]"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2sp+")\n")
	add(nv2, "(*[3]"+v2t+")(<nil>)\n")

	// Array containing interfaces.
	v3i0 := valOne
	v3 := [3]any{v3i0, int(2), uint(3)}
	v3i0Len := strconv.Itoa(len(v3i0))
	v3Len := strconv.Itoa(len(v3))
	v3Cap := strconv.Itoa(cap(v3))
	nv3 := (*[3]any)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "[3]" + typeInterface
	v3t2 := typeString
	v3t3 := typeInt
	v3t4 := typeUint
	v3s := "(len=" + v3Len + " cap=" + v3Cap + ") {\n (" + v3t2 + ") " +
		"(len=" + v3i0Len + ") \"one\",\n (" + v3t3 + ") 2,\n (" +
		v3t4 + ") 3\n}"
	add(v3, "("+v3t+") "+v3s+"\n")
	add(pv3, "(*"+v3t+")("+v3Addr+")("+v3s+")\n")
	add(&pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")("+v3s+")\n")
	add(nv3, "(*"+v3t+")(<nil>)\n")

	// Array containing bytes.
	v4 := [34]byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30,
		0x31, 0x32,
	}
	v4Len := strconv.Itoa(len(v4))
	v4Cap := strconv.Itoa(cap(v4))
	nv4 := (*[34]byte)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "[34]" + typeUint8
	v4s := "(len=" + v4Len + " cap=" + v4Cap + ") " +
		"{\n 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20" +
		"  |............... |\n" +
		" 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30" +
		"  |!\"#$%&'()*+,-./0|\n" +
		" 00000020  31 32                                           " +
		"  |12|\n}"
	add(v4, "("+v4t+") "+v4s+"\n")
	add(pv4, "(*"+v4t+")("+v4Addr+")("+v4s+")\n")
	add(&pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")("+v4s+")\n")
	add(nv4, "(*"+v4t+")(<nil>)\n")

	return tests
}

func sliceDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Slice containing standard float32 values.
	v := []float32{3.14, 6.28, 12.56}
	vLen := strconv.Itoa(len(v))
	vCap := strconv.Itoa(cap(v))
	nv := (*[]float32)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeFloat32
	vs := "(len=" + vLen + " cap=" + vCap + ") {\n (" + vt + ") 3.14,\n (" +
		vt + ") 6.28,\n (" + vt + ") 12.56\n}"
	add(v, "([]"+vt+") "+vs+"\n")
	add(pv, "(*[]"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**[]"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*[]"+vt+")(<nil>)\n")

	// Slice containing type with custom formatter on pointer receiver only.
	v2i0 := pstringer("1")
	v2i1 := pstringer("2")
	v2i2 := pstringer("3")
	v2 := []pstringer{v2i0, v2i1, v2i2}
	v2i0Len := strconv.Itoa(len(v2i0))
	v2i1Len := strconv.Itoa(len(v2i1))
	v2i2Len := strconv.Itoa(len(v2i2))
	v2Len := strconv.Itoa(len(v2))
	v2Cap := strconv.Itoa(cap(v2))
	nv2 := (*[]pstringer)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typePstringer
	v2s := "(len=" + v2Len + " cap=" + v2Cap + ") {\n (" + v2t + ") (len=" +
		v2i0Len + ") stringer 1,\n (" + v2t + ") (len=" + v2i1Len +
		") stringer 2,\n (" + v2t + ") (len=" + v2i2Len + ") " +
		"stringer 3\n}"
	add(v2, "([]"+v2t+") "+v2s+"\n")
	add(pv2, "(*[]"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**[]"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*[]"+v2t+")(<nil>)\n")

	// Slice containing interfaces.
	v3i0 := valOne
	v3 := []any{v3i0, int(2), uint(3), nil}
	v3i0Len := strconv.Itoa(len(v3i0))
	v3Len := strconv.Itoa(len(v3))
	v3Cap := strconv.Itoa(cap(v3))
	nv3 := (*[]any)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "[]" + typeInterface
	v3t2 := typeString
	v3t3 := typeInt
	v3t4 := typeUint
	v3t5 := typeInterface
	v3s := "(len=" + v3Len + " cap=" + v3Cap + ") {\n (" + v3t2 + ") " +
		"(len=" + v3i0Len + ") \"one\",\n (" + v3t3 + ") 2,\n (" +
		v3t4 + ") 3,\n (" + v3t5 + ") <nil>\n}"
	add(v3, "("+v3t+") "+v3s+"\n")
	add(pv3, "(*"+v3t+")("+v3Addr+")("+v3s+")\n")
	add(&pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")("+v3s+")\n")
	add(nv3, "(*"+v3t+")(<nil>)\n")

	// Slice containing bytes.
	v4 := []byte{
		0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18,
		0x19, 0x1a, 0x1b, 0x1c, 0x1d, 0x1e, 0x1f, 0x20,
		0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28,
		0x29, 0x2a, 0x2b, 0x2c, 0x2d, 0x2e, 0x2f, 0x30,
		0x31, 0x32,
	}
	v4Len := strconv.Itoa(len(v4))
	v4Cap := strconv.Itoa(cap(v4))
	nv4 := (*[]byte)(nil)
	pv4 := &v4
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "[]" + typeUint8
	v4s := "(len=" + v4Len + " cap=" + v4Cap + ") " +
		"{\n 00000000  11 12 13 14 15 16 17 18  19 1a 1b 1c 1d 1e 1f 20" +
		"  |............... |\n" +
		" 00000010  21 22 23 24 25 26 27 28  29 2a 2b 2c 2d 2e 2f 30" +
		"  |!\"#$%&'()*+,-./0|\n" +
		" 00000020  31 32                                           " +
		"  |12|\n}"
	add(v4, "("+v4t+") "+v4s+"\n")
	add(pv4, "(*"+v4t+")("+v4Addr+")("+v4s+")\n")
	add(&pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")("+v4s+")\n")
	add(nv4, "(*"+v4t+")(<nil>)\n")

	// Nil slice.
	v5 := []int(nil)
	nv5 := (*[]int)(nil)
	pv5 := &v5
	v5Addr := fmt.Sprintf("%p", pv5)
	pv5Addr := fmt.Sprintf("%p", &pv5)
	v5t := "[]" + typeInt
	v5s := typeNil
	add(v5, "("+v5t+") "+v5s+"\n")
	add(pv5, "(*"+v5t+")("+v5Addr+")("+v5s+")\n")
	add(&pv5, "(**"+v5t+")("+pv5Addr+"->"+v5Addr+")("+v5s+")\n")
	add(nv5, "(*"+v5t+")(<nil>)\n")

	return tests
}

func stringDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Standard string.
	v := valTest
	vLen := strconv.Itoa(len(v))
	nv := (*string)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeString
	vs := "(len=" + vLen + ") \"test\""
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	return tests
}

func interfaceDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Nil interface.
	var v any
	nv := (*any)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeInterface
	vs := typeNil
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Sub-interface.
	v2 := any(uint16(65535))
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeUint16
	v2s := "65535"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")

	return tests
}

func mapDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Map with string keys and int vals.
	k := valOne
	kk := "two"
	m := map[string]int{k: 1, kk: 2}
	klen := strconv.Itoa(len(k)) // not kLen to shut golint up
	kkLen := strconv.Itoa(len(kk))
	mLen := strconv.Itoa(len(m))
	nilMap := map[string]int(nil)
	nm := (*map[string]int)(nil)
	pm := &m
	mAddr := fmt.Sprintf("%p", pm)
	pmAddr := fmt.Sprintf("%p", &pm)
	mt := "map[string]int"
	mt1 := typeString
	mt2 := typeInt
	ms := "(len=" + mLen + ") {\n (" + mt1 + ") (len=" + klen + ") " +
		"\"one\": (" + mt2 + ") 1,\n (" + mt1 + ") (len=" + kkLen +
		") \"two\": (" + mt2 + ") 2\n}"
	ms2 := "(len=" + mLen + ") {\n (" + mt1 + ") (len=" + kkLen + ") " +
		"\"two\": (" + mt2 + ") 2,\n (" + mt1 + ") (len=" + klen +
		") \"one\": (" + mt2 + ") 1\n}"
	add(m, "("+mt+") "+ms+"\n", "("+mt+") "+ms2+"\n")
	add(pm, "(*"+mt+")("+mAddr+")("+ms+")\n",
		"(*"+mt+")("+mAddr+")("+ms2+")\n")
	add(&pm, "(**"+mt+")("+pmAddr+"->"+mAddr+")("+ms+")\n",
		"(**"+mt+")("+pmAddr+"->"+mAddr+")("+ms2+")\n")
	add(nm, "(*"+mt+")(<nil>)\n")
	add(nilMap, "("+mt+") <nil>\n")

	// Map with custom formatter type on pointer receiver only keys and vals.
	k2 := pstringer(valOne)
	v2 := pstringer("1")
	m2 := map[pstringer]pstringer{k2: v2}
	k2Len := strconv.Itoa(len(k2))
	v2Len := strconv.Itoa(len(v2))
	m2Len := strconv.Itoa(len(m2))
	nilMap2 := map[pstringer]pstringer(nil)
	nm2 := (*map[pstringer]pstringer)(nil)
	pm2 := &m2
	m2Addr := fmt.Sprintf("%p", pm2)
	pm2Addr := fmt.Sprintf("%p", &pm2)
	m2t := "map[spew_test.pstringer]spew_test.pstringer"
	m2t1 := typePstringer
	m2t2 := typePstringer
	m2s := "(len=" + m2Len + ") {\n (" + m2t1 + ") (len=" + k2Len + ") " +
		"stringer one: (" + m2t2 + ") (len=" + v2Len + ") stringer 1\n}"
	if spew.UnsafeDisabled {
		m2s = "(len=" + m2Len + ") {\n (" + m2t1 + ") (len=" + k2Len +
			") " + "\"one\": (" + m2t2 + ") (len=" + v2Len +
			") \"1\"\n}"
	}
	add(m2, "("+m2t+") "+m2s+"\n")
	add(pm2, "(*"+m2t+")("+m2Addr+")("+m2s+")\n")
	add(&pm2, "(**"+m2t+")("+pm2Addr+"->"+m2Addr+")("+m2s+")\n")
	add(nm2, "(*"+m2t+")(<nil>)\n")
	add(nilMap2, "("+m2t+") <nil>\n")

	// Map with interface keys and values.
	k3 := valOne
	k3Len := strconv.Itoa(len(k3))
	m3 := map[any]any{k3: 1}
	m3Len := strconv.Itoa(len(m3))
	nilMap3 := map[any]any(nil)
	nm3 := (*map[any]any)(nil)
	pm3 := &m3
	m3Addr := fmt.Sprintf("%p", pm3)
	pm3Addr := fmt.Sprintf("%p", &pm3)
	m3t := "map[" + typeInterface + "]" + typeInterface
	m3t1 := typeString
	m3t2 := typeInt
	m3s := "(len=" + m3Len + ") {\n (" + m3t1 + ") (len=" + k3Len + ") " +
		"\"one\": (" + m3t2 + ") 1\n}"
	add(m3, "("+m3t+") "+m3s+"\n")
	add(pm3, "(*"+m3t+")("+m3Addr+")("+m3s+")\n")
	add(&pm3, "(**"+m3t+")("+pm3Addr+"->"+m3Addr+")("+m3s+")\n")
	add(nm3, "(*"+m3t+")(<nil>)\n")
	add(nilMap3, "("+m3t+") <nil>\n")

	// Map with nil interface value.
	k4 := "nil"
	k4Len := strconv.Itoa(len(k4))
	m4 := map[string]any{k4: nil}
	m4Len := strconv.Itoa(len(m4))
	nilMap4 := map[string]any(nil)
	nm4 := (*map[string]any)(nil)
	pm4 := &m4
	m4Addr := fmt.Sprintf("%p", pm4)
	pm4Addr := fmt.Sprintf("%p", &pm4)
	m4t := "map[string]" + typeInterface
	m4t1 := typeString
	m4t2 := typeInterface
	m4s := "(len=" + m4Len + ") {\n (" + m4t1 + ") (len=" + k4Len + ")" +
		" \"nil\": (" + m4t2 + ") <nil>\n}"
	add(m4, "("+m4t+") "+m4s+"\n")
	add(pm4, "(*"+m4t+")("+m4Addr+")("+m4s+")\n")
	add(&pm4, "(**"+m4t+")("+pm4Addr+"->"+m4Addr+")("+m4s+")\n")
	add(nm4, "(*"+m4t+")(<nil>)\n")
	add(nilMap4, "("+m4t+") <nil>\n")

	return tests
}

func structDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
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
	vt := typeS1
	vt2 := typeInt8
	vt3 := typeUint8
	vs := "{\n a: (" + vt2 + ") 127,\n b: (" + vt3 + ") 255\n}"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

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
	v2t := typeS2
	v2t2 := typeS1
	v2t3 := typeInt8
	v2t4 := typeUint8
	v2t5 := typeBool
	v2s := "{\n s1: (" + v2t2 + ") {\n  a: (" + v2t3 + ") 127,\n  b: (" +
		v2t4 + ") 255\n },\n b: (" + v2t5 + ") true\n}"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*"+v2t+")(<nil>)\n")

	// Struct that contains custom type with Stringer pointer interface via both
	// exported and unexported fields.
	type s3 struct {
		s pstringer
		S pstringer
	}
	v3 := s3{valTest, "test2"}
	nv3 := (*s3)(nil)
	pv3 := &v3
	v3Addr := fmt.Sprintf("%p", pv3)
	pv3Addr := fmt.Sprintf("%p", &pv3)
	v3t := "spew_test.s3"
	v3t2 := typePstringer
	v3s := "{\n s: (" + v3t2 + ") (len=4) stringer test,\n S: (" + v3t2 +
		") (len=5) stringer test2\n}"
	v3sp := v3s
	if spew.UnsafeDisabled {
		v3s = "{\n s: (" + v3t2 + ") (len=4) \"test\",\n S: (" +
			v3t2 + ") (len=5) \"test2\"\n}"
		v3sp = "{\n s: (" + v3t2 + ") (len=4) \"test\",\n S: (" +
			v3t2 + ") (len=5) stringer test2\n}"
	}
	add(v3, "("+v3t+") "+v3s+"\n")
	add(pv3, "(*"+v3t+")("+v3Addr+")("+v3sp+")\n")
	add(&pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")("+v3sp+")\n")
	add(nv3, "(*"+v3t+")(<nil>)\n")

	// Struct that contains embedded struct and field to same struct.
	e := embed{"embedstr"}
	eLen := strconv.Itoa(len("embedstr"))
	v4 := embedwrap{embed: &e, e: &e}
	nv4 := (*embedwrap)(nil)
	pv4 := &v4
	eAddr := fmt.Sprintf("%p", &e)
	v4Addr := fmt.Sprintf("%p", pv4)
	pv4Addr := fmt.Sprintf("%p", &pv4)
	v4t := "spew_test.embedwrap"
	v4t2 := "spew_test.embed"
	v4t3 := typeString
	v4s := "{\n embed: (*" + v4t2 + ")(" + eAddr + ")({\n  a: (" + v4t3 +
		") (len=" + eLen + ") \"embedstr\"\n }),\n e: (*" + v4t2 +
		")(" + eAddr + ")({\n  a: (" + v4t3 + ") (len=" + eLen + ")" +
		" \"embedstr\"\n })\n}"
	add(v4, "("+v4t+") "+v4s+"\n")
	add(pv4, "(*"+v4t+")("+v4Addr+")("+v4s+")\n")
	add(&pv4, "(**"+v4t+")("+pv4Addr+"->"+v4Addr+")("+v4s+")\n")
	add(nv4, "(*"+v4t+")(<nil>)\n")

	return tests
}

func uintptrDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Null pointer.
	v := uintptr(0)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeUintptr
	vs := typeNil
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")

	// Address of real variable.
	i := 1
	v2 := uintptr(unsafe.Pointer(&i))
	nv2 := (*uintptr)(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeUintptr
	v2s := fmt.Sprintf("%p", &i)
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*"+v2t+")(<nil>)\n")

	return tests
}

func unsafePointerDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Null pointer.
	v := unsafe.Pointer(nil)
	nv := (*unsafe.Pointer)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeUnsafePointer
	vs := typeNil
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Address of real variable.
	i := 1
	v2 := unsafe.Pointer(&i)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeUnsafePointer
	v2s := fmt.Sprintf("%p", &i)
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	return tests
}

func chanDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Nil channel.
	var v chan int
	pv := &v
	nv := (*chan int)(nil)
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := typeChanInt
	vs := typeNil
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Real channel.
	v2 := make(chan int)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := typeChanInt
	v2s := fmt.Sprintf("%p", v2)
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")

	return tests
}

func funcDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Function with no params and no returns.
	v := func() {}
	nv := (*func())(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "func()"
	vs := fmt.Sprintf("%p", v)
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	// Function with param and no returns.
	v2 := TestDump
	nv2 := (*func(*testing.T))(nil)
	pv2 := &v2
	v2Addr := fmt.Sprintf("%p", pv2)
	pv2Addr := fmt.Sprintf("%p", &pv2)
	v2t := "func(*testing.T)"
	v2s := fmt.Sprintf("%p", v2)
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s+")\n")
	add(nv2, "(*"+v2t+")(<nil>)\n")

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
	add(v3, "("+v3t+") "+v3s+"\n")
	add(pv3, "(*"+v3t+")("+v3Addr+")("+v3s+")\n")
	add(&pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")("+v3s+")\n")
	add(nv3, "(*"+v3t+")(<nil>)\n")

	return tests
}

func circularDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
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
	vs := "{\n c: (*" + vt + ")(" + vAddr + ")({\n  c: (*" + vt + ")(" +
		vAddr + ")(<already shown>)\n })\n}"
	vs2 := "{\n c: (*" + vt + ")(" + vAddr + ")(<already shown>)\n}"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs2+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs2+")\n")

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
	v2s := "{\n ps2: (*" + v2t2 + ")(" + ts2Addr + ")({\n  ps1: (*" + v2t +
		")(" + v2Addr + ")({\n   ps2: (*" + v2t2 + ")(" + ts2Addr +
		")(<already shown>)\n  })\n })\n}"
	v2s2 := "{\n ps2: (*" + v2t2 + ")(" + ts2Addr + ")({\n  ps1: (*" + v2t +
		")(" + v2Addr + ")(<already shown>)\n })\n}"
	add(v2, "("+v2t+") "+v2s+"\n")
	add(pv2, "(*"+v2t+")("+v2Addr+")("+v2s2+")\n")
	add(&pv2, "(**"+v2t+")("+pv2Addr+"->"+v2Addr+")("+v2s2+")\n")

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
	v3s := "{\n ps2: (*" + v3t2 + ")(" + tic2Addr + ")({\n  ps3: (*" + v3t3 +
		")(" + tic3Addr + ")({\n   ps1: (*" + v3t + ")(" + v3Addr +
		")({\n    ps2: (*" + v3t2 + ")(" + tic2Addr +
		")(<already shown>)\n   })\n  })\n })\n}"
	v3s2 := "{\n ps2: (*" + v3t2 + ")(" + tic2Addr + ")({\n  ps3: (*" + v3t3 +
		")(" + tic3Addr + ")({\n   ps1: (*" + v3t + ")(" + v3Addr +
		")(<already shown>)\n  })\n })\n}"
	add(v3, "("+v3t+") "+v3s+"\n")
	add(pv3, "(*"+v3t+")("+v3Addr+")("+v3s2+")\n")
	add(&pv3, "(**"+v3t+")("+pv3Addr+"->"+v3Addr+")("+v3s2+")\n")

	return tests
}

func panicDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Type that panics in its Stringer interface.
	v := panicer(127)
	nv := (*panicer)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "spew_test.panicer"
	vs := "(PANIC=test panic)127"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	return tests
}

func errorDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	// Type that has a custom Error interface.
	v := customError(127)
	nv := (*customError)(nil)
	pv := &v
	vAddr := fmt.Sprintf("%p", pv)
	pvAddr := fmt.Sprintf("%p", &pv)
	vt := "spew_test.customError"
	vs := "error: 127"
	add(v, "("+vt+") "+vs+"\n")
	add(pv, "(*"+vt+")("+vAddr+")("+vs+")\n")
	add(&pv, "(**"+vt+")("+pvAddr+"->"+vAddr+")("+vs+")\n")
	add(nv, "(*"+vt+")(<nil>)\n")

	return tests
}

type (
	embeddedTime struct {
		time.Time
	}
	embeddedTimePtr struct { // this type is an example where the unsafeReflectValue does not work well
		*time.Time
	}
	redeclaredTime         time.Time
	redeclaredTimePtr      *time.Time
	aliasedTime            = time.Time
	embeddedRedeclaredTime struct {
		redeclaredTime
	}
)

func timeDumpTests() []dumpTest {
	var tests []dumpTest
	add := func(in any, wants ...string) {
		tests = append(tests, dumpTest{in, wants})
	}

	ts := time.Date(2006, time.January, 2, 15, 4, 5, 999999999, time.UTC)
	alias := aliasedTime(ts) //nolint:unconvert // we want to prove here that aliased types don't matter
	tsAddr := fmt.Sprintf("%p", &ts)
	es := embeddedTime{
		Time: ts,
	}
	ps := embeddedTimePtr{
		Time: &ts,
	}
	var tptr *time.Time
	panick := embeddedTimePtr{
		Time: nil,
	}
	rtptr := redeclaredTimePtr(&ts)
	ppts := ptr(ptr(ts))
	ppAddr := fmt.Sprintf("%p", ppts)
	ppIAddr := fmt.Sprintf("%p", *ppts)
	er := embeddedRedeclaredTime{
		redeclaredTime: redeclaredTime(ts),
	}

	add(
		// simple time.Time
		ts,
		"(time.Time) 2006-01-02 15:04:05.999999999 +0000 UTC\n",
	)
	add(
		// aliases are ignored at runtime
		alias,
		"(time.Time) 2006-01-02 15:04:05.999999999 +0000 UTC\n",
	)
	add(
		// pointer to time.Time
		&ts,
		"(*time.Time)("+tsAddr+")(2006-01-02 15:04:05.999999999 +0000 UTC)\n",
	)
	add(
		// struct with embedded time.Time
		es,
		"(spew_test.embeddedTime) 2006-01-02 15:04:05.999999999 +0000 UTC\n",
	)
	add(
		// struct with embedded pointer to time.Time
		ps,
		"(spew_test.embeddedTimePtr) 2006-01-02 15:04:05.999999999 +0000 UTC\n",
	)
	add(
		// nil time.Time
		tptr,
		"(*time.Time)(<nil>)\n",
	)
	add(
		// **time.Time
		ppts,
		"(**time.Time)("+ppAddr+"->"+ppIAddr+")(2006-01-02 15:04:05.999999999 +0000 UTC)\n",
	)
	add(
		panick, // this is a stringer, but the inner member that implements String() string is nil
		"(spew_test.embeddedTimePtr) (PANIC=runtime error: invalid memory address or nil pointer dereference){\n Time: (*time.Time)(<nil>)\n}\n",
	)
	add(
		// redeclared type convertible to time.Time
		redeclaredTime(ts),
		"(spew_test.redeclaredTime) 2006-01-02 15:04:05.999999999 +0000 UTC\n",
	)
	add(
		// redeclared type convertible to *time.Time
		//
		// NOTE: the information about the original (redeclared) type is lost. This is due to
		// how displaying pointer type information is displayed (i.e. using v.Elem().Type()).
		rtptr,
		"(*time.Time)("+tsAddr+")(2006-01-02 15:04:05.999999999 +0000 UTC)\n",
	)
	add(
		// embedded redeclared type convertible to time.Time
		er,
		"(spew_test.embeddedRedeclaredTime) {\n"+
			" redeclaredTime: (spew_test.redeclaredTime) 2006-01-02 15:04:05.999999999 +0000 UTC\n}\n",
	)

	return tests
}
