// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.
// Generated on 2026-01-19 (version fbbb078) using codegen version v2.1.9-0.20260119215714-fbbb0787fd81+dirty [sha: fbbb0787fd8131d63f280f85b14e47f7c0dc8ee0]

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

func TestConditionf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Conditionf(t, func() bool { return true }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Conditionf(mock, func() bool { return false }, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Condition should call FailNow()")
		}
	})
}

func TestContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Containsf(t, []string{"A", "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Containsf(mock, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Contains should call FailNow()")
		}
	})
}

func TestDirExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		DirExistsf(t, filepath.Join(testDataPath(), "existing_dir"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirExistsf(mock, filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("DirExists should call FailNow()")
		}
	})
}

func TestDirNotExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		DirNotExistsf(t, filepath.Join(testDataPath(), "non_existing_dir"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirNotExistsf(mock, filepath.Join(testDataPath(), "existing_dir"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("DirNotExists should call FailNow()")
		}
	})
}

func TestElementsMatchf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		ElementsMatchf(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ElementsMatch should call FailNow()")
		}
	})
}

func TestElementsMatchTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		ElementsMatchTf(t, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchTf(mock, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ElementsMatchT should call FailNow()")
		}
	})
}

func TestEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Emptyf(t, "", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Emptyf(mock, "not empty", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Empty should call FailNow()")
		}
	})
}

func TestEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Equalf(t, 123, 123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Equalf(mock, 123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Equal should call FailNow()")
		}
	})
}

func TestEqualErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		EqualErrorf(t, ErrTest, "assert.ErrTest general error for testing", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualErrorf(mock, ErrTest, "wrong error message", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualError should call FailNow()")
		}
	})
}

func TestEqualExportedValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		EqualExportedValuesf(t, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualExportedValuesf(mock, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualExportedValues should call FailNow()")
		}
	})
}

func TestEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		EqualTf(t, 123, 123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualTf(mock, 123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualT should call FailNow()")
		}
	})
}

func TestEqualValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		EqualValuesf(t, uint32(123), int32(123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualValuesf(mock, uint32(123), int32(456), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualValues should call FailNow()")
		}
	})
}

func TestErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Errorf(t, ErrTest, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Errorf(mock, nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Error should call FailNow()")
		}
	})
}

func TestErrorAsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		ErrorAsf(t, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAsf(mock, ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorAs should call FailNow()")
		}
	})
}

func TestErrorContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		ErrorContainsf(t, ErrTest, "general error", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorContainsf(mock, ErrTest, "not in message", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorContains should call FailNow()")
		}
	})
}

func TestErrorIsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		ErrorIsf(t, fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorIsf(mock, ErrTest, io.EOF, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorIs should call FailNow()")
		}
	})
}

func TestEventuallyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Eventuallyf(t, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Eventuallyf(mock, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Eventually should call FailNow()")
		}
	})
}

func TestEventuallyWithTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		EventuallyWithTf(t, func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EventuallyWithTf(mock, func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EventuallyWithT should call FailNow()")
		}
	})
}

func TestExactlyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Exactlyf(t, int32(123), int32(123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Exactlyf(mock, int32(123), int64(123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Exactly should call FailNow()")
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
			t.Error("Fail should call FailNow()")
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
			t.Error("FailNow should call FailNow()")
		}
	})
}

func TestFalsef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Falsef(t, 1 == 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Falsef(mock, 1 == 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("False should call FailNow()")
		}
	})
}

func TestFalseTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		FalseTf(t, 1 == 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FalseTf(mock, 1 == 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FalseT should call FailNow()")
		}
	})
}

func TestFileEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		FileEmptyf(t, filepath.Join(testDataPath(), "empty_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileEmptyf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileEmpty should call FailNow()")
		}
	})
}

func TestFileExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		FileExistsf(t, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileExistsf(mock, filepath.Join(testDataPath(), "non_existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileExists should call FailNow()")
		}
	})
}

func TestFileNotEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		FileNotEmptyf(t, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotEmptyf(mock, filepath.Join(testDataPath(), "empty_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileNotEmpty should call FailNow()")
		}
	})
}

func TestFileNotExistsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		FileNotExistsf(t, filepath.Join(testDataPath(), "non_existing_file"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotExistsf(mock, filepath.Join(testDataPath(), "existing_file"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileNotExists should call FailNow()")
		}
	})
}

