// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

package assert

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

		mock := new(mockT)
		a := New(mock)
		result := a.Condition(func() bool { return true })
		if !result {
			t.Error("Assertions.Condition should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Condition(func() bool { return false })
		if result {
			t.Error("Assertions.Condition should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Condition should mark test as failed")
		}
	})
}

func TestAssertionsContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Contains([]string{"A", "B"}, "A")
		if !result {
			t.Error("Assertions.Contains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Contains([]string{"A", "B"}, "C")
		if result {
			t.Error("Assertions.Contains should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Contains should mark test as failed")
		}
	})
}

func TestAssertionsDirExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirExists(filepath.Join(testDataPath(), "existing_dir"))
		if !result {
			t.Error("Assertions.DirExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirExists(filepath.Join(testDataPath(), "non_existing_dir"))
		if result {
			t.Error("Assertions.DirExists should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.DirExists should mark test as failed")
		}
	})
}

func TestAssertionsDirNotExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirNotExists(filepath.Join(testDataPath(), "non_existing_dir"))
		if !result {
			t.Error("Assertions.DirNotExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirNotExists(filepath.Join(testDataPath(), "existing_dir"))
		if result {
			t.Error("Assertions.DirNotExists should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.DirNotExists should mark test as failed")
		}
	})
}

func TestAssertionsElementsMatch(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatch([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if !result {
			t.Error("Assertions.ElementsMatch should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatch([]int{1, 2, 3}, []int{1, 2, 4})
		if result {
			t.Error("Assertions.ElementsMatch should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ElementsMatch should mark test as failed")
		}
	})
}

func TestAssertionsEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Empty("")
		if !result {
			t.Error("Assertions.Empty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Empty("not empty")
		if result {
			t.Error("Assertions.Empty should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Empty should mark test as failed")
		}
	})
}

func TestAssertionsEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Equal(123, 123)
		if !result {
			t.Error("Assertions.Equal should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Equal(123, 456)
		if result {
			t.Error("Assertions.Equal should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Equal should mark test as failed")
		}
	})
}

func TestAssertionsEqualError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualError(ErrTest, "assert.ErrTest general error for testing")
		if !result {
			t.Error("Assertions.EqualError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualError(ErrTest, "wrong error message")
		if result {
			t.Error("Assertions.EqualError should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualError should mark test as failed")
		}
	})
}

func TestAssertionsEqualExportedValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualExportedValues(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
		if !result {
			t.Error("Assertions.EqualExportedValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualExportedValues(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1})
		if result {
			t.Error("Assertions.EqualExportedValues should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualExportedValues should mark test as failed")
		}
	})
}

func TestAssertionsEqualValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualValues(uint32(123), int32(123))
		if !result {
			t.Error("Assertions.EqualValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualValues(uint32(123), int32(456))
		if result {
			t.Error("Assertions.EqualValues should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualValues should mark test as failed")
		}
	})
}

func TestAssertionsError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Error(ErrTest)
		if !result {
			t.Error("Assertions.Error should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Error(nil)
		if result {
			t.Error("Assertions.Error should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Error should mark test as failed")
		}
	})
}

func TestAssertionsErrorAs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAs(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if !result {
			t.Error("Assertions.ErrorAs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAs(ErrTest, new(*dummyError))
		if result {
			t.Error("Assertions.ErrorAs should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorAs should mark test as failed")
		}
	})
}

func TestAssertionsErrorContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorContains(ErrTest, "general error")
		if !result {
			t.Error("Assertions.ErrorContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorContains(ErrTest, "not in message")
		if result {
			t.Error("Assertions.ErrorContains should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorContains should mark test as failed")
		}
	})
}

func TestAssertionsErrorIs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorIs(fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		if !result {
			t.Error("Assertions.ErrorIs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorIs(ErrTest, io.EOF)
		if result {
			t.Error("Assertions.ErrorIs should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorIs should mark test as failed")
		}
	})
}

func TestAssertionsEventually(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Eventually(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		if !result {
			t.Error("Assertions.Eventually should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Eventually(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("Assertions.Eventually should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Eventually should mark test as failed")
		}
	})
}

func TestAssertionsEventuallyWith(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EventuallyWith(func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		if !result {
			t.Error("Assertions.EventuallyWith should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EventuallyWith(func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("Assertions.EventuallyWith should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EventuallyWith should mark test as failed")
		}
	})
}

func TestAssertionsExactly(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Exactly(int32(123), int32(123))
		if !result {
			t.Error("Assertions.Exactly should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Exactly(int32(123), int64(123))
		if result {
			t.Error("Assertions.Exactly should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Exactly should mark test as failed")
		}
	})
}

func TestAssertionsFail(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Fail("failed")
		if result {
			t.Error("Assertions.Fail should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Fail should mark test as failed")
		}
	})
}

func TestAssertionsFailNow(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		result := a.FailNow("failed")
		if result {
			t.Error("Assertions.FailNow should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FailNow should call FailNow()")
		}
	})
}

func TestAssertionsFalse(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.False(1 == 0)
		if !result {
			t.Error("Assertions.False should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.False(1 == 1)
		if result {
			t.Error("Assertions.False should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.False should mark test as failed")
		}
	})
}

func TestAssertionsFileEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileEmpty(filepath.Join(testDataPath(), "empty_file"))
		if !result {
			t.Error("Assertions.FileEmpty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileEmpty(filepath.Join(testDataPath(), "existing_file"))
		if result {
			t.Error("Assertions.FileEmpty should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileEmpty should mark test as failed")
		}
	})
}

func TestAssertionsFileExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileExists(filepath.Join(testDataPath(), "existing_file"))
		if !result {
			t.Error("Assertions.FileExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileExists(filepath.Join(testDataPath(), "non_existing_file"))
		if result {
			t.Error("Assertions.FileExists should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileExists should mark test as failed")
		}
	})
}

func TestAssertionsFileNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotEmpty(filepath.Join(testDataPath(), "existing_file"))
		if !result {
			t.Error("Assertions.FileNotEmpty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotEmpty(filepath.Join(testDataPath(), "empty_file"))
		if result {
			t.Error("Assertions.FileNotEmpty should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileNotEmpty should mark test as failed")
		}
	})
}

func TestAssertionsFileNotExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotExists(filepath.Join(testDataPath(), "non_existing_file"))
		if !result {
			t.Error("Assertions.FileNotExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotExists(filepath.Join(testDataPath(), "existing_file"))
		if result {
			t.Error("Assertions.FileNotExists should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileNotExists should mark test as failed")
		}
	})
}

func TestAssertionsGreater(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Greater(2, 1)
		if !result {
			t.Error("Assertions.Greater should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Greater(1, 2)
		if result {
			t.Error("Assertions.Greater should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Greater should mark test as failed")
		}
	})
}

func TestAssertionsGreaterOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqual(2, 1)
		if !result {
			t.Error("Assertions.GreaterOrEqual should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqual(1, 2)
		if result {
			t.Error("Assertions.GreaterOrEqual should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqual should mark test as failed")
		}
	})
}

func TestAssertionsHTTPBodyContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyContains(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!")
		if !result {
			t.Error("Assertions.HTTPBodyContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyContains(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!")
		if result {
			t.Error("Assertions.HTTPBodyContains should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPBodyContains should mark test as failed")
		}
	})
}

func TestAssertionsHTTPBodyNotContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyNotContains(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!")
		if !result {
			t.Error("Assertions.HTTPBodyNotContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyNotContains(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!")
		if result {
			t.Error("Assertions.HTTPBodyNotContains should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPBodyNotContains should mark test as failed")
		}
	})
}

func TestAssertionsHTTPError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPError(httpError, "GET", "/", nil)
		if !result {
			t.Error("Assertions.HTTPError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPError(httpOK, "GET", "/", nil)
		if result {
			t.Error("Assertions.HTTPError should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPError should mark test as failed")
		}
	})
}

func TestAssertionsHTTPRedirect(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPRedirect(httpRedirect, "GET", "/", nil)
		if !result {
			t.Error("Assertions.HTTPRedirect should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPRedirect(httpError, "GET", "/", nil)
		if result {
			t.Error("Assertions.HTTPRedirect should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPRedirect should mark test as failed")
		}
	})
}

func TestAssertionsHTTPStatusCode(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPStatusCode(httpOK, "GET", "/", nil, http.StatusOK)
		if !result {
			t.Error("Assertions.HTTPStatusCode should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPStatusCode(httpError, "GET", "/", nil, http.StatusOK)
		if result {
			t.Error("Assertions.HTTPStatusCode should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPStatusCode should mark test as failed")
		}
	})
}

func TestAssertionsHTTPSuccess(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPSuccess(httpOK, "GET", "/", nil)
		if !result {
			t.Error("Assertions.HTTPSuccess should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPSuccess(httpError, "GET", "/", nil)
		if result {
			t.Error("Assertions.HTTPSuccess should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPSuccess should mark test as failed")
		}
	})
}

func TestAssertionsImplements(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Implements(ptr(dummyInterface), new(testing.T))
		if !result {
			t.Error("Assertions.Implements should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Implements((*error)(nil), new(testing.T))
		if result {
			t.Error("Assertions.Implements should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Implements should mark test as failed")
		}
	})
}

func TestAssertionsInDelta(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDelta(1.0, 1.01, 0.02)
		if !result {
			t.Error("Assertions.InDelta should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDelta(1.0, 1.1, 0.05)
		if result {
			t.Error("Assertions.InDelta should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDelta should mark test as failed")
		}
	})
}

func TestAssertionsInDeltaMapValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaMapValues(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
		if !result {
			t.Error("Assertions.InDeltaMapValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaMapValues(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05)
		if result {
			t.Error("Assertions.InDeltaMapValues should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDeltaMapValues should mark test as failed")
		}
	})
}

