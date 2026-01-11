// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// Condition uses a [Comparison] to assert a complex condition.
//
// # Usage
//
//	assertions.Condition(t, func() bool { return myCondition })
//
// # Examples
//
//	success:  func() bool { return true }
//	failure:  func() bool { return false }
func Condition(t T, comp Comparison, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	result := comp()
	if !result {
		Fail(t, "condition failed", msgAndArgs...)
	}

	return result
}

// Eventually asserts that the given condition will be met in waitFor time,
// periodically checking the target function on each tick.
//
// [Eventually] waits until the condition returns true, for at most waitFor,
// or until the parent context of the test is cancelled.
//
// If the condition takes longer than waitFor to complete, [Eventually] fails
// but waits for the current condition execution to finish before returning.
//
// For long-running conditions to be interrupted early, check [testing.T.Context]
// which is cancelled on test failure.
//
// # Usage
//
//	assertions.Eventually(t, func() bool { return true }, time.Second, 10*time.Millisecond)
//
// # Concurrency
//
// The condition function is never executed in parallel: only one goroutine executes it.
// It may write to variables outside its scope without triggering race conditions.
//
// A blocking condition will cause [Eventually] to hang until it returns.
//
// # Examples
//
//	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
func Eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return eventually(t, condition, waitFor, tick, msgAndArgs...)
}

// Never asserts that the given condition is never satisfied within waitFor time,
// periodically checking the target function at each tick.
//
// [Never] is the opposite of [Eventually]. It succeeds if the waitFor timeout
// is reached without the condition ever returning true.
//
// If the parent context is cancelled before the timeout, [Never] fails.
//
// # Usage
//
//	assertions.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond)
//
// # Concurrency
//
// The condition function is never executed in parallel: only one goroutine executes it.
// It may write to variables outside its scope without triggering race conditions.
//
// A blocking condition will cause [Never] to hang until it returns.
//
// # Examples
//
//	success:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
func Never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return never(t, condition, waitFor, tick, msgAndArgs...)
}

// EventuallyWithT asserts that the given condition will be met in waitFor time,
// periodically checking the target function at each tick.
//
// In contrast to [Eventually], the condition function is supplied with a [CollectT]
// to accumulate errors from calling other assertions.
//
// The condition is considered "met" if no errors are raised in a tick.
// The supplied [CollectT] collects all errors from one tick.
//
// If the condition is not met before waitFor, the collected errors from the
// last tick are copied to t.
//
// Calling [CollectT.FailNow] cancels the condition immediately and fails the assertion.
//
// # Usage
//
//	externalValue := false
//	go func() {
//		time.Sleep(8*time.Second)
//		externalValue = true
//	}()
//
//	assertions.EventuallyWithT(t, func(c *assertions.CollectT) {
//		// add assertions as needed; any assertion failure will fail the current tick
//		assertions.True(c, externalValue, "expected 'externalValue' to be true")
//	}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
//
// # Concurrency
//
// The condition function is never executed in parallel: only one goroutine executes it.
// It may write to variables outside its scope without triggering race conditions.
//
// # Examples
//
//	success: func(c *CollectT) { True(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//	failure: func(c *CollectT) { False(c,true) }, 100*time.Millisecond, 20*time.Millisecond
func EventuallyWithT(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return eventuallyWithT(t, condition, waitFor, tick, msgAndArgs...)
}

func eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return pollCondition(t,
		condition, waitFor, tick,
		pollOptions{
			mode:        pollUntilTrue,
			failMessage: "condition never satisfied",
		},
		msgAndArgs...,
	)
}

func never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return pollCondition(t,
		condition, waitFor, tick,
		pollOptions{
			mode:        pollUntilTimeout,
			failMessage: "condition satisfied",
		},
		msgAndArgs...,
	)
}

