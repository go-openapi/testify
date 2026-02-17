// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

package assert

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
// Upon failure, the test [T] is marked as failed and continues execution.
func Conditionf(t T, comp Comparison, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Condition(t, comp, forwardArgs(msg, args))
}

// Containsf is the same as [Contains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Containsf(t T, s any, contains any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Contains(t, s, contains, forwardArgs(msg, args))
}

// DirExistsf is the same as [DirExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func DirExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(t, path, forwardArgs(msg, args))
}

// DirNotExistsf is the same as [DirNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func DirNotExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.DirNotExists(t, path, forwardArgs(msg, args))
}

// ElementsMatchf is the same as [ElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(t, listA, listB, forwardArgs(msg, args))
}

// ElementsMatchTf is the same as [ElementsMatchT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatchT[E](t, listA, listB, forwardArgs(msg, args))
}

// Emptyf is the same as [Empty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Emptyf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Empty(t, object, forwardArgs(msg, args))
}

// Equalf is the same as [Equal], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Equalf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Equal(t, expected, actual, forwardArgs(msg, args))
}

// EqualErrorf is the same as [EqualError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualErrorf(t T, err error, errString string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(t, err, errString, forwardArgs(msg, args))
}

// EqualExportedValuesf is the same as [EqualExportedValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualExportedValuesf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(t, expected, actual, forwardArgs(msg, args))
}

// EqualTf is the same as [EqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualT[V](t, expected, actual, forwardArgs(msg, args))
}

// EqualValuesf is the same as [EqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualValuesf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(t, expected, actual, forwardArgs(msg, args))
}

// Errorf is the same as [Error], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Errorf(t T, err error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Error(t, err, forwardArgs(msg, args))
}

// ErrorAsf is the same as [ErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorAsf(t T, err error, target any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(t, err, target, forwardArgs(msg, args))
}

// ErrorContainsf is the same as [ErrorContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorContainsf(t T, err error, contains string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(t, err, contains, forwardArgs(msg, args))
}

// ErrorIsf is the same as [ErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorIsf(t T, err error, target error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(t, err, target, forwardArgs(msg, args))
}

// Eventuallyf is the same as [Eventually], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Eventuallyf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Eventually(t, condition, waitFor, tick, forwardArgs(msg, args))
}

// EventuallyWithf is the same as [EventuallyWith], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EventuallyWithf(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWith(t, condition, waitFor, tick, forwardArgs(msg, args))
}

// Exactlyf is the same as [Exactly], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Exactlyf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(t, expected, actual, forwardArgs(msg, args))
}

// Failf is the same as [Fail], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Failf(t T, failureMessage string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Fail(t, failureMessage, forwardArgs(msg, args))
}

// FailNowf is the same as [FailNow], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FailNowf(t T, failureMessage string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(t, failureMessage, forwardArgs(msg, args))
}

// Falsef is the same as [False], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Falsef(t T, value bool, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.False(t, value, forwardArgs(msg, args))
}

// FalseTf is the same as [FalseT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FalseTf[B Boolean](t T, value B, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FalseT[B](t, value, forwardArgs(msg, args))
}

// FileEmptyf is the same as [FileEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileEmptyf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(t, path, forwardArgs(msg, args))
}

// FileExistsf is the same as [FileExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(t, path, forwardArgs(msg, args))
}

// FileNotEmptyf is the same as [FileNotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileNotEmptyf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(t, path, forwardArgs(msg, args))
}

// FileNotExistsf is the same as [FileNotExists], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileNotExistsf(t T, path string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileNotExists(t, path, forwardArgs(msg, args))
}

// Greaterf is the same as [Greater], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Greaterf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Greater(t, e1, e2, forwardArgs(msg, args))
}

// GreaterOrEqualf is the same as [GreaterOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func GreaterOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(t, e1, e2, forwardArgs(msg, args))
}

// GreaterOrEqualTf is the same as [GreaterOrEqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func GreaterOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqualT[Orderable](t, e1, e2, forwardArgs(msg, args))
}

// GreaterTf is the same as [GreaterT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func GreaterTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.GreaterT[Orderable](t, e1, e2, forwardArgs(msg, args))
}

// HTTPBodyContainsf is the same as [HTTPBodyContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPBodyContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(t, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPBodyNotContainsf is the same as [HTTPBodyNotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPBodyNotContainsf(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(t, handler, method, url, values, str, forwardArgs(msg, args))
}

// HTTPErrorf is the same as [HTTPError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPErrorf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(t, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPRedirectf is the same as [HTTPRedirect], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPRedirectf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(t, handler, method, url, values, forwardArgs(msg, args))
}

// HTTPStatusCodef is the same as [HTTPStatusCode], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPStatusCodef(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(t, handler, method, url, values, statuscode, forwardArgs(msg, args))
}

