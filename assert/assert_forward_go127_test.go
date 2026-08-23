// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.27

package assert

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

		mock := new(mockT)
		a := New(mock)
		result := a.BlockedT(make(chan struct{}))
		if !result {
			t.Error("Assertions.BlockedT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.BlockedT(sendChanMessage())
		if result {
			t.Error("Assertions.BlockedT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.BlockedT should mark test as failed")
		}
	})
}

func TestAssertionsConsistently(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Consistently(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond)
		if !result {
			t.Error("Assertions.Consistently should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Consistently(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("Assertions.Consistently should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Consistently should mark test as failed")
		}
	})
}

func TestAssertionsElementsMatchT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatchT([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if !result {
			t.Error("Assertions.ElementsMatchT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatchT([]int{1, 2, 3}, []int{1, 2, 4})
		if result {
			t.Error("Assertions.ElementsMatchT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ElementsMatchT should mark test as failed")
		}
	})
}

func TestAssertionsEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualT(123, 123)
		if !result {
			t.Error("Assertions.EqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualT(123, 456)
		if result {
			t.Error("Assertions.EqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualT should mark test as failed")
		}
	})
}

func TestAssertionsErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAsType(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if !result {
			t.Error("Assertions.ErrorAsType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAsType(ErrTest, new(*dummyError))
		if result {
			t.Error("Assertions.ErrorAsType should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorAsType should mark test as failed")
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

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EventuallyWith(func(c *CollectT) { c.Cancel() }, 100*time.Millisecond, 20*time.Millisecond)
		if result {
			t.Error("Assertions.EventuallyWith should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EventuallyWith should mark test as failed")
		}
	})
}

func TestAssertionsFalseT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FalseT(1 == 0)
		if !result {
			t.Error("Assertions.FalseT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FalseT(1 == 1)
		if result {
			t.Error("Assertions.FalseT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FalseT should mark test as failed")
		}
	})
}

func TestAssertionsGreaterOrEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqualT(2, 1)
		if !result {
			t.Error("Assertions.GreaterOrEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqualT(1, 2)
		if result {
			t.Error("Assertions.GreaterOrEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqualT should mark test as failed")
		}
	})
}

func TestAssertionsGreaterT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterT(2, 1)
		if !result {
			t.Error("Assertions.GreaterT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterT(1, 2)
		if result {
			t.Error("Assertions.GreaterT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.GreaterT should mark test as failed")
		}
	})
}

func TestAssertionsInDeltaT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaT(1.0, 1.01, 0.02)
		if !result {
			t.Error("Assertions.InDeltaT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaT(1.0, 1.1, 0.05)
		if result {
			t.Error("Assertions.InDeltaT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDeltaT should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilonSymmetricT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSymmetricT(100.0, 101.0, 0.02)
		if !result {
			t.Error("Assertions.InEpsilonSymmetricT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSymmetricT(100.0, 110.0, 0.05)
		if result {
			t.Error("Assertions.InEpsilonSymmetricT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilonSymmetricT should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilonT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonT(100.0, 101.0, 0.02)
		if !result {
			t.Error("Assertions.InEpsilonT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonT(100.0, 110.0, 0.05)
		if result {
			t.Error("Assertions.InEpsilonT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilonT should mark test as failed")
		}
	})
}

func TestAssertionsIsDecreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasingT([]int{3, 2, 1})
		if !result {
			t.Error("Assertions.IsDecreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasingT([]int{1, 2, 3})
		if result {
			t.Error("Assertions.IsDecreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsDecreasingT should mark test as failed")
		}
	})
}

func TestAssertionsIsIncreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasingT([]int{1, 2, 3})
		if !result {
			t.Error("Assertions.IsIncreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasingT([]int{1, 1, 2})
		if result {
			t.Error("Assertions.IsIncreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsIncreasingT should mark test as failed")
		}
	})
}

func TestAssertionsIsNonDecreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasingT([]int{1, 1, 2})
		if !result {
			t.Error("Assertions.IsNonDecreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasingT([]int{2, 1, 0})
		if result {
			t.Error("Assertions.IsNonDecreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasingT should mark test as failed")
		}
	})
}

func TestAssertionsIsNonIncreasingT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasingT([]int{2, 1, 1})
		if !result {
			t.Error("Assertions.IsNonIncreasingT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasingT([]int{1, 2, 3})
		if result {
			t.Error("Assertions.IsNonIncreasingT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasingT should mark test as failed")
		}
	})
}

func TestAssertionsIsNotOfTypeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotOfTypeT[myType](123.123)
		if !result {
			t.Error("Assertions.IsNotOfTypeT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotOfTypeT[myType](myType(123.123))
		if result {
			t.Error("Assertions.IsNotOfTypeT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNotOfTypeT should mark test as failed")
		}
	})
}

func TestAssertionsIsOfTypeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsOfTypeT[myType](myType(123.123))
		if !result {
			t.Error("Assertions.IsOfTypeT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsOfTypeT[myType](123.123)
		if result {
			t.Error("Assertions.IsOfTypeT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsOfTypeT should mark test as failed")
		}
	})
}

func TestAssertionsJSONEqT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqT(`{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`))
		if !result {
			t.Error("Assertions.JSONEqT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqT(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`)
		if result {
			t.Error("Assertions.JSONEqT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONEqT should mark test as failed")
		}
	})
}

func TestAssertionsJSONMarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONMarshalAsT([]byte(`{"A": "a"}`), dummyStruct{A: "a"})
		if !result {
			t.Error("Assertions.JSONMarshalAsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONMarshalAsT(`[{"foo": "bar"}, {"hello": "world"}]`, 1)
		if result {
			t.Error("Assertions.JSONMarshalAsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONMarshalAsT should mark test as failed")
		}
	})
}

func TestAssertionsJSONUnmarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONUnmarshalAsT(dummyStruct{A: "a"}, []byte(`{"A": "a"}`))
		if !result {
			t.Error("Assertions.JSONUnmarshalAsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONUnmarshalAsT(1, `[{"foo": "bar"}, {"hello": "world"}]`)
		if result {
			t.Error("Assertions.JSONUnmarshalAsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONUnmarshalAsT should mark test as failed")
		}
	})
}

func TestAssertionsLessOrEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqualT(1, 2)
		if !result {
			t.Error("Assertions.LessOrEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqualT(2, 1)
		if result {
			t.Error("Assertions.LessOrEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.LessOrEqualT should mark test as failed")
		}
	})
}

func TestAssertionsLessT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessT(1, 2)
		if !result {
			t.Error("Assertions.LessT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessT(2, 1)
		if result {
			t.Error("Assertions.LessT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.LessT should mark test as failed")
		}
	})
}

func TestAssertionsMapContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapContainsT(map[string]string{"A": "B"}, "A")
		if !result {
			t.Error("Assertions.MapContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapContainsT(map[string]string{"A": "B"}, "C")
		if result {
			t.Error("Assertions.MapContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapContainsT should mark test as failed")
		}
	})
}

func TestAssertionsMapEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapEqualT(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"})
		if !result {
			t.Error("Assertions.MapEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapEqualT(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"})
		if result {
			t.Error("Assertions.MapEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapEqualT should mark test as failed")
		}
	})
}

func TestAssertionsMapNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotContainsT(map[string]string{"A": "B"}, "C")
		if !result {
			t.Error("Assertions.MapNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotContainsT(map[string]string{"A": "B"}, "A")
		if result {
			t.Error("Assertions.MapNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapNotContainsT should mark test as failed")
		}
	})
}

func TestAssertionsMapNotEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotEqualT(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"})
		if !result {
			t.Error("Assertions.MapNotEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotEqualT(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"})
		if result {
			t.Error("Assertions.MapNotEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapNotEqualT should mark test as failed")
		}
	})
}

