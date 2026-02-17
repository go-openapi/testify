// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

package assert

import (
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// Assertions exposes all assertion functions as methods.
//
// NOTE: assertion methods with parameterized types (generics) are not supported as methods.
//
// Upon failure, the test [T] is marked as failed and continues execution.
type Assertions struct {
	T
}

// New makes a new [Assertions] object for the specified [T] (e.g. [testing.T]).
func New(t T) *Assertions {
	return &Assertions{
		T: t,
	}
}

// Condition is the same as [Condition], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Condition(comp Comparison, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Condition(a.T, comp, msgAndArgs...)
}

// Conditionf is the same as [Assertions.Condition], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Conditionf(comp Comparison, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Condition(a.T, comp, forwardArgs(msg, args))
}

// Contains is the same as [Contains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Contains(s any, contains any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Contains(a.T, s, contains, msgAndArgs...)
}

// Containsf is the same as [Assertions.Contains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Containsf(s any, contains any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Contains(a.T, s, contains, forwardArgs(msg, args))
}

// DirExists is the same as [DirExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(a.T, path, msgAndArgs...)
}

// DirExistsf is the same as [Assertions.DirExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(a.T, path, forwardArgs(msg, args))
}

// DirNotExists is the same as [DirNotExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirNotExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirNotExists(a.T, path, msgAndArgs...)
}

// DirNotExistsf is the same as [Assertions.DirNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirNotExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirNotExists(a.T, path, forwardArgs(msg, args))
}

// ElementsMatch is the same as [ElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatch(listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(a.T, listA, listB, msgAndArgs...)
}

// ElementsMatchf is the same as [Assertions.ElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatchf(listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(a.T, listA, listB, forwardArgs(msg, args))
}

// Empty is the same as [Empty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Empty(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Empty(a.T, object, msgAndArgs...)
}

// Emptyf is the same as [Assertions.Empty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Emptyf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Empty(a.T, object, forwardArgs(msg, args))
}

// Equal is the same as [Equal], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Equal(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Equal(a.T, expected, actual, msgAndArgs...)
}

// Equalf is the same as [Assertions.Equal], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Equalf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Equal(a.T, expected, actual, forwardArgs(msg, args))
}

// EqualError is the same as [EqualError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualError(err error, errString string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(a.T, err, errString, msgAndArgs...)
}

// EqualErrorf is the same as [Assertions.EqualError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualErrorf(err error, errString string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(a.T, err, errString, forwardArgs(msg, args))
}

// EqualExportedValues is the same as [EqualExportedValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualExportedValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(a.T, expected, actual, msgAndArgs...)
}

// EqualExportedValuesf is the same as [Assertions.EqualExportedValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualExportedValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(a.T, expected, actual, forwardArgs(msg, args))
}

// EqualValues is the same as [EqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(a.T, expected, actual, msgAndArgs...)
}

// EqualValuesf is the same as [Assertions.EqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(a.T, expected, actual, forwardArgs(msg, args))
}

// Error is the same as [Error], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Error(err error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Error(a.T, err, msgAndArgs...)
}

// Errorf is the same as [Assertions.Error], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Errorf(err error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Error(a.T, err, forwardArgs(msg, args))
}

// ErrorAs is the same as [ErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAs(err error, target any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(a.T, err, target, msgAndArgs...)
}

// ErrorAsf is the same as [Assertions.ErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAsf(err error, target any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(a.T, err, target, forwardArgs(msg, args))
}

// ErrorContains is the same as [ErrorContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorContains(err error, contains string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(a.T, err, contains, msgAndArgs...)
}

// ErrorContainsf is the same as [Assertions.ErrorContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorContainsf(err error, contains string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(a.T, err, contains, forwardArgs(msg, args))
}

// ErrorIs is the same as [ErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(a.T, err, target, msgAndArgs...)
}

// ErrorIsf is the same as [Assertions.ErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(a.T, err, target, forwardArgs(msg, args))
}

// Eventually is the same as [Eventually], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Eventually(a.T, condition, waitFor, tick, msgAndArgs...)
}

// Eventuallyf is the same as [Assertions.Eventually], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Eventually(a.T, condition, waitFor, tick, forwardArgs(msg, args))
}

// EventuallyWith is the same as [EventuallyWith], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EventuallyWith(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWith(a.T, condition, waitFor, tick, msgAndArgs...)
}