func TestAssertionsInDeltaSlice(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaSlice([]float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
		if !result {
			t.Error("Assertions.InDeltaSlice should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaSlice([]float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05)
		if result {
			t.Error("Assertions.InDeltaSlice should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDeltaSlice should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilon(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilon(100.0, 101.0, 0.02)
		if !result {
			t.Error("Assertions.InEpsilon should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilon(100.0, 110.0, 0.05)
		if result {
			t.Error("Assertions.InEpsilon should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilon should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilonSlice(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSlice([]float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
		if !result {
			t.Error("Assertions.InEpsilonSlice should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSlice([]float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05)
		if result {
			t.Error("Assertions.InEpsilonSlice should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilonSlice should mark test as failed")
		}
	})
}

func TestAssertionsIsDecreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasing([]int{3, 2, 1})
		if !result {
			t.Error("Assertions.IsDecreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasing([]int{1, 2, 3})
		if result {
			t.Error("Assertions.IsDecreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsDecreasing should mark test as failed")
		}
	})
}

func TestAssertionsIsIncreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasing([]int{1, 2, 3})
		if !result {
			t.Error("Assertions.IsIncreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasing([]int{1, 1, 2})
		if result {
			t.Error("Assertions.IsIncreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsIncreasing should mark test as failed")
		}
	})
}

func TestAssertionsIsNonDecreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasing([]int{1, 1, 2})
		if !result {
			t.Error("Assertions.IsNonDecreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasing([]int{2, 1, 0})
		if result {
			t.Error("Assertions.IsNonDecreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasing should mark test as failed")
		}
	})
}

func TestAssertionsIsNonIncreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasing([]int{2, 1, 1})
		if !result {
			t.Error("Assertions.IsNonIncreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasing([]int{1, 2, 3})
		if result {
			t.Error("Assertions.IsNonIncreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasing should mark test as failed")
		}
	})
}

func TestAssertionsIsNotType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotType(int32(123), int64(456))
		if !result {
			t.Error("Assertions.IsNotType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotType(123, 456)
		if result {
			t.Error("Assertions.IsNotType should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNotType should mark test as failed")
		}
	})
}

func TestAssertionsIsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsType(123, 456)
		if !result {
			t.Error("Assertions.IsType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsType(int32(123), int64(456))
		if result {
			t.Error("Assertions.IsType should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsType should mark test as failed")
		}
	})
}

func TestAssertionsJSONEq(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEq(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
		if !result {
			t.Error("Assertions.JSONEq should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEq(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		if result {
			t.Error("Assertions.JSONEq should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONEq should mark test as failed")
		}
	})
}

func TestAssertionsJSONEqBytes(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqBytes([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
		if !result {
			t.Error("Assertions.JSONEqBytes should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqBytes([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`))
		if result {
			t.Error("Assertions.JSONEqBytes should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONEqBytes should mark test as failed")
		}
	})
}

func TestAssertionsKind(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Kind(reflect.String, "hello")
		if !result {
			t.Error("Assertions.Kind should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Kind(reflect.String, 0)
		if result {
			t.Error("Assertions.Kind should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Kind should mark test as failed")
		}
	})
}

func TestAssertionsLen(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Len([]string{"A", "B"}, 2)
		if !result {
			t.Error("Assertions.Len should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Len([]string{"A", "B"}, 1)
		if result {
			t.Error("Assertions.Len should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Len should mark test as failed")
		}
	})
}

func TestAssertionsLess(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Less(1, 2)
		if !result {
			t.Error("Assertions.Less should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Less(2, 1)
		if result {
			t.Error("Assertions.Less should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Less should mark test as failed")
		}
	})
}

func TestAssertionsLessOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqual(1, 2)
		if !result {
			t.Error("Assertions.LessOrEqual should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqual(2, 1)
		if result {
			t.Error("Assertions.LessOrEqual should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.LessOrEqual should mark test as failed")
		}
	})
}

func TestAssertionsNegative(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Negative(-1)
		if !result {
			t.Error("Assertions.Negative should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Negative(1)
		if result {
			t.Error("Assertions.Negative should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Negative should mark test as failed")
		}
	})
}

func TestAssertionsNever(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Never(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		if !result {
			t.Error("Assertions.Never should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Never(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("Assertions.Never should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Never should mark test as failed")
		}
	})
}

func TestAssertionsNil(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Nil(nil)
		if !result {
			t.Error("Assertions.Nil should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Nil("not nil")
		if result {
			t.Error("Assertions.Nil should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Nil should mark test as failed")
		}
	})
}

func TestAssertionsNoError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoError(nil)
		if !result {
			t.Error("Assertions.NoError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoError(ErrTest)
		if result {
			t.Error("Assertions.NoError should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NoError should mark test as failed")
		}
	})
}

func TestAssertionsNoFileDescriptorLeak(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoFileDescriptorLeak(func() {})
		if !result {
			t.Error("Assertions.NoFileDescriptorLeak should return true on success")
		}
	})
}

func TestAssertionsNoGoRoutineLeak(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoGoRoutineLeak(func() {})
		if !result {
			t.Error("Assertions.NoGoRoutineLeak should return true on success")
		}
	})
}

func TestAssertionsNotContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotContains([]string{"A", "B"}, "C")
		if !result {
			t.Error("Assertions.NotContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotContains([]string{"A", "B"}, "B")
		if result {
			t.Error("Assertions.NotContains should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotContains should mark test as failed")
		}
	})
}

func TestAssertionsNotElementsMatch(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatch([]int{1, 2, 3}, []int{1, 2, 4})
		if !result {
			t.Error("Assertions.NotElementsMatch should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatch([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if result {
			t.Error("Assertions.NotElementsMatch should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotElementsMatch should mark test as failed")
		}
	})
}

func TestAssertionsNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEmpty("not empty")
		if !result {
			t.Error("Assertions.NotEmpty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEmpty("")
		if result {
			t.Error("Assertions.NotEmpty should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEmpty should mark test as failed")
		}
	})
}

func TestAssertionsNotEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqual(123, 456)
		if !result {
			t.Error("Assertions.NotEqual should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqual(123, 123)
		if result {
			t.Error("Assertions.NotEqual should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEqual should mark test as failed")
		}
	})
}

func TestAssertionsNotEqualValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualValues(uint32(123), int32(456))
		if !result {
			t.Error("Assertions.NotEqualValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualValues(uint32(123), int32(123))
		if result {
			t.Error("Assertions.NotEqualValues should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEqualValues should mark test as failed")
		}
	})
}

func TestAssertionsNotErrorAs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAs(ErrTest, new(*dummyError))
		if !result {
			t.Error("Assertions.NotErrorAs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAs(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if result {
			t.Error("Assertions.NotErrorAs should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotErrorAs should mark test as failed")
		}
	})
}

func TestAssertionsNotErrorIs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorIs(ErrTest, io.EOF)
		if !result {
			t.Error("Assertions.NotErrorIs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorIs(fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		if result {
			t.Error("Assertions.NotErrorIs should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotErrorIs should mark test as failed")
		}
	})
}

func TestAssertionsNotImplements(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotImplements((*error)(nil), new(testing.T))
		if !result {
			t.Error("Assertions.NotImplements should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotImplements(ptr(dummyInterface), new(testing.T))
		if result {
			t.Error("Assertions.NotImplements should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotImplements should mark test as failed")
		}
	})
}

func TestAssertionsNotKind(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotKind(reflect.String, 0)
		if !result {
			t.Error("Assertions.NotKind should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotKind(reflect.String, "hello")
		if result {
			t.Error("Assertions.NotKind should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotKind should mark test as failed")
		}
	})
}

