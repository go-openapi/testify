// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/v2/codegen; DO NOT EDIT.
// Generated on 2026-01-02 (version v1.2.2-760-g97c29e3) using codegen version master [sha: 97c29e3dbfc40800a080863ceea81db0cfd6e858]

package assert

import (
	"net/http"
	"net/url"
	"time"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

// Condition uses a Comparison to assert a complex condition.
//
// # Usage
//
//	assertions.Condition(t, func() bool { return myCondition })
//
// # Examples
//
//	success:  func() bool { return true }
//	failure:  func() bool { return false }
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Condition(t T, comp Comparison, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Condition(t, comp, msgAndArgs...)
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
// # Usage
//
//	assertions.Contains(t, "Hello World", "World")
//	assertions.Contains(t, []string{"Hello", "World"}, "World")
//	assertions.Contains(t, map[string]string{"Hello": "World"}, "Hello")
//
// # Examples
//
//	success: []string{"A","B"}, "A"
//	failure: []string{"A","B"}, "C"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Contains(t T, s any, contains any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Contains(t, s, contains, msgAndArgs...)
}

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
//
// # Usage
//
//	assertions.DirExists(t, "path/to/directory")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"existing_dir")
//	failure: filepath.Join(testDataPath(),"non_existing_dir")
//
// Upon failure, the test [T] is marked as failed and continues execution.
func DirExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.DirExists(t, path, msgAndArgs...)
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// # Usage
//
//	assertions.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2])
//
// # Examples
//
//	success: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
//	failure: []int{1, 2, 3}, []int{1, 2, 4}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ElementsMatch(t, listA, listB, msgAndArgs...)
}

// Empty asserts that the given value is "empty".
//
// Zero values are "empty".
//
// Arrays are "empty" if every element is the zero value of the type (stricter than "empty").
//
// Slices, maps and channels with zero length are "empty".
//
// Pointer values are "empty" if the pointer is nil or if the pointed value is "empty".
//
// # Usage
//
//	assertions.Empty(t, obj)
//
// # Examples
//
//	success: ""
//	failure: "not empty"
//
// Upon failure, the test [T] is marked as failed and continues execution.
//
// [Zero values]: https://go.dev/ref/spec#The_zero_value
func Empty(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Empty(t, object, msgAndArgs...)
}

// Equal asserts that two objects are equal.
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
//
// Function equality cannot be determined and will always fail.
//
// # Usage
//
//	assertions.Equal(t, 123, 123)
//
// # Examples
//
//	success: 123, 123
//	failure: 123, 456
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Equal(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Equal(t, expected, actual, msgAndArgs...)
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualError(t T, theError error, errString string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualError(t, theError, errString, msgAndArgs...)
}

// EqualExportedValues asserts that the types of two objects are equal and their public
// fields are also equal. This is useful for comparing structs that have private fields
// that could potentially differ.
//
// # Usage
//
//	 type S struct {
//		Exported     	int
//		notExported   	int
//	 }
//	assertions.EqualExportedValues(t, S{1, 2}, S{1, 3}) => true
//	assertions.EqualExportedValues(t, S{1, 2}, S{2, 3}) => false
//
// # Examples
//
//	success: &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}
//	failure:  &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualExportedValues(t, expected, actual, msgAndArgs...)
}

// EqualValues asserts that two objects are equal or convertible to the larger
// type and equal.
//
// # Usage
//
//	assertions.EqualValues(t, uint32(123), int32(123))
//
// # Examples
//
//	success: uint32(123), int32(123)
//	failure: uint32(123), int32(456)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EqualValues(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EqualValues(t, expected, actual, msgAndArgs...)
}

