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

func TestCondition(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Condition(mock, func() bool { return true })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Condition(mock, func() bool { return false })
		// require functions don't return a value
		if !mock.failed {
			t.Error("Condition should call FailNow()")
		}
	})
}

func TestContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Contains(mock, []string{"A", "B"}, "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Contains(mock, []string{"A", "B"}, "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Contains should call FailNow()")
		}
	})
}

func TestDirExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirExists(mock, filepath.Join(testDataPath(), "existing_dir"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirExists(mock, filepath.Join(testDataPath(), "non_existing_dir"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("DirExists should call FailNow()")
		}
	})
}

func TestDirNotExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirNotExists(mock, filepath.Join(testDataPath(), "non_existing_dir"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		DirNotExists(mock, filepath.Join(testDataPath(), "existing_dir"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("DirNotExists should call FailNow()")
		}
	})
}

func TestElementsMatch(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatch(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatch(mock, []int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
		if !mock.failed {
			t.Error("ElementsMatch should call FailNow()")
		}
	})
}

func TestElementsMatchT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchT(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ElementsMatchT(mock, []int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
		if !mock.failed {
			t.Error("ElementsMatchT should call FailNow()")
		}
	})
}

func TestEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Empty(mock, "")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Empty(mock, "not empty")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Empty should call FailNow()")
		}
	})
}

func TestEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Equal(mock, 123, 123)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Equal(mock, 123, 456)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Equal should call FailNow()")
		}
	})
}

func TestEqualError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualError(mock, ErrTest, "assert.ErrTest general error for testing")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualError(mock, ErrTest, "wrong error message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualError should call FailNow()")
		}
	})
}

func TestEqualExportedValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualExportedValues(mock, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "a", b: 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualExportedValues(mock, &dummyStruct{A: "a", b: 1}, &dummyStruct{A: "b", b: 1})
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualExportedValues should call FailNow()")
		}
	})
}

func TestEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualT(mock, 123, 123)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualT(mock, 123, 456)
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualT should call FailNow()")
		}
	})
}

func TestEqualValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualValues(mock, uint32(123), int32(123))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EqualValues(mock, uint32(123), int32(456))
		// require functions don't return a value
		if !mock.failed {
			t.Error("EqualValues should call FailNow()")
		}
	})
}

func TestError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Error(mock, ErrTest)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Error(mock, nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Error should call FailNow()")
		}
	})
}

func TestErrorAs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAs(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAs(mock, ErrTest, new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorAs should call FailNow()")
		}
	})
}

func TestErrorContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorContains(mock, ErrTest, "general error")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorContains(mock, ErrTest, "not in message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorContains should call FailNow()")
		}
	})
}

func TestErrorIs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorIs(mock, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorIs(mock, ErrTest, io.EOF)
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorIs should call FailNow()")
		}
	})
}

func TestEventually(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Eventually(mock, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Eventually(mock, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Eventually should call FailNow()")
		}
	})
}

func TestEventuallyWith(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EventuallyWith(mock, func(c *CollectT) { True(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		EventuallyWith(mock, func(c *CollectT) { False(c, true) }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("EventuallyWith should call FailNow()")
		}
	})
}

func TestExactly(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Exactly(mock, int32(123), int32(123))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Exactly(mock, int32(123), int64(123))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Exactly should call FailNow()")
		}
	})
}

func TestFail(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Fail(mock, "failed")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Fail should call FailNow()")
		}
	})
}

func TestFailNow(t *testing.T) {
	t.Parallel()

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FailNow(mock, "failed")
		// require functions don't return a value
		if !mock.failed {
			t.Error("FailNow should call FailNow()")
		}
	})
}

func TestFalse(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		False(mock, 1 == 0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		False(mock, 1 == 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("False should call FailNow()")
		}
	})
}

func TestFalseT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FalseT(mock, 1 == 0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FalseT(mock, 1 == 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("FalseT should call FailNow()")
		}
	})
}

func TestFileEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileEmpty(mock, filepath.Join(testDataPath(), "empty_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileEmpty(mock, filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileEmpty should call FailNow()")
		}
	})
}

func TestFileExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileExists(mock, filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileExists(mock, filepath.Join(testDataPath(), "non_existing_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileExists should call FailNow()")
		}
	})
}

func TestFileNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotEmpty(mock, filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotEmpty(mock, filepath.Join(testDataPath(), "empty_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileNotEmpty should call FailNow()")
		}
	})
}

func TestFileNotExists(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotExists(mock, filepath.Join(testDataPath(), "non_existing_file"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		FileNotExists(mock, filepath.Join(testDataPath(), "existing_file"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("FileNotExists should call FailNow()")
		}
	})
}

func TestGreater(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Greater(mock, 2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Greater(mock, 1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Greater should call FailNow()")
		}
	})
}

func TestGreaterOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqual(mock, 2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqual(mock, 1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterOrEqual should call FailNow()")
		}
	})
}

func TestGreaterOrEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualT(mock, 2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterOrEqualT(mock, 1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterOrEqualT should call FailNow()")
		}
	})
}

func TestGreaterT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterT(mock, 2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		GreaterT(mock, 1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("GreaterT should call FailNow()")
		}
	})
}

func TestHTTPBodyContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyContains(mock, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, World!")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyContains(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, World!")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPBodyContains should call FailNow()")
		}
	})
}

func TestHTTPBodyNotContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyNotContains(mock, httpBody, "GET", "/", url.Values{"name": []string{"World"}}, "Hello, Bob!")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPBodyNotContains(mock, httpBody, "GET", "/", url.Values{"name": []string{"Bob"}}, "Hello, Bob!")
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPBodyNotContains should call FailNow()")
		}
	})
}

func TestHTTPError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPError(mock, httpError, "GET", "/", nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPError(mock, httpOK, "GET", "/", nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPError should call FailNow()")
		}
	})
}

func TestHTTPRedirect(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPRedirect(mock, httpRedirect, "GET", "/", nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPRedirect(mock, httpError, "GET", "/", nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPRedirect should call FailNow()")
		}
	})
}

func TestHTTPStatusCode(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPStatusCode(mock, httpOK, "GET", "/", nil, http.StatusOK)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPStatusCode(mock, httpError, "GET", "/", nil, http.StatusOK)
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPStatusCode should call FailNow()")
		}
	})
}

func TestHTTPSuccess(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPSuccess(mock, httpOK, "GET", "/", nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		HTTPSuccess(mock, httpError, "GET", "/", nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("HTTPSuccess should call FailNow()")
		}
	})
}

func TestImplements(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Implements(mock, ptr(dummyInterface), new(testing.T))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Implements(mock, (*error)(nil), new(testing.T))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Implements should call FailNow()")
		}
	})
}

func TestInDelta(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDelta(mock, 1.0, 1.01, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDelta(mock, 1.0, 1.1, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDelta should call FailNow()")
		}
	})
}

func TestInDeltaMapValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaMapValues(mock, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.01}, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaMapValues(mock, map[string]float64{"a": 1.0}, map[string]float64{"a": 1.1}, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaMapValues should call FailNow()")
		}
	})
}

func TestInDeltaSlice(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaSlice(mock, []float64{1.0, 2.0}, []float64{1.01, 2.01}, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaSlice(mock, []float64{1.0, 2.0}, []float64{1.1, 2.1}, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaSlice should call FailNow()")
		}
	})
}

func TestInDeltaT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaT(mock, 1.0, 1.01, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InDeltaT(mock, 1.0, 1.1, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("InDeltaT should call FailNow()")
		}
	})
}

func TestInEpsilon(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilon(mock, 100.0, 101.0, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilon(mock, 100.0, 110.0, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilon should call FailNow()")
		}
	})
}

func TestInEpsilonSlice(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonSlice(mock, []float64{100.0, 200.0}, []float64{101.0, 202.0}, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonSlice(mock, []float64{100.0, 200.0}, []float64{110.0, 220.0}, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilonSlice should call FailNow()")
		}
	})
}

func TestInEpsilonT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonT(mock, 100.0, 101.0, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		InEpsilonT(mock, 100.0, 110.0, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("InEpsilonT should call FailNow()")
		}
	})
}

func TestIsDecreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasing(mock, []int{3, 2, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasing(mock, []int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsDecreasing should call FailNow()")
		}
	})
}

func TestIsDecreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingT(mock, []int{3, 2, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsDecreasingT(mock, []int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsDecreasingT should call FailNow()")
		}
	})
}

func TestIsIncreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasing(mock, []int{1, 2, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasing(mock, []int{1, 1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsIncreasing should call FailNow()")
		}
	})
}

func TestIsIncreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingT(mock, []int{1, 2, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsIncreasingT(mock, []int{1, 1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsIncreasingT should call FailNow()")
		}
	})
}

func TestIsNonDecreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasing(mock, []int{1, 1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasing(mock, []int{2, 1, 0})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonDecreasing should call FailNow()")
		}
	})
}

func TestIsNonDecreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingT(mock, []int{1, 1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonDecreasingT(mock, []int{2, 1, 0})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonDecreasingT should call FailNow()")
		}
	})
}

func TestIsNonIncreasing(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasing(mock, []int{2, 1, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasing(mock, []int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonIncreasing should call FailNow()")
		}
	})
}

func TestIsNonIncreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingT(mock, []int{2, 1, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNonIncreasingT(mock, []int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNonIncreasingT should call FailNow()")
		}
	})
}

func TestIsNotOfTypeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotOfTypeT[myType](mock, 123.123)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotOfTypeT[myType](mock, myType(123.123))
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNotOfTypeT should call FailNow()")
		}
	})
}

func TestIsNotType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotType(mock, int32(123), int64(456))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsNotType(mock, 123, 456)
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsNotType should call FailNow()")
		}
	})
}

func TestIsOfTypeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsOfTypeT[myType](mock, myType(123.123))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsOfTypeT[myType](mock, 123.123)
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsOfTypeT should call FailNow()")
		}
	})
}

func TestIsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsType(mock, 123, 456)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		IsType(mock, int32(123), int64(456))
		// require functions don't return a value
		if !mock.failed {
			t.Error("IsType should call FailNow()")
		}
	})
}

func TestJSONEq(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEq(mock, `{"hello": "world", "foo": "bar"}`, `{"foo": "bar", "hello": "world"}`)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEq(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEq should call FailNow()")
		}
	})
}

func TestJSONEqBytes(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqBytes(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`{"foo": "bar", "hello": "world"}`))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqBytes(mock, []byte(`{"hello": "world", "foo": "bar"}`), []byte(`[{"foo": "bar"}, {"hello": "world"}]`))
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEqBytes should call FailNow()")
		}
	})
}

func TestJSONEqT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqT(mock, `{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONEqT(mock, `{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONEqT should call FailNow()")
		}
	})
}

func TestJSONMarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONMarshalAsT(mock, []byte(`{"A": "a"}`), dummyStruct{A: "a"})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONMarshalAsT(mock, `[{"foo": "bar"}, {"hello": "world"}]`, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONMarshalAsT should call FailNow()")
		}
	})
}

func TestJSONUnmarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONUnmarshalAsT(mock, dummyStruct{A: "a"}, []byte(`{"A": "a"}`))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		JSONUnmarshalAsT(mock, 1, `[{"foo": "bar"}, {"hello": "world"}]`)
		// require functions don't return a value
		if !mock.failed {
			t.Error("JSONUnmarshalAsT should call FailNow()")
		}
	})
}

func TestKind(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Kind(mock, reflect.String, "hello")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Kind(mock, reflect.String, 0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Kind should call FailNow()")
		}
	})
}

func TestLen(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Len(mock, []string{"A", "B"}, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Len(mock, []string{"A", "B"}, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Len should call FailNow()")
		}
	})
}

func TestLess(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Less(mock, 1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Less(mock, 2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Less should call FailNow()")
		}
	})
}

func TestLessOrEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqual(mock, 1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqual(mock, 2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessOrEqual should call FailNow()")
		}
	})
}

func TestLessOrEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualT(mock, 1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessOrEqualT(mock, 2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessOrEqualT should call FailNow()")
		}
	})
}

func TestLessT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessT(mock, 1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		LessT(mock, 2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("LessT should call FailNow()")
		}
	})
}

func TestMapContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapContainsT(mock, map[string]string{"A": "B"}, "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapContainsT(mock, map[string]string{"A": "B"}, "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("MapContainsT should call FailNow()")
		}
	})
}

func TestMapNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapNotContainsT(mock, map[string]string{"A": "B"}, "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		MapNotContainsT(mock, map[string]string{"A": "B"}, "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("MapNotContainsT should call FailNow()")
		}
	})
}

func TestNegative(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Negative(mock, -1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Negative(mock, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Negative should call FailNow()")
		}
	})
}

func TestNegativeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NegativeT(mock, -1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NegativeT(mock, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NegativeT should call FailNow()")
		}
	})
}