func eventuallyWithT(t T, collectCondition func(collector *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var lastCollectedErrors []error
	var cancelFunc func() // will be set by pollCondition via onSetup

	condition := func() bool {
		collector := new(CollectT).withCancelFunc(cancelFunc)
		collectCondition(collector)
		if collector.failed() {
			lastCollectedErrors = collector.collected()
			return false
		}

		return true
	}

	copyCollected := func(tt T) {
		for _, err := range lastCollectedErrors {
			tt.Errorf("%v", err)
		}
	}

	return pollCondition(t,
		condition, waitFor, tick,
		pollOptions{
			mode:        pollUntilTrue,
			failMessage: "condition never satisfied",
			onFailure:   copyCollected,
			onSetup:     func(cancel func()) { cancelFunc = cancel },
		},
		msgAndArgs...,
	)
}

// pollMode determines how the condition polling should behave.
type pollMode int

const (
	// pollUntilTrue succeeds when condition returns true (for Eventually).
	pollUntilTrue pollMode = iota
	// pollUntilTimeout succeeds when timeout is reached without condition being true (for Never).
	pollUntilTimeout
)

// pollOptions configures the condition polling behavior.
type pollOptions struct {
	mode        pollMode
	failMessage string              // error message added at the end of the stack
	onFailure   func(t T)           // called on failure (e.g., to copy collected errors)
	onSetup     func(cancel func()) // called after context setup to expose cancel function
}

// pollCondition is the common implementation for eventually, never, and eventuallyWithT.
// It polls a condition function at regular intervals until success or timeout.
//
//nolint:gocognit,gocyclo,cyclop // A refactoring is planned for this complex function.
func pollCondition(t T, condition func() bool, waitFor, tick time.Duration, opts pollOptions, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var parentCtx context.Context
	if withContext, ok := t.(contextualizer); ok {
		parentCtx = withContext.Context()
	}
	if parentCtx == nil {
		parentCtx = context.Background()
	}

	// For pollUntilTimeout (Never), we detach from parent cancellation
	// so that timeout reaching is a success, not a failure.
	var ctx context.Context
	var cancel context.CancelFunc
	if opts.mode == pollUntilTimeout {
		ctx, cancel = context.WithTimeout(context.WithoutCancel(parentCtx), waitFor)
	} else {
		ctx, cancel = context.WithTimeout(parentCtx, waitFor)
	}
	defer cancel()

	// Allow caller to capture the cancel function (for eventuallyWithT's CollectT)
	if opts.onSetup != nil {
		opts.onSetup(cancel)
	}

	var reported atomic.Bool
	failFunc := func(reason string) {
		if reported.CompareAndSwap(false, true) {
			if reason != "" {
				t.Errorf("%s", reason)
			}
			Fail(t, opts.failMessage, msgAndArgs...)
		}
	}

	conditionChan := make(chan func() bool, 1)
	doneChan := make(chan struct{})

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	// Check the condition once first on the initial call.
	conditionChan <- condition

	var wg sync.WaitGroup

	// Goroutine 1: Poll for the condition at every tick
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			if opts.mode == pollUntilTimeout {
				// For Never: check parent context separately
				select {
				case <-parentCtx.Done():
					failFunc(parentCtx.Err().Error())
					return
				case <-ctx.Done():
					return // timeout reached = success for Never
				case <-doneChan:
					return
				case <-ticker.C:
					select {
					case <-parentCtx.Done():
						failFunc(parentCtx.Err().Error())
						return
					case <-ctx.Done():
						return
					case <-doneChan:
						return
					case conditionChan <- condition:
					}
				}
			} else {
				// For Eventually: parent cancellation flows through ctx
				select {
				case <-ctx.Done():
					failFunc(ctx.Err().Error())
					return
				case <-doneChan:
					return
				case <-ticker.C:
					select {
					case <-ctx.Done():
						failFunc(ctx.Err().Error())
						return
					case <-doneChan:
						return
					case conditionChan <- condition:
					}
				}
			}
		}
	}()

	// Goroutine 2: Execute the condition and check results
	wg.Add(1)
	go func() {
		defer wg.Done()

		for {
			if opts.mode == pollUntilTimeout {
				select {
				case <-parentCtx.Done():
					failFunc(parentCtx.Err().Error())
					return
				case <-ctx.Done():
					return // timeout = success
				case fn := <-conditionChan:
					if fn() {
						close(doneChan) // condition true = failure for Never
						return
					}
				}
			} else {
				select {
				case <-ctx.Done():
					failFunc(ctx.Err().Error())
					return
				case fn := <-conditionChan:
					if fn() {
						close(doneChan) // condition true = success
						return
					}
				}
			}
		}
	}()

	wg.Wait()

	// Determine success based on mode
	select {
	case <-doneChan:
		if opts.mode == pollUntilTimeout {
			// For Never: doneChan closed means condition became true
			// But if timeout was reached first (ctx.Err != nil), it's still a success
			if ctx.Err() != nil {
				return true
			}
			// Condition became true before timeout = failure
			failFunc("")
			return false
		}
		// For Eventually: doneChan closed means condition became true
		if ctx.Err() != nil {
			// Timeout occurred before or during success
			if opts.onFailure != nil {
				opts.onFailure(t)
			}
			return false
		}
		return true
	default:
		// doneChan not closed
		if opts.mode == pollUntilTimeout {
			// For Never: timeout reached without condition being true = success
			// We should return a success, unless the parent context has failed.
			return parentCtx.Err() == nil
		}

		// opts.mode = pollUntilTrue
		// For Eventually: should not reach here (failFunc already called)
		if opts.onFailure != nil {
			opts.onFailure(t)
		}

		return false
	}
}

// CollectT implements the [T] interface and collects all errors.
//
// [CollectT] is specifically intended to be used with [EventuallyWithT] and
// should not be used outside of that context.
type CollectT struct {
	// Domain: condition
	//
	// Maintainer:
	//   1. FailNow() no longer just exits the go routine, but cancels the context of the caller instead before exiting.
	//   2. We no longer establish the distinction between c.error nil or empty. Non-empty is an error, full stop.
	//   2. Deprecated methods have been removed.

	// A slice of errors. Non-empty slice denotes a failure.
	// A c.FailNow() will thee lose accumulated errors
	errors []error

	// cancelContext cancels the parent context on FailNow()
	cancelContext func()
}

// Helper is like [testing.T.Helper] but does nothing.
func (*CollectT) Helper() {}

// Errorf collects the error.
func (c *CollectT) Errorf(format string, args ...any) {
	c.errors = append(c.errors, fmt.Errorf(format, args...))
}

// FailNow records a failure and cancels the parent [EventuallyWithT] context,
// before exiting the current go routine with [runtime.Goexit].
//
// This causes the assertion to fail immediately without waiting for a timeout.
func (c *CollectT) FailNow() {
	c.cancelContext()
	c.errors = append(c.errors, errors.New("failed now")) // so c.failed() is true (currently lost as not owned by another go routine)
	runtime.Goexit()
}

func (c *CollectT) failed() bool {
	return len(c.errors) != 0
}

func (c *CollectT) collected() []error {
	return c.errors
}

func (c *CollectT) withCancelFunc(cancel func()) *CollectT {
	c.cancelContext = cancel

	return c
}