// EventuallyWithf is the same as [Assertions.EventuallyWith], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EventuallyWithf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWith(a.T, condition, waitFor, tick, forwardArgs(msg, args))
}

// Exactly is the same as [Exactly], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Exactly(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(a.T, expected, actual, msgAndArgs...)
}

// Exactlyf is the same as [Assertions.Exactly], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Exactlyf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(a.T, expected, actual, forwardArgs(msg, args))
}

// Fail is the same as [Fail], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Fail(a.T, failureMessage, msgAndArgs...)
}

// Failf is the same as [Assertions.Fail], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Failf(failureMessage string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Fail(a.T, failureMessage, forwardArgs(msg, args))
}

// FailNow is the same as [FailNow], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(a.T, failureMessage, msgAndArgs...)
}

// FailNowf is the same as [Assertions.FailNow], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(a.T, failureMessage, forwardArgs(msg, args))
}

// False is the same as [False], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) False(value bool, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.False(a.T, value, msgAndArgs...)
}

// Falsef is the same as [Assertions.False], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Falsef(value bool, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.False(a.T, value, forwardArgs(msg, args))
}

// FileEmpty is the same as [FileEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileEmpty(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(a.T, path, msgAndArgs...)
}

// FileEmptyf is the same as [Assertions.FileEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileEmptyf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(a.T, path, forwardArgs(msg, args))
}

// FileExists is the same as [FileExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(a.T, path, msgAndArgs...)
}

// FileExistsf is the same as [Assertions.FileExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(a.T, path, forwardArgs(msg, args))
}

// FileNotEmpty is the same as [FileNotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotEmpty(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(a.T, path, msgAndArgs...)
}

// FileNotEmptyf is the same as [Assertions.FileNotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotEmptyf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(a.T, path, forwardArgs(msg, args))
}

// FileNotExists is the same as [FileNotExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotExists(a.T, path, msgAndArgs...)
}

// FileNotExistsf is the same as [Assertions.FileNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotExists(a.T, path, forwardArgs(msg, args))
}

// Greater is the same as [Greater], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Greater(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Greater(a.T, e1, e2, msgAndArgs...)
}

// Greaterf is the same as [Assertions.Greater], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Greaterf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Greater(a.T, e1, e2, forwardArgs(msg, args))
}

// GreaterOrEqual is the same as [GreaterOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqual(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(a.T, e1, e2, msgAndArgs...)
}

// GreaterOrEqualf is the same as [Assertions.GreaterOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqualf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(a.T, e1, e2, forwardArgs(msg, args))
}

// HTTPBodyContains is the same as [HTTPBodyContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(a.T, handler, method, url, values, str, msgAndArgs...)
}

// HTTPBodyContainsf is the same as [Assertions.HTTPBodyContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(a.T, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPBodyNotContains is the same as [HTTPBodyNotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(a.T, handler, method, url, values, str, msgAndArgs...)
}

// HTTPBodyNotContainsf is the same as [Assertions.HTTPBodyNotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(a.T, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPError is the same as [HTTPError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(a.T, handler, method, url, values, msgAndArgs...)
}

// HTTPErrorf is the same as [Assertions.HTTPError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(a.T, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPRedirect is the same as [HTTPRedirect], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(a.T, handler, method, url, values, msgAndArgs...)
}

// HTTPRedirectf is the same as [Assertions.HTTPRedirect], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(a.T, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPStatusCode is the same as [HTTPStatusCode], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(a.T, handler, method, url, values, statuscode, msgAndArgs...)
}

// HTTPStatusCodef is the same as [Assertions.HTTPStatusCode], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(a.T, handler, method, url, values, statuscode, forwardArgs(msg, args))
}

// HTTPSuccess is the same as [HTTPSuccess], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(a.T, handler, method, url, values, msgAndArgs...)
}

// HTTPSuccessf is the same as [Assertions.HTTPSuccess], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(a.T, handler, method, url, values, forwardArgs(msg, args))
}

// Implements is the same as [Implements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Implements(interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Implements(a.T, interfaceObject, object, msgAndArgs...)
}

// Implementsf is the same as [Assertions.Implements], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Implementsf(interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Implements(a.T, interfaceObject, object, forwardArgs(msg, args))
}

