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

func TestConditionf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Conditionf(t, func() bool { return true }, "test message")
		if !result {
			t.Error("Conditionf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Conditionf(mock, func() bool { return false }, "test message")
		if result {
			t.Error("Conditionf should return false on failure")
		}
		if !mock.failed {
			t.Error("Condition should mark test as failed")
		}
	})
}

func TestContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Containsf(t, []string{"A", "B"}, "A", "test message")
		if !result {
			t.Error("Containsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Containsf(mock, []string{"A", "B"}, "C", "test message")
		if result {
			t.Error("Containsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Contains should mark test as failed")
		}
	})
}

func TestDirExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := DirExistsf(t, filepath.Join(testDataPath(), "existing_dir"), "test message")
		if !result {
			t.Error("DirExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := DirExistsf(mock, filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		if result {
			t.Error("DirExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("DirExists should mark test as failed")
		}
	})
}

func TestDirNotExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := DirNotExistsf(t, filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		if !result {
			t.Error("DirNotExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := DirNotExistsf(mock, filepath.Join(testDataPath(), "existing_dir"), "test message")
		if result {
			t.Error("DirNotExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("DirNotExists should mark test as failed")
		}
	})
}

func TestElementsMatchf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ElementsMatchf(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if !result {
			t.Error("ElementsMatchf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ElementsMatchf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if result {
			t.Error("ElementsMatchf should return false on failure")
		}
		if !mock.failed {
			t.Error("ElementsMatch should mark test as failed")
		}
	})
}

func TestElementsMatchTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ElementsMatchTf(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if !result {
			t.Error("ElementsMatchTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ElementsMatchTf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if result {
			t.Error("ElementsMatchTf should return false on failure")
		}
		if !mock.failed {
			t.Error("ElementsMatchT should mark test as failed")
		}
	})
}

func TestEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Emptyf(t, "", "test message")
		if !result {
			t.Error("Emptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Emptyf(mock, "not empty", "test message")
		if result {
			t.Error("Emptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Empty should mark test as failed")
		}
	})
}

func TestEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Equalf(t, 123, 123, "test message")
		if !result {
			t.Error("Equalf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Equalf(mock, 123, 456, "test message")
		if result {
			t.Error("Equalf should return false on failure")
		}
		if !mock.failed {
			t.Error("Equal should mark test as failed")
		}
	})
}

func TestEqualErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualErrorf(t, ErrTest, "assert.ErrTest general error for testing", "test message")
		if !result {
			t.Error("EqualErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualErrorf(mock, ErrTest, "wrong error message", "test message")
		if result {
			t.Error("EqualErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualError should mark test as failed")
		}
	})
}

func TestEqualExportedValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualExportedValuesf(t, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}, "test message")
		if !result {
			t.Error("EqualExportedValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualExportedValuesf(mock, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}, "test message")
		if result {
			t.Error("EqualExportedValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualExportedValues should mark test as failed")
		}
	})
}

func TestEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualTf(t, 123, 123, "test message")
		if !result {
			t.Error("EqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualTf(mock, 123, 456, "test message")
		if result {
			t.Error("EqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualT should mark test as failed")
		}
	})
}

func TestEqualValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EqualValuesf(t, uint32(123), int32(123), "test message")
		if !result {
			t.Error("EqualValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EqualValuesf(mock, uint32(123), int32(456), "test message")
		if result {
			t.Error("EqualValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("EqualValues should mark test as failed")
		}
	})
}

func TestErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Errorf(t, ErrTest, "test message")
		if !result {
			t.Error("Errorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Errorf(mock, nil, "test message")
		if result {
			t.Error("Errorf should return false on failure")
		}
		if !mock.failed {
			t.Error("Error should mark test as failed")
		}
	})
}

func TestErrorAsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ErrorAsf(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if !result {
			t.Error("ErrorAsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorAsf(mock, ErrTest, new(*dummyError), "test message")
		if result {
			t.Error("ErrorAsf should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorAs should mark test as failed")
		}
	})
}

func TestErrorContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ErrorContainsf(t, ErrTest, "general error", "test message")
		if !result {
			t.Error("ErrorContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorContainsf(mock, ErrTest, "not in message", "test message")
		if result {
			t.Error("ErrorContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorContains should mark test as failed")
		}
	})
}

func TestErrorIsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := ErrorIsf(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		if !result {
			t.Error("ErrorIsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorIsf(mock, ErrTest, io.EOF, "test message")
		if result {
			t.Error("ErrorIsf should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorIs should mark test as failed")
		}
	})
}

func TestEventuallyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Eventuallyf(t, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if !result {
			t.Error("Eventuallyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Eventuallyf(mock, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("Eventuallyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Eventually should mark test as failed")
		}
	})
}

func TestEventuallyWithTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := EventuallyWithTf(t, func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if !result {
			t.Error("EventuallyWithTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := EventuallyWithTf(mock, func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("EventuallyWithTf should return false on failure")
		}
		if !mock.failed {
			t.Error("EventuallyWithT should mark test as failed")
		}
	})
}

func TestExactlyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Exactlyf(t, int32(123), int32(123), "test message")
		if !result {
			t.Error("Exactlyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Exactlyf(mock, int32(123), int64(123), "test message")
		if result {
			t.Error("Exactlyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Exactly should mark test as failed")
		}
	})
}

func TestFailf(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Failf(mock, "failed", "test message")
		if result {
			t.Error("Failf should return false on failure")
		}
		if !mock.failed {
			t.Error("Fail should mark test as failed")
		}
	})
}

func TestFailNowf(t *testing.T) {
	t.Parallel()
	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		result := FailNowf(mock, "failed", "test message")
		if result {
			t.Error("FailNowf should return false on failure")
		}
		if !mock.failed {
			t.Error("FailNow should call FailNow()")
		}
	})
}

func TestFalsef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Falsef(t, 1 == 0, "test message")
		if !result {
			t.Error("Falsef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Falsef(mock, 1 == 1, "test message")
		if result {
			t.Error("Falsef should return false on failure")
		}
		if !mock.failed {
			t.Error("False should mark test as failed")
		}
	})
}

func TestFalseTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FalseTf(t, 1 == 0, "test message")
		if !result {
			t.Error("FalseTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FalseTf(mock, 1 == 1, "test message")
		if result {
			t.Error("FalseTf should return false on failure")
		}
		if !mock.failed {
			t.Error("FalseT should mark test as failed")
		}
	})
}

func TestFileEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileEmptyf(t, filepath.Join(testDataPath(), "empty_file"), "test message")
		if !result {
			t.Error("FileEmptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileEmptyf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		if result {
			t.Error("FileEmptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("FileEmpty should mark test as failed")
		}
	})
}

func TestFileExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileExistsf(t, filepath.Join(testDataPath(), "existing_file"), "test message")
		if !result {
			t.Error("FileExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileExistsf(mock, filepath.Join(testDataPath(), "non_existing_file"), "test message")
		if result {
			t.Error("FileExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("FileExists should mark test as failed")
		}
	})
}

func TestFileNotEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileNotEmptyf(t, filepath.Join(testDataPath(), "existing_file"), "test message")
		if !result {
			t.Error("FileNotEmptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileNotEmptyf(mock, filepath.Join(testDataPath(), "empty_file"), "test message")
		if result {
			t.Error("FileNotEmptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("FileNotEmpty should mark test as failed")
		}
	})
}

func TestFileNotExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := FileNotExistsf(t, filepath.Join(testDataPath(), "non_existing_file"), "test message")
		if !result {
			t.Error("FileNotExistsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := FileNotExistsf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		if result {
			t.Error("FileNotExistsf should return false on failure")
		}
		if !mock.failed {
			t.Error("FileNotExists should mark test as failed")
		}
	})
}

func TestGreaterf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Greaterf(t, 2, 1, "test message")
		if !result {
			t.Error("Greaterf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Greaterf(mock, 1, 2, "test message")
		if result {
			t.Error("Greaterf should return false on failure")
		}
		if !mock.failed {
			t.Error("Greater should mark test as failed")
		}
	})
}

func TestGreaterOrEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := GreaterOrEqualf(t, 2, 1, "test message")
		if !result {
			t.Error("GreaterOrEqualf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := GreaterOrEqualf(mock, 1, 2, "test message")
		if result {
			t.Error("GreaterOrEqualf should return false on failure")
		}
		if !mock.failed {
			t.Error("GreaterOrEqual should mark test as failed")
		}
	})
}

func TestGreaterOrEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := GreaterOrEqualTf(t, 2, 1, "test message")
		if !result {
			t.Error("GreaterOrEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := GreaterOrEqualTf(mock, 1, 2, "test message")
		if result {
			t.Error("GreaterOrEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("GreaterOrEqualT should mark test as failed")
		}
	})
}

func TestGreaterTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := GreaterTf(t, 2, 1, "test message")
		if !result {
			t.Error("GreaterTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := GreaterTf(mock, 1, 2, "test message")
		if result {
			t.Error("GreaterTf should return false on failure")
		}
		if !mock.failed {
			t.Error("GreaterT should mark test as failed")
		}
	})
}

func TestHTTPBodyContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPBodyContainsf(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!", "test message")
		if !result {
			t.Error("HTTPBodyContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPBodyContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!", "test message")
		if result {
			t.Error("HTTPBodyContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPBodyContains should mark test as failed")
		}
	})
}

func TestHTTPBodyNotContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPBodyNotContainsf(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!", "test message")
		if !result {
			t.Error("HTTPBodyNotContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPBodyNotContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!", "test message")
		if result {
			t.Error("HTTPBodyNotContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPBodyNotContains should mark test as failed")
		}
	})
}

func TestHTTPErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPErrorf(t, httpError, "GET", "/", nil, "test message")
		if !result {
			t.Error("HTTPErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPErrorf(mock, httpOK, "GET", "/", nil, "test message")
		if result {
			t.Error("HTTPErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPError should mark test as failed")
		}
	})
}

func TestHTTPRedirectf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPRedirectf(t, httpRedirect, "GET", "/", nil, "test message")
		if !result {
			t.Error("HTTPRedirectf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPRedirectf(mock, httpError, "GET", "/", nil, "test message")
		if result {
			t.Error("HTTPRedirectf should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPRedirect should mark test as failed")
		}
	})
}

func TestHTTPStatusCodef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPStatusCodef(t, httpOK, "GET", "/", nil, http.StatusOK, "test message")
		if !result {
			t.Error("HTTPStatusCodef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPStatusCodef(mock, httpError, "GET", "/", nil, http.StatusOK, "test message")
		if result {
			t.Error("HTTPStatusCodef should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPStatusCode should mark test as failed")
		}
	})
}

func TestHTTPSuccessf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := HTTPSuccessf(t, httpOK, "GET", "/", nil, "test message")
		if !result {
			t.Error("HTTPSuccessf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := HTTPSuccessf(mock, httpError, "GET", "/", nil, "test message")
		if result {
			t.Error("HTTPSuccessf should return false on failure")
		}
		if !mock.failed {
			t.Error("HTTPSuccess should mark test as failed")
		}
	})
}

func TestImplementsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Implementsf(t, ptr(dummyInterface), new(testing.T), "test message")
		if !result {
			t.Error("Implementsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Implementsf(mock, (*error)(nil), new(testing.T), "test message")
		if result {
			t.Error("Implementsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Implements should mark test as failed")
		}
	})
}

func TestInDeltaf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDeltaf(t, 1.0, 1.01, 0.02, "test message")
		if !result {
			t.Error("InDeltaf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaf(mock, 1.0, 1.1, 0.05, "test message")
		if result {
			t.Error("InDeltaf should return false on failure")
		}
		if !mock.failed {
			t.Error("InDelta should mark test as failed")
		}
	})
}

func TestInDeltaMapValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDeltaMapValuesf(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02, "test message")
		if !result {
			t.Error("InDeltaMapValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaMapValuesf(mock, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05, "test message")
		if result {
			t.Error("InDeltaMapValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("InDeltaMapValues should mark test as failed")
		}
	})
}

func TestInDeltaSlicef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDeltaSlicef(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02, "test message")
		if !result {
			t.Error("InDeltaSlicef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaSlicef(mock, []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05, "test message")
		if result {
			t.Error("InDeltaSlicef should return false on failure")
		}
		if !mock.failed {
			t.Error("InDeltaSlice should mark test as failed")
		}
	})
}

func TestInDeltaTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InDeltaTf(t, 1.0, 1.01, 0.02, "test message")
		if !result {
			t.Error("InDeltaTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InDeltaTf(mock, 1.0, 1.1, 0.05, "test message")
		if result {
			t.Error("InDeltaTf should return false on failure")
		}
		if !mock.failed {
			t.Error("InDeltaT should mark test as failed")
		}
	})
}

func TestInEpsilonf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InEpsilonf(t, 100.0, 101.0, 0.02, "test message")
		if !result {
			t.Error("InEpsilonf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilonf(mock, 100.0, 110.0, 0.05, "test message")
		if result {
			t.Error("InEpsilonf should return false on failure")
		}
		if !mock.failed {
			t.Error("InEpsilon should mark test as failed")
		}
	})
}

func TestInEpsilonSlicef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InEpsilonSlicef(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02, "test message")
		if !result {
			t.Error("InEpsilonSlicef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilonSlicef(mock, []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05, "test message")
		if result {
			t.Error("InEpsilonSlicef should return false on failure")
		}
		if !mock.failed {
			t.Error("InEpsilonSlice should mark test as failed")
		}
	})
}

func TestInEpsilonTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := InEpsilonTf(t, 100.0, 101.0, 0.02, "test message")
		if !result {
			t.Error("InEpsilonTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := InEpsilonTf(mock, 100.0, 110.0, 0.05, "test message")
		if result {
			t.Error("InEpsilonTf should return false on failure")
		}
		if !mock.failed {
			t.Error("InEpsilonT should mark test as failed")
		}
	})
}

func TestIsDecreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsDecreasingf(t, []int{3, 2, 1}, "test message")
		if !result {
			t.Error("IsDecreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsDecreasingf(mock, []int{1, 2, 3}, "test message")
		if result {
			t.Error("IsDecreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsDecreasing should mark test as failed")
		}
	})
}

func TestIsDecreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsDecreasingTf(t, []int{3, 2, 1}, "test message")
		if !result {
			t.Error("IsDecreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsDecreasingTf(mock, []int{1, 2, 3}, "test message")
		if result {
			t.Error("IsDecreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsDecreasingT should mark test as failed")
		}
	})
}

func TestIsIncreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsIncreasingf(t, []int{1, 2, 3}, "test message")
		if !result {
			t.Error("IsIncreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsIncreasingf(mock, []int{1, 1, 2}, "test message")
		if result {
			t.Error("IsIncreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsIncreasing should mark test as failed")
		}
	})
}

func TestIsIncreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsIncreasingTf(t, []int{1, 2, 3}, "test message")
		if !result {
			t.Error("IsIncreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsIncreasingTf(mock, []int{1, 1, 2}, "test message")
		if result {
			t.Error("IsIncreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsIncreasingT should mark test as failed")
		}
	})
}

func TestIsNonDecreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonDecreasingf(t, []int{1, 1, 2}, "test message")
		if !result {
			t.Error("IsNonDecreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonDecreasingf(mock, []int{2, 1, 0}, "test message")
		if result {
			t.Error("IsNonDecreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonDecreasing should mark test as failed")
		}
	})
}

func TestIsNonDecreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonDecreasingTf(t, []int{1, 1, 2}, "test message")
		if !result {
			t.Error("IsNonDecreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonDecreasingTf(mock, []int{2, 1, 0}, "test message")
		if result {
			t.Error("IsNonDecreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonDecreasingT should mark test as failed")
		}
	})
}

func TestIsNonIncreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonIncreasingf(t, []int{2, 1, 1}, "test message")
		if !result {
			t.Error("IsNonIncreasingf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonIncreasingf(mock, []int{1, 2, 3}, "test message")
		if result {
			t.Error("IsNonIncreasingf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonIncreasing should mark test as failed")
		}
	})
}

func TestIsNonIncreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNonIncreasingTf(t, []int{2, 1, 1}, "test message")
		if !result {
			t.Error("IsNonIncreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNonIncreasingTf(mock, []int{1, 2, 3}, "test message")
		if result {
			t.Error("IsNonIncreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNonIncreasingT should mark test as failed")
		}
	})
}

func TestIsNotOfTypeTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNotOfTypeTf[myType](t, 123.123, "test message")
		if !result {
			t.Error("IsNotOfTypeTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNotOfTypeTf[myType](mock, myType(123.123), "test message")
		if result {
			t.Error("IsNotOfTypeTf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNotOfTypeT should mark test as failed")
		}
	})
}

func TestIsNotTypef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsNotTypef(t, int32(123), int64(456), "test message")
		if !result {
			t.Error("IsNotTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsNotTypef(mock, 123, 456, "test message")
		if result {
			t.Error("IsNotTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("IsNotType should mark test as failed")
		}
	})
}

func TestIsOfTypeTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsOfTypeTf[myType](t, myType(123.123), "test message")
		if !result {
			t.Error("IsOfTypeTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsOfTypeTf[myType](mock, 123.123, "test message")
		if result {
			t.Error("IsOfTypeTf should return false on failure")
		}
		if !mock.failed {
			t.Error("IsOfTypeT should mark test as failed")
		}
	})
}

func TestIsTypef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := IsTypef(t, 123, 456, "test message")
		if !result {
			t.Error("IsTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := IsTypef(mock, int32(123), int64(456), "test message")
		if result {
			t.Error("IsTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("IsType should mark test as failed")
		}
	})
}

func TestJSONEqf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "test message")
		if !result {
			t.Error("JSONEqf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := JSONEqf(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		if result {
			t.Error("JSONEqf should return false on failure")
		}
		if !mock.failed {
			t.Error("JSONEq should mark test as failed")
		}
	})
}

func TestJSONEqBytesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := JSONEqBytesf(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		if !result {
			t.Error("JSONEqBytesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := JSONEqBytesf(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`), "test message")
		if result {
			t.Error("JSONEqBytesf should return false on failure")
		}
		if !mock.failed {
			t.Error("JSONEqBytes should mark test as failed")
		}
	})
}

func TestJSONEqTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := JSONEqTf(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		if !result {
			t.Error("JSONEqTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := JSONEqTf(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		if result {
			t.Error("JSONEqTf should return false on failure")
		}
		if !mock.failed {
			t.Error("JSONEqT should mark test as failed")
		}
	})
}

func TestKindf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Kindf(t, reflect.String, "hello", "test message")
		if !result {
			t.Error("Kindf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Kindf(mock, reflect.String, 0, "test message")
		if result {
			t.Error("Kindf should return false on failure")
		}
		if !mock.failed {
			t.Error("Kind should mark test as failed")
		}
	})
}

func TestLenf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Lenf(t, []string{"A", "B"}, 2, "test message")
		if !result {
			t.Error("Lenf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Lenf(mock, []string{"A", "B"}, 1, "test message")
		if result {
			t.Error("Lenf should return false on failure")
		}
		if !mock.failed {
			t.Error("Len should mark test as failed")
		}
	})
}

func TestLessf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Lessf(t, 1, 2, "test message")
		if !result {
			t.Error("Lessf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Lessf(mock, 2, 1, "test message")
		if result {
			t.Error("Lessf should return false on failure")
		}
		if !mock.failed {
			t.Error("Less should mark test as failed")
		}
	})
}

func TestLessOrEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := LessOrEqualf(t, 1, 2, "test message")
		if !result {
			t.Error("LessOrEqualf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := LessOrEqualf(mock, 2, 1, "test message")
		if result {
			t.Error("LessOrEqualf should return false on failure")
		}
		if !mock.failed {
			t.Error("LessOrEqual should mark test as failed")
		}
	})
}

func TestLessOrEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := LessOrEqualTf(t, 1, 2, "test message")
		if !result {
			t.Error("LessOrEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := LessOrEqualTf(mock, 2, 1, "test message")
		if result {
			t.Error("LessOrEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("LessOrEqualT should mark test as failed")
		}
	})
}

func TestLessTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := LessTf(t, 1, 2, "test message")
		if !result {
			t.Error("LessTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := LessTf(mock, 2, 1, "test message")
		if result {
			t.Error("LessTf should return false on failure")
		}
		if !mock.failed {
			t.Error("LessT should mark test as failed")
		}
	})
}

func TestMapContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := MapContainsTf(t, map[string]string{"A": "B"}, "A", "test message")
		if !result {
			t.Error("MapContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := MapContainsTf(mock, map[string]string{"A": "B"}, "C", "test message")
		if result {
			t.Error("MapContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("MapContainsT should mark test as failed")
		}
	})
}

func TestMapNotContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := MapNotContainsTf(t, map[string]string{"A": "B"}, "C", "test message")
		if !result {
			t.Error("MapNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := MapNotContainsTf(mock, map[string]string{"A": "B"}, "A", "test message")
		if result {
			t.Error("MapNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("MapNotContainsT should mark test as failed")
		}
	})
}

func TestNegativef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Negativef(t, -1, "test message")
		if !result {
			t.Error("Negativef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Negativef(mock, 1, "test message")
		if result {
			t.Error("Negativef should return false on failure")
		}
		if !mock.failed {
			t.Error("Negative should mark test as failed")
		}
	})
}

func TestNegativeTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NegativeTf(t, -1, "test message")
		if !result {
			t.Error("NegativeTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NegativeTf(mock, 1, "test message")
		if result {
			t.Error("NegativeTf should return false on failure")
		}
		if !mock.failed {
			t.Error("NegativeT should mark test as failed")
		}
	})
}

func TestNeverf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Neverf(t, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if !result {
			t.Error("Neverf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Neverf(mock, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("Neverf should return false on failure")
		}
		if !mock.failed {
			t.Error("Never should mark test as failed")
		}
	})
}

func TestNilf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Nilf(t, nil, "test message")
		if !result {
			t.Error("Nilf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Nilf(mock, "not nil", "test message")
		if result {
			t.Error("Nilf should return false on failure")
		}
		if !mock.failed {
			t.Error("Nil should mark test as failed")
		}
	})
}

func TestNoErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NoErrorf(t, nil, "test message")
		if !result {
			t.Error("NoErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NoErrorf(mock, ErrTest, "test message")
		if result {
			t.Error("NoErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("NoError should mark test as failed")
		}
	})
}

func TestNotContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotContainsf(t, []string{"A", "B"}, "C", "test message")
		if !result {
			t.Error("NotContainsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotContainsf(mock, []string{"A", "B"}, "B", "test message")
		if result {
			t.Error("NotContainsf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotContains should mark test as failed")
		}
	})
}

