// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

package assert_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/assert"
)

func ExampleCondition() {
	t := new(testing.T) // should come from testing, e.g. func TestCondition(t *testing.T)
	success := assert.Condition(t, func() bool {
		return true
	})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleConsistently() {
	t := new(testing.T) // should come from testing, e.g. func TestConsistently(t *testing.T)
	success := assert.Consistently(t, func() bool {
		return true
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleContains() {
	t := new(testing.T) // should come from testing, e.g. func TestContains(t *testing.T)
	success := assert.Contains(t, []string{"A", "B"}, "A")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleDirExists() {
	t := new(testing.T) // should come from testing, e.g. func TestDirExists(t *testing.T)
	success := assert.DirExists(t, filepath.Join(testDataPath(), "existing_dir"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleDirNotExists() {
	t := new(testing.T) // should come from testing, e.g. func TestDirNotExists(t *testing.T)
	success := assert.DirNotExists(t, filepath.Join(testDataPath(), "non_existing_dir"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleElementsMatch() {
	t := new(testing.T) // should come from testing, e.g. func TestElementsMatch(t *testing.T)
	success := assert.ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleElementsMatchT() {
	t := new(testing.T) // should come from testing, e.g. func TestElementsMatchT(t *testing.T)
	success := assert.ElementsMatchT(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEmpty() {
	t := new(testing.T) // should come from testing, e.g. func TestEmpty(t *testing.T)
	success := assert.Empty(t, "")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEqual() {
	t := new(testing.T) // should come from testing, e.g. func TestEqual(t *testing.T)
	success := assert.Equal(t, 123, 123)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEqualError() {
	t := new(testing.T) // should come from testing, e.g. func TestEqualError(t *testing.T)
	success := assert.EqualError(t, assert.ErrTest, "assert.ErrTest general error for testing")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEqualExportedValues() {
	t := new(testing.T) // should come from testing, e.g. func TestEqualExportedValues(t *testing.T)
	success := assert.EqualExportedValues(t, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEqualT() {
	t := new(testing.T) // should come from testing, e.g. func TestEqualT(t *testing.T)
	success := assert.EqualT(t, 123, 123)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEqualValues() {
	t := new(testing.T) // should come from testing, e.g. func TestEqualValues(t *testing.T)
	success := assert.EqualValues(t, uint32(123), int32(123))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleError() {
	t := new(testing.T) // should come from testing, e.g. func TestError(t *testing.T)
	success := assert.Error(t, assert.ErrTest)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleErrorAs() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorAs(t *testing.T)
	success := assert.ErrorAs(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleErrorContains() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorContains(t *testing.T)
	success := assert.ErrorContains(t, assert.ErrTest, "general error")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleErrorIs() {
	t := new(testing.T) // should come from testing, e.g. func TestErrorIs(t *testing.T)
	success := assert.ErrorIs(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEventually() {
	t := new(testing.T) // should come from testing, e.g. func TestEventually(t *testing.T)
	success := assert.Eventually(t, func() bool {
		return true
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleEventuallyWith() {
	t := new(testing.T) // should come from testing, e.g. func TestEventuallyWith(t *testing.T)
	success := assert.EventuallyWith(t, func(c *assert.CollectT) {
		assert.True(c, true)
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleExactly() {
	t := new(testing.T) // should come from testing, e.g. func TestExactly(t *testing.T)
	success := assert.Exactly(t, int32(123), int32(123))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

// func ExampleFail() {
// no success example available. Please add some examples to produce a testable example.
// }

// func ExampleFailNow() {
// no success example available. Please add some examples to produce a testable example.
// }

func ExampleFalse() {
	t := new(testing.T) // should come from testing, e.g. func TestFalse(t *testing.T)
	success := assert.False(t, 1 == 0)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleFalseT() {
	t := new(testing.T) // should come from testing, e.g. func TestFalseT(t *testing.T)
	success := assert.FalseT(t, 1 == 0)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleFileEmpty() {
	t := new(testing.T) // should come from testing, e.g. func TestFileEmpty(t *testing.T)
	success := assert.FileEmpty(t, filepath.Join(testDataPath(), "empty_file"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleFileExists() {
	t := new(testing.T) // should come from testing, e.g. func TestFileExists(t *testing.T)
	success := assert.FileExists(t, filepath.Join(testDataPath(), "existing_file"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleFileNotEmpty() {
	t := new(testing.T) // should come from testing, e.g. func TestFileNotEmpty(t *testing.T)
	success := assert.FileNotEmpty(t, filepath.Join(testDataPath(), "existing_file"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleFileNotExists() {
	t := new(testing.T) // should come from testing, e.g. func TestFileNotExists(t *testing.T)
	success := assert.FileNotExists(t, filepath.Join(testDataPath(), "non_existing_file"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleGreater() {
	t := new(testing.T) // should come from testing, e.g. func TestGreater(t *testing.T)
	success := assert.Greater(t, 2, 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleGreaterOrEqual() {
	t := new(testing.T) // should come from testing, e.g. func TestGreaterOrEqual(t *testing.T)
	success := assert.GreaterOrEqual(t, 2, 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleGreaterOrEqualT() {
	t := new(testing.T) // should come from testing, e.g. func TestGreaterOrEqualT(t *testing.T)
	success := assert.GreaterOrEqualT(t, 2, 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleGreaterT() {
	t := new(testing.T) // should come from testing, e.g. func TestGreaterT(t *testing.T)
	success := assert.GreaterT(t, 2, 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleHTTPBodyContains() {
	t := new(testing.T) // should come from testing, e.g. func TestHTTPBodyContains(t *testing.T)
	success := assert.HTTPBodyContains(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleHTTPBodyNotContains() {
	t := new(testing.T) // should come from testing, e.g. func TestHTTPBodyNotContains(t *testing.T)
	success := assert.HTTPBodyNotContains(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleHTTPError() {
	t := new(testing.T) // should come from testing, e.g. func TestHTTPError(t *testing.T)
	success := assert.HTTPError(t, httpError, "GET", "/", nil)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleHTTPRedirect() {
	t := new(testing.T) // should come from testing, e.g. func TestHTTPRedirect(t *testing.T)
	success := assert.HTTPRedirect(t, httpRedirect, "GET", "/", nil)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleHTTPStatusCode() {
	t := new(testing.T) // should come from testing, e.g. func TestHTTPStatusCode(t *testing.T)
	success := assert.HTTPStatusCode(t, httpOK, "GET", "/", nil, http.StatusOK)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleHTTPSuccess() {
	t := new(testing.T) // should come from testing, e.g. func TestHTTPSuccess(t *testing.T)
	success := assert.HTTPSuccess(t, httpOK, "GET", "/", nil)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleImplements() {
	t := new(testing.T) // should come from testing, e.g. func TestImplements(t *testing.T)
	success := assert.Implements(t, ptr(dummyInterface), new(testing.T))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleInDelta() {
	t := new(testing.T) // should come from testing, e.g. func TestInDelta(t *testing.T)
	success := assert.InDelta(t, 1.0, 1.01, 0.02)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleInDeltaMapValues() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaMapValues(t *testing.T)
	success := assert.InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleInDeltaSlice() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaSlice(t *testing.T)
	success := assert.InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleInDeltaT() {
	t := new(testing.T) // should come from testing, e.g. func TestInDeltaT(t *testing.T)
	success := assert.InDeltaT(t, 1.0, 1.01, 0.02)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleInEpsilon() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilon(t *testing.T)
	success := assert.InEpsilon(t, 100.0, 101.0, 0.02)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleInEpsilonSlice() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilonSlice(t *testing.T)
	success := assert.InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleInEpsilonT() {
	t := new(testing.T) // should come from testing, e.g. func TestInEpsilonT(t *testing.T)
	success := assert.InEpsilonT(t, 100.0, 101.0, 0.02)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsDecreasing() {
	t := new(testing.T) // should come from testing, e.g. func TestIsDecreasing(t *testing.T)
	success := assert.IsDecreasing(t, []int{3, 2, 1})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsDecreasingT() {
	t := new(testing.T) // should come from testing, e.g. func TestIsDecreasingT(t *testing.T)
	success := assert.IsDecreasingT(t, []int{3, 2, 1})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsIncreasing() {
	t := new(testing.T) // should come from testing, e.g. func TestIsIncreasing(t *testing.T)
	success := assert.IsIncreasing(t, []int{1, 2, 3})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsIncreasingT() {
	t := new(testing.T) // should come from testing, e.g. func TestIsIncreasingT(t *testing.T)
	success := assert.IsIncreasingT(t, []int{1, 2, 3})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsNonDecreasing() {
	t := new(testing.T) // should come from testing, e.g. func TestIsNonDecreasing(t *testing.T)
	success := assert.IsNonDecreasing(t, []int{1, 1, 2})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsNonDecreasingT() {
	t := new(testing.T) // should come from testing, e.g. func TestIsNonDecreasingT(t *testing.T)
	success := assert.IsNonDecreasingT(t, []int{1, 1, 2})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsNonIncreasing() {
	t := new(testing.T) // should come from testing, e.g. func TestIsNonIncreasing(t *testing.T)
	success := assert.IsNonIncreasing(t, []int{2, 1, 1})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsNonIncreasingT() {
	t := new(testing.T) // should come from testing, e.g. func TestIsNonIncreasingT(t *testing.T)
	success := assert.IsNonIncreasingT(t, []int{2, 1, 1})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsNotOfTypeT() {
	t := new(testing.T) // should come from testing, e.g. func TestIsNotOfTypeT(t *testing.T)
	success := assert.IsNotOfTypeT[myType](t, 123.123)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsNotType() {
	t := new(testing.T) // should come from testing, e.g. func TestIsNotType(t *testing.T)
	success := assert.IsNotType(t, int32(123), int64(456))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsOfTypeT() {
	t := new(testing.T) // should come from testing, e.g. func TestIsOfTypeT(t *testing.T)
	success := assert.IsOfTypeT[myType](t, myType(123.123))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleIsType() {
	t := new(testing.T) // should come from testing, e.g. func TestIsType(t *testing.T)
	success := assert.IsType(t, 123, 456)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleJSONEq() {
	t := new(testing.T) // should come from testing, e.g. func TestJSONEq(t *testing.T)
	success := assert.JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleJSONEqBytes() {
	t := new(testing.T) // should come from testing, e.g. func TestJSONEqBytes(t *testing.T)
	success := assert.JSONEqBytes(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleJSONEqT() {
	t := new(testing.T) // should come from testing, e.g. func TestJSONEqT(t *testing.T)
	success := assert.JSONEqT(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleJSONMarshalAsT() {
	t := new(testing.T) // should come from testing, e.g. func TestJSONMarshalAsT(t *testing.T)
	success := assert.JSONMarshalAsT(t, []byte(`{"A": "a"}`), dummyStruct{A: "a"})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleJSONUnmarshalAsT() {
	t := new(testing.T) // should come from testing, e.g. func TestJSONUnmarshalAsT(t *testing.T)
	success := assert.JSONUnmarshalAsT(t, dummyStruct{A: "a"}, []byte(`{"A": "a"}`))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleKind() {
	t := new(testing.T) // should come from testing, e.g. func TestKind(t *testing.T)
	success := assert.Kind(t, reflect.String, "hello")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleLen() {
	t := new(testing.T) // should come from testing, e.g. func TestLen(t *testing.T)
	success := assert.Len(t, []string{"A", "B"}, 2)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleLess() {
	t := new(testing.T) // should come from testing, e.g. func TestLess(t *testing.T)
	success := assert.Less(t, 1, 2)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleLessOrEqual() {
	t := new(testing.T) // should come from testing, e.g. func TestLessOrEqual(t *testing.T)
	success := assert.LessOrEqual(t, 1, 2)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleLessOrEqualT() {
	t := new(testing.T) // should come from testing, e.g. func TestLessOrEqualT(t *testing.T)
	success := assert.LessOrEqualT(t, 1, 2)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleLessT() {
	t := new(testing.T) // should come from testing, e.g. func TestLessT(t *testing.T)
	success := assert.LessT(t, 1, 2)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleMapContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestMapContainsT(t *testing.T)
	success := assert.MapContainsT(t, map[string]string{"A": "B"}, "A")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleMapNotContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestMapNotContainsT(t *testing.T)
	success := assert.MapNotContainsT(t, map[string]string{"A": "B"}, "C")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNegative() {
	t := new(testing.T) // should come from testing, e.g. func TestNegative(t *testing.T)
	success := assert.Negative(t, -1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNegativeT() {
	t := new(testing.T) // should come from testing, e.g. func TestNegativeT(t *testing.T)
	success := assert.NegativeT(t, -1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNever() {
	t := new(testing.T) // should come from testing, e.g. func TestNever(t *testing.T)
	success := assert.Never(t, func() bool {
		return false
	}, 100*time.Millisecond, 20*time.Millisecond)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNil() {
	t := new(testing.T) // should come from testing, e.g. func TestNil(t *testing.T)
	success := assert.Nil(t, nil)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNoError() {
	t := new(testing.T) // should come from testing, e.g. func TestNoError(t *testing.T)
	success := assert.NoError(t, nil)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNoFileDescriptorLeak() {
	if runtime.GOOS != "linux" {
		// This example is only runnable on linux. On other platforms, the assertion skips the test.
		// We force the expected output below, so that tests don't fail on other platforms.
		fmt.Println("success: true")

		return
	}

	t := new(testing.T) // should come from testing, e.g. func TestNoFileDescriptorLeak(t *testing.T)
	success := assert.NoFileDescriptorLeak(t, func() {
	})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNoGoRoutineLeak() {
	t := new(testing.T) // should come from testing, e.g. func TestNoGoRoutineLeak(t *testing.T)
	success := assert.NoGoRoutineLeak(t, func() {
	})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotContains() {
	t := new(testing.T) // should come from testing, e.g. func TestNotContains(t *testing.T)
	success := assert.NotContains(t, []string{"A", "B"}, "C")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotElementsMatch() {
	t := new(testing.T) // should come from testing, e.g. func TestNotElementsMatch(t *testing.T)
	success := assert.NotElementsMatch(t, []int{1, 2, 3}, []int{1, 2, 4})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotElementsMatchT() {
	t := new(testing.T) // should come from testing, e.g. func TestNotElementsMatchT(t *testing.T)
	success := assert.NotElementsMatchT(t, []int{1, 2, 3}, []int{1, 2, 4})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotEmpty() {
	t := new(testing.T) // should come from testing, e.g. func TestNotEmpty(t *testing.T)
	success := assert.NotEmpty(t, "not empty")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotEqual() {
	t := new(testing.T) // should come from testing, e.g. func TestNotEqual(t *testing.T)
	success := assert.NotEqual(t, 123, 456)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotEqualT() {
	t := new(testing.T) // should come from testing, e.g. func TestNotEqualT(t *testing.T)
	success := assert.NotEqualT(t, 123, 456)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotEqualValues() {
	t := new(testing.T) // should come from testing, e.g. func TestNotEqualValues(t *testing.T)
	success := assert.NotEqualValues(t, uint32(123), int32(456))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotErrorAs() {
	t := new(testing.T) // should come from testing, e.g. func TestNotErrorAs(t *testing.T)
	success := assert.NotErrorAs(t, assert.ErrTest, new(*dummyError))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotErrorIs() {
	t := new(testing.T) // should come from testing, e.g. func TestNotErrorIs(t *testing.T)
	success := assert.NotErrorIs(t, assert.ErrTest, io.EOF)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotImplements() {
	t := new(testing.T) // should come from testing, e.g. func TestNotImplements(t *testing.T)
	success := assert.NotImplements(t, (*error)(nil), new(testing.T))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotKind() {
	t := new(testing.T) // should come from testing, e.g. func TestNotKind(t *testing.T)
	success := assert.NotKind(t, reflect.String, 0)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotNil() {
	t := new(testing.T) // should come from testing, e.g. func TestNotNil(t *testing.T)
	success := assert.NotNil(t, "not nil")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotPanics() {
	t := new(testing.T) // should come from testing, e.g. func TestNotPanics(t *testing.T)
	success := assert.NotPanics(t, func() {
	})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotRegexp() {
	t := new(testing.T) // should come from testing, e.g. func TestNotRegexp(t *testing.T)
	success := assert.NotRegexp(t, "^start", "not starting")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotRegexpT() {
	t := new(testing.T) // should come from testing, e.g. func TestNotRegexpT(t *testing.T)
	success := assert.NotRegexpT(t, "^start", "not starting")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotSame() {
	t := new(testing.T) // should come from testing, e.g. func TestNotSame(t *testing.T)
	success := assert.NotSame(t, &staticVar, ptr("static string"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotSameT() {
	t := new(testing.T) // should come from testing, e.g. func TestNotSameT(t *testing.T)
	success := assert.NotSameT(t, &staticVar, ptr("static string"))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotSortedT() {
	t := new(testing.T) // should come from testing, e.g. func TestNotSortedT(t *testing.T)
	success := assert.NotSortedT(t, []int{3, 1, 3})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotSubset() {
	t := new(testing.T) // should come from testing, e.g. func TestNotSubset(t *testing.T)
	success := assert.NotSubset(t, []int{1, 2, 3}, []int{4, 5})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleNotZero() {
	t := new(testing.T) // should come from testing, e.g. func TestNotZero(t *testing.T)
	success := assert.NotZero(t, 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExamplePanics() {
	t := new(testing.T) // should come from testing, e.g. func TestPanics(t *testing.T)
	success := assert.Panics(t, func() {
		panic("panicking")
	})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExamplePanicsWithError() {
	t := new(testing.T) // should come from testing, e.g. func TestPanicsWithError(t *testing.T)
	success := assert.PanicsWithError(t, assert.ErrTest.Error(), func() {
		panic(assert.ErrTest)
	})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExamplePanicsWithValue() {
	t := new(testing.T) // should come from testing, e.g. func TestPanicsWithValue(t *testing.T)
	success := assert.PanicsWithValue(t, "panicking", func() {
		panic("panicking")
	})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExamplePositive() {
	t := new(testing.T) // should come from testing, e.g. func TestPositive(t *testing.T)
	success := assert.Positive(t, 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExamplePositiveT() {
	t := new(testing.T) // should come from testing, e.g. func TestPositiveT(t *testing.T)
	success := assert.PositiveT(t, 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleRegexp() {
	t := new(testing.T) // should come from testing, e.g. func TestRegexp(t *testing.T)
	success := assert.Regexp(t, "^start", "starting")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleRegexpT() {
	t := new(testing.T) // should come from testing, e.g. func TestRegexpT(t *testing.T)
	success := assert.RegexpT(t, "^start", "starting")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSame() {
	t := new(testing.T) // should come from testing, e.g. func TestSame(t *testing.T)
	success := assert.Same(t, &staticVar, staticVarPtr)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSameT() {
	t := new(testing.T) // should come from testing, e.g. func TestSameT(t *testing.T)
	success := assert.SameT(t, &staticVar, staticVarPtr)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSeqContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestSeqContainsT(t *testing.T)
	success := assert.SeqContainsT(t, slices.Values([]string{"A", "B"}), "A")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSeqNotContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestSeqNotContainsT(t *testing.T)
	success := assert.SeqNotContainsT(t, slices.Values([]string{"A", "B"}), "C")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSliceContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceContainsT(t *testing.T)
	success := assert.SliceContainsT(t, []string{"A", "B"}, "A")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSliceNotContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceNotContainsT(t *testing.T)
	success := assert.SliceNotContainsT(t, []string{"A", "B"}, "C")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSliceNotSubsetT() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceNotSubsetT(t *testing.T)
	success := assert.SliceNotSubsetT(t, []int{1, 2, 3}, []int{4, 5})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSliceSubsetT() {
	t := new(testing.T) // should come from testing, e.g. func TestSliceSubsetT(t *testing.T)
	success := assert.SliceSubsetT(t, []int{1, 2, 3}, []int{1, 2})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSortedT() {
	t := new(testing.T) // should come from testing, e.g. func TestSortedT(t *testing.T)
	success := assert.SortedT(t, []int{1, 1, 3})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleStringContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestStringContainsT(t *testing.T)
	success := assert.StringContainsT(t, "AB", "A")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleStringNotContainsT() {
	t := new(testing.T) // should come from testing, e.g. func TestStringNotContainsT(t *testing.T)
	success := assert.StringNotContainsT(t, "AB", "C")
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleSubset() {
	t := new(testing.T) // should come from testing, e.g. func TestSubset(t *testing.T)
	success := assert.Subset(t, []int{1, 2, 3}, []int{1, 2})
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleTrue() {
	t := new(testing.T) // should come from testing, e.g. func TestTrue(t *testing.T)
	success := assert.True(t, 1 == 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleTrueT() {
	t := new(testing.T) // should come from testing, e.g. func TestTrueT(t *testing.T)
	success := assert.TrueT(t, 1 == 1)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleWithinDuration() {
	t := new(testing.T) // should come from testing, e.g. func TestWithinDuration(t *testing.T)
	success := assert.WithinDuration(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

func ExampleWithinRange() {
	t := new(testing.T) // should come from testing, e.g. func TestWithinRange(t *testing.T)
	success := assert.WithinRange(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

// func ExampleYAMLEq() {
// no success example available. Please add some examples to produce a testable example.
// }

// func ExampleYAMLEqBytes() {
// no success example available. Please add some examples to produce a testable example.
// }

// func ExampleYAMLEqT() {
// no success example available. Please add some examples to produce a testable example.
// }

// func ExampleYAMLMarshalAsT() {
// no success example available. Please add some examples to produce a testable example.
// }

// func ExampleYAMLUnmarshalAsT() {
// no success example available. Please add some examples to produce a testable example.
// }

func ExampleZero() {
	t := new(testing.T) // should come from testing, e.g. func TestZero(t *testing.T)
	success := assert.Zero(t, 0)
	fmt.Printf("success: %t\n", success)

	// Output: success: true
}

// Test helpers (also in the tests for package assert.
//
// This code is duplicated because the current test is run as a separate test package: assert_test.

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
	dummyInterface assert.T
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

type myType float64
