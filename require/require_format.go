// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

package require

import (
	"iter"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// Conditionf is the same as [Condition], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Containsf is the same as [Contains], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// DirExistsf is the same as [DirExists], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// DirNotExistsf is the same as [DirNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func DirNotExistsf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.DirNotExists(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// ElementsMatchf is the same as [ElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// ElementsMatchTf is the same as [ElementsMatchT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatchT[E](t, listA, listB, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Emptyf is the same as [Empty], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Equalf is the same as [Equal], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// EqualErrorf is the same as [EqualError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualErrorf(t T, err error, errString string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualError(t, err, errString, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// EqualExportedValuesf is the same as [EqualExportedValues], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// EqualTf is the same as [EqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualT[V](t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// EqualValuesf is the same as [EqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Errorf is the same as [Error], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// ErrorAsf is the same as [ErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// ErrorContainsf is the same as [ErrorContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorContainsf(t T, err error, contains string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorContains(t, err, contains, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// ErrorIsf is the same as [ErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Eventuallyf is the same as [Eventually], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// EventuallyWithf is the same as [EventuallyWith], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EventuallyWithf(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWith(t, condition, waitFor, tick, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Exactlyf is the same as [Exactly], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Failf is the same as [Fail], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Failf(t T, failureMessage string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	_ = assertions.Fail(t, failureMessage, forwardArgs(msg, args))

	t.FailNow()
}

// FailNowf is the same as [FailNow], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FailNowf(t T, failureMessage string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	_ = assertions.FailNow(t, failureMessage, forwardArgs(msg, args))

	t.FailNow()
}

