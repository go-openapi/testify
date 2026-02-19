// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"runtime/debug"
)

// PanicAssertionFunc is a common function prototype when validating a panic value.  Can be useful
// for table driven tests.
type PanicAssertionFunc func(t T, f func(), msgAndArgs ...any) bool

// Panics asserts that the code inside the specified function panics.
//
// # Usage
//
//	assertions.Panics(t, func(){ GoCrazy() })
//
// # Examples
//
//	success: func() { panic("panicking") }
//	failure: func() { }
func Panics(t T, f func(), msgAndArgs ...any) bool {
	// Domain: panic
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if funcDidPanic, panicValue, _ := didPanic(f); !funcDidPanic {
		return Fail(t, fmt.Sprintf("func should panic\n\tPanic value:\t%#v", panicValue), msgAndArgs...)
	}

	return true
}

// PanicsWithValue asserts that the code inside the specified function panics,
// and that the recovered panic value equals the expected panic value.
//
// # Usage
//
//	assertions.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
//
// # Examples
//
//	success: "panicking", func() { panic("panicking") }
//	failure: "panicking", func() { }
func PanicsWithValue(t T, expected any, f func(), msgAndArgs ...any) bool {
	// Domain: panic
	if h, ok := t.(H); ok {
		h.Helper()
	}

	funcDidPanic, panicValue, panickedStack := didPanic(f)
	if !funcDidPanic {
		return Fail(t, fmt.Sprintf("func should panic\n\tPanic value:\t%#v", panicValue), msgAndArgs...)
	}
	if panicValue != expected {
		return Fail(t, fmt.Sprintf("func should panic with value:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s", expected, panicValue, panickedStack), msgAndArgs...)
	}

	return true
}

// PanicsWithError asserts that the code inside the specified function panics,
// and that the recovered panic value is an error that satisfies the [EqualError] comparison.
//
// # Usage
//
//	assertions.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
//
// # Examples
//
//	success: ErrTest.Error(), func() { panic(ErrTest) }
//	failure: ErrTest.Error(), func() { }
func PanicsWithError(t T, errString string, f func(), msgAndArgs ...any) bool {
	// Domain: panic
	if h, ok := t.(H); ok {
		h.Helper()
	}

	funcDidPanic, panicValue, panickedStack := didPanic(f)
	if !funcDidPanic {
		return Fail(t, fmt.Sprintf("func should panic\n\tPanic value:\t%#v", panicValue), msgAndArgs...)
	}
	panicErr, isError := panicValue.(error)
	if !isError || panicErr.Error() != errString {
		msg := fmt.Sprintf("func should panic with error message:\t%#v\n", errString)
		if isError {
			msg += fmt.Sprintf("\tError message:\t%#v\n", panicErr.Error())
		}
		msg += fmt.Sprintf("\tPanic value:\t%#v\n", panicValue)
		msg += fmt.Sprintf("\tPanic stack:\t%s\n", panickedStack)
		return Fail(t, msg, msgAndArgs...)
	}

	return true
}

// NotPanics asserts that the code inside the specified function does NOT panic.
//
// # Usage
//
//	assertions.NotPanics(t, func(){ RemainCalm() })
//
// # Examples
//
//	success: func() { }
//	failure: func() { panic("panicking") }
func NotPanics(t T, f func(), msgAndArgs ...any) bool {
	// Domain: panic
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if funcDidPanic, panicValue, panickedStack := didPanic(f); funcDidPanic {
		return Fail(t, fmt.Sprintf("func should not panic\n\tPanic value:\t%v\n\tPanic stack:\t%s", panicValue, panickedStack), msgAndArgs...)
	}

	return true
}

// didPanic returns true if the function passed to it panics. Otherwise, it returns false.
func didPanic(f func()) (didPanic bool, message any, stack string) {
	didPanic = true

	defer func() {
		message = recover()
		if didPanic {
			stack = string(debug.Stack())
		}
		// Go 1.21 introduces runtime.PanicNilError on panic(nil),
		// so maintain the same logic going forward (https://github.com/golang/go/issues/25448).
		if err, ok := message.(error); ok {
			if err.Error() == "panic called with nil argument" {
				message = nil
			}
		}
	}()

	// call the target function
	f()
	didPanic = false

	return
}
