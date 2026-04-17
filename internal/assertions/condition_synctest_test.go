// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"testing/synctest"
	"time"
)

// ===========================================================================
// Dual-path test runner.
// ===========================================================================

// runDualPath runs fn twice: once with real time, once inside a synctest
// bubble. fn should use plain (non-wrapped) conditions and a mock T so
// that failures can be verified without polluting the outer test.
//
// NOTE: fn MUST NOT call t.Parallel(). [synctest.Test] forbids t.Parallel()
// inside a bubble. This restriction is transparent to real users of the
// async assertions — users do not call t.Parallel() inside condition
// functions.
func runDualPath(t *testing.T, name string, fn func(t *testing.T)) {
	t.Helper()
	t.Run(name+"/real-time", fn)
	t.Run(name+"/synctest", func(t *testing.T) {
		synctest.Test(t, fn)
	})
}

// ===========================================================================
// Dual-path tests.
// ===========================================================================

// TestConditionDualPath_EventuallyBehavior exercises [Eventually]'s core
// behavior through both real-time and bubble-wrapped test harnesses.
// Using the harness-level bubble means the mock captures failures even
// under fake time — the best of both worlds for behavior parity tests.
func TestConditionDualPath_EventuallyBehavior(t *testing.T) {
	runDualPath(t, "succeeds on first call", func(t *testing.T) {
		mock := new(errorsCapturingT)
		if !Eventually(mock, func() bool { return true }, testTimeout, testTick) {
			t.Error("expected success")
		}
		if len(mock.errors) != 0 {
			t.Errorf("expected no errors, got %d", len(mock.errors))
		}
	})

	runDualPath(t, "fails on persistent false", func(t *testing.T) {
		mock := new(errorsCapturingT)
		if Eventually(mock, func() bool { return false }, testTimeout, testTick) {
			t.Error("expected failure")
		}
		if len(mock.errors) == 0 {
			t.Error("expected mock to capture at least one error")
		}
	})

	runDualPath(t, "succeeds after a few ticks", func(t *testing.T) {
		mock := new(errorsCapturingT)
		var counter int
		var mu sync.Mutex
		cond := func() bool {
			mu.Lock()
			defer mu.Unlock()
			counter++

			return counter >= 3
		}

		if !Eventually(mock, cond, testTimeout, testTick) {
			t.Error("expected success")
		}
		mu.Lock()
		got := counter
		mu.Unlock()
		if got < 3 {
			t.Errorf("expected at least 3 calls, got %d", got)
		}
	})
}

// TestConditionDualPath_EventuallyWithErrorBehavior exercises the
// context/error-returning variant of [Eventually] through both paths.
func TestConditionDualPath_EventuallyWithErrorBehavior(t *testing.T) {
	runDualPath(t, "succeeds after returning transient errors", func(t *testing.T) {
		mock := new(errorsCapturingT)
		state := 0
		cond := func(_ context.Context) error {
			defer func() { state++ }()
			if state < 2 {
				return errors.New("not ready yet")
			}

			return nil
		}
		if !Eventually(mock, cond, testTimeout, testTick) {
			t.Error("expected Eventually to return true")
		}
	})

	runDualPath(t, "fails on persistent error", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(_ context.Context) error {
			return errors.New("persistent error")
		}
		if Eventually(mock, cond, testTimeout, testTick) {
			t.Error("expected Eventually to return false")
		}
	})

	runDualPath(t, "receives non-nil context", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(ctx context.Context) error {
			if ctx == nil {
				return errors.New("expected non-nil context")
			}

			return nil
		}
		if !Eventually(mock, cond, testTimeout, testTick) {
			t.Error("expected Eventually to return true")
		}
	})
}

