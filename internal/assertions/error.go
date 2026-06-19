// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var _ error = TestExampleError("")

// TestExampleError is a sentinel error type that may be used for testing.
type TestExampleError string

func (e TestExampleError) Error() string {
	return string(e)
}

// ErrTest is an error instance useful for testing.
//
// If the code does not care about error specifics, and only needs
// to return the error as an example, this error may be used to make
// the test code more readable.
const ErrTest TestExampleError = "assert.ErrTest general error for testing"

// NoError asserts that a function returned a nil error (i.e. no error).
//
// # Usage
//
//	actualObj, err := SomeFunction()
//	if assert.NoError(t, err) {
//		assertions.Equal(t, expectedObj, actualObj)
//	}
//
// # Examples
//
//	success: nil
//	failure: ErrTest
func NoError(t T, err error, msgAndArgs ...any) bool {
	// Domain: error
	if err != nil {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		return Fail(t, "Received unexpected error:\n"+truncatingFormat("%+v", err), msgAndArgs...)
	}

	return true
}

// Error asserts that a function returned a non-nil error (i.e. an error).
//
// # Usage
//
//	actualObj, err := SomeFunction()
//	assertions.Error(t, err)
//
// # Examples
//
//	success: ErrTest
//	failure: nil
func Error(t T, err error, msgAndArgs ...any) bool {
	// Domain: error
	// Opposite: NoError
	if err == nil {
		if h, ok := t.(H); ok {
			h.Helper()
		}
		return Fail(t, "An error is expected but got nil.", msgAndArgs...)
	}

	return true
}

// EqualError asserts that a function returned a non-nil error (i.e. an error)
// and that it is equal to the provided error.
//
// # Usage
//
//	actualObj, err := SomeFunction()
//	assertions.EqualError(t, err,  expectedErrorString)
//
// # Examples
//
//	success: ErrTest, "assert.ErrTest general error for testing"
//	failure: ErrTest, "wrong error message"
func EqualError(t T, err error, errString string, msgAndArgs ...any) bool {
	// Domain: error
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if !Error(t, err, msgAndArgs...) {
		return false
	}
	expected := errString
	actual := err.Error()
	// don't need to use deep equals here, we know they are both strings
	if expected != actual {
		return Fail(t, fmt.Sprintf("Error message not equal:\n"+
			"expected: %q\n"+
			"actual  : %s", expected, truncatingFormat("%q", actual)), msgAndArgs...)
	}
	return true
}

// ErrorContains asserts that a function returned a non-nil error (i.e. an
// error) and that the error contains the specified substring.
//
// # Usage
//
//	actualObj, err := SomeFunction()
//	assertions.ErrorContains(t, err,  expectedErrorSubString)
//
// # Examples
//
//	success: ErrTest, "general error"
//	failure: ErrTest, "not in message"
func ErrorContains(t T, err error, contains string, msgAndArgs ...any) bool {
	// Domain: error
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if !Error(t, err, msgAndArgs...) {
		return false
	}

	actual := err.Error()
	if !strings.Contains(actual, contains) {
		return Fail(t, fmt.Sprintf("Error %s does not contain %#v", truncatingFormat("%#v", actual), contains), msgAndArgs...)
	}

	return true
}

// ErrorIs asserts that at least one of the errors in err's chain matches target.
//
// This is a wrapper for [errors.Is].
//
// # Usage
//
//	assertions.ErrorIs(t, err, io.EOF)
//
// # Examples
//
//	success: fmt.Errorf("wrap: %w", io.EOF), io.EOF
//	failure: ErrTest, io.EOF
func ErrorIs(t T, err, target error, msgAndArgs ...any) bool {
	// Domain: error
	// Opposite: NotErrorIs
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if errors.Is(err, target) {
		return true
	}

	var expectedText string
	if target != nil {
		expectedText = target.Error()
		if err == nil {
			return Fail(t, fmt.Sprintf("Expected error with %q in chain but got nil.", expectedText), msgAndArgs...)
		}
	}

	chain := buildErrorChainString(err, false)

	return Fail(t, fmt.Sprintf("Target error should be in err chain:\n"+
		"expected: %s\n"+
		"in chain: %s", truncatingFormat("%q", expectedText), truncatingFormat("%s", chain),
	), msgAndArgs...)
}

