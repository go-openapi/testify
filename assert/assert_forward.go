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
// NOTE: generic assertions are exposed as methods only when built with go1.27 or newer,
// the first release that accepts type parameters on methods.
//
// Upon failure, the test [T] is marked as failed and continues execution.
type Assertions struct {
	T

	o any
}

// New makes a new [Assertions] object for the specified [T] (e.g. [testing.T]).
//
// It may be tuned using [Option].
func New(t T, opts ...Option) *Assertions {
	return &Assertions{
		T: t,
		o: assertions.BuildOptions(opts),
	}
}

// Blocked is the same as [Blocked], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Blocked(ch any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Blocked(a.T, ch, append(msgAndArgs, a.o)...)
}

// Blockedf is the same as [Assertions.Blocked], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Blockedf(ch any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Blocked(a.T, ch, forwardArgs(msg, args, a.o)...)
}

// Condition is the same as [Condition], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Condition(comp func() bool, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Condition(a.T, comp, append(msgAndArgs, a.o)...)
}

// Conditionf is the same as [Assertions.Condition], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Conditionf(comp func() bool, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Condition(a.T, comp, forwardArgs(msg, args, a.o)...)
}

// Contains is the same as [Contains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Contains(s any, contains any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Contains(a.T, s, contains, append(msgAndArgs, a.o)...)
}

// Containsf is the same as [Assertions.Contains], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Containsf(s any, contains any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Contains(a.T, s, contains, forwardArgs(msg, args, a.o)...)
}

// DirExists is the same as [DirExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(a.T, path, append(msgAndArgs, a.o)...)
}

// DirExistsf is the same as [Assertions.DirExists], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(a.T, path, forwardArgs(msg, args, a.o)...)
}

// DirNotExists is the same as [DirNotExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirNotExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirNotExists(a.T, path, append(msgAndArgs, a.o)...)
}

// DirNotExistsf is the same as [Assertions.DirNotExists], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) DirNotExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.DirNotExists(a.T, path, forwardArgs(msg, args, a.o)...)
}

// ElementsMatch is the same as [ElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatch(listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(a.T, listA, listB, append(msgAndArgs, a.o)...)
}

// ElementsMatchf is the same as [Assertions.ElementsMatch], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ElementsMatchf(listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(a.T, listA, listB, forwardArgs(msg, args, a.o)...)
}

// Empty is the same as [Empty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Empty(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Empty(a.T, object, append(msgAndArgs, a.o)...)
}

// Emptyf is the same as [Assertions.Empty], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Emptyf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Empty(a.T, object, forwardArgs(msg, args, a.o)...)
}

// Equal is the same as [Equal], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Equal(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Equal(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// Equalf is the same as [Assertions.Equal], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Equalf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Equal(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// EqualError is the same as [EqualError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualError(err error, errString string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(a.T, err, errString, append(msgAndArgs, a.o)...)
}

// EqualErrorf is the same as [Assertions.EqualError], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualErrorf(err error, errString string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(a.T, err, errString, forwardArgs(msg, args, a.o)...)
}

// EqualExportedValues is the same as [EqualExportedValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualExportedValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// EqualExportedValuesf is the same as [Assertions.EqualExportedValues], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualExportedValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// EqualValues is the same as [EqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// EqualValuesf is the same as [Assertions.EqualValues], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) EqualValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// Error is the same as [Error], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Error(err error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Error(a.T, err, append(msgAndArgs, a.o)...)
}

// Errorf is the same as [Assertions.Error], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Errorf(err error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Error(a.T, err, forwardArgs(msg, args, a.o)...)
}

// ErrorAs is the same as [ErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAs(err error, target any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(a.T, err, target, append(msgAndArgs, a.o)...)
}

// ErrorAsf is the same as [Assertions.ErrorAs], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorAsf(err error, target any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(a.T, err, target, forwardArgs(msg, args, a.o)...)
}

// ErrorContains is the same as [ErrorContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorContains(err error, contains string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(a.T, err, contains, append(msgAndArgs, a.o)...)
}

// ErrorContainsf is the same as [Assertions.ErrorContains], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorContainsf(err error, contains string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(a.T, err, contains, forwardArgs(msg, args, a.o)...)
}

