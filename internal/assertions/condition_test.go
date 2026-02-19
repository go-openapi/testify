// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
	"errors"
	"iter"
	"slices"
	"sort"
	"strings"
	"sync"
	"testing"
	"time"
)

const (
	testTimeout = 100 * time.Millisecond
	testTick    = 20 * time.Millisecond
)

func TestCondition(t *testing.T) {
	t.Parallel()

	t.Run("condition should be true", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if !Condition(mock, func() bool { return true }, "Truth") {
			t.Error("condition should return true")
		}
	})

	t.Run("condition should be false", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		if Condition(mock, func() bool { return false }, "Lie") {
			t.Error("condition should return false")
		}
	})
}

func TestConditionEventually(t *testing.T) {
	t.Parallel()

	t.Run("condition should Eventually be false", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func() bool {
			return false
		}

		if Eventually(mock, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return false")
		}
	})

	t.Run("condition should Eventually be true", func(t *testing.T) {
		t.Parallel()

		state := 0
		condition := func() bool {
			defer func() {
				state++
			}()
			return state == 2
		}

		if !Eventually(t, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return true")
		}
	})
}

func TestConditionEventuallyWithError(t *testing.T) {
	t.Parallel()

	t.Run("condition should eventually return no error", func(t *testing.T) {
		t.Parallel()

		state := 0
		condition := func(_ context.Context) error {
			defer func() { state++ }()
			if state < 2 {
				return errors.New("not ready yet")
			}

			return nil
		}

		if !Eventually(t, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return true")
		}
	})

	t.Run("condition should eventually fail on persistent error", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(_ context.Context) error {
			return errors.New("persistent error")
		}

		if Eventually(mock, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return false")
		}
	})

	t.Run("condition should use provided context", func(t *testing.T) {
		t.Parallel()

		condition := func(ctx context.Context) error {
			if ctx == nil {
				return errors.New("expected non-nil context")
			}

			return nil
		}

		if !Eventually(t, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return true")
		}
	})
}

// Check that a long running condition doesn't block Eventually.
//
// See issue 805 (and its long tail of following issues).
func TestConditionEventuallyTimeout(t *testing.T) {
	t.Parallel()

	t.Run("should fail on timeout", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		// A condition function that returns after the Eventually timeout
		condition := func() bool {
			time.Sleep(100 * time.Millisecond)
			return true
		}

		if Eventually(mock, condition, time.Millisecond, time.Microsecond) {
			t.Error("expected Eventually to return false on timeout")
		}
	})

	t.Run("should fail on parent test failed", func(t *testing.T) {
		t.Parallel()

		parentCtx, failParent := context.WithCancel(context.WithoutCancel(t.Context()))
		mock := new(errorsCapturingT).WithContext(parentCtx)

		condition := func() bool {
			time.Sleep(testTick)
			failParent() // this cancels the parent context (e.g. mocks failing the parent test)
			time.Sleep(2 * testTick)

			return true
		}

		if Eventually(mock, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return false when parent test fails")
		}

		t.Run("reported errors should include the context cancellation", func(t *testing.T) {
			// assert how this failure is reported
			if len(mock.errors) != 2 {
				t.Errorf("expected 2 error messages (1 for context canceled, 1 for never met condition), got %d", len(mock.errors))
			}

			var hasContextCancelled, hasFailedCondition bool
			for _, err := range mock.errors {
				msg := err.Error()
				switch {
				case strings.Contains(msg, "context canceled"):
					hasContextCancelled = true
				case strings.Contains(msg, "never satisfied"):
					hasFailedCondition = true
				}
			}
			if !hasContextCancelled {
				t.Error("expected a context cancelled error")
			}
			if !hasFailedCondition {
				t.Error("expected a condition never satisfied error")
			}
		})
	})
}

func TestConditionEventuallySucceedQuickly(t *testing.T) {
	t.Parallel()

	t.Run("should succeed before the first tick", func(t *testing.T) {
		mock := new(errorsCapturingT)
		condition := func() bool { return true }

		// By making the tick longer than the total duration, we expect that this test would fail if
		// we didn't check the condition before the first tick elapses.
		if !Eventually(mock, condition, testTimeout, 1*time.Second) {
			t.Error("expected Eventually to return true before first tick")
		}
	})
}