func TestAssertionsNotNil(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotNil("not nil")
		if !result {
			t.Error("Assertions.NotNil should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotNil(nil)
		if result {
			t.Error("Assertions.NotNil should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotNil should mark test as failed")
		}
	})
}

func TestAssertionsNotPanics(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotPanics(func() {})
		if !result {
			t.Error("Assertions.NotPanics should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotPanics(func() { panic("panicking") })
		if result {
			t.Error("Assertions.NotPanics should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotPanics should mark test as failed")
		}
	})
}

func TestAssertionsNotRegexp(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexp("^start", "not starting")
		if !result {
			t.Error("Assertions.NotRegexp should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexp("^start", "starting")
		if result {
			t.Error("Assertions.NotRegexp should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotRegexp should mark test as failed")
		}
	})
}

func TestAssertionsNotSame(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSame(&staticVar, ptr("static string"))
		if !result {
			t.Error("Assertions.NotSame should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSame(&staticVar, staticVarPtr)
		if result {
			t.Error("Assertions.NotSame should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSame should mark test as failed")
		}
	})
}

func TestAssertionsNotSubset(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSubset([]int{1, 2, 3}, []int{4, 5})
		if !result {
			t.Error("Assertions.NotSubset should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSubset([]int{1, 2, 3}, []int{1, 2})
		if result {
			t.Error("Assertions.NotSubset should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSubset should mark test as failed")
		}
	})
}

func TestAssertionsNotZero(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotZero(1)
		if !result {
			t.Error("Assertions.NotZero should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotZero(0)
		if result {
			t.Error("Assertions.NotZero should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotZero should mark test as failed")
		}
	})
}

func TestAssertionsPanics(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() { panic("panicking") })
		if !result {
			t.Error("Assertions.Panics should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {})
		if result {
			t.Error("Assertions.Panics should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Panics should mark test as failed")
		}
	})
}

func TestAssertionsPanicsWithError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithError(ErrTest.Error(), func() { panic(ErrTest) })
		if !result {
			t.Error("Assertions.PanicsWithError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithError(ErrTest.Error(), func() {})
		if result {
			t.Error("Assertions.PanicsWithError should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.PanicsWithError should mark test as failed")
		}
	})
}

func TestAssertionsPanicsWithValue(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithValue("panicking", func() { panic("panicking") })
		if !result {
			t.Error("Assertions.PanicsWithValue should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithValue("panicking", func() {})
		if result {
			t.Error("Assertions.PanicsWithValue should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.PanicsWithValue should mark test as failed")
		}
	})
}

func TestAssertionsPositive(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Positive(1)
		if !result {
			t.Error("Assertions.Positive should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Positive(-1)
		if result {
			t.Error("Assertions.Positive should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Positive should mark test as failed")
		}
	})
}

func TestAssertionsRegexp(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Regexp("^start", "starting")
		if !result {
			t.Error("Assertions.Regexp should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Regexp("^start", "not starting")
		if result {
			t.Error("Assertions.Regexp should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Regexp should mark test as failed")
		}
	})
}

func TestAssertionsSame(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Same(&staticVar, staticVarPtr)
		if !result {
			t.Error("Assertions.Same should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Same(&staticVar, ptr("static string"))
		if result {
			t.Error("Assertions.Same should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Same should mark test as failed")
		}
	})
}

func TestAssertionsSubset(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Subset([]int{1, 2, 3}, []int{1, 2})
		if !result {
			t.Error("Assertions.Subset should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Subset([]int{1, 2, 3}, []int{4, 5})
		if result {
			t.Error("Assertions.Subset should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Subset should mark test as failed")
		}
	})
}

func TestAssertionsTrue(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.True(1 == 1)
		if !result {
			t.Error("Assertions.True should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.True(1 == 0)
		if result {
			t.Error("Assertions.True should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.True should mark test as failed")
		}
	})
}

func TestAssertionsWithinDuration(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinDuration(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second)
		if !result {
			t.Error("Assertions.WithinDuration should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinDuration(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second)
		if result {
			t.Error("Assertions.WithinDuration should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.WithinDuration should mark test as failed")
		}
	})
}

func TestAssertionsWithinRange(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinRange(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		if !result {
			t.Error("Assertions.WithinRange should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinRange(time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		if result {
			t.Error("Assertions.WithinRange should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.WithinRange should mark test as failed")
		}
	})
}

func TestAssertionsYAMLEq(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLEq("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLEq should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLEq should panic as expected")
		}
	})
}

func TestAssertionsYAMLEqBytes(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLEqBytes([]byte("key: value"), []byte("key: value"))
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLEqBytes should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLEqBytes should panic as expected")
		}
	})
}

func TestAssertionsZero(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Zero(0)
		if !result {
			t.Error("Assertions.Zero should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Zero(1)
		if result {
			t.Error("Assertions.Zero should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Zero should mark test as failed")
		}
	})
}

func TestAssertionsConditionf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Conditionf(func() bool { return true }, "test message")
		if !result {
			t.Error("Assertions.Conditionf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Conditionf(func() bool { return false }, "test message")
		if result {
			t.Error("Assertions.Conditionf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Conditionf should mark test as failed")
		}
	})
}

func TestAssertionsContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Containsf([]string{"A", "B"}, "A", "test message")
		if !result {
			t.Error("Assertions.Containsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Containsf([]string{"A", "B"}, "C", "test message")
		if result {
			t.Error("Assertions.Containsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Containsf should mark test as failed")
		}
	})
}

func TestAssertionsDirExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirExistsf(filepath.Join(testDataPath(), "existing_dir"), "test message")
		if !result {
			t.Error("Assertions.DirExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirExistsf(filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		if result {
			t.Error("Assertions.DirExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.DirExistsf should mark test as failed")
		}
	})
}

func TestAssertionsDirNotExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirNotExistsf(filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		if !result {
			t.Error("Assertions.DirNotExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.DirNotExistsf(filepath.Join(testDataPath(), "existing_dir"), "test message")
		if result {
			t.Error("Assertions.DirNotExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.DirNotExistsf should mark test as failed")
		}
	})
}