// ErrorIs is the same as [ErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(a.T, err, target, append(msgAndArgs, a.o)...)
}

// ErrorIsf is the same as [Assertions.ErrorIs], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(a.T, err, target, forwardArgs(msg, args, a.o)...)
}

// ErrorNotContains is the same as [ErrorNotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorNotContains(err error, contains string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorNotContains(a.T, err, contains, append(msgAndArgs, a.o)...)
}

// ErrorNotContainsf is the same as [Assertions.ErrorNotContains], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) ErrorNotContainsf(err error, contains string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.ErrorNotContains(a.T, err, contains, forwardArgs(msg, args, a.o)...)
}

// Exactly is the same as [Exactly], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Exactly(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// Exactlyf is the same as [Assertions.Exactly], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Exactlyf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// Fail is the same as [Fail], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Fail(a.T, failureMessage, append(msgAndArgs, a.o)...)
}

// Failf is the same as [Assertions.Fail], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Failf(failureMessage string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Fail(a.T, failureMessage, forwardArgs(msg, args, a.o)...)
}

// FailNow is the same as [FailNow], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(a.T, failureMessage, append(msgAndArgs, a.o)...)
}

// FailNowf is the same as [Assertions.FailNow], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(a.T, failureMessage, forwardArgs(msg, args, a.o)...)
}

// False is the same as [False], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) False(value bool, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.False(a.T, value, append(msgAndArgs, a.o)...)
}

// Falsef is the same as [Assertions.False], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Falsef(value bool, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.False(a.T, value, forwardArgs(msg, args, a.o)...)
}

// FileEmpty is the same as [FileEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileEmpty(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(a.T, path, append(msgAndArgs, a.o)...)
}

// FileEmptyf is the same as [Assertions.FileEmpty], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileEmptyf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(a.T, path, forwardArgs(msg, args, a.o)...)
}

// FileExists is the same as [FileExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(a.T, path, append(msgAndArgs, a.o)...)
}

// FileExistsf is the same as [Assertions.FileExists], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(a.T, path, forwardArgs(msg, args, a.o)...)
}

// FileNotEmpty is the same as [FileNotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotEmpty(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(a.T, path, append(msgAndArgs, a.o)...)
}

// FileNotEmptyf is the same as [Assertions.FileNotEmpty], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotEmptyf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(a.T, path, forwardArgs(msg, args, a.o)...)
}

// FileNotExists is the same as [FileNotExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotExists(path string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotExists(a.T, path, append(msgAndArgs, a.o)...)
}

// FileNotExistsf is the same as [Assertions.FileNotExists], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) FileNotExistsf(path string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.FileNotExists(a.T, path, forwardArgs(msg, args, a.o)...)
}

// Greater is the same as [Greater], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Greater(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Greater(a.T, e1, e2, append(msgAndArgs, a.o)...)
}

// Greaterf is the same as [Assertions.Greater], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Greaterf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Greater(a.T, e1, e2, forwardArgs(msg, args, a.o)...)
}

// GreaterOrEqual is the same as [GreaterOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqual(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(a.T, e1, e2, append(msgAndArgs, a.o)...)
}

// GreaterOrEqualf is the same as [Assertions.GreaterOrEqual], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) GreaterOrEqualf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(a.T, e1, e2, forwardArgs(msg, args, a.o)...)
}

// HTTPBodyContains is the same as [HTTPBodyContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(a.T, handler, method, url, values, str, append(msgAndArgs, a.o)...)
}

// HTTPBodyContainsf is the same as [Assertions.HTTPBodyContains], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(a.T, handler, method, url, values, str, forwardArgs(msg, args, a.o)...)
}

// HTTPBodyNotContains is the same as [HTTPBodyNotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(a.T, handler, method, url, values, str, append(msgAndArgs, a.o)...)
}

// HTTPBodyNotContainsf is the same as [Assertions.HTTPBodyNotContains], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(a.T, handler, method, url, values, str, forwardArgs(msg, args, a.o)...)
}

// HTTPError is the same as [HTTPError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(a.T, handler, method, url, values, append(msgAndArgs, a.o)...)
}