func TestGreaterf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Greaterf(t, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Greaterf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Greater should call FailNow()")
		}
	})
}

func TestGreaterOrEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		GreaterOrEqualf(t, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterOrEqual should call FailNow()")
		}
	})
}

func TestGreaterOrEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		GreaterOrEqualTf(t, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualTf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterOrEqualT should call FailNow()")
		}
	})
}

func TestGreaterTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		GreaterTf(t, 2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterTf(mock, 1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterT should call FailNow()")
		}
	})
}

func TestHTTPBodyContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		HTTPBodyContainsf(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPBodyContains should call FailNow()")
		}
	})
}

func TestHTTPBodyNotContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		HTTPBodyNotContainsf(t, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyNotContainsf(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPBodyNotContains should call FailNow()")
		}
	})
}

func TestHTTPErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		HTTPErrorf(t, httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPErrorf(mock, httpOK, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPError should call FailNow()")
		}
	})
}

func TestHTTPRedirectf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		HTTPRedirectf(t, httpRedirect, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPRedirectf(mock, httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPRedirect should call FailNow()")
		}
	})
}

func TestHTTPStatusCodef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		HTTPStatusCodef(t, httpOK, "GET", "/", nil, http.StatusOK, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPStatusCodef(mock, httpError, "GET", "/", nil, http.StatusOK, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPStatusCode should call FailNow()")
		}
	})
}

func TestHTTPSuccessf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		HTTPSuccessf(t, httpOK, "GET", "/", nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPSuccessf(mock, httpError, "GET", "/", nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPSuccess should call FailNow()")
		}
	})
}

func TestImplementsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Implementsf(t, ptr(dummyInterface), new(testing.T), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Implementsf(mock, (*error)(nil), new(testing.T), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Implements should call FailNow()")
		}
	})
}

func TestInDeltaf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		InDeltaf(t, 1.0, 1.01, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaf(mock, 1.0, 1.1, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDelta should call FailNow()")
		}
	})
}

func TestInDeltaMapValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		InDeltaMapValuesf(t, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaMapValuesf(mock, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaMapValues should call FailNow()")
		}
	})
}

func TestInDeltaSlicef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		InDeltaSlicef(t, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaSlicef(mock, []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaSlice should call FailNow()")
		}
	})
}

func TestInDeltaTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		InDeltaTf(t, 1.0, 1.01, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaTf(mock, 1.0, 1.1, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaT should call FailNow()")
		}
	})
}

func TestInEpsilonf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		InEpsilonf(t, 100.0, 101.0, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonf(mock, 100.0, 110.0, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilon should call FailNow()")
		}
	})
}

func TestInEpsilonSlicef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		InEpsilonSlicef(t, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonSlicef(mock, []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilonSlice should call FailNow()")
		}
	})
}

func TestInEpsilonTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		InEpsilonTf(t, 100.0, 101.0, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonTf(mock, 100.0, 110.0, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilonT should call FailNow()")
		}
	})
}

func TestIsDecreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsDecreasingf(t, []int{3, 2, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsDecreasing should call FailNow()")
		}
	})
}

func TestIsDecreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsDecreasingTf(t, []int{3, 2, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingTf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsDecreasingT should call FailNow()")
		}
	})
}

func TestIsIncreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsIncreasingf(t, []int{1, 2, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingf(mock, []int{1, 1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsIncreasing should call FailNow()")
		}
	})
}

func TestIsIncreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsIncreasingTf(t, []int{1, 2, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingTf(mock, []int{1, 1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsIncreasingT should call FailNow()")
		}
	})
}

func TestIsNonDecreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsNonDecreasingf(t, []int{1, 1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingf(mock, []int{2, 1, 0}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonDecreasing should call FailNow()")
		}
	})
}

func TestIsNonDecreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsNonDecreasingTf(t, []int{1, 1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingTf(mock, []int{2, 1, 0}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonDecreasingT should call FailNow()")
		}
	})
}

func TestIsNonIncreasingf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsNonIncreasingf(t, []int{2, 1, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonIncreasing should call FailNow()")
		}
	})
}

func TestIsNonIncreasingTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsNonIncreasingTf(t, []int{2, 1, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingTf(mock, []int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonIncreasingT should call FailNow()")
		}
	})
}

func TestIsNotTypef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsNotTypef(t, int32(123), int64(456), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotTypef(mock, 123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNotType should call FailNow()")
		}
	})
}

