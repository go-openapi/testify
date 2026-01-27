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
	"slices"
	"testing"
	"time"
)

func TestCondition(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Condition(t, func() bool { return true })
		if !result {
			t.Error("Condition should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Condition(mock, func() bool { return false })
		if result {
			t.Error("Condition should return false on failure")
		}
		if !mock.failed {
			t.Error("Condition should mark test as failed")
		}
	})
}

func TestContains(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Contains(t, []string{"A", "B"}, "A")
		if !result {
			t.Error("Contains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Contains(mock, []string{"A", "B"}, "C")
		if result {
			t.Error("Contains should return false on failure")
		}
		if !mock.failed {
			t.Error("Contains should mark test as failed")
		}
	})
}

func TestDirExists(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := DirExists(t, filepath.Join(testDataPath(), "existing_dir"))
		if !result {
			t.Error("DirExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := DirExists(mock, filepath.Join(testDataPath(), "non_existing_dir"))
		if result {
			t.Error("DirExists should return false on failure")
		}
		if !mock.failed {
			t.Error("DirExists should mark test as failed")
		}
	})
}

func TestDirNotExists(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := DirNotExists(t, filepath.Join(testDataPath(), "non_existing_dir"))
		if !result {
			t.Error("DirNotExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := DirNotExists(mock, filepath.Join(testDataPath(), "existing_dir"))
		if result {
			t.Error("DirNotExists should return false on failure")
		}
		if !mock.failed {
			t.Error("DirNotExists should mark test as failed")
		}
	})
}

func TestElementsMatch(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ElementsMatch(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if !result {
			t.Error("ElementsMatch should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ElementsMatch(mock, []int{1, 2, 3}, []int{1, 2, 4})
		if result {
			t.Error("ElementsMatch should return false on failure")
		}
		if !mock.failed {
			t.Error("ElementsMatch should mark test as failed")
		}
	})
}

func TestElementsMatchT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ElementsMatchT(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if !result {
			t.Error("ElementsMatchT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ElementsMatchT(mock, []int{1, 2, 3}, []int{1, 2, 4})
		if result {
			t.Error("ElementsMatchT should return false on failure")
		}
		if !mock.failed {
			t.Error("ElementsMatchT should mark test as failed")
		}
	})
}

func TestEmpty(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Empty(t, "")
		if !result {
			t.Error("Empty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Empty(mock, "not empty")
		if result {
			t.Error("Empty should return false on failure")
		}
		if !mock.failed {
			t.Error("Empty should mark test as failed")
		}
	})
}

func TestEqual(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Equal(t, 123, 123)
		if !result {
			t.Error("Equal should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Equal(mock, 123, 456)
		if result {
			t.Error("Equal should return false on failure")
		}
		if !mock.failed {
			t.Error("Equal should mark test as failed")
		}
	})
}

func TestEqualError(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualError(t, ErrTest, "assert.ErrTest general error for testing")
		if !result {
			t.Error("EqualError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualError(mock, ErrTest, "wrong error message")
		if result {
			t.Error("EqualError should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualError should mark test as failed")
		}
	})
}

func TestEqualExportedValues(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualExportedValues(t, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
		if !result {
			t.Error("EqualExportedValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualExportedValues(mock, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1})
		if result {
			t.Error("EqualExportedValues should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualExportedValues should mark test as failed")
		}
	})
}

func TestEqualT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualT(t, 123, 123)
		if !result {
			t.Error("EqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualT(mock, 123, 456)
		if result {
			t.Error("EqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualT should mark test as failed")
		}
	})
}

func TestEqualValues(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualValues(t, uint32(123), int32(123))
		if !result {
			t.Error("EqualValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualValues(mock, uint32(123), int32(456))
		if result {
			t.Error("EqualValues should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualValues should mark test as failed")
		}
	})
}

func TestError(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Error(t, ErrTest)
		if !result {
			t.Error("Error should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Error(mock, nil)
		if result {
			t.Error("Error should return false on failure")
		}
		if !mock.failed {
			t.Error("Error should mark test as failed")
		}
	})
}

func TestErrorAs(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ErrorAs(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if !result {
			t.Error("ErrorAs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorAs(mock, ErrTest, new(*dummyError))
		if result {
			t.Error("ErrorAs should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorAs should mark test as failed")
		}
	})
}

func TestErrorContains(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ErrorContains(t, ErrTest, "general error")
		if !result {
			t.Error("ErrorContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorContains(mock, ErrTest, "not in message")
		if result {
			t.Error("ErrorContains should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorContains should mark test as failed")
		}
	})
}

func TestErrorIs(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ErrorIs(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		if !result {
			t.Error("ErrorIs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorIs(mock, ErrTest, io.EOF)
		if result {
			t.Error("ErrorIs should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorIs should mark test as failed")
		}
	})
}

func TestEventually(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Eventually(t, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		if !result {
			t.Error("Eventually should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Eventually(mock, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("Eventually should return false on failure")
		}
		if !mock.failed {
			t.Error("Eventually should mark test as failed")
		}
	})
}

func TestEventuallyWith(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EventuallyWith(t, func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		if !result {
			t.Error("EventuallyWith should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EventuallyWith(mock, func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("EventuallyWith should return false on failure")
		}
		if !mock.failed {
			t.Error("EventuallyWith should mark test as failed")
		}
	})
}

func TestExactly(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Exactly(t, int32(123), int32(123))
		if !result {
			t.Error("Exactly should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Exactly(mock, int32(123), int64(123))
		if result {
			t.Error("Exactly should return false on failure")
		}
		if !mock.failed {
			t.Error("Exactly should mark test as failed")
		}
	})
}

func TestFail(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Fail(mock, "failed")
		if result {
			t.Error("Fail should return false on failure")
		}
		if !mock.failed {
			t.Error("Fail should mark test as failed")
		}
	})
}

func TestFailNow(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		result := FailNow(mock, "failed")
		if result {
			t.Error("FailNow should return false on failure")
		}
		if !mock.failed {
			t.Error("FailNow should call FailNow()")
		}
	})
}

