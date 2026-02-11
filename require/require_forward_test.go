// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

package require

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func TestAssertionsCondition(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Condition(func() bool { return true })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Condition(func() bool { return false })
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Condition should call FailNow()")
		}
	})
}

func TestAssertionsContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Contains([]string{"A", "B"}, "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Contains([]string{"A", "B"}, "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Contains should call FailNow()")
		}
	})
}

func TestAssertionsDirExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirExists(filepath.Join(testDataPath(), "existing_dir"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirExists(filepath.Join(testDataPath(), "non_existing_dir"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.DirExists should call FailNow()")
		}
	})
}

func TestAssertionsDirNotExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirNotExists(filepath.Join(testDataPath(), "non_existing_dir"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirNotExists(filepath.Join(testDataPath(), "existing_dir"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.DirNotExists should call FailNow()")
		}
	})
}

func TestAssertionsElementsMatch(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatch([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatch([]int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ElementsMatch should call FailNow()")
		}
	})
}

func TestAssertionsEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Empty("")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Empty("not empty")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Empty should call FailNow()")
		}
	})
}

func TestAssertionsEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Equal(123, 123)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Equal(123, 456)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Equal should call FailNow()")
		}
	})
}

func TestAssertionsEqualError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualError(ErrTest, "assert.ErrTest general error for testing")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualError(ErrTest, "wrong error message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualError should call FailNow()")
		}
	})
}

func TestAssertionsEqualExportedValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualExportedValues(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualExportedValues(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualExportedValues should call FailNow()")
		}
	})
}

func TestAssertionsEqualValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualValues(uint32(123), int32(123))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualValues(uint32(123), int32(456))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualValues should call FailNow()")
		}
	})
}

func TestAssertionsError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Error(ErrTest)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Error(nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Error should call FailNow()")
		}
	})
}

func TestAssertionsErrorAs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAs(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAs(ErrTest, new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorAs should call FailNow()")
		}
	})
}

func TestAssertionsErrorContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorContains(ErrTest, "general error")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorContains(ErrTest, "not in message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorContains should call FailNow()")
		}
	})
}

func TestAssertionsErrorIs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorIs(fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorIs(ErrTest, io.EOF)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorIs should call FailNow()")
		}
	})
}

func TestAssertionsEventually(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Eventually(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Eventually(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Eventually should call FailNow()")
		}
	})
}

func TestAssertionsEventuallyWith(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EventuallyWith(func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EventuallyWith(func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EventuallyWith should call FailNow()")
		}
	})
}

func TestAssertionsExactly(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Exactly(int32(123), int32(123))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Exactly(int32(123), int64(123))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Exactly should call FailNow()")
		}
	})
}

func TestAssertionsFail(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Fail("failed")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Fail should call FailNow()")
		}
	})
}

func TestAssertionsFailNow(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FailNow("failed")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FailNow should call FailNow()")
		}
	})
}

func TestAssertionsFalse(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.False(1 == 0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.False(1 == 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.False should call FailNow()")
		}
	})
}

func TestAssertionsFileEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileEmpty(filepath.Join(testDataPath(), "empty_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileEmpty(filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileEmpty should call FailNow()")
		}
	})
}

func TestAssertionsFileExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileExists(filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileExists(filepath.Join(testDataPath(), "non_existing_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileExists should call FailNow()")
		}
	})
}

func TestAssertionsFileNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotEmpty(filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotEmpty(filepath.Join(testDataPath(), "empty_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileNotEmpty should call FailNow()")
		}
	})
}

func TestAssertionsFileNotExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotExists(filepath.Join(testDataPath(), "non_existing_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotExists(filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileNotExists should call FailNow()")
		}
	})
}

func TestAssertionsGreater(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Greater(2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Greater(1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Greater should call FailNow()")
		}
	})
}

func TestAssertionsGreaterOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqual(2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqual(1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqual should call FailNow()")
		}
	})
}

func TestAssertionsHTTPBodyContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyContains(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyContains(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPBodyContains should call FailNow()")
		}
	})
}

func TestAssertionsHTTPBodyNotContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyNotContains(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyNotContains(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPBodyNotContains should call FailNow()")
		}
	})
}