// HTTPErrorf is the same as [Assertions.HTTPError], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(a.T, handler, method, url, values, forwardArgs(msg, args, a.o)...)
}

// HTTPRedirect is the same as [HTTPRedirect], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(a.T, handler, method, url, values, append(msgAndArgs, a.o)...)
}

// HTTPRedirectf is the same as [Assertions.HTTPRedirect], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(a.T, handler, method, url, values, forwardArgs(msg, args, a.o)...)
}

// HTTPStatusCode is the same as [HTTPStatusCode], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(a.T, handler, method, url, values, statuscode, append(msgAndArgs, a.o)...)
}

// HTTPStatusCodef is the same as [Assertions.HTTPStatusCode], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(a.T, handler, method, url, values, statuscode, forwardArgs(msg, args, a.o)...)
}

// HTTPSuccess is the same as [HTTPSuccess], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(a.T, handler, method, url, values, append(msgAndArgs, a.o)...)
}

// HTTPSuccessf is the same as [Assertions.HTTPSuccess], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(a.T, handler, method, url, values, forwardArgs(msg, args, a.o)...)
}

// Implements is the same as [Implements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Implements(interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Implements(a.T, interfaceObject, object, append(msgAndArgs, a.o)...)
}

// Implementsf is the same as [Assertions.Implements], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Implementsf(interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Implements(a.T, interfaceObject, object, forwardArgs(msg, args, a.o)...)
}

// InDelta is the same as [InDelta], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDelta(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(a.T, expected, actual, delta, append(msgAndArgs, a.o)...)
}

// InDeltaf is the same as [Assertions.InDelta], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaf(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(a.T, expected, actual, delta, forwardArgs(msg, args, a.o)...)
}

// InDeltaMapValues is the same as [InDeltaMapValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaMapValues(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(a.T, expected, actual, delta, append(msgAndArgs, a.o)...)
}

// InDeltaMapValuesf is the same as [Assertions.InDeltaMapValues], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaMapValuesf(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(a.T, expected, actual, delta, forwardArgs(msg, args, a.o)...)
}

// InDeltaSlice is the same as [InDeltaSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaSlice(expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(a.T, expected, actual, delta, append(msgAndArgs, a.o)...)
}

// InDeltaSlicef is the same as [Assertions.InDeltaSlice], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InDeltaSlicef(expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(a.T, expected, actual, delta, forwardArgs(msg, args, a.o)...)
}

// InEpsilon is the same as [InEpsilon], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilon(expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(a.T, expected, actual, epsilon, append(msgAndArgs, a.o)...)
}

// InEpsilonf is the same as [Assertions.InEpsilon], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonf(expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(a.T, expected, actual, epsilon, forwardArgs(msg, args, a.o)...)
}

// InEpsilonSlice is the same as [InEpsilonSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSlice(expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(a.T, expected, actual, epsilon, append(msgAndArgs, a.o)...)
}

// InEpsilonSlicef is the same as [Assertions.InEpsilonSlice], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSlicef(expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(a.T, expected, actual, epsilon, forwardArgs(msg, args, a.o)...)
}

// InEpsilonSymmetric is the same as [InEpsilonSymmetric], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSymmetric(x any, y any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSymmetric(a.T, x, y, epsilon, append(msgAndArgs, a.o)...)
}

// InEpsilonSymmetricf is the same as [Assertions.InEpsilonSymmetric], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) InEpsilonSymmetricf(x any, y any, epsilon float64, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSymmetric(a.T, x, y, epsilon, forwardArgs(msg, args, a.o)...)
}

// IsDecreasing is the same as [IsDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(a.T, collection, append(msgAndArgs, a.o)...)
}

// IsDecreasingf is the same as [Assertions.IsDecreasing], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsDecreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(a.T, collection, forwardArgs(msg, args, a.o)...)
}

// IsIncreasing is the same as [IsIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(a.T, collection, append(msgAndArgs, a.o)...)
}

// IsIncreasingf is the same as [Assertions.IsIncreasing], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsIncreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(a.T, collection, forwardArgs(msg, args, a.o)...)
}