// InDelta is the same as [InDelta], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDelta(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(a.T, expected, actual, delta, msgAndArgs...)
}

// InDeltaf is the same as [Assertions.InDelta], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaf(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(a.T, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaMapValues is the same as [InDeltaMapValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaMapValues(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(a.T, expected, actual, delta, msgAndArgs...)
}

// InDeltaMapValuesf is the same as [Assertions.InDeltaMapValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaMapValuesf(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(a.T, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaSlice is the same as [InDeltaSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaSlice(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(a.T, expected, actual, delta, msgAndArgs...)
}

// InDeltaSlicef is the same as [Assertions.InDeltaSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaSlicef(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(a.T, expected, actual, delta, forwardArgs(msg, args))
}

// InEpsilon is the same as [InEpsilon], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilon(expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(a.T, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonf is the same as [Assertions.InEpsilon], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonf(expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(a.T, expected, actual, epsilon, forwardArgs(msg, args))
}

// InEpsilonSlice is the same as [InEpsilonSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSlice(expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(a.T, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonSlicef is the same as [Assertions.InEpsilonSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSlicef(expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(a.T, expected, actual, epsilon, forwardArgs(msg, args))
}

// IsDecreasing is the same as [IsDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(a.T, collection, msgAndArgs...)
}

// IsDecreasingf is the same as [Assertions.IsDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(a.T, collection, forwardArgs(msg, args))
}

// IsIncreasing is the same as [IsIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(a.T, collection, msgAndArgs...)
}

// IsIncreasingf is the same as [Assertions.IsIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(a.T, collection, forwardArgs(msg, args))
}

// IsNonDecreasing is the same as [IsNonDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(a.T, collection, msgAndArgs...)
}

// IsNonDecreasingf is the same as [Assertions.IsNonDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(a.T, collection, forwardArgs(msg, args))
}

// IsNonIncreasing is the same as [IsNonIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(a.T, collection, msgAndArgs...)
}

// IsNonIncreasingf is the same as [Assertions.IsNonIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(a.T, collection, forwardArgs(msg, args))
}

// IsNotType is the same as [IsNotType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotType(theType any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(a.T, theType, object, msgAndArgs...)
}

// IsNotTypef is the same as [Assertions.IsNotType], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotTypef(theType any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(a.T, theType, object, forwardArgs(msg, args))
}

// IsType is the same as [IsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsType(expectedType any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsType(a.T, expectedType, object, msgAndArgs...)
}

// IsTypef is the same as [Assertions.IsType], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsTypef(expectedType any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsType(a.T, expectedType, object, forwardArgs(msg, args))
}

// JSONEq is the same as [JSONEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(a.T, expected, actual, msgAndArgs...)
}

// JSONEqf is the same as [Assertions.JSONEq], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(a.T, expected, actual, forwardArgs(msg, args))
}

// JSONEqBytes is the same as [JSONEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqBytes(expected []byte, actual []byte, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(a.T, expected, actual, msgAndArgs...)
}

// JSONEqBytesf is the same as [Assertions.JSONEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqBytesf(expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(a.T, expected, actual, forwardArgs(msg, args))
}

// Kind is the same as [Kind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Kind(expectedKind reflect.Kind, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Kind(a.T, expectedKind, object, msgAndArgs...)
}

// Kindf is the same as [Assertions.Kind], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Kindf(expectedKind reflect.Kind, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Kind(a.T, expectedKind, object, forwardArgs(msg, args))
}

// Len is the same as [Len], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Len(object any, length int, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Len(a.T, object, length, msgAndArgs...)
}

// Lenf is the same as [Assertions.Len], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Lenf(object any, length int, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Len(a.T, object, length, forwardArgs(msg, args))
}

// Less is the same as [Less], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Less(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Less(a.T, e1, e2, msgAndArgs...)
}

// Lessf is the same as [Assertions.Less], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Lessf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Less(a.T, e1, e2, forwardArgs(msg, args))
}

// LessOrEqual is the same as [LessOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqual(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(a.T, e1, e2, msgAndArgs...)
}

// LessOrEqualf is the same as [Assertions.LessOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqualf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(a.T, e1, e2, forwardArgs(msg, args))
}

// Negative is the same as [Negative], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Negative(e any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Negative(a.T, e, msgAndArgs...)
}

// Negativef is the same as [Assertions.Negative], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Negativef(e any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Negative(a.T, e, forwardArgs(msg, args))
}