func TestAssertionsHTTPError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPError(httpError, "GET", "/", nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPError(httpOK, "GET", "/", nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPError should call FailNow()")
		}
	})
}

func TestAssertionsHTTPRedirect(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPRedirect(httpRedirect, "GET", "/", nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPRedirect(httpError, "GET", "/", nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPRedirect should call FailNow()")
		}
	})
}

func TestAssertionsHTTPStatusCode(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPStatusCode(httpOK, "GET", "/", nil, http.StatusOK)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPStatusCode(httpError, "GET", "/", nil, http.StatusOK)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPStatusCode should call FailNow()")
		}
	})
}

func TestAssertionsHTTPSuccess(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPSuccess(httpOK, "GET", "/", nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPSuccess(httpError, "GET", "/", nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPSuccess should call FailNow()")
		}
	})
}

func TestAssertionsImplements(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Implements(ptr(dummyInterface), new(testing.T))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Implements((*error)(nil), new(testing.T))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Implements should call FailNow()")
		}
	})
}

func TestAssertionsInDelta(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDelta(1.0, 1.01, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDelta(1.0, 1.1, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDelta should call FailNow()")
		}
	})
}

func TestAssertionsInDeltaMapValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaMapValues(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaMapValues(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDeltaMapValues should call FailNow()")
		}
	})
}

func TestAssertionsInDeltaSlice(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaSlice([]float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaSlice([]float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDeltaSlice should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilon(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilon(100.0, 101.0, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilon(100.0, 110.0, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilon should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilonSlice(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSlice([]float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSlice([]float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilonSlice should call FailNow()")
		}
	})
}

func TestAssertionsIsDecreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasing([]int{3, 2, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasing([]int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsDecreasing should call FailNow()")
		}
	})
}

func TestAssertionsIsIncreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasing([]int{1, 2, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasing([]int{1, 1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsIncreasing should call FailNow()")
		}
	})
}

func TestAssertionsIsNonDecreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasing([]int{1, 1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasing([]int{2, 1, 0})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasing should call FailNow()")
		}
	})
}

func TestAssertionsIsNonIncreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasing([]int{2, 1, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasing([]int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasing should call FailNow()")
		}
	})
}

func TestAssertionsIsNotType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotType(int32(123), int64(456))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotType(123, 456)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNotType should call FailNow()")
		}
	})
}

func TestAssertionsIsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsType(123, 456)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsType(int32(123), int64(456))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsType should call FailNow()")
		}
	})
}

func TestAssertionsJSONEq(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEq(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONEq should call FailNow()")
		}
	})
}

func TestAssertionsJSONEqBytes(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqBytes([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqBytes([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONEqBytes should call FailNow()")
		}
	})
}

func TestAssertionsKind(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Kind(reflect.String, "hello")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Kind(reflect.String, 0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Kind should call FailNow()")
		}
	})
}

func TestAssertionsLen(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Len([]string{"A", "B"}, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Len([]string{"A", "B"}, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Len should call FailNow()")
		}
	})
}

func TestAssertionsLess(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Less(1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Less(2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Less should call FailNow()")
		}
	})
}

func TestAssertionsLessOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqual(1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqual(2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.LessOrEqual should call FailNow()")
		}
	})
}

func TestAssertionsNegative(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Negative(-1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Negative(1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Negative should call FailNow()")
		}
	})
}

func TestAssertionsNever(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Never(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Never(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Never should call FailNow()")
		}
	})
}

func TestAssertionsNil(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Nil(nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Nil("not nil")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Nil should call FailNow()")
		}
	})
}

func TestAssertionsNoError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoError(nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoError(ErrTest)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NoError should call FailNow()")
		}
	})
}

func TestAssertionsNoFileDescriptorLeak(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoFileDescriptorLeak(func() {})
		// require functions don't return a value
	})
}

func TestAssertionsNoGoRoutineLeak(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoGoRoutineLeak(func() {})
		// require functions don't return a value
	})
}

func TestAssertionsNotContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotContains([]string{"A", "B"}, "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotContains([]string{"A", "B"}, "B")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotContains should call FailNow()")
		}
	})
}