// IsNonDecreasing is the same as [IsNonDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(a.T, collection, append(msgAndArgs, a.o)...)
}

// IsNonDecreasingf is the same as [Assertions.IsNonDecreasing], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonDecreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(a.T, collection, forwardArgs(msg, args, a.o)...)
}

// IsNonIncreasing is the same as [IsNonIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasing(collection any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(a.T, collection, append(msgAndArgs, a.o)...)
}

// IsNonIncreasingf is the same as [Assertions.IsNonIncreasing], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNonIncreasingf(collection any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(a.T, collection, forwardArgs(msg, args, a.o)...)
}

// IsNotType is the same as [IsNotType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotType(theType any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(a.T, theType, object, append(msgAndArgs, a.o)...)
}

// IsNotTypef is the same as [Assertions.IsNotType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsNotTypef(theType any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(a.T, theType, object, forwardArgs(msg, args, a.o)...)
}

// IsType is the same as [IsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsType(expectedType any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsType(a.T, expectedType, object, append(msgAndArgs, a.o)...)
}

// IsTypef is the same as [Assertions.IsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) IsTypef(expectedType any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.IsType(a.T, expectedType, object, forwardArgs(msg, args, a.o)...)
}

// JSONEq is the same as [JSONEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// JSONEqf is the same as [Assertions.JSONEq], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// JSONEqBytes is the same as [JSONEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqBytes(expected []byte, actual []byte, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// JSONEqBytesf is the same as [Assertions.JSONEqBytes], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) JSONEqBytesf(expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// Kind is the same as [Kind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Kind(expectedKind reflect.Kind, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Kind(a.T, expectedKind, object, append(msgAndArgs, a.o)...)
}

// Kindf is the same as [Assertions.Kind], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Kindf(expectedKind reflect.Kind, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Kind(a.T, expectedKind, object, forwardArgs(msg, args, a.o)...)
}

// Len is the same as [Len], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Len(object any, length int, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Len(a.T, object, length, append(msgAndArgs, a.o)...)
}

// Lenf is the same as [Assertions.Len], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Lenf(object any, length int, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Len(a.T, object, length, forwardArgs(msg, args, a.o)...)
}

// Less is the same as [Less], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Less(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Less(a.T, e1, e2, append(msgAndArgs, a.o)...)
}

// Lessf is the same as [Assertions.Less], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Lessf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Less(a.T, e1, e2, forwardArgs(msg, args, a.o)...)
}

// LessOrEqual is the same as [LessOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqual(e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(a.T, e1, e2, append(msgAndArgs, a.o)...)
}

// LessOrEqualf is the same as [Assertions.LessOrEqual], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) LessOrEqualf(e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(a.T, e1, e2, forwardArgs(msg, args, a.o)...)
}

// Negative is the same as [Negative], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Negative(e any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Negative(a.T, e, append(msgAndArgs, a.o)...)
}

// Negativef is the same as [Assertions.Negative], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Negativef(e any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Negative(a.T, e, forwardArgs(msg, args, a.o)...)
}

// Nil is the same as [Nil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Nil(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Nil(a.T, object, append(msgAndArgs, a.o)...)
}

// Nilf is the same as [Assertions.Nil], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Nilf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Nil(a.T, object, forwardArgs(msg, args, a.o)...)
}

// NoError is the same as [NoError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoError(err error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoError(a.T, err, append(msgAndArgs, a.o)...)
}

// NoErrorf is the same as [Assertions.NoError], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoErrorf(err error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoError(a.T, err, forwardArgs(msg, args, a.o)...)
}

// NoFileDescriptorLeak is the same as [NoFileDescriptorLeak], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoFileDescriptorLeak(tested func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoFileDescriptorLeak(a.T, tested, append(msgAndArgs, a.o)...)
}

// NoFileDescriptorLeakf is the same as [Assertions.NoFileDescriptorLeak], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoFileDescriptorLeakf(tested func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoFileDescriptorLeak(a.T, tested, forwardArgs(msg, args, a.o)...)
}

// NoGoRoutineLeak is the same as [NoGoRoutineLeak], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoGoRoutineLeak(tested func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoGoRoutineLeak(a.T, tested, append(msgAndArgs, a.o)...)
}

