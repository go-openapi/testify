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

// Condition uses a [Comparison] to assert a complex condition.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Condition(t T, comp Comparison, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Condition(t, comp, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Contains(t T, s any, contains any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Contains(t, s, contains, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func DirExists(t T, path string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.DirExists(t, path, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// DirNotExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
//
// # Usage
//
//	assertions.DirNotExists(t, "path/to/directory")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"non_existing_dir")
//	failure: filepath.Join(testDataPath(),"existing_dir")
//
// Upon failure, the test [T] is marked as failed and stops execution.
func DirNotExists(t T, path string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.DirNotExists(t, path, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// # Usage
//
//	assertions.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
//
// # Examples
//
//	success: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
//	failure: []int{1, 2, 3}, []int{1, 2, 4}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ElementsMatch(t T, listA any, listB any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatch(t, listA, listB, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// ElementsMatchT asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// # Usage
//
//	assertions.ElementsMatchT(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
//
// # Examples
//
//	success: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
//	failure: []int{1, 2, 3}, []int{1, 2, 4}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ElementsMatchT[E](t, listA, listB, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
//
// [Zero values]: https://go.dev/ref/spec#The_zero_value
func Empty(t T, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Empty(t, object, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Equal(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Equal(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualError(t T, err error, errString string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualError(t, err, errString, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// EqualExportedValues asserts that the types of two objects are equal and their public
// fields are also equal.
//
// This is useful for comparing structs that have private fields that could potentially differ.
//
// Function equality cannot be determined and will always fail.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualExportedValues(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualExportedValues(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// EqualT asserts that two objects of the same comparable type are equal.
//
// Pointer variable equality is determined based on the equality of the memory addresses (unlike [Equal], but like [Same]).
//
// Functions, slices and maps are not comparable. See also [ComparisonOperators].
//
// If you need to compare values of non-comparable types, or compare pointers by the value they point to,
// use [Equal] instead.
//
// # Usage
//
//	assertions.EqualT(t, 123, 123)
//
// # Examples
//
//	success: 123, 123
//	failure: 123, 456
//
// Upon failure, the test [T] is marked as failed and stops execution.
//
// [ComparisonOperators]: https://go.dev/ref/spec#Comparison_operators.
func EqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualT[V](t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// EqualValues asserts that two objects are equal or convertible to the larger
// type and equal.
//
// Function equality cannot be determined and will always fail.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func EqualValues(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EqualValues(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Error asserts that a function returned a non-nil error (i.e. an error).
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Error(t T, err error, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Error(t, err, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorAs(t T, err error, target any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAs(t, err, target, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorContains(t T, err error, contains string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorContains(t, err, contains, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorIs(t T, err error, target error, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorIs(t, err, target, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Eventually asserts that the given condition will be met in waitFor time,
// periodically checking the target function on each tick.
//
// [Eventually] waits until the condition returns true, for at most waitFor,
// or until the parent context of the test is cancelled.
//
// If the condition takes longer than waitFor to complete, [Eventually] fails
// but waits for the current condition execution to finish before returning.
//
// For long-running conditions to be interrupted early, check [testing.T.Context]
// which is cancelled on test failure.
//
// # Usage
//
//	assertions.Eventually(t, func() bool { return true }, time.Second, 10*time.Millisecond)
//
// # Concurrency
//
// The condition function is never executed in parallel: only one goroutine executes it.
// It may write to variables outside its scope without triggering race conditions.
//
// A blocking condition will cause [Eventually] to hang until it returns.
//
// # Examples
//
//	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Eventually(t, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// EventuallyWith asserts that the given condition will be met in waitFor time,
// periodically checking the target function at each tick.
//
// In contrast to [Eventually], the condition function is supplied with a [CollectT]
// to accumulate errors from calling other assertions.
//
// The condition is considered "met" if no errors are raised in a tick.
// The supplied [CollectT] collects all errors from one tick.
//
// If the condition is not met before waitFor, the collected errors from the
// last tick are copied to t.
//
// Calling [CollectT.FailNow] cancels the condition immediately and fails the assertion.
//
// # Usage
//
//	externalValue := false
//	go func() {
//		time.Sleep(8*time.Second)
//		externalValue = true
//	}()
//
//	assertions.EventuallyWith(t, func(c *assertions.CollectT) {
//		// add assertions as needed; any assertion failure will fail the current tick
//		assertions.True(c, externalValue, "expected 'externalValue' to be true")
//	},
//	10*time.Second,
//	1*time.Second,
//	"external state has not changed to 'true'; still false",
//	)
//
// # Concurrency
//
// The condition function is never executed in parallel: only one goroutine executes it.
// It may write to variables outside its scope without triggering race conditions.
//
// # Examples
//
//	success: func(c *CollectT) { True(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//	failure: func(c *CollectT) { False(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//
// Upon failure, the test [T] is marked as failed and stops execution.
func EventuallyWith(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.EventuallyWith(t, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Exactly(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Exactly(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Fail(t T, failureMessage string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	_ = assertions.Fail(t, failureMessage, msgAndArgs...)

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func FailNow(t T, failureMessage string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	_ = assertions.FailNow(t, failureMessage, msgAndArgs...)

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func False(t T, value bool, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.False(t, value, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// FalseT asserts that the specified value is false.
//
// The type constraint [Boolean] accepts any type which underlying type is bool.
//
// # Usage
//
//	 type B bool
//	 var b B = true
//
//		assertions.FalseT(t, b)
//
// # Examples
//
//	success: 1 == 0
//	failure: 1 == 1
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FalseT[B Boolean](t T, value B, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FalseT[B](t, value, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func FileEmpty(t T, path string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileEmpty(t, path, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func FileExists(t T, path string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileExists(t, path, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func FileNotEmpty(t T, path string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileNotEmpty(t, path, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// FileNotExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
//
// # Usage
//
//	assertions.FileNotExists(t, "path/to/file")
//
// # Examples
//
//	success: filepath.Join(testDataPath(),"non_existing_file")
//	failure: filepath.Join(testDataPath(),"existing_file")
//
// Upon failure, the test [T] is marked as failed and stops execution.
func FileNotExists(t T, path string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.FileNotExists(t, path, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Greater asserts that the first element is strictly greater than the second.
//
// Both elements must be of the same type in the [reflect.Kind] sense.
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Greater(t T, e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Greater(t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// GreaterOrEqual asserts that the first element is greater than or equal to the second.
//
// See also [Greater].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func GreaterOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqual(t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// GreaterOrEqualT asserts that for two elements of the same type,
// the first element is greater than or equal to the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [GreaterOrEqualT] with [*time.Time].
//
// [GreaterOrEqualT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
// use [GreaterOrEqual] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.GreaterOrEqualT(t, 2, 1)
//	assertions.GreaterOrEqualT(t, 2, 2)
//	assertions.GreaterOrEqualT(t, "b", "a")
//	assertions.GreaterOrEqualT(t, "b", "b")
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
//
// Upon failure, the test [T] is marked as failed and stops execution.
func GreaterOrEqualT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterOrEqualT[Orderable](t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// GreaterT asserts that for two elements of the same type,
// the first element is strictly greater than the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [GreaterT] with [*time.Time].
//
// [GreaterT] ensures type safety at build time. If you need to compare values with a dynamically assigned type, use [Greater] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.GreaterT(t, 2, 1)
//	assertions.GreaterT(t, float64(2), float64(1))
//	assertions.GreaterT(t, "b", "a")
//	assertions.GreaterT(t, time.Date(2026,1,1,0,0,0,0,nil), time.Now())
//
// # Examples
//
//	success: 2, 1
//	failure: 1, 2
//
// Upon failure, the test [T] is marked as failed and stops execution.
func GreaterT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.GreaterT[Orderable](t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPBodyContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyContains(t, handler, method, url, values, str, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPBodyNotContains(t T, handler http.HandlerFunc, method string, url string, values url.Values, str any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPBodyNotContains(t, handler, method, url, values, str, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPError(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPError(t, handler, method, url, values, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPRedirect(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPRedirect(t, handler, method, url, values, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPStatusCode(t T, handler http.HandlerFunc, method string, url string, values url.Values, statuscode int, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPStatusCode(t, handler, method, url, values, statuscode, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func HTTPSuccess(t T, handler http.HandlerFunc, method string, url string, values url.Values, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.HTTPSuccess(t, handler, method, url, values, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Implements(t T, interfaceObject any, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Implements(t, interfaceObject, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// InDelta asserts that the two numerals are within delta of each other.
//
// Delta must be greater than or equal to zero.
//
// Expected and actual values should convert to float64.
// To compare large integers that can't be represented accurately as float64 (e.g. uint64),
// prefer [InDeltaT] to preserve the original type.
//
// # Behavior with IEEE floating point arithmetic
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
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
// Upon failure, the test [T] is marked as failed and stops execution.
func InDelta(t T, expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDelta(t, expected, actual, delta, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// InDeltaMapValues is the same as [InDelta], but it compares all values between two maps. Both maps must have exactly the same keys.
//
// See [InDelta].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func InDeltaMapValues(t T, expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaMapValues(t, expected, actual, delta, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// InDeltaSlice is the same as [InDelta], except it compares two slices.
//
// See [InDelta].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func InDeltaSlice(t T, expected any, actual any, delta float64, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaSlice(t, expected, actual, delta, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// InDeltaT asserts that the two numerals of the same type numerical type are within delta of each other.
//
// [InDeltaT] accepts any go numeric type, including integer types.
//
// The main difference with [InDelta] is that the delta is expressed with the same type as the values, not necessarily a float64.
//
// Delta must be greater than or equal to zero.
//
// # Behavior with IEEE floating point arithmetic
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
//
// # Usage
//
// assertions.InDeltaT(t, math.Pi, 22/7.0, 0.01)
//
// # Examples
//
//	success: 1.0, 1.01, 0.02
//	failure: 1.0, 1.1, 0.05
//
// Upon failure, the test [T] is marked as failed and stops execution.
func InDeltaT[Number Measurable](t T, expected Number, actual Number, delta Number, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InDeltaT[Number](t, expected, actual, delta, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// InEpsilon asserts that expected and actual have a relative error less than epsilon.
//
// # Behavior with IEEE floating point arithmetic
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
//
// Edge case: for very large integers that do not convert accurately to a float64 (e.g. uint64), prefer [InDeltaT].
//
// Formula:
//   - If expected == 0: fail if |actual - expected| > epsilon
//   - If expected != 0: fail if |actual - expected| > epsilon * |expected|
//
// This allows [InEpsilonT] to work naturally across the full numeric range including zero.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func InEpsilon(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilon(t, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// InEpsilonSlice is the same as [InEpsilon], except it compares each value from two slices.
//
// See [InEpsilon].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func InEpsilonSlice(t T, expected any, actual any, epsilon float64, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonSlice(t, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// InEpsilonT asserts that expected and actual have a relative error less than epsilon.
//
// When expected is zero, epsilon is interpreted as an absolute error threshold,
// since relative error is mathematically undefined for zero values.
//
// Unlike [InDeltaT], which preserves the original type, [InEpsilonT] converts the expected and actual
// numbers to float64, since the relative error doesn't make sense as an integer.
//
// # Behavior with IEEE floating point arithmetic
//
//   - expected NaN is matched only by a NaN, e.g. this works: InDeltaT(math.NaN(), math.Sqrt(-1), 0.0)
//   - expected +Inf is matched only by a +Inf
//   - expected -Inf is matched only by a -Inf
//
// Edge case: for very large integers that do not convert accurately to a float64 (e.g. uint64), prefer [InDeltaT].
//
// Formula:
//   - If expected == 0: fail if |actual - expected| > epsilon
//   - If expected != 0: fail if |actual - expected| > epsilon * |expected|
//
// This allows [InEpsilonT] to work naturally across the full numeric range including zero.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func InEpsilonT[Number Measurable](t T, expected Number, actual Number, epsilon float64, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.InEpsilonT[Number](t, expected, actual, epsilon, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsDecreasing asserts that the collection is strictly decreasing.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func IsDecreasing(t T, collection any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasing(t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsDecreasingT asserts that a slice of [Ordered] is strictly decreasing.
//
// # Usage
//
//	assertions.IsDecreasingT(t, []int{2, 1, 0})
//	assertions.IsDecreasingT(t, []float{2, 1})
//	assertions.IsDecreasingT(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{3, 2, 1}
//	failure: []int{1, 2, 3}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsDecreasingT[OrderedSlice, E](t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsIncreasing asserts that the collection is strictly increasing.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func IsIncreasing(t T, collection any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasing(t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsIncreasingT asserts that a slice of [Ordered] is strictly increasing.
//
// # Usage
//
//	assertions.IsIncreasingT(t, []int{1, 2, 3})
//	assertions.IsIncreasingT(t, []float{1, 2})
//	assertions.IsIncreasingT(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 2, 3}
//	failure: []int{1, 1, 2}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsIncreasingT[OrderedSlice, E](t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsNonDecreasing asserts that the collection is not strictly decreasing.
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
//	failure: []int{2, 1, 0}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonDecreasing(t T, collection any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasing(t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsNonDecreasingT asserts that a slice of [Ordered] is not decreasing.
//
// # Usage
//
//	assertions.IsNonDecreasingT(t, []int{1, 1, 2})
//	assertions.IsNonDecreasingT(t, []float{1, 2})
//	assertions.IsNonDecreasingT(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 1, 2}
//	failure: []int{2, 1, 0}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonDecreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonDecreasingT[OrderedSlice, E](t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonIncreasing(t T, collection any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasing(t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsNonIncreasingT asserts that a slice of [Ordered] is NOT strictly increasing.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNonIncreasingT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNonIncreasingT[OrderedSlice, E](t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsNotOfTypeT asserts that an object is not of a given type.
//
// # Usage
//
//	assertions.IsOfType[MyType](t,myVar)
//
// # Examples
//
//	success: 123.123
//	failure: myType(123.123)
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNotOfTypeT[EType any](t T, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNotOfTypeT[EType](t, object, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func IsNotType(t T, theType any, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsNotType(t, theType, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// IsOfTypeT asserts that an object is of a given type.
//
// # Usage
//
//	assertions.IsOfTypeT[MyType](t,myVar)
//
// # Examples
//
//	success: myType(123.123)
//	failure: 123.123
//
// Upon failure, the test [T] is marked as failed and stops execution.
func IsOfTypeT[EType any](t T, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsOfTypeT[EType](t, object, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func IsType(t T, expectedType any, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.IsType(t, expectedType, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// JSONEq asserts that two JSON strings are semantically equivalent.
//
// Expected and actual must be valid JSON.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONEq(t T, expected string, actual string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEq(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// JSONEqBytes asserts that two JSON slices of bytes are semantically equivalent.
//
// Expected and actual must be valid JSON.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqBytes(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// JSONEqT asserts that two JSON documents are semantically equivalent.
//
// The expected and actual arguments may be string or []byte. They do not need to be of the same type.
//
// Expected and actual must be valid JSON.
//
// # Usage
//
//	assertions.JSONEqT(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
//
// # Examples
//
//	success: `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`)
//	failure: `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONEqT[EDoc, ADoc](t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// JSONMarshalAsT wraps [JSONEq] after [json.Marshal].
//
// The input JSON may be a string or []byte.
//
// It fails if the marshaling returns an error or if the expected JSON bytes differ semantically
// from the expected ones.
//
// # Usage
//
//	actual := struct {
//		A int `json:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.JSONUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	success: []byte(`{"A": "a"}`), dummyStruct{A: "a"}
//	failure: `[{"foo": "bar"}, {"hello": "world"}]`, 1
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONMarshalAsT[EDoc](t, expected, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// JSONUnmarshalAsT wraps [Equal] after [json.Unmarshal].
//
// The input JSON may be a string or []byte.
//
// It fails if the unmarshaling returns an error or if the resulting object is not equal to the expected one.
//
// Be careful not to wrap the expected object into an "any" interface if this is not what you expected:
// the unmarshaling would take this type to unmarshal as a map[string]any.
//
// # Usage
//
//	expected := struct {
//		A int `json:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.JSONUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	success: dummyStruct{A: "a"} , []byte(`{"A": "a"}`)
//	failure: 1, `[{"foo": "bar"}, {"hello": "world"}]`
//
// Upon failure, the test [T] is marked as failed and stops execution.
func JSONUnmarshalAsT[Object any, ADoc Text](t T, expected Object, jazon ADoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.JSONUnmarshalAsT[Object, ADoc](t, expected, jazon, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Kind asserts that the [reflect.Kind] of a given object matches the expected [reflect.Kind].
//
// Kind reflects the concrete value stored in the object. The nil value (or interface with nil value)
// are comparable to [reflect.Invalid]. See also [reflect.Value.Kind].
//
// # Usage
//
//	assertions.Kind(t, reflect.String, "Hello World")
//
// # Examples
//
//	success: reflect.String, "hello"
//	failure: reflect.String, 0
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Kind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Kind(t, expectedKind, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Len asserts that the specified object has specific length.
//
// Len also fails if the object has a type that len() does not accept.
//
// The asserted object can be a string, a slice, a map, an array, pointer to array or a channel.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Len(t T, object any, length int, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Len(t, object, length, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Less asserts that the first element is strictly less than the second.
//
// Both elements must be of the same type in the [reflect.Kind] sense.
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Less(t T, e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Less(t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func LessOrEqual(t T, e1 any, e2 any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqual(t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// LessOrEqualT asserts that for two elements of the same type, the first element is less than or equal to the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [LessOrEqualT] with [*time.Time].
//
// [LessOrEqualT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
// use [LessOrEqual] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you should use [LessOrEqual] instead.
//
// # Usage
//
//	assertions.LessOrEqualT(t, 1, 2)
//	assertions.LessOrEqualT(t, 2, 2)
//	assertions.LessOrEqualT(t, "a", "b")
//	assertions.LessOrEqualT(t, "b", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
//
// Upon failure, the test [T] is marked as failed and stops execution.
func LessOrEqualT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.LessOrEqualT[Orderable](t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// LessT asserts that for two elements of the same type, the first element is strictly less than the second.
//
// The [Ordered] type can be any of Go's [cmp.Ordered] (strings, numeric types),
// []byte (uses [bytes.Compare]) and [time.Time] (uses [time.Time.Compare].
//
// Notice that pointers are not [Ordered], but uintptr are. So you can't call [LessT] with [*time.Time].
//
// [LessT] ensures type safety at build time. If you need to compare values with a dynamically assigned type,
// use [Less] instead.
//
// To compare values that need a type conversion (e.g. float32 against float64), you need to convert types beforehand.
//
// # Usage
//
//	assertions.LessT(t, 1, 2)
//	assertions.LessT(t, float64(1), float64(2))
//	assertions.LessT(t, "a", "b")
//
// # Examples
//
//	success: 1, 2
//	failure: 2, 1
//
// Upon failure, the test [T] is marked as failed and stops execution.
func LessT[Orderable Ordered](t T, e1 Orderable, e2 Orderable, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.LessT[Orderable](t, e1, e2, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// MapContainsT asserts that the specified map contains a key.
//
// Go native comparable types are explained there: [comparable-types].
//
// # Usage
//
//	assertions.MapContainsT(t, map[string]string{"Hello": "x","World": "y"}, "World")
//
// # Examples
//
//	success: map[string]string{"A": "B"}, "A"
//	failure: map[string]string{"A": "B"}, "C"
//
// Upon failure, the test [T] is marked as failed and stops execution.
//
// [comparable-types]: https://go.dev/blog/comparable
func MapContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.MapContainsT[Map, K, V](t, m, key, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// MapNotContainsT asserts that the specified map does not contain a key.
//
// # Usage
//
//	assertions.MapNotContainsT(t, map[string]string{"Hello": "x","World": "y"}, "hi")
//
// # Examples
//
//	success: map[string]string{"A": "B"}, "C"
//	failure: map[string]string{"A": "B"}, "A"
//
// Upon failure, the test [T] is marked as failed and stops execution.
func MapNotContainsT[Map ~map[K]V, K comparable, V any](t T, m Map, key K, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.MapNotContainsT[Map, K, V](t, m, key, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Negative(t T, e any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Negative(t, e, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NegativeT asserts that the specified element of a signed numeric type is strictly negative.
//
// # Usage
//
//	assertions.NegativeT(t, -1)
//	assertions.NegativeT(t, -1.23)
//
// # Examples
//
//	success: -1
//	failure: 1
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NegativeT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NegativeT[SignedNumber](t, e, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Never asserts that the given condition is never satisfied within waitFor time,
// periodically checking the target function at each tick.
//
// [Never] is the opposite of [Eventually]. It succeeds if the waitFor timeout
// is reached without the condition ever returning true.
//
// If the parent context is cancelled before the timeout, [Never] fails.
//
// # Usage
//
//	assertions.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond)
//
// # Concurrency
//
// The condition function is never executed in parallel: only one goroutine executes it.
// It may write to variables outside its scope without triggering race conditions.
//
// A blocking condition will cause [Never] to hang until it returns.
//
// # Examples
//
//	success:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Never(t, condition, waitFor, tick, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Nil(t T, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Nil(t, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NoError asserts that a function returned a nil error (i.e. no error).
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NoError(t T, err error, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoError(t, err, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NoFileDescriptorLeak ensures that no file descriptor leaks from inside the tested function.
//
// This assertion works on Linux only (via /proc/self/fd).
// On other platforms, the test is skipped.
//
// NOTE: this assertion is not compatible with parallel tests.
// File descriptors are a process-wide resource; concurrent tests
// opening files would cause false positives.
//
// Sockets, pipes, and anonymous inodes are filtered out by default,
// as these are typically managed by the Go runtime.
//
// # Concurrency
//
// [NoFileDescriptorLeak] is not compatible with parallel tests.
// File descriptors are a process-wide resource; any concurrent I/O
// from other goroutines may cause false positives.
//
// Calls to [NoFileDescriptorLeak] are serialized with a mutex
// to prevent multiple leak checks from interfering with each other.
//
// # Usage
//
//	NoFileDescriptorLeak(t, func() {
//		// code that should not leak file descriptors
//	})
//
// # Examples
//
//	success: func() {}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NoFileDescriptorLeak(t T, tested func(), msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoFileDescriptorLeak(t, tested, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NoGoRoutineLeak ensures that no goroutine did leak from inside the tested function.
//
// NOTE: only the go routines spawned from inside the tested function are checked for leaks.
// No filter or configuration is needed to exclude "known go routines".
//
// Resource cleanup should be done inside the tested function, and not using [testing.T.Cleanup],
// as t.Cleanup is called after the leak check.
//
// # Edge cases
//
//   - if the tested function panics leaving behind leaked goroutines, these are detected.
//   - if the tested function calls runtime.Goexit (e.g. from [testing.T.FailNow]) leaving behind leaked goroutines,
//     these are detected.
//   - if a panic occurs in one of the leaked go routines, it cannot be recovered with certainty and
//     the calling program will usually panic.
//
// # Concurrency
//
// [NoGoRoutineLeak] may be used safely in parallel tests.
//
// # Usage
//
//	NoGoRoutineLeak(t, func() {
//		...
//	},
//	"should not leak any go routine",
//	)
//
// # Examples
//
//	success: func() {}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NoGoRoutineLeak(t T, tested func(), msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NoGoRoutineLeak(t, tested, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotContains(t T, s any, contains any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotContains(t, s, contains, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should not match.
// This is an inverse of ElementsMatch.
//
// # Usage
//
//	assertions.NotElementsMatch(t, []int{1, 1, 2, 3}, []int{1, 1, 2, 3}) -> false
//	assertions.NotElementsMatch(t, []int{1, 1, 2, 3}, []int{1, 2, 3}) -> true
//	assertions.NotElementsMatch(t, []int{1, 2, 3}, []int{1, 2, 4}) -> true
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{1, 2, 4}
//	failure: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotElementsMatch(t T, listA any, listB any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatch(t, listA, listB, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotElementsMatchT asserts that the specified listA(array, slice...) is NOT equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should not match.
// This is an inverse of ElementsMatch.
//
// # Usage
//
//	assertions.NotElementsMatchT(t, []int{1, 1, 2, 3}, []int{1, 1, 2, 3}) -> false
//	assertions.NotElementsMatchT(t, []int{1, 1, 2, 3}, []int{1, 2, 3}) -> true
//	assertions.NotElementsMatchT(t, []int{1, 2, 3}, []int{1, 2, 4}) -> true
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{1, 2, 4}
//	failure: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotElementsMatchT[E comparable](t T, listA []E, listB []E, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotElementsMatchT[E](t, listA, listB, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEmpty(t T, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEmpty(t, object, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Function equality cannot be determined and will always fail.
//
// # Examples
//
//	success: 123, 456
//	failure: 123, 123
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEqual(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqual(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotEqualT asserts that the specified values of the same comparable type are NOT equal.
//
// See [EqualT].
//
// # Usage
//
//	assertions.NotEqualT(t, obj1, obj2)
//
// # Examples
//
//	success: 123, 456
//	failure: 123, 123
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEqualT[V comparable](t T, expected V, actual V, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualT[V](t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotEqualValues asserts that two objects are not equal even when converted to the same type.
//
// Function equality cannot be determined and will always fail.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotEqualValues(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotEqualValues(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotErrorAs(t T, err error, target any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAs(t, err, target, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotErrorIs(t T, err error, target error, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorIs(t, err, target, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotImplements(t T, interfaceObject any, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotImplements(t, interfaceObject, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotKind asserts that the [reflect.Kind] of a given object does not match the expected [reflect.Kind].
//
// Kind reflects the concrete value stored in the object. The nil value (or interface with nil value)
// are comparable to [reflect.Invalid]. See also [reflect.Value.Kind].
//
// # Usage
//
//	assertions.NotKind(t, reflect.Int, "Hello World")
//
// # Examples
//
//	success: reflect.String, 0
//	failure: reflect.String, "hello"
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotKind(t T, expectedKind reflect.Kind, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotKind(t, expectedKind, object, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotNil(t T, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotNil(t, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotPanics asserts that the code inside the specified function does NOT panic.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotPanics(t T, f func(), msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotPanics(t, f, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotRegexp asserts that a specified regular expression does not match a string.
//
// See [Regexp].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotRegexp(t T, rx any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexp(t, rx, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotRegexpT asserts that a specified regular expression does not match a string.
//
// See [RegexpT].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotRegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotRegexpT[Rex, ADoc](t, rx, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotSame asserts that two pointers do not reference the same object.
//
// See [Same].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSame(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSame(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotSameT asserts that two pointers do not reference the same object.
//
// See [SameT].
//
// # Usage
//
//	assertions.NotSameT(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, ptr("static string")
//	failure: &staticVar, staticVarPtr
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSameT[P](t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// NotSortedT asserts that the slice of [Ordered] is NOT sorted (i.e. non-strictly increasing).
//
// Unlike [IsDecreasingT], it accepts slices that are neither increasing nor decreasing.
//
// # Usage
//
//	assertions.NotSortedT(t, []int{3, 2, 3})
//	assertions.NotSortedT(t, []float{2, 1})
//	assertions.NotSortedT(t, []string{"b", "a"})
//
// # Examples
//
//	success: []int{3, 1, 3}
//	failure: []int{1, 4, 8}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSortedT[OrderedSlice, E](t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotSubset(t T, list any, subset any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotSubset(t, list, subset, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func NotZero(t T, i any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotZero(t, i, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Panics asserts that the code inside the specified function panics.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Panics(t T, f func(), msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Panics(t, f, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// PanicsWithError asserts that the code inside the specified function panics,
// and that the recovered panic value is an error that satisfies the EqualError comparison.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func PanicsWithError(t T, errString string, f func(), msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithError(t, errString, f, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// PanicsWithValue asserts that the code inside the specified function panics,
// and that the recovered panic value equals the expected panic value.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func PanicsWithValue(t T, expected any, f func(), msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PanicsWithValue(t, expected, f, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Positive(t T, e any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Positive(t, e, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// PositiveT asserts that the specified element of a signed numeric type is strictly positive.
//
// # Usage
//
//	assertions.PositiveT(t, 1)
//	assertions.PositiveT(t, 1.23)
//
// # Examples
//
//	success: 1
//	failure: -1
//
// Upon failure, the test [T] is marked as failed and stops execution.
func PositiveT[SignedNumber SignedNumeric](t T, e SignedNumber, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.PositiveT[SignedNumber](t, e, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Regexp asserts that a specified regular expression matches a string.
//
// The regular expression may be passed as a [regexp.Regexp], a string or a []byte and will be compiled.
//
// The actual argument to be matched may be a string, []byte or anything that prints as a string with [fmt.Sprint].
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Regexp(t T, rx any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Regexp(t, rx, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// RegexpT asserts that a specified regular expression matches a string.
//
// The actual argument to be matched may be a string or []byte.
//
// See [Regexp].
//
// # Examples
//
//	success: "^start", "starting"
//	failure: "^start", "not starting"
//
// Upon failure, the test [T] is marked as failed and stops execution.
func RegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.RegexpT[Rex, ADoc](t, rx, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Same asserts that two pointers reference the same object.
//
// Both arguments must be pointer variables.
//
// Pointer variable sameness is determined based on the equality of both type and value.
//
// Unlike [Equal] pointers, [Same] pointers point to the same memory address.
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Same(t T, expected any, actual any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Same(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SameT asserts that two pointers of the same type reference the same object.
//
// See [Same].
//
// # Usage
//
//	assertions.SameT(t, ptr1, ptr2)
//
// # Examples
//
//	success: &staticVar, staticVarPtr
//	failure: &staticVar, ptr("static string")
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SameT[P any](t T, expected *P, actual *P, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SameT[P](t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SeqContainsT asserts that the specified iterator contains a comparable element.
//
// The sequence may not be consumed entirely: the iteration stops as soon as the specified element is found.
//
// Go native comparable types are explained there: [comparable-types].
//
// # Usage
//
//	assertions.SeqContainsT(t, slices.Values([]{"Hello","World"}), "World")
//
// # Examples
//
//	success: slices.Values([]string{"A","B"}), "A"
//	failure: slices.Values([]string{"A","B"}), "C"
//
// Upon failure, the test [T] is marked as failed and stops execution.
//
// [comparable-types]: https://go.dev/blog/comparable
func SeqContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SeqContainsT[E](t, iter, element, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SeqNotContainsT asserts that the specified iterator does not contain a comparable element.
//
// See [SeqContainsT].
//
// # Usage
//
//	assertions.SeqContainsT(t, slices.Values([]{"Hello","World"}), "World")
//
// # Examples
//
//	success: slices.Values([]string{"A","B"}), "C"
//	failure: slices.Values([]string{"A","B"}), "A"
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SeqNotContainsT[E comparable](t T, iter iter.Seq[E], element E, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SeqNotContainsT[E](t, iter, element, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SliceContainsT asserts that the specified slice contains a comparable element.
//
// Go native comparable types are explained there: [comparable-types].
//
// # Usage
//
//	assertions.SliceContainsT(t, []{"Hello","World"}, "World")
//
// # Examples
//
//	success: []string{"A","B"}, "A"
//	failure: []string{"A","B"}, "C"
//
// Upon failure, the test [T] is marked as failed and stops execution.
//
// [comparable-types]: https://go.dev/blog/comparable
func SliceContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceContainsT[Slice, E](t, s, element, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SliceNotContainsT asserts that the specified slice does not contain a comparable element.
//
// See [SliceContainsT].
//
// # Usage
//
//	assertions.SliceNotContainsT(t, []{"Hello","World"}, "hi")
//
// # Examples
//
//	success: []string{"A","B"}, "C"
//	failure: []string{"A","B"}, "A"
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SliceNotContainsT[Slice ~[]E, E comparable](t T, s Slice, element E, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotContainsT[Slice, E](t, s, element, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SliceNotSubsetT asserts that a slice of comparable elements does not contain all the elements given in the subset.
//
// # Usage
//
//	assertions.SliceNotSubsetT(t, []int{1, 2, 3}, []int{1, 4})
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{4, 5}
//	failure: []int{1, 2, 3}, []int{1, 2}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SliceNotSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceNotSubsetT[Slice, E](t, list, subset, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SliceSubsetT asserts that a slice of comparable elements contains all the elements given in the subset.
//
// # Usage
//
//	assertions.SliceSubsetT(t, []int{1, 2, 3}, []int{1, 2})
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{1, 2}
//	failure: []int{1, 2, 3}, []int{4, 5}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SliceSubsetT[Slice ~[]E, E comparable](t T, list Slice, subset Slice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SliceSubsetT[Slice, E](t, list, subset, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// SortedT asserts that the slice of [Ordered] is sorted (i.e. non-strictly increasing).
//
// Unlike [IsIncreasingT], it accepts elements to be equal.
//
// # Usage
//
//	assertions.SortedT(t, []int{1, 2, 3})
//	assertions.SortedT(t, []float{1, 2})
//	assertions.SortedT(t, []string{"a", "b"})
//
// # Examples
//
//	success: []int{1, 1, 3}
//	failure: []int{1, 4, 2}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func SortedT[OrderedSlice ~[]E, E Ordered](t T, collection OrderedSlice, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.SortedT[OrderedSlice, E](t, collection, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// StringContainsT asserts that a string contains the specified substring.
//
// Strings may be go strings or []byte according to the type constraint [Text].
//
// # Usage
//
//	assertions.StringContainsT(t, "Hello World", "World")
//
// # Examples
//
//	success: "AB", "A"
//	failure: "AB", "C"
//
// Upon failure, the test [T] is marked as failed and stops execution.
func StringContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.StringContainsT[ADoc, EDoc](t, str, substring, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// StringNotContainsT asserts that a string does not contain the specified substring.
//
// See [StringContainsT].
//
// # Usage
//
//	assertions.StringNotContainsT(t, "Hello World", "hi")
//
// # Examples
//
//	success: "AB", "C"
//	failure: "AB", "A"
//
// Upon failure, the test [T] is marked as failed and stops execution.
func StringNotContainsT[ADoc, EDoc Text](t T, str ADoc, substring EDoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.StringNotContainsT[ADoc, EDoc](t, str, substring, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// Subset asserts that the list (array, slice, or map) contains all elements
// given in the subset (array, slice, or map).
//
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
// nil values are considered as empty sets.
//
// # Usage
//
//	assertions.Subset(t, []int{1, 2, 3}, []int{1, 2})
//	assertions.Subset(t, []string{"x": 1, "y": 2}, []string{"x": 1})
//	assertions.Subset(t, []int{1, 2, 3}, map[int]string{1: "one", 2: "two"})
//	assertions.Subset(t, map[string]int{"x": 1, "y": 2}, []string{"x"})
//
// # Examples
//
//	success: []int{1, 2, 3}, []int{1, 2}
//	failure: []int{1, 2, 3}, []int{4, 5}
//
// Upon failure, the test [T] is marked as failed and stops execution.
func Subset(t T, list any, subset any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Subset(t, list, subset, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func True(t T, value bool, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.True(t, value, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// TrueT asserts that the specified value is true.
//
// The type constraint [Boolean] accepts any type which underlying type is bool.
//
// # Usage
//
//	type B bool
//	var b B = true
//
//	assertions.True(t, b)
//
// # Examples
//
//	success: 1 == 1
//	failure: 1 == 0
//
// Upon failure, the test [T] is marked as failed and stops execution.
func TrueT[B Boolean](t T, value B, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.TrueT[B](t, value, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func WithinDuration(t T, expected time.Time, actual time.Time, delta time.Duration, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.WithinDuration(t, expected, actual, delta, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func WithinRange(t T, actual time.Time, start time.Time, end time.Time, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.WithinRange(t, actual, start, end, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// YAMLEq asserts that two YAML strings are equivalent.
//
// See [YAMLEqBytes].
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLEq(t T, expected string, actual string, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEq(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// YAMLEqBytes asserts that two YAML slices of bytes are equivalent.
//
// Expected and actual must be valid YAML.
//
// # Important
//
// By default, this function is disabled and will panic.
//
// To enable it, you should add a blank import like so:
//
//	import(
//	  "github.com/go-openapi/testify/enable/yaml/v2"
//	)
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
//	panic: []byte("key: value"), []byte("key: value")
//	should panic without the yaml feature enabled.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLEqBytes(t T, expected []byte, actual []byte, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqBytes(t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// YAMLEqT asserts that two YAML documents are equivalent.
//
// The expected and actual arguments may be string or []byte. They do not need to be of the same type.
//
// See [YAMLEqBytes].
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLEqT[EDoc, ADoc Text](t T, expected EDoc, actual ADoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLEqT[EDoc, ADoc](t, expected, actual, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// YAMLMarshalAsT wraps [YAMLEq] after [yaml.Marshal].
//
// The input YAML may be a string or []byte.
//
// It fails if the marshaling returns an error or if the expected YAML bytes differ semantically
// from the expected ones.
//
// # Usage
//
//	actual := struct {
//		A int `yaml:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.YAMLUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLMarshalAsT[EDoc Text](t T, expected EDoc, object any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLMarshalAsT[EDoc](t, expected, object, msgAndArgs...) {
		return
	}

	t.FailNow()
}

// YAMLUnmarshalAsT wraps [Equal] after [yaml.Unmarshal].
//
// The input YAML may be a string or []byte.
//
// It fails if the unmarshaling returns an error or if the resulting object is not equal to the expected one.
//
// Be careful not to wrap the expected object into an "any" interface if this is not what you expected:
// the unmarshaling would take this type to unmarshal as a map[string]any.
//
// # Usage
//
//	expected := struct {
//		A int `yaml:"a"`
//	}{
//		A: 10,
//	}
//
//	assertions.YAMLUnmarshalAsT(t,expected, `{"a": 10}`)
//
// # Examples
//
//	panic: "key: value", "key: value"
//	should panic without the yaml feature enabled.
//
// Upon failure, the test [T] is marked as failed and stops execution.
func YAMLUnmarshalAsT[Object any, ADoc Text](t T, expected Object, jazon ADoc, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.YAMLUnmarshalAsT[Object, ADoc](t, expected, jazon, msgAndArgs...) {
		return
	}

	t.FailNow()
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
// Upon failure, the test [T] is marked as failed and stops execution.
func Zero(t T, i any, msgAndArgs ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.Zero(t, i, msgAndArgs...) {
		return
	}

	t.FailNow()
}
