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

// Conditionf is the same as [Condition], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Conditionf(t T, comp Comparison, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Condition(t, comp, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Containsf is the same as [Contains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Containsf(t T, s any, contains any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Contains(t, s, contains, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// DirExistsf is the same as [DirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func DirExistsf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.DirExists(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// ElementsMatchf is the same as [ElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ElementsMatchf(t T, listA any, listB any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatch(t, listA, listB, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Emptyf is the same as [Empty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Emptyf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Empty(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Equalf is the same as [Equal], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Equalf(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Equal(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// EqualErrorf is the same as [EqualError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualErrorf(t T, theError error, errString string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualError(t, theError, errString, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// EqualExportedValuesf is the same as [EqualExportedValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualExportedValuesf(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualExportedValues(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// EqualValuesf is the same as [EqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualValuesf(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualValues(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Errorf is the same as [Error], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Errorf(t T, err error, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Error(t, err, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// ErrorAsf is the same as [ErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorAsf(t T, err error, target any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAs(t, err, target, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// ErrorContainsf is the same as [ErrorContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorContainsf(t T, theError error, contains string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorContains(t, theError, contains, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// ErrorIsf is the same as [ErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorIsf(t T, err error, target error, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorIs(t, err, target, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Eventuallyf is the same as [Eventually], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Eventuallyf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Eventually(t, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// EventuallyWithTf is the same as [EventuallyWithT], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EventuallyWithTf(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWithT(t, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Exactlyf is the same as [Exactly], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Exactlyf(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Exactly(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Failf is the same as [Fail], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Failf(t T, failureMessage string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	_ = assertions.Fail(t, failureMessage, forwardArgs(msg, args))

	t.FailNow()
}

// FailNowf is the same as [FailNow], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FailNowf(t T, failureMessage string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	_ = assertions.FailNow(t, failureMessage, forwardArgs(msg, args))

	t.FailNow()
}

