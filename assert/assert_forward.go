// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.
// Generated on 2026-01-11 (version e6b0793) using codegen version v2.1.9-0.20260111152118-e6b0793ba519+dirty [sha: e6b0793ba519fb22dc1887392e1465649a5a95ff]

package assert

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// Assertions exposes all assertion functions as methods.
//
// NOTE: assertion methods with parameterized types (generics) are not supported as methods.
//
// Upon failure, the test [T] is marked as failed and continues execution.
type Assertions struct {
	t T
}

// New makes a new [Assertions] object for the specified [T] (e.g. [testing.T]).
func New(t T) *Assertions {
	return &Assertions{
		t: t,
	}
}

// Condition is the same as [Condition], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Condition(comp Comparison, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Condition(a.t, comp, msgAndArgs...)
}

// Conditionf is the same as [Assertions.Condition], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Conditionf(comp Comparison, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Condition(a.t, comp, forwardArgs(msg, args))
}

// Contains is the same as [Contains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Contains(s any, contains any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Contains(a.t, s, contains, msgAndArgs...)
}

// Containsf is the same as [Assertions.Contains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Containsf(s any, contains any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Contains(a.t, s, contains, forwardArgs(msg, args))
}

// DirExists is the same as [DirExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(a.t, path, msgAndArgs...)
}

// DirExistsf is the same as [Assertions.DirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(a.t, path, forwardArgs(msg, args))
}

// ElementsMatch is the same as [ElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatch(listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(a.t, listA, listB, msgAndArgs...)
}

// ElementsMatchf is the same as [Assertions.ElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatchf(listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(a.t, listA, listB, forwardArgs(msg, args))
}

// Empty is the same as [Empty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Empty(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Empty(a.t, object, msgAndArgs...)
}

// Emptyf is the same as [Assertions.Empty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Emptyf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Empty(a.t, object, forwardArgs(msg, args))
}

// Equal is the same as [Equal], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Equal(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Equal(a.t, expected, actual, msgAndArgs...)
}

// Equalf is the same as [Assertions.Equal], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Equalf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Equal(a.t, expected, actual, forwardArgs(msg, args))
}

// EqualError is the same as [EqualError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(a.t, theError, errString, msgAndArgs...)
}

// EqualErrorf is the same as [Assertions.EqualError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(a.t, theError, errString, forwardArgs(msg, args))
}

// EqualExportedValues is the same as [EqualExportedValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualExportedValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(a.t, expected, actual, msgAndArgs...)
}

// EqualExportedValuesf is the same as [Assertions.EqualExportedValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualExportedValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(a.t, expected, actual, forwardArgs(msg, args))
}

// EqualValues is the same as [EqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(a.t, expected, actual, msgAndArgs...)
}

// EqualValuesf is the same as [Assertions.EqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(a.t, expected, actual, forwardArgs(msg, args))
}

// Error is the same as [Error], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Error(err error, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Error(a.t, err, msgAndArgs...)
}

// Errorf is the same as [Assertions.Error], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Errorf(err error, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Error(a.t, err, forwardArgs(msg, args))
}

// ErrorAs is the same as [ErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAs(err error, target any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(a.t, err, target, msgAndArgs...)
}

// ErrorAsf is the same as [Assertions.ErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAsf(err error, target any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(a.t, err, target, forwardArgs(msg, args))
}

// ErrorContains is the same as [ErrorContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorContains(theError error, contains string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(a.t, theError, contains, msgAndArgs...)
}

// ErrorContainsf is the same as [Assertions.ErrorContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorContainsf(theError error, contains string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(a.t, theError, contains, forwardArgs(msg, args))
}

// ErrorIs is the same as [ErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(a.t, err, target, msgAndArgs...)
}

// ErrorIsf is the same as [Assertions.ErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(a.t, err, target, forwardArgs(msg, args))
}

// Eventually is the same as [Eventually], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Eventually(a.t, condition, waitFor, tick, msgAndArgs...)
}

// Eventuallyf is the same as [Assertions.Eventually], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Eventually(a.t, condition, waitFor, tick, forwardArgs(msg, args))
}