func TestAssertionsNotElementsMatch(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatch([]int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatch([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotElementsMatch should call FailNow()")
		}
	})
}

func TestAssertionsNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEmpty("not empty")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEmpty("")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEmpty should call FailNow()")
		}
	})
}

func TestAssertionsNotEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqual(123, 456)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqual(123, 123)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEqual should call FailNow()")
		}
	})
}

func TestAssertionsNotEqualValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualValues(uint32(123), int32(456))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualValues(uint32(123), int32(123))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEqualValues should call FailNow()")
		}
	})
}

func TestAssertionsNotErrorAs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAs(ErrTest, new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAs(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotErrorAs should call FailNow()")
		}
	})
}

func TestAssertionsNotErrorIs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorIs(ErrTest, io.EOF)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorIs(fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotErrorIs should call FailNow()")
		}
	})
}

func TestAssertionsNotImplements(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotImplements((*error)(nil), new(testing.T))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotImplements(ptr(dummyInterface), new(testing.T))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotImplements should call FailNow()")
		}
	})
}

func TestAssertionsNotKind(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotKind(reflect.String, 0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotKind(reflect.String, "hello")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotKind should call FailNow()")
		}
	})
}

func TestAssertionsNotNil(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotNil("not nil")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotNil(nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotNil should call FailNow()")
		}
	})
}

func TestAssertionsNotPanics(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotPanics(func() {})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotPanics(func() { panic("panicking") })
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotPanics should call FailNow()")
		}
	})
}

func TestAssertionsNotRegexp(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexp("^start", "not starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexp("^start", "starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotRegexp should call FailNow()")
		}
	})
}

func TestAssertionsNotSame(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSame(&staticVar, ptr("static string"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSame(&staticVar, staticVarPtr)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSame should call FailNow()")
		}
	})
}

func TestAssertionsNotSubset(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSubset([]int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSubset([]int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSubset should call FailNow()")
		}
	})
}

func TestAssertionsNotZero(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotZero(1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotZero(0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotZero should call FailNow()")
		}
	})
}

func TestAssertionsPanics(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() { panic("panicking") })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Panics should call FailNow()")
		}
	})
}

func TestAssertionsPanicsWithError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithError(ErrTest.Error(), func() { panic(ErrTest) })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithError(ErrTest.Error(), func() {})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.PanicsWithError should call FailNow()")
		}
	})
}

func TestAssertionsPanicsWithValue(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithValue("panicking", func() { panic("panicking") })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithValue("panicking", func() {})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.PanicsWithValue should call FailNow()")
		}
	})
}

func TestAssertionsPositive(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Positive(1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Positive(-1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Positive should call FailNow()")
		}
	})
}

func TestAssertionsRegexp(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Regexp("^start", "starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Regexp("^start", "not starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Regexp should call FailNow()")
		}
	})
}

func TestAssertionsSame(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Same(&staticVar, staticVarPtr)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Same(&staticVar, ptr("static string"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Same should call FailNow()")
		}
	})
}

func TestAssertionsSubset(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Subset([]int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Subset([]int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Subset should call FailNow()")
		}
	})
}

func TestAssertionsTrue(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.True(1 == 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.True(1 == 0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.True should call FailNow()")
		}
	})
}

func TestAssertionsWithinDuration(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinDuration(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinDuration(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.WithinDuration should call FailNow()")
		}
	})
}

func TestAssertionsWithinRange(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinRange(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinRange(time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.WithinRange should call FailNow()")
		}
	})
}

func TestAssertionsYAMLEq(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLEq("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLEq should panic as expected")
		}
	})
}

func TestAssertionsYAMLEqBytes(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLEqBytes([]byte("key: value"), []byte("key: value"))
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLEqBytes should panic as expected")
		}
	})
}

func TestAssertionsZero(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Zero(0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Zero(1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Zero should call FailNow()")
		}
	})
}

func TestAssertionsConditionf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Conditionf(func() bool { return true }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Conditionf(func() bool { return false }, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Conditionf should call FailNow()")
		}
	})
}

func TestAssertionsContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Containsf([]string{"A", "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Containsf([]string{"A", "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Containsf should call FailNow()")
		}
	})
}

func TestAssertionsDirExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirExistsf(filepath.Join(testDataPath(), "existing_dir"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirExistsf(filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.DirExistsf should call FailNow()")
		}
	})
}

func TestAssertionsDirNotExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirNotExistsf(filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.DirNotExistsf(filepath.Join(testDataPath(), "existing_dir"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.DirNotExistsf should call FailNow()")
		}
	})
}

func TestAssertionsElementsMatchf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatchf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatchf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ElementsMatchf should call FailNow()")
		}
	})
}

func TestAssertionsEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Emptyf("", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Emptyf("not empty", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Emptyf should call FailNow()")
		}
	})
}

func TestAssertionsEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Equalf(123, 123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Equalf(123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Equalf should call FailNow()")
		}
	})
}

func TestAssertionsEqualErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualErrorf(ErrTest, "assert.ErrTest general error for testing", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualErrorf(ErrTest, "wrong error message", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualErrorf should call FailNow()")
		}
	})
}

func TestAssertionsEqualExportedValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualExportedValuesf(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualExportedValuesf(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualExportedValuesf should call FailNow()")
		}
	})
}

func TestAssertionsEqualValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualValuesf(uint32(123), int32(123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualValuesf(uint32(123), int32(456), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualValuesf should call FailNow()")
		}
	})
}

func TestAssertionsErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Errorf(ErrTest, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Errorf(nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Errorf should call FailNow()")
		}
	})
}

func TestAssertionsErrorAsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAsf(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAsf(ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorAsf should call FailNow()")
		}
	})
}

func TestAssertionsErrorContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorContainsf(ErrTest, "general error", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorContainsf(ErrTest, "not in message", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorContainsf should call FailNow()")
		}
	})
}

func TestAssertionsErrorIsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorIsf(fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorIsf(ErrTest, io.EOF, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorIsf should call FailNow()")
		}
	})
}

func TestAssertionsEventuallyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Eventuallyf(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Eventuallyf(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Eventuallyf should call FailNow()")
		}
	})
}

func TestAssertionsEventuallyWithf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EventuallyWithf(func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EventuallyWithf(func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EventuallyWithf should call FailNow()")
		}
	})
}

func TestAssertionsExactlyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Exactlyf(int32(123), int32(123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Exactlyf(int32(123), int64(123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Exactlyf should call FailNow()")
		}
	})
}

func TestAssertionsFailf(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Failf("failed", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Failf should call FailNow()")
		}
	})
}

func TestAssertionsFailNowf(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FailNowf("failed", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FailNowf should call FailNow()")
		}
	})
}

func TestAssertionsFalsef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Falsef(1 == 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Falsef(1 == 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Falsef should call FailNow()")
		}
	})
}

func TestAssertionsFileEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileEmptyf(filepath.Join(testDataPath(), "empty_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileEmptyf(filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileEmptyf should call FailNow()")
		}
	})
}

func TestAssertionsFileExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileExistsf(filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileExistsf(filepath.Join(testDataPath(), "non_existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileExistsf should call FailNow()")
		}
	})
}

func TestAssertionsFileNotEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotEmptyf(filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotEmptyf(filepath.Join(testDataPath(), "empty_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileNotEmptyf should call FailNow()")
		}
	})
}

func TestAssertionsFileNotExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotExistsf(filepath.Join(testDataPath(), "non_existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FileNotExistsf(filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FileNotExistsf should call FailNow()")
		}
	})
}

func TestAssertionsGreaterf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Greaterf(2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Greaterf(1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Greaterf should call FailNow()")
		}
	})
}

func TestAssertionsGreaterOrEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqualf(2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqualf(1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqualf should call FailNow()")
		}
	})
}

func TestAssertionsHTTPBodyContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyContainsf(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyContainsf(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPBodyContainsf should call FailNow()")
		}
	})
}

func TestAssertionsHTTPBodyNotContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyNotContainsf(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPBodyNotContainsf(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPBodyNotContainsf should call FailNow()")
		}
	})
}

func TestAssertionsHTTPErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPErrorf(httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPErrorf(httpOK, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPErrorf should call FailNow()")
		}
	})
}

func TestAssertionsHTTPRedirectf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPRedirectf(httpRedirect, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPRedirectf(httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPRedirectf should call FailNow()")
		}
	})
}

func TestAssertionsHTTPStatusCodef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPStatusCodef(httpOK, "GET", "/", nil, http.StatusOK, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPStatusCodef(httpError, "GET", "/", nil, http.StatusOK, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPStatusCodef should call FailNow()")
		}
	})
}

func TestAssertionsHTTPSuccessf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPSuccessf(httpOK, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.HTTPSuccessf(httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.HTTPSuccessf should call FailNow()")
		}
	})
}

func TestAssertionsImplementsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Implementsf(ptr(dummyInterface), new(testing.T), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Implementsf((*error)(nil), new(testing.T), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Implementsf should call FailNow()")
		}
	})
}

func TestAssertionsInDeltaf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaf(1.0, 1.01, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaf(1.0, 1.1, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDeltaf should call FailNow()")
		}
	})
}

func TestAssertionsInDeltaMapValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaMapValuesf(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaMapValuesf(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDeltaMapValuesf should call FailNow()")
		}
	})
}

func TestAssertionsInDeltaSlicef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaSlicef([]float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaSlicef([]float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDeltaSlicef should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilonf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonf(100.0, 101.0, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonf(100.0, 110.0, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilonf should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilonSlicef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSlicef([]float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSlicef([]float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilonSlicef should call FailNow()")
		}
	})
}

func TestAssertionsIsDecreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasingf([]int{3, 2, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasingf([]int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsDecreasingf should call FailNow()")
		}
	})
}

func TestAssertionsIsIncreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasingf([]int{1, 2, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasingf([]int{1, 1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsIncreasingf should call FailNow()")
		}
	})
}

func TestAssertionsIsNonDecreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasingf([]int{1, 1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasingf([]int{2, 1, 0}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasingf should call FailNow()")
		}
	})
}

func TestAssertionsIsNonIncreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasingf([]int{2, 1, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasingf([]int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasingf should call FailNow()")
		}
	})
}

func TestAssertionsIsNotTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotTypef(int32(123), int64(456), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotTypef(123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNotTypef should call FailNow()")
		}
	})
}

func TestAssertionsIsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsTypef(123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsTypef(int32(123), int64(456), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsTypef should call FailNow()")
		}
	})
}

func TestAssertionsJSONEqf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONEqf should call FailNow()")
		}
	})
}

func TestAssertionsJSONEqBytesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqBytesf([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqBytesf([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONEqBytesf should call FailNow()")
		}
	})
}

func TestAssertionsKindf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Kindf(reflect.String, "hello", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Kindf(reflect.String, 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Kindf should call FailNow()")
		}
	})
}

func TestAssertionsLenf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Lenf([]string{"A", "B"}, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Lenf([]string{"A", "B"}, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Lenf should call FailNow()")
		}
	})
}

func TestAssertionsLessf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Lessf(1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Lessf(2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Lessf should call FailNow()")
		}
	})
}

func TestAssertionsLessOrEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqualf(1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqualf(2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.LessOrEqualf should call FailNow()")
		}
	})
}

func TestAssertionsNegativef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Negativef(-1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Negativef(1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Negativef should call FailNow()")
		}
	})
}

func TestAssertionsNeverf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Neverf(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Neverf(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Neverf should call FailNow()")
		}
	})
}

func TestAssertionsNilf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Nilf(nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Nilf("not nil", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Nilf should call FailNow()")
		}
	})
}

func TestAssertionsNoErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoErrorf(nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoErrorf(ErrTest, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NoErrorf should call FailNow()")
		}
	})
}

func TestAssertionsNoFileDescriptorLeakf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoFileDescriptorLeakf(func() {}, "test message")
		// require functions don't return a value
	})
}

func TestAssertionsNoGoRoutineLeakf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NoGoRoutineLeakf(func() {}, "test message")
		// require functions don't return a value
	})
}

func TestAssertionsNotContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotContainsf([]string{"A", "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotContainsf([]string{"A", "B"}, "B", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotContainsf should call FailNow()")
		}
	})
}