func TestAssertionsNegativeT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NegativeT(-1)
		if !result {
			t.Error("Assertions.NegativeT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NegativeT(1)
		if result {
			t.Error("Assertions.NegativeT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NegativeT should mark test as failed")
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

func TestAssertionsNotBlockedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotBlockedT(sendChanMessage())
		if !result {
			t.Error("Assertions.NotBlockedT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotBlockedT(make(chan struct{}))
		if result {
			t.Error("Assertions.NotBlockedT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotBlockedT should mark test as failed")
		}
	})
}

func TestAssertionsNotElementsMatchT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatchT([]int{1, 2, 3}, []int{1, 2, 4})
		if !result {
			t.Error("Assertions.NotElementsMatchT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatchT([]int{1, 3, 2, 3}, []int{1, 3, 3, 2})
		if result {
			t.Error("Assertions.NotElementsMatchT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotElementsMatchT should mark test as failed")
		}
	})
}

func TestAssertionsNotEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualT(123, 456)
		if !result {
			t.Error("Assertions.NotEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualT(123, 123)
		if result {
			t.Error("Assertions.NotEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEqualT should mark test as failed")
		}
	})
}

func TestAssertionsNotErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAsType(ErrTest, new(*dummyError))
		if !result {
			t.Error("Assertions.NotErrorAsType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAsType(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if result {
			t.Error("Assertions.NotErrorAsType should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotErrorAsType should mark test as failed")
		}
	})
}

func TestAssertionsNotRegexpT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexpT("^start", "not starting")
		if !result {
			t.Error("Assertions.NotRegexpT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexpT("^start", "starting")
		if result {
			t.Error("Assertions.NotRegexpT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotRegexpT should mark test as failed")
		}
	})
}

func TestAssertionsNotSameT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSameT(&staticVar, ptr("static string"))
		if !result {
			t.Error("Assertions.NotSameT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSameT(&staticVar, staticVarPtr)
		if result {
			t.Error("Assertions.NotSameT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSameT should mark test as failed")
		}
	})
}

func TestAssertionsNotSortedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSortedT([]int{3, 1, 3})
		if !result {
			t.Error("Assertions.NotSortedT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSortedT([]int{1, 4, 8})
		if result {
			t.Error("Assertions.NotSortedT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSortedT should mark test as failed")
		}
	})
}

func TestAssertionsPositiveT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PositiveT(1)
		if !result {
			t.Error("Assertions.PositiveT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PositiveT(-1)
		if result {
			t.Error("Assertions.PositiveT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.PositiveT should mark test as failed")
		}
	})
}

func TestAssertionsRegexpT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.RegexpT("^start", "starting")
		if !result {
			t.Error("Assertions.RegexpT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.RegexpT("^start", "not starting")
		if result {
			t.Error("Assertions.RegexpT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.RegexpT should mark test as failed")
		}
	})
}

func TestAssertionsSameT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SameT(&staticVar, staticVarPtr)
		if !result {
			t.Error("Assertions.SameT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SameT(&staticVar, ptr("static string"))
		if result {
			t.Error("Assertions.SameT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SameT should mark test as failed")
		}
	})
}

func TestAssertionsSeqContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqContainsT(slices.Values([]string{"A", "B"}), "A")
		if !result {
			t.Error("Assertions.SeqContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqContainsT(slices.Values([]string{"A", "B"}), "C")
		if result {
			t.Error("Assertions.SeqContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SeqContainsT should mark test as failed")
		}
	})
}

func TestAssertionsSeqNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqNotContainsT(slices.Values([]string{"A", "B"}), "C")
		if !result {
			t.Error("Assertions.SeqNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqNotContainsT(slices.Values([]string{"A", "B"}), "A")
		if result {
			t.Error("Assertions.SeqNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SeqNotContainsT should mark test as failed")
		}
	})
}

func TestAssertionsSliceContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceContainsT([]string{"A", "B"}, "A")
		if !result {
			t.Error("Assertions.SliceContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceContainsT([]string{"A", "B"}, "C")
		if result {
			t.Error("Assertions.SliceContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceContainsT should mark test as failed")
		}
	})
}

func TestAssertionsSliceEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceEqualT([]string{"Hello", "World"}, []string{"Hello", "World"})
		if !result {
			t.Error("Assertions.SliceEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceEqualT([]string{"Hello", "World"}, []string{"Hello"})
		if result {
			t.Error("Assertions.SliceEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceEqualT should mark test as failed")
		}
	})
}

func TestAssertionsSliceNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotContainsT([]string{"A", "B"}, "C")
		if !result {
			t.Error("Assertions.SliceNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotContainsT([]string{"A", "B"}, "A")
		if result {
			t.Error("Assertions.SliceNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceNotContainsT should mark test as failed")
		}
	})
}

func TestAssertionsSliceNotEqualT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotEqualT([]string{"Hello", "World"}, []string{"Hello"})
		if !result {
			t.Error("Assertions.SliceNotEqualT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotEqualT([]string{"Hello", "World"}, []string{"Hello", "World"})
		if result {
			t.Error("Assertions.SliceNotEqualT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceNotEqualT should mark test as failed")
		}
	})
}

func TestAssertionsSliceNotSubsetT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotSubsetT([]int{1, 2, 3}, []int{4, 5})
		if !result {
			t.Error("Assertions.SliceNotSubsetT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotSubsetT([]int{1, 2, 3}, []int{1, 2})
		if result {
			t.Error("Assertions.SliceNotSubsetT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceNotSubsetT should mark test as failed")
		}
	})
}

func TestAssertionsSliceSubsetT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceSubsetT([]int{1, 2, 3}, []int{1, 2})
		if !result {
			t.Error("Assertions.SliceSubsetT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceSubsetT([]int{1, 2, 3}, []int{4, 5})
		if result {
			t.Error("Assertions.SliceSubsetT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceSubsetT should mark test as failed")
		}
	})
}

func TestAssertionsSortedT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SortedT([]int{1, 1, 3})
		if !result {
			t.Error("Assertions.SortedT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SortedT([]int{1, 4, 2})
		if result {
			t.Error("Assertions.SortedT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SortedT should mark test as failed")
		}
	})
}

func TestAssertionsStringContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringContainsT("AB", "A")
		if !result {
			t.Error("Assertions.StringContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringContainsT("AB", "C")
		if result {
			t.Error("Assertions.StringContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.StringContainsT should mark test as failed")
		}
	})
}

func TestAssertionsStringNotContainsT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringNotContainsT("AB", "C")
		if !result {
			t.Error("Assertions.StringNotContainsT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringNotContainsT("AB", "A")
		if result {
			t.Error("Assertions.StringNotContainsT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.StringNotContainsT should mark test as failed")
		}
	})
}

func TestAssertionsTrueT(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.TrueT(1 == 1)
		if !result {
			t.Error("Assertions.TrueT should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.TrueT(1 == 0)
		if result {
			t.Error("Assertions.TrueT should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.TrueT should mark test as failed")
		}
	})
}

func TestAssertionsYAMLEqT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLEqT("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLEqT should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLEqT should panic as expected")
		}
	})
}

func TestAssertionsYAMLMarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLMarshalAsT("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLMarshalAsT should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLMarshalAsT should panic as expected")
		}
	})
}

func TestAssertionsYAMLUnmarshalAsT(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLUnmarshalAsT("key: value", "key: value")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLUnmarshalAsT should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLUnmarshalAsT should panic as expected")
		}
	})
}

func TestAssertionsBlockedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.BlockedTf(make(chan struct{}), "test message")
		if !result {
			t.Error("Assertions.BlockedTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.BlockedTf(sendChanMessage(), "test message")
		if result {
			t.Error("Assertions.BlockedTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.BlockedTf should mark test as failed")
		}
	})
}

func TestAssertionsConsistentlyf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Consistentlyf(func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if !result {
			t.Error("Assertions.Consistentlyf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Consistentlyf(func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("Assertions.Consistentlyf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.Consistentlyf should mark test as failed")
		}
	})
}

func TestAssertionsElementsMatchTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatchTf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if !result {
			t.Error("Assertions.ElementsMatchTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ElementsMatchTf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if result {
			t.Error("Assertions.ElementsMatchTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ElementsMatchTf should mark test as failed")
		}
	})
}

func TestAssertionsEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualTf(123, 123, "test message")
		if !result {
			t.Error("Assertions.EqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EqualTf(123, 456, "test message")
		if result {
			t.Error("Assertions.EqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EqualTf should mark test as failed")
		}
	})
}

func TestAssertionsErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAsTypef(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if !result {
			t.Error("Assertions.ErrorAsTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.ErrorAsTypef(ErrTest, new(*dummyError), "test message")
		if result {
			t.Error("Assertions.ErrorAsTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.ErrorAsTypef should mark test as failed")
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

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.EventuallyWithf(func(c *CollectT) { c.Cancel() }, 100*time.Millisecond, 20*time.Millisecond, "test message")
		if result {
			t.Error("Assertions.EventuallyWithf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.EventuallyWithf should mark test as failed")
		}
	})
}

func TestAssertionsFalseTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FalseTf(1 == 0, "test message")
		if !result {
			t.Error("Assertions.FalseTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.FalseTf(1 == 1, "test message")
		if result {
			t.Error("Assertions.FalseTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.FalseTf should mark test as failed")
		}
	})
}

func TestAssertionsGreaterOrEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqualTf(2, 1, "test message")
		if !result {
			t.Error("Assertions.GreaterOrEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterOrEqualTf(1, 2, "test message")
		if result {
			t.Error("Assertions.GreaterOrEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.GreaterOrEqualTf should mark test as failed")
		}
	})
}

func TestAssertionsGreaterTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterTf(2, 1, "test message")
		if !result {
			t.Error("Assertions.GreaterTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.GreaterTf(1, 2, "test message")
		if result {
			t.Error("Assertions.GreaterTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.GreaterTf should mark test as failed")
		}
	})
}

func TestAssertionsInDeltaTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaTf(1.0, 1.01, 0.02, "test message")
		if !result {
			t.Error("Assertions.InDeltaTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InDeltaTf(1.0, 1.1, 0.05, "test message")
		if result {
			t.Error("Assertions.InDeltaTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InDeltaTf should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilonSymmetricTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSymmetricTf(100.0, 101.0, 0.02, "test message")
		if !result {
			t.Error("Assertions.InEpsilonSymmetricTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonSymmetricTf(100.0, 110.0, 0.05, "test message")
		if result {
			t.Error("Assertions.InEpsilonSymmetricTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilonSymmetricTf should mark test as failed")
		}
	})
}

func TestAssertionsInEpsilonTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonTf(100.0, 101.0, 0.02, "test message")
		if !result {
			t.Error("Assertions.InEpsilonTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.InEpsilonTf(100.0, 110.0, 0.05, "test message")
		if result {
			t.Error("Assertions.InEpsilonTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.InEpsilonTf should mark test as failed")
		}
	})
}

func TestAssertionsIsDecreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasingTf([]int{3, 2, 1}, "test message")
		if !result {
			t.Error("Assertions.IsDecreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsDecreasingTf([]int{1, 2, 3}, "test message")
		if result {
			t.Error("Assertions.IsDecreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsDecreasingTf should mark test as failed")
		}
	})
}

func TestAssertionsIsIncreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasingTf([]int{1, 2, 3}, "test message")
		if !result {
			t.Error("Assertions.IsIncreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsIncreasingTf([]int{1, 1, 2}, "test message")
		if result {
			t.Error("Assertions.IsIncreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsIncreasingTf should mark test as failed")
		}
	})
}

func TestAssertionsIsNonDecreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasingTf([]int{1, 1, 2}, "test message")
		if !result {
			t.Error("Assertions.IsNonDecreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonDecreasingTf([]int{2, 1, 0}, "test message")
		if result {
			t.Error("Assertions.IsNonDecreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonDecreasingTf should mark test as failed")
		}
	})
}

func TestAssertionsIsNonIncreasingTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasingTf([]int{2, 1, 1}, "test message")
		if !result {
			t.Error("Assertions.IsNonIncreasingTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNonIncreasingTf([]int{1, 2, 3}, "test message")
		if result {
			t.Error("Assertions.IsNonIncreasingTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNonIncreasingTf should mark test as failed")
		}
	})
}

func TestAssertionsIsNotOfTypeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotOfTypeTf[myType](123.123, "test message")
		if !result {
			t.Error("Assertions.IsNotOfTypeTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsNotOfTypeTf[myType](myType(123.123), "test message")
		if result {
			t.Error("Assertions.IsNotOfTypeTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsNotOfTypeTf should mark test as failed")
		}
	})
}

func TestAssertionsIsOfTypeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsOfTypeTf[myType](myType(123.123), "test message")
		if !result {
			t.Error("Assertions.IsOfTypeTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.IsOfTypeTf[myType](123.123, "test message")
		if result {
			t.Error("Assertions.IsOfTypeTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.IsOfTypeTf should mark test as failed")
		}
	})
}

func TestAssertionsJSONEqTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqTf(`{"hello": "world", "foo": "bar"}`, []byte(`{"foo": "bar", "hello": "world"}`), "test message")
		if !result {
			t.Error("Assertions.JSONEqTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONEqTf(`{"hello": "world", "foo": "bar"}`, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		if result {
			t.Error("Assertions.JSONEqTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONEqTf should mark test as failed")
		}
	})
}

func TestAssertionsJSONMarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONMarshalAsTf([]byte(`{"A": "a"}`), dummyStruct{A: "a"}, "test message")
		if !result {
			t.Error("Assertions.JSONMarshalAsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONMarshalAsTf(`[{"foo": "bar"}, {"hello": "world"}]`, 1, "test message")
		if result {
			t.Error("Assertions.JSONMarshalAsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONMarshalAsTf should mark test as failed")
		}
	})
}

func TestAssertionsJSONUnmarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONUnmarshalAsTf(dummyStruct{A: "a"}, []byte(`{"A": "a"}`), "test message")
		if !result {
			t.Error("Assertions.JSONUnmarshalAsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.JSONUnmarshalAsTf(1, `[{"foo": "bar"}, {"hello": "world"}]`, "test message")
		if result {
			t.Error("Assertions.JSONUnmarshalAsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.JSONUnmarshalAsTf should mark test as failed")
		}
	})
}

func TestAssertionsLessOrEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqualTf(1, 2, "test message")
		if !result {
			t.Error("Assertions.LessOrEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessOrEqualTf(2, 1, "test message")
		if result {
			t.Error("Assertions.LessOrEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.LessOrEqualTf should mark test as failed")
		}
	})
}

func TestAssertionsLessTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessTf(1, 2, "test message")
		if !result {
			t.Error("Assertions.LessTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.LessTf(2, 1, "test message")
		if result {
			t.Error("Assertions.LessTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.LessTf should mark test as failed")
		}
	})
}

func TestAssertionsMapContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapContainsTf(map[string]string{"A": "B"}, "A", "test message")
		if !result {
			t.Error("Assertions.MapContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapContainsTf(map[string]string{"A": "B"}, "C", "test message")
		if result {
			t.Error("Assertions.MapContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsMapEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapEqualTf(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		if !result {
			t.Error("Assertions.MapEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapEqualTf(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		if result {
			t.Error("Assertions.MapEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapEqualTf should mark test as failed")
		}
	})
}

func TestAssertionsMapNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotContainsTf(map[string]string{"A": "B"}, "C", "test message")
		if !result {
			t.Error("Assertions.MapNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotContainsTf(map[string]string{"A": "B"}, "A", "test message")
		if result {
			t.Error("Assertions.MapNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapNotContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsMapNotEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotEqualTf(map[string]string{"2": "Hello", "1": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		if !result {
			t.Error("Assertions.MapNotEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.MapNotEqualTf(map[string]string{"1": "Hello", "2": "World"}, map[string]string{"1": "Hello", "2": "World"}, "test message")
		if result {
			t.Error("Assertions.MapNotEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.MapNotEqualTf should mark test as failed")
		}
	})
}

func TestAssertionsNegativeTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NegativeTf(-1, "test message")
		if !result {
			t.Error("Assertions.NegativeTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NegativeTf(1, "test message")
		if result {
			t.Error("Assertions.NegativeTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NegativeTf should mark test as failed")
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

func TestAssertionsNotBlockedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotBlockedTf(sendChanMessage(), "test message")
		if !result {
			t.Error("Assertions.NotBlockedTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotBlockedTf(make(chan struct{}), "test message")
		if result {
			t.Error("Assertions.NotBlockedTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotBlockedTf should mark test as failed")
		}
	})
}

func TestAssertionsNotElementsMatchTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatchTf([]int{1, 2, 3}, []int{1, 2, 4}, "test message")
		if !result {
			t.Error("Assertions.NotElementsMatchTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotElementsMatchTf([]int{1, 3, 2, 3}, []int{1, 3, 3, 2}, "test message")
		if result {
			t.Error("Assertions.NotElementsMatchTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotElementsMatchTf should mark test as failed")
		}
	})
}

func TestAssertionsNotEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualTf(123, 456, "test message")
		if !result {
			t.Error("Assertions.NotEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotEqualTf(123, 123, "test message")
		if result {
			t.Error("Assertions.NotEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotEqualTf should mark test as failed")
		}
	})
}

func TestAssertionsNotErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAsTypef(ErrTest, new(*dummyError), "test message")
		if !result {
			t.Error("Assertions.NotErrorAsTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotErrorAsTypef(fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if result {
			t.Error("Assertions.NotErrorAsTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotErrorAsTypef should mark test as failed")
		}
	})
}

func TestAssertionsNotRegexpTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexpTf("^start", "not starting", "test message")
		if !result {
			t.Error("Assertions.NotRegexpTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotRegexpTf("^start", "starting", "test message")
		if result {
			t.Error("Assertions.NotRegexpTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotRegexpTf should mark test as failed")
		}
	})
}

func TestAssertionsNotSameTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSameTf(&staticVar, ptr("static string"), "test message")
		if !result {
			t.Error("Assertions.NotSameTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSameTf(&staticVar, staticVarPtr, "test message")
		if result {
			t.Error("Assertions.NotSameTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSameTf should mark test as failed")
		}
	})
}

func TestAssertionsNotSortedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSortedTf([]int{3, 1, 3}, "test message")
		if !result {
			t.Error("Assertions.NotSortedTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.NotSortedTf([]int{1, 4, 8}, "test message")
		if result {
			t.Error("Assertions.NotSortedTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.NotSortedTf should mark test as failed")
		}
	})
}

func TestAssertionsPositiveTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PositiveTf(1, "test message")
		if !result {
			t.Error("Assertions.PositiveTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.PositiveTf(-1, "test message")
		if result {
			t.Error("Assertions.PositiveTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.PositiveTf should mark test as failed")
		}
	})
}

func TestAssertionsRegexpTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.RegexpTf("^start", "starting", "test message")
		if !result {
			t.Error("Assertions.RegexpTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.RegexpTf("^start", "not starting", "test message")
		if result {
			t.Error("Assertions.RegexpTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.RegexpTf should mark test as failed")
		}
	})
}

func TestAssertionsSameTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SameTf(&staticVar, staticVarPtr, "test message")
		if !result {
			t.Error("Assertions.SameTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SameTf(&staticVar, ptr("static string"), "test message")
		if result {
			t.Error("Assertions.SameTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SameTf should mark test as failed")
		}
	})
}

func TestAssertionsSeqContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqContainsTf(slices.Values([]string{"A", "B"}), "A", "test message")
		if !result {
			t.Error("Assertions.SeqContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqContainsTf(slices.Values([]string{"A", "B"}), "C", "test message")
		if result {
			t.Error("Assertions.SeqContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SeqContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsSeqNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqNotContainsTf(slices.Values([]string{"A", "B"}), "C", "test message")
		if !result {
			t.Error("Assertions.SeqNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SeqNotContainsTf(slices.Values([]string{"A", "B"}), "A", "test message")
		if result {
			t.Error("Assertions.SeqNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SeqNotContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsSliceContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceContainsTf([]string{"A", "B"}, "A", "test message")
		if !result {
			t.Error("Assertions.SliceContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceContainsTf([]string{"A", "B"}, "C", "test message")
		if result {
			t.Error("Assertions.SliceContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsSliceEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceEqualTf([]string{"Hello", "World"}, []string{"Hello", "World"}, "test message")
		if !result {
			t.Error("Assertions.SliceEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceEqualTf([]string{"Hello", "World"}, []string{"Hello"}, "test message")
		if result {
			t.Error("Assertions.SliceEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceEqualTf should mark test as failed")
		}
	})
}

func TestAssertionsSliceNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotContainsTf([]string{"A", "B"}, "C", "test message")
		if !result {
			t.Error("Assertions.SliceNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotContainsTf([]string{"A", "B"}, "A", "test message")
		if result {
			t.Error("Assertions.SliceNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceNotContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsSliceNotEqualTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotEqualTf([]string{"Hello", "World"}, []string{"Hello"}, "test message")
		if !result {
			t.Error("Assertions.SliceNotEqualTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotEqualTf([]string{"Hello", "World"}, []string{"Hello", "World"}, "test message")
		if result {
			t.Error("Assertions.SliceNotEqualTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceNotEqualTf should mark test as failed")
		}
	})
}

func TestAssertionsSliceNotSubsetTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotSubsetTf([]int{1, 2, 3}, []int{4, 5}, "test message")
		if !result {
			t.Error("Assertions.SliceNotSubsetTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceNotSubsetTf([]int{1, 2, 3}, []int{1, 2}, "test message")
		if result {
			t.Error("Assertions.SliceNotSubsetTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceNotSubsetTf should mark test as failed")
		}
	})
}

func TestAssertionsSliceSubsetTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceSubsetTf([]int{1, 2, 3}, []int{1, 2}, "test message")
		if !result {
			t.Error("Assertions.SliceSubsetTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SliceSubsetTf([]int{1, 2, 3}, []int{4, 5}, "test message")
		if result {
			t.Error("Assertions.SliceSubsetTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SliceSubsetTf should mark test as failed")
		}
	})
}

func TestAssertionsSortedTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SortedTf([]int{1, 1, 3}, "test message")
		if !result {
			t.Error("Assertions.SortedTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.SortedTf([]int{1, 4, 2}, "test message")
		if result {
			t.Error("Assertions.SortedTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.SortedTf should mark test as failed")
		}
	})
}

func TestAssertionsStringContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringContainsTf("AB", "A", "test message")
		if !result {
			t.Error("Assertions.StringContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringContainsTf("AB", "C", "test message")
		if result {
			t.Error("Assertions.StringContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.StringContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsStringNotContainsTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringNotContainsTf("AB", "C", "test message")
		if !result {
			t.Error("Assertions.StringNotContainsTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.StringNotContainsTf("AB", "A", "test message")
		if result {
			t.Error("Assertions.StringNotContainsTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.StringNotContainsTf should mark test as failed")
		}
	})
}

func TestAssertionsTrueTf(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.TrueTf(1 == 1, "test message")
		if !result {
			t.Error("Assertions.TrueTf should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.TrueTf(1 == 0, "test message")
		if result {
			t.Error("Assertions.TrueTf should return false on failure")
		}
		if !mock.failed {
			t.Error("Assertions.TrueTf should mark test as failed")
		}
	})
}

func TestAssertionsYAMLEqTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLEqTf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLEqTf should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLEqTf should panic as expected")
		}
	})
}

func TestAssertionsYAMLMarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLMarshalAsTf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLMarshalAsTf should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLMarshalAsTf should panic as expected")
		}
	})
}

func TestAssertionsYAMLUnmarshalAsTf(t *testing.T) {
	t.Parallel()

	t.Run("panic", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		a := New(mock)
		result := a.Panics(func() {
			a.YAMLUnmarshalAsTf("key: value", "key: value", "test message")
		}, "should panic without the yaml feature enabled.")
		if !result {
			t.Error("Assertions.YAMLUnmarshalAsTf should return true on panic")
		}
		if mock.failed {
			t.Error("Assertions.YAMLUnmarshalAsTf should panic as expected")
		}
	})
}
