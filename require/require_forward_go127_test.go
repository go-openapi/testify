// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.27

package require

import (
	"fmt"
	"slices"
	"testing"
	"time"
)

func TestAssertionsBlockedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.BlockedT(make(chan struct{}))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.BlockedT(sendChanMessage())
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.BlockedT should call FailNow()")
		}
	})
}

func TestAssertionsConsistently(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Consistently(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Consistently(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Consistently should call FailNow()")
		}
	})
}

func TestAssertionsElementsMatchT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatchT([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatchT([]int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ElementsMatchT should call FailNow()")
		}
	})
}

func TestAssertionsEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualT(123, 123)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualT(123, 456)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualT should call FailNow()")
		}
	})
}

func TestAssertionsErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAsType(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAsType(ErrTest, new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorAsType should call FailNow()")
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

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EventuallyWith(func(c *CollectT) { c.Cancel() }, 100*time.Millisecond, 20*time.Millisecond)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EventuallyWith should call FailNow()")
		}
	})
}

func TestAssertionsFalseT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FalseT(1 == 0)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FalseT(1 == 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FalseT should call FailNow()")
		}
	})
}

func TestAssertionsGreaterOrEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqualT(2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqualT(1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqualT should call FailNow()")
		}
	})
}

func TestAssertionsGreaterT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterT(2, 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterT(1, 2)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.GreaterT should call FailNow()")
		}
	})
}

func TestAssertionsInDeltaT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaT(1.0, 1.01, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaT(1.0, 1.1, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDeltaT should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilonSymmetricT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSymmetricT(100.0, 101.0, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSymmetricT(100.0, 110.0, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilonSymmetricT should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilonT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonT(100.0, 101.0, 0.02)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonT(100.0, 110.0, 0.05)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilonT should call FailNow()")
		}
	})
}

func TestAssertionsIsDecreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasingT([]int{3, 2, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasingT([]int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsDecreasingT should call FailNow()")
		}
	})
}

func TestAssertionsIsIncreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasingT([]int{1, 2, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasingT([]int{1, 1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsIncreasingT should call FailNow()")
		}
	})
}

func TestAssertionsIsNonDecreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasingT([]int{1, 1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasingT([]int{2, 1, 0})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasingT should call FailNow()")
		}
	})
}

func TestAssertionsIsNonIncreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasingT([]int{2, 1, 1})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasingT([]int{1, 2, 3})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasingT should call FailNow()")
		}
	})
}

func TestAssertionsIsNotOfTypeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotOfTypeT[myType](123.123)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotOfTypeT[myType](myType(123.123))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNotOfTypeT should call FailNow()")
		}
	})
}

func TestAssertionsIsOfTypeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsOfTypeT[myType](myType(123.123))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsOfTypeT[myType](123.123)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsOfTypeT should call FailNow()")
		}
	})
}