// TestConditionDualPath_ConsistentlyWithErrorBehavior exercises the
// context/error-returning variant of [Consistently] through both paths.
func TestConditionDualPath_ConsistentlyWithErrorBehavior(t *testing.T) {
	runDualPath(t, "succeeds when always nil", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(_ context.Context) error { return nil }
		if !Consistently(mock, cond, testTimeout, testTick) {
			t.Error("expected Consistently to return true")
		}
	})

	runDualPath(t, "fails on persistent error", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(_ context.Context) error {
			return errors.New("something went wrong")
		}
		if Consistently(mock, cond, testTimeout, testTick) {
			t.Error("expected Consistently to return false")
		}
	})

	runDualPath(t, "fails when error appears on second call", func(t *testing.T) {
		mock := new(errorsCapturingT)
		// Channel created inside fn — under synctest, it is bubble-owned.
		returns := make(chan error, 2)
		returns <- nil
		returns <- errors.New("something went wrong")
		defer close(returns)

		cond := func(_ context.Context) error { return <-returns }
		if Consistently(mock, cond, testTimeout, testTick) {
			t.Error("expected Consistently to return false")
		}
	})
}

// TestConditionDualPath_NeverBehavior exercises [Never] through both paths.
func TestConditionDualPath_NeverBehavior(t *testing.T) {
	runDualPath(t, "succeeds when condition never true", func(t *testing.T) {
		mock := new(errorsCapturingT)
		if !Never(mock, func() bool { return false }, testTimeout, testTick) {
			t.Error("expected Never to return true")
		}
	})

	runDualPath(t, "fails when condition becomes true", func(t *testing.T) {
		mock := new(errorsCapturingT)
		var counter int
		var mu sync.Mutex
		cond := func() bool {
			mu.Lock()
			defer mu.Unlock()
			counter++

			return counter == 2
		}

		if Never(mock, cond, testTimeout, testTick) {
			t.Error("expected Never to return false")
		}
	})
}

// TestConditionDualPath_ConsistentlyBehavior exercises [Consistently] through both paths.
func TestConditionDualPath_ConsistentlyBehavior(t *testing.T) {
	runDualPath(t, "succeeds when condition always true", func(t *testing.T) {
		mock := new(errorsCapturingT)
		if !Consistently(mock, func() bool { return true }, testTimeout, testTick) {
			t.Error("expected Consistently to return true")
		}
	})

	runDualPath(t, "fails when condition becomes false", func(t *testing.T) {
		mock := new(errorsCapturingT)
		var counter int
		var mu sync.Mutex
		cond := func() bool {
			mu.Lock()
			defer mu.Unlock()
			counter++

			return counter < 3
		}

		if Consistently(mock, cond, testTimeout, testTick) {
			t.Error("expected Consistently to return false")
		}
	})
}

// TestConditionDualPath_EventuallyWithBehavior exercises [EventuallyWith] through both paths.
func TestConditionDualPath_EventuallyWithBehavior(t *testing.T) {
	runDualPath(t, "succeeds when no errors collected", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(_ *CollectT) {}
		if !EventuallyWith(mock, cond, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return true")
		}
	})

	runDualPath(t, "fails when errors persistently collected", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(c *CollectT) { Fail(c, "boom") }
		if EventuallyWith(mock, cond, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return false")
		}
	})
}

// TestConditionDualPath_EventuallyWithContextBehavior exercises the
// context variant of [EventuallyWith] (`func(ctx, *CollectT)`) through
// both paths.
func TestConditionDualPath_EventuallyWithContextBehavior(t *testing.T) {
	runDualPath(t, "succeeds after a few calls via context variant", func(t *testing.T) {
		mock := new(errorsCapturingT)
		counter := 0
		cond := func(_ context.Context, c *CollectT) {
			counter++
			True(c, counter == 2)
		}
		if !EventuallyWith(mock, cond, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return true")
		}
		if len(mock.errors) != 0 {
			t.Errorf("expected 0 errors, got %d", len(mock.errors))
		}
		if counter != 2 {
			t.Errorf("expected exactly 2 calls, got %d", counter)
		}
	})

	runDualPath(t, "fails on persistent collected failure via context variant", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(_ context.Context, c *CollectT) {
			Fail(c, "fixed failure")
		}
		if EventuallyWith(mock, cond, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return false")
		}
	})

	runDualPath(t, "receives non-nil context", func(t *testing.T) {
		mock := new(errorsCapturingT)
		cond := func(ctx context.Context, c *CollectT) {
			if ctx == nil {
				Fail(c, "expected non-nil context")
			}
		}
		if !EventuallyWith(mock, cond, testTimeout, testTick) {
			t.Error("expected EventuallyWith to return true")
		}
	})
}