// HTTPSuccessf is the same as [HTTPSuccess], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPSuccessf(t T, handler http.HandlerFunc, method string, url string, values url.Values, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(t, handler, method, url, values, forwardArgs(msg, args))
}

// Implementsf is the same as [Implements], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Implementsf(t T, interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Implements(t, interfaceObject, object, forwardArgs(msg, args))
}

// InDeltaf is the same as [InDelta], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaf(t T, expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(t, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaMapValuesf is the same as [InDeltaMapValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaMapValuesf(t T, expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(t, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaSlicef is the same as [InDeltaSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaSlicef(t T, expected any, actual any, delta float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(t, expected, actual, delta, forwardArgs(msg, args))
}

// InDeltaTf is the same as [InDeltaT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaTf[Number Measurable](t T, expected Number, actual Number, delta Number, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaT[Number](t, expected, actual, delta, forwardArgs(msg, args))
}

// InEpsilonf is the same as [InEpsilon], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InEpsilonf(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(t, expected, actual, epsilon, forwardArgs(msg, args))
}

// InEpsilonSlicef is the same as [InEpsilonSlice], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InEpsilonSlicef(t T, expected any, actual any, epsilon float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(t, expected, actual, epsilon, forwardArgs(msg, args))
}

// InEpsilonTf is the same as [InEpsilonT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InEpsilonTf[Number Measurable](t T, expected Number, actual Number, epsilon float64, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonT[Number](t, expected, actual, epsilon, forwardArgs(msg, args))
}

// IsDecreasingf is the same as [IsDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsDecreasingf(t T, collection any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(t, collection, forwardArgs(msg, args))
}

// IsDecreasingTf is the same as [IsDecreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args))
}

// IsIncreasingf is the same as [IsIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsIncreasingf(t T, collection any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(t, collection, forwardArgs(msg, args))
}

// IsIncreasingTf is the same as [IsIncreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args))
}

// IsNonDecreasingf is the same as [IsNonDecreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonDecreasingf(t T, collection any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(t, collection, forwardArgs(msg, args))
}

// IsNonDecreasingTf is the same as [IsNonDecreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonDecreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args))
}

// IsNonIncreasingf is the same as [IsNonIncreasing], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonIncreasingf(t T, collection any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(t, collection, forwardArgs(msg, args))
}

// IsNonIncreasingTf is the same as [IsNonIncreasingT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonIncreasingTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasingT[OrderedSlice, E](t, collection, forwardArgs(msg, args))
}

// IsNotOfTypeTf is the same as [IsNotOfTypeT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNotOfTypeTf[EType any](t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNotOfTypeT[EType](t, object, forwardArgs(msg, args))
}

// IsNotTypef is the same as [IsNotType], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNotTypef(t T, theType any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(t, theType, object, forwardArgs(msg, args))
}

// IsOfTypeTf is the same as [IsOfTypeT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsOfTypeTf[EType any](t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsOfTypeT[EType](t, object, forwardArgs(msg, args))
}

// IsTypef is the same as [IsType], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsTypef(t T, expectedType any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsType(t, expectedType, object, forwardArgs(msg, args))
}

// JSONEqf is the same as [JSONEq], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONEqf(t T, expected string, actual string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(t, expected, actual, forwardArgs(msg, args))
}

// JSONEqBytesf is the same as [JSONEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(t, expected, actual, forwardArgs(msg, args))
}

// JSONEqTf is the same as [JSONEqT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqT[EDoc, ADoc](t, expected, actual, forwardArgs(msg, args))
}

// JSONMarshalAsTf is the same as [JSONMarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONMarshalAsT[EDoc](t, expected, object, forwardArgs(msg, args))
}

// JSONUnmarshalAsTf is the same as [JSONUnmarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONUnmarshalAsTf[Object any, ADoc Text](t T, expected Object, jazon ADoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONUnmarshalAsT[Object, ADoc](t, expected, jazon, forwardArgs(msg, args))
}

// Kindf is the same as [Kind], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Kindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Kind(t, expectedKind, object, forwardArgs(msg, args))
}

// Lenf is the same as [Len], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Lenf(t T, object any, length int, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Len(t, object, length, forwardArgs(msg, args))
}

// Lessf is the same as [Less], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Lessf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Less(t, e1, e2, forwardArgs(msg, args))
}

// LessOrEqualf is the same as [LessOrEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func LessOrEqualf(t T, e1 any, e2 any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(t, e1, e2, forwardArgs(msg, args))
}

// LessOrEqualTf is the same as [LessOrEqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func LessOrEqualTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqualT[Orderable](t, e1, e2, forwardArgs(msg, args))
}

