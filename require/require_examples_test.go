// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/v2/codegen@master [sha: bb2c19fba6c03f46cb643b3bcdc1d647ea1453ab]; DO NOT EDIT.

package require_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

func ExampleCondition() {
	t := new(testing.T)
	require.Condition(t, func() bool { return true })
	fmt.Println("passed")

	// Output: passed
}

func ExampleContains() {
	t := new(testing.T)
	require.Contains(t, []string{"A", "B"}, "A")
	fmt.Println("passed")

	// Output: passed
}

func ExampleDirExists() {
	t := new(testing.T)
	require.DirExists(t, filepath.Join(testDataPath(), "existing_dir"))
	fmt.Println("passed")

	// Output: passed
}

func ExampleElementsMatch() {
	t := new(testing.T)
	require.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	fmt.Println("passed")

	// Output: passed
}

func ExampleEmpty() {
	t := new(testing.T)
	require.Empty(t, "")
	fmt.Println("passed")

	// Output: passed
}

func ExampleEqual() {
	t := new(testing.T)
	require.Equal(t, 123, 123)
	fmt.Println("passed")

	// Output: passed
}

func ExampleEqualError() {
	t := new(testing.T)
	require.EqualError(t, require.ErrTest, "assert.ErrTest general error for testing")
	fmt.Println("passed")

	// Output: passed
}