func TestFalse(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := False(t, 1 == 0)
		if !result {
			t.Error("False should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := False(mock, 1 == 1)
		if result {
			t.Error("False should return false on failure")
		}
		if !mock.failed {
			t.Error("False should mark test as failed")
		}
	})
}

func TestFalseT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FalseT(t, 1 == 0)
		if !result {
			t.Error("FalseT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FalseT(mock, 1 == 1)
		if result {
			t.Error("FalseT should return false on failure")
		}
		if !mock.failed {
			t.Error("FalseT should mark test as failed")
		}
	})
}

func TestFileEmpty(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileEmpty(t, filepath.Join(testDataPath(), "empty_file"))
		if !result {
			t.Error("FileEmpty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileEmpty(mock, filepath.Join(testDataPath(), "existing_file"))
		if result {
			t.Error("FileEmpty should return false on failure")
		}
		if !mock.failed {
			t.Error("FileEmpty should mark test as failed")
		}
	})
}

func TestFileExists(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileExists(t, filepath.Join(testDataPath(), "existing_file"))
		if !result {
			t.Error("FileExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileExists(mock, filepath.Join(testDataPath(), "non_existing_file"))
		if result {
			t.Error("FileExists should return false on failure")
		}
		if !mock.failed {
			t.Error("FileExists should mark test as failed")
		}
	})
}

func TestFileNotEmpty(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileNotEmpty(t, filepath.Join(testDataPath(), "existing_file"))
		if !result {
			t.Error("FileNotEmpty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileNotEmpty(mock, filepath.Join(testDataPath(), "empty_file"))
		if result {
			t.Error("FileNotEmpty should return false on failure")
		}
		if !mock.failed {
			t.Error("FileNotEmpty should mark test as failed")
		}
	})
}

func TestFileNotExists(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileNotExists(t, filepath.Join(testDataPath(), "non_existing_file"))
		if !result {
			t.Error("FileNotExists should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileNotExists(mock, filepath.Join(testDataPath(), "existing_file"))
		if result {
			t.Error("FileNotExists should return false on failure")
		}
		if !mock.failed {
			t.Error("FileNotExists should mark test as failed")
		}
	})
}

func TestGreater(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Greater(t, 2, 1)
		if !result {
			t.Error("Greater should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Greater(mock, 1, 2)
		if result {
			t.Error("Greater should return false on failure")
		}
		if !mock.failed {
			t.Error("Greater should mark test as failed")
		}
	})
}

func TestGreaterOrEqual(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := GreaterOrEqual(t, 2, 1)
		if !result {
			t.Error("GreaterOrEqual should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := GreaterOrEqual(mock, 1, 2)
		if result {
			t.Error("GreaterOrEqual should return false on failure")
		}
		if !mock.failed {
			t.Error("GreaterOrEqual should mark test as failed")
		}
	})
}

func TestGreaterOrEqualT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := GreaterOrEqualT(t, 2, 1)
		if !result {
			t.Error("GreaterOrEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := GreaterOrEqualT(mock, 1, 2)
		if result {
			t.Error("GreaterOrEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("GreaterOrEqualT should mark test as failed")
		}
	})
}