func TestAssertionsElementsMatchf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatchf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if !result {
			t.Error("Assertions.ElementsMatchf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatchf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if result {
			t.Error("Assertions.ElementsMatchf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ElementsMatchf should mark test as failed")
		}
	})
}

func TestAssertionsEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Emptyf("", "test message")
		if !result {
			t.Error("Assertions.Emptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Emptyf("not empty", "test message")
		if result {
			t.Error("Assertions.Emptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Emptyf should mark test as failed")
		}
	})
}

func TestAssertionsEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Equalf(123, 123, "test message")
		if !result {
			t.Error("Assertions.Equalf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Equalf(123, 456, "test message")
		if result {
			t.Error("Assertions.Equalf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Equalf should mark test as failed")
		}
	})
}

func TestAssertionsEqualErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualErrorf(ErrTest, "assert.ErrTest general error for testing", "test message")
		if !result {
			t.Error("Assertions.EqualErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualErrorf(ErrTest, "wrong error message", "test message")
		if result {
			t.Error("Assertions.EqualErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualErrorf should mark test as failed")
		}
	})
}

func TestAssertionsEqualExportedValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualExportedValuesf(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}, "test message")
		if !result {
			t.Error("Assertions.EqualExportedValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualExportedValuesf(&dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}, "test message")
		if result {
			t.Error("Assertions.EqualExportedValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualExportedValuesf should mark test as failed")
		}
	})
}

func TestAssertionsEqualValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualValuesf(uint32(123), int32(123), "test message")
		if !result {
			t.Error("Assertions.EqualValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualValuesf(uint32(123), int32(456), "test message")
		if result {
			t.Error("Assertions.EqualValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualValuesf should mark test as failed")
		}
	})
}

func TestAssertionsErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Errorf(ErrTest, "test message")
		if !result {
			t.Error("Assertions.Errorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Errorf(nil, "test message")
		if result {
			t.Error("Assertions.Errorf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Errorf should mark test as failed")
		}
	})
}

func TestAssertionsErrorAsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAsf(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if !result {
			t.Error("Assertions.ErrorAsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAsf(ErrTest, new(*dummyError), "test message")
		if result {
			t.Error("Assertions.ErrorAsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorAsf should mark test as failed")
		}
	})
}

func TestAssertionsErrorContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorContainsf(ErrTest, "general error", "test message")
		if !result {
			t.Error("Assertions.ErrorContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorContainsf(ErrTest, "not in message", "test message")
		if result {
			t.Error("Assertions.ErrorContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorContainsf should mark test as failed")
		}
	})
}

func TestAssertionsErrorIsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorIsf(fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		if !result {
			t.Error("Assertions.ErrorIsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorIsf(ErrTest, io.EOF, "test message")
		if result {
			t.Error("Assertions.ErrorIsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorIsf should mark test as failed")
		}
	})
}

func TestAssertionsEventuallyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Eventuallyf(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if !result {
			t.Error("Assertions.Eventuallyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Eventuallyf(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("Assertions.Eventuallyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Eventuallyf should mark test as failed")
		}
	})
}

func TestAssertionsEventuallyWithf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EventuallyWithf(func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if !result {
			t.Error("Assertions.EventuallyWithf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EventuallyWithf(func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("Assertions.EventuallyWithf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EventuallyWithf should mark test as failed")
		}
	})
}

func TestAssertionsExactlyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Exactlyf(int32(123), int32(123), "test message")
		if !result {
			t.Error("Assertions.Exactlyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Exactlyf(int32(123), int64(123), "test message")
		if result {
			t.Error("Assertions.Exactlyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Exactlyf should mark test as failed")
		}
	})
}

func TestAssertionsFailf(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Failf("failed", "test message")
		if result {
			t.Error("Assertions.Failf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Failf should mark test as failed")
		}
	})
}

func TestAssertionsFailNowf(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		result := a.FailNowf("failed", "test message")
		if result {
			t.Error("Assertions.FailNowf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FailNowf should call FailNow()")
		}
	})
}

func TestAssertionsFalsef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Falsef(1 == 0, "test message")
		if !result {
			t.Error("Assertions.Falsef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Falsef(1 == 1, "test message")
		if result {
			t.Error("Assertions.Falsef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Falsef should mark test as failed")
		}
	})
}

func TestAssertionsFileEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileEmptyf(filepath.Join(testDataPath(), "empty_file"), "test message")
		if !result {
			t.Error("Assertions.FileEmptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileEmptyf(filepath.Join(testDataPath(), "existing_file"), "test message")
		if result {
			t.Error("Assertions.FileEmptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileEmptyf should mark test as failed")
		}
	})
}

func TestAssertionsFileExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileExistsf(filepath.Join(testDataPath(), "existing_file"), "test message")
		if !result {
			t.Error("Assertions.FileExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileExistsf(filepath.Join(testDataPath(), "non_existing_file"), "test message")
		if result {
			t.Error("Assertions.FileExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileExistsf should mark test as failed")
		}
	})
}

func TestAssertionsFileNotEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotEmptyf(filepath.Join(testDataPath(), "existing_file"), "test message")
		if !result {
			t.Error("Assertions.FileNotEmptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotEmptyf(filepath.Join(testDataPath(), "empty_file"), "test message")
		if result {
			t.Error("Assertions.FileNotEmptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileNotEmptyf should mark test as failed")
		}
	})
}

func TestAssertionsFileNotExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotExistsf(filepath.Join(testDataPath(), "non_existing_file"), "test message")
		if !result {
			t.Error("Assertions.FileNotExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FileNotExistsf(filepath.Join(testDataPath(), "existing_file"), "test message")
		if result {
			t.Error("Assertions.FileNotExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FileNotExistsf should mark test as failed")
		}
	})
}