// Error asserts that a function returned a non-nil error (ie. an error).
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Error(t T, err error, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Error(t, err, msgAndArgs...)
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorAs(t T, err error, target any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAs(t, err, target, msgAndArgs...)
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorContains(t T, theError error, contains string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorContains(t, theError, contains, msgAndArgs...)
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorIs(t T, err error, target error, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorIs(t, err, target, msgAndArgs...)
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
// # Usage
//
//	assertions.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
//
// # Examples
//
//	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Eventually(t, condition, waitFor, tick, msgAndArgs...)
}

// EventuallyWithT asserts that given condition will be met in waitFor time,
// periodically checking target function each tick. In contrast to Eventually,
// it supplies a CollectT to the condition function, so that the condition
// function can use the CollectT to call other assertions.
// The condition is considered "met" if no errors are raised in a tick.
// The supplied CollectT collects all errors from one tick (if there are any).
// If the condition is not met before waitFor, the collected errors of
// the last tick are copied to t.
//
// # Usage
//
//	externalValue := false
//	go func() {
//		time.Sleep(8*time.Second)
//		externalValue = true
//	}()
//	assertions.EventuallyWithT(t, func(c *assertions.CollectT) {
//		// add assertions as needed; any assertion failure will fail the current tick
//		assertions.True(c, externalValue, "expected 'externalValue' to be true")
//	}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
//
// # Examples
//
//	success: func(c *CollectT) { True(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//	failure: func(c *CollectT) { False(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//
// Upon failure, the test [T] is marked as failed and continues execution.
func EventuallyWithT(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.EventuallyWithT(t, condition, waitFor, tick, msgAndArgs...)
}

// Exactly asserts that two objects are equal in value and type.
//
// # Usage
//
//	assertions.Exactly(t, int32(123), int64(123))
//
// # Examples
//
//	success: int32(123), int32(123)
//	failure: int32(123), int64(123)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Exactly(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Exactly(t, expected, actual, msgAndArgs...)
}

// Fail reports a failure through.
//
// # Usage
//
//	assertions.Fail(t, "failed")
//
// # Examples
//
//	failure: "failed"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Fail(t T, failureMessage string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Fail(t, failureMessage, msgAndArgs...)
}

// FailNow fails test.
//
// # Usage
//
//	assertions.FailNow(t, "failed")
//
// # Examples
//
//	failure: "failed"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FailNow(t T, failureMessage string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FailNow(t, failureMessage, msgAndArgs...)
}

// False asserts that the specified value is false.
//
// # Usage
//
//	assertions.False(t, myBool)
//
// # Examples
//
//	success: 1 == 0
//	failure: 1 == 1
//
// Upon failure, the test [T] is marked as failed and continues execution.
func False(t T, value bool, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.False(t, value, msgAndArgs...)
}

// FileEmpty checks whether a file exists in the given path and is empty.
// It fails if the file is not empty, if the path points to a directory or there is an error when trying to check the file.
//
// # Usage
//
//	assertions.FileEmpty(t, "path/to/file")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"empty_file")
//	failure: filepath.Join(testDataPath(),"existing_file")
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileEmpty(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileEmpty(t, path, msgAndArgs...)
}

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
//
// # Usage
//
//	assertions.FileExists(t, "path/to/file")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"existing_file")
//	failure: filepath.Join(testDataPath(),"non_existing_file")
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileExists(t, path, msgAndArgs...)
}

// FileNotEmpty checks whether a file exists in the given path and is not empty.
// It fails if the file is empty, if the path points to a directory or there is an error when trying to check the file.
//
// # Usage
//
//	assertions.FileNotEmpty(t, "path/to/file")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"existing_file")
//	failure: filepath.Join(testDataPath(),"empty_file")
//
// Upon failure, the test [T] is marked as failed and continues execution.
func FileNotEmpty(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.FileNotEmpty(t, path, msgAndArgs...)
}

// Greater asserts that the first element is strictly greater than the second.
//
// # Usage
//
//	assertions.Greater(t, 2, 1)
//	assertions.Greater(t, float64(2), float64(1))
//	assertions.Greater(t, "b", "a")
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Greater(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Greater(t, e1, e2, msgAndArgs...)
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second.
//
// # Usage
//
//	assertions.GreaterOrEqual(t, 2, 1)
//	assertions.GreaterOrEqual(t, 2, 2)
//	assertions.GreaterOrEqual(t, "b", "a")
//	assertions.GreaterOrEqual(t, "b", "b")
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
//
// Upon failure, the test [T] is marked as failed and continues execution.
func GreaterOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.GreaterOrEqual(t, e1, e2, msgAndArgs...)
}

