// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/v2/codegen@master [sha: bb2c19fba6c03f46cb643b3bcdc1d647ea1453ab]; DO NOT EDIT.

package assert

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// Conditionf is the same as [Condition], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Conditionf(t T, comp Comparison, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Condition(t, comp, forwardArgs(msg, args))
}

// Containsf is the same as [Contains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Containsf(t T, s any, contains any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Contains(t, s, contains, forwardArgs(msg, args))
}

// DirExistsf is the same as [DirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func DirExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(t, path, forwardArgs(msg, args))
}

// ElementsMatchf is the same as [ElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(t, listA, listB, forwardArgs(msg, args))
}

// Emptyf is the same as [Empty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Emptyf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Empty(t, object, forwardArgs(msg, args))
}

// Equalf is the same as [Equal], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Equalf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Equal(t, expected, actual, forwardArgs(msg, args))
}

// EqualErrorf is the same as [EqualError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualErrorf(t T, theError error, errString string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(t, theError, errString, forwardArgs(msg, args))
}

// EqualExportedValuesf is the same as [EqualExportedValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualExportedValuesf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(t, expected, actual, forwardArgs(msg, args))
}

// EqualValuesf is the same as [EqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualValuesf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(t, expected, actual, forwardArgs(msg, args))
}

// Errorf is the same as [Error], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Errorf(t T, err error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Error(t, err, forwardArgs(msg, args))
}

// ErrorAsf is the same as [ErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorAsf(t T, err error, target any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(t, err, target, forwardArgs(msg, args))
}

// ErrorContainsf is the same as [ErrorContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorContainsf(t T, theError error, contains string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(t, theError, contains, forwardArgs(msg, args))
}

// ErrorIsf is the same as [ErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorIsf(t T, err error, target error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(t, err, target, forwardArgs(msg, args))
}

// Eventuallyf is the same as [Eventually], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Eventuallyf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Eventually(t, condition, waitFor, tick, forwardArgs(msg, args))
}

// EventuallyWithTf is the same as [EventuallyWithT], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EventuallyWithTf(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWithT(t, condition, waitFor, tick, forwardArgs(msg, args))
}

// Exactlyf is the same as [Exactly], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Exactlyf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(t, expected, actual, forwardArgs(msg, args))
}

// Failf is the same as [Fail], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Failf(t T, failureMessage string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Fail(t, failureMessage, forwardArgs(msg, args))
}

// FailNowf is the same as [FailNow], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FailNowf(t T, failureMessage string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(t, failureMessage, forwardArgs(msg, args))
}

// Falsef is the same as [False], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Falsef(t T, value bool, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.False(t, value, forwardArgs(msg, args))
}

// FileEmptyf is the same as [FileEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileEmptyf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(t, path, forwardArgs(msg, args))
}

// FileExistsf is the same as [FileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(t, path, forwardArgs(msg, args))
}

// FileNotEmptyf is the same as [FileNotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileNotEmptyf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(t, path, forwardArgs(msg, args))
}

// Greaterf is the same as [Greater], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Greaterf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Greater(t, e1, e2, forwardArgs(msg, args))
}

// GreaterOrEqualf is the same as [GreaterOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func GreaterOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(t, e1, e2, forwardArgs(msg, args))
}

// HTTPBodyContainsf is the same as [HTTPBodyContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPBodyContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(t, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPBodyNotContainsf is the same as [HTTPBodyNotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPBodyNotContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(t, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPErrorf is the same as [HTTPError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPErrorf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(t, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPRedirectf is the same as [HTTPRedirect], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPRedirectf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(t, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPStatusCodef is the same as [HTTPStatusCode], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPStatusCodef(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(t, handler, method, url, values, statuscode, forwardArgs(msg, args))
}

// HTTPSuccessf is the same as [HTTPSuccess], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPSuccessf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(t, handler, method, url, values, forwardArgs(msg, args))
}

// Implementsf is the same as [Implements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Implementsf(t T, interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Implements(t, interfaceObject, object, forwardArgs(msg, args))
}

// InDeltaf is the same as [InDelta], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaf(t T, expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(t, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaMapValuesf is the same as [InDeltaMapValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaMapValuesf(t T, expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(t, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaSlicef is the same as [InDeltaSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaSlicef(t T, expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(t, expected, actual, delta, forwardArgs(msg, args))
}

// InEpsilonf is the same as [InEpsilon], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InEpsilonf(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(t, expected, actual, epsilon, forwardArgs(msg, args))
}

// InEpsilonSlicef is the same as [InEpsilonSlice], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InEpsilonSlicef(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(t, expected, actual, epsilon, forwardArgs(msg, args))
}

// IsDecreasingf is the same as [IsDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsDecreasingf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(t, object, forwardArgs(msg, args))
}

// IsIncreasingf is the same as [IsIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsIncreasingf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(t, object, forwardArgs(msg, args))
}

// IsNonDecreasingf is the same as [IsNonDecreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonDecreasingf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(t, object, forwardArgs(msg, args))
}

// IsNonIncreasingf is the same as [IsNonIncreasing], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonIncreasingf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(t, object, forwardArgs(msg, args))
}

// IsNotTypef is the same as [IsNotType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNotTypef(t T, theType any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(t, theType, object, forwardArgs(msg, args))
}