func TestGreaterT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := GreaterT(t, 2, 1)
		if !result {
			t.Error("GreaterT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := GreaterT(mock, 1, 2)
		if result {
			t.Error("GreaterT should return false on failure")
		}
		if !mock.failed {
			t.Error("GreaterT should mark test as failed")
		}
	})
}

func TestHTTPBodyContains(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPBodyContains(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!")
		if !result {
			t.Error("HTTPBodyContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPBodyContains(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!")
		if result {
			t.Error("HTTPBodyContains should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPBodyContains should mark test as failed")
		}
	})
}

func TestHTTPBodyNotContains(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPBodyNotContains(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!")
		if !result {
			t.Error("HTTPBodyNotContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPBodyNotContains(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!")
		if result {
			t.Error("HTTPBodyNotContains should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPBodyNotContains should mark test as failed")
		}
	})
}

func TestHTTPError(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPError(t, httpError, "GET", "/", nil)
		if !result {
			t.Error("HTTPError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPError(mock, httpOK, "GET", "/", nil)
		if result {
			t.Error("HTTPError should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPError should mark test as failed")
		}
	})
}

func TestHTTPRedirect(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPRedirect(t, httpRedirect, "GET", "/", nil)
		if !result {
			t.Error("HTTPRedirect should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPRedirect(mock, httpError, "GET", "/", nil)
		if result {
			t.Error("HTTPRedirect should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPRedirect should mark test as failed")
		}
	})
}

func TestHTTPStatusCode(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPStatusCode(t, httpOK, "GET", "/", nil, http.StatusOK)
		if !result {
			t.Error("HTTPStatusCode should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPStatusCode(mock, httpError, "GET", "/", nil, http.StatusOK)
		if result {
			t.Error("HTTPStatusCode should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPStatusCode should mark test as failed")
		}
	})
}

func TestHTTPSuccess(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPSuccess(t, httpOK, "GET", "/", nil)
		if !result {
			t.Error("HTTPSuccess should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPSuccess(mock, httpError, "GET", "/", nil)
		if result {
			t.Error("HTTPSuccess should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPSuccess should mark test as failed")
		}
	})
}

func TestImplements(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Implements(t, ptr(dummyInterface), new(testing.T))
		if !result {
			t.Error("Implements should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Implements(mock, (*error)(nil), new(testing.T))
		if result {
			t.Error("Implements should return false on failure")
		}
		if !mock.failed {
			t.Error("Implements should mark test as failed")
		}
	})
}

func TestInDelta(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDelta(t, 1.0, 1.01, 0.02)
		if !result {
			t.Error("InDelta should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDelta(mock, 1.0, 1.1, 0.05)
		if result {
			t.Error("InDelta should return false on failure")
		}
		if !mock.failed {
			t.Error("InDelta should mark test as failed")
		}
	})
}

