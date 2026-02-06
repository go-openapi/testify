// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"iter"
	"reflect"
	"slices"
	"testing"
)

func TestTypeImplements(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if !Implements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject)) {
		t.Error("Implements method should return true: AssertionTesterConformingObject implements AssertionTesterInterface")
	}
	if Implements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject)) {
		t.Error("Implements method should return false: AssertionTesterNonConformingObject does not implements AssertionTesterInterface")
	}
	if Implements(mock, (*AssertionTesterInterface)(nil), nil) {
		t.Error("Implements method should return false: nil does not implement AssertionTesterInterface")
	}
}

func TestTypeNotImplements(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if !NotImplements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterNonConformingObject)) {
		t.Error("NotImplements method should return true: AssertionTesterNonConformingObject does not implement AssertionTesterInterface")
	}
	if NotImplements(mock, (*AssertionTesterInterface)(nil), new(AssertionTesterConformingObject)) {
		t.Error("NotImplements method should return false: AssertionTesterConformingObject implements AssertionTesterInterface")
	}
	if NotImplements(mock, (*AssertionTesterInterface)(nil), nil) {
		t.Error("NotImplements method should return false: nil can't be checked to be implementing AssertionTesterInterface or not")
	}
}

func TestTypeIsType(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if !IsType(mock, new(AssertionTesterConformingObject), new(AssertionTesterConformingObject)) {
		t.Error("IsType should return true: AssertionTesterConformingObject is the same type as AssertionTesterConformingObject")
	}
	if IsType(mock, new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject)) {
		t.Error("IsType should return false: AssertionTesterConformingObject is not the same type as AssertionTesterNonConformingObject")
	}
}

func TestTypeNotIsType(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	if !IsNotType(mock, new(AssertionTesterConformingObject), new(AssertionTesterNonConformingObject)) {
		t.Error("NotIsType should return true: AssertionTesterConformingObject is not the same type as AssertionTesterNonConformingObject")
	}
	if IsNotType(mock, new(AssertionTesterConformingObject), new(AssertionTesterConformingObject)) {
		t.Error("NotIsType should return false: AssertionTesterConformingObject is the same type as AssertionTesterConformingObject")
	}
}

func TestTypeIsOfTypeT(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	type myType float64
	var myVar myType = 1.2
	f := 1.2

	if !IsOfTypeT[myType](mock, myVar) {
		t.Errorf("expected myVar to be of type %T", myVar)
	}
	if IsNotOfTypeT[myType](mock, myVar) {
		t.Errorf("expected IsNotOfTypeT to return false for myVar of type %T", myVar)
	}
	if IsOfTypeT[myType](mock, f) {
		t.Errorf("expected f (%T) not to be of type %T", f, myVar)
	}
	if !IsNotOfTypeT[myType](mock, f) {
		t.Errorf("expected IsNotOfTypeT to return true for f (%T) vs %T", f, myVar)
	}
}

func TestTypeZero(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	for test := range typeZeros() {
		if !Zero(mock, test) {
			t.Errorf("expected %#v to be the zero value for %T", test, test)
		}
	}

	for test := range typeNonZeros() {
		if Zero(mock, test) {
			t.Errorf("expected %#v to NOT be the zero value for %T", test, test)
		}
	}
}

func TestTypeNotZero(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	for test := range typeZeros() {
		if NotZero(mock, test) {
			t.Errorf("expected NotZero to return false for zero value %#v (%T)", test, test)
		}
	}

	for test := range typeNonZeros() {
		if !NotZero(mock, test) {
			t.Errorf("expected NotZero to return true for non-zero value %#v (%T)", test, test)
		}
	}
}

func TestTypeKind(t *testing.T) {
	t.Parallel()

	for tt := range kindCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			mock := new(mockT)
			result := Kind(mock, tt.expectedKind, tt.value)
			resultNot := NotKind(mock, tt.expectedKind, tt.value)

			if tt.result {
				if !result {
					t.Errorf("expected kind of %T to be %q, but Kind reported %t", tt.value, tt.expectedKind, result)
				}
				if resultNot {
					t.Errorf("expected kind of %T to be %q, but NotKind reported %t", tt.value, tt.expectedKind, resultNot)
				}

				return
			}

			// expected: false
			if result {
				t.Errorf("expected kind of %T NOT to be %q, but Kind reported %t", tt.value, tt.expectedKind, result)
			}
			if !resultNot {
				t.Errorf("expected kind of %T NOT to be %q, but NotKind reported %t", tt.value, tt.expectedKind, resultNot)
			}
		})
	}
}