func TestAssertionsGreaterf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Greaterf(2, 1, "test message")
		if !result {
			t.Error("Assertions.Greaterf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Greaterf(1, 2, "test message")
		if result {
			t.Error("Assertions.Greaterf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Greaterf should mark test as failed")
		}
	})
}

func TestAssertionsGreaterOrEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqualf(2, 1, "test message")
		if !result {
			t.Error("Assertions.GreaterOrEqualf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqualf(1, 2, "test message")
		if result {
			t.Error("Assertions.GreaterOrEqualf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqualf should mark test as failed")
		}
	})
}

func TestAssertionsHTTPBodyContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyContainsf(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!", "test message")
		if !result {
			t.Error("Assertions.HTTPBodyContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyContainsf(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!", "test message")
		if result {
			t.Error("Assertions.HTTPBodyContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPBodyContainsf should mark test as failed")
		}
	})
}

func TestAssertionsHTTPBodyNotContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyNotContainsf(httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!", "test message")
		if !result {
			t.Error("Assertions.HTTPBodyNotContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPBodyNotContainsf(httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!", "test message")
		if result {
			t.Error("Assertions.HTTPBodyNotContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPBodyNotContainsf should mark test as failed")
		}
	})
}

func TestAssertionsHTTPErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPErrorf(httpError, "GET", "/", nil, "test message")
		if !result {
			t.Error("Assertions.HTTPErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPErrorf(httpOK, "GET", "/", nil, "test message")
		if result {
			t.Error("Assertions.HTTPErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPErrorf should mark test as failed")
		}
	})
}

func TestAssertionsHTTPRedirectf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPRedirectf(httpRedirect, "GET", "/", nil, "test message")
		if !result {
			t.Error("Assertions.HTTPRedirectf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPRedirectf(httpError, "GET", "/", nil, "test message")
		if result {
			t.Error("Assertions.HTTPRedirectf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPRedirectf should mark test as failed")
		}
	})
}

func TestAssertionsHTTPStatusCodef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPStatusCodef(httpOK, "GET", "/", nil, http.StatusOK, "test message")
		if !result {
			t.Error("Assertions.HTTPStatusCodef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPStatusCodef(httpError, "GET", "/", nil, http.StatusOK, "test message")
		if result {
			t.Error("Assertions.HTTPStatusCodef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPStatusCodef should mark test as failed")
		}
	})
}

func TestAssertionsHTTPSuccessf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPSuccessf(httpOK, "GET", "/", nil, "test message")
		if !result {
			t.Error("Assertions.HTTPSuccessf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.HTTPSuccessf(httpError, "GET", "/", nil, "test message")
		if result {
			t.Error("Assertions.HTTPSuccessf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.HTTPSuccessf should mark test as failed")
		}
	})
}

func TestAssertionsImplementsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Implementsf(ptr(dummyInterface), new(testing.T), "test message")
		if !result {
			t.Error("Assertions.Implementsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Implementsf((*error)(nil), new(testing.T), "test message")
		if result {
			t.Error("Assertions.Implementsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Implementsf should mark test as failed")
		}
	})
}

func TestAssertionsInDeltaf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaf(1.0, 1.01, 0.02, "test message")
		if !result {
			t.Error("Assertions.InDeltaf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaf(1.0, 1.1, 0.05, "test message")
		if result {
			t.Error("Assertions.InDeltaf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDeltaf should mark test as failed")
		}
	})
}

func TestAssertionsInDeltaMapValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaMapValuesf(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02, "test message")
		if !result {
			t.Error("Assertions.InDeltaMapValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaMapValuesf(map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05, "test message")
		if result {
			t.Error("Assertions.InDeltaMapValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDeltaMapValuesf should mark test as failed")
		}
	})
}

func TestAssertionsInDeltaSlicef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaSlicef([]float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02, "test message")
		if !result {
			t.Error("Assertions.InDeltaSlicef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaSlicef([]float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05, "test message")
		if result {
			t.Error("Assertions.InDeltaSlicef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDeltaSlicef should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilonf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonf(100.0, 101.0, 0.02, "test message")
		if !result {
			t.Error("Assertions.InEpsilonf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonf(100.0, 110.0, 0.05, "test message")
		if result {
			t.Error("Assertions.InEpsilonf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilonf should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilonSlicef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSlicef([]float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02, "test message")
		if !result {
			t.Error("Assertions.InEpsilonSlicef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSlicef([]float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05, "test message")
		if result {
			t.Error("Assertions.InEpsilonSlicef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilonSlicef should mark test as failed")
		}
	})
}

func TestAssertionsIsDecreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasingf([]int{3, 2, 1}, "test message")
		if !result {
			t.Error("Assertions.IsDecreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasingf([]int{1, 2, 3}, "test message")
		if result {
			t.Error("Assertions.IsDecreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsDecreasingf should mark test as failed")
		}
	})
}

func TestAssertionsIsIncreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasingf([]int{1, 2, 3}, "test message")
		if !result {
			t.Error("Assertions.IsIncreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasingf([]int{1, 1, 2}, "test message")
		if result {
			t.Error("Assertions.IsIncreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsIncreasingf should mark test as failed")
		}
	})
}

func TestAssertionsIsNonDecreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasingf([]int{1, 1, 2}, "test message")
		if !result {
			t.Error("Assertions.IsNonDecreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasingf([]int{2, 1, 0}, "test message")
		if result {
			t.Error("Assertions.IsNonDecreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasingf should mark test as failed")
		}
	})
}

func TestAssertionsIsNonIncreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasingf([]int{2, 1, 1}, "test message")
		if !result {
			t.Error("Assertions.IsNonIncreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasingf([]int{1, 2, 3}, "test message")
		if result {
			t.Error("Assertions.IsNonIncreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasingf should mark test as failed")
		}
	})
}

func TestAssertionsIsNotTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotTypef(int32(123), int64(456), "test message")
		if !result {
			t.Error("Assertions.IsNotTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotTypef(123, 456, "test message")
		if result {
			t.Error("Assertions.IsNotTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNotTypef should mark test as failed")
		}
	})
}

func TestAssertionsIsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsTypef(123, 456, "test message")
		if !result {
			t.Error("Assertions.IsTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsTypef(int32(123), int64(456), "test message")
		if result {
			t.Error("Assertions.IsTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsTypef should mark test as failed")
		}
	})
}

func TestAssertionsJSONEqf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "test message")
		if !result {
			t.Error("Assertions.JSONEqf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqf(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		if result {
			t.Error("Assertions.JSONEqf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONEqf should mark test as failed")
		}
	})
}

func TestAssertionsJSONEqBytesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqBytesf([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		if !result {
			t.Error("Assertions.JSONEqBytesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqBytesf([]byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`), "test message")
		if result {
			t.Error("Assertions.JSONEqBytesf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONEqBytesf should mark test as failed")
		}
	})
}

func TestAssertionsKindf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Kindf(reflect.String, "hello", "test message")
		if !result {
			t.Error("Assertions.Kindf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Kindf(reflect.String, 0, "test message")
		if result {
			t.Error("Assertions.Kindf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Kindf should mark test as failed")
		}
	})
}

func TestAssertionsLenf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Lenf([]string{"A", "B"}, 2, "test message")
		if !result {
			t.Error("Assertions.Lenf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Lenf([]string{"A", "B"}, 1, "test message")
		if result {
			t.Error("Assertions.Lenf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Lenf should mark test as failed")
		}
	})
}

func TestAssertionsLessf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Lessf(1, 2, "test message")
		if !result {
			t.Error("Assertions.Lessf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Lessf(2, 1, "test message")
		if result {
			t.Error("Assertions.Lessf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Lessf should mark test as failed")
		}
	})
}

func TestAssertionsLessOrEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqualf(1, 2, "test message")
		if !result {
			t.Error("Assertions.LessOrEqualf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqualf(2, 1, "test message")
		if result {
			t.Error("Assertions.LessOrEqualf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.LessOrEqualf should mark test as failed")
		}
	})
}

func TestAssertionsNegativef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Negativef(-1, "test message")
		if !result {
			t.Error("Assertions.Negativef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Negativef(1, "test message")
		if result {
			t.Error("Assertions.Negativef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Negativef should mark test as failed")
		}
	})
}

func TestAssertionsNeverf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Neverf(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if !result {
			t.Error("Assertions.Neverf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Neverf(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("Assertions.Neverf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Neverf should mark test as failed")
		}
	})
}

func TestAssertionsNilf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Nilf(nil, "test message")
		if !result {
			t.Error("Assertions.Nilf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Nilf("not nil", "test message")
		if result {
			t.Error("Assertions.Nilf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Nilf should mark test as failed")
		}
	})
}

func TestAssertionsNoErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoErrorf(nil, "test message")
		if !result {
			t.Error("Assertions.NoErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoErrorf(ErrTest, "test message")
		if result {
			t.Error("Assertions.NoErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NoErrorf should mark test as failed")
		}
	})
}

func TestAssertionsNoFileDescriptorLeakf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoFileDescriptorLeakf(func() {}, "test message")
		if !result {
			t.Error("Assertions.NoFileDescriptorLeakf should return true on success")
		}
	})
}

func TestAssertionsNoGoRoutineLeakf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NoGoRoutineLeakf(func() {}, "test message")
		if !result {
			t.Error("Assertions.NoGoRoutineLeakf should return true on success")
		}
	})
}

func TestAssertionsNotContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotContainsf([]string{"A", "B"}, "C", "test message")
		if !result {
			t.Error("Assertions.NotContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotContainsf([]string{"A", "B"}, "B", "test message")
		if result {
			t.Error("Assertions.NotContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotContainsf should mark test as failed")
		}
	})
}

func TestAssertionsNotElementsMatchf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatchf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if !result {
			t.Error("Assertions.NotElementsMatchf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatchf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if result {
			t.Error("Assertions.NotElementsMatchf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotElementsMatchf should mark test as failed")
		}
	})
}

func TestAssertionsNotEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEmptyf("not empty", "test message")
		if !result {
			t.Error("Assertions.NotEmptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEmptyf("", "test message")
		if result {
			t.Error("Assertions.NotEmptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEmptyf should mark test as failed")
		}
	})
}

func TestAssertionsNotEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualf(123, 456, "test message")
		if !result {
			t.Error("Assertions.NotEqualf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualf(123, 123, "test message")
		if result {
			t.Error("Assertions.NotEqualf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEqualf should mark test as failed")
		}
	})
}

func TestAssertionsNotEqualValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualValuesf(uint32(123), int32(456), "test message")
		if !result {
			t.Error("Assertions.NotEqualValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualValuesf(uint32(123), int32(123), "test message")
		if result {
			t.Error("Assertions.NotEqualValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEqualValuesf should mark test as failed")
		}
	})
}

func TestAssertionsNotErrorAsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAsf(ErrTest, new(*dummyError), "test message")
		if !result {
			t.Error("Assertions.NotErrorAsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAsf(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if result {
			t.Error("Assertions.NotErrorAsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotErrorAsf should mark test as failed")
		}
	})
}

func TestAssertionsNotErrorIsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorIsf(ErrTest, io.EOF, "test message")
		if !result {
			t.Error("Assertions.NotErrorIsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorIsf(fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		if result {
			t.Error("Assertions.NotErrorIsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotErrorIsf should mark test as failed")
		}
	})
}

func TestAssertionsNotImplementsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotImplementsf((*error)(nil), new(testing.T), "test message")
		if !result {
			t.Error("Assertions.NotImplementsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotImplementsf(ptr(dummyInterface), new(testing.T), "test message")
		if result {
			t.Error("Assertions.NotImplementsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotImplementsf should mark test as failed")
		}
	})
}