// EventuallyWithT is the same as [EventuallyWithT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EventuallyWithT(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWithT(a.t, condition, waitFor, tick, msgAndArgs...)
}

// EventuallyWithTf is the same as [Assertions.EventuallyWithT], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EventuallyWithTf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWithT(a.t, condition, waitFor, tick, forwardArgs(msg, args))
}

// Exactly is the same as [Exactly], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Exactly(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(a.t, expected, actual, msgAndArgs...)
}

// Exactlyf is the same as [Assertions.Exactly], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Exactlyf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(a.t, expected, actual, forwardArgs(msg, args))
}

// Fail is the same as [Fail], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Fail(a.t, failureMessage, msgAndArgs...)
}

// Failf is the same as [Assertions.Fail], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Failf(failureMessage string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Fail(a.t, failureMessage, forwardArgs(msg, args))
}

// FailNow is the same as [FailNow], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(a.t, failureMessage, msgAndArgs...)
}

// FailNowf is the same as [Assertions.FailNow], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(a.t, failureMessage, forwardArgs(msg, args))
}

// False is the same as [False], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) False(value bool, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.False(a.t, value, msgAndArgs...)
}

// Falsef is the same as [Assertions.False], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Falsef(value bool, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.False(a.t, value, forwardArgs(msg, args))
}

// FileEmpty is the same as [FileEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileEmpty(path string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(a.t, path, msgAndArgs...)
}

// FileEmptyf is the same as [Assertions.FileEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileEmptyf(path string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(a.t, path, forwardArgs(msg, args))
}

// FileExists is the same as [FileExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(a.t, path, msgAndArgs...)
}

// FileExistsf is the same as [Assertions.FileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(a.t, path, forwardArgs(msg, args))
}

// FileNotEmpty is the same as [FileNotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotEmpty(path string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(a.t, path, msgAndArgs...)
}

// FileNotEmptyf is the same as [Assertions.FileNotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotEmptyf(path string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(a.t, path, forwardArgs(msg, args))
}

// Greater is the same as [Greater], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Greater(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Greater(a.t, e1, e2, msgAndArgs...)
}

// Greaterf is the same as [Assertions.Greater], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Greaterf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Greater(a.t, e1, e2, forwardArgs(msg, args))
}

// GreaterOrEqual is the same as [GreaterOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqual(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(a.t, e1, e2, msgAndArgs...)
}

// GreaterOrEqualf is the same as [Assertions.GreaterOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqualf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(a.t, e1, e2, forwardArgs(msg, args))
}

// HTTPBodyContains is the same as [HTTPBodyContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(a.t, handler, method, url, values, str, msgAndArgs...)
}

// HTTPBodyContainsf is the same as [Assertions.HTTPBodyContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(a.t, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPBodyNotContains is the same as [HTTPBodyNotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(a.t, handler, method, url, values, str, msgAndArgs...)
}

// HTTPBodyNotContainsf is the same as [Assertions.HTTPBodyNotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(a.t, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPError is the same as [HTTPError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(a.t, handler, method, url, values, msgAndArgs...)
}

// HTTPErrorf is the same as [Assertions.HTTPError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(a.t, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPRedirect is the same as [HTTPRedirect], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(a.t, handler, method, url, values, msgAndArgs...)
}

// HTTPRedirectf is the same as [Assertions.HTTPRedirect], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(a.t, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPStatusCode is the same as [HTTPStatusCode], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(a.t, handler, method, url, values, statuscode, msgAndArgs...)
}

// HTTPStatusCodef is the same as [Assertions.HTTPStatusCode], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(a.t, handler, method, url, values, statuscode, forwardArgs(msg, args))
}

// HTTPSuccess is the same as [HTTPSuccess], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(a.t, handler, method, url, values, msgAndArgs...)
}

// HTTPSuccessf is the same as [Assertions.HTTPSuccess], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(a.t, handler, method, url, values, forwardArgs(msg, args))
}

// Implements is the same as [Implements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Implements(interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Implements(a.t, interfaceObject, object, msgAndArgs...)
}