// TestConditionDualPath_EventuallySucceedQuickly verifies that Eventually
// checks the condition BEFORE the first tick — by using a tick longer than
// the total duration, only the initial-check path can succeed.
func TestConditionDualPath_EventuallySucceedQuickly(t *testing.T) {
	t.Parallel()
	runDualPath(t, "should succeed before the first tick", dualEventuallySucceedBeforeFirstTick)
}

func dualEventuallySucceedBeforeFirstTick(t *testing.T) {
	mock := new(errorsCapturingT)
	cond := func() bool { return true }

	// Tick longer than the total duration: only the initial check can succeed.
	if !Eventually(mock, cond, testTimeout, 1*time.Second) {
		t.Error("expected Eventually to return true before first tick")
	}
}

// TestConditionDualPath_EventuallyTimeoutBehavior verifies that Eventually
// fails correctly when the condition is slower than the timeout (issue 805)
// and when the parent context is cancelled. Both subtests run under real
// time and inside a synctest bubble — reassurance that no semantic shift
// occurs when switching modes.
func TestConditionDualPath_EventuallyTimeoutBehavior(t *testing.T) {
	t.Parallel()
	runDualPath(t, "should fail on timeout", dualEventuallyTimeoutOnSlowCondition)
	runDualPath(t, "should fail when parent context is cancelled", dualEventuallyTimeoutOnParentCancellation)
}

func dualEventuallyTimeoutOnSlowCondition(t *testing.T) {
	mock := new(errorsCapturingT)
	// Condition returns long after the Eventually timeout.
	cond := func() bool {
		time.Sleep(100 * time.Millisecond)

		return true
	}

	if Eventually(mock, cond, time.Millisecond, time.Microsecond) {
		t.Error("expected Eventually to return false on timeout")
	}
}

func dualEventuallyTimeoutOnParentCancellation(t *testing.T) {
	parentCtx, failParent := context.WithCancel(context.WithoutCancel(t.Context()))
	mock := new(errorsCapturingT).WithContext(parentCtx)

	cond := func() bool {
		time.Sleep(testTick)
		failParent() // cancels the parent context mid-assertion
		time.Sleep(2 * testTick)

		return true
	}

	if Eventually(mock, cond, testTimeout, testTick) {
		t.Error("expected Eventually to return false when parent context is cancelled")
	}

	// Flattened from the original nested t.Run: nested subtests are forbidden
	// inside a synctest bubble. These assertions verify that the reported
	// errors include both the context-cancellation cause and the
	// "never satisfied" marker.
	if len(mock.errors) != 2 {
		t.Errorf("expected 2 error messages (1 for context canceled, 1 for never satisfied), got %d", len(mock.errors))
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
}

// TestConditionDualPath_PollUntilTimeoutBehavior exercises the shared
// poll-until-timeout subtests for [Never] and [Consistently] through both
// real-time and synctest paths. These subtests cover timing-independent
// invariants (initial-check before first tick, flipped-condition failure,
// parent-context cancellation) and benefit from dual-path for determinism.
func TestConditionDualPath_PollUntilTimeoutBehavior(t *testing.T) {
	t.Parallel()
	for c := range pollUntilTimeoutCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			runDualPath(t, "succeed with constant good value", dualPollCaseConstantGood(c))
			runDualPath(t, "succeed on timeout with slow bad value", dualPollCaseSlowBad(c))
			runDualPath(t, "fail when condition flips on second call", dualPollCaseFlipOnSecond(c))
			runDualPath(t, "fail before first tick with constant bad value", dualPollCaseBadBeforeFirstTick(c))
			runDualPath(t, "fail when parent context is cancelled", dualPollCaseParentCancelled(c))
		})
	}
}

// ---------------------------------------------------------------------------
// Dual-path subtest bodies for PollUntilTimeout (Never / Consistently).
// ---------------------------------------------------------------------------

func dualPollCaseConstantGood(c pollUntilTimeoutCase) func(*testing.T) {
	return func(t *testing.T) {
		mock := new(errorsCapturingT)
		if !c.assertion(mock, func() bool { return c.goodValue }, testTimeout, testTick) {
			t.Errorf("expected %s to return true", c.name)
		}
	}
}