func TestAssertionsJSONEqT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqT(`{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqT(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONEqT should call FailNow()")
		}
	})
}

func TestAssertionsJSONMarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONMarshalAsT([]byte(`{"A": "a"}`), dummyStruct{A: "a"})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONMarshalAsT(`[{"foo": "bar"}, {"hello": "world"}]`, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONMarshalAsT should call FailNow()")
		}
	})
}

func TestAssertionsJSONUnmarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONUnmarshalAsT(dummyStruct{A: "a"}, []byte(`{"A": "a"}`))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONUnmarshalAsT(1, `[{"foo": "bar"}, {"hello": "world"}]`)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONUnmarshalAsT should call FailNow()")
		}
	})
}

func TestAssertionsLessOrEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqualT(1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqualT(2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.LessOrEqualT should call FailNow()")
		}
	})
}

func TestAssertionsLessT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessT(1, 2)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessT(2, 1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.LessT should call FailNow()")
		}
	})
}

func TestAssertionsMapContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapContainsT(map[string]string{"A": "B"}, "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapContainsT(map[string]string{"A": "B"}, "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapContainsT should call FailNow()")
		}
	})
}

func TestAssertionsMapEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapEqualT(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapEqualT(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapEqualT should call FailNow()")
		}
	})
}

func TestAssertionsMapNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotContainsT(map[string]string{"A": "B"}, "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotContainsT(map[string]string{"A": "B"}, "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapNotContainsT should call FailNow()")
		}
	})
}

func TestAssertionsMapNotEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotEqualT(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotEqualT(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapNotEqualT should call FailNow()")
		}
	})
}

func TestAssertionsNegativeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NegativeT(-1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NegativeT(1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NegativeT should call FailNow()")
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

func TestAssertionsNotBlockedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotBlockedT(sendChanMessage())
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotBlockedT(make(chan struct{}))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotBlockedT should call FailNow()")
		}
	})
}

func TestAssertionsNotElementsMatchT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatchT([]int{1, 2, 3}, []int{1, 2, 4})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatchT([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotElementsMatchT should call FailNow()")
		}
	})
}

func TestAssertionsNotEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualT(123, 456)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualT(123, 123)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEqualT should call FailNow()")
		}
	})
}

func TestAssertionsNotErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAsType(ErrTest, new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAsType(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotErrorAsType should call FailNow()")
		}
	})
}

func TestAssertionsNotRegexpT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexpT("^start", "not starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexpT("^start", "starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotRegexpT should call FailNow()")
		}
	})
}

func TestAssertionsNotSameT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSameT(&staticVar, ptr("static string"))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSameT(&staticVar, staticVarPtr)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSameT should call FailNow()")
		}
	})
}

func TestAssertionsNotSortedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSortedT([]int{3, 1, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSortedT([]int{1, 4, 8})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSortedT should call FailNow()")
		}
	})
}

func TestAssertionsPositiveT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PositiveT(1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PositiveT(-1)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.PositiveT should call FailNow()")
		}
	})
}

func TestAssertionsRegexpT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.RegexpT("^start", "starting")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.RegexpT("^start", "not starting")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.RegexpT should call FailNow()")
		}
	})
}

func TestAssertionsSameT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SameT(&staticVar, staticVarPtr)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SameT(&staticVar, ptr("static string"))
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SameT should call FailNow()")
		}
	})
}

func TestAssertionsSeqContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqContainsT(slices.Values([]string{"A", "B"}), "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqContainsT(slices.Values([]string{"A", "B"}), "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SeqContainsT should call FailNow()")
		}
	})
}

func TestAssertionsSeqNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqNotContainsT(slices.Values([]string{"A", "B"}), "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqNotContainsT(slices.Values([]string{"A", "B"}), "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SeqNotContainsT should call FailNow()")
		}
	})
}

func TestAssertionsSliceContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceContainsT([]string{"A", "B"}, "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceContainsT([]string{"A", "B"}, "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceContainsT should call FailNow()")
		}
	})
}

func TestAssertionsSliceEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceEqualT([]string{"Hello", "World"}, []string{"Hello", "World"})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceEqualT([]string{"Hello", "World"}, []string{"Hello"})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceEqualT should call FailNow()")
		}
	})
}

func TestAssertionsSliceNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotContainsT([]string{"A", "B"}, "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotContainsT([]string{"A", "B"}, "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceNotContainsT should call FailNow()")
		}
	})
}

func TestAssertionsSliceNotEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotEqualT([]string{"Hello", "World"}, []string{"Hello"})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotEqualT([]string{"Hello", "World"}, []string{"Hello", "World"})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceNotEqualT should call FailNow()")
		}
	})
}

func TestAssertionsSliceNotSubsetT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotSubsetT([]int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotSubsetT([]int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceNotSubsetT should call FailNow()")
		}
	})
}

func TestAssertionsSliceSubsetT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceSubsetT([]int{1, 2, 3}, []int{1, 2})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceSubsetT([]int{1, 2, 3}, []int{4, 5})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceSubsetT should call FailNow()")
		}
	})
}

func TestAssertionsSortedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SortedT([]int{1, 1, 3})
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SortedT([]int{1, 4, 2})
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SortedT should call FailNow()")
		}
	})
}

func TestAssertionsStringContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringContainsT("AB", "A")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringContainsT("AB", "C")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.StringContainsT should call FailNow()")
		}
	})
}

func TestAssertionsStringNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringNotContainsT("AB", "C")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringNotContainsT("AB", "A")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.StringNotContainsT should call FailNow()")
		}
	})
}

func TestAssertionsTrueT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.TrueT(1 == 1)
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.TrueT(1 == 0)
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.TrueT should call FailNow()")
		}
	})
}

func TestAssertionsYAMLEqT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLEqT("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLEqT should panic as expected")
		}
	})
}

func TestAssertionsYAMLMarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLMarshalAsT("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLMarshalAsT should panic as expected")
		}
	})
}

func TestAssertionsYAMLUnmarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLUnmarshalAsT("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLUnmarshalAsT should panic as expected")
		}
	})
}

func TestAssertionsBlockedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.BlockedTf(make(chan struct{}), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.BlockedTf(sendChanMessage(), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.BlockedTf should call FailNow()")
		}
	})
}

func TestAssertionsConsistentlyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Consistentlyf(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Consistentlyf(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.Consistentlyf should call FailNow()")
		}
	})
}

func TestAssertionsElementsMatchTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatchTf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ElementsMatchTf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ElementsMatchTf should call FailNow()")
		}
	})
}

func TestAssertionsEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualTf(123, 123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EqualTf(123, 456, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EqualTf should call FailNow()")
		}
	})
}

func TestAssertionsErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAsTypef(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.ErrorAsTypef(ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.ErrorAsTypef should call FailNow()")
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

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.EventuallyWithf(func(c *CollectT) { c.Cancel() }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.EventuallyWithf should call FailNow()")
		}
	})
}

func TestAssertionsFalseTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FalseTf(1 == 0, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.FalseTf(1 == 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.FalseTf should call FailNow()")
		}
	})
}

func TestAssertionsGreaterOrEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqualTf(2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterOrEqualTf(1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqualTf should call FailNow()")
		}
	})
}

func TestAssertionsGreaterTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterTf(2, 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.GreaterTf(1, 2, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.GreaterTf should call FailNow()")
		}
	})
}

func TestAssertionsInDeltaTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaTf(1.0, 1.01, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InDeltaTf(1.0, 1.1, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InDeltaTf should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilonSymmetricTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSymmetricTf(100.0, 101.0, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonSymmetricTf(100.0, 110.0, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilonSymmetricTf should call FailNow()")
		}
	})
}

func TestAssertionsInEpsilonTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonTf(100.0, 101.0, 0.02, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.InEpsilonTf(100.0, 110.0, 0.05, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.InEpsilonTf should call FailNow()")
		}
	})
}

func TestAssertionsIsDecreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasingTf([]int{3, 2, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsDecreasingTf([]int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsDecreasingTf should call FailNow()")
		}
	})
}

func TestAssertionsIsIncreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasingTf([]int{1, 2, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsIncreasingTf([]int{1, 1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsIncreasingTf should call FailNow()")
		}
	})
}

func TestAssertionsIsNonDecreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasingTf([]int{1, 1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonDecreasingTf([]int{2, 1, 0}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasingTf should call FailNow()")
		}
	})
}

func TestAssertionsIsNonIncreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasingTf([]int{2, 1, 1}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNonIncreasingTf([]int{1, 2, 3}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasingTf should call FailNow()")
		}
	})
}

func TestAssertionsIsNotOfTypeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotOfTypeTf[myType](123.123, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsNotOfTypeTf[myType](myType(123.123), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsNotOfTypeTf should call FailNow()")
		}
	})
}

func TestAssertionsIsOfTypeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsOfTypeTf[myType](myType(123.123), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.IsOfTypeTf[myType](123.123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.IsOfTypeTf should call FailNow()")
		}
	})
}

func TestAssertionsJSONEqTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqTf(`{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONEqTf(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONEqTf should call FailNow()")
		}
	})
}

func TestAssertionsJSONMarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONMarshalAsTf([]byte(`{"A": "a"}`), dummyStruct{A: "a"}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONMarshalAsTf(`[{"foo": "bar"}, {"hello": "world"}]`, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONMarshalAsTf should call FailNow()")
		}
	})
}

func TestAssertionsJSONUnmarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONUnmarshalAsTf(dummyStruct{A: "a"}, []byte(`{"A": "a"}`), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.JSONUnmarshalAsTf(1, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.JSONUnmarshalAsTf should call FailNow()")
		}
	})
}

func TestAssertionsLessOrEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqualTf(1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessOrEqualTf(2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.LessOrEqualTf should call FailNow()")
		}
	})
}

func TestAssertionsLessTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessTf(1, 2, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.LessTf(2, 1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.LessTf should call FailNow()")
		}
	})
}

func TestAssertionsMapContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapContainsTf(map[string]string{"A": "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapContainsTf(map[string]string{"A": "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsMapEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapEqualTf(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapEqualTf(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapEqualTf should call FailNow()")
		}
	})
}

func TestAssertionsMapNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotContainsTf(map[string]string{"A": "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotContainsTf(map[string]string{"A": "B"}, "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapNotContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsMapNotEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotEqualTf(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.MapNotEqualTf(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.MapNotEqualTf should call FailNow()")
		}
	})
}

func TestAssertionsNegativeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NegativeTf(-1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NegativeTf(1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NegativeTf should call FailNow()")
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

func TestAssertionsNotBlockedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotBlockedTf(sendChanMessage(), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotBlockedTf(make(chan struct{}), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotBlockedTf should call FailNow()")
		}
	})
}

func TestAssertionsNotElementsMatchTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatchTf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotElementsMatchTf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotElementsMatchTf should call FailNow()")
		}
	})
}

func TestAssertionsNotEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualTf(123, 456, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotEqualTf(123, 123, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotEqualTf should call FailNow()")
		}
	})
}

func TestAssertionsNotErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAsTypef(ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotErrorAsTypef(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotErrorAsTypef should call FailNow()")
		}
	})
}

func TestAssertionsNotRegexpTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexpTf("^start", "not starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotRegexpTf("^start", "starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotRegexpTf should call FailNow()")
		}
	})
}

func TestAssertionsNotSameTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSameTf(&staticVar, ptr("static string"), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSameTf(&staticVar, staticVarPtr, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSameTf should call FailNow()")
		}
	})
}

func TestAssertionsNotSortedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSortedTf([]int{3, 1, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.NotSortedTf([]int{1, 4, 8}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.NotSortedTf should call FailNow()")
		}
	})
}

func TestAssertionsPositiveTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PositiveTf(1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.PositiveTf(-1, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.PositiveTf should call FailNow()")
		}
	})
}

func TestAssertionsRegexpTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.RegexpTf("^start", "starting", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.RegexpTf("^start", "not starting", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.RegexpTf should call FailNow()")
		}
	})
}

func TestAssertionsSameTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SameTf(&staticVar, staticVarPtr, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SameTf(&staticVar, ptr("static string"), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SameTf should call FailNow()")
		}
	})
}

func TestAssertionsSeqContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqContainsTf(slices.Values([]string{"A", "B"}), "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqContainsTf(slices.Values([]string{"A", "B"}), "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SeqContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsSeqNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqNotContainsTf(slices.Values([]string{"A", "B"}), "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SeqNotContainsTf(slices.Values([]string{"A", "B"}), "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SeqNotContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsSliceContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceContainsTf([]string{"A", "B"}, "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceContainsTf([]string{"A", "B"}, "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsSliceEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceEqualTf([]string{"Hello", "World"}, []string{"Hello", "World"}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceEqualTf([]string{"Hello", "World"}, []string{"Hello"}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceEqualTf should call FailNow()")
		}
	})
}

func TestAssertionsSliceNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotContainsTf([]string{"A", "B"}, "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotContainsTf([]string{"A", "B"}, "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceNotContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsSliceNotEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotEqualTf([]string{"Hello", "World"}, []string{"Hello"}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotEqualTf([]string{"Hello", "World"}, []string{"Hello", "World"}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceNotEqualTf should call FailNow()")
		}
	})
}

func TestAssertionsSliceNotSubsetTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotSubsetTf([]int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceNotSubsetTf([]int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceNotSubsetTf should call FailNow()")
		}
	})
}

func TestAssertionsSliceSubsetTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceSubsetTf([]int{1, 2, 3}, []int{1, 2}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SliceSubsetTf([]int{1, 2, 3}, []int{4, 5}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SliceSubsetTf should call FailNow()")
		}
	})
}

func TestAssertionsSortedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SortedTf([]int{1, 1, 3}, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.SortedTf([]int{1, 4, 2}, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.SortedTf should call FailNow()")
		}
	})
}

func TestAssertionsStringContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringContainsTf("AB", "A", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringContainsTf("AB", "C", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.StringContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsStringNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringNotContainsTf("AB", "C", "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.StringNotContainsTf("AB", "A", "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.StringNotContainsTf should call FailNow()")
		}
	})
}

func TestAssertionsTrueTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.TrueTf(1 == 1, "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.TrueTf(1 == 0, "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("Assertions.TrueTf should call FailNow()")
		}
	})
}

func TestAssertionsYAMLEqTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLEqTf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLEqTf should panic as expected")
		}
	})
}

func TestAssertionsYAMLMarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLMarshalAsTf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLMarshalAsTf should panic as expected")
		}
	})
}

func TestAssertionsYAMLUnmarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		a := New(mock)
		a.Panics(func() {
			a.YAMLUnmarshalAsTf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if mock.failed {
			t.Error("Assertions.YAMLUnmarshalAsTf should panic as expected")
		}
	})
}
