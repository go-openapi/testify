// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

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
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Condition(comp Comparison, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Condition(a.T, comp, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Conditionf is the same as [Assertions.Condition], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Conditionf(comp Comparison, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Condition(a.T, comp, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Contains is the same as [Contains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Contains(s any, contains any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Contains(a.T, s, contains, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Containsf is the same as [Assertions.Contains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Containsf(s any, contains any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Contains(a.T, s, contains, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// DirExists is the same as [DirExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) DirExists(path string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.DirExists(a.T, path, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// DirExistsf is the same as [Assertions.DirExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) DirExistsf(path string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.DirExists(a.T, path, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// DirNotExists is the same as [DirNotExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) DirNotExists(path string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.DirNotExists(a.T, path, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// DirNotExistsf is the same as [Assertions.DirNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) DirNotExistsf(path string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.DirNotExists(a.T, path, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// ElementsMatch is the same as [ElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ElementsMatch(listA any, listB any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatch(a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// ElementsMatchf is the same as [Assertions.ElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ElementsMatchf(listA any, listB any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatch(a.T, listA, listB, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Empty is the same as [Empty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Empty(object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Empty(a.T, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Emptyf is the same as [Assertions.Empty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Emptyf(object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Empty(a.T, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Equal is the same as [Equal], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Equal(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Equal(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Equalf is the same as [Assertions.Equal], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Equalf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Equal(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// EqualError is the same as [EqualError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualError(err error, errString string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualError(a.T, err, errString, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// EqualErrorf is the same as [Assertions.EqualError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualErrorf(err error, errString string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualError(a.T, err, errString, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// EqualExportedValues is the same as [EqualExportedValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualExportedValues(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualExportedValues(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// EqualExportedValuesf is the same as [Assertions.EqualExportedValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualExportedValuesf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualExportedValues(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// EqualValues is the same as [EqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualValues(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualValues(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// EqualValuesf is the same as [Assertions.EqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EqualValuesf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EqualValues(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Error is the same as [Error], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Error(err error, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Error(a.T, err, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Errorf is the same as [Assertions.Error], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Errorf(err error, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Error(a.T, err, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// ErrorAs is the same as [ErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorAs(err error, target any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAs(a.T, err, target, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// ErrorAsf is the same as [Assertions.ErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorAsf(err error, target any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAs(a.T, err, target, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// ErrorContains is the same as [ErrorContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorContains(err error, contains string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorContains(a.T, err, contains, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// ErrorContainsf is the same as [Assertions.ErrorContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorContainsf(err error, contains string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorContains(a.T, err, contains, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// ErrorIs is the same as [ErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorIs(err error, target error, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorIs(a.T, err, target, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// ErrorIsf is the same as [Assertions.ErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) ErrorIsf(err error, target error, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.ErrorIs(a.T, err, target, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Eventually is the same as [Eventually], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Eventually(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Eventually(a.T, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Eventuallyf is the same as [Assertions.Eventually], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Eventuallyf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Eventually(a.T, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// EventuallyWith is the same as [EventuallyWith], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EventuallyWith(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWith(a.T, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// EventuallyWithf is the same as [Assertions.EventuallyWith], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) EventuallyWithf(condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWith(a.T, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Exactly is the same as [Exactly], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Exactly(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Exactly(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Exactlyf is the same as [Assertions.Exactly], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Exactlyf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Exactly(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Fail is the same as [Fail], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Fail(failureMessage string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	_ = assertions.Fail(a.T, failureMessage, msgAndArgs...)

	a.T.FailNow()
}

// Failf is the same as [Assertions.Fail], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Failf(failureMessage string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	_ = assertions.Fail(a.T, failureMessage, forwardArgs(msg, args))

	a.T.FailNow()
}

// FailNow is the same as [FailNow], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FailNow(failureMessage string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	_ = assertions.FailNow(a.T, failureMessage, msgAndArgs...)

	a.T.FailNow()
}

// FailNowf is the same as [Assertions.FailNow], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FailNowf(failureMessage string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	_ = assertions.FailNow(a.T, failureMessage, forwardArgs(msg, args))

	a.T.FailNow()
}

// False is the same as [False], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) False(value bool, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.False(a.T, value, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Falsef is the same as [Assertions.False], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Falsef(value bool, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.False(a.T, value, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// FileEmpty is the same as [FileEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileEmpty(path string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileEmpty(a.T, path, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// FileEmptyf is the same as [Assertions.FileEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileEmptyf(path string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileEmpty(a.T, path, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// FileExists is the same as [FileExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileExists(path string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileExists(a.T, path, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// FileExistsf is the same as [Assertions.FileExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileExistsf(path string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileExists(a.T, path, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// FileNotEmpty is the same as [FileNotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileNotEmpty(path string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileNotEmpty(a.T, path, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// FileNotEmptyf is the same as [Assertions.FileNotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileNotEmptyf(path string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileNotEmpty(a.T, path, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// FileNotExists is the same as [FileNotExists], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileNotExists(path string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileNotExists(a.T, path, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// FileNotExistsf is the same as [Assertions.FileNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) FileNotExistsf(path string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.FileNotExists(a.T, path, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Greater is the same as [Greater], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Greater(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Greater(a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Greaterf is the same as [Assertions.Greater], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Greaterf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Greater(a.T, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// GreaterOrEqual is the same as [GreaterOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterOrEqual(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqual(a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// GreaterOrEqualf is the same as [Assertions.GreaterOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) GreaterOrEqualf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqual(a.T, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// HTTPBodyContains is the same as [HTTPBodyContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyContains(a.T, handler, method, url, values, str, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// HTTPBodyContainsf is the same as [Assertions.HTTPBodyContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyContains(a.T, handler, method, url, values, str, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// HTTPBodyNotContains is the same as [HTTPBodyNotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyNotContains(handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyNotContains(a.T, handler, method, url, values, str, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// HTTPBodyNotContainsf is the same as [Assertions.HTTPBodyNotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPBodyNotContainsf(handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyNotContains(a.T, handler, method, url, values, str, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// HTTPError is the same as [HTTPError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPError(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPError(a.T, handler, method, url, values, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// HTTPErrorf is the same as [Assertions.HTTPError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPErrorf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPError(a.T, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// HTTPRedirect is the same as [HTTPRedirect], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPRedirect(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPRedirect(a.T, handler, method, url, values, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// HTTPRedirectf is the same as [Assertions.HTTPRedirect], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPRedirectf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPRedirect(a.T, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// HTTPStatusCode is the same as [HTTPStatusCode], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPStatusCode(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPStatusCode(a.T, handler, method, url, values, statuscode, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// HTTPStatusCodef is the same as [Assertions.HTTPStatusCode], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPStatusCodef(handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPStatusCode(a.T, handler, method, url, values, statuscode, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// HTTPSuccess is the same as [HTTPSuccess], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPSuccess(handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPSuccess(a.T, handler, method, url, values, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// HTTPSuccessf is the same as [Assertions.HTTPSuccess], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) HTTPSuccessf(handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.HTTPSuccess(a.T, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Implements is the same as [Implements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Implements(interfaceObject any, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Implements(a.T, interfaceObject, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Implementsf is the same as [Assertions.Implements], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Implementsf(interfaceObject any, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Implements(a.T, interfaceObject, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// InDelta is the same as [InDelta], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDelta(expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDelta(a.T, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InDeltaf is the same as [Assertions.InDelta], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaf(expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDelta(a.T, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// InDeltaMapValues is the same as [InDeltaMapValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaMapValues(expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaMapValues(a.T, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InDeltaMapValuesf is the same as [Assertions.InDeltaMapValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaMapValuesf(expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaMapValues(a.T, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// InDeltaSlice is the same as [InDeltaSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaSlice(expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaSlice(a.T, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InDeltaSlicef is the same as [Assertions.InDeltaSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InDeltaSlicef(expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaSlice(a.T, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// InEpsilon is the same as [InEpsilon], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilon(expected any, actual any, epsilon float64, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilon(a.T, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InEpsilonf is the same as [Assertions.InEpsilon], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonf(expected any, actual any, epsilon float64, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilon(a.T, expected, actual, epsilon, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// InEpsilonSlice is the same as [InEpsilonSlice], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonSlice(expected any, actual any, epsilon float64, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSlice(a.T, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// InEpsilonSlicef is the same as [Assertions.InEpsilonSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) InEpsilonSlicef(expected any, actual any, epsilon float64, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSlice(a.T, expected, actual, epsilon, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// IsDecreasing is the same as [IsDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsDecreasing(collection any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasing(a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsDecreasingf is the same as [Assertions.IsDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsDecreasingf(collection any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasing(a.T, collection, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// IsIncreasing is the same as [IsIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsIncreasing(collection any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasing(a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsIncreasingf is the same as [Assertions.IsIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsIncreasingf(collection any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasing(a.T, collection, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// IsNonDecreasing is the same as [IsNonDecreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonDecreasing(collection any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasing(a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsNonDecreasingf is the same as [Assertions.IsNonDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonDecreasingf(collection any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasing(a.T, collection, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// IsNonIncreasing is the same as [IsNonIncreasing], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonIncreasing(collection any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasing(a.T, collection, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsNonIncreasingf is the same as [Assertions.IsNonIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNonIncreasingf(collection any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasing(a.T, collection, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// IsNotType is the same as [IsNotType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNotType(theType any, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNotType(a.T, theType, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsNotTypef is the same as [Assertions.IsNotType], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsNotTypef(theType any, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsNotType(a.T, theType, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// IsType is the same as [IsType], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsType(expectedType any, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsType(a.T, expectedType, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// IsTypef is the same as [Assertions.IsType], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) IsTypef(expectedType any, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.IsType(a.T, expectedType, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// JSONEq is the same as [JSONEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEq(expected string, actual string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONEq(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// JSONEqf is the same as [Assertions.JSONEq], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqf(expected string, actual string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONEq(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// JSONEqBytes is the same as [JSONEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqBytes(expected []byte, actual []byte, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqBytes(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// JSONEqBytesf is the same as [Assertions.JSONEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) JSONEqBytesf(expected []byte, actual []byte, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqBytes(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Kind is the same as [Kind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Kind(expectedKind reflect.Kind, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Kind(a.T, expectedKind, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Kindf is the same as [Assertions.Kind], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Kindf(expectedKind reflect.Kind, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Kind(a.T, expectedKind, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Len is the same as [Len], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Len(object any, length int, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Len(a.T, object, length, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Lenf is the same as [Assertions.Len], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Lenf(object any, length int, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Len(a.T, object, length, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Less is the same as [Less], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Less(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Less(a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Lessf is the same as [Assertions.Less], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Lessf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Less(a.T, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// LessOrEqual is the same as [LessOrEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessOrEqual(e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqual(a.T, e1, e2, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// LessOrEqualf is the same as [Assertions.LessOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) LessOrEqualf(e1 any, e2 any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqual(a.T, e1, e2, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Negative is the same as [Negative], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Negative(e any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Negative(a.T, e, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Negativef is the same as [Assertions.Negative], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Negativef(e any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Negative(a.T, e, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Never is the same as [Never], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Never(condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Never(a.T, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Neverf is the same as [Assertions.Never], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Neverf(condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Never(a.T, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Nil is the same as [Nil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Nil(object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Nil(a.T, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Nilf is the same as [Assertions.Nil], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Nilf(object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Nil(a.T, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NoError is the same as [NoError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoError(err error, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NoError(a.T, err, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NoErrorf is the same as [Assertions.NoError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoErrorf(err error, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NoError(a.T, err, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NoFileDescriptorLeak is the same as [NoFileDescriptorLeak], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoFileDescriptorLeak(tested func(), msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NoFileDescriptorLeak(a.T, tested, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NoFileDescriptorLeakf is the same as [Assertions.NoFileDescriptorLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoFileDescriptorLeakf(tested func(), msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NoFileDescriptorLeak(a.T, tested, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NoGoRoutineLeak is the same as [NoGoRoutineLeak], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoGoRoutineLeak(tested func(), msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NoGoRoutineLeak(a.T, tested, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NoGoRoutineLeakf is the same as [Assertions.NoGoRoutineLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NoGoRoutineLeakf(tested func(), msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NoGoRoutineLeak(a.T, tested, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotContains is the same as [NotContains], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotContains(s any, contains any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotContains(a.T, s, contains, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotContainsf is the same as [Assertions.NotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotContainsf(s any, contains any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotContains(a.T, s, contains, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotElementsMatch is the same as [NotElementsMatch], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotElementsMatch(listA any, listB any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatch(a.T, listA, listB, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotElementsMatchf is the same as [Assertions.NotElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotElementsMatchf(listA any, listB any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatch(a.T, listA, listB, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotEmpty is the same as [NotEmpty], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEmpty(object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEmpty(a.T, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotEmptyf is the same as [Assertions.NotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEmptyf(object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEmpty(a.T, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotEqual is the same as [NotEqual], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqual(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEqual(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotEqualf is the same as [Assertions.NotEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEqual(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotEqualValues is the same as [NotEqualValues], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualValues(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualValues(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotEqualValuesf is the same as [Assertions.NotEqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotEqualValuesf(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualValues(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotErrorAs is the same as [NotErrorAs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorAs(err error, target any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAs(a.T, err, target, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotErrorAsf is the same as [Assertions.NotErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorAsf(err error, target any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAs(a.T, err, target, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotErrorIs is the same as [NotErrorIs], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorIs(err error, target error, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorIs(a.T, err, target, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotErrorIsf is the same as [Assertions.NotErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotErrorIsf(err error, target error, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorIs(a.T, err, target, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotImplements is the same as [NotImplements], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotImplements(interfaceObject any, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotImplements(a.T, interfaceObject, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotImplementsf is the same as [Assertions.NotImplements], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotImplementsf(interfaceObject any, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotImplements(a.T, interfaceObject, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotKind is the same as [NotKind], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotKind(expectedKind reflect.Kind, object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotKind(a.T, expectedKind, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotKindf is the same as [Assertions.NotKind], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotKindf(expectedKind reflect.Kind, object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotKind(a.T, expectedKind, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotNil is the same as [NotNil], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotNil(object any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotNil(a.T, object, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotNilf is the same as [Assertions.NotNil], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotNilf(object any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotNil(a.T, object, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotPanics is the same as [NotPanics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotPanics(f func(), msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotPanics(a.T, f, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotPanicsf is the same as [Assertions.NotPanics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotPanicsf(f func(), msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotPanics(a.T, f, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotRegexp is the same as [NotRegexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotRegexp(rx any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexp(a.T, rx, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotRegexpf is the same as [Assertions.NotRegexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotRegexpf(rx any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexp(a.T, rx, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotSame is the same as [NotSame], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSame(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSame(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotSamef is the same as [Assertions.NotSame], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSamef(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSame(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotSubset is the same as [NotSubset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSubset(list any, subset any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSubset(a.T, list, subset, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotSubsetf is the same as [Assertions.NotSubset], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotSubsetf(list any, subset any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotSubset(a.T, list, subset, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// NotZero is the same as [NotZero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotZero(i any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotZero(a.T, i, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// NotZerof is the same as [Assertions.NotZero], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) NotZerof(i any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.NotZero(a.T, i, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Panics is the same as [Panics], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Panics(f func(), msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Panics(a.T, f, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Panicsf is the same as [Assertions.Panics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Panicsf(f func(), msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Panics(a.T, f, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// PanicsWithError is the same as [PanicsWithError], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithError(errString string, f func(), msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithError(a.T, errString, f, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// PanicsWithErrorf is the same as [Assertions.PanicsWithError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithErrorf(errString string, f func(), msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithError(a.T, errString, f, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// PanicsWithValue is the same as [PanicsWithValue], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithValue(expected any, f func(), msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithValue(a.T, expected, f, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// PanicsWithValuef is the same as [Assertions.PanicsWithValue], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) PanicsWithValuef(expected any, f func(), msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithValue(a.T, expected, f, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Positive is the same as [Positive], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Positive(e any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Positive(a.T, e, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Positivef is the same as [Assertions.Positive], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Positivef(e any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Positive(a.T, e, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Regexp is the same as [Regexp], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Regexp(rx any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Regexp(a.T, rx, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Regexpf is the same as [Assertions.Regexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Regexpf(rx any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Regexp(a.T, rx, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Same is the same as [Same], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Same(expected any, actual any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Same(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Samef is the same as [Assertions.Same], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Samef(expected any, actual any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Same(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Subset is the same as [Subset], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Subset(list any, subset any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Subset(a.T, list, subset, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Subsetf is the same as [Assertions.Subset], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Subsetf(list any, subset any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Subset(a.T, list, subset, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// True is the same as [True], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) True(value bool, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.True(a.T, value, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Truef is the same as [Assertions.True], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Truef(value bool, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.True(a.T, value, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// WithinDuration is the same as [WithinDuration], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinDuration(expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.WithinDuration(a.T, expected, actual, delta, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// WithinDurationf is the same as [Assertions.WithinDuration], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinDurationf(expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.WithinDuration(a.T, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// WithinRange is the same as [WithinRange], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinRange(actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.WithinRange(a.T, actual, start, end, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// WithinRangef is the same as [Assertions.WithinRange], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) WithinRangef(actual time.Time, start time.Time, end time.Time, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.WithinRange(a.T, actual, start, end, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// YAMLEq is the same as [YAMLEq], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEq(expected string, actual string, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEq(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// YAMLEqf is the same as [Assertions.YAMLEq], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEqf(expected string, actual string, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEq(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// YAMLEqBytes is the same as [YAMLEqBytes], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEqBytes(expected []byte, actual []byte, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqBytes(a.T, expected, actual, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// YAMLEqBytesf is the same as [Assertions.YAMLEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) YAMLEqBytesf(expected []byte, actual []byte, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqBytes(a.T, expected, actual, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}

// Zero is the same as [Zero], as a method rather than a package-level function.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Zero(i any, msgAndArgs ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Zero(a.T, i, msgAndArgs...) {
		return
	}

	a.T.FailNow()
}

// Zerof is the same as [Assertions.Zero], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func (a *Assertions) Zerof(i any, msg string, args ...any) {
	if h, ok := a.T.(H); ok {
		h.Helper()
	}
	if assertions.Zero(a.T, i, forwardArgs(msg, args)) {
		return
	}

	a.T.FailNow()
}