// HTTPBodyContains asserts that a specified handler returns a body that contains a string.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPBodyContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// # Examples
//
//	success: httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!"
//	failure: httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPBodyContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyContains(t, handler, method, url, values, str, msgAndArgs...)
}

// HTTPBodyNotContains asserts that a specified handler returns a
// body that does not contain a string.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPBodyNotContains(t, myHandler, "GET", "www.google.com", nil, "I'm Feeling Lucky")
//
// # Examples
//
//	success: httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!"
//	failure: httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPBodyNotContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPBodyNotContains(t, handler, method, url, values, str, msgAndArgs...)
}

// HTTPError asserts that a specified handler returns an error status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPError(t, myHandler, "POST", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// # Examples
//
//	success: httpError, "GET", "/", nil
//	failure: httpOK, "GET", "/", nil
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPError(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPError(t, handler, method, url, values, msgAndArgs...)
}

// HTTPRedirect asserts that a specified handler returns a redirect status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPRedirect(t, myHandler, "GET", "/a/b/c", url.Values{"a": []string{"b", "c"}}
//
// # Examples
//
//	success: httpRedirect, "GET", "/", nil
//	failure: httpError, "GET", "/", nil
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPRedirect(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPRedirect(t, handler, method, url, values, msgAndArgs...)
}

// HTTPStatusCode asserts that a specified handler returns a specified status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPStatusCode(t, myHandler, "GET", "/notImplemented", nil, 501)
//
// # Examples
//
//	success: httpOK, "GET", "/", nil, http.StatusOK
//	failure: httpError, "GET", "/", nil, http.StatusOK
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPStatusCode(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPStatusCode(t, handler, method, url, values, statuscode, msgAndArgs...)
}

// HTTPSuccess asserts that a specified handler returns a success status code.
//
// Returns whether the assertion was successful (true) or not (false).
//
// # Usage
//
//	assertions.HTTPSuccess(t, myHandler, "POST", "http://www.google.com", nil)
//
// # Examples
//
//	success: httpOK, "GET", "/", nil
//	failure: httpError, "GET", "/", nil
//
// Upon failure, the test [T] is marked as failed and continues execution.
func HTTPSuccess(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.HTTPSuccess(t, handler, method, url, values, msgAndArgs...)
}

// Implements asserts that an object is implemented by the specified interface.
//
// # Usage
//
//	assertions.Implements(t, (*MyInterface)(nil), new(MyObject))
//
// # Examples
//
//	success: ptr(dummyInterface), new(testing.T)
//	failure: (*error)(nil), new(testing.T)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Implements(t T, interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Implements(t, interfaceObject, object, msgAndArgs...)
}

// InDelta asserts that the two numerals are within delta of each other.
//
// # Usage
//
// assertions.InDelta(t, math.Pi, 22/7.0, 0.01)
//
// # Examples
//
//	success: 1.0, 1.01, 0.02
//	failure: 1.0, 1.1, 0.05
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDelta(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDelta(t, expected, actual, delta, msgAndArgs...)
}

// InDeltaMapValues is the same as InDelta, but it compares all values between two maps. Both maps must have exactly the same keys.
//
// # Usage
//
//	assertions.InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
//
// # Examples
//
//	success: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02
//	failure: map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaMapValues(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...)
}

// InDeltaSlice is the same as InDelta, except it compares two slices.
//
// # Usage
//
//	assertions.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
//
// # Examples
//
//	success: []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02
//	failure: []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InDeltaSlice(t T, expected any, actual any, delta float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InDeltaSlice(t, expected, actual, delta, msgAndArgs...)
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon.
//
// # Usage
//
//	assertions.InEpsilon(t, 100.0, 101.0, 0.02)
//
// # Examples
//
//	success: 100.0, 101.0, 0.02
//	failure: 100.0, 110.0, 0.05
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InEpsilon(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilon(t, expected, actual, epsilon, msgAndArgs...)
}