// Falsef is the same as [False], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// FalseTf is the same as [FalseT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FalseTf[B Boolean](t T, value B, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FalseT[B](t, value, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// FileEmptyf is the same as [FileEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// FileExistsf is the same as [FileExists], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// FileNotEmptyf is the same as [FileNotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// FileNotExistsf is the same as [FileNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FileNotExistsf(t T, path string, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileNotExists(t, path, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Greaterf is the same as [Greater], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// GreaterOrEqualf is the same as [GreaterOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// GreaterOrEqualTf is the same as [GreaterOrEqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func GreaterOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqualT[Orderable](t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// GreaterTf is the same as [GreaterT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func GreaterTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterT[Orderable](t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// HTTPBodyContainsf is the same as [HTTPBodyContains], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// HTTPBodyNotContainsf is the same as [HTTPBodyNotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// HTTPErrorf is the same as [HTTPError], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// HTTPRedirectf is the same as [HTTPRedirect], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// HTTPStatusCodef is the same as [HTTPStatusCode], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// HTTPSuccessf is the same as [HTTPSuccess], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Implementsf is the same as [Implements], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// InDeltaf is the same as [InDelta], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// InDeltaMapValuesf is the same as [InDeltaMapValues], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// InDeltaSlicef is the same as [InDeltaSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// InDeltaTf is the same as [InDeltaT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InDeltaTf[Number Measurable](t T, expected Number, actual Number, delta Number, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaT[Number](t, expected, actual, delta, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// InEpsilonf is the same as [InEpsilon], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// InEpsilonSlicef is the same as [InEpsilonSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// InEpsilonTf is the same as [InEpsilonT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InEpsilonTf[Number Measurable](t T, expected Number, actual Number, epsilon float64, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonT[Number](t, expected, actual, epsilon, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsDecreasingf is the same as [IsDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsDecreasingf(t T, collection any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasing(t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsDecreasingTf is the same as [IsDecreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsIncreasingf is the same as [IsIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsIncreasingf(t T, collection any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasing(t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsIncreasingTf is the same as [IsIncreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNonDecreasingf is the same as [IsNonDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonDecreasingf(t T, collection any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasing(t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNonDecreasingTf is the same as [IsNonDecreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNonIncreasingf is the same as [IsNonIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonIncreasingf(t T, collection any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasing(t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNonIncreasingTf is the same as [IsNonIncreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNotOfTypeTf is the same as [IsNotOfTypeT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNotOfTypeTf[EType any](t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNotOfTypeT[EType](t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsNotTypef is the same as [IsNotType], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// IsOfTypeTf is the same as [IsOfTypeT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsOfTypeTf[EType any](t T, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsOfTypeT[EType](t, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// IsTypef is the same as [IsType], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// JSONEqf is the same as [JSONEq], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// JSONEqBytesf is the same as [JSONEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// JSONEqTf is the same as [JSONEqT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqT[EDoc, ADoc](t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// JSONMarshalAsTf is the same as [JSONMarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONMarshalAsT[EDoc](t, expected, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// JSONUnmarshalAsTf is the same as [JSONUnmarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONUnmarshalAsTf[Object any, ADoc Text](t T, expected Object, jazon ADoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONUnmarshalAsT[Object, ADoc](t, expected, jazon, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Kindf is the same as [Kind], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Lenf is the same as [Len], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Lessf is the same as [Less], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// LessOrEqualf is the same as [LessOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// LessOrEqualTf is the same as [LessOrEqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func LessOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqualT[Orderable](t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// LessTf is the same as [LessT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func LessTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.LessT[Orderable](t, e1, e2, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// MapContainsTf is the same as [MapContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func MapContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.MapContainsT[Map, K, V](t, m, key, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// MapNotContainsTf is the same as [MapNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func MapNotContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.MapNotContainsT[Map, K, V](t, m, key, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Negativef is the same as [Negative], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NegativeTf is the same as [NegativeT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NegativeTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NegativeT[SignedNumber](t, e, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Neverf is the same as [Never], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Nilf is the same as [Nil], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NoErrorf is the same as [NoError], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NoFileDescriptorLeakf is the same as [NoFileDescriptorLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NoFileDescriptorLeakf(t T, tested func(), msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoFileDescriptorLeak(t, tested, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NoGoRoutineLeakf is the same as [NoGoRoutineLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NoGoRoutineLeakf(t T, tested func(), msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoGoRoutineLeak(t, tested, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotContainsf is the same as [NotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotElementsMatchf is the same as [NotElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotElementsMatchTf is the same as [NotElementsMatchT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatchT[E](t, listA, listB, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotEmptyf is the same as [NotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotEqualf is the same as [NotEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotEqualTf is the same as [NotEqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualT[V](t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotEqualValuesf is the same as [NotEqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotErrorAsf is the same as [NotErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotErrorIsf is the same as [NotErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotImplementsf is the same as [NotImplements], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotKindf is the same as [NotKind], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotNilf is the same as [NotNil], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotPanicsf is the same as [NotPanics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotPanicsf(t T, f func(), msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotPanics(t, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotRegexpf is the same as [NotRegexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotRegexpf(t T, rx any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexp(t, rx, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotRegexpTf is the same as [NotRegexpT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotRegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexpT[Rex, ADoc](t, rx, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotSamef is the same as [NotSame], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotSameTf is the same as [NotSameT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSameTf[P any](t T, expected *P, actual *P, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSameT[P](t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotSortedTf is the same as [NotSortedT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSortedT[OrderedSlice, E](t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// NotSubsetf is the same as [NotSubset], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// NotZerof is the same as [NotZero], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Panicsf is the same as [Panics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Panicsf(t T, f func(), msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Panics(t, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// PanicsWithErrorf is the same as [PanicsWithError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func PanicsWithErrorf(t T, errString string, f func(), msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithError(t, errString, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// PanicsWithValuef is the same as [PanicsWithValue], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func PanicsWithValuef(t T, expected any, f func(), msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithValue(t, expected, f, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Positivef is the same as [Positive], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// PositiveTf is the same as [PositiveT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func PositiveTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PositiveT[SignedNumber](t, e, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Regexpf is the same as [Regexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Regexpf(t T, rx any, actual any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Regexp(t, rx, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// RegexpTf is the same as [RegexpT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func RegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.RegexpT[Rex, ADoc](t, rx, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Samef is the same as [Same], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// SameTf is the same as [SameT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SameTf[P any](t T, expected *P, actual *P, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SameT[P](t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// SeqContainsTf is the same as [SeqContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SeqContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SeqContainsT[E](t, iter, element, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// SeqNotContainsTf is the same as [SeqNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SeqNotContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SeqNotContainsT[E](t, iter, element, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// SliceContainsTf is the same as [SliceContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SliceContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceContainsT[Slice, E](t, s, element, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// SliceNotContainsTf is the same as [SliceNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SliceNotContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotContainsT[Slice, E](t, s, element, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// SliceNotSubsetTf is the same as [SliceNotSubsetT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SliceNotSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotSubsetT[Slice, E](t, list, subset, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// SliceSubsetTf is the same as [SliceSubsetT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SliceSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceSubsetT[Slice, E](t, list, subset, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// SortedTf is the same as [SortedT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SortedT[OrderedSlice, E](t, collection, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// StringContainsTf is the same as [StringContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func StringContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.StringContainsT[ADoc, EDoc](t, str, substring, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// StringNotContainsTf is the same as [StringNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func StringNotContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.StringNotContainsT[ADoc, EDoc](t, str, substring, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Subsetf is the same as [Subset], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// Truef is the same as [True], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// TrueTf is the same as [TrueT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func TrueTf[B Boolean](t T, value B, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.TrueT[B](t, value, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// WithinDurationf is the same as [WithinDuration], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// WithinRangef is the same as [WithinRange], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// YAMLEqf is the same as [YAMLEq], but it accepts a format msg string to format arguments like [fmt.Printf].
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

// YAMLEqBytesf is the same as [YAMLEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqBytes(t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// YAMLEqTf is the same as [YAMLEqT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqT[EDoc, ADoc](t, expected, actual, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// YAMLMarshalAsTf is the same as [YAMLMarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLMarshalAsT[EDoc](t, expected, object, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// YAMLUnmarshalAsTf is the same as [YAMLUnmarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLUnmarshalAsTf[Object any, ADoc Text](t T, expected Object, jazon ADoc, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLUnmarshalAsT[Object, ADoc](t, expected, jazon, forwardArgs(msg, args)) {
		return
	}

	t.FailNow()
}

// Zerof is the same as [Zero], but it accepts a format msg string to format arguments like [fmt.Printf].
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