func ExampleEqualExportedValues() {
	t := new(testing.T)
	require.EqualExportedValues(t, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
	fmt.Println("passed")

	// Output: passed
}

func ExampleEqualValues() {
	t := new(testing.T)
	require.EqualValues(t, uint32(123), int32(123))
	fmt.Println("passed")

	// Output: passed
}

func ExampleError() {
	t := new(testing.T)
	require.Error(t, require.ErrTest)
	fmt.Println("passed")

	// Output: passed
}

func ExampleErrorAs() {
	t := new(testing.T)
	require.ErrorAs(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
	fmt.Println("passed")

	// Output: passed
}

func ExampleErrorContains() {
	t := new(testing.T)
	require.ErrorContains(t, require.ErrTest, "general error")
	fmt.Println("passed")

	// Output: passed
}

func ExampleErrorIs() {
	t := new(testing.T)
	require.ErrorIs(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
	fmt.Println("passed")

	// Output: passed
}

func ExampleEventually() {
	t := new(testing.T)
	require.Eventually(t, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Println("passed")

	// Output: passed
}

func ExampleEventuallyWithT() {
	t := new(testing.T)
	require.EventuallyWithT(t, func(c *require.CollectT) { require.True(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Println("passed")

	// Output: passed
}

func ExampleExactly() {
	t := new(testing.T)
	require.Exactly(t, int32(123), int32(123))
	fmt.Println("passed")

	// Output: passed
}

// func ExampleFail() {
// no success example available. Please add some examples to produce a testable example.
// }

// func ExampleFailNow() {
// no success example available. Please add some examples to produce a testable example.
// }

func ExampleFalse() {
	t := new(testing.T)
	require.False(t, 1 == 0)
	fmt.Println("passed")

	// Output: passed
}

func ExampleFileEmpty() {
	t := new(testing.T)
	require.FileEmpty(t, filepath.Join(testDataPath(), "empty_file"))
	fmt.Println("passed")

	// Output: passed
}

func ExampleFileExists() {
	t := new(testing.T)
	require.FileExists(t, filepath.Join(testDataPath(), "existing_file"))
	fmt.Println("passed")

	// Output: passed
}

func ExampleFileNotEmpty() {
	t := new(testing.T)
	require.FileNotEmpty(t, filepath.Join(testDataPath(), "existing_file"))
	fmt.Println("passed")

	// Output: passed
}

func ExampleGreater() {
	t := new(testing.T)
	require.Greater(t, 2, 1)
	fmt.Println("passed")

	// Output: passed
}

func ExampleGreaterOrEqual() {
	t := new(testing.T)
	require.GreaterOrEqual(t, 2, 1)
	fmt.Println("passed")

	// Output: passed
}

func ExampleHTTPBodyContains() {
	t := new(testing.T)
	require.HTTPBodyContains(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!")
	fmt.Println("passed")

	// Output: passed
}

func ExampleHTTPBodyNotContains() {
	t := new(testing.T)
	require.HTTPBodyNotContains(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!")
	fmt.Println("passed")

	// Output: passed
}

func ExampleHTTPError() {
	t := new(testing.T)
	require.HTTPError(t, httpError, "GET", "/", nil)
	fmt.Println("passed")

	// Output: passed
}

func ExampleHTTPRedirect() {
	t := new(testing.T)
	require.HTTPRedirect(t, httpRedirect, "GET", "/", nil)
	fmt.Println("passed")

	// Output: passed
}

func ExampleHTTPStatusCode() {
	t := new(testing.T)
	require.HTTPStatusCode(t, httpOK, "GET", "/", nil, http.StatusOK)
	fmt.Println("passed")

	// Output: passed
}

func ExampleHTTPSuccess() {
	t := new(testing.T)
	require.HTTPSuccess(t, httpOK, "GET", "/", nil)
	fmt.Println("passed")

	// Output: passed
}

func ExampleImplements() {
	t := new(testing.T)
	require.Implements(t, ptr(dummyInterface), new(testing.T))
	fmt.Println("passed")

	// Output: passed
}

func ExampleInDelta() {
	t := new(testing.T)
	require.InDelta(t, 1.0, 1.01, 0.02)
	fmt.Println("passed")

	// Output: passed
}

func ExampleInDeltaMapValues() {
	t := new(testing.T)
	require.InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
	fmt.Println("passed")

	// Output: passed
}

func ExampleInDeltaSlice() {
	t := new(testing.T)
	require.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
	fmt.Println("passed")

	// Output: passed
}

func ExampleInEpsilon() {
	t := new(testing.T)
	require.InEpsilon(t, 100.0, 101.0, 0.02)
	fmt.Println("passed")

	// Output: passed
}

func ExampleInEpsilonSlice() {
	t := new(testing.T)
	require.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
	fmt.Println("passed")

	// Output: passed
}

func ExampleIsDecreasing() {
	t := new(testing.T)
	require.IsDecreasing(t, []int{3, 2, 1})
	fmt.Println("passed")

	// Output: passed
}

func ExampleIsIncreasing() {
	t := new(testing.T)
	require.IsIncreasing(t, []int{1, 2, 3})
	fmt.Println("passed")

	// Output: passed
}

func ExampleIsNonDecreasing() {
	t := new(testing.T)
	require.IsNonDecreasing(t, []int{1, 1, 2})
	fmt.Println("passed")

	// Output: passed
}

func ExampleIsNonIncreasing() {
	t := new(testing.T)
	require.IsNonIncreasing(t, []int{2, 1, 1})
	fmt.Println("passed")

	// Output: passed
}

func ExampleIsNotType() {
	t := new(testing.T)
	require.IsNotType(t, int32(123), int64(456))
	fmt.Println("passed")

	// Output: passed
}

func ExampleIsType() {
	t := new(testing.T)
	require.IsType(t, 123, 456)
	fmt.Println("passed")

	// Output: passed
}

func ExampleJSONEq() {
	t := new(testing.T)
	require.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	fmt.Println("passed")

	// Output: passed
}

func ExampleJSONEqBytes() {
	t := new(testing.T)
	require.JSONEqBytes(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
	fmt.Println("passed")

	// Output: passed
}

func ExampleLen() {
	t := new(testing.T)
	require.Len(t, []string{"A", "B"}, 2)
	fmt.Println("passed")

	// Output: passed
}

func ExampleLess() {
	t := new(testing.T)
	require.Less(t, 1, 2)
	fmt.Println("passed")

	// Output: passed
}

func ExampleLessOrEqual() {
	t := new(testing.T)
	require.LessOrEqual(t, 1, 2)
	fmt.Println("passed")

	// Output: passed
}

func ExampleNegative() {
	t := new(testing.T)
	require.Negative(t, -1)
	fmt.Println("passed")

	// Output: passed
}

func ExampleNever() {
	t := new(testing.T)
	require.Never(t, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Println("passed")

	// Output: passed
}

func ExampleNil() {
	t := new(testing.T)
	require.Nil(t, nil)
	fmt.Println("passed")

	// Output: passed
}

func ExampleNoDirExists() {
	t := new(testing.T)
	require.NoDirExists(t, filepath.Join(testDataPath(), "non_existing_dir"))
	fmt.Println("passed")

	// Output: passed
}

func ExampleNoError() {
	t := new(testing.T)
	require.NoError(t, nil)
	fmt.Println("passed")

	// Output: passed
}

func ExampleNoFileExists() {
	t := new(testing.T)
	require.NoFileExists(t, filepath.Join(testDataPath(), "non_existing_file"))
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotContains() {
	t := new(testing.T)
	require.NotContains(t, []string{"A", "B"}, "C")
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotElementsMatch() {
	t := new(testing.T)
	require.NotElementsMatch(t, []int{1, 2, 3}, []int{1, 2, 4})
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotEmpty() {
	t := new(testing.T)
	require.NotEmpty(t, "not empty")
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotEqual() {
	t := new(testing.T)
	require.NotEqual(t, 123, 456)
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotEqualValues() {
	t := new(testing.T)
	require.NotEqualValues(t, uint32(123), int32(456))
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotErrorAs() {
	t := new(testing.T)
	require.NotErrorAs(t, require.ErrTest, new(*dummyError))
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotErrorIs() {
	t := new(testing.T)
	require.NotErrorIs(t, require.ErrTest, io.EOF)
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotImplements() {
	t := new(testing.T)
	require.NotImplements(t, (*error)(nil), new(testing.T))
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotNil() {
	t := new(testing.T)
	require.NotNil(t, "not nil")
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotPanics() {
	t := new(testing.T)
	require.NotPanics(t, func() {})
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotRegexp() {
	t := new(testing.T)
	require.NotRegexp(t, "^start", "not starting")
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotSame() {
	t := new(testing.T)
	require.NotSame(t, &staticVar, ptr("static string"))
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotSubset() {
	t := new(testing.T)
	require.NotSubset(t, []int{1, 2, 3}, []int{4, 5})
	fmt.Println("passed")

	// Output: passed
}

func ExampleNotZero() {
	t := new(testing.T)
	require.NotZero(t, 1)
	fmt.Println("passed")

	// Output: passed
}

func ExamplePanics() {
	t := new(testing.T)
	require.Panics(t, func() { panic("panicking") })
	fmt.Println("passed")

	// Output: passed
}

func ExamplePanicsWithError() {
	t := new(testing.T)
	require.PanicsWithError(t, require.ErrTest.Error(), func() { panic(require.ErrTest) })
	fmt.Println("passed")

	// Output: passed
}

func ExamplePanicsWithValue() {
	t := new(testing.T)
	require.PanicsWithValue(t, "panicking", func() { panic("panicking") })
	fmt.Println("passed")

	// Output: passed
}

func ExamplePositive() {
	t := new(testing.T)
	require.Positive(t, 1)
	fmt.Println("passed")

	// Output: passed
}

func ExampleRegexp() {
	t := new(testing.T)
	require.Regexp(t, "^start", "starting")
	fmt.Println("passed")

	// Output: passed
}

func ExampleSame() {
	t := new(testing.T)
	require.Same(t, &staticVar, staticVarPtr)
	fmt.Println("passed")

	// Output: passed
}

func ExampleSubset() {
	t := new(testing.T)
	require.Subset(t, []int{1, 2, 3}, []int{1, 2})
	fmt.Println("passed")

	// Output: passed
}

func ExampleTrue() {
	t := new(testing.T)
	require.True(t, 1 == 1)
	fmt.Println("passed")

	// Output: passed
}

func ExampleWithinDuration() {
	t := new(testing.T)
	require.WithinDuration(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second)
	fmt.Println("passed")

	// Output: passed
}

func ExampleWithinRange() {
	t := new(testing.T)
	require.WithinRange(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
	fmt.Println("passed")

	// Output: passed
}

// func ExampleYAMLEq() {
// no success example available. Please add some examples to produce a testable example.
// }

func ExampleZero() {
	t := new(testing.T)
	require.Zero(t, 0)
	fmt.Println("passed")

	// Output: passed
}

// Test helpers (also in the tests for package require
//
// This code is duplicated because the current test is run as a separate test package: require_test

func testDataPath() string {
	return filepath.Join("..", "internal", "assertions", "testdata")
}

func httpOK(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func httpError(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func httpRedirect(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func httpBody(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	_, _ = fmt.Fprintf(w, "Hello, %s!", name)
}

//nolint:gochecknoglobals // this is on purpose to share a common pointer when testing
var (
	staticVar      = "static string"
	staticVarPtr   = &staticVar
	dummyInterface require.T
)

func ptr[T any](value T) *T {
	p := value

	return &p
}

type dummyStruct struct {
	A string
	b int
}

type dummyError struct {
}

func (d *dummyError) Error() string {
	return "dummy error"
}