func TestAssertionsNotElementsMatchf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatchf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatchf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotElementsMatchf should call FailNow()")
		}
	})
}

func TestAssertionsNotEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEmptyf("not empty", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEmptyf("", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEmptyf should call FailNow()")
		}
	})
}

func TestAssertionsNotEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualf(123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualf(123, 123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEqualf should call FailNow()")
		}
	})
}

func TestAssertionsNotEqualValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualValuesf(uint32(123), int32(456), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualValuesf(uint32(123), int32(123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEqualValuesf should call FailNow()")
		}
	})
}

func TestAssertionsNotErrorAsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAsf(ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAsf(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotErrorAsf should call FailNow()")
		}
	})
}

func TestAssertionsNotErrorIsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorIsf(ErrTest, io.EOF, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorIsf(fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotErrorIsf should call FailNow()")
		}
	})
}

func TestAssertionsNotImplementsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotImplementsf((*error)(nil), new(testing.T), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotImplementsf(ptr(dummyInterface), new(testing.T), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotImplementsf should call FailNow()")
		}
	})
}

func TestAssertionsNotKindf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotKindf(reflect.String, 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotKindf(reflect.String, "hello", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotKindf should call FailNow()")
		}
	})
}

func TestAssertionsNotNilf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotNilf("not nil", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotNilf(nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotNilf should call FailNow()")
		}
	})
}

func TestAssertionsNotPanicsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotPanicsf(func() {}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotPanicsf(func() { panic("panicking") }, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotPanicsf should call FailNow()")
		}
	})
}

func TestAssertionsNotRegexpf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexpf("^start", "not starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexpf("^start", "starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotRegexpf should call FailNow()")
		}
	})
}

func TestAssertionsNotSamef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSamef(&staticVar, ptr("static string"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSamef(&staticVar, staticVarPtr, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSamef should call FailNow()")
		}
	})
}

func TestAssertionsNotSubsetf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSubsetf([]int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSubsetf([]int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSubsetf should call FailNow()")
		}
	})
}

func TestAssertionsNotZerof(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotZerof(1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotZerof(0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotZerof should call FailNow()")
		}
	})
}

func TestAssertionsPanicsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panicsf(func() { panic("panicking") }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panicsf(func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Panicsf should call FailNow()")
		}
	})
}

func TestAssertionsPanicsWithErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithErrorf(ErrTest.Error(), func() { panic(ErrTest) }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithErrorf(ErrTest.Error(), func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.PanicsWithErrorf should call FailNow()")
		}
	})
}

func TestAssertionsPanicsWithValuef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithValuef("panicking", func() { panic("panicking") }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PanicsWithValuef("panicking", func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.PanicsWithValuef should call FailNow()")
		}
	})
}

func TestAssertionsPositivef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Positivef(1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Positivef(-1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Positivef should call FailNow()")
		}
	})
}

func TestAssertionsRegexpf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Regexpf("^start", "starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Regexpf("^start", "not starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Regexpf should call FailNow()")
		}
	})
}

func TestAssertionsSamef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Samef(&staticVar, staticVarPtr, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Samef(&staticVar, ptr("static string"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Samef should call FailNow()")
		}
	})
}

func TestAssertionsSubsetf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Subsetf([]int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Subsetf([]int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Subsetf should call FailNow()")
		}
	})
}

func TestAssertionsTruef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Truef(1 == 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Truef(1 == 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Truef should call FailNow()")
		}
	})
}

func TestAssertionsWithinDurationf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinDurationf(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinDurationf(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.WithinDurationf should call FailNow()")
		}
	})
}

func TestAssertionsWithinRangef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinRangef(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.WithinRangef(time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.WithinRangef should call FailNow()")
		}
	})
}

func TestAssertionsYAMLEqf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLEqf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLEqf should panic as expected")
		}
	})
}

func TestAssertionsYAMLEqBytesf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLEqBytesf([]byte("key: value"), []byte("key: value"), "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLEqBytesf should panic as expected")
		}
	})
}

func TestAssertionsZerof(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Zerof(0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Zerof(1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Zerof should call FailNow()")
		}
	})
}