func TestInDeltaMapValues(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDeltaMapValues(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
		if !result {
			t.Error("InDeltaMapValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaMapValues(mock, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05)
		if result {
			t.Error("InDeltaMapValues should return false on failure")
		}
		if !mock.failed {
			t.Error("InDeltaMapValues should mark test as failed")
		}
	})
}

func TestInDeltaSlice(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDeltaSlice(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
		if !result {
			t.Error("InDeltaSlice should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaSlice(mock, []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05)
		if result {
			t.Error("InDeltaSlice should return false on failure")
		}
		if !mock.failed {
			t.Error("InDeltaSlice should mark test as failed")
		}
	})
}

func TestInDeltaT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDeltaT(t, 1.0, 1.01, 0.02)
		if !result {
			t.Error("InDeltaT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaT(mock, 1.0, 1.1, 0.05)
		if result {
			t.Error("InDeltaT should return false on failure")
		}
		if !mock.failed {
			t.Error("InDeltaT should mark test as failed")
		}
	})
}

func TestInEpsilon(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InEpsilon(t, 100.0, 101.0, 0.02)
		if !result {
			t.Error("InEpsilon should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilon(mock, 100.0, 110.0, 0.05)
		if result {
			t.Error("InEpsilon should return false on failure")
		}
		if !mock.failed {
			t.Error("InEpsilon should mark test as failed")
		}
	})
}

func TestInEpsilonSlice(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InEpsilonSlice(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
		if !result {
			t.Error("InEpsilonSlice should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilonSlice(mock, []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05)
		if result {
			t.Error("InEpsilonSlice should return false on failure")
		}
		if !mock.failed {
			t.Error("InEpsilonSlice should mark test as failed")
		}
	})
}

func TestInEpsilonT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InEpsilonT(t, 100.0, 101.0, 0.02)
		if !result {
			t.Error("InEpsilonT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilonT(mock, 100.0, 110.0, 0.05)
		if result {
			t.Error("InEpsilonT should return false on failure")
		}
		if !mock.failed {
			t.Error("InEpsilonT should mark test as failed")
		}
	})
}

func TestIsDecreasing(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsDecreasing(t, []int{3, 2, 1})
		if !result {
			t.Error("IsDecreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsDecreasing(mock, []int{1, 2, 3})
		if result {
			t.Error("IsDecreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("IsDecreasing should mark test as failed")
		}
	})
}

func TestIsDecreasingT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsDecreasingT(t, []int{3, 2, 1})
		if !result {
			t.Error("IsDecreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsDecreasingT(mock, []int{1, 2, 3})
		if result {
			t.Error("IsDecreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("IsDecreasingT should mark test as failed")
		}
	})
}

func TestIsIncreasing(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsIncreasing(t, []int{1, 2, 3})
		if !result {
			t.Error("IsIncreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsIncreasing(mock, []int{1, 1, 2})
		if result {
			t.Error("IsIncreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("IsIncreasing should mark test as failed")
		}
	})
}

func TestIsIncreasingT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsIncreasingT(t, []int{1, 2, 3})
		if !result {
			t.Error("IsIncreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsIncreasingT(mock, []int{1, 1, 2})
		if result {
			t.Error("IsIncreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("IsIncreasingT should mark test as failed")
		}
	})
}

func TestIsNonDecreasing(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonDecreasing(t, []int{1, 1, 2})
		if !result {
			t.Error("IsNonDecreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonDecreasing(mock, []int{2, 1, 0})
		if result {
			t.Error("IsNonDecreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonDecreasing should mark test as failed")
		}
	})
}

func TestIsNonDecreasingT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonDecreasingT(t, []int{1, 1, 2})
		if !result {
			t.Error("IsNonDecreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonDecreasingT(mock, []int{2, 1, 0})
		if result {
			t.Error("IsNonDecreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonDecreasingT should mark test as failed")
		}
	})
}

func TestIsNonIncreasing(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonIncreasing(t, []int{2, 1, 1})
		if !result {
			t.Error("IsNonIncreasing should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonIncreasing(mock, []int{1, 2, 3})
		if result {
			t.Error("IsNonIncreasing should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonIncreasing should mark test as failed")
		}
	})
}

func TestIsNonIncreasingT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonIncreasingT(t, []int{2, 1, 1})
		if !result {
			t.Error("IsNonIncreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonIncreasingT(mock, []int{1, 2, 3})
		if result {
			t.Error("IsNonIncreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonIncreasingT should mark test as failed")
		}
	})
}

func TestIsNotOfTypeT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNotOfTypeT[myType](t, 123.123)
		if !result {
			t.Error("IsNotOfTypeT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNotOfTypeT[myType](mock, myType(123.123))
		if result {
			t.Error("IsNotOfTypeT should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNotOfTypeT should mark test as failed")
		}
	})
}

func TestIsNotType(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNotType(t, int32(123), int64(456))
		if !result {
			t.Error("IsNotType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNotType(mock, 123, 456)
		if result {
			t.Error("IsNotType should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNotType should mark test as failed")
		}
	})
}

func TestIsOfTypeT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsOfTypeT[myType](t, myType(123.123))
		if !result {
			t.Error("IsOfTypeT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsOfTypeT[myType](mock, 123.123)
		if result {
			t.Error("IsOfTypeT should return false on failure")
		}
		if !mock.failed {
			t.Error("IsOfTypeT should mark test as failed")
		}
	})
}

func TestIsType(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsType(t, 123, 456)
		if !result {
			t.Error("IsType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsType(mock, int32(123), int64(456))
		if result {
			t.Error("IsType should return false on failure")
		}
		if !mock.failed {
			t.Error("IsType should mark test as failed")
		}
	})
}

func TestJSONEq(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := JSONEq(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
		if !result {
			t.Error("JSONEq should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := JSONEq(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		if result {
			t.Error("JSONEq should return false on failure")
		}
		if !mock.failed {
			t.Error("JSONEq should mark test as failed")
		}
	})
}

func TestJSONEqBytes(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := JSONEqBytes(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
		if !result {
			t.Error("JSONEqBytes should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := JSONEqBytes(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`))
		if result {
			t.Error("JSONEqBytes should return false on failure")
		}
		if !mock.failed {
			t.Error("JSONEqBytes should mark test as failed")
		}
	})
}

func TestJSONEqT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := JSONEqT(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
		if !result {
			t.Error("JSONEqT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := JSONEqT(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		if result {
			t.Error("JSONEqT should return false on failure")
		}
		if !mock.failed {
			t.Error("JSONEqT should mark test as failed")
		}
	})
}