// LessTf is the same as [LessT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func LessTf[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.LessT[Orderable](t, e1, e2, forwardArgs(msg, args))
}

// MapContainsTf is the same as [MapContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func MapContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.MapContainsT[Map, K, V](t, m, key, forwardArgs(msg, args))
}

// MapNotContainsTf is the same as [MapNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func MapNotContainsTf[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.MapNotContainsT[Map, K, V](t, m, key, forwardArgs(msg, args))
}

// Negativef is the same as [Negative], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Negativef(t T, e any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Negative(t, e, forwardArgs(msg, args))
}

// NegativeTf is the same as [NegativeT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NegativeTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NegativeT[SignedNumber](t, e, forwardArgs(msg, args))
}

// Neverf is the same as [Never], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Neverf(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Never(t, condition, waitFor, tick, forwardArgs(msg, args))
}

// Nilf is the same as [Nil], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Nilf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Nil(t, object, forwardArgs(msg, args))
}

// NoErrorf is the same as [NoError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoErrorf(t T, err error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoError(t, err, forwardArgs(msg, args))
}

// NoFileDescriptorLeakf is the same as [NoFileDescriptorLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoFileDescriptorLeakf(t T, tested func(), msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoFileDescriptorLeak(t, tested, forwardArgs(msg, args))
}

// NoGoRoutineLeakf is the same as [NoGoRoutineLeak], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoGoRoutineLeakf(t T, tested func(), msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoGoRoutineLeak(t, tested, forwardArgs(msg, args))
}

// NotContainsf is the same as [NotContains], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotContainsf(t T, s any, contains any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(t, s, contains, forwardArgs(msg, args))
}

// NotElementsMatchf is the same as [NotElementsMatch], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotElementsMatchf(t T, listA any, listB any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(t, listA, listB, forwardArgs(msg, args))
}

// NotElementsMatchTf is the same as [NotElementsMatchT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotElementsMatchTf[E comparable](t T, listA []E, listB []E, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatchT[E](t, listA, listB, forwardArgs(msg, args))
}

// NotEmptyf is the same as [NotEmpty], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEmptyf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(t, object, forwardArgs(msg, args))
}

// NotEqualf is the same as [NotEqual], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEqualf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(t, expected, actual, forwardArgs(msg, args))
}

// NotEqualTf is the same as [NotEqualT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEqualTf[V comparable](t T, expected V, actual V, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualT[V](t, expected, actual, forwardArgs(msg, args))
}

// NotEqualValuesf is the same as [NotEqualValues], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEqualValuesf(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(t, expected, actual, forwardArgs(msg, args))
}

// NotErrorAsf is the same as [NotErrorAs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotErrorAsf(t T, err error, target any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(t, err, target, forwardArgs(msg, args))
}

// NotErrorIsf is the same as [NotErrorIs], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotErrorIsf(t T, err error, target error, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(t, err, target, forwardArgs(msg, args))
}

// NotImplementsf is the same as [NotImplements], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotImplementsf(t T, interfaceObject any, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(t, interfaceObject, object, forwardArgs(msg, args))
}

// NotKindf is the same as [NotKind], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotKindf(t T, expectedKind reflect.Kind, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotKind(t, expectedKind, object, forwardArgs(msg, args))
}

// NotNilf is the same as [NotNil], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotNilf(t T, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(t, object, forwardArgs(msg, args))
}

// NotPanicsf is the same as [NotPanics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotPanicsf(t T, f func(), msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(t, f, forwardArgs(msg, args))
}

// NotRegexpf is the same as [NotRegexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotRegexpf(t T, rx any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(t, rx, actual, forwardArgs(msg, args))
}

// NotRegexpTf is the same as [NotRegexpT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotRegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexpT[Rex, ADoc](t, rx, actual, forwardArgs(msg, args))
}

// NotSamef is the same as [NotSame], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSamef(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(t, expected, actual, forwardArgs(msg, args))
}

// NotSameTf is the same as [NotSameT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSameTf[P any](t T, expected *P, actual *P, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSameT[P](t, expected, actual, forwardArgs(msg, args))
}

// NotSortedTf is the same as [NotSortedT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSortedT[OrderedSlice, E](t, collection, forwardArgs(msg, args))
}

// NotSubsetf is the same as [NotSubset], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSubsetf(t T, list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(t, list, subset, forwardArgs(msg, args))
}

// NotZerof is the same as [NotZero], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotZerof(t T, i any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(t, i, forwardArgs(msg, args))
}

// Panicsf is the same as [Panics], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Panicsf(t T, f func(), msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Panics(t, f, forwardArgs(msg, args))
}

// PanicsWithErrorf is the same as [PanicsWithError], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func PanicsWithErrorf(t T, errString string, f func(), msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(t, errString, f, forwardArgs(msg, args))
}

