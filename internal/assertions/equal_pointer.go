// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "fmt"

// Same asserts that two pointers reference the same object.
//
// Both arguments must be pointer variables.
//
// Pointer variable sameness is determined based on the equality of both type and value.
//
// Unlike [Equal] pointers, [Same] pointers point to the same memory address.
//
// # Usage
//
//	assertions.Same(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, staticVarPtr
//	failure: &staticVar, ptr("static string")
func Same(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	same, ok := samePointers(expected, actual)
	if !ok {
		return Fail(t, "Both arguments must be pointers", msgAndArgs...)
	}

	if !same {
		// both are pointers but not the same type & pointing to the same address
		return Fail(t, fmt.Sprintf("Not same: \n"+
			"expected: %[2]s (%[1]T)(%[1]p)\n"+
			"actual  : %[4]s (%[3]T)(%[3]p)", expected, truncatingFormat("%#v", expected), actual, truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

// SameT asserts that two pointers of the same type reference the same object.
//
// See [Same].
//
// # Usage
//
//	assertions.SameT(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, staticVarPtr
//	failure: &staticVar, ptr("static string")
func SameT[P any](t T, expected, actual *P, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if expected != actual {
		return Fail(t, fmt.Sprintf("Not same: \n"+
			"expected: %[2]s (%[1]T)(%[1]p)\n"+
			"actual  : %[4]s (%[3]T)(%[3]p)", expected, truncatingFormat("%#v", expected), actual, truncatingFormat("%#v", actual)), msgAndArgs...)
	}

	return true
}

// NotSame asserts that two pointers do not reference the same object.
//
// See [Same].
//
// # Usage
//
//	assertions.NotSame(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, ptr("static string")
//	failure: &staticVar, staticVarPtr
func NotSame(t T, expected, actual any, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	same, ok := samePointers(expected, actual)
	if !ok {
		// fails when the arguments are not pointers
		return Fail(t, "Both arguments must be pointers", msgAndArgs...)
	}

	if same {
		return Fail(t, fmt.Sprintf(
			"Expected and actual point to the same object: %p %s",
			expected, truncatingFormat("%#v", expected)), msgAndArgs...)
	}
	return true
}

// NotSameT asserts that two pointers do not reference the same object.
//
// See [SameT].
//
// # Usage
//
//	assertions.NotSameT(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, ptr("static string")
//	failure: &staticVar, staticVarPtr
func NotSameT[P any](t T, expected, actual *P, msgAndArgs ...any) bool {
	// Domain: equality
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if expected == actual {
		return Fail(t, fmt.Sprintf(
			"Expected and actual point to the same object: %p %s",
			expected, truncatingFormat("%#v", expected)), msgAndArgs...)
	}

	return true
}