func TestAssertionsNotKindf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotKindf(reflect.String, 0, "test message")
		if !result {
			t.Error("Assertions.NotKindf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotKindf(reflect.String, "hello", "test message")
		if result {
			t.Error("Assertions.NotKindf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotKindf should mark test as failed")
		}
	})
}

func TestAssertionsNotNilf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotNilf("not nil", "test message")
		if !result {
			t.Error("Assertions.NotNilf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotNilf(nil, "test message")
		if result {
			t.Error("Assertions.NotNilf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotNilf should mark test as failed")
		}
	})
}

func TestAssertionsNotPanicsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotPanicsf(func() {}, "test message")
		if !result {
			t.Error("Assertions.NotPanicsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotPanicsf(func() { panic("panicking") }, "test message")
		if result {
			t.Error("Assertions.NotPanicsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotPanicsf should mark test as failed")
		}
	})
}

func TestAssertionsNotRegexpf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexpf("^start", "not starting", "test message")
		if !result {
			t.Error("Assertions.NotRegexpf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexpf("^start", "starting", "test message")
		if result {
			t.Error("Assertions.NotRegexpf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotRegexpf should mark test as failed")
		}
	})
}

func TestAssertionsNotSamef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSamef(&staticVar, ptr("static string"), "test message")
		if !result {
			t.Error("Assertions.NotSamef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSamef(&staticVar, staticVarPtr, "test message")
		if result {
			t.Error("Assertions.NotSamef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSamef should mark test as failed")
		}
	})
}

func TestAssertionsNotSubsetf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSubsetf([]int{1, 2, 3}, []int{4, 5}, "test message")
		if !result {
			t.Error("Assertions.NotSubsetf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSubsetf([]int{1, 2, 3}, []int{1, 2}, "test message")
		if result {
			t.Error("Assertions.NotSubsetf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSubsetf should mark test as failed")
		}
	})
}

func TestAssertionsNotZerof(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotZerof(1, "test message")
		if !result {
			t.Error("Assertions.NotZerof should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotZerof(0, "test message")
		if result {
			t.Error("Assertions.NotZerof should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotZerof should mark test as failed")
		}
	})
}

func TestAssertionsPanicsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panicsf(func() { panic("panicking") }, "test message")
		if !result {
			t.Error("Assertions.Panicsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panicsf(func() {}, "test message")
		if result {
			t.Error("Assertions.Panicsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Panicsf should mark test as failed")
		}
	})
}

func TestAssertionsPanicsWithErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithErrorf(ErrTest.Error(), func() { panic(ErrTest) }, "test message")
		if !result {
			t.Error("Assertions.PanicsWithErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithErrorf(ErrTest.Error(), func() {}, "test message")
		if result {
			t.Error("Assertions.PanicsWithErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.PanicsWithErrorf should mark test as failed")
		}
	})
}

func TestAssertionsPanicsWithValuef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithValuef("panicking", func() { panic("panicking") }, "test message")
		if !result {
			t.Error("Assertions.PanicsWithValuef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PanicsWithValuef("panicking", func() {}, "test message")
		if result {
			t.Error("Assertions.PanicsWithValuef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.PanicsWithValuef should mark test as failed")
		}
	})
}

func TestAssertionsPositivef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Positivef(1, "test message")
		if !result {
			t.Error("Assertions.Positivef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Positivef(-1, "test message")
		if result {
			t.Error("Assertions.Positivef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Positivef should mark test as failed")
		}
	})
}

func TestAssertionsRegexpf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Regexpf("^start", "starting", "test message")
		if !result {
			t.Error("Assertions.Regexpf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Regexpf("^start", "not starting", "test message")
		if result {
			t.Error("Assertions.Regexpf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Regexpf should mark test as failed")
		}
	})
}

func TestAssertionsSamef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Samef(&staticVar, staticVarPtr, "test message")
		if !result {
			t.Error("Assertions.Samef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Samef(&staticVar, ptr("static string"), "test message")
		if result {
			t.Error("Assertions.Samef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Samef should mark test as failed")
		}
	})
}

func TestAssertionsSubsetf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Subsetf([]int{1, 2, 3}, []int{1, 2}, "test message")
		if !result {
			t.Error("Assertions.Subsetf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Subsetf([]int{1, 2, 3}, []int{4, 5}, "test message")
		if result {
			t.Error("Assertions.Subsetf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Subsetf should mark test as failed")
		}
	})
}

func TestAssertionsTruef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Truef(1 == 1, "test message")
		if !result {
			t.Error("Assertions.Truef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Truef(1 == 0, "test message")
		if result {
			t.Error("Assertions.Truef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Truef should mark test as failed")
		}
	})
}

func TestAssertionsWithinDurationf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinDurationf(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second, "test message")
		if !result {
			t.Error("Assertions.WithinDurationf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinDurationf(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second, "test message")
		if result {
			t.Error("Assertions.WithinDurationf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.WithinDurationf should mark test as failed")
		}
	})
}

func TestAssertionsWithinRangef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinRangef(time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		if !result {
			t.Error("Assertions.WithinRangef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.WithinRangef(time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		if result {
			t.Error("Assertions.WithinRangef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.WithinRangef should mark test as failed")
		}
	})
}

func TestAssertionsYAMLEqf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLEqf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLEqf should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLEqf should panic as expected")
		}
	})
}

func TestAssertionsYAMLEqBytesf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLEqBytesf([]byte("key: value"), []byte("key: value"), "test message")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLEqBytesf should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLEqBytesf should panic as expected")
		}
	})
}

func TestAssertionsZerof(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Zerof(0, "test message")
		if !result {
			t.Error("Assertions.Zerof should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Zerof(1, "test message")
		if result {
			t.Error("Assertions.Zerof should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Zerof should mark test as failed")
		}
	})
}
