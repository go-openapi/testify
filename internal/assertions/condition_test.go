// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"testing"
	"time"
)

func TestCondition(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	if !Condition(mock, func() bool { return true }, "Truth") {
		t.Error("Condition should return true")
	}

	if Condition(mock, func() bool { return false }, "Lie") {
		t.Error("Condition should return false")
	}
}

func TestConditionEventually(t *testing.T) {
	t.Parallel()

	t.Run("condition should Eventually be false", func(t *testing.T) {
		t.Parallel()
		mock := new(testing.T)

		condition := func() bool {
			return false
		}

		False(t, Eventually(mock, condition, 100*time.Millisecond, 20*time.Millisecond))
	})

	t.Run("condition should Eventually be true", func(t *testing.T) {
		t.Parallel()
		mock := new(testing.T)

		state := 0
		condition := func() bool {
			defer func() {
				state++
			}()
			return state == 2
		}

		True(t, Eventually(mock, condition, 100*time.Millisecond, 20*time.Millisecond))
	})
}

func TestConditionEventuallyWithTFalse(t *testing.T) {
	t.Parallel()
	mock := new(errorsCapturingT)

	condition := func(collect *CollectT) {
		Fail(collect, "condition fixed failure")
	}

	False(t, EventuallyWithT(mock, condition, 100*time.Millisecond, 20*time.Millisecond))
	Len(t, mock.errors, 2)
}

func TestConditionEventuallyWithTTrue(t *testing.T) {
	t.Parallel()
	mock := new(errorsCapturingT)

	counter := 0
	condition := func(collect *CollectT) {
		counter++
		True(collect, counter == 2)
	}

	True(t, EventuallyWithT(mock, condition, 100*time.Millisecond, 20*time.Millisecond))
	Len(t, mock.errors, 0)
	Equal(t, 2, counter, "Condition is expected to be called 2 times")
}

func TestConditionEventuallyWithT_ConcurrencySafe(t *testing.T) {
	t.Parallel()
	mock := new(errorsCapturingT)

	condition := func(collect *CollectT) {
		Fail(collect, "condition fixed failure")
	}

	// To trigger race conditions, we run EventuallyWithT with a nanosecond tick.
	False(t, EventuallyWithT(mock, condition, 100*time.Millisecond, time.Nanosecond))
	Len(t, mock.errors, 2)
}

func TestConditionEventuallyWithT_ReturnsTheLatestFinishedConditionErrors(t *testing.T) {
	t.Parallel()
	mock := new(errorsCapturingT)

	// We'll use a channel to control whether a condition should sleep or not.
	mustSleep := make(chan bool, 2)
	mustSleep <- false
	mustSleep <- true
	close(mustSleep)

	condition := func(collect *CollectT) {
		if <-mustSleep {
			// Sleep to ensure that the second condition runs longer than timeout.
			time.Sleep(time.Second)
			return
		}

		// The first condition will fail. We expect to get this error as a result.
		Fail(collect, "condition fixed failure")
	}

	False(t, EventuallyWithT(mock, condition, 100*time.Millisecond, 20*time.Millisecond))
	Len(t, mock.errors, 2)
}

func TestConditionEventuallyWithTFailNow(t *testing.T) {
	t.Parallel()
	mock := new(CollectT)

	condition := func(collect *CollectT) {
		collect.FailNow()
	}

	False(t, EventuallyWithT(mock, condition, 100*time.Millisecond, 20*time.Millisecond))
	Len(t, mock.errors, 1)
}

// Check that a long running condition doesn't block Eventually.
// See issue 805 (and its long tail of following issues).
func TestConditionEventuallyTimeout(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	NotPanics(t, func() {
		done, done2 := make(chan struct{}), make(chan struct{})

		// A condition function that returns after the Eventually timeout
		condition := func() bool {
			// Wait until Eventually times out and terminates
			<-done
			close(done2)
			return true
		}

		False(t, Eventually(mock, condition, time.Millisecond, time.Microsecond))

		close(done)
		<-done2
	})
}

func TestConditionEventuallySucceedQuickly(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	condition := func() bool { return true }

	// By making the tick longer than the total duration, we expect that this test would fail if
	// we didn't check the condition before the first tick elapses.
	True(t, Eventually(mock, condition, 100*time.Millisecond, time.Second))
}

func TestConditionEventuallyWithTSucceedQuickly(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	condition := func(*CollectT) {}

	// By making the tick longer than the total duration, we expect that this test would fail if
	// we didn't check the condition before the first tick elapses.
	True(t, EventuallyWithT(mock, condition, 100*time.Millisecond, time.Second))
}

func TestConditionNeverFalse(t *testing.T) {
	t.Parallel()

	condition := func() bool {
		return false
	}

	True(t, Never(t, condition, 100*time.Millisecond, 20*time.Millisecond))
}

// TestNeverTrue checks Never with a condition that returns true on second call.
func TestConditionNeverTrue(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	// A list of values returned by condition.
	// Channel protects against concurrent access.
	returns := make(chan bool, 2)
	returns <- false
	returns <- true
	defer close(returns)

	// Will return true on second call.
	condition := func() bool {
		return <-returns
	}

	False(t, Never(mock, condition, 100*time.Millisecond, 20*time.Millisecond))
}

func TestConditionNeverFailQuickly(t *testing.T) {
	t.Parallel()
	mock := new(testing.T)

	// By making the tick longer than the total duration, we expect that this test would fail if
	// we didn't check the condition before the first tick elapses.
	condition := func() bool { return true }
	False(t, Never(mock, condition, 100*time.Millisecond, time.Second))
}

// errorsCapturingT is a mock implementation of TestingT that captures errors reported with Errorf.
type errorsCapturingT struct {
	errors []error
}

// Helper is like [testing.T.Helper] but does nothing.
func (errorsCapturingT) Helper() {}

func (t *errorsCapturingT) Errorf(format string, args ...any) {
	t.errors = append(t.errors, fmt.Errorf(format, args...))
}