// Never is the same as [Never], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Never(a.T, condition, waitFor, tick, msgAndArgs...)
}

// Neverf is the same as [Assertions.Never], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Never(a.T, condition, waitFor, tick, forwardArgs(msg, args))
}

// Nil is the same as [Nil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Nil(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Nil(a.T, object, msgAndArgs...)
}

// Nilf is the same as [Assertions.Nil], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Nilf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Nil(a.T, object, forwardArgs(msg, args))
}

// NoError is the same as [NoError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoError(err error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoError(a.T, err, msgAndArgs...)
}

// NoErrorf is the same as [Assertions.NoError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoErrorf(err error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoError(a.T, err, forwardArgs(msg, args))
}

// NoFileDescriptorLeak is the same as [NoFileDescriptorLeak], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoFileDescriptorLeak(tested func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoFileDescriptorLeak(a.T, tested, msgAndArgs...)
}

// NoFileDescriptorLeakf is the same as [Assertions.NoFileDescriptorLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoFileDescriptorLeakf(tested func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoFileDescriptorLeak(a.T, tested, forwardArgs(msg, args))
}

// NoGoRoutineLeak is the same as [NoGoRoutineLeak], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoGoRoutineLeak(tested func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoGoRoutineLeak(a.T, tested, msgAndArgs...)
}

// NoGoRoutineLeakf is the same as [Assertions.NoGoRoutineLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoGoRoutineLeakf(tested func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoGoRoutineLeak(a.T, tested, forwardArgs(msg, args))
}

// NotContains is the same as [NotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotContains(s any, contains any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(a.T, s, contains, msgAndArgs...)
}

// NotContainsf is the same as [Assertions.NotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotContainsf(s any, contains any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(a.T, s, contains, forwardArgs(msg, args))
}

// NotElementsMatch is the same as [NotElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatch(listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(a.T, listA, listB, msgAndArgs...)
}

// NotElementsMatchf is the same as [Assertions.NotElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatchf(listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(a.T, listA, listB, forwardArgs(msg, args))
}

// NotEmpty is the same as [NotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEmpty(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(a.T, object, msgAndArgs...)
}

// NotEmptyf is the same as [Assertions.NotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEmptyf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(a.T, object, forwardArgs(msg, args))
}

// NotEqual is the same as [NotEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqual(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(a.T, expected, actual, msgAndArgs...)
}

// NotEqualf is the same as [Assertions.NotEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(a.T, expected, actual, forwardArgs(msg, args))
}

// NotEqualValues is the same as [NotEqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(a.T, expected, actual, msgAndArgs...)
}

// NotEqualValuesf is the same as [Assertions.NotEqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(a.T, expected, actual, forwardArgs(msg, args))
}

// NotErrorAs is the same as [NotErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAs(err error, target any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(a.T, err, target, msgAndArgs...)
}

// NotErrorAsf is the same as [Assertions.NotErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAsf(err error, target any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(a.T, err, target, forwardArgs(msg, args))
}

// NotErrorIs is the same as [NotErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(a.T, err, target, msgAndArgs...)
}

// NotErrorIsf is the same as [Assertions.NotErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(a.T, err, target, forwardArgs(msg, args))
}

// NotImplements is the same as [NotImplements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotImplements(interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(a.T, interfaceObject, object, msgAndArgs...)
}

// NotImplementsf is the same as [Assertions.NotImplements], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotImplementsf(interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(a.T, interfaceObject, object, forwardArgs(msg, args))
}

// NotKind is the same as [NotKind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotKind(expectedKind reflect.Kind, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotKind(a.T, expectedKind, object, msgAndArgs...)
}

// NotKindf is the same as [Assertions.NotKind], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotKindf(expectedKind reflect.Kind, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotKind(a.T, expectedKind, object, forwardArgs(msg, args))
}

// NotNil is the same as [NotNil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotNil(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(a.T, object, msgAndArgs...)
}

// NotNilf is the same as [Assertions.NotNil], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotNilf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(a.T, object, forwardArgs(msg, args))
}

// NotPanics is the same as [NotPanics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotPanics(f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(a.T, f, msgAndArgs...)
}

// NotPanicsf is the same as [Assertions.NotPanics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotPanicsf(f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(a.T, f, forwardArgs(msg, args))
}