func TestConditionEventuallyNoLeak(t *testing.T) {
	t.Parallel()

	t.Run("should output messages in a determined order", func(t *testing.T) {
		t.Parallel()

		/* Original output (replaced by integers) from https://github.com/stretchr/testify/issues/1611
		   condition_test.go:150: 2026-01-11 12:11:49.34854116 +0100 CET m=+0.000641595 Condition: inEventually = true
		   condition_test.go:152: 2026-01-11 12:11:49.84944055 +0100 CET m=+0.501540975 Condition: inEventually = true
		   condition_test.go:147: 2026-01-11 12:11:49.849484723 +0100 CET m=+0.501585149 Condition: end.
		   condition_test.go:160: 2026-01-11 12:11:49.849500022 +0100 CET m=+0.501600447 Eventually done
		   condition_test.go:163: 2026-01-11 12:11:49.849508218 +0100 CET m=+0.501608643 End of TestConditionEventuallyNoLeak/should_output_messages_in_a_determined_order
		*/
		mock := new(errorsCapturingT)
		done := make(chan struct{}, 1)
		recordedActions := make([]int, 0, 5)
		var mx sync.Mutex
		record := func(action int) {
			mx.Lock()
			defer mx.Unlock()

			recordedActions = append(recordedActions, action)
		}

		inEventually := true
		Eventually(mock,
			func() bool {
				defer func() {
					record(2)
					done <- struct{}{}
				}()
				if inEventually {
					record(0)
				}
				time.Sleep(5 * testTimeout)
				if inEventually {
					record(1)
				}
				return true
			},
			testTimeout,
			testTick,
		)

		inEventually = false
		record(3)

		<-done
		record(4)
		record(5)

		const expectedActions = 6
		if len(recordedActions) != expectedActions {
			t.Errorf("expected %d actions to be recorded, got %d", expectedActions, len(recordedActions))
		}
		if !sort.IntsAreSorted(recordedActions) {
			t.Errorf("expected recorded actions to be ordered, got %v", recordedActions)
		}
	})

	t.Run("should not leak a go routine for condition execution", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		done := make(chan bool, 1)

		inEventually := true
		Eventually(mock,
			func() bool {
				defer func() {
					done <- inEventually
				}()
				time.Sleep(5 * testTimeout)

				return true
			},
			testTimeout,
			testTick,
		)

		inEventually = false
		result := <-done
		if !result {
			t.Error("Condition should end while Eventually still runs.")
		}
	})
}

func TestConditionEventuallyWith(t *testing.T) {
	t.Parallel()

	t.Run("should complete with false", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		counter := 0
		condition := func(collect *CollectT) {
			counter++
			Fail(collect, "condition fixed failure")
			Fail(collect, "another condition fixed failure")
		}

		if EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return false")
		}

		const expectedErrors = 4
		if len(mock.errors) < expectedErrors-1 || len(mock.errors) > expectedErrors { // it may be 3 or 4, depending on how the test schedules
			t.Errorf("expected %d errors (2 from condition, 2 from Eventually), got %d", expectedErrors, len(mock.errors))
		}

		expectedCalls := int(testTimeout / testTick)
		if counter < expectedCalls-1 || counter > expectedCalls+1 { // it may be 4, 5 or 6 depending on how the test schedules
			t.Errorf("expected %d calls to the condition, but got %d", expectedCalls, counter)
		}
	})

	t.Run("should complete with true", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		counter := 0
		condition := func(collect *CollectT) {
			counter++
			True(collect, counter == 2)
		}

		if !EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return true")
		}
		if len(mock.errors) != 0 {
			t.Errorf("expected 0 errors, got %d", len(mock.errors))
		}
		const expectedCalls = 2
		if expectedCalls != counter {
			t.Errorf("expected condition to be called %d times, got %d", expectedCalls, counter)
		}
	})

	t.Run("should complete with fail, on a nanosecond tick", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(collect *CollectT) {
			Fail(collect, "condition fixed failure")
		}

		// To trigger race conditions, we run EventuallyWith with a nanosecond tick.
		if EventuallyWith(mock, condition, testTimeout, time.Nanosecond) {
			t.Error("expected EventuallyWith to return false")
		}
		const expectedErrors = 3
		if len(mock.errors) != expectedErrors {
			t.Errorf("expected %d errors (1 from condition, 2 from Eventually), got %d", expectedErrors, len(mock.errors))
		}
	})

	t.Run("should complete with fail, with latest failed condition", func(t *testing.T) {
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

		if EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return false")
		}
		const expectedErrors = 3
		if len(mock.errors) != expectedErrors {
			t.Errorf("expected %d errors (1 from condition, 2 from Eventually), got %d", expectedErrors, len(mock.errors))
		}
	})

	t.Run("should complete with success, with the ticker never used", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(*CollectT) {}

		// By making the tick longer than the total duration, we expect that this test would fail if
		// we didn't check the condition before the first tick elapses.
		if !EventuallyWith(mock, condition, testTimeout, time.Second) {
			t.Error("expected EventuallyWith to return true")
		}
	})

	t.Run("should fail with a call to collect.FailNow", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		counter := 0

		// The call to FailNow cancels the execution context of EventuallyWith.
		// so we don't have to wait for the timeout.
		condition := func(collect *CollectT) {
			counter++
			collect.FailNow()
		}

		if EventuallyWith(mock, condition, 30*time.Minute, testTick) {
			t.Error("expected EventuallyWith to return false")
		}
		const expectedErrors = 2
		if len(mock.errors) != expectedErrors {
			t.Errorf("expected %d errors (0 accumulated + 2 from EventuallyWith), got %d", expectedErrors, len(mock.errors))
		}
		if counter != 1 {
			t.Errorf("expected the condition function to have been called only once, but got: %d", counter)
		}
	})
}