// NoGoRoutineLeakf is the same as [Assertions.NoGoRoutineLeak], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NoGoRoutineLeakf(tested func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NoGoRoutineLeak(a.T, tested, forwardArgs(msg, args, a.o)...)
}

// NotBlocked is the same as [NotBlocked], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotBlocked(ch any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotBlocked(a.T, ch, append(msgAndArgs, a.o)...)
}

// NotBlockedf is the same as [Assertions.NotBlocked], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotBlockedf(ch any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotBlocked(a.T, ch, forwardArgs(msg, args, a.o)...)
}

// NotContains is the same as [NotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotContains(s any, contains any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(a.T, s, contains, append(msgAndArgs, a.o)...)
}

// NotContainsf is the same as [Assertions.NotContains], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotContainsf(s any, contains any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(a.T, s, contains, forwardArgs(msg, args, a.o)...)
}

// NotElementsMatch is the same as [NotElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatch(listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(a.T, listA, listB, append(msgAndArgs, a.o)...)
}

// NotElementsMatchf is the same as [Assertions.NotElementsMatch], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotElementsMatchf(listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(a.T, listA, listB, forwardArgs(msg, args, a.o)...)
}

// NotEmpty is the same as [NotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEmpty(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(a.T, object, append(msgAndArgs, a.o)...)
}

// NotEmptyf is the same as [Assertions.NotEmpty], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEmptyf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(a.T, object, forwardArgs(msg, args, a.o)...)
}

// NotEqual is the same as [NotEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqual(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// NotEqualf is the same as [Assertions.NotEqual], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// NotEqualValues is the same as [NotEqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualValues(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// NotEqualValuesf is the same as [Assertions.NotEqualValues], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotEqualValuesf(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// NotErrorAs is the same as [NotErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAs(err error, target any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(a.T, err, target, append(msgAndArgs, a.o)...)
}

// NotErrorAsf is the same as [Assertions.NotErrorAs], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorAsf(err error, target any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(a.T, err, target, forwardArgs(msg, args, a.o)...)
}

// NotErrorIs is the same as [NotErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(a.T, err, target, append(msgAndArgs, a.o)...)
}

// NotErrorIsf is the same as [Assertions.NotErrorIs], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(a.T, err, target, forwardArgs(msg, args, a.o)...)
}

// NotImplements is the same as [NotImplements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotImplements(interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(a.T, interfaceObject, object, append(msgAndArgs, a.o)...)
}

// NotImplementsf is the same as [Assertions.NotImplements], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotImplementsf(interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(a.T, interfaceObject, object, forwardArgs(msg, args, a.o)...)
}

// NotKind is the same as [NotKind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotKind(expectedKind reflect.Kind, object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotKind(a.T, expectedKind, object, append(msgAndArgs, a.o)...)
}

// NotKindf is the same as [Assertions.NotKind], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotKindf(expectedKind reflect.Kind, object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotKind(a.T, expectedKind, object, forwardArgs(msg, args, a.o)...)
}

// NotNil is the same as [NotNil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotNil(object any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(a.T, object, append(msgAndArgs, a.o)...)
}

// NotNilf is the same as [Assertions.NotNil], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotNilf(object any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(a.T, object, forwardArgs(msg, args, a.o)...)
}

// NotPanics is the same as [NotPanics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotPanics(f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(a.T, f, append(msgAndArgs, a.o)...)
}

// NotPanicsf is the same as [Assertions.NotPanics], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotPanicsf(f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(a.T, f, forwardArgs(msg, args, a.o)...)
}

// NotRegexp is the same as [NotRegexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexp(rx any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(a.T, rx, actual, append(msgAndArgs, a.o)...)
}

// NotRegexpf is the same as [Assertions.NotRegexp], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotRegexpf(rx any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(a.T, rx, actual, forwardArgs(msg, args, a.o)...)
}

// NotSame is the same as [NotSame], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSame(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// NotSamef is the same as [Assertions.NotSame], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSamef(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// NotSubset is the same as [NotSubset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSubset(list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(a.T, list, subset, append(msgAndArgs, a.o)...)
}