func TestKind(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Kind(t, reflect.String, "hello")
		if !result {
			t.Error("Kind should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Kind(mock, reflect.String, 0)
		if result {
			t.Error("Kind should return false on failure")
		}
		if !mock.failed {
			t.Error("Kind should mark test as failed")
		}
	})
}

func TestLen(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Len(t, []string{"A", "B"}, 2)
		if !result {
			t.Error("Len should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Len(mock, []string{"A", "B"}, 1)
		if result {
			t.Error("Len should return false on failure")
		}
		if !mock.failed {
			t.Error("Len should mark test as failed")
		}
	})
}

func TestLess(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Less(t, 1, 2)
		if !result {
			t.Error("Less should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Less(mock, 2, 1)
		if result {
			t.Error("Less should return false on failure")
		}
		if !mock.failed {
			t.Error("Less should mark test as failed")
		}
	})
}

func TestLessOrEqual(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := LessOrEqual(t, 1, 2)
		if !result {
			t.Error("LessOrEqual should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := LessOrEqual(mock, 2, 1)
		if result {
			t.Error("LessOrEqual should return false on failure")
		}
		if !mock.failed {
			t.Error("LessOrEqual should mark test as failed")
		}
	})
}

func TestLessOrEqualT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := LessOrEqualT(t, 1, 2)
		if !result {
			t.Error("LessOrEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := LessOrEqualT(mock, 2, 1)
		if result {
			t.Error("LessOrEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("LessOrEqualT should mark test as failed")
		}
	})
}

func TestLessT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := LessT(t, 1, 2)
		if !result {
			t.Error("LessT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := LessT(mock, 2, 1)
		if result {
			t.Error("LessT should return false on failure")
		}
		if !mock.failed {
			t.Error("LessT should mark test as failed")
		}
	})
}

func TestMapContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := MapContainsT(t, map[string]string{"A": "B"}, "A")
		if !result {
			t.Error("MapContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := MapContainsT(mock, map[string]string{"A": "B"}, "C")
		if result {
			t.Error("MapContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("MapContainsT should mark test as failed")
		}
	})
}

func TestMapNotContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := MapNotContainsT(t, map[string]string{"A": "B"}, "C")
		if !result {
			t.Error("MapNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := MapNotContainsT(mock, map[string]string{"A": "B"}, "A")
		if result {
			t.Error("MapNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("MapNotContainsT should mark test as failed")
		}
	})
}

func TestNegative(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Negative(t, -1)
		if !result {
			t.Error("Negative should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Negative(mock, 1)
		if result {
			t.Error("Negative should return false on failure")
		}
		if !mock.failed {
			t.Error("Negative should mark test as failed")
		}
	})
}

func TestNegativeT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NegativeT(t, -1)
		if !result {
			t.Error("NegativeT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NegativeT(mock, 1)
		if result {
			t.Error("NegativeT should return false on failure")
		}
		if !mock.failed {
			t.Error("NegativeT should mark test as failed")
		}
	})
}

func TestNever(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Never(t, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		if !result {
			t.Error("Never should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Never(mock, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("Never should return false on failure")
		}
		if !mock.failed {
			t.Error("Never should mark test as failed")
		}
	})
}

func TestNil(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Nil(t, nil)
		if !result {
			t.Error("Nil should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Nil(mock, "not nil")
		if result {
			t.Error("Nil should return false on failure")
		}
		if !mock.failed {
			t.Error("Nil should mark test as failed")
		}
	})
}

func TestNoError(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NoError(t, nil)
		if !result {
			t.Error("NoError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NoError(mock, ErrTest)
		if result {
			t.Error("NoError should return false on failure")
		}
		if !mock.failed {
			t.Error("NoError should mark test as failed")
		}
	})
}

func TestNotContains(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotContains(t, []string{"A", "B"}, "C")
		if !result {
			t.Error("NotContains should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotContains(mock, []string{"A", "B"}, "B")
		if result {
			t.Error("NotContains should return false on failure")
		}
		if !mock.failed {
			t.Error("NotContains should mark test as failed")
		}
	})
}

func TestNotElementsMatch(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotElementsMatch(t, []int{1, 2, 3}, []int{1, 2, 4})
		if !result {
			t.Error("NotElementsMatch should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotElementsMatch(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if result {
			t.Error("NotElementsMatch should return false on failure")
		}
		if !mock.failed {
			t.Error("NotElementsMatch should mark test as failed")
		}
	})
}

