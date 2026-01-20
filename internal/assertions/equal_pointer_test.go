package assertions

import (
	"fmt"
	"iter"
	"slices"
	"testing"
)

// Pointer identity tests (Same, SameT, NotSame, NotSameT).
func TestEqualPointers(t *testing.T) {
	t.Parallel()

	for tc := range unifiedPointerPairCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			expected, actual := tc.makeValues()

			if !tc.reflectionOnly {
				t.Run("with Same", testPointerAssertion(tc, sameKind, Same, expected, actual))
				t.Run("with NotSame", testPointerAssertion(tc, notSameKind, NotSame, expected, actual))

				// Generic variants - type dispatch
				t.Run("with SameT", testPointerAssertionT(tc, sameTKind, expected, actual))
				t.Run("with NotSameT", testPointerAssertionT(tc, notSameTKind, expected, actual))

				return
			}

			// Reflection-only cases (non-pointer args, different types, one nil)
			t.Run("with Same (reflection)", testPointerAssertion(tc, sameKind, Same, expected, actual))
			t.Run("with NotSame (reflection)", testPointerAssertion(tc, notSameKind, NotSame, expected, actual))
		})
	}
}

type pointerPairTestCase struct {
	name           string
	makeValues     func() (expected, actual any)
	relationship   pairRelationship
	reflectionOnly bool
}

func unifiedPointerPairCases() iter.Seq[pointerPairTestCase] {
	const hello = "hello"

	return slices.Values([]pointerPairTestCase{
		// Both nil
		{"both-nil/ptr", func() (any, any) { return (*int)(nil), (*int)(nil) }, bothNil, false},
		{"both-nil/different-types", func() (any, any) { return (*int)(nil), (*string)(nil) }, differentTypes, true},

		// One nil
		{"one-nil/first", func() (any, any) { v := 42; return (*int)(nil), &v }, oneNil, true},
		{"one-nil/second", func() (any, any) { v := 42; return &v, (*int)(nil) }, oneNil, true},

		// Same identity - both point to same address
		{"same-identity/int", func() (any, any) { v := 42; return &v, &v }, sameIdentity, false},
		{"same-identity/string", func() (any, any) { s := hello; return &s, &s }, sameIdentity, false},
		{"same-identity/float64", func() (any, any) { f := 3.14; return &f, &f }, sameIdentity, false},

		// Different identity - point to different adresses
		{"different-identity/equal-values", func() (any, any) { v1, v2 := 42, 42; return &v1, &v2 }, differentIdentity, false},
		{"different-identity/different-values", func() (any, any) { v1, v2 := 42, 43; return &v1, &v2 }, differentIdentity, false},

		// Different types (reflection-only)
		{"different-types/int-string", func() (any, any) {
			i, s := 1, hello
			return &i, &s
		}, differentTypes, true},

		// Edge cases (always false)
		{"not-pointer/right", func() (any, any) { v1 := 12; return &v1, 1 }, notPointer, true},
		{"not-pointer/left", func() (any, any) { v1 := 12; return 1, &v1 }, notPointer, true},
		{"not-pointer/both", func() (any, any) { return 1, 1 }, notPointer, true},
	})
}

func testPointerAssertion(
	tc pointerPairTestCase,
	kind pointerAssertionKind,
	pointerAssertion func(T, any, any, ...any) bool,
	expected, actual any,
) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := pointerAssertion(mock, expected, actual)
		shouldPass := expectedStatusForPointerAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

type pointerAssertionKind int

const (
	sameKind pointerAssertionKind = iota
	sameTKind
	notSameKind
	notSameTKind
)

type pairRelationship int

const (
	bothNil pairRelationship = iota
	oneNil
	sameIdentity
	differentIdentity
	differentTypes
	notPointer
)

func expectedStatusForPointerAssertion(kind pointerAssertionKind, relationship pairRelationship) bool {
	positive := kind == sameKind || kind == sameTKind

	switch relationship {
	case notPointer:
		return false
	case sameIdentity, bothNil:
		// Two nil pointers of the same type are considered "same" in Go
		return positive
	case oneNil, differentIdentity, differentTypes:
		return !positive
	default:
		panic(fmt.Errorf("test case configuration error: invalid pairRelationship: %d", relationship))
	}
}

func testPointerAssertionT(tc pointerPairTestCase, kind pointerAssertionKind, expected, actual any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		stop := func(expected string, actual any) {
			t.Fatalf("test case error: expected=%s, actual=%T", expected, actual)
		}

		// Type switch with safety check
		//
		// Add more supported types to the switch with new test cases
		var result bool
		switch exp := expected.(type) {
		case *int:
			act, ok := actual.(*int)
			if !ok {
				stop("*int", actual)
			}
			result = testPointerGenericAssertion(mock, kind, exp, act)
		case *string:
			act, ok := actual.(*string)
			if !ok {
				stop("*string", actual)
			}
			result = testPointerGenericAssertion(mock, kind, exp, act)
		case *float64:
			act, ok := actual.(*float64)
			if !ok {
				stop("*float64", actual)
			}
			result = testPointerGenericAssertion(mock, kind, exp, act)
		default:
			t.Fatalf("unsupported type: %T", expected)
		}

		shouldPass := expectedStatusForPointerAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

func testPointerGenericAssertion[P any](mock T, kind pointerAssertionKind, expected, actual *P) bool {
	switch kind {
	case sameTKind:
		return SameT(mock, expected, actual)
	case notSameTKind:
		return NotSameT(mock, expected, actual)
	default:
		panic(fmt.Errorf("test case configuration error: invalid pointerAssertionKind: %d", kind))
	}
}