// PanicsWithValuef is the same as [PanicsWithValue], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func PanicsWithValuef(t T, expected any, f func(), msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(t, expected, f, forwardArgs(msg, args))
}

// Positivef is the same as [Positive], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Positivef(t T, e any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Positive(t, e, forwardArgs(msg, args))
}

// PositiveTf is the same as [PositiveT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func PositiveTf[SignedNumber SignedNumeric](t T, e SignedNumber, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.PositiveT[SignedNumber](t, e, forwardArgs(msg, args))
}

// Regexpf is the same as [Regexp], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Regexpf(t T, rx any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(t, rx, actual, forwardArgs(msg, args))
}

// RegexpTf is the same as [RegexpT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func RegexpTf[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.RegexpT[Rex, ADoc](t, rx, actual, forwardArgs(msg, args))
}

// Samef is the same as [Same], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Samef(t T, expected any, actual any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Same(t, expected, actual, forwardArgs(msg, args))
}

// SameTf is the same as [SameT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SameTf[P any](t T, expected *P, actual *P, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SameT[P](t, expected, actual, forwardArgs(msg, args))
}

// SeqContainsTf is the same as [SeqContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SeqContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SeqContainsT[E](t, iter, element, forwardArgs(msg, args))
}

// SeqNotContainsTf is the same as [SeqNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SeqNotContainsTf[E comparable](t T, iter iter.Seq[E], element E, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SeqNotContainsT[E](t, iter, element, forwardArgs(msg, args))
}

// SliceContainsTf is the same as [SliceContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SliceContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SliceContainsT[Slice, E](t, s, element, forwardArgs(msg, args))
}

// SliceNotContainsTf is the same as [SliceNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SliceNotContainsTf[Slice ~[]E, E comparable](t T, s Slice, element E, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotContainsT[Slice, E](t, s, element, forwardArgs(msg, args))
}

// SliceNotSubsetTf is the same as [SliceNotSubsetT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SliceNotSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SliceNotSubsetT[Slice, E](t, list, subset, forwardArgs(msg, args))
}

// SliceSubsetTf is the same as [SliceSubsetT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SliceSubsetTf[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SliceSubsetT[Slice, E](t, list, subset, forwardArgs(msg, args))
}

// SortedTf is the same as [SortedT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func SortedTf[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.SortedT[OrderedSlice, E](t, collection, forwardArgs(msg, args))
}

// StringContainsTf is the same as [StringContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func StringContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.StringContainsT[ADoc, EDoc](t, str, substring, forwardArgs(msg, args))
}

// StringNotContainsTf is the same as [StringNotContainsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func StringNotContainsTf[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.StringNotContainsT[ADoc, EDoc](t, str, substring, forwardArgs(msg, args))
}

// Subsetf is the same as [Subset], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Subsetf(t T, list any, subset any, msg string, args ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Subset(t, list, subset, forwardArgs(msg, args))
}

// Truef is the same as [True], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Truef(t T, value bool, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.True(t, value, forwardArgs(msg, args))
}

// TrueTf is the same as [TrueT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func TrueTf[B Boolean](t T, value B, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.TrueT[B](t, value, forwardArgs(msg, args))
}

// WithinDurationf is the same as [WithinDuration], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func WithinDurationf(t T, expected time.Time, actual time.Time, delta time.Duration, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(t, expected, actual, delta, forwardArgs(msg, args))
}

// WithinRangef is the same as [WithinRange], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func WithinRangef(t T, actual time.Time, start time.Time, end time.Time, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(t, actual, start, end, forwardArgs(msg, args))
}

// YAMLEqf is the same as [YAMLEq], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func YAMLEqf(t T, expected string, actual string, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(t, expected, actual, forwardArgs(msg, args))
}

// YAMLEqBytesf is the same as [YAMLEqBytes], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func YAMLEqBytesf(t T, expected []byte, actual []byte, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqBytes(t, expected, actual, forwardArgs(msg, args))
}

// YAMLEqTf is the same as [YAMLEqT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func YAMLEqTf[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEqT[EDoc, ADoc](t, expected, actual, forwardArgs(msg, args))
}

// YAMLMarshalAsTf is the same as [YAMLMarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func YAMLMarshalAsTf[EDoc Text](t T, expected EDoc, object any, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLMarshalAsT[EDoc](t, expected, object, forwardArgs(msg, args))
}

// YAMLUnmarshalAsTf is the same as [YAMLUnmarshalAsT], but it accepts a format msg string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func YAMLUnmarshalAsTf[Object any, ADoc Text](t T, expected Object, jazon ADoc, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLUnmarshalAsT[Object, ADoc](t, expected, jazon, forwardArgs(msg, args))
}

// Zerof is the same as [Zero], but it accepts a format msg string to format arguments like [fmt.Printf].
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