// Implementsf is the same as [Assertions.Implements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Implementsf(interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Implements(a.t, interfaceObject, object, forwardArgs(msg, args))
}

// InDelta is the same as [InDelta], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDelta(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(a.t, expected, actual, delta, msgAndArgs...)
}

// InDeltaf is the same as [Assertions.InDelta], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaf(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(a.t, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaMapValues is the same as [InDeltaMapValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaMapValues(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(a.t, expected, actual, delta, msgAndArgs...)
}

// InDeltaMapValuesf is the same as [Assertions.InDeltaMapValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaMapValuesf(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(a.t, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaSlice is the same as [InDeltaSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaSlice(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(a.t, expected, actual, delta, msgAndArgs...)
}

// InDeltaSlicef is the same as [Assertions.InDeltaSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaSlicef(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(a.t, expected, actual, delta, forwardArgs(msg, args))
}

// InEpsilon is the same as [InEpsilon], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilon(expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(a.t, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonf is the same as [Assertions.InEpsilon], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonf(expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(a.t, expected, actual, epsilon, forwardArgs(msg, args))
}

// InEpsilonSlice is the same as [InEpsilonSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSlice(expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(a.t, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonSlicef is the same as [Assertions.InEpsilonSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSlicef(expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(a.t, expected, actual, epsilon, forwardArgs(msg, args))
}

// IsDecreasing is the same as [IsDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasing(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(a.t, object, msgAndArgs...)
}

// IsDecreasingf is the same as [Assertions.IsDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasingf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(a.t, object, forwardArgs(msg, args))
}

// IsIncreasing is the same as [IsIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasing(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(a.t, object, msgAndArgs...)
}

// IsIncreasingf is the same as [Assertions.IsIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasingf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(a.t, object, forwardArgs(msg, args))
}

// IsNonDecreasing is the same as [IsNonDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasing(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(a.t, object, msgAndArgs...)
}

// IsNonDecreasingf is the same as [Assertions.IsNonDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasingf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(a.t, object, forwardArgs(msg, args))
}

// IsNonIncreasing is the same as [IsNonIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasing(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(a.t, object, msgAndArgs...)
}

// IsNonIncreasingf is the same as [Assertions.IsNonIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasingf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(a.t, object, forwardArgs(msg, args))
}

// IsNotType is the same as [IsNotType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotType(theType any, object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(a.t, theType, object, msgAndArgs...)
}

// IsNotTypef is the same as [Assertions.IsNotType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotTypef(theType any, object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(a.t, theType, object, forwardArgs(msg, args))
}

// IsType is the same as [IsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsType(expectedType any, object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsType(a.t, expectedType, object, msgAndArgs...)
}

// IsTypef is the same as [Assertions.IsType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsTypef(expectedType any, object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.IsType(a.t, expectedType, object, forwardArgs(msg, args))
}

// JSONEq is the same as [JSONEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(a.t, expected, actual, msgAndArgs...)
}

// JSONEqf is the same as [Assertions.JSONEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(a.t, expected, actual, forwardArgs(msg, args))
}

// JSONEqBytes is the same as [JSONEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqBytes(expected []byte, actual []byte, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(a.t, expected, actual, msgAndArgs...)
}

// JSONEqBytesf is the same as [Assertions.JSONEqBytes], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqBytesf(expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(a.t, expected, actual, forwardArgs(msg, args))
}

// Len is the same as [Len], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Len(object any, length int, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Len(a.t, object, length, msgAndArgs...)
}

// Lenf is the same as [Assertions.Len], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Lenf(object any, length int, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Len(a.t, object, length, forwardArgs(msg, args))
}

// Less is the same as [Less], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Less(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Less(a.t, e1, e2, msgAndArgs...)
}

// Lessf is the same as [Assertions.Less], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Lessf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Less(a.t, e1, e2, forwardArgs(msg, args))
}

// LessOrEqual is the same as [LessOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqual(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(a.t, e1, e2, msgAndArgs...)
}

