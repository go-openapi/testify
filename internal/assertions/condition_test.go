// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"iter"
	"slices"
	"sort"
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

// This test is deliberately NOT dual-path: it asserts that there are no leaking go routines
// when real time tickers are used. This is naturally verified when running in a syntest bubble.
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

// TestConditionEventuallyWith keeps only the nanosecond-tick "race trigger"
// subtest of [EventuallyWith]. All behavior-oriented subtests have been
// migrated to [TestConditionDualPath_EventuallyWithCollectBehavior] in
// condition_synctest_test.go, where they run under both real time and a
// synctest bubble.
//
// This test is deliberately NOT dual-path: it uses a nanosecond tick
// to force real-time scheduling races between the poller, the ticker,
// and the condition goroutine. Under synctest, ticks fire deterministically
// from a fake clock — so there are no real races to exercise. Keeping this
// test real-time-only preserves its purpose as a smoke test against
// concurrency regressions that only manifest under real scheduler pressure.
func TestConditionEventuallyWith(t *testing.T) {
	t.Parallel()

	t.Run("should complete with fail, on a nanosecond tick (real-time race trigger)", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		cond := func(c *CollectT) {
			Fail(c, "condition fixed failure")
		}

		// Nanosecond tick to provoke real-time scheduling races.
		if EventuallyWith(mock, cond, testTimeout, time.Nanosecond) {
			t.Error("expected EventuallyWith to return false")
		}
		const expectedErrors = 3
		if len(mock.errors) != expectedErrors {
			t.Errorf("expected %d errors (1 from condition, 2 from Eventually), got %d", expectedErrors, len(mock.errors))
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
			assertion: Never[func() bool],
			goodValue: false, // Never succeeds when the condition always returns false ("never true")
		},
		{
			name:      "Consistently",
			assertion: Consistently[func() bool],
			goodValue: true, // Consistently succeeds when the condition always returns true ("always true")
		},
	})
}

func TestConditionPanicRecovery(t *testing.T) {
	t.Parallel()

	t.Run("Eventually survives a panicking condition and retries", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		var counter int
		var mu sync.Mutex

		condition := func() bool {
			mu.Lock()
			counter++
			n := counter
			mu.Unlock()
			if n < 3 {
				panic("boom")
			}

			return true
		}

		if !Eventually(mock, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return true after recovering from panics")
		}
		mu.Lock()
		got := counter
		mu.Unlock()
		if got < 3 {
			t.Errorf("expected at least 3 calls, got %d", got)
		}
	})

	t.Run("Eventually fails when condition always panics", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func() bool {
			panic("persistent failure")
		}

		if Eventually(mock, condition, testTimeout, testTick) {
			t.Error("expected Eventually to return false when condition always panics")
		}
	})

	t.Run("Never fails when condition panics", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func() bool {
			panic("unexpected")
		}

		if Never(mock, condition, testTimeout, testTick) {
			t.Error("expected Never to return false when condition panics")
		}
	})

	t.Run("Consistently fails when condition panics", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func() bool {
			panic("unexpected")
		}

		if Consistently(mock, condition, testTimeout, testTick) {
			t.Error("expected Consistently to return false when condition panics")
		}
	})

	t.Run("EventuallyWith survives a panicking condition and retries", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		var counter int
		var mu sync.Mutex

		condition := func(_ *CollectT) {
			mu.Lock()
			counter++
			n := counter
			mu.Unlock()
			if n < 3 {
				panic("boom in collect")
			}
		}

		if !EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return true after recovering from panics")
		}
		mu.Lock()
		got := counter
		mu.Unlock()
		if got < 3 {
			t.Errorf("expected at least 3 calls, got %d", got)
		}
	})

	t.Run("EventuallyWith fails when condition always panics", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		condition := func(_ *CollectT) {
			panic("always panics")
		}

		if EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return false when condition always panics")
		}
	})

	t.Run("EventuallyWith collects panic error via sentinel", func(t *testing.T) {
		t.Parallel()

		mock := new(errorsCapturingT)
		var counter int
		var mu sync.Mutex

		condition := func(collect *CollectT) {
			mu.Lock()
			counter++
			n := counter
			mu.Unlock()

			if n == 1 {
				panic("boom on first tick")
			}
			// Subsequent ticks fail normally, preserving the panic error
			// from the first tick in lastCollectedErrors.
			Fail(collect, "still failing")
		}

		if EventuallyWith(mock, condition, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return false")
		}
	})

	t.Run("errConditionPanicked sentinel is detectable with errors.Is", func(t *testing.T) {
		t.Parallel()

		err := fmt.Errorf("%w: %v", errConditionPanicked, "test panic")
		if !errors.Is(err, errConditionPanicked) {
			t.Error("expected errors.Is to detect errConditionPanicked sentinel")
		}
	})
}
