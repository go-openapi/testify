// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
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

		mock := new(testing.T)
		if !Condition(mock, func() bool { return true }, "Truth") {
			t.Error("condition should return true")
		}
	})

	t.Run("condition should be false", func(t *testing.T) {
		t.Parallel()

		mock := new(testing.T)
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

		False(t, Eventually(mock, condition, testTimeout, testTick))
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

		True(t, Eventually(t, condition, testTimeout, testTick))
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
			time.Sleep(5 * time.Millisecond)
			return true
		}

		False(t, Eventually(mock, condition, time.Millisecond, time.Microsecond))
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

		False(t, Eventually(mock, condition, testTimeout, testTick))

		t.Run("reported errors should include the context cancellation", func(t *testing.T) {
			// assert how this failure is reported
			Len(t, mock.errors, 2, "expected to have 2 error messages: 1 for the context canceled, 1 for the never met condition")

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
			True(t, hasContextCancelled, "expected a context cancelled error")
			True(t, hasFailedCondition, "expected a condition never satisfied error")
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
		True(t, Eventually(mock, condition, testTimeout, 1*time.Second))
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
		Len(t, recordedActions, expectedActions, "expected 6 actions to be recorded during this execution", "got:", len(recordedActions))
		True(t, sort.IntsAreSorted(recordedActions), "expected recorded actions to be ordered")
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
		True(t, result, "Condition should end while Eventually still runs.")
	})
}

func TestConditionEventuallyWithT(t *testing.T) {
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

		False(t, EventuallyWithT(mock, condition, testTimeout, testTick))

		const expectedErrors = 4
		Len(t, mock.errors, expectedErrors, "expected 2 errors from the condition, and 2 additional errors from Eventually")

		expectedCalls := int(testTimeout / testTick)
		if counter != expectedCalls && counter != expectedCalls+1 { // it may be 5 or 6 depending on how the test schedules
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

		True(t, EventuallyWithT(mock, condition, testTimeout, testTick))
		Len(t, mock.errors, 0)
		const expectedCalls = 2
		Equal(t, expectedCalls, counter, "Condition is expected to have been called 2 times")
	})

	t.Run("should complete with fail, on a nanosecond tick", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(collect *CollectT) {
			Fail(collect, "condition fixed failure")
		}

		// To trigger race conditions, we run EventuallyWithT with a nanosecond tick.
		False(t, EventuallyWithT(mock, condition, testTimeout, time.Nanosecond))
		const expectedErrors = 3
		Len(t, mock.errors, expectedErrors, "expected 1 errors from the condition, and 2 additional errors from Eventually")
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

		False(t, EventuallyWithT(mock, condition, testTimeout, testTick))
		const expectedErrors = 3
		Len(t, mock.errors, expectedErrors, "expected 1 errors from the condition, and 2 additional errors from Eventually")
	})

	t.Run("should complete with success, with the ticker never used", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(*CollectT) {}

		// By making the tick longer than the total duration, we expect that this test would fail if
		// we didn't check the condition before the first tick elapses.
		True(t, EventuallyWithT(mock, condition, testTimeout, time.Second))
	})

	t.Run("should fail with a call to collect.FailNow", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		counter := 0

		// The call to FailNow cancels the execution context of EventuallyWithT.
		// so we don't have to wait for the timeout.
		condition := func(collect *CollectT) {
			counter++
			collect.FailNow()
		}

		False(t, EventuallyWithT(mock, condition, 30*time.Minute, testTick))
		const expectedErrors = 2
		Len(t, mock.errors, expectedErrors) // we have 0 accumulated error + 2 errors from EventuallyWithT (includes the timeout)
		if counter != 1 {
			t.Errorf("expected the condition function to have been called only once, but got: %d", counter)
		}
	})
}

func TestConditionNever(t *testing.T) {
	t.Parallel()

	t.Run("should never be true", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func() bool {
			return false
		}

		True(t, Never(mock, condition, testTimeout, testTick))
	})

	t.Run("should never be true, on timeout", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func() bool {
			time.Sleep(2 * testTick)
			// eventually returns true, after timeout
			return true
		}

		True(t, Never(mock, condition, testTick, 1*time.Millisecond))
	})

	t.Run("should never be true fails", func(t *testing.T) {
		// checks Never with a condition that returns true on second call.
		t.Parallel()

		mock := new(errorsCapturingT)
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

		False(t, Never(mock, condition, testTimeout, testTick))
	})

	t.Run("should never be true fails, with ticker never triggered", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		// By making the tick longer than the total duration, we expect that this test would fail if
		// we didn't check the condition before the first tick elapses.
		condition := func() bool { return true }
		False(t, Never(mock, condition, testTimeout, time.Second))
	})

	t.Run("should never be true fails, with parent test failing", func(t *testing.T) {
		t.Parallel()

		parentCtx, failParent := context.WithCancel(context.WithoutCancel(t.Context()))
		mock := new(errorsCapturingT).WithContext(parentCtx)
		condition := func() bool {
			failParent() // cancels the parent context, which results in Never to fail
			return false
		}
		False(t, Never(mock, condition, testTimeout, time.Second))
	})
}