// NotRegexp is the same as [NotRegexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexp(rx any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(a.T, rx, actual, msgAndArgs...)
}

// NotRegexpf is the same as [Assertions.NotRegexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexpf(rx any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(a.T, rx, actual, forwardArgs(msg, args))
}

// NotSame is the same as [NotSame], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSame(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(a.T, expected, actual, msgAndArgs...)
}

// NotSamef is the same as [Assertions.NotSame], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSamef(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(a.T, expected, actual, forwardArgs(msg, args))
}

// NotSubset is the same as [NotSubset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSubset(list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(a.T, list, subset, msgAndArgs...)
}

// NotSubsetf is the same as [Assertions.NotSubset], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSubsetf(list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(a.T, list, subset, forwardArgs(msg, args))
}

// NotZero is the same as [NotZero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotZero(i any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(a.T, i, msgAndArgs...)
}

// NotZerof is the same as [Assertions.NotZero], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotZerof(i any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(a.T, i, forwardArgs(msg, args))
}

// Panics is the same as [Panics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Panics(f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Panics(a.T, f, msgAndArgs...)
}

// Panicsf is the same as [Assertions.Panics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Panicsf(f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Panics(a.T, f, forwardArgs(msg, args))
}

// PanicsWithError is the same as [PanicsWithError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithError(errString string, f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(a.T, errString, f, msgAndArgs...)
}

// PanicsWithErrorf is the same as [Assertions.PanicsWithError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithErrorf(errString string, f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(a.T, errString, f, forwardArgs(msg, args))
}

// PanicsWithValue is the same as [PanicsWithValue], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithValue(expected any, f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(a.T, expected, f, msgAndArgs...)
}

// PanicsWithValuef is the same as [Assertions.PanicsWithValue], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithValuef(expected any, f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(a.T, expected, f, forwardArgs(msg, args))
}

// Positive is the same as [Positive], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Positive(e any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Positive(a.T, e, msgAndArgs...)
}

// Positivef is the same as [Assertions.Positive], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Positivef(e any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Positive(a.T, e, forwardArgs(msg, args))
}

// Regexp is the same as [Regexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Regexp(rx any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(a.T, rx, actual, msgAndArgs...)
}

// Regexpf is the same as [Assertions.Regexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Regexpf(rx any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(a.T, rx, actual, forwardArgs(msg, args))
}

// Same is the same as [Same], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Same(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Same(a.T, expected, actual, msgAndArgs...)
}

// Samef is the same as [Assertions.Same], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Samef(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Same(a.T, expected, actual, forwardArgs(msg, args))
}

// Subset is the same as [Subset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Subset(list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Subset(a.T, list, subset, msgAndArgs...)
}

// Subsetf is the same as [Assertions.Subset], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Subsetf(list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Subset(a.T, list, subset, forwardArgs(msg, args))
}

// True is the same as [True], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) True(value bool, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.True(a.T, value, msgAndArgs...)
}

// Truef is the same as [Assertions.True], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Truef(value bool, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.True(a.T, value, forwardArgs(msg, args))
}

// WithinDuration is the same as [WithinDuration], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(a.T, expected, actual, delta, msgAndArgs...)
}

// WithinDurationf is the same as [Assertions.WithinDuration], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(a.T, expected, actual, delta, forwardArgs(msg, args))
}

// WithinRange is the same as [WithinRange], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(a.T, actual, start, end, msgAndArgs...)
}

// WithinRangef is the same as [Assertions.WithinRange], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(a.T, actual, start, end, forwardArgs(msg, args))
}

// YAMLEq is the same as [YAMLEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(a.T, expected, actual, msgAndArgs...)
}

// YAMLEqf is the same as [Assertions.YAMLEq], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(a.T, expected, actual, forwardArgs(msg, args))
}

// YAMLEqBytes is the same as [YAMLEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqBytes(expected []byte, actual []byte, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqBytes(a.T, expected, actual, msgAndArgs...)
}

// YAMLEqBytesf is the same as [Assertions.YAMLEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqBytesf(expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqBytes(a.T, expected, actual, forwardArgs(msg, args))
}

// Zero is the same as [Zero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Zero(i any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Zero(a.T, i, msgAndArgs...)
}

// Zerof is the same as [Assertions.Zero], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Zerof(i any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Zero(a.T, i, forwardArgs(msg, args))
}