func TestNotElementsMatchT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotElementsMatchT(t, []int{1, 2, 3}, []int{1, 2, 4})
		if !result {
			t.Error("NotElementsMatchT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotElementsMatchT(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if result {
			t.Error("NotElementsMatchT should return false on failure")
		}
		if !mock.failed {
			t.Error("NotElementsMatchT should mark test as failed")
		}
	})
}

func TestNotEmpty(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEmpty(t, "not empty")
		if !result {
			t.Error("NotEmpty should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEmpty(mock, "")
		if result {
			t.Error("NotEmpty should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEmpty should mark test as failed")
		}
	})
}

func TestNotEqual(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEqual(t, 123, 456)
		if !result {
			t.Error("NotEqual should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEqual(mock, 123, 123)
		if result {
			t.Error("NotEqual should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEqual should mark test as failed")
		}
	})
}

func TestNotEqualT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEqualT(t, 123, 456)
		if !result {
			t.Error("NotEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEqualT(mock, 123, 123)
		if result {
			t.Error("NotEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEqualT should mark test as failed")
		}
	})
}

func TestNotEqualValues(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEqualValues(t, uint32(123), int32(456))
		if !result {
			t.Error("NotEqualValues should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEqualValues(mock, uint32(123), int32(123))
		if result {
			t.Error("NotEqualValues should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEqualValues should mark test as failed")
		}
	})
}

func TestNotErrorAs(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotErrorAs(t, ErrTest, new(*dummyError))
		if !result {
			t.Error("NotErrorAs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorAs(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if result {
			t.Error("NotErrorAs should return false on failure")
		}
		if !mock.failed {
			t.Error("NotErrorAs should mark test as failed")
		}
	})
}

func TestNotErrorIs(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotErrorIs(t, ErrTest, io.EOF)
		if !result {
			t.Error("NotErrorIs should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorIs(mock, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		if result {
			t.Error("NotErrorIs should return false on failure")
		}
		if !mock.failed {
			t.Error("NotErrorIs should mark test as failed")
		}
	})
}

func TestNotImplements(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotImplements(t, (*error)(nil), new(testing.T))
		if !result {
			t.Error("NotImplements should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotImplements(mock, ptr(dummyInterface), new(testing.T))
		if result {
			t.Error("NotImplements should return false on failure")
		}
		if !mock.failed {
			t.Error("NotImplements should mark test as failed")
		}
	})
}

func TestNotKind(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotKind(t, reflect.String, 0)
		if !result {
			t.Error("NotKind should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotKind(mock, reflect.String, "hello")
		if result {
			t.Error("NotKind should return false on failure")
		}
		if !mock.failed {
			t.Error("NotKind should mark test as failed")
		}
	})
}

func TestNotNil(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotNil(t, "not nil")
		if !result {
			t.Error("NotNil should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotNil(mock, nil)
		if result {
			t.Error("NotNil should return false on failure")
		}
		if !mock.failed {
			t.Error("NotNil should mark test as failed")
		}
	})
}

func TestNotPanics(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotPanics(t, func() {})
		if !result {
			t.Error("NotPanics should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotPanics(mock, func() { panic("panicking") })
		if result {
			t.Error("NotPanics should return false on failure")
		}
		if !mock.failed {
			t.Error("NotPanics should mark test as failed")
		}
	})
}

func TestNotRegexp(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotRegexp(t, "^start", "not starting")
		if !result {
			t.Error("NotRegexp should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotRegexp(mock, "^start", "starting")
		if result {
			t.Error("NotRegexp should return false on failure")
		}
		if !mock.failed {
			t.Error("NotRegexp should mark test as failed")
		}
	})
}

func TestNotRegexpT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotRegexpT(t, "^start", "not starting")
		if !result {
			t.Error("NotRegexpT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotRegexpT(mock, "^start", "starting")
		if result {
			t.Error("NotRegexpT should return false on failure")
		}
		if !mock.failed {
			t.Error("NotRegexpT should mark test as failed")
		}
	})
}

func TestNotSame(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSame(t, &staticVar, ptr("static string"))
		if !result {
			t.Error("NotSame should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSame(mock, &staticVar, staticVarPtr)
		if result {
			t.Error("NotSame should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSame should mark test as failed")
		}
	})
}

func TestNotSameT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSameT(t, &staticVar, ptr("static string"))
		if !result {
			t.Error("NotSameT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSameT(mock, &staticVar, staticVarPtr)
		if result {
			t.Error("NotSameT should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSameT should mark test as failed")
		}
	})
}

func TestNotSortedT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSortedT(t, []int{3, 1, 3})
		if !result {
			t.Error("NotSortedT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSortedT(mock, []int{1, 4, 8})
		if result {
			t.Error("NotSortedT should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSortedT should mark test as failed")
		}
	})
}