// InEpsilonSlice is the same as InEpsilon, except it compares each value from two slices.
//
// # Usage
//
//	assertions.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
//
// # Examples
//
//	success: []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02
//	failure: []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05
//
// Upon failure, the test [T] is marked as failed and continues execution.
func InEpsilonSlice(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.InEpsilonSlice(t, expected, actual, epsilon, msgAndArgs...)
}

// IsDecreasing asserts that the collection is decreasing.
//
// # Usage
//
//	assertions.IsDecreasing(t, []int{2, 1, 0})
//	assertions.IsDecreasing(t, []float{2, 1})
//	assertions.IsDecreasing(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{3, 2, 1}
//	failure: []int{1, 2, 3}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsDecreasing(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsDecreasing(t, object, msgAndArgs...)
}

// IsIncreasing asserts that the collection is increasing.
//
// # Usage
//
//	assertions.IsIncreasing(t, []int{1, 2, 3})
//	assertions.IsIncreasing(t, []float{1, 2})
//	assertions.IsIncreasing(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 2, 3}
//	failure: []int{1, 1, 2}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsIncreasing(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsIncreasing(t, object, msgAndArgs...)
}

// IsNonDecreasing asserts that the collection is not decreasing.
//
// # Usage
//
//	assertions.IsNonDecreasing(t, []int{1, 1, 2})
//	assertions.IsNonDecreasing(t, []float{1, 2})
//	assertions.IsNonDecreasing(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 1, 2}
//	failure: []int{2, 1, 1}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonDecreasing(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonDecreasing(t, object, msgAndArgs...)
}

// IsNonIncreasing asserts that the collection is not increasing.
//
// # Usage
//
//	assertions.IsNonIncreasing(t, []int{2, 1, 1})
//	assertions.IsNonIncreasing(t, []float{2, 1})
//	assertions.IsNonIncreasing(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{2, 1, 1}
//	failure: []int{1, 2, 3}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNonIncreasing(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNonIncreasing(t, object, msgAndArgs...)
}

// IsNotType asserts that the specified objects are not of the same type.
//
// # Usage
//
//	assertions.IsNotType(t, &NotMyStruct{}, &MyStruct{})
//
// # Examples
//
//	success: int32(123), int64(456)
//	failure: 123, 456
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsNotType(t T, theType any, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsNotType(t, theType, object, msgAndArgs...)
}

// IsType asserts that the specified objects are of the same type.
//
// # Usage
//
//	assertions.IsType(t, &MyStruct{}, &MyStruct{})
//
// # Examples
//
//	success: 123, 456
//	failure: int32(123), int64(456)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func IsType(t T, expectedType any, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.IsType(t, expectedType, object, msgAndArgs...)
}

// JSONEq asserts that two JSON strings are equivalent.
//
// # Usage
//
//	assertions.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
//
// # Examples
//
//	success: `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`
//	failure: `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONEq(t T, expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEq(t, expected, actual, msgAndArgs...)
}

// JSONEqBytes asserts that two JSON byte slices are equivalent.
//
// # Usage
//
//	assertions.JSONEqBytes(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
//
// # Examples
//
//	success: []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`)
//	failure: []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func JSONEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.JSONEqBytes(t, expected, actual, msgAndArgs...)
}

// Len asserts that the specified object has specific length.
//
// Len also fails if the object has a type that len() does not accept.
//
// The asserted object can be a string, a slice, a map, an array or a channel.
//
// See also [reflect.Len].
//
// # Usage
//
//	assertions.Len(t, mySlice, 3)
//	assertions.Len(t, myString, 4)
//	assertions.Len(t, myMap, 5)
//
// # Examples
//
//	success: []string{"A","B"}, 2
//	failure: []string{"A","B"}, 1
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Len(t T, object any, length int, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Len(t, object, length, msgAndArgs...)
}