func dualPollCaseSlowBad(c pollUntilTimeoutCase) func(*testing.T) {
	return func(t *testing.T) {
		mock := new(errorsCapturingT)
		badValue := !c.goodValue
		cond := func() bool {
			time.Sleep(2 * testTick)

			return badValue // returns bad value, but only after timeout
		}
		if !c.assertion(mock, cond, testTick, 1*time.Millisecond) {
			t.Errorf("expected %s to return true on timeout", c.name)
		}
	}
}

func dualPollCaseFlipOnSecond(c pollUntilTimeoutCase) func(*testing.T) {
	return func(t *testing.T) {
		mock := new(errorsCapturingT)
		// Channel created inside fn — bubble-owned under synctest.
		badValue := !c.goodValue
		returns := make(chan bool, 2)
		returns <- c.goodValue
		returns <- badValue
		defer close(returns)

		cond := func() bool { return <-returns }
		if c.assertion(mock, cond, testTimeout, testTick) {
			t.Errorf("expected %s to return false", c.name)
		}
	}
}

func dualPollCaseBadBeforeFirstTick(c pollUntilTimeoutCase) func(*testing.T) {
	return func(t *testing.T) {
		mock := new(errorsCapturingT)
		badValue := !c.goodValue
		// Tick longer than total duration: the initial-check path must detect
		// the bad value before any tick elapses.
		if c.assertion(mock, func() bool { return badValue }, testTimeout, time.Second) {
			t.Errorf("expected %s to return false", c.name)
		}
	}
}

func dualPollCaseParentCancelled(c pollUntilTimeoutCase) func(*testing.T) {
	return func(t *testing.T) {
		parentCtx, failParent := context.WithCancel(context.WithoutCancel(t.Context()))
		mock := new(errorsCapturingT).WithContext(parentCtx)
		cond := func() bool {
			failParent() // cancels the parent context

			return c.goodValue
		}
		if c.assertion(mock, cond, testTimeout, time.Second) {
			t.Errorf("expected %s to return false when parent context is cancelled", c.name)
		}
	}
}

// TestConditionDualPath_EventuallyWithCollectBehavior exercises the
// behavioral subtests of [EventuallyWith] through both paths. The
// [CollectT]-based invariants (FailNow retries, Cancel short-circuits,
// initial-check before first tick, etc.) are independent of timing and
// work identically under real time and fake time.
//
// The nanosecond-tick "race trigger" subtest of [EventuallyWith] is NOT
// migrated here — see [TestConditionEventuallyWith] for the rationale.
func TestConditionDualPath_EventuallyWithCollectBehavior(t *testing.T) {
	runDualPath(t, "should complete with false (tolerant count)", dualEventuallyWithCompleteFalse)
	runDualPath(t, "should complete with true", dualEventuallyWithCompleteTrue)
	runDualPath(t, "should complete with fail, with latest failed condition", dualEventuallyWithFailLatest)
	runDualPath(t, "should complete with success, with the ticker never used", dualEventuallyWithCompleteSuccess)
	runDualPath(t, "collect.FailNow only fails the current tick (poller retries)", dualEventuallyWithFailNowRetries)
	runDualPath(t, "collect.FailNow allows convergence on a later tick", dualEventuallyWithFailNowConverges)
	runDualPath(t, "collect.Cancel aborts the whole assertion immediately", dualEventuallyWithCancelAborts)
	runDualPath(t, "collect.Cancelf aborts with a custom message", dualEventuallyWithCancelfAborts)
}

// ---------------------------------------------------------------------------
// Dual-path subtest bodies for EventuallyWith.
//
// These are invoked by runDualPath — both under real time and inside a
// synctest bubble. They MUST NOT call t.Parallel(): synctest forbids it.
// ---------------------------------------------------------------------------

func dualEventuallyWithCompleteFalse(t *testing.T) {
	mock := new(errorsCapturingT)
	counter := 0
	cond := func(c *CollectT) {
		counter++
		Fail(c, "condition fixed failure")
		Fail(c, "another condition fixed failure")
	}

	if EventuallyWith(mock, cond, testTimeout, testTick) {
		t.Error("expected EventuallyWith to return false")
	}

	// Real-time path has scheduler jitter; synctest path is exact.
	// Both paths fall within this tolerance.
	const expectedErrors = 4
	if len(mock.errors) < expectedErrors-1 || len(mock.errors) > expectedErrors {
		t.Errorf("expected %d errors (2 from condition, 2 from Eventually), got %d", expectedErrors, len(mock.errors))
	}

	expectedCalls := int(testTimeout / testTick)
	if counter < expectedCalls-1 || counter > expectedCalls+1 {
		t.Errorf("expected %d calls, got %d", expectedCalls, counter)
	}
}

