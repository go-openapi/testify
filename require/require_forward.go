// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.
// Generated on 2026-01-11 (version ca82e58) using codegen version v2.1.9-0.20260111184010-ca82e58db12c+dirty [sha: ca82e58db12cbb61bfcae58c3684b3add9599d10]

package require

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
// Upon failure, the test [T] is marked as failed and stops execution.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Condition(comp Comparison, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Condition(a.t, comp, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Conditionf is the same as [Assertions.Condition], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Conditionf(comp Comparison, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Condition(a.t, comp, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Contains is the same as [Contains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Contains(s any, contains any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Contains(a.t, s, contains, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Containsf is the same as [Assertions.Contains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Containsf(s any, contains any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Contains(a.t, s, contains, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// DirExists is the same as [DirExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) DirExists(path string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.DirExists(a.t, path, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// DirExistsf is the same as [Assertions.DirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) DirExistsf(path string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.DirExists(a.t, path, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// ElementsMatch is the same as [ElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ElementsMatch(listA any, listB any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatch(a.t, listA, listB, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// ElementsMatchf is the same as [Assertions.ElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ElementsMatchf(listA any, listB any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatch(a.t, listA, listB, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Empty is the same as [Empty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Empty(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Empty(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Emptyf is the same as [Assertions.Empty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Emptyf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Empty(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Equal is the same as [Equal], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Equal(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Equal(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Equalf is the same as [Assertions.Equal], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Equalf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Equal(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// EqualError is the same as [EqualError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualError(theError error, errString string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EqualError(a.t, theError, errString, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// EqualErrorf is the same as [Assertions.EqualError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualErrorf(theError error, errString string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EqualError(a.t, theError, errString, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// EqualExportedValues is the same as [EqualExportedValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualExportedValues(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EqualExportedValues(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// EqualExportedValuesf is the same as [Assertions.EqualExportedValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualExportedValuesf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EqualExportedValues(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// EqualValues is the same as [EqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualValues(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EqualValues(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// EqualValuesf is the same as [Assertions.EqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualValuesf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EqualValues(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Error is the same as [Error], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Error(err error, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Error(a.t, err, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Errorf is the same as [Assertions.Error], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Errorf(err error, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Error(a.t, err, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// ErrorAs is the same as [ErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorAs(err error, target any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAs(a.t, err, target, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// ErrorAsf is the same as [Assertions.ErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorAsf(err error, target any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAs(a.t, err, target, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// ErrorContains is the same as [ErrorContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorContains(theError error, contains string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorContains(a.t, theError, contains, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// ErrorContainsf is the same as [Assertions.ErrorContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorContainsf(theError error, contains string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorContains(a.t, theError, contains, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// ErrorIs is the same as [ErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorIs(a.t, err, target, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// ErrorIsf is the same as [Assertions.ErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorIs(a.t, err, target, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Eventually is the same as [Eventually], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Eventually(a.t, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Eventuallyf is the same as [Assertions.Eventually], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Eventually(a.t, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// EventuallyWithT is the same as [EventuallyWithT], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EventuallyWithT(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWithT(a.t, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// EventuallyWithTf is the same as [Assertions.EventuallyWithT], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EventuallyWithTf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWithT(a.t, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Exactly is the same as [Exactly], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Exactly(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Exactly(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Exactlyf is the same as [Assertions.Exactly], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Exactlyf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Exactly(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Fail is the same as [Fail], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	_ = assertions.Fail(a.t, failureMessage, msgAndArgs...)

	a.t.FailNow()
}

// Failf is the same as [Assertions.Fail], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Failf(failureMessage string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	_ = assertions.Fail(a.t, failureMessage, forwardArgs(msg, args))

	a.t.FailNow()
}

// FailNow is the same as [FailNow], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	_ = assertions.FailNow(a.t, failureMessage, msgAndArgs...)

	a.t.FailNow()
}

// FailNowf is the same as [Assertions.FailNow], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	_ = assertions.FailNow(a.t, failureMessage, forwardArgs(msg, args))

	a.t.FailNow()
}