func TestNotElementsMatchf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotElementsMatchf(t, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if !result {
			t.Error("NotElementsMatchf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotElementsMatchf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if result {
			t.Error("NotElementsMatchf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotElementsMatch should mark test as failed")
		}
	})
}

func TestNotElementsMatchTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotElementsMatchTf(t, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if !result {
			t.Error("NotElementsMatchTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotElementsMatchTf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if result {
			t.Error("NotElementsMatchTf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotElementsMatchT should mark test as failed")
		}
	})
}

func TestNotEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEmptyf(t, "not empty", "test message")
		if !result {
			t.Error("NotEmptyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEmptyf(mock, "", "test message")
		if result {
			t.Error("NotEmptyf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEmpty should mark test as failed")
		}
	})
}

func TestNotEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEqualf(t, 123, 456, "test message")
		if !result {
			t.Error("NotEqualf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEqualf(mock, 123, 123, "test message")
		if result {
			t.Error("NotEqualf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEqual should mark test as failed")
		}
	})
}

func TestNotEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEqualTf(t, 123, 456, "test message")
		if !result {
			t.Error("NotEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEqualTf(mock, 123, 123, "test message")
		if result {
			t.Error("NotEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEqualT should mark test as failed")
		}
	})
}

func TestNotEqualValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotEqualValuesf(t, uint32(123), int32(456), "test message")
		if !result {
			t.Error("NotEqualValuesf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotEqualValuesf(mock, uint32(123), int32(123), "test message")
		if result {
			t.Error("NotEqualValuesf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotEqualValues should mark test as failed")
		}
	})
}

func TestNotErrorAsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotErrorAsf(t, ErrTest, new(*dummyError), "test message")
		if !result {
			t.Error("NotErrorAsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorAsf(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if result {
			t.Error("NotErrorAsf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotErrorAs should mark test as failed")
		}
	})
}

func TestNotErrorIsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotErrorIsf(t, ErrTest, io.EOF, "test message")
		if !result {
			t.Error("NotErrorIsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorIsf(mock, fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		if result {
			t.Error("NotErrorIsf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotErrorIs should mark test as failed")
		}
	})
}

func TestNotImplementsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotImplementsf(t, (*error)(nil), new(testing.T), "test message")
		if !result {
			t.Error("NotImplementsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotImplementsf(mock, ptr(dummyInterface), new(testing.T), "test message")
		if result {
			t.Error("NotImplementsf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotImplements should mark test as failed")
		}
	})
}

func TestNotKindf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotKindf(t, reflect.String, 0, "test message")
		if !result {
			t.Error("NotKindf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotKindf(mock, reflect.String, "hello", "test message")
		if result {
			t.Error("NotKindf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotKind should mark test as failed")
		}
	})
}

func TestNotNilf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotNilf(t, "not nil", "test message")
		if !result {
			t.Error("NotNilf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotNilf(mock, nil, "test message")
		if result {
			t.Error("NotNilf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotNil should mark test as failed")
		}
	})
}

func TestNotPanicsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotPanicsf(t, func() {}, "test message")
		if !result {
			t.Error("NotPanicsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotPanicsf(mock, func() { panic("panicking") }, "test message")
		if result {
			t.Error("NotPanicsf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotPanics should mark test as failed")
		}
	})
}

func TestNotRegexpf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotRegexpf(t, "^start", "not starting", "test message")
		if !result {
			t.Error("NotRegexpf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotRegexpf(mock, "^start", "starting", "test message")
		if result {
			t.Error("NotRegexpf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotRegexp should mark test as failed")
		}
	})
}

func TestNotRegexpTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotRegexpTf(t, "^start", "not starting", "test message")
		if !result {
			t.Error("NotRegexpTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotRegexpTf(mock, "^start", "starting", "test message")
		if result {
			t.Error("NotRegexpTf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotRegexpT should mark test as failed")
		}
	})
}

func TestNotSamef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSamef(t, &staticVar, ptr("static string"), "test message")
		if !result {
			t.Error("NotSamef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSamef(mock, &staticVar, staticVarPtr, "test message")
		if result {
			t.Error("NotSamef should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSame should mark test as failed")
		}
	})
}

func TestNotSameTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSameTf(t, &staticVar, ptr("static string"), "test message")
		if !result {
			t.Error("NotSameTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSameTf(mock, &staticVar, staticVarPtr, "test message")
		if result {
			t.Error("NotSameTf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSameT should mark test as failed")
		}
	})
}

func TestNotSortedTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSortedTf(t, []int{3, 1, 3}, "test message")
		if !result {
			t.Error("NotSortedTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSortedTf(mock, []int{1, 4, 8}, "test message")
		if result {
			t.Error("NotSortedTf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSortedT should mark test as failed")
		}
	})
}

