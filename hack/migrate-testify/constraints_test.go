// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"go/token"
	"go/types"
	"iter"
	"slices"
	"testing"
)

type constraintTestCase struct {
	name       string
	typ        types.Type
	constraint constraintKind
	expected   bool
}

func constraintTestCases() iter.Seq[constraintTestCase] {
	// Create some test types.
	intType := types.Typ[types.Int]
	int64Type := types.Typ[types.Int64]
	float64Type := types.Typ[types.Float64]
	stringType := types.Typ[types.String]
	boolType := types.Typ[types.Bool]
	uint8Type := types.Typ[types.Uint8]
	complex128Type := types.Typ[types.Complex128]

	// []byte
	byteSlice := types.NewSlice(types.Typ[types.Byte])

	// *int
	intPtr := types.NewPointer(intType)

	// interface{}
	emptyIface := types.NewInterfaceType(nil, nil)

	return slices.Values([]constraintTestCase{
		// comparable
		{name: "int is comparable", typ: intType, constraint: constraintComparable, expected: true},
		{name: "string is comparable", typ: stringType, constraint: constraintComparable, expected: true},
		{name: "bool is comparable", typ: boolType, constraint: constraintComparable, expected: true},

		// ordered
		{name: "int is ordered", typ: intType, constraint: constraintOrdered, expected: true},
		{name: "float64 is ordered", typ: float64Type, constraint: constraintOrdered, expected: true},
		{name: "string is ordered", typ: stringType, constraint: constraintOrdered, expected: true},
		{name: "[]byte is ordered", typ: byteSlice, constraint: constraintOrdered, expected: true},
		{name: "bool is NOT ordered", typ: boolType, constraint: constraintOrdered, expected: false},

		// text
		{name: "string is text", typ: stringType, constraint: constraintText, expected: true},
		{name: "[]byte is text", typ: byteSlice, constraint: constraintText, expected: true},
		{name: "int is NOT text", typ: intType, constraint: constraintText, expected: false},

		// signedNumeric
		{name: "int is signedNumeric", typ: intType, constraint: constraintSignedNumeric, expected: true},
		{name: "int64 is signedNumeric", typ: int64Type, constraint: constraintSignedNumeric, expected: true},
		{name: "float64 is signedNumeric", typ: float64Type, constraint: constraintSignedNumeric, expected: true},
		{name: "uint8 is NOT signedNumeric", typ: uint8Type, constraint: constraintSignedNumeric, expected: false},

		// measurable
		{name: "int is measurable", typ: intType, constraint: constraintMeasurable, expected: true},
		{name: "uint8 is measurable", typ: uint8Type, constraint: constraintMeasurable, expected: true},
		{name: "float64 is measurable", typ: float64Type, constraint: constraintMeasurable, expected: true},
		{name: "complex128 is NOT measurable", typ: complex128Type, constraint: constraintMeasurable, expected: false},
		{name: "string is NOT measurable", typ: stringType, constraint: constraintMeasurable, expected: false},

		// boolean
		{name: "bool is boolean", typ: boolType, constraint: constraintBoolean, expected: true},
		{name: "int is NOT boolean", typ: intType, constraint: constraintBoolean, expected: false},

		// pointer
		{name: "*int is pointer", typ: intPtr, constraint: constraintPointer, expected: true},
		{name: "int is NOT pointer", typ: intType, constraint: constraintPointer, expected: false},

		// deepComparable — same as comparable but excludes pointers and structs with pointers
		{name: "int is deepComparable", typ: intType, constraint: constraintDeepComparable, expected: true},
		{name: "string is deepComparable", typ: stringType, constraint: constraintDeepComparable, expected: true},
		{name: "bool is deepComparable", typ: boolType, constraint: constraintDeepComparable, expected: true},
		{name: "*int is NOT deepComparable", typ: intPtr, constraint: constraintDeepComparable, expected: false},

		// any/interface{} should not satisfy anything
		{name: "interface{} is NOT comparable", typ: emptyIface, constraint: constraintComparable, expected: false},
		{name: "interface{} is NOT ordered", typ: emptyIface, constraint: constraintOrdered, expected: false},
		{name: "interface{} is NOT deepComparable", typ: emptyIface, constraint: constraintDeepComparable, expected: false},
	})
}