func TestNotSubset(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSubset(t, []int{1, 2, 3}, []int{4, 5})
		if !result {
			t.Error("NotSubset should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSubset(mock, []int{1, 2, 3}, []int{1, 2})
		if result {
			t.Error("NotSubset should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSubset should mark test as failed")
		}
	})
}

func TestNotZero(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotZero(t, 1)
		if !result {
			t.Error("NotZero should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotZero(mock, 0)
		if result {
			t.Error("NotZero should return false on failure")
		}
		if !mock.failed {
			t.Error("NotZero should mark test as failed")
		}
	})
}

func TestPanics(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Panics(t, func() { panic("panicking") })
		if !result {
			t.Error("Panics should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Panics(mock, func() {})
		if result {
			t.Error("Panics should return false on failure")
		}
		if !mock.failed {
			t.Error("Panics should mark test as failed")
		}
	})
}

func TestPanicsWithError(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := PanicsWithError(t, ErrTest.Error(), func() { panic(ErrTest) })
		if !result {
			t.Error("PanicsWithError should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := PanicsWithError(mock, ErrTest.Error(), func() {})
		if result {
			t.Error("PanicsWithError should return false on failure")
		}
		if !mock.failed {
			t.Error("PanicsWithError should mark test as failed")
		}
	})
}

func TestPanicsWithValue(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := PanicsWithValue(t, "panicking", func() { panic("panicking") })
		if !result {
			t.Error("PanicsWithValue should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := PanicsWithValue(mock, "panicking", func() {})
		if result {
			t.Error("PanicsWithValue should return false on failure")
		}
		if !mock.failed {
			t.Error("PanicsWithValue should mark test as failed")
		}
	})
}

func TestPositive(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Positive(t, 1)
		if !result {
			t.Error("Positive should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Positive(mock, -1)
		if result {
			t.Error("Positive should return false on failure")
		}
		if !mock.failed {
			t.Error("Positive should mark test as failed")
		}
	})
}

func TestPositiveT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := PositiveT(t, 1)
		if !result {
			t.Error("PositiveT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := PositiveT(mock, -1)
		if result {
			t.Error("PositiveT should return false on failure")
		}
		if !mock.failed {
			t.Error("PositiveT should mark test as failed")
		}
	})
}

func TestRegexp(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Regexp(t, "^start", "starting")
		if !result {
			t.Error("Regexp should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Regexp(mock, "^start", "not starting")
		if result {
			t.Error("Regexp should return false on failure")
		}
		if !mock.failed {
			t.Error("Regexp should mark test as failed")
		}
	})
}

func TestRegexpT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := RegexpT(t, "^start", "starting")
		if !result {
			t.Error("RegexpT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := RegexpT(mock, "^start", "not starting")
		if result {
			t.Error("RegexpT should return false on failure")
		}
		if !mock.failed {
			t.Error("RegexpT should mark test as failed")
		}
	})
}

func TestSame(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Same(t, &staticVar, staticVarPtr)
		if !result {
			t.Error("Same should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Same(mock, &staticVar, ptr("static string"))
		if result {
			t.Error("Same should return false on failure")
		}
		if !mock.failed {
			t.Error("Same should mark test as failed")
		}
	})
}

func TestSameT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SameT(t, &staticVar, staticVarPtr)
		if !result {
			t.Error("SameT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SameT(mock, &staticVar, ptr("static string"))
		if result {
			t.Error("SameT should return false on failure")
		}
		if !mock.failed {
			t.Error("SameT should mark test as failed")
		}
	})
}

func TestSeqContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SeqContainsT(t, slices.Values([]string{"A", "B"}), "A")
		if !result {
			t.Error("SeqContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SeqContainsT(mock, slices.Values([]string{"A", "B"}), "C")
		if result {
			t.Error("SeqContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("SeqContainsT should mark test as failed")
		}
	})
}

func TestSeqNotContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SeqNotContainsT(t, slices.Values([]string{"A", "B"}), "C")
		if !result {
			t.Error("SeqNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SeqNotContainsT(mock, slices.Values([]string{"A", "B"}), "A")
		if result {
			t.Error("SeqNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("SeqNotContainsT should mark test as failed")
		}
	})
}

func TestSliceContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceContainsT(t, []string{"A", "B"}, "A")
		if !result {
			t.Error("SliceContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceContainsT(mock, []string{"A", "B"}, "C")
		if result {
			t.Error("SliceContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceContainsT should mark test as failed")
		}
	})
}

func TestSliceNotContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceNotContainsT(t, []string{"A", "B"}, "C")
		if !result {
			t.Error("SliceNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceNotContainsT(mock, []string{"A", "B"}, "A")
		if result {
			t.Error("SliceNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceNotContainsT should mark test as failed")
		}
	})
}