func TestTypeDiffEmptyCases(t *testing.T) {
	t.Parallel()

	cases := []struct {
		a, b any
	}{
		{nil, nil},
		{struct{ foo string }{}, nil},
		{nil, struct{ foo string }{}},
		{1, 2},
		{1, 2},
		{[]int{1}, []bool{true}},
	}
	for _, tc := range cases {
		if result := diff(tc.a, tc.b); result != "" {
			t.Errorf("expected empty diff for (%v, %v), got %q", tc.a, tc.b, result)
		}
	}
}

func TestTypeErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, typeFailCases())
}

// =======================================
// TestTypeErrorMessages
// =======================================

func typeFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "Zero/large-slice-truncated",
			assertion:    func(t T) bool { return Zero(t, make([]int, 1_000_000)) },
			wantContains: []string{"Should be zero, but was", "<... truncated>"},
		},
	})
}

// =======================================
// TestTypeIsZero
// =======================================

func typeZeros() iter.Seq[any] {
	return slices.Values([]any{
		false,
		byte(0),
		complex64(0),
		complex128(0),
		float32(0),
		float64(0),
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		rune(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		uintptr(0),
		"",
		[0]any{},
		[]any(nil),
		struct{ x int }{},
		(*any)(nil),
		(func())(nil),
		nil,
		any(nil),
		map[any]any(nil),
		(chan any)(nil),
		(<-chan any)(nil),
		(chan<- any)(nil),
	})
}

func typeNonZeros() iter.Seq[any] {
	var i int

	return slices.Values([]any{
		true,
		byte(1),
		complex64(1),
		complex128(1),
		float32(1),
		float64(1),
		int(1),
		int8(1),
		int16(1),
		int32(1),
		int64(1),
		rune(1),
		uint(1),
		uint8(1),
		uint16(1),
		uint32(1),
		uint64(1),
		uintptr(1),
		"s",
		[1]any{1},
		[]any{},
		struct{ x int }{1},
		(&i),
		(func() {}),
		any(1),
		map[any]any{},
		(make(chan any)),
		(<-chan any)(make(chan any)),
		(chan<- any)(make(chan any)),
	})
}

type kindCase struct {
	expectedKind reflect.Kind
	value        any
	result       bool
	name         string
}

func kindCases() iter.Seq[kindCase] {
	var iface any = "string"

	return slices.Values([]kindCase{
		// True cases
		{reflect.String, "Hello World", true, "is string"},
		{reflect.Int, 123, true, "is int"},
		{reflect.Array, [6]int{2, 3, 5, 7, 11, 13}, true, "is array"},
		{reflect.Func, Kind, true, "is func"},
		{reflect.Float64, 0.0345, true, "is float64"},
		{reflect.Map, make(map[string]int), true, "is map"},
		{reflect.Bool, true, true, "is bool"},
		{reflect.Ptr, new(int), true, "is pointer"},
		// False cases
		{reflect.String, 13, false, "not string"},
		{reflect.Int, [6]int{2, 3, 5, 7, 11, 13}, false, "not int"},
		{reflect.Float64, 12, false, "not float64"},
		{reflect.Bool, make(map[string]int), false, "not bool"},
		// Edge cases
		// True
		{reflect.Invalid, any(nil), true, "legitimate expectation of reflect.Invalid (any)"},
		{reflect.Ptr, (*any)(nil), true, "legitimate expectation of reflect.Pointer (*any)"},
		{reflect.Invalid, (error)(nil), true, "legitimate expectation of reflect.Invalid (error)"},
		{reflect.Invalid, nil, true, "legitimate nil input"},
		// False
		{reflect.Interface, iface, false, "interface returns concrete type (any)"},
		{reflect.Interface, errors.New("stuff"), false, "interface returns concrete type (error)"},
		{reflect.Invalid, "string", false, "wrong expectation of reflect.Invalid"},
		{reflect.Ptr, nil, false, "nil input"},
	})
}