func TestNever(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Never(mock, func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Never(mock, func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Never should call FailNow()")
		}
	})
}

func TestNil(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Nil(mock, nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Nil(mock, "not nil")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Nil should call FailNow()")
		}
	})
}

func TestNoError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoError(mock, nil)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoError(mock, ErrTest)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NoError should call FailNow()")
		}
	})
}

func TestNoFileDescriptorLeak(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoFileDescriptorLeak(mock, func() {})
		// require functions don't return a value
	})
}

func TestNoGoRoutineLeak(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NoGoRoutineLeak(mock, func() {})
		// require functions don't return a value
	})
}

func TestNotContains(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotContains(mock, []string{"A", "B"}, "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotContains(mock, []string{"A", "B"}, "B")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotContains should call FailNow()")
		}
	})
}

func TestNotElementsMatch(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatch(mock, []int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatch(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotElementsMatch should call FailNow()")
		}
	})
}

func TestNotElementsMatchT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchT(mock, []int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotElementsMatchT(mock, []int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotElementsMatchT should call FailNow()")
		}
	})
}

func TestNotEmpty(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEmpty(mock, "not empty")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEmpty(mock, "")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEmpty should call FailNow()")
		}
	})
}

func TestNotEqual(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqual(mock, 123, 456)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqual(mock, 123, 123)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqual should call FailNow()")
		}
	})
}

func TestNotEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualT(mock, 123, 456)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualT(mock, 123, 123)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqualT should call FailNow()")
		}
	})
}

func TestNotEqualValues(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualValues(mock, uint32(123), int32(456))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotEqualValues(mock, uint32(123), int32(123))
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotEqualValues should call FailNow()")
		}
	})
}

func TestNotErrorAs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAs(mock, ErrTest, new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAs(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorAs should call FailNow()")
		}
	})
}

func TestNotErrorIs(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorIs(mock, ErrTest, io.EOF)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorIs(mock, fmt.Errorf("wrap: %w", io.EOF), io.EOF)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorIs should call FailNow()")
		}
	})
}

func TestNotImplements(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotImplements(mock, (*error)(nil), new(testing.T))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotImplements(mock, ptr(dummyInterface), new(testing.T))
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotImplements should call FailNow()")
		}
	})
}

func TestNotKind(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotKind(mock, reflect.String, 0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotKind(mock, reflect.String, "hello")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotKind should call FailNow()")
		}
	})
}

func TestNotNil(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotNil(mock, "not nil")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotNil(mock, nil)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotNil should call FailNow()")
		}
	})
}

func TestNotPanics(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotPanics(mock, func() {})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotPanics(mock, func() { panic("panicking") })
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotPanics should call FailNow()")
		}
	})
}

func TestNotRegexp(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexp(mock, "^start", "not starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexp(mock, "^start", "starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotRegexp should call FailNow()")
		}
	})
}

func TestNotRegexpT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpT(mock, "^start", "not starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotRegexpT(mock, "^start", "starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotRegexpT should call FailNow()")
		}
	})
}

func TestNotSame(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSame(mock, &staticVar, ptr("static string"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSame(mock, &staticVar, staticVarPtr)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSame should call FailNow()")
		}
	})
}

func TestNotSameT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSameT(mock, &staticVar, ptr("static string"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSameT(mock, &staticVar, staticVarPtr)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSameT should call FailNow()")
		}
	})
}

func TestNotSortedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSortedT(mock, []int{3, 1, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSortedT(mock, []int{1, 4, 8})
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSortedT should call FailNow()")
		}
	})
}

func TestNotSubset(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSubset(mock, []int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotSubset(mock, []int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotSubset should call FailNow()")
		}
	})
}

func TestNotZero(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotZero(mock, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotZero(mock, 0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotZero should call FailNow()")
		}
	})
}

func TestPanics(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(mock, func() { panic("panicking") })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(mock, func() {})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Panics should call FailNow()")
		}
	})
}

func TestPanicsWithError(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithError(mock, ErrTest.Error(), func() { panic(ErrTest) })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithError(mock, ErrTest.Error(), func() {})
		// require functions don't return a value
		if !mock.failed {
			t.Error("PanicsWithError should call FailNow()")
		}
	})
}

func TestPanicsWithValue(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithValue(mock, "panicking", func() { panic("panicking") })
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PanicsWithValue(mock, "panicking", func() {})
		// require functions don't return a value
		if !mock.failed {
			t.Error("PanicsWithValue should call FailNow()")
		}
	})
}