func TestIsTypef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		IsTypef(t, 123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsTypef(mock, int32(123), int64(456), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsType should call FailNow()")
		}
	})
}

func TestJSONEqf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		JSONEqf(t, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqf(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEq should call FailNow()")
		}
	})
}

func TestJSONEqBytesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		JSONEqBytesf(t, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqBytesf(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEqBytes should call FailNow()")
		}
	})
}

func TestJSONEqTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		JSONEqTf(t, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqTf(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEqT should call FailNow()")
		}
	})
}

func TestKindf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Kindf(t, reflect.String, "hello", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Kindf(mock, reflect.String, 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Kind should call FailNow()")
		}
	})
}

func TestLenf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Lenf(t, []string{"A", "B"}, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Lenf(mock, []string{"A", "B"}, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Len should call FailNow()")
		}
	})
}

func TestLessf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Lessf(t, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Lessf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Less should call FailNow()")
		}
	})
}

func TestLessOrEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		LessOrEqualf(t, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessOrEqual should call FailNow()")
		}
	})
}

func TestLessOrEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		LessOrEqualTf(t, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualTf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessOrEqualT should call FailNow()")
		}
	})
}

func TestLessTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		LessTf(t, 1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessTf(mock, 2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessT should call FailNow()")
		}
	})
}

func TestMapContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		MapContainsTf(t, map[string]string{"A": "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapContainsTf(mock, map[string]string{"A": "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("MapContainsT should call FailNow()")
		}
	})
}

func TestMapNotContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		MapNotContainsTf(t, map[string]string{"A": "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapNotContainsTf(mock, map[string]string{"A": "B"}, "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("MapNotContainsT should call FailNow()")
		}
	})
}

func TestNegativef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Negativef(t, -1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Negativef(mock, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Negative should call FailNow()")
		}
	})
}

func TestNegativeTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NegativeTf(t, -1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NegativeTf(mock, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NegativeT should call FailNow()")
		}
	})
}

func TestNeverf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Neverf(t, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Neverf(mock, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Never should call FailNow()")
		}
	})
}

func TestNilf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Nilf(t, nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Nilf(mock, "not nil", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Nil should call FailNow()")
		}
	})
}

func TestNoErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NoErrorf(t, nil, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoErrorf(mock, ErrTest, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NoError should call FailNow()")
		}
	})
}

func TestNotContainsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotContainsf(t, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotContainsf(mock, []string{"A", "B"}, "B", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotContains should call FailNow()")
		}
	})
}

func TestNotElementsMatchf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotElementsMatchf(t, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotElementsMatch should call FailNow()")
		}
	})
}

func TestNotElementsMatchTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotElementsMatchTf(t, []int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchTf(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotElementsMatchT should call FailNow()")
		}
	})
}

func TestNotEmptyf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotEmptyf(t, "not empty", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEmptyf(mock, "", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEmpty should call FailNow()")
		}
	})
}

func TestNotEqualf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotEqualf(t, 123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualf(mock, 123, 123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqual should call FailNow()")
		}
	})
}

func TestNotEqualTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotEqualTf(t, 123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualTf(mock, 123, 123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqualT should call FailNow()")
		}
	})
}

func TestNotEqualValuesf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotEqualValuesf(t, uint32(123), int32(456), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualValuesf(mock, uint32(123), int32(123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqualValues should call FailNow()")
		}
	})
}

func TestNotErrorAsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotErrorAsf(t, ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAsf(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorAs should call FailNow()")
		}
	})
}

func TestNotErrorIsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotErrorIsf(t, ErrTest, io.EOF, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorIsf(mock, fmt.Errorf("wrap: %w", io.EOF), io.EOF, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorIs should call FailNow()")
		}
	})
}

func TestNotImplementsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotImplementsf(t, (*error)(nil), new(testing.T), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotImplementsf(mock, ptr(dummyInterface), new(testing.T), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotImplements should call FailNow()")
		}
	})
}

func TestNotKindf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotKindf(t, reflect.String, 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotKindf(mock, reflect.String, "hello", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotKind should call FailNow()")
		}
	})
}

func TestNotNilf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotNilf(t, "not nil", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotNilf(mock, nil, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotNil should call FailNow()")
		}
	})
}

func TestNotPanicsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotPanicsf(t, func() {}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotPanicsf(mock, func() { panic("panicking") }, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotPanics should call FailNow()")
		}
	})
}

func TestNotRegexpf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotRegexpf(t, "^start", "not starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpf(mock, "^start", "starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotRegexp should call FailNow()")
		}
	})
}

func TestNotRegexpTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotRegexpTf(t, "^start", "not starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpTf(mock, "^start", "starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotRegexpT should call FailNow()")
		}
	})
}

func TestNotSamef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotSamef(t, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSamef(mock, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSame should call FailNow()")
		}
	})
}

func TestNotSameTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotSameTf(t, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSameTf(mock, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSameT should call FailNow()")
		}
	})
}

func TestNotSortedTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotSortedTf(t, []int{3, 1, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSortedTf(mock, []int{1, 4, 8}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSortedT should call FailNow()")
		}
	})
}

func TestNotSubsetf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotSubsetf(t, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSubsetf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSubset should call FailNow()")
		}
	})
}

func TestNotZerof(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		NotZerof(t, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotZerof(mock, 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotZero should call FailNow()")
		}
	})
}

func TestPanicsf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Panicsf(t, func() { panic("panicking") }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panicsf(mock, func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Panics should call FailNow()")
		}
	})
}

func TestPanicsWithErrorf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		PanicsWithErrorf(t, ErrTest.Error(), func() { panic(ErrTest) }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithErrorf(mock, ErrTest.Error(), func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("PanicsWithError should call FailNow()")
		}
	})
}

func TestPanicsWithValuef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		PanicsWithValuef(t, "panicking", func() { panic("panicking") }, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithValuef(mock, "panicking", func() {}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("PanicsWithValue should call FailNow()")
		}
	})
}

func TestPositivef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Positivef(t, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Positivef(mock, -1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Positive should call FailNow()")
		}
	})
}

func TestPositiveTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		PositiveTf(t, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PositiveTf(mock, -1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("PositiveT should call FailNow()")
		}
	})
}

func TestRegexpf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Regexpf(t, "^start", "starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Regexpf(mock, "^start", "not starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Regexp should call FailNow()")
		}
	})
}

func TestRegexpTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		RegexpTf(t, "^start", "starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		RegexpTf(mock, "^start", "not starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("RegexpT should call FailNow()")
		}
	})
}

func TestSamef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Samef(t, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Samef(mock, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Same should call FailNow()")
		}
	})
}

func TestSameTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		SameTf(t, &staticVar, staticVarPtr, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SameTf(mock, &staticVar, ptr("static string"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SameT should call FailNow()")
		}
	})
}

func TestSliceContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		SliceContainsTf(t, []string{"A", "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceContainsTf(mock, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceContainsT should call FailNow()")
		}
	})
}

func TestSliceNotContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		SliceNotContainsTf(t, []string{"A", "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotContainsTf(mock, []string{"A", "B"}, "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceNotContainsT should call FailNow()")
		}
	})
}

func TestSliceNotSubsetTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		SliceNotSubsetTf(t, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotSubsetTf(mock, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceNotSubsetT should call FailNow()")
		}
	})
}

func TestSliceSubsetTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		SliceSubsetTf(t, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceSubsetTf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceSubsetT should call FailNow()")
		}
	})
}

func TestSortedTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		SortedTf(t, []int{1, 1, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SortedTf(mock, []int{1, 4, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SortedT should call FailNow()")
		}
	})
}

func TestStringContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		StringContainsTf(t, "AB", "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringContainsTf(mock, "AB", "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("StringContainsT should call FailNow()")
		}
	})
}

func TestStringNotContainsTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		StringNotContainsTf(t, "AB", "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringNotContainsTf(mock, "AB", "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("StringNotContainsT should call FailNow()")
		}
	})
}

func TestSubsetf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Subsetf(t, []int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Subsetf(mock, []int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Subset should call FailNow()")
		}
	})
}

func TestTruef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		Truef(t, 1 == 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Truef(mock, 1 == 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("True should call FailNow()")
		}
	})
}

func TestTrueTf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		TrueTf(t, 1 == 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		TrueTf(mock, 1 == 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("TrueT should call FailNow()")
		}
	})
}

func TestWithinDurationf(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		WithinDurationf(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinDurationf(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("WithinDuration should call FailNow()")
		}
	})
}

func TestWithinRangef(t *testing.T) {
	t.Parallel()
	t.Run("success", func(t *testing.T) {
		t.Parallel()
		WithinRangef(t, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinRangef(mock, time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("WithinRange should call FailNow()")
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
		Zerof(t, 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Zerof(mock, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Zero should call FailNow()")
		}
	})
}