func TestConditionPollUntilTimeout(t *testing.T) {
	t.Parallel()

	for c := range pollUntilTimeoutCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			badValue := !c.goodValue

			t.Run("should succeed with constant good value", func(t *testing.T) {
				t.Parallel()

				mock := new(errorsCapturingT)
				if !c.assertion(mock, func() bool { return c.goodValue }, testTimeout, testTick) {
					t.Errorf("expected %s to return true", c.name)
				}
			})

			t.Run("should succeed on timeout with slow bad value", func(t *testing.T) {
				t.Parallel()

				mock := new(errorsCapturingT)
				condition := func() bool {
					time.Sleep(2 * testTick)
					return badValue // returns bad value, but only after timeout
				}

				if !c.assertion(mock, condition, testTick, 1*time.Millisecond) {
					t.Errorf("expected %s to return true on timeout", c.name)
				}
			})

			t.Run("should fail when condition flips on second call", func(t *testing.T) {
				t.Parallel()

				mock := new(errorsCapturingT)
				returns := make(chan bool, 2)
				returns <- c.goodValue
				returns <- badValue
				defer close(returns)

				condition := func() bool { return <-returns }

				if c.assertion(mock, condition, testTimeout, testTick) {
					t.Errorf("expected %s to return false", c.name)
				}
			})

			t.Run("should fail before first tick with constant bad value", func(t *testing.T) {
				t.Parallel()

				mock := new(errorsCapturingT)
				// By making the tick longer than the total duration, we expect that this test would fail if
				// we didn't check the condition before the first tick elapses.
				if c.assertion(mock, func() bool { return badValue }, testTimeout, time.Second) {
					t.Errorf("expected %s to return false", c.name)
				}
			})

			t.Run("should fail when parent test fails", func(t *testing.T) {
				t.Parallel()

				parentCtx, failParent := context.WithCancel(context.WithoutCancel(t.Context()))
				mock := new(errorsCapturingT).WithContext(parentCtx)
				condition := func() bool {
					failParent() // cancels the parent context
					return c.goodValue
				}
				if c.assertion(mock, condition, testTimeout, time.Second) {
					t.Errorf("expected %s to return false when parent test fails", c.name)
				}
			})
		})
	}
}

func TestConditionConsistentlyWithError(t *testing.T) {
	t.Parallel()

	t.Run("should succeed when condition always returns nil", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(_ context.Context) error {
			return nil // no error = condition not triggered
		}

		if !Consistently(mock, condition, testTimeout, testTick) {
			t.Error("expected Consistently to return true when condition never returns an error")
		}
	})

	t.Run("should fail when condition returns an error", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(_ context.Context) error {
			return errors.New("something went wrong")
		}

		if Consistently(mock, condition, testTimeout, testTick) {
			t.Error("expected Consistently to return false when condition returns an error")
		}
	})

	t.Run("should fail when error is returned on second call", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		returns := make(chan error, 2)
		returns <- nil
		returns <- errors.New("something went wrong")
		defer close(returns)

		condition := func(_ context.Context) error {
			return <-returns
		}

		if Consistently(mock, condition, testTimeout, testTick) {
			t.Error("expected Consistently to return false")
		}
	})
}

func TestConditionEventuallyWithContext(t *testing.T) {
	t.Parallel()

	t.Run("should complete with true using context variant", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		counter := 0
		condition := func(_ context.Context, collect *CollectT) {
			counter++
			True(collect, counter == 2)
		}

		if !EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return true")
		}
		if len(mock.errors) != 0 {
			t.Errorf("expected 0 errors, got %d", len(mock.errors))
		}
		const expectedCalls = 2
		if expectedCalls != counter {
			t.Errorf("expected condition to be called %d times, got %d", expectedCalls, counter)
		}
	})

	t.Run("should complete with false using context variant", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(_ context.Context, collect *CollectT) {
			Fail(collect, "condition fixed failure")
		}

		if EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return false")
		}
	})

	t.Run("should receive a non-nil context", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(ctx context.Context, collect *CollectT) {
			if ctx == nil {
				Fail(collect, "expected non-nil context")
			}
		}

		if !EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return true")
		}
	})
}

func TestConditionErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, conditionFailCases())
}

func conditionFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:      "Condition/false",
			assertion: func(t T) bool { return Condition(t, func() bool { return false }) },
			wantError: "condition failed",
		},
	})
}

// pollUntilTimeoutAssertion is the common signature for Never and Consistently,
// both of which poll until timeout using func() bool conditions.
type pollUntilTimeoutAssertion func(T, func() bool, time.Duration, time.Duration, ...any) bool

// pollUntilTimeoutCase parameterizes the shared tests for Never and Consistently.
type pollUntilTimeoutCase struct {
	name      string
	assertion pollUntilTimeoutAssertion
	goodValue bool // the value the condition returns when "holding": false for Never, true for Consistently
}

func pollUntilTimeoutCases() iter.Seq[pollUntilTimeoutCase] {
	return slices.Values([]pollUntilTimeoutCase{
		{
			name:      "Never",
			assertion: Never,
			goodValue: false, // Never succeeds when the condition always returns false ("never true")
		},
		{
			name:      "Consistently",
			assertion: Consistently[func() bool],
			goodValue: true, // Consistently succeeds when the condition always returns true ("always true")
		},
	})
}