// NotSubsetf is the same as [Assertions.NotSubset], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotSubsetf(list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(a.T, list, subset, forwardArgs(msg, args, a.o)...)
}

// NotZero is the same as [NotZero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotZero(i any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(a.T, i, append(msgAndArgs, a.o)...)
}

// NotZerof is the same as [Assertions.NotZero], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) NotZerof(i any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(a.T, i, forwardArgs(msg, args, a.o)...)
}

// Panics is the same as [Panics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Panics(f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Panics(a.T, f, append(msgAndArgs, a.o)...)
}

// Panicsf is the same as [Assertions.Panics], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Panicsf(f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Panics(a.T, f, forwardArgs(msg, args, a.o)...)
}

// PanicsWithError is the same as [PanicsWithError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithError(errString string, f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(a.T, errString, f, append(msgAndArgs, a.o)...)
}

// PanicsWithErrorf is the same as [Assertions.PanicsWithError], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithErrorf(errString string, f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(a.T, errString, f, forwardArgs(msg, args, a.o)...)
}

// PanicsWithValue is the same as [PanicsWithValue], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithValue(expected any, f func(), msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(a.T, expected, f, append(msgAndArgs, a.o)...)
}

// PanicsWithValuef is the same as [Assertions.PanicsWithValue], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) PanicsWithValuef(expected any, f func(), msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(a.T, expected, f, forwardArgs(msg, args, a.o)...)
}

// Positive is the same as [Positive], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Positive(e any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Positive(a.T, e, append(msgAndArgs, a.o)...)
}

// Positivef is the same as [Assertions.Positive], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Positivef(e any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Positive(a.T, e, forwardArgs(msg, args, a.o)...)
}

// Regexp is the same as [Regexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Regexp(rx any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(a.T, rx, actual, append(msgAndArgs, a.o)...)
}

// Regexpf is the same as [Assertions.Regexp], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Regexpf(rx any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(a.T, rx, actual, forwardArgs(msg, args, a.o)...)
}

// Same is the same as [Same], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Same(expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Same(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// Samef is the same as [Assertions.Same], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Samef(expected any, actual any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Same(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// Subset is the same as [Subset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Subset(list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Subset(a.T, list, subset, append(msgAndArgs, a.o)...)
}

// Subsetf is the same as [Assertions.Subset], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Subsetf(list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Subset(a.T, list, subset, forwardArgs(msg, args, a.o)...)
}

// True is the same as [True], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) True(value bool, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.True(a.T, value, append(msgAndArgs, a.o)...)
}

// Truef is the same as [Assertions.True], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Truef(value bool, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.True(a.T, value, forwardArgs(msg, args, a.o)...)
}

// WithinDuration is the same as [WithinDuration], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(a.T, expected, actual, delta, append(msgAndArgs, a.o)...)
}

// WithinDurationf is the same as [Assertions.WithinDuration], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(a.T, expected, actual, delta, forwardArgs(msg, args, a.o)...)
}

// WithinRange is the same as [WithinRange], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(a.T, actual, start, end, append(msgAndArgs, a.o)...)
}

// WithinRangef is the same as [Assertions.WithinRange], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(a.T, actual, start, end, forwardArgs(msg, args, a.o)...)
}

// YAMLEq is the same as [YAMLEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// YAMLEqf is the same as [Assertions.YAMLEq], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// YAMLEqBytes is the same as [YAMLEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqBytes(expected []byte, actual []byte, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqBytes(a.T, expected, actual, append(msgAndArgs, a.o)...)
}

// YAMLEqBytesf is the same as [Assertions.YAMLEqBytes], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) YAMLEqBytesf(expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqBytes(a.T, expected, actual, forwardArgs(msg, args, a.o)...)
}

// Zero is the same as [Zero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Zero(i any, msgAndArgs ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Zero(a.T, i, append(msgAndArgs, a.o)...)
}

// Zerof is the same as [Assertions.Zero], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func (a *Assertions) Zerof(i any, msg string, args ...any) bool {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	return assertions.Zero(a.T, i, forwardArgs(msg, args, a.o)...)
}