func TestSatisfiesConstraint(t *testing.T) {
	t.Parallel()

	for c := range constraintTestCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			got := satisfiesConstraint(c.typ, c.constraint)
			if got != c.expected {
				t.Errorf("satisfiesConstraint(%v, %v) = %v, want %v", c.typ, c.constraint, got, c.expected)
			}
		})
	}
}

func TestSameType(t *testing.T) {
	t.Parallel()

	intType := types.Typ[types.Int]
	int64Type := types.Typ[types.Int64]
	stringType := types.Typ[types.String]

	if !sameType(intType, intType) {
		t.Error("int should be same type as int")
	}
	if sameType(intType, int64Type) {
		t.Error("int should NOT be same type as int64")
	}
	if sameType(intType, stringType) {
		t.Error("int should NOT be same type as string")
	}
}

func TestDeepComparable(t *testing.T) {
	t.Parallel()

	// Struct with no pointers — deep-comparable.
	plainStruct := types.NewStruct([]*types.Var{
		types.NewVar(token.NoPos, nil, "X", types.Typ[types.Int]),
		types.NewVar(token.NoPos, nil, "Y", types.Typ[types.String]),
	}, nil)
	if !isDeepComparable(plainStruct) {
		t.Error("struct{X int; Y string} should be deep-comparable")
	}

	// Struct with a pointer field — NOT deep-comparable.
	ptrStruct := types.NewStruct([]*types.Var{
		types.NewVar(token.NoPos, nil, "X", types.Typ[types.Int]),
		types.NewVar(token.NoPos, nil, "P", types.NewPointer(types.Typ[types.Int])),
	}, nil)
	if isDeepComparable(ptrStruct) {
		t.Error("struct{X int; P *int} should NOT be deep-comparable")
	}

	// Struct with an interface field — NOT deep-comparable.
	ifaceStruct := types.NewStruct([]*types.Var{
		types.NewVar(token.NoPos, nil, "X", types.Typ[types.Int]),
		types.NewVar(token.NoPos, nil, "E", types.NewInterfaceType(nil, nil)),
	}, nil)
	if isDeepComparable(ifaceStruct) {
		t.Error("struct{X int; E interface{}} should NOT be deep-comparable")
	}

	// Array of int — deep-comparable.
	intArray := types.NewArray(types.Typ[types.Int], 3)
	if !isDeepComparable(intArray) {
		t.Error("[3]int should be deep-comparable")
	}

	// Array of *int — NOT deep-comparable.
	ptrArray := types.NewArray(types.NewPointer(types.Typ[types.Int]), 3)
	if isDeepComparable(ptrArray) {
		t.Error("[3]*int should NOT be deep-comparable")
	}
}

func TestIsSliceType(t *testing.T) {
	t.Parallel()

	intSlice := types.NewSlice(types.Typ[types.Int])
	elem, ok := isSliceType(intSlice)
	if !ok {
		t.Fatal("expected []int to be a slice type")
	}
	if !sameType(elem, types.Typ[types.Int]) {
		t.Errorf("expected slice element to be int, got %v", elem)
	}

	_, ok = isSliceType(types.Typ[types.Int])
	if ok {
		t.Error("int should not be a slice type")
	}
}

func TestIsMapType(t *testing.T) {
	t.Parallel()

	m := types.NewMap(types.Typ[types.String], types.Typ[types.Int])
	key, val, ok := isMapType(m)
	if !ok {
		t.Fatal("expected map[string]int to be a map type")
	}
	if !sameType(key, types.Typ[types.String]) {
		t.Errorf("expected key to be string, got %v", key)
	}
	if !sameType(val, types.Typ[types.Int]) {
		t.Errorf("expected val to be int, got %v", val)
	}
}