func dualEventuallyWithCompleteTrue(t *testing.T) {
	mock := new(errorsCapturingT)
	counter := 0
	cond := func(c *CollectT) {
		counter++
		True(c, counter == 2)
	}

	if !EventuallyWith(mock, cond, testTimeout, testTick) {
		t.Error("expected EventuallyWith to return true")
	}
	if len(mock.errors) != 0 {
		t.Errorf("expected 0 errors, got %d", len(mock.errors))
	}
	if counter != 2 {
		t.Errorf("expected condition to be called 2 times, got %d", counter)
	}
}

func dualEventuallyWithFailLatest(t *testing.T) {
	mock := new(errorsCapturingT)
	// Channel created inside fn — bubble-owned under synctest.
	mustSleep := make(chan bool, 2)
	mustSleep <- false
	mustSleep <- true
	close(mustSleep)

	cond := func(c *CollectT) {
		if <-mustSleep {
			// Ensure the second condition runs longer than the timeout.
			// Real time: 1s real sleep. Fake time: 1s fake sleep.
			time.Sleep(time.Second)

			return
		}
		Fail(c, "condition fixed failure")
	}

	if EventuallyWith(mock, cond, testTimeout, testTick) {
		t.Error("expected EventuallyWith to return false")
	}
	const expectedErrors = 3
	if len(mock.errors) != expectedErrors {
		t.Errorf("expected %d errors (1 from condition, 2 from Eventually), got %d", expectedErrors, len(mock.errors))
	}
}

func dualEventuallyWithCompleteSuccess(t *testing.T) {
	mock := new(errorsCapturingT)
	cond := func(*CollectT) {}

	// Tick longer than total duration: the initial-check path must succeed.
	if !EventuallyWith(mock, cond, testTimeout, time.Second) {
		t.Error("expected EventuallyWith to return true")
	}
}

func dualEventuallyWithFailNowRetries(t *testing.T) {
	mock := new(errorsCapturingT)
	var counter int
	var mu sync.Mutex

	cond := func(c *CollectT) {
		mu.Lock()
		counter++
		mu.Unlock()
		c.FailNow()
	}

	if EventuallyWith(mock, cond, testTimeout, testTick) {
		t.Error("expected EventuallyWith to return false")
	}
	mu.Lock()
	got := counter
	mu.Unlock()
	if got < 2 {
		t.Errorf("expected the condition to be retried multiple times, got %d call(s)", got)
	}
}

func dualEventuallyWithFailNowConverges(t *testing.T) {
	mock := new(errorsCapturingT)
	var counter int
	var mu sync.Mutex

	cond := func(c *CollectT) {
		mu.Lock()
		counter++
		n := counter
		mu.Unlock()
		if n < 3 {
			c.FailNow()
		}
	}

	if !EventuallyWith(mock, cond, testTimeout, testTick) {
		t.Error("expected EventuallyWith to eventually return true")
	}
	if len(mock.errors) != 0 {
		t.Errorf("expected no errors reported on parent t after success, got %d: %v", len(mock.errors), mock.errors)
	}
}

func dualEventuallyWithCancelAborts(t *testing.T) {
	mock := new(errorsCapturingT)
	var counter int
	var mu sync.Mutex

	// The 30-minute timeout must NOT be waited on: in the real-time path
	// Cancel must short-circuit quickly; in the synctest path the fake
	// timeout costs zero real time anyway.
	cond := func(c *CollectT) {
		mu.Lock()
		counter++
		mu.Unlock()
		c.Cancel()
	}

	start := time.Now()
	if EventuallyWith(mock, cond, 30*time.Minute, testTick) {
		t.Error("expected EventuallyWith to return false")
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Errorf("expected Cancel to short-circuit, but EventuallyWith took %s", elapsed)
	}
	mu.Lock()
	got := counter
	mu.Unlock()
	if got != 1 {
		t.Errorf("expected the condition to be called only once, got: %d", got)
	}
	if len(mock.errors) == 0 {
		t.Error("expected at least one error reported on parent t after Cancel")
	}
}

