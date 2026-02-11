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
	"slices"
	"testing"
	"time"
)

func TestConditionf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Conditionf(mock, func() bool { return true }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Conditionf(mock, func() bool { return false }, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Conditionf should call FailNow()")
		}
	})
}

func TestContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Containsf(mock, []string{"A", "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Containsf(mock, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Containsf should call FailNow()")
		}
	})
}

func TestDirExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirExistsf(mock, filepath.Join(testDataPath(), "existing_dir"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirExistsf(mock, filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("DirExistsf should call FailNow()")
		}
	})
}

func TestDirNotExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirNotExistsf(mock, filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirNotExistsf(mock, filepath.Join(testDataPath(), "existing_dir"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("DirNotExistsf should call FailNow()")
		}
	})
}

func TestElementsMatchf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ElementsMatchf should call FailNow()")
		}
	})
}

func TestElementsMatchTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchTf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchTf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ElementsMatchTf should call FailNow()")
		}
	})
}

func TestEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Emptyf(mock, "", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Emptyf(mock, "not empty", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Emptyf should call FailNow()")
		}
	})
}

func TestEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Equalf(mock, 123, 123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Equalf(mock, 123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Equalf should call FailNow()")
		}
	})
}

func TestEqualErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualErrorf(mock, ErrTest, "assert.ErrTest general error for testing", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualErrorf(mock, ErrTest, "wrong error message", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualErrorf should call FailNow()")
		}
	})
}

func TestEqualExportedValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualExportedValuesf(mock, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualExportedValuesf(mock, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualExportedValuesf should call FailNow()")
		}
	})
}

func TestEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualTf(mock, 123, 123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualTf(mock, 123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualTf should call FailNow()")
		}
	})
}

func TestEqualValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualValuesf(mock, uint32(123), int32(123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualValuesf(mock, uint32(123), int32(456), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualValuesf should call FailNow()")
		}
	})
}

func TestErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Errorf(mock, ErrTest, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Errorf(mock, nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Errorf should call FailNow()")
		}
	})
}

func TestErrorAsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAsf(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAsf(mock, ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorAsf should call FailNow()")
		}
	})
}

func TestErrorContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorContainsf(mock, ErrTest, "general error", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorContainsf(mock, ErrTest, "not in message", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorContainsf should call FailNow()")
		}
	})
}

func TestErrorIsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorIsf(mock, fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorIsf(mock, ErrTest, io.EOF, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorIsf should call FailNow()")
		}
	})
}

func TestEventuallyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Eventuallyf(mock, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Eventuallyf(mock, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Eventuallyf should call FailNow()")
		}
	})
}

func TestEventuallyWithf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EventuallyWithf(mock, func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EventuallyWithf(mock, func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EventuallyWithf should call FailNow()")
		}
	})
}

func TestExactlyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Exactlyf(mock, int32(123), int32(123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Exactlyf(mock, int32(123), int64(123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Exactlyf should call FailNow()")
		}
	})
}

func TestFailf(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Failf(mock, "failed", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Failf should call FailNow()")
		}
	})
}

func TestFailNowf(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FailNowf(mock, "failed", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FailNowf should call FailNow()")
		}
	})
}

func TestFalsef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Falsef(mock, 1 == 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Falsef(mock, 1 == 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Falsef should call FailNow()")
		}
	})
}

func TestFalseTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FalseTf(mock, 1 == 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FalseTf(mock, 1 == 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FalseTf should call FailNow()")
		}
	})
}

func TestFileEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileEmptyf(mock, filepath.Join(testDataPath(), "empty_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileEmptyf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileEmptyf should call FailNow()")
		}
	})
}

func TestFileExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileExistsf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileExistsf(mock, filepath.Join(testDataPath(), "non_existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileExistsf should call FailNow()")
		}
	})
}

func TestFileNotEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotEmptyf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotEmptyf(mock, filepath.Join(testDataPath(), "empty_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileNotEmptyf should call FailNow()")
		}
	})
}

func TestFileNotExistsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotExistsf(mock, filepath.Join(testDataPath(), "non_existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotExistsf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileNotExistsf should call FailNow()")
		}
	})
}

func TestGreaterf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Greaterf(mock, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Greaterf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Greaterf should call FailNow()")
		}
	})
}

func TestGreaterOrEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualf(mock, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterOrEqualf should call FailNow()")
		}
	})
}

func TestGreaterOrEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualTf(mock, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualTf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterOrEqualTf should call FailNow()")
		}
	})
}

func TestGreaterTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterTf(mock, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterTf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterTf should call FailNow()")
		}
	})
}

func TestHTTPBodyContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPBodyContainsf should call FailNow()")
		}
	})
}

func TestHTTPBodyNotContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyNotContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyNotContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPBodyNotContainsf should call FailNow()")
		}
	})
}

func TestHTTPErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPErrorf(mock, httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPErrorf(mock, httpOK, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPErrorf should call FailNow()")
		}
	})
}

func TestHTTPRedirectf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPRedirectf(mock, httpRedirect, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPRedirectf(mock, httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPRedirectf should call FailNow()")
		}
	})
}

func TestHTTPStatusCodef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPStatusCodef(mock, httpOK, "GET", "/", nil, http.StatusOK, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPStatusCodef(mock, httpError, "GET", "/", nil, http.StatusOK, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPStatusCodef should call FailNow()")
		}
	})
}

func TestHTTPSuccessf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPSuccessf(mock, httpOK, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPSuccessf(mock, httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPSuccessf should call FailNow()")
		}
	})
}

func TestImplementsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Implementsf(mock, ptr(dummyInterface), new(testing.T), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Implementsf(mock, (*error)(nil), new(testing.T), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Implementsf should call FailNow()")
		}
	})
}

func TestInDeltaf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaf(mock, 1.0, 1.01, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaf(mock, 1.0, 1.1, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaf should call FailNow()")
		}
	})
}

func TestInDeltaMapValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaMapValuesf(mock, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaMapValuesf(mock, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaMapValuesf should call FailNow()")
		}
	})
}

func TestInDeltaSlicef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaSlicef(mock, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaSlicef(mock, []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaSlicef should call FailNow()")
		}
	})
}

func TestInDeltaTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaTf(mock, 1.0, 1.01, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaTf(mock, 1.0, 1.1, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaTf should call FailNow()")
		}
	})
}

func TestInEpsilonf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonf(mock, 100.0, 101.0, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonf(mock, 100.0, 110.0, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilonf should call FailNow()")
		}
	})
}

func TestInEpsilonSlicef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonSlicef(mock, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonSlicef(mock, []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilonSlicef should call FailNow()")
		}
	})
}

func TestInEpsilonTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonTf(mock, 100.0, 101.0, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonTf(mock, 100.0, 110.0, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilonTf should call FailNow()")
		}
	})
}

func TestIsDecreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingf(mock, []int{3, 2, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsDecreasingf should call FailNow()")
		}
	})
}

func TestIsDecreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingTf(mock, []int{3, 2, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingTf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsDecreasingTf should call FailNow()")
		}
	})
}

func TestIsIncreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingf(mock, []int{1, 1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsIncreasingf should call FailNow()")
		}
	})
}

func TestIsIncreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingTf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingTf(mock, []int{1, 1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsIncreasingTf should call FailNow()")
		}
	})
}

func TestIsNonDecreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingf(mock, []int{1, 1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingf(mock, []int{2, 1, 0}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonDecreasingf should call FailNow()")
		}
	})
}

func TestIsNonDecreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingTf(mock, []int{1, 1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingTf(mock, []int{2, 1, 0}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonDecreasingTf should call FailNow()")
		}
	})
}

func TestIsNonIncreasingf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingf(mock, []int{2, 1, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonIncreasingf should call FailNow()")
		}
	})
}

func TestIsNonIncreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingTf(mock, []int{2, 1, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingTf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonIncreasingTf should call FailNow()")
		}
	})
}

func TestIsNotOfTypeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotOfTypeTf[myType](mock, 123.123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotOfTypeTf[myType](mock, myType(123.123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNotOfTypeTf should call FailNow()")
		}
	})
}

func TestIsNotTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotTypef(mock, int32(123), int64(456), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotTypef(mock, 123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNotTypef should call FailNow()")
		}
	})
}

func TestIsOfTypeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsOfTypeTf[myType](mock, myType(123.123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsOfTypeTf[myType](mock, 123.123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsOfTypeTf should call FailNow()")
		}
	})
}

func TestIsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsTypef(mock, 123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsTypef(mock, int32(123), int64(456), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsTypef should call FailNow()")
		}
	})
}

func TestJSONEqf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqf(mock, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqf(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEqf should call FailNow()")
		}
	})
}

func TestJSONEqBytesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqBytesf(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqBytesf(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEqBytesf should call FailNow()")
		}
	})
}

func TestJSONEqTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqTf(mock, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqTf(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEqTf should call FailNow()")
		}
	})
}

func TestJSONMarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONMarshalAsTf(mock, []byte(`{"A": "a"}`), dummyStruct{A: "a"}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONMarshalAsTf(mock, `[{"foo": "bar"}, {"hello": "world"}]`, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONMarshalAsTf should call FailNow()")
		}
	})
}

func TestJSONUnmarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONUnmarshalAsTf(mock, dummyStruct{A: "a"}, []byte(`{"A": "a"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONUnmarshalAsTf(mock, 1, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONUnmarshalAsTf should call FailNow()")
		}
	})
}

func TestKindf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Kindf(mock, reflect.String, "hello", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Kindf(mock, reflect.String, 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Kindf should call FailNow()")
		}
	})
}

func TestLenf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Lenf(mock, []string{"A", "B"}, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Lenf(mock, []string{"A", "B"}, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Lenf should call FailNow()")
		}
	})
}

func TestLessf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Lessf(mock, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Lessf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Lessf should call FailNow()")
		}
	})
}

func TestLessOrEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualf(mock, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessOrEqualf should call FailNow()")
		}
	})
}

func TestLessOrEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualTf(mock, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualTf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessOrEqualTf should call FailNow()")
		}
	})
}

func TestLessTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessTf(mock, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessTf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessTf should call FailNow()")
		}
	})
}

func TestMapContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapContainsTf(mock, map[string]string{"A": "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapContainsTf(mock, map[string]string{"A": "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("MapContainsTf should call FailNow()")
		}
	})
}

func TestMapNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapNotContainsTf(mock, map[string]string{"A": "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapNotContainsTf(mock, map[string]string{"A": "B"}, "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("MapNotContainsTf should call FailNow()")
		}
	})
}

func TestNegativef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Negativef(mock, -1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Negativef(mock, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Negativef should call FailNow()")
		}
	})
}

func TestNegativeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NegativeTf(mock, -1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NegativeTf(mock, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NegativeTf should call FailNow()")
		}
	})
}

func TestNeverf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Neverf(mock, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Neverf(mock, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Neverf should call FailNow()")
		}
	})
}

func TestNilf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Nilf(mock, nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Nilf(mock, "not nil", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Nilf should call FailNow()")
		}
	})
}

func TestNoErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoErrorf(mock, nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoErrorf(mock, ErrTest, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NoErrorf should call FailNow()")
		}
	})
}

func TestNoFileDescriptorLeakf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoFileDescriptorLeakf(mock, func() {}, "test message")
		// require functions don't return a value
	})
}

func TestNoGoRoutineLeakf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoGoRoutineLeakf(mock, func() {}, "test message")
		// require functions don't return a value
	})
}

func TestNotContainsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotContainsf(mock, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotContainsf(mock, []string{"A", "B"}, "B", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotContainsf should call FailNow()")
		}
	})
}

func TestNotElementsMatchf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotElementsMatchf should call FailNow()")
		}
	})
}

func TestNotElementsMatchTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchTf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchTf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotElementsMatchTf should call FailNow()")
		}
	})
}

func TestNotEmptyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEmptyf(mock, "not empty", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEmptyf(mock, "", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEmptyf should call FailNow()")
		}
	})
}

func TestNotEqualf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualf(mock, 123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualf(mock, 123, 123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqualf should call FailNow()")
		}
	})
}

func TestNotEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualTf(mock, 123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualTf(mock, 123, 123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqualTf should call FailNow()")
		}
	})
}

func TestNotEqualValuesf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualValuesf(mock, uint32(123), int32(456), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualValuesf(mock, uint32(123), int32(123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqualValuesf should call FailNow()")
		}
	})
}

func TestNotErrorAsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAsf(mock, ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAsf(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorAsf should call FailNow()")
		}
	})
}

func TestNotErrorIsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorIsf(mock, ErrTest, io.EOF, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorIsf(mock, fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorIsf should call FailNow()")
		}
	})
}

func TestNotImplementsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotImplementsf(mock, (*error)(nil), new(testing.T), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotImplementsf(mock, ptr(dummyInterface), new(testing.T), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotImplementsf should call FailNow()")
		}
	})
}

func TestNotKindf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotKindf(mock, reflect.String, 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotKindf(mock, reflect.String, "hello", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotKindf should call FailNow()")
		}
	})
}

func TestNotNilf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotNilf(mock, "not nil", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotNilf(mock, nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotNilf should call FailNow()")
		}
	})
}

func TestNotPanicsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotPanicsf(mock, func() {}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotPanicsf(mock, func() { panic("panicking") }, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotPanicsf should call FailNow()")
		}
	})
}

func TestNotRegexpf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpf(mock, "^start", "not starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpf(mock, "^start", "starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotRegexpf should call FailNow()")
		}
	})
}

func TestNotRegexpTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpTf(mock, "^start", "not starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpTf(mock, "^start", "starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotRegexpTf should call FailNow()")
		}
	})
}

func TestNotSamef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSamef(mock, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSamef(mock, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSamef should call FailNow()")
		}
	})
}

func TestNotSameTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSameTf(mock, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSameTf(mock, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSameTf should call FailNow()")
		}
	})
}

func TestNotSortedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSortedTf(mock, []int{3, 1, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSortedTf(mock, []int{1, 4, 8}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSortedTf should call FailNow()")
		}
	})
}

func TestNotSubsetf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSubsetf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSubsetf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSubsetf should call FailNow()")
		}
	})
}

func TestNotZerof(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotZerof(mock, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotZerof(mock, 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotZerof should call FailNow()")
		}
	})
}

func TestPanicsf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panicsf(mock, func() { panic("panicking") }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panicsf(mock, func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Panicsf should call FailNow()")
		}
	})
}

func TestPanicsWithErrorf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithErrorf(mock, ErrTest.Error(), func() { panic(ErrTest) }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithErrorf(mock, ErrTest.Error(), func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("PanicsWithErrorf should call FailNow()")
		}
	})
}

func TestPanicsWithValuef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithValuef(mock, "panicking", func() { panic("panicking") }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithValuef(mock, "panicking", func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("PanicsWithValuef should call FailNow()")
		}
	})
}

func TestPositivef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Positivef(mock, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Positivef(mock, -1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Positivef should call FailNow()")
		}
	})
}

func TestPositiveTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PositiveTf(mock, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PositiveTf(mock, -1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("PositiveTf should call FailNow()")
		}
	})
}

func TestRegexpf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Regexpf(mock, "^start", "starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Regexpf(mock, "^start", "not starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Regexpf should call FailNow()")
		}
	})
}

func TestRegexpTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		RegexpTf(mock, "^start", "starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		RegexpTf(mock, "^start", "not starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("RegexpTf should call FailNow()")
		}
	})
}

func TestSamef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Samef(mock, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Samef(mock, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Samef should call FailNow()")
		}
	})
}

func TestSameTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SameTf(mock, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SameTf(mock, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SameTf should call FailNow()")
		}
	})
}

func TestSeqContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqContainsTf(mock, slices.Values([]string{"A", "B"}), "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqContainsTf(mock, slices.Values([]string{"A", "B"}), "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SeqContainsTf should call FailNow()")
		}
	})
}

func TestSeqNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqNotContainsTf(mock, slices.Values([]string{"A", "B"}), "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqNotContainsTf(mock, slices.Values([]string{"A", "B"}), "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SeqNotContainsTf should call FailNow()")
		}
	})
}

func TestSliceContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceContainsTf(mock, []string{"A", "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceContainsTf(mock, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceContainsTf should call FailNow()")
		}
	})
}

func TestSliceNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotContainsTf(mock, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotContainsTf(mock, []string{"A", "B"}, "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceNotContainsTf should call FailNow()")
		}
	})
}

func TestSliceNotSubsetTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotSubsetTf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotSubsetTf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceNotSubsetTf should call FailNow()")
		}
	})
}

func TestSliceSubsetTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceSubsetTf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceSubsetTf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceSubsetTf should call FailNow()")
		}
	})
}

func TestSortedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SortedTf(mock, []int{1, 1, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SortedTf(mock, []int{1, 4, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SortedTf should call FailNow()")
		}
	})
}

func TestStringContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringContainsTf(mock, "AB", "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringContainsTf(mock, "AB", "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("StringContainsTf should call FailNow()")
		}
	})
}

func TestStringNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringNotContainsTf(mock, "AB", "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringNotContainsTf(mock, "AB", "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("StringNotContainsTf should call FailNow()")
		}
	})
}

func TestSubsetf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Subsetf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Subsetf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Subsetf should call FailNow()")
		}
	})
}

func TestTruef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Truef(mock, 1 == 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Truef(mock, 1 == 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Truef should call FailNow()")
		}
	})
}

func TestTrueTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		TrueTf(mock, 1 == 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		TrueTf(mock, 1 == 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("TrueTf should call FailNow()")
		}
	})
}

func TestWithinDurationf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinDurationf(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinDurationf(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("WithinDurationf should call FailNow()")
		}
	})
}

func TestWithinRangef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinRangef(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinRangef(mock, time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("WithinRangef should call FailNow()")
		}
	})
}

func TestYAMLEqf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLEqf(mock, "key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLEqf should panic as expected")
		}
	})
}

func TestYAMLEqBytesf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLEqBytesf(mock, []byte("key: value"), []byte("key: value"), "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLEqBytesf should panic as expected")
		}
	})
}

func TestYAMLEqTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLEqTf(mock, "key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLEqTf should panic as expected")
		}
	})
}

func TestYAMLMarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLMarshalAsTf(mock, "key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLMarshalAsTf should panic as expected")
		}
	})
}

func TestYAMLUnmarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLUnmarshalAsTf(mock, "key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLUnmarshalAsTf should panic as expected")
		}
	})
}

func TestZerof(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Zerof(mock, 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Zerof(mock, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Zerof should call FailNow()")
		}
	})
}