// NotErrorIs asserts that none of the errors in err's chain matches target.
//
// This is a wrapper for [errors.Is].
//
// # Usage
//
//	assertions.NotErrorIs(t, err, io.EOF)
//
// # Examples
//
//	success: ErrTest, io.EOF
//	failure: fmt.Errorf("wrap: %w", io.EOF), io.EOF
func NotErrorIs(t T, err, target error, msgAndArgs ...any) bool {
	// Domain: error
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if !errors.Is(err, target) {
		return true
	}

	var expectedText string
	if target != nil {
		expectedText = target.Error()
	}

	chain := buildErrorChainString(err, false)

	return Fail(t, fmt.Sprintf("Target error should not be in err chain:\n"+
		"found: %s\n"+
		"in chain: %s", truncatingFormat("%q", expectedText), truncatingFormat("%s", chain),
	), msgAndArgs...)
}

// ErrorAs asserts that at least one of the errors in err's chain matches target, and if so, sets target to that error value.
//
// This is a wrapper for [errors.As].
//
// # Usage
//
//	assertions.ErrorAs(t, err, &target)
//
// # Examples
//
//	success: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
//	failure: ErrTest, new(*dummyError)
func ErrorAs(t T, err error, target any, msgAndArgs ...any) bool {
	// Domain: error
	// Opposite: NotErrorAs
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if errors.As(err, target) {
		return true
	}

	expectedType := reflect.TypeOf(target).Elem().String()
	if err == nil {
		return Fail(t, fmt.Sprintf("An error is expected but got nil.\n"+
			"expected: %s", expectedType), msgAndArgs...)
	}

	chain := buildErrorChainString(err, true)

	return Fail(t, fmt.Sprintf("Should be in error chain:\n"+
		"expected: %s\n"+
		"in chain: %s", expectedType, truncatingFormat("%s", chain),
	), msgAndArgs...)
}

// NotErrorAs asserts that none of the errors in err's chain matches target,
// but if so, sets target to that error value.
//
// # Usage
//
//	assertions.NotErrorAs(t, err, &target)
//
// # Examples
//
//	success: ErrTest, new(*dummyError)
//	failure: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
func NotErrorAs(t T, err error, target any, msgAndArgs ...any) bool {
	// Domain: error
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if !errors.As(err, target) {
		return true
	}

	chain := buildErrorChainString(err, true)

	return Fail(t, fmt.Sprintf("Target error should not be in err chain:\n"+
		"found: %s\n"+
		"in chain: %s", reflect.TypeOf(target).Elem().String(), truncatingFormat("%s", chain),
	), msgAndArgs...)
}

func unwrapAll(err error) []error {
	return appendUnwrapped(nil, err, make(map[error]struct{}))
}

// appendUnwrapped flattens an error chain, guarding against cyclic chains that
// would otherwise recurse until the goroutine stack overflows.
//
// Only comparable errors are tracked: using an incomparable error as a map key
// would panic, so those simply skip cycle detection (mirroring how the standard
// library's errors.Is guards its own comparisons).
func appendUnwrapped(errs []error, err error, visited map[error]struct{}) []error {
	if err != nil && reflect.TypeOf(err).Comparable() {
		if _, ok := visited[err]; ok {
			return errs // cyclic error chain: stop here
		}
		visited[err] = struct{}{}
		defer delete(visited, err)
	}

	errs = append(errs, err)
	switch x := err.(type) { //nolint:errorlint // false positive: this type switch is checking for interfaces
	case interface{ Unwrap() error }:
		next := x.Unwrap()
		if next == nil {
			return errs
		}
		errs = appendUnwrapped(errs, next, visited)
	case interface{ Unwrap() []error }:
		for _, next := range x.Unwrap() {
			errs = appendUnwrapped(errs, next, visited)
		}
	}
	return errs
}

func buildErrorChainString(err error, withType bool) string {
	if err == nil {
		return ""
	}

	var chain strings.Builder
	errs := unwrapAll(err)
	for i := range errs {
		if i != 0 {
			chain.WriteString("\n\t")
		}
		fmt.Fprintf(&chain, "%q", errs[i].Error())
		if withType {
			fmt.Fprintf(&chain, " (%T)", errs[i])
		}
	}

	return chain.String()
}
