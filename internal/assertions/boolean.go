// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

// True asserts that the specified value is true.
//
// # Usage
//
//	assertions.True(t, myBool)
//
// # Examples
//
//	success: 1 == 1
//	failure: 1 == 0
func True(t T, value bool, msgAndArgs ...any) bool {
	// Domain: boolean
	if !value {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		return Fail(t, "Should be true", msgAndArgs...)
	}

	return true
}

// TrueT asserts that the specified value is true.
//
// # Usage
//
//	type B bool
//	var b B = true
//
//	assertions.True(t, b)
//
// # Examples
//
//	success: 1 == 1
//	failure: 1 == 0
func TrueT[B Boolean](t T, value B, msgAndArgs ...any) bool {
	// Domain: boolean
	if !bool(value) {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		return Fail(t, "Should be true", msgAndArgs...)
	}

	return true
}

// False asserts that the specified value is false.
//
// # Usage
//
//	assertions.False(t, myBool)
//
// # Examples
//
//	success: 1 == 0
//	failure: 1 == 1
func False(t T, value bool, msgAndArgs ...any) bool {
	// Domain: boolean
	if value {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		return Fail(t, "Should be false", msgAndArgs...)
	}

	return true
}

// FalseT asserts that the specified value is false.
//
// # Usage
//
//	 type B bool
//	 var b B = true
//
//		assertions.FalseT(t, b)
//
// # Examples
//
//	success: 1 == 0
//	failure: 1 == 1
func FalseT[B Boolean](t T, value B, msgAndArgs ...any) bool {
	// Domain: boolean
	if bool(value) {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		return Fail(t, "Should be true", msgAndArgs...)
	}

	return true
}