// IsTypef is the same as [IsType], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsTypef(t T, expectedType any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsType(t, expectedType, object, forwardArgs(msg, args))
}

// JSONEqf is the same as [JSONEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONEqf(t T, expected string, actual string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(t, expected, actual, forwardArgs(msg, args))
}

// JSONEqBytesf is the same as [JSONEqBytes], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(t, expected, actual, forwardArgs(msg, args))
}

// Lenf is the same as [Len], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Lenf(t T, object any, length int, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Len(t, object, length, forwardArgs(msg, args))
}

// Lessf is the same as [Less], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Lessf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Less(t, e1, e2, forwardArgs(msg, args))
}

// LessOrEqualf is the same as [LessOrEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func LessOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(t, e1, e2, forwardArgs(msg, args))
}

// Negativef is the same as [Negative], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Negativef(t T, e any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Negative(t, e, forwardArgs(msg, args))
}

// Neverf is the same as [Never], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Neverf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Never(t, condition, waitFor, tick, forwardArgs(msg, args))
}

// Nilf is the same as [Nil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Nilf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Nil(t, object, forwardArgs(msg, args))
}

// NoDirExistsf is the same as [NoDirExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoDirExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoDirExists(t, path, forwardArgs(msg, args))
}

// NoErrorf is the same as [NoError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoErrorf(t T, err error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoError(t, err, forwardArgs(msg, args))
}

// NoFileExistsf is the same as [NoFileExists], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoFileExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoFileExists(t, path, forwardArgs(msg, args))
}

// NotContainsf is the same as [NotContains], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotContainsf(t T, s any, contains any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(t, s, contains, forwardArgs(msg, args))
}

// NotElementsMatchf is the same as [NotElementsMatch], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(t, listA, listB, forwardArgs(msg, args))
}

// NotEmptyf is the same as [NotEmpty], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEmptyf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(t, object, forwardArgs(msg, args))
}

// NotEqualf is the same as [NotEqual], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEqualf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(t, expected, actual, forwardArgs(msg, args))
}

// NotEqualValuesf is the same as [NotEqualValues], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(t, expected, actual, forwardArgs(msg, args))
}

// NotErrorAsf is the same as [NotErrorAs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotErrorAsf(t T, err error, target any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(t, err, target, forwardArgs(msg, args))
}

// NotErrorIsf is the same as [NotErrorIs], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotErrorIsf(t T, err error, target error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(t, err, target, forwardArgs(msg, args))
}

// NotImplementsf is the same as [NotImplements], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotImplementsf(t T, interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(t, interfaceObject, object, forwardArgs(msg, args))
}

// NotNilf is the same as [NotNil], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotNilf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(t, object, forwardArgs(msg, args))
}

// NotPanicsf is the same as [NotPanics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotPanicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(t, f, forwardArgs(msg, args))
}

// NotRegexpf is the same as [NotRegexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotRegexpf(t T, rx any, str any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(t, rx, str, forwardArgs(msg, args))
}

// NotSamef is the same as [NotSame], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSamef(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(t, expected, actual, forwardArgs(msg, args))
}

// NotSubsetf is the same as [NotSubset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSubsetf(t T, list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(t, list, subset, forwardArgs(msg, args))
}

// NotZerof is the same as [NotZero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotZerof(t T, i any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(t, i, forwardArgs(msg, args))
}

// Panicsf is the same as [Panics], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Panicsf(t T, f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Panics(t, f, forwardArgs(msg, args))
}

// PanicsWithErrorf is the same as [PanicsWithError], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func PanicsWithErrorf(t T, errString string, f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(t, errString, f, forwardArgs(msg, args))
}

// PanicsWithValuef is the same as [PanicsWithValue], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func PanicsWithValuef(t T, expected any, f assertions.PanicTestFunc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(t, expected, f, forwardArgs(msg, args))
}

// Positivef is the same as [Positive], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Positivef(t T, e any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Positive(t, e, forwardArgs(msg, args))
}

// Regexpf is the same as [Regexp], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Regexpf(t T, rx any, str any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(t, rx, str, forwardArgs(msg, args))
}

// Samef is the same as [Same], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Samef(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Same(t, expected, actual, forwardArgs(msg, args))
}

// Subsetf is the same as [Subset], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Subsetf(t T, list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Subset(t, list, subset, forwardArgs(msg, args))
}

// Truef is the same as [True], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Truef(t T, value bool, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.True(t, value, forwardArgs(msg, args))
}

// WithinDurationf is the same as [WithinDuration], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func WithinDurationf(t T, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(t, expected, actual, delta, forwardArgs(msg, args))
}

// WithinRangef is the same as [WithinRange], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func WithinRangef(t T, actual time.Time, start time.Time, end time.Time, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(t, actual, start, end, forwardArgs(msg, args))
}

// YAMLEqf is the same as [YAMLEq], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func YAMLEqf(t T, expected string, actual string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(t, expected, actual, forwardArgs(msg, args))
}

// Zerof is the same as [Zero], but accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Zerof(t T, i any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Zero(t, i, forwardArgs(msg, args))
}

func forwardArgs(msg string, args []any) []any {
	result := make([]any, len(args)+1)
	result[0] = msg
	copy(result[1:], args)

	return result
}