// Less asserts that the first element is strictly less than the second.
//
// # Usage
//
//	assertions.Less(t, 1, 2)
//	assertions.Less(t, float64(1), float64(2))
//	assertions.Less(t, "a", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Less(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Less(t, e1, e2, msgAndArgs...)
}

// LessOrEqual asserts that the first element is less than or equal to the second.
//
// # Usage
//
//	assertions.LessOrEqual(t, 1, 2)
//	assertions.LessOrEqual(t, 2, 2)
//	assertions.LessOrEqual(t, "a", "b")
//	assertions.LessOrEqual(t, "b", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
//
// Upon failure, the test [T] is marked as failed and continues execution.
func LessOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.LessOrEqual(t, e1, e2, msgAndArgs...)
}

// Negative asserts that the specified element is strictly negative.
//
// # Usage
//
//	assertions.Negative(t, -1)
//	assertions.Negative(t, -1.23)
//
// # Examples
//
//	success: -1
//	failure: 1
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Negative(t T, e any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Negative(t, e, msgAndArgs...)
}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
// # Usage
//
//	assertions.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
//
// # Examples
//
//	success:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Never(t, condition, waitFor, tick, msgAndArgs...)
}

// Nil asserts that the specified object is nil.
//
// # Usage
//
//	assertions.Nil(t, err)
//
// # Examples
//
//	success: nil
//	failure: "not nil"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Nil(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Nil(t, object, msgAndArgs...)
}

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
//
// # Usage
//
//	assertions.NoDirExists(t, "path/to/directory")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"non_existing_dir")
//	failure: filepath.Join(testDataPath(),"existing_dir")
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoDirExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoDirExists(t, path, msgAndArgs...)
}

// NoError asserts that a function returned a nil error (ie. no error).
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoError(t T, err error, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoError(t, err, msgAndArgs...)
}

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
//
// # Usage
//
//	assertions.NoFileExists(t, "path/to/file")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"non_existing_file")
//	failure: filepath.Join(testDataPath(),"existing_file")
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NoFileExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NoFileExists(t, path, msgAndArgs...)
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
// # Usage
//
//	assertions.NotContains(t, "Hello World", "Earth")
//	assertions.NotContains(t, ["Hello", "World"], "Earth")
//	assertions.NotContains(t, {"Hello": "World"}, "Earth")
//
// # Examples
//
//	success: []string{"A","B"}, "C"
//	failure: []string{"A","B"}, "B"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotContains(t T, s any, contains any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotContains(t, s, contains, msgAndArgs...)
}

// NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should not match.
// This is an inverse of ElementsMatch.
//
// # Usage
//
//	assertions.NotElementsMatch(t, [1, 1, 2, 3], [1, 1, 2, 3]) -> false
//	assertions.NotElementsMatch(t, [1, 1, 2, 3], [1, 2, 3]) -> true
//	assertions.NotElementsMatch(t, [1, 2, 3], [1, 2, 4]) -> true
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{1, 2, 4}
//	failure: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotElementsMatch(t T, listA any, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotElementsMatch(t, listA, listB, msgAndArgs...)
}

// NotEmpty asserts that the specified object is NOT [Empty].
//
// # Usage
//
//	if assert.NotEmpty(t, obj) {
//		assertions.Equal(t, "two", obj[1])
//	}
//
// # Examples
//
//	success: "not empty"
//	failure: ""
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEmpty(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEmpty(t, object, msgAndArgs...)
}

// NotEqual asserts that the specified values are NOT equal.
//
// # Usage
//
//	assertions.NotEqual(t, obj1, obj2)
//
// Pointer variable equality is determined based on the equality of the
// referenced values (as opposed to the memory addresses).
//
// # Examples
//
//	success: 123, 456
//	failure: 123, 123
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEqual(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqual(t, expected, actual, msgAndArgs...)
}