func TestPositive(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Positive(mock, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Positive(mock, -1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Positive should call FailNow()")
		}
	})
}

func TestPositiveT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PositiveT(mock, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		PositiveT(mock, -1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("PositiveT should call FailNow()")
		}
	})
}

func TestRegexp(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Regexp(mock, "^start", "starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Regexp(mock, "^start", "not starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Regexp should call FailNow()")
		}
	})
}

func TestRegexpT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		RegexpT(mock, "^start", "starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		RegexpT(mock, "^start", "not starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("RegexpT should call FailNow()")
		}
	})
}

func TestSame(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Same(mock, &staticVar, staticVarPtr)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Same(mock, &staticVar, ptr("static string"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Same should call FailNow()")
		}
	})
}

func TestSameT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SameT(mock, &staticVar, staticVarPtr)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SameT(mock, &staticVar, ptr("static string"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("SameT should call FailNow()")
		}
	})
}

func TestSeqContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqContainsT(mock, slices.Values([]string{"A", "B"}), "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqContainsT(mock, slices.Values([]string{"A", "B"}), "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SeqContainsT should call FailNow()")
		}
	})
}

func TestSeqNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqNotContainsT(mock, slices.Values([]string{"A", "B"}), "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SeqNotContainsT(mock, slices.Values([]string{"A", "B"}), "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SeqNotContainsT should call FailNow()")
		}
	})
}

func TestSliceContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceContainsT(mock, []string{"A", "B"}, "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceContainsT(mock, []string{"A", "B"}, "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceContainsT should call FailNow()")
		}
	})
}

func TestSliceNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotContainsT(mock, []string{"A", "B"}, "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotContainsT(mock, []string{"A", "B"}, "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceNotContainsT should call FailNow()")
		}
	})
}

func TestSliceNotSubsetT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotSubsetT(mock, []int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceNotSubsetT(mock, []int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceNotSubsetT should call FailNow()")
		}
	})
}

func TestSliceSubsetT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceSubsetT(mock, []int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SliceSubsetT(mock, []int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
		if !mock.failed {
			t.Error("SliceSubsetT should call FailNow()")
		}
	})
}

func TestSortedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SortedT(mock, []int{1, 1, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		SortedT(mock, []int{1, 4, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("SortedT should call FailNow()")
		}
	})
}

func TestStringContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringContainsT(mock, "AB", "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringContainsT(mock, "AB", "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("StringContainsT should call FailNow()")
		}
	})
}

func TestStringNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringNotContainsT(mock, "AB", "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		StringNotContainsT(mock, "AB", "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("StringNotContainsT should call FailNow()")
		}
	})
}

func TestSubset(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Subset(mock, []int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Subset(mock, []int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Subset should call FailNow()")
		}
	})
}

func TestTrue(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		True(mock, 1 == 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		True(mock, 1 == 0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("True should call FailNow()")
		}
	})
}

func TestTrueT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		TrueT(mock, 1 == 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		TrueT(mock, 1 == 0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("TrueT should call FailNow()")
		}
	})
}

func TestWithinDuration(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinDuration(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinDuration(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second)
		// require functions don't return a value
		if !mock.failed {
			t.Error("WithinDuration should call FailNow()")
		}
	})
}

func TestWithinRange(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinRange(mock, time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		WithinRange(mock, time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC))
		// require functions don't return a value
		if !mock.failed {
			t.Error("WithinRange should call FailNow()")
		}
	})
}

func TestYAMLEq(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLEq(mock, "key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLEq should panic as expected")
		}
	})
}

func TestYAMLEqBytes(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLEqBytes(mock, []byte("key: value"), []byte("key: value"))
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLEqBytes should panic as expected")
		}
	})
}

func TestYAMLEqT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLEqT(mock, "key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLEqT should panic as expected")
		}
	})
}

func TestYAMLMarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLMarshalAsT(mock, "key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLMarshalAsT should panic as expected")
		}
	})
}

func TestYAMLUnmarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Panics(t, func() {
			YAMLUnmarshalAsT(mock, "key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("YAMLUnmarshalAsT should panic as expected")
		}
	})
}

func TestZero(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Zero(mock, 0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		Zero(mock, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Zero should call FailNow()")
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