func dualEventuallyWithCancelfAborts(t *testing.T) {
	mock := new(errorsCapturingT)
	var counter int
	var mu sync.Mutex

	cond := func(c *CollectT) {
		mu.Lock()
		counter++
		mu.Unlock()
		c.Cancelf("upstream %s is gone", "service-x")
	}

	start := time.Now()
	if EventuallyWith(mock, cond, 30*time.Minute, testTick) {
		t.Error("expected EventuallyWith to return false")
	}
	if elapsed := time.Since(start); elapsed > 5*time.Second {
		t.Errorf("expected Cancelf to short-circuit, but EventuallyWith took %s", elapsed)
	}
	mu.Lock()
	got := counter
	mu.Unlock()
	if got != 1 {
		t.Errorf("expected the condition to be called once, got %d", got)
	}

	foundCustom := false
	for _, err := range mock.errors {
		if strings.Contains(err.Error(), "upstream service-x is gone") {
			foundCustom = true

			break
		}
	}
	if !foundCustom {
		t.Errorf("expected custom Cancelf message in errors, got: %v", mock.errors)
	}
}

// ===========================================================================
// API-level tests — verify the WithSynctest wrapper types activate the
// internal bubble when t is a concrete *testing.T.
// ===========================================================================

// TestSynctest_EventuallyDeterministicCount verifies that the WithSynctest
// wrapper activates a bubble, enabling exact tick counts under fake time.
func TestSynctest_EventuallyDeterministicCount(t *testing.T) {
	var counter int
	var mu sync.Mutex
	const target = 3

	cond := WithSynctest(func() bool {
		mu.Lock()
		defer mu.Unlock()
		counter++

		return counter == target
	})

	if !Eventually(t, cond, testTimeout, testTick) {
		t.Fatal("expected Eventually to succeed")
	}
	mu.Lock()
	got := counter
	mu.Unlock()
	if got != target {
		t.Errorf("expected exactly %d calls under fake time, got %d", target, got)
	}
}

// TestSynctest_EventuallyHugeTimeoutInstant verifies that a 1-hour timeout
// with a 1-minute tick completes in milliseconds of real wall time when
// the condition succeeds — proving fake time is used.
func TestSynctest_EventuallyHugeTimeoutInstant(t *testing.T) {
	var counter int
	var mu sync.Mutex
	cond := WithSynctest(func() bool {
		mu.Lock()
		defer mu.Unlock()
		counter++

		return counter == 5
	})

	start := time.Now()
	if !Eventually(t, cond, 1*time.Hour, 1*time.Minute) {
		t.Fatal("expected success")
	}
	if elapsed := time.Since(start); elapsed > 1*time.Second {
		t.Errorf("expected fake time to be instant, took %s", elapsed)
	}
}

// TestSynctest_EventuallyContextVariant verifies the [WithSynctestContext]
// wrapper (context/error condition form).
func TestSynctest_EventuallyContextVariant(t *testing.T) {
	var counter int
	var mu sync.Mutex
	cond := WithSynctestContext(func(_ context.Context) error {
		mu.Lock()
		defer mu.Unlock()
		counter++
		if counter < 3 {
			return errors.New("not yet")
		}

		return nil
	})

	if !Eventually(t, cond, testTimeout, testTick) {
		t.Fatal("expected success")
	}
	mu.Lock()
	got := counter
	mu.Unlock()
	if got != 3 {
		t.Errorf("expected exactly 3 calls, got %d", got)
	}
}

// TestSynctest_NeverUsesBubble verifies [Never] activates a bubble with the
// [WithSynctest] wrapper.
func TestSynctest_NeverUsesBubble(t *testing.T) {
	cond := WithSynctest(func() bool { return false })
	start := time.Now()
	if !Never(t, cond, 1*time.Hour, 1*time.Minute) {
		t.Fatal("expected Never to succeed")
	}
	if elapsed := time.Since(start); elapsed > 1*time.Second {
		t.Errorf("expected fake time to be instant, took %s", elapsed)
	}
}