// NotEqualValues asserts that two objects are not equal even when converted to the same type.
//
// # Usage
//
//	assertions.NotEqualValues(t, obj1, obj2)
//
// # Examples
//
//	success: uint32(123), int32(456)
//	failure: uint32(123), int32(123)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotEqualValues(t, expected, actual, msgAndArgs...)
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotErrorAs(t T, err error, target any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAs(t, err, target, msgAndArgs...)
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
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotErrorIs(t T, err error, target error, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorIs(t, err, target, msgAndArgs...)
}

// NotImplements asserts that an object does not implement the specified interface.
//
// # Usage
//
//	assertions.NotImplements(t, (*MyInterface)(nil), new(MyObject))
//
// # Examples
//
//	success: (*error)(nil), new(testing.T)
//	failure: ptr(dummyInterface), new(testing.T)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotImplements(t T, interfaceObject any, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotImplements(t, interfaceObject, object, msgAndArgs...)
}

// NotNil asserts that the specified object is not nil.
//
// # Usage
//
// assertions.NotNil(t, err)
//
// # Examples
//
//	success: "not nil"
//	failure: nil
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotNil(t T, object any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotNil(t, object, msgAndArgs...)
}

// NotPanics asserts that the code inside the specified PanicTestFunc does NOT panic.
//
// # Usage
//
//	assertions.NotPanics(t, func(){ RemainCalm() })
//
// # Examples
//
//	success: func() { }
//	failure: func() { panic("panicking") }
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotPanics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotPanics(t, f, msgAndArgs...)
}

// NotRegexp asserts that a specified regexp does not match a string.
//
// # Usage
//
//	assertions.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//	assertions.NotRegexp(t, "^start", "it's not starting")
//
// # Examples
//
//	success: "^start", "not starting"
//	failure: "^start", "starting"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotRegexp(t T, rx any, str any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotRegexp(t, rx, str, msgAndArgs...)
}

// NotSame asserts that two pointers do not reference the same object.
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
//
// # Usage
//
//	assertions.NotSame(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, ptr("static string")
//	failure: &staticVar, staticVarPtr
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSame(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSame(t, expected, actual, msgAndArgs...)
}

// NotSubset asserts that the list (array, slice, or map) does NOT contain all
// elements given in the subset (array, slice, or map).
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
// # Usage
//
//	assertions.NotSubset(t, [1, 3, 4], [1, 2])
//	assertions.NotSubset(t, {"x": 1, "y": 2}, {"z": 3})
//	assertions.NotSubset(t, [1, 3, 4], {1: "one", 2: "two"})
//	assertions.NotSubset(t, {"x": 1, "y": 2}, ["z"])
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{4, 5}
//	failure: []int{1, 2, 3}, []int{1, 2}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotSubset(t T, list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotSubset(t, list, subset, msgAndArgs...)
}

// NotZero asserts that i is not the zero value for its type.
//
// # Usage
//
//	assertions.NotZero(t, obj)
//
// # Examples
//
//	success: 1
//	failure: 0
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotZero(t T, i any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotZero(t, i, msgAndArgs...)
}

// Panics asserts that the code inside the specified PanicTestFunc panics.
//
// # Usage
//
//	assertions.Panics(t, func(){ GoCrazy() })
//
// # Examples
//
//	success: func() { panic("panicking") }
//	failure: func() { }
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Panics(t T, f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Panics(t, f, msgAndArgs...)
}

// PanicsWithError asserts that the code inside the specified PanicTestFunc
// panics, and that the recovered panic value is an error that satisfies the
// EqualError comparison.
//
// # Usage
//
//	assertions.PanicsWithError(t, "crazy error", func(){ GoCrazy() })
//
// # Examples
//
//	success: ErrTest.Error(), func() { panic(ErrTest) }
//	failure: ErrTest.Error(), func() { }
//
// Upon failure, the test [T] is marked as failed and continues execution.
func PanicsWithError(t T, errString string, f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithError(t, errString, f, msgAndArgs...)
}

