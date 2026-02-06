// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"testing"
)

// Unary assertion tests (Nil, NotNil, Empty, NotEmpty).
func TestEqualUnaryAssertions(t *testing.T) {
	t.Parallel()

	for tc := range unifiedUnaryCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			t.Run("with Nil", testUnaryAssertion(tc, nilKind, Nil))
			t.Run("with NotNil", testUnaryAssertion(tc, notNilKind, NotNil))
			t.Run("with Empty", testUnaryAssertion(tc, emptyKind, Empty))
			t.Run("with NotEmpty", testUnaryAssertion(tc, notEmptyKind, NotEmpty))
		})
	}
}

func TestEqualUnaryErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, equalUnaryFailCases())
}

// ============================================================================
// TestEqualUnaryAssertions
// ============================================================================

type unaryTestCase struct {
	name     string
	object   any
	category objectCategory
}

func unifiedUnaryCases() iter.Seq[unaryTestCase] {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	x := 1
	xP := &x
	z := 0
	zP := &z
	var arr [1]int

	type TString string
	type TStruct struct {
		x int
	}
	type FStruct struct {
		x func()
	}

	return slices.Values([]unaryTestCase{
		// Nil category
		{"nil/nil-ptr", (*int)(nil), nilCategory},
		{"nil/nil-slice", []int(nil), nilCategory},
		{"nil/nil-interface", (any)(nil), nilCategory},
		{"nil/nil-struct-ptr", (*struct{})(nil), nilCategory},

		// Empty non-nil category
		{"empty/slice", []int{}, emptyNonNil},
		{"empty/string", "", emptyNonNil},
		{"empty/zero-int", 0, emptyNonNil},
		{"empty/zero-bool", false, emptyNonNil},
		{"empty/channel", make(chan struct{}), emptyNonNil},
		{"empty/zero-struct", TStruct{}, emptyNonNil},
		{"empty/aliased-string", TString(""), emptyNonNil},
		{"empty/zero-array", [1]int{}, emptyNonNil},
		{"empty/zero-ptr", zP, emptyNonNil},
		{"empty/zero-struct-ptr", &TStruct{}, emptyNonNil},
		{"empty/zero-array-ptr", &arr, emptyNonNil},
		{"empty/rune", '\u0000', emptyNonNil},
		{"empty/complex", 0i, emptyNonNil},
		{"empty/error", errors.New(""), emptyNonNil},
		{"empty/struct-with-func", FStruct{x: nil}, emptyNonNil},

		// Non-empty comparable category
		{"non-empty/int", 42, nonEmptyComparable},
		{"non-empty/rune", 'A', nonEmptyComparable},
		{"non-empty/string", "hello", nonEmptyComparable},
		{"non-empty/bool", true, nonEmptyComparable},
		{"non-empty/slice", []int{1}, nonEmptyComparable},
		{"non-empty/channel", chWithValue, nonEmptyComparable},
		{"non-empty/struct", TStruct{x: 1}, nonEmptyComparable},
		{"non-empty/aliased-string", TString("abc"), nonEmptyComparable},
		{"non-empty/ptr", xP, nonEmptyComparable},
		{"non-empty/array", [1]int{42}, nonEmptyComparable},

		// Non-empty non-comparable category
		{"non-empty/error", errors.New("something"), nonEmptyNonComparable},
		{"non-empty/slice-error", []error{errors.New("")}, nonEmptyNonComparable},
		{"non-empty/slice-nil-error", []error{nil}, nonEmptyNonComparable},
		{"non-empty/slice-zero", []int{0}, nonEmptyNonComparable},
		{"non-empty/slice-nil", []*int{nil}, nonEmptyNonComparable},
		{"non-empty/struct-with-func", FStruct{x: func() {}}, nonEmptyNonComparable},
	})
}

type unaryAssertionKind int

const (
	nilKind unaryAssertionKind = iota
	notNilKind
	emptyKind
	notEmptyKind
)

type objectCategory int

const (
	nilCategory objectCategory = iota
	emptyNonNil
	nonEmptyComparable
	nonEmptyNonComparable
)

// expectedStatusForUnaryAssertion returns the expected semantics for a given assertion (Nil, Empty, ...)
// and a given category of input.
func expectedStatusForUnaryAssertion(kind unaryAssertionKind, category objectCategory) bool {
	switch kind {
	case nilKind:
		return category == nilCategory
	case notNilKind:
		return category != nilCategory
	case emptyKind:
		return category == nilCategory || category == emptyNonNil
	case notEmptyKind:
		return category == nonEmptyComparable || category == nonEmptyNonComparable
	default:
		panic(fmt.Errorf("test case configuration error: invalid unaryAssertionKind: %d", kind))
	}
}

func testUnaryAssertion(tc unaryTestCase, kind unaryAssertionKind, unaryAssertion func(T, any, ...any) bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := unaryAssertion(mock, tc.object)
		shouldPass := expectedStatusForUnaryAssertion(kind, tc.category)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

// ============================================================================
// TestEqualUnaryErrorMessages
// ============================================================================

func equalUnaryFailCases() iter.Seq[failCase] {
	chWithValue := make(chan struct{}, 1)
	chWithValue <- struct{}{}
	x := 1
	xP := &x

	type TString string
	type TStruct struct {
		x int
	}

	return slices.Values([]failCase{
		{
			name:      "Empty/non-empty-string",
			assertion: func(t T) bool { return Empty(t, "something") },
			wantError: "Should be empty, but was something",
		},
		{
			name:      "Empty/non-nil-error",
			assertion: func(t T) bool { return Empty(t, errors.New("something")) },
			wantError: "Should be empty, but was something",
		},
		{
			name:      "Empty/non-empty-string-array",
			assertion: func(t T) bool { return Empty(t, []string{"something"}) },
			wantError: "Should be empty, but was [something]",
		},
		{
			name:      "Empty/non-zero-int",
			assertion: func(t T) bool { return Empty(t, 1) },
			wantError: "Should be empty, but was 1",
		},
		{
			name:      "Empty/true-value",
			assertion: func(t T) bool { return Empty(t, true) },
			wantError: "Should be empty, but was true",
		},
		{
			name:         "Empty/channel-with-values",
			assertion:    func(t T) bool { return Empty(t, chWithValue) },
			wantContains: []string{"Should be empty, but was"},
		},
		{
			name:      "Empty/struct-with-values",
			assertion: func(t T) bool { return Empty(t, TStruct{x: 1}) },
			wantError: "Should be empty, but was {1}",
		},
		{
			name:      "Empty/aliased-string",
			assertion: func(t T) bool { return Empty(t, TString("abc")) },
			wantError: "Should be empty, but was abc",
		},
		{
			name:         "Empty/ptr-to-non-nil",
			assertion:    func(t T) bool { return Empty(t, xP) },
			wantContains: []string{"Should be empty, but was"},
		},
		{
			name:      "Empty/non-zero-array",
			assertion: func(t T) bool { return Empty(t, [1]int{42}) },
			wantError: "Should be empty, but was [42]",
		},
		{
			name:         "Empty/whitespace-string",
			assertion:    func(t T) bool { return Empty(t, "   ") },
			wantContains: []string{"Should be empty, but was"},
		},
		{
			name:         "Empty/newline-string",
			assertion:    func(t T) bool { return Empty(t, "\n") },
			wantContains: []string{"Should be empty, but was"},
		},
		{
			name:         "Empty/non-printable-char",
			assertion:    func(t T) bool { return Empty(t, "\u00a0") },
			wantContains: []string{"Should be empty, but was"},
		},
	})
}