// TestSynctest_ConsistentlyUsesBubble verifies [Consistently] activates a bubble.
func TestSynctest_ConsistentlyUsesBubble(t *testing.T) {
	cond := WithSynctest(func() bool { return true })
	start := time.Now()
	if !Consistently(t, cond, 1*time.Hour, 1*time.Minute) {
		t.Fatal("expected Consistently to succeed")
	}
	if elapsed := time.Since(start); elapsed > 1*time.Second {
		t.Errorf("expected fake time to be instant, took %s", elapsed)
	}
}

// TestSynctest_EventuallyWithUsesBubble verifies [EventuallyWith] activates
// a bubble with the [WithSynctestCollect] wrapper.
func TestSynctest_EventuallyWithUsesBubble(t *testing.T) {
	var counter int
	var mu sync.Mutex
	cond := WithSynctestCollect(func(c *CollectT) {
		mu.Lock()
		counter++
		n := counter
		mu.Unlock()
		True(c, n >= 3)
	})

	if !EventuallyWith(t, cond, testTimeout, testTick) {
		t.Fatal("expected EventuallyWith to succeed")
	}
	mu.Lock()
	got := counter
	mu.Unlock()
	if got != 3 {
		t.Errorf("expected exactly 3 calls, got %d", got)
	}
}

// TestSynctest_EventuallyWithContextVariant verifies the
// [WithSynctestCollectContext] wrapper.
func TestSynctest_EventuallyWithContextVariant(t *testing.T) {
	var counter int
	var mu sync.Mutex
	cond := WithSynctestCollectContext(func(_ context.Context, c *CollectT) {
		mu.Lock()
		counter++
		n := counter
		mu.Unlock()
		True(c, n >= 3)
	})

	if !EventuallyWith(t, cond, testTimeout, testTick) {
		t.Fatal("expected EventuallyWith to succeed")
	}
}

// TestSynctest_FallbackOnMock verifies that passing a non-*testing.T mock
// with WithSynctest falls back to real-time polling (no bubble).
func TestSynctest_FallbackOnMock(t *testing.T) {
	mock := new(errorsCapturingT)
	var counter int
	var mu sync.Mutex
	cond := WithSynctest(func() bool {
		mu.Lock()
		defer mu.Unlock()
		counter++

		return false
	})

	start := time.Now()
	if Eventually(mock, cond, testTimeout, testTick) {
		t.Error("expected Eventually to return false on mock")
	}
	elapsed := time.Since(start)
	// Real-time polling: should take close to testTimeout.
	if elapsed < testTimeout/2 {
		t.Errorf("expected real-time polling on mock path, took only %s", elapsed)
	}
}

// TestSynctest_PanicRecovery verifies panic recovery works through the
// bubble — recoverCondition treats a panic as a failed tick, and the
// poller retries.
func TestSynctest_PanicRecovery(t *testing.T) {
	var counter int
	var mu sync.Mutex
	cond := WithSynctest(func() bool {
		mu.Lock()
		counter++
		n := counter
		mu.Unlock()
		if n < 3 {
			panic(fmt.Sprintf("boom %d", n))
		}

		return true
	})

	if !Eventually(t, cond, testTimeout, testTick) {
		t.Fatal("expected Eventually to succeed after recovering panics")
	}
}

// TestSynctest_SlowConditionNoLeak verifies no goroutine leak when the
// condition sleeps longer than the tick interval. If a goroutine leaked,
// synctest.Test would deadlock waiting for it to exit.
func TestSynctest_SlowConditionNoLeak(t *testing.T) {
	var counter int
	var mu sync.Mutex
	cond := WithSynctest(func() bool {
		mu.Lock()
		counter++
		n := counter
		mu.Unlock()
		// time.Sleep advances fake clock when all goroutines block durably.
		time.Sleep(2 * testTick)

		return n >= 2
	})

	start := time.Now()
	if !Eventually(t, cond, testTimeout, testTick) {
		t.Fatal("expected success")
	}
	if elapsed := time.Since(start); elapsed > 1*time.Second {
		t.Errorf("expected fake-time sleeps to cost no real time, took %s", elapsed)
	}
}