// LessOrEqualf is the same as [Assertions.LessOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqualf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(a.t, e1, e2, forwardArgs(msg, args))
}

// Negative is the same as [Negative], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Negative(e any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Negative(a.t, e, msgAndArgs...)
}

// Negativef is the same as [Assertions.Negative], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Negativef(e any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Negative(a.t, e, forwardArgs(msg, args))
}

// Never is the same as [Never], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Never(a.t, condition, waitFor, tick, msgAndArgs...)
}

// Neverf is the same as [Assertions.Never], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Never(a.t, condition, waitFor, tick, forwardArgs(msg, args))
}

// Nil is the same as [Nil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Nil(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Nil(a.t, object, msgAndArgs...)
}

// Nilf is the same as [Assertions.Nil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Nilf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Nil(a.t, object, forwardArgs(msg, args))
}

// NoDirExists is the same as [NoDirExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoDirExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NoDirExists(a.t, path, msgAndArgs...)
}

// NoDirExistsf is the same as [Assertions.NoDirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoDirExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NoDirExists(a.t, path, forwardArgs(msg, args))
}

// NoError is the same as [NoError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoError(err error, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NoError(a.t, err, msgAndArgs...)
}

// NoErrorf is the same as [Assertions.NoError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoErrorf(err error, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NoError(a.t, err, forwardArgs(msg, args))
}

// NoFileExists is the same as [NoFileExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoFileExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NoFileExists(a.t, path, msgAndArgs...)
}

// NoFileExistsf is the same as [Assertions.NoFileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoFileExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NoFileExists(a.t, path, forwardArgs(msg, args))
}

// NotContains is the same as [NotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotContains(s any, contains any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(a.t, s, contains, msgAndArgs...)
}

// NotContainsf is the same as [Assertions.NotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotContainsf(s any, contains any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(a.t, s, contains, forwardArgs(msg, args))
}

// NotElementsMatch is the same as [NotElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatch(listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(a.t, listA, listB, msgAndArgs...)
}

// NotElementsMatchf is the same as [Assertions.NotElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatchf(listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(a.t, listA, listB, forwardArgs(msg, args))
}

// NotEmpty is the same as [NotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEmpty(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(a.t, object, msgAndArgs...)
}

// NotEmptyf is the same as [Assertions.NotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEmptyf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(a.t, object, forwardArgs(msg, args))
}

// NotEqual is the same as [NotEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqual(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(a.t, expected, actual, msgAndArgs...)
}

// NotEqualf is the same as [Assertions.NotEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(a.t, expected, actual, forwardArgs(msg, args))
}

// NotEqualValues is the same as [NotEqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(a.t, expected, actual, msgAndArgs...)
}

// NotEqualValuesf is the same as [Assertions.NotEqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(a.t, expected, actual, forwardArgs(msg, args))
}

// NotErrorAs is the same as [NotErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAs(err error, target any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(a.t, err, target, msgAndArgs...)
}

// NotErrorAsf is the same as [Assertions.NotErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAsf(err error, target any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(a.t, err, target, forwardArgs(msg, args))
}

// NotErrorIs is the same as [NotErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(a.t, err, target, msgAndArgs...)
}

// NotErrorIsf is the same as [Assertions.NotErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(a.t, err, target, forwardArgs(msg, args))
}

// NotImplements is the same as [NotImplements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotImplements(interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(a.t, interfaceObject, object, msgAndArgs...)
}

// NotImplementsf is the same as [Assertions.NotImplements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotImplementsf(interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(a.t, interfaceObject, object, forwardArgs(msg, args))
}

// NotNil is the same as [NotNil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotNil(object any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(a.t, object, msgAndArgs...)
}

// NotNilf is the same as [Assertions.NotNil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotNilf(object any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(a.t, object, forwardArgs(msg, args))
}

// NotPanics is the same as [NotPanics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotPanics(f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(a.t, f, msgAndArgs...)
}

// NotPanicsf is the same as [Assertions.NotPanics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotPanicsf(f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(a.t, f, forwardArgs(msg, args))
}