// Falsef is the same as [False], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Falsef(t T, value bool, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.False(t, value, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// FileEmptyf is the same as [FileEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FileEmptyf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileEmpty(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// FileExistsf is the same as [FileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FileExistsf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileExists(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// FileNotEmptyf is the same as [FileNotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FileNotEmptyf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileNotEmpty(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Greaterf is the same as [Greater], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Greaterf(t T, e1 any, e2 any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Greater(t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// GreaterOrEqualf is the same as [GreaterOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func GreaterOrEqualf(t T, e1 any, e2 any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqual(t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// HTTPBodyContainsf is the same as [HTTPBodyContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPBodyContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyContains(t, handler, method, url, values, str, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// HTTPBodyNotContainsf is the same as [HTTPBodyNotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPBodyNotContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyNotContains(t, handler, method, url, values, str, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// HTTPErrorf is the same as [HTTPError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPErrorf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPError(t, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// HTTPRedirectf is the same as [HTTPRedirect], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPRedirectf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPRedirect(t, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// HTTPStatusCodef is the same as [HTTPStatusCode], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPStatusCodef(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPStatusCode(t, handler, method, url, values, statuscode, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// HTTPSuccessf is the same as [HTTPSuccess], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPSuccessf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPSuccess(t, handler, method, url, values, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Implementsf is the same as [Implements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Implementsf(t T, interfaceObject any, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Implements(t, interfaceObject, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// InDeltaf is the same as [InDelta], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InDeltaf(t T, expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDelta(t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// InDeltaMapValuesf is the same as [InDeltaMapValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InDeltaMapValuesf(t T, expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaMapValues(t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// InDeltaSlicef is the same as [InDeltaSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InDeltaSlicef(t T, expected any, actual any, delta float64, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaSlice(t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// InEpsilonf is the same as [InEpsilon], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InEpsilonf(t T, expected any, actual any, epsilon float64, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilon(t, expected, actual, epsilon, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// InEpsilonSlicef is the same as [InEpsilonSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InEpsilonSlicef(t T, expected any, actual any, epsilon float64, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSlice(t, expected, actual, epsilon, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsDecreasingf is the same as [IsDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsDecreasingf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasing(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsIncreasingf is the same as [IsIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsIncreasingf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasing(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNonDecreasingf is the same as [IsNonDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonDecreasingf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasing(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNonIncreasingf is the same as [IsNonIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonIncreasingf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasing(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNotTypef is the same as [IsNotType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNotTypef(t T, theType any, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNotType(t, theType, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsTypef is the same as [IsType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsTypef(t T, expectedType any, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsType(t, expectedType, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// JSONEqf is the same as [JSONEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONEqf(t T, expected string, actual string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEq(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// JSONEqBytesf is the same as [JSONEqBytes], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqBytes(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Kindf is the same as [Kind], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Kindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Kind(t, expectedKind, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Lenf is the same as [Len], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Lenf(t T, object any, length int, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Len(t, object, length, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Lessf is the same as [Less], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Lessf(t T, e1 any, e2 any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Less(t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// LessOrEqualf is the same as [LessOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func LessOrEqualf(t T, e1 any, e2 any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqual(t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Negativef is the same as [Negative], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Negativef(t T, e any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Negative(t, e, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Neverf is the same as [Never], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Neverf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Never(t, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Nilf is the same as [Nil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Nilf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Nil(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NoDirExistsf is the same as [NoDirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NoDirExistsf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoDirExists(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NoErrorf is the same as [NoError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NoErrorf(t T, err error, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoError(t, err, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NoFileExistsf is the same as [NoFileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NoFileExistsf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoFileExists(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotContainsf is the same as [NotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotContainsf(t T, s any, contains any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotContains(t, s, contains, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotElementsMatchf is the same as [NotElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotElementsMatchf(t T, listA any, listB any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatch(t, listA, listB, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotEmptyf is the same as [NotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEmptyf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEmpty(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotEqualf is the same as [NotEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEqualf(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqual(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotEqualValuesf is the same as [NotEqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEqualValuesf(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualValues(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotErrorAsf is the same as [NotErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotErrorAsf(t T, err error, target any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAs(t, err, target, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotErrorIsf is the same as [NotErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotErrorIsf(t T, err error, target error, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorIs(t, err, target, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotImplementsf is the same as [NotImplements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotImplementsf(t T, interfaceObject any, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotImplements(t, interfaceObject, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotKindf is the same as [NotKind], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotKindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotKind(t, expectedKind, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotNilf is the same as [NotNil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotNilf(t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotNil(t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotPanicsf is the same as [NotPanics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotPanicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotPanics(t, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotRegexpf is the same as [NotRegexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotRegexpf(t T, rx any, str any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexp(t, rx, str, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotSamef is the same as [NotSame], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSamef(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSame(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotSubsetf is the same as [NotSubset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSubsetf(t T, list any, subset any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSubset(t, list, subset, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotZerof is the same as [NotZero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotZerof(t T, i any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotZero(t, i, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Panicsf is the same as [Panics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Panicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Panics(t, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// PanicsWithErrorf is the same as [PanicsWithError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func PanicsWithErrorf(t T, errString string, f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithError(t, errString, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// PanicsWithValuef is the same as [PanicsWithValue], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func PanicsWithValuef(t T, expected any, f assertions.PanicTestFunc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithValue(t, expected, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Positivef is the same as [Positive], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Positivef(t T, e any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Positive(t, e, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Regexpf is the same as [Regexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Regexpf(t T, rx any, str any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Regexp(t, rx, str, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Samef is the same as [Same], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Samef(t T, expected any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Same(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Subsetf is the same as [Subset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Subsetf(t T, list any, subset any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Subset(t, list, subset, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Truef is the same as [True], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Truef(t T, value bool, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.True(t, value, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// WithinDurationf is the same as [WithinDuration], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func WithinDurationf(t T, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.WithinDuration(t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// WithinRangef is the same as [WithinRange], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func WithinRangef(t T, actual time.Time, start time.Time, end time.Time, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.WithinRange(t, actual, start, end, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// YAMLEqf is the same as [YAMLEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLEqf(t T, expected string, actual string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEq(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Zerof is the same as [Zero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Zerof(t T, i any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Zero(t, i, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

func forwardArgs(msg string, args []any) []any {
	result := make([]any, len(args)+1)
	result[0] = msg
	copy(result[1:], args)

	return result
}