func TestSliceNotSubsetT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceNotSubsetT(t, []int{1, 2, 3}, []int{4, 5})
		if !result {
			t.Error("SliceNotSubsetT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceNotSubsetT(mock, []int{1, 2, 3}, []int{1, 2})
		if result {
			t.Error("SliceNotSubsetT should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceNotSubsetT should mark test as failed")
		}
	})
}

func TestSliceSubsetT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceSubsetT(t, []int{1, 2, 3}, []int{1, 2})
		if !result {
			t.Error("SliceSubsetT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceSubsetT(mock, []int{1, 2, 3}, []int{4, 5})
		if result {
			t.Error("SliceSubsetT should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceSubsetT should mark test as failed")
		}
	})
}

func TestSortedT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SortedT(t, []int{1, 1, 3})
		if !result {
			t.Error("SortedT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SortedT(mock, []int{1, 4, 2})
		if result {
			t.Error("SortedT should return false on failure")
		}
		if !mock.failed {
			t.Error("SortedT should mark test as failed")
		}
	})
}

func TestStringContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := StringContainsT(t, "AB", "A")
		if !result {
			t.Error("StringContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := StringContainsT(mock, "AB", "C")
		if result {
			t.Error("StringContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("StringContainsT should mark test as failed")
		}
	})
}

func TestStringNotContainsT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := StringNotContainsT(t, "AB", "C")
		if !result {
			t.Error("StringNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := StringNotContainsT(mock, "AB", "A")
		if result {
			t.Error("StringNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("StringNotContainsT should mark test as failed")
		}
	})
}

func TestSubset(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Subset(t, []int{1, 2, 3}, []int{1, 2})
		if !result {
			t.Error("Subset should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Subset(mock, []int{1, 2, 3}, []int{4, 5})
		if result {
			t.Error("Subset should return false on failure")
		}
		if !mock.failed {
			t.Error("Subset should mark test as failed")
		}
	})
}

func TestTrue(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := True(t, 1 == 1)
		if !result {
			t.Error("True should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := True(mock, 1 == 0)
		if result {
			t.Error("True should return false on failure")
		}
		if !mock.failed {
			t.Error("True should mark test as failed")
		}
	})
}

func TestTrueT(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := TrueT(t, 1 == 1)
		if !result {
			t.Error("TrueT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := TrueT(mock, 1 == 0)
		if result {
			t.Error("TrueT should return false on failure")
		}
		if !mock.failed {
			t.Error("TrueT should mark test as failed")
		}
	})
}

func TestWithinDuration(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := WithinDuration(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second)
		if !result {
			t.Error("WithinDuration should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := WithinDuration(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second)
		if result {
			t.Error("WithinDuration should return false on failure")
		}
		if !mock.failed {
			t.Error("WithinDuration should mark test as failed")
		}
	})
}

func TestWithinRange(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := WithinRange(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		if !result {
			t.Error("WithinRange should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := WithinRange(mock, time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		if result {
			t.Error("WithinRange should return false on failure")
		}
		if !mock.failed {
			t.Error("WithinRange should mark test as failed")
		}
	})
}

func TestYAMLEq(t *testing.T) {
	t.Parallel()
	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		Panics(t, func() {
			YAMLEq(t, "key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
	})
}

func TestYAMLEqBytes(t *testing.T) {
	t.Parallel()
	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		Panics(t, func() {
			YAMLEqBytes(t, []byte("key: value"), []byte("key: value"))
		}, "should panic without the yaml feature enabled.")
	})
}

func TestYAMLEqT(t *testing.T) {
	t.Parallel()
	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		Panics(t, func() {
			YAMLEqT(t, "key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
	})
}

func TestZero(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Zero(t, 0)
		if !result {
			t.Error("Zero should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Zero(mock, 1)
		if result {
			t.Error("Zero should return false on failure")
		}
		if !mock.failed {
			t.Error("Zero should mark test as failed")
		}
	})
}

// mockT is a mock testing.T for assertion tests
type mockT struct {
	failed bool
}

func (m *mockT) Helper() {}

func (m *mockT) Errorf(format string, args ...any) {
	m.failed = true
}

type mockFailNowT struct {
	failed bool
}

// Helper is like [testing.T.Helper] but does nothing.
func (mockFailNowT) Helper() {}

func (m *mockFailNowT) Errorf(format string, args ...any) {
	_ = format
	_ = args
}

func (m *mockFailNowT) FailNow() {
	m.failed = true
}

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
	dummyInterface T
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