// False is the same as [False], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) False(value bool, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.False(a.t, value, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Falsef is the same as [Assertions.False], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Falsef(value bool, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.False(a.t, value, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// FileEmpty is the same as [FileEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileEmpty(path string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.FileEmpty(a.t, path, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// FileEmptyf is the same as [Assertions.FileEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileEmptyf(path string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.FileEmpty(a.t, path, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// FileExists is the same as [FileExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileExists(path string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.FileExists(a.t, path, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// FileExistsf is the same as [Assertions.FileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileExistsf(path string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.FileExists(a.t, path, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// FileNotEmpty is the same as [FileNotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileNotEmpty(path string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.FileNotEmpty(a.t, path, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// FileNotEmptyf is the same as [Assertions.FileNotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileNotEmptyf(path string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.FileNotEmpty(a.t, path, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Greater is the same as [Greater], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Greater(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Greater(a.t, e1, e2, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Greaterf is the same as [Assertions.Greater], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Greaterf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Greater(a.t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// GreaterOrEqual is the same as [GreaterOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterOrEqual(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqual(a.t, e1, e2, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// GreaterOrEqualf is the same as [Assertions.GreaterOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterOrEqualf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqual(a.t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// HTTPBodyContains is the same as [HTTPBodyContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyContains(a.t, handler, method, url, values, str, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// HTTPBodyContainsf is the same as [Assertions.HTTPBodyContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyContains(a.t, handler, method, url, values, str, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// HTTPBodyNotContains is the same as [HTTPBodyNotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyNotContains(a.t, handler, method, url, values, str, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// HTTPBodyNotContainsf is the same as [Assertions.HTTPBodyNotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyNotContains(a.t, handler, method, url, values, str, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// HTTPError is the same as [HTTPError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPError(a.t, handler, method, url, values, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// HTTPErrorf is the same as [Assertions.HTTPError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPError(a.t, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// HTTPRedirect is the same as [HTTPRedirect], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPRedirect(a.t, handler, method, url, values, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// HTTPRedirectf is the same as [Assertions.HTTPRedirect], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPRedirect(a.t, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// HTTPStatusCode is the same as [HTTPStatusCode], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPStatusCode(a.t, handler, method, url, values, statuscode, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// HTTPStatusCodef is the same as [Assertions.HTTPStatusCode], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPStatusCode(a.t, handler, method, url, values, statuscode, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// HTTPSuccess is the same as [HTTPSuccess], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPSuccess(a.t, handler, method, url, values, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// HTTPSuccessf is the same as [Assertions.HTTPSuccess], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPSuccess(a.t, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Implements is the same as [Implements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Implements(interfaceObject any, object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Implements(a.t, interfaceObject, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Implementsf is the same as [Assertions.Implements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Implementsf(interfaceObject any, object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Implements(a.t, interfaceObject, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// InDelta is the same as [InDelta], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDelta(expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InDelta(a.t, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// InDeltaf is the same as [Assertions.InDelta], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaf(expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InDelta(a.t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// InDeltaMapValues is the same as [InDeltaMapValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaMapValues(expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaMapValues(a.t, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// InDeltaMapValuesf is the same as [Assertions.InDeltaMapValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaMapValuesf(expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaMapValues(a.t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// InDeltaSlice is the same as [InDeltaSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaSlice(expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaSlice(a.t, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// InDeltaSlicef is the same as [Assertions.InDeltaSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaSlicef(expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaSlice(a.t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// InEpsilon is the same as [InEpsilon], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilon(expected any, actual any, epsilon float64, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilon(a.t, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// InEpsilonf is the same as [Assertions.InEpsilon], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonf(expected any, actual any, epsilon float64, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilon(a.t, expected, actual, epsilon, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// InEpsilonSlice is the same as [InEpsilonSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonSlice(expected any, actual any, epsilon float64, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSlice(a.t, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// InEpsilonSlicef is the same as [Assertions.InEpsilonSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonSlicef(expected any, actual any, epsilon float64, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSlice(a.t, expected, actual, epsilon, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// IsDecreasing is the same as [IsDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsDecreasing(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasing(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// IsDecreasingf is the same as [Assertions.IsDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsDecreasingf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasing(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// IsIncreasing is the same as [IsIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsIncreasing(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasing(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// IsIncreasingf is the same as [Assertions.IsIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsIncreasingf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasing(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// IsNonDecreasing is the same as [IsNonDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonDecreasing(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasing(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// IsNonDecreasingf is the same as [Assertions.IsNonDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonDecreasingf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasing(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// IsNonIncreasing is the same as [IsNonIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonIncreasing(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasing(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// IsNonIncreasingf is the same as [Assertions.IsNonIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonIncreasingf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasing(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// IsNotType is the same as [IsNotType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNotType(theType any, object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsNotType(a.t, theType, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// IsNotTypef is the same as [Assertions.IsNotType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNotTypef(theType any, object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsNotType(a.t, theType, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// IsType is the same as [IsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsType(expectedType any, object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsType(a.t, expectedType, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// IsTypef is the same as [Assertions.IsType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsTypef(expectedType any, object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.IsType(a.t, expectedType, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// JSONEq is the same as [JSONEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEq(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// JSONEqf is the same as [Assertions.JSONEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEq(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// JSONEqBytes is the same as [JSONEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqBytes(expected []byte, actual []byte, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqBytes(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// JSONEqBytesf is the same as [Assertions.JSONEqBytes], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqBytesf(expected []byte, actual []byte, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqBytes(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Kind is the same as [Kind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Kind(expectedKind reflect.Kind, object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Kind(a.t, expectedKind, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Kindf is the same as [Assertions.Kind], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Kindf(expectedKind reflect.Kind, object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Kind(a.t, expectedKind, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Len is the same as [Len], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Len(object any, length int, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Len(a.t, object, length, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Lenf is the same as [Assertions.Len], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Lenf(object any, length int, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Len(a.t, object, length, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Less is the same as [Less], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Less(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Less(a.t, e1, e2, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Lessf is the same as [Assertions.Less], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Lessf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Less(a.t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// LessOrEqual is the same as [LessOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessOrEqual(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqual(a.t, e1, e2, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// LessOrEqualf is the same as [Assertions.LessOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessOrEqualf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqual(a.t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Negative is the same as [Negative], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Negative(e any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Negative(a.t, e, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Negativef is the same as [Assertions.Negative], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Negativef(e any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Negative(a.t, e, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Never is the same as [Never], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Never(a.t, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Neverf is the same as [Assertions.Never], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Never(a.t, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Nil is the same as [Nil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Nil(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Nil(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Nilf is the same as [Assertions.Nil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Nilf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Nil(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NoDirExists is the same as [NoDirExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoDirExists(path string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NoDirExists(a.t, path, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NoDirExistsf is the same as [Assertions.NoDirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoDirExistsf(path string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NoDirExists(a.t, path, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NoError is the same as [NoError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoError(err error, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NoError(a.t, err, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NoErrorf is the same as [Assertions.NoError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoErrorf(err error, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NoError(a.t, err, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NoFileExists is the same as [NoFileExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoFileExists(path string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NoFileExists(a.t, path, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NoFileExistsf is the same as [Assertions.NoFileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoFileExistsf(path string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NoFileExists(a.t, path, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotContains is the same as [NotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotContains(s any, contains any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotContains(a.t, s, contains, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotContainsf is the same as [Assertions.NotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotContainsf(s any, contains any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotContains(a.t, s, contains, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotElementsMatch is the same as [NotElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotElementsMatch(listA any, listB any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatch(a.t, listA, listB, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotElementsMatchf is the same as [Assertions.NotElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotElementsMatchf(listA any, listB any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatch(a.t, listA, listB, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotEmpty is the same as [NotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEmpty(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotEmpty(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotEmptyf is the same as [Assertions.NotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEmptyf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotEmpty(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotEqual is the same as [NotEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqual(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqual(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotEqualf is the same as [Assertions.NotEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqual(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotEqualValues is the same as [NotEqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualValues(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualValues(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotEqualValuesf is the same as [Assertions.NotEqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualValuesf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualValues(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotErrorAs is the same as [NotErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorAs(err error, target any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAs(a.t, err, target, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotErrorAsf is the same as [Assertions.NotErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorAsf(err error, target any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAs(a.t, err, target, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotErrorIs is the same as [NotErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorIs(a.t, err, target, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotErrorIsf is the same as [Assertions.NotErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorIs(a.t, err, target, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotImplements is the same as [NotImplements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotImplements(interfaceObject any, object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotImplements(a.t, interfaceObject, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotImplementsf is the same as [Assertions.NotImplements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotImplementsf(interfaceObject any, object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotImplements(a.t, interfaceObject, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotKind is the same as [NotKind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotKind(expectedKind reflect.Kind, object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotKind(a.t, expectedKind, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotKindf is the same as [Assertions.NotKind], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotKindf(expectedKind reflect.Kind, object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotKind(a.t, expectedKind, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotNil is the same as [NotNil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotNil(object any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotNil(a.t, object, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotNilf is the same as [Assertions.NotNil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotNilf(object any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotNil(a.t, object, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotPanics is the same as [NotPanics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotPanics(f assertions.PanicTestFunc, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotPanics(a.t, f, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotPanicsf is the same as [Assertions.NotPanics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotPanicsf(f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotPanics(a.t, f, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotRegexp is the same as [NotRegexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotRegexp(rx any, str any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexp(a.t, rx, str, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotRegexpf is the same as [Assertions.NotRegexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotRegexpf(rx any, str any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexp(a.t, rx, str, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotSame is the same as [NotSame], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSame(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotSame(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotSamef is the same as [Assertions.NotSame], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSamef(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotSame(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotSubset is the same as [NotSubset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSubset(list any, subset any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotSubset(a.t, list, subset, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotSubsetf is the same as [Assertions.NotSubset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSubsetf(list any, subset any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotSubset(a.t, list, subset, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// NotZero is the same as [NotZero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotZero(i any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotZero(a.t, i, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// NotZerof is the same as [Assertions.NotZero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotZerof(i any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.NotZero(a.t, i, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Panics is the same as [Panics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Panics(f assertions.PanicTestFunc, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Panics(a.t, f, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Panicsf is the same as [Assertions.Panics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Panicsf(f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Panics(a.t, f, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// PanicsWithError is the same as [PanicsWithError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithError(errString string, f assertions.PanicTestFunc, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithError(a.t, errString, f, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// PanicsWithErrorf is the same as [Assertions.PanicsWithError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithErrorf(errString string, f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithError(a.t, errString, f, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// PanicsWithValue is the same as [PanicsWithValue], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithValue(expected any, f assertions.PanicTestFunc, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithValue(a.t, expected, f, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// PanicsWithValuef is the same as [Assertions.PanicsWithValue], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithValuef(expected any, f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithValue(a.t, expected, f, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Positive is the same as [Positive], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Positive(e any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Positive(a.t, e, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Positivef is the same as [Assertions.Positive], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Positivef(e any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Positive(a.t, e, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Regexp is the same as [Regexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Regexp(rx any, str any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Regexp(a.t, rx, str, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Regexpf is the same as [Assertions.Regexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Regexpf(rx any, str any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Regexp(a.t, rx, str, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Same is the same as [Same], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Same(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Same(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Samef is the same as [Assertions.Same], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Samef(expected any, actual any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Same(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Subset is the same as [Subset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Subset(list any, subset any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Subset(a.t, list, subset, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Subsetf is the same as [Assertions.Subset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Subsetf(list any, subset any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Subset(a.t, list, subset, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// True is the same as [True], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) True(value bool, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.True(a.t, value, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Truef is the same as [Assertions.True], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Truef(value bool, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.True(a.t, value, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// WithinDuration is the same as [WithinDuration], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.WithinDuration(a.t, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// WithinDurationf is the same as [Assertions.WithinDuration], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.WithinDuration(a.t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// WithinRange is the same as [WithinRange], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.WithinRange(a.t, actual, start, end, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// WithinRangef is the same as [Assertions.WithinRange], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.WithinRange(a.t, actual, start, end, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// YAMLEq is the same as [YAMLEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEq(a.t, expected, actual, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// YAMLEqf is the same as [Assertions.YAMLEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEq(a.t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}

// Zero is the same as [Zero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Zero(i any, msgAndArgs ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Zero(a.t, i, msgAndArgs...) {
		return
	}

	a.t.FailNow()
}

// Zerof is the same as [Assertions.Zero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Zerof(i any, msg string, args ...any) {
	if h, ok := a.t.(H); ok {
		h.Helper()
	}
	if assertions.Zero(a.t, i, forwardArgs(msg, args)) {
		return
	}

	a.t.FailNow()
}