// PanicsWithValue asserts that the code inside the specified PanicTestFunc panics, and that
// the recovered panic value equals the expected panic value.
//
// # Usage
//
//	assertions.PanicsWithValue(t, "crazy error", func(){ GoCrazy() })
//
// # Examples
//
//	success: "panicking", func() { panic("panicking") }
//	failure: "panicking", func() { }
//
// Upon failure, the test [T] is marked as failed and continues execution.
func PanicsWithValue(t T, expected any, f assertions.PanicTestFunc, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.PanicsWithValue(t, expected, f, msgAndArgs...)
}

// Positive asserts that the specified element is strictly positive.
//
// # Usage
//
//	assertions.Positive(t, 1)
//	assertions.Positive(t, 1.23)
//
// # Examples
//
//	success: 1
//	failure: -1
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Positive(t T, e any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Positive(t, e, msgAndArgs...)
}

// Regexp asserts that a specified regexp matches a string.
//
// # Usage
//
//	assertions.Regexp(t, regexp.MustCompile("start"), "it's starting")
//	assertions.Regexp(t, "start...$", "it's not starting")
//
// # Examples
//
//	success: "^start", "starting"
//	failure: "^start", "not starting"
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Regexp(t T, rx any, str any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Regexp(t, rx, str, msgAndArgs...)
}

// Same asserts that two pointers reference the same object.
//
// Both arguments must be pointer variables. Pointer variable sameness is
// determined based on the equality of both type and value.
//
// # Usage
//
//	assertions.Same(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, staticVarPtr
//	failure: &staticVar, ptr("static string")
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Same(t T, expected any, actual any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Same(t, expected, actual, msgAndArgs...)
}

// Subset asserts that the list (array, slice, or map) contains all elements
// given in the subset (array, slice, or map).
//
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
// # Usage
//
//	assertions.Subset(t, [1, 2, 3], [1, 2])
//	assertions.Subset(t, {"x": 1, "y": 2}, {"x": 1})
//	assertions.Subset(t, [1, 2, 3], {1: "one", 2: "two"})
//	assertions.Subset(t, {"x": 1, "y": 2}, ["x"])
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{1, 2}
//	failure: []int{1, 2, 3}, []int{4, 5}
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Subset(t T, list any, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Subset(t, list, subset, msgAndArgs...)
}

// True asserts that the specified value is true.
//
// # Usage
//
//	assertions.True(t, myBool)
//
// # Examples
//
//	success: 1 == 1
//	failure: 1 == 0
//
// Upon failure, the test [T] is marked as failed and continues execution.
func True(t T, value bool, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.True(t, value, msgAndArgs...)
}

// WithinDuration asserts that the two times are within duration delta of each other.
//
// # Usage
//
//	assertions.WithinDuration(t, time.Now(), 10*time.Second)
//
// # Examples
//
//	success: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second
//	failure: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second
//
// Upon failure, the test [T] is marked as failed and continues execution.
func WithinDuration(t T, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.WithinDuration(t, expected, actual, delta, msgAndArgs...)
}

// WithinRange asserts that a time is within a time range (inclusive).
//
// # Usage
//
//	assertions.WithinRange(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
//
// # Examples
//
//	success: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
//	failure: time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
//
// Upon failure, the test [T] is marked as failed and continues execution.
func WithinRange(t T, actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.WithinRange(t, actual, start, end, msgAndArgs...)
}

// YAMLEq asserts that the first documents in the two YAML strings are equivalent.
//
// # Usage
//
//	expected := `---
//	key: value
//	---
//	key: this is a second document, it is not evaluated
//	`
//	actual := `---
//	key: value
//	---
//	key: this is a subsequent document, it is not evaluated
//	`
//	assertions.YAMLEq(t, expected, actual)
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
//
// Upon failure, the test [T] is marked as failed and continues execution.
func YAMLEq(t T, expected string, actual string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.YAMLEq(t, expected, actual, msgAndArgs...)
}

// Zero asserts that i is the zero value for its type.
//
// # Usage
//
//	assertions.Zero(t, obj)
//
// # Examples
//
//	success: 0
//	failure: 1
//
// Upon failure, the test [T] is marked as failed and continues execution.
func Zero(t T, i any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.Zero(t, i, msgAndArgs...)
}