// NotRegexp is the same as [NotRegexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexp(rx any, str any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(a.t, rx, str, msgAndArgs...)
}

// NotRegexpf is the same as [Assertions.NotRegexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexpf(rx any, str any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(a.t, rx, str, forwardArgs(msg, args))
}

// NotSame is the same as [NotSame], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSame(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(a.t, expected, actual, msgAndArgs...)
}

// NotSamef is the same as [Assertions.NotSame], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSamef(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(a.t, expected, actual, forwardArgs(msg, args))
}

// NotSubset is the same as [NotSubset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSubset(list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(a.t, list, subset, msgAndArgs...)
}

// NotSubsetf is the same as [Assertions.NotSubset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSubsetf(list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(a.t, list, subset, forwardArgs(msg, args))
}

// NotZero is the same as [NotZero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotZero(i any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(a.t, i, msgAndArgs...)
}

// NotZerof is the same as [Assertions.NotZero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotZerof(i any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(a.t, i, forwardArgs(msg, args))
}

// Panics is the same as [Panics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Panics(f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Panics(a.t, f, msgAndArgs...)
}

// Panicsf is the same as [Assertions.Panics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Panicsf(f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Panics(a.t, f, forwardArgs(msg, args))
}

// PanicsWithError is the same as [PanicsWithError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithError(errString string, f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(a.t, errString, f, msgAndArgs...)
}

// PanicsWithErrorf is the same as [Assertions.PanicsWithError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithErrorf(errString string, f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(a.t, errString, f, forwardArgs(msg, args))
}

// PanicsWithValue is the same as [PanicsWithValue], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithValue(expected any, f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(a.t, expected, f, msgAndArgs...)
}

// PanicsWithValuef is the same as [Assertions.PanicsWithValue], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithValuef(expected any, f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(a.t, expected, f, forwardArgs(msg, args))
}

// Positive is the same as [Positive], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Positive(e any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Positive(a.t, e, msgAndArgs...)
}

// Positivef is the same as [Assertions.Positive], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Positivef(e any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Positive(a.t, e, forwardArgs(msg, args))
}

// Regexp is the same as [Regexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Regexp(rx any, str any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(a.t, rx, str, msgAndArgs...)
}

// Regexpf is the same as [Assertions.Regexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Regexpf(rx any, str any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(a.t, rx, str, forwardArgs(msg, args))
}

// Same is the same as [Same], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Same(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Same(a.t, expected, actual, msgAndArgs...)
}

// Samef is the same as [Assertions.Same], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Samef(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Same(a.t, expected, actual, forwardArgs(msg, args))
}

// Subset is the same as [Subset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Subset(list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Subset(a.t, list, subset, msgAndArgs...)
}

// Subsetf is the same as [Assertions.Subset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Subsetf(list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Subset(a.t, list, subset, forwardArgs(msg, args))
}

// True is the same as [True], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) True(value bool, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.True(a.t, value, msgAndArgs...)
}

// Truef is the same as [Assertions.True], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Truef(value bool, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.True(a.t, value, forwardArgs(msg, args))
}

// WithinDuration is the same as [WithinDuration], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(a.t, expected, actual, delta, msgAndArgs...)
}

// WithinDurationf is the same as [Assertions.WithinDuration], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(a.t, expected, actual, delta, forwardArgs(msg, args))
}

// WithinRange is the same as [WithinRange], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(a.t, actual, start, end, msgAndArgs...)
}

// WithinRangef is the same as [Assertions.WithinRange], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(a.t, actual, start, end, forwardArgs(msg, args))
}

// YAMLEq is the same as [YAMLEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(a.t, expected, actual, msgAndArgs...)
}

// YAMLEqf is the same as [Assertions.YAMLEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(a.t, expected, actual, forwardArgs(msg, args))
}

// Zero is the same as [Zero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Zero(i any, msgAndArgs ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Zero(a.t, i, msgAndArgs...)
}

// Zerof is the same as [Assertions.Zero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Zerof(i any, msg string, args ...any) bool {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	return assertions.Zero(a.t, i, forwardArgs(msg, args))
}