func TestNotSubsetf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotSubsetf(t, []int{1, 2, 3}, []int{4, 5}, "test message")
		if !result {
			t.Error("NotSubsetf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotSubsetf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		if result {
			t.Error("NotSubsetf should return false on failure")
		}
		if !mock.failed {
			t.Error("NotSubset should mark test as failed")
		}
	})
}

func TestNotZerof(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := NotZerof(t, 1, "test message")
		if !result {
			t.Error("NotZerof should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotZerof(mock, 0, "test message")
		if result {
			t.Error("NotZerof should return false on failure")
		}
		if !mock.failed {
			t.Error("NotZero should mark test as failed")
		}
	})
}

func TestPanicsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Panicsf(t, func() { panic("panicking") }, "test message")
		if !result {
			t.Error("Panicsf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Panicsf(mock, func() {}, "test message")
		if result {
			t.Error("Panicsf should return false on failure")
		}
		if !mock.failed {
			t.Error("Panics should mark test as failed")
		}
	})
}

func TestPanicsWithErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := PanicsWithErrorf(t, ErrTest.Error(), func() { panic(ErrTest) }, "test message")
		if !result {
			t.Error("PanicsWithErrorf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := PanicsWithErrorf(mock, ErrTest.Error(), func() {}, "test message")
		if result {
			t.Error("PanicsWithErrorf should return false on failure")
		}
		if !mock.failed {
			t.Error("PanicsWithError should mark test as failed")
		}
	})
}

func TestPanicsWithValuef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := PanicsWithValuef(t, "panicking", func() { panic("panicking") }, "test message")
		if !result {
			t.Error("PanicsWithValuef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := PanicsWithValuef(mock, "panicking", func() {}, "test message")
		if result {
			t.Error("PanicsWithValuef should return false on failure")
		}
		if !mock.failed {
			t.Error("PanicsWithValue should mark test as failed")
		}
	})
}

func TestPositivef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Positivef(t, 1, "test message")
		if !result {
			t.Error("Positivef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Positivef(mock, -1, "test message")
		if result {
			t.Error("Positivef should return false on failure")
		}
		if !mock.failed {
			t.Error("Positive should mark test as failed")
		}
	})
}

func TestPositiveTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := PositiveTf(t, 1, "test message")
		if !result {
			t.Error("PositiveTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := PositiveTf(mock, -1, "test message")
		if result {
			t.Error("PositiveTf should return false on failure")
		}
		if !mock.failed {
			t.Error("PositiveT should mark test as failed")
		}
	})
}

func TestRegexpf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Regexpf(t, "^start", "starting", "test message")
		if !result {
			t.Error("Regexpf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Regexpf(mock, "^start", "not starting", "test message")
		if result {
			t.Error("Regexpf should return false on failure")
		}
		if !mock.failed {
			t.Error("Regexp should mark test as failed")
		}
	})
}

func TestRegexpTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := RegexpTf(t, "^start", "starting", "test message")
		if !result {
			t.Error("RegexpTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := RegexpTf(mock, "^start", "not starting", "test message")
		if result {
			t.Error("RegexpTf should return false on failure")
		}
		if !mock.failed {
			t.Error("RegexpT should mark test as failed")
		}
	})
}

func TestSamef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Samef(t, &staticVar, staticVarPtr, "test message")
		if !result {
			t.Error("Samef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Samef(mock, &staticVar, ptr("static string"), "test message")
		if result {
			t.Error("Samef should return false on failure")
		}
		if !mock.failed {
			t.Error("Same should mark test as failed")
		}
	})
}

func TestSameTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SameTf(t, &staticVar, staticVarPtr, "test message")
		if !result {
			t.Error("SameTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SameTf(mock, &staticVar, ptr("static string"), "test message")
		if result {
			t.Error("SameTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SameT should mark test as failed")
		}
	})
}

func TestSeqContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SeqContainsTf(t, slices.Values([]string{"A", "B"}), "A", "test message")
		if !result {
			t.Error("SeqContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SeqContainsTf(mock, slices.Values([]string{"A", "B"}), "C", "test message")
		if result {
			t.Error("SeqContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SeqContainsT should mark test as failed")
		}
	})
}

func TestSeqNotContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SeqNotContainsTf(t, slices.Values([]string{"A", "B"}), "C", "test message")
		if !result {
			t.Error("SeqNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SeqNotContainsTf(mock, slices.Values([]string{"A", "B"}), "A", "test message")
		if result {
			t.Error("SeqNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SeqNotContainsT should mark test as failed")
		}
	})
}

func TestSliceContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceContainsTf(t, []string{"A", "B"}, "A", "test message")
		if !result {
			t.Error("SliceContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceContainsTf(mock, []string{"A", "B"}, "C", "test message")
		if result {
			t.Error("SliceContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceContainsT should mark test as failed")
		}
	})
}

func TestSliceNotContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceNotContainsTf(t, []string{"A", "B"}, "C", "test message")
		if !result {
			t.Error("SliceNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceNotContainsTf(mock, []string{"A", "B"}, "A", "test message")
		if result {
			t.Error("SliceNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceNotContainsT should mark test as failed")
		}
	})
}

func TestSliceNotSubsetTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceNotSubsetTf(t, []int{1, 2, 3}, []int{4, 5}, "test message")
		if !result {
			t.Error("SliceNotSubsetTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceNotSubsetTf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		if result {
			t.Error("SliceNotSubsetTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceNotSubsetT should mark test as failed")
		}
	})
}

func TestSliceSubsetTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SliceSubsetTf(t, []int{1, 2, 3}, []int{1, 2}, "test message")
		if !result {
			t.Error("SliceSubsetTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SliceSubsetTf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		if result {
			t.Error("SliceSubsetTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SliceSubsetT should mark test as failed")
		}
	})
}

func TestSortedTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := SortedTf(t, []int{1, 1, 3}, "test message")
		if !result {
			t.Error("SortedTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := SortedTf(mock, []int{1, 4, 2}, "test message")
		if result {
			t.Error("SortedTf should return false on failure")
		}
		if !mock.failed {
			t.Error("SortedT should mark test as failed")
		}
	})
}

func TestStringContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := StringContainsTf(t, "AB", "A", "test message")
		if !result {
			t.Error("StringContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := StringContainsTf(mock, "AB", "C", "test message")
		if result {
			t.Error("StringContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("StringContainsT should mark test as failed")
		}
	})
}

func TestStringNotContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := StringNotContainsTf(t, "AB", "C", "test message")
		if !result {
			t.Error("StringNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := StringNotContainsTf(mock, "AB", "A", "test message")
		if result {
			t.Error("StringNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("StringNotContainsT should mark test as failed")
		}
	})
}

func TestSubsetf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Subsetf(t, []int{1, 2, 3}, []int{1, 2}, "test message")
		if !result {
			t.Error("Subsetf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Subsetf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		if result {
			t.Error("Subsetf should return false on failure")
		}
		if !mock.failed {
			t.Error("Subset should mark test as failed")
		}
	})
}

func TestTruef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Truef(t, 1 == 1, "test message")
		if !result {
			t.Error("Truef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Truef(mock, 1 == 0, "test message")
		if result {
			t.Error("Truef should return false on failure")
		}
		if !mock.failed {
			t.Error("True should mark test as failed")
		}
	})
}

func TestTrueTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := TrueTf(t, 1 == 1, "test message")
		if !result {
			t.Error("TrueTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := TrueTf(mock, 1 == 0, "test message")
		if result {
			t.Error("TrueTf should return false on failure")
		}
		if !mock.failed {
			t.Error("TrueT should mark test as failed")
		}
	})
}

func TestWithinDurationf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := WithinDurationf(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second, "test message")
		if !result {
			t.Error("WithinDurationf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := WithinDurationf(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second, "test message")
		if result {
			t.Error("WithinDurationf should return false on failure")
		}
		if !mock.failed {
			t.Error("WithinDuration should mark test as failed")
		}
	})
}

func TestWithinRangef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := WithinRangef(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		if !result {
			t.Error("WithinRangef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := WithinRangef(mock, time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		if result {
			t.Error("WithinRangef should return false on failure")
		}
		if !mock.failed {
			t.Error("WithinRange should mark test as failed")
		}
	})
}

func TestYAMLEqf(t *testing.T) {
	t.Parallel()
	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		Panicsf(t, func() {
			YAMLEqf(t, "key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
	})
}

func TestYAMLEqBytesf(t *testing.T) {
	t.Parallel()
	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		Panicsf(t, func() {
			YAMLEqBytesf(t, []byte("key: value"), []byte("key: value"), "test message")
		}, "should panic without the yaml feature enabled.")
	})
}

func TestYAMLEqTf(t *testing.T) {
	t.Parallel()
	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		Panicsf(t, func() {
			YAMLEqTf(t, "key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
	})
}

func TestZerof(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		result := Zerof(t, 0, "test message")
		if !result {
			t.Error("Zerof should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := Zerof(mock, 1, "test message")
		if result {
			t.Error("Zerof should return false on failure")
		}
		if !mock.failed {
			t.Error("Zero should mark test as failed")
		}
	})
}
