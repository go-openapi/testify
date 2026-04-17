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
	"testing"
	"testing/synctest"
	"time"
)

// Condition uses a comparison function to assert a complex condition.
//
// # Usage
//
//	assertions.Condition(t, func() bool { return myCondition })
//
// # Examples
//
//	success:  func() bool { return true }
//	failure:  func() bool { return false }
func Condition(t T, comp func() bool, msgAndArgs ...any) bool {
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

// Eventually asserts that the given condition will be met before timeout,
// periodically checking the target function on each tick.
//
// [Eventually] waits until the condition returns true, at most until timeout,
// or until the parent context of the test is cancelled.
//
// If the condition takes longer than the timeout to complete, [Eventually] fails
// but waits for the current condition execution to finish before returning.
//
// For long-running conditions to be interrupted early, check [testing.T.Context]
// which is cancelled on test failure.
//
// # Usage
//
//	assertions.Eventually(t, func() bool { return true }, time.Second, 10*time.Millisecond)
//
// # Alternative condition signature
//
// The simplest form of condition is:
//
//	func() bool
//
// To build more complex cases, a condition may also be defined as:
//
//	func(context.Context) error
//
// It fails when an error has always been returned up to timeout (equivalent semantics to func() bool returns false),
// expressing "eventually returns no error (nil)".
//
// It will be executed with the context of the assertion, which inherits the [testing.T.Context] and
// is cancelled on timeout.
//
// The semantics of the three available async assertions read as follows.
//
//   - [Eventually] (func() bool) : "eventually returns true"
//
//   - [Never] (func() bool) : "never returns true"
//
//   - [Consistently] (func() bool): "always returns true"
//
//   - [Eventually] (func(ctx) error) : "eventually returns nil"
//
//   - [Never] (func(ctx) error) : not supported, use [Consistently] instead (avoids confusion with double negation)
//
//   - [Consistently] (func(ctx) error): "always returns nil"
//
// # Concurrency
//
// The condition function is always executed serially by a single goroutine. It is always executed at least once.
//
// It may thus write to variables outside its scope without triggering race conditions.
//
// A blocking condition will cause [Eventually] to hang until it returns.
//
// Notice that time ticks may be skipped if the condition takes longer than the tick interval.
//
// # Panic recovery
//
// If the condition panics, the panic is recovered and treated as a failed tick
// (equivalent to returning false or a non-nil error). For [Eventually], this means
// the poller retries on the next tick — if a later tick succeeds, the assertion
// succeeds. For [Never] and [Consistently], a panic is treated as the condition
// erroring, which causes immediate failure.
//
// The recovered panic is wrapped as an error with the sentinel [errConditionPanicked],
// detectable with [errors.Is].
//
// # Attention point
//
// Time-based tests may be flaky in a resource-constrained environment such as a CI runner and may produce
// counter-intuitive results, such as ticks or timeouts not firing in time as expected.
//
// To avoid flaky tests, always make sure that ticks and timeouts differ by at least an order of magnitude (tick <<
// timeout).
//
// # Synctest (opt-in)
//
// Wrap the condition with [WithSynctest] (or [WithSynctestContext]) to run
// the polling loop inside a [testing/synctest] bubble, which uses a fake
// clock. This eliminates timing-induced flakiness and makes the tick count
// deterministic. See [WithSynctest] for the constraints (no real I/O in
// the condition, requires `*testing.T`).
//
// # Examples
//
//	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
func Eventually[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	// Domain: condition
	// Opposite: Never
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return eventually(t, condition, timeout, tick, msgAndArgs...)
}

// Never asserts that the given condition is never satisfied until timeout,
// periodically checking the target function at each tick.
//
// [Never] is the opposite of [Eventually] ("at least once").
// It succeeds if the timeout is reached without the condition ever returning true.
//
// If the parent context is cancelled before the timeout, [Never] fails.
//
// # Usage
//
//	assertions.Never(t, func() bool { return false }, time.Second, 10*time.Millisecond)
//
// See also [Eventually] for details about using context, concurrency, and panic recovery.
//
// # Alternative condition signature
//
// The simplest form of condition is:
//
//	func() bool
//
// Use [Consistently] instead if you want to use a condition returning an error.
//
// # Panic recovery
//
// A panicking condition is treated as an error, causing [Never] to fail immediately.
// See [Eventually] for details.
//
// # Concurrency
//
// See [Eventually].
//
// # Attention point
//
// See [Eventually].
//
// # Synctest (opt-in)
//
// Wrap the condition with [WithSynctest] to run the polling loop inside a
// [testing/synctest] bubble, which uses a fake clock. This eliminates
// timing-induced flakiness and makes the tick count deterministic. See
// [WithSynctest] for the constraints (no real I/O in the condition,
// requires [*testing.T]). Note: [Never] does not accept the context/error
// form of condition, so [WithSynctestContext] does not apply here.
//
// # Examples
//
//	success:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
func Never[C NeverConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return never(t, condition, timeout, tick, msgAndArgs...)
}

// Consistently asserts that the given condition is always satisfied until timeout,
// periodically checking the target function at each tick.
//
// [Consistently] ("always") imposes a stronger constraint than [Eventually] ("at least once"):
// it checks at every tick that every occurrence of the condition is satisfied, whereas
// [Eventually] succeeds on the first occurrence of a successful condition.
//
// # Usage
//
//	assertions.Consistently(t, func() bool { return true }, time.Second, 10*time.Millisecond)
//
// See also [Eventually] for details about using context, concurrency, and panic recovery.
//
// # Alternative condition signature
//
// The simplest form of condition is:
//
//	func() bool
//
// The semantics of the assertion are "always returns true".
//
// To build more complex cases, a condition may also be defined as:
//
//	func(context.Context) error
//
// It fails as soon as an error is returned before timeout expressing "always returns no error (nil)"
//
// This is consistent with [Eventually] expressing "eventually returns no error (nil)".
//
// It will be executed with the context of the assertion, which inherits the [testing.T.Context] and
// is cancelled on timeout.
//
// # Panic recovery
//
// A panicking condition is treated as an error, causing [Consistently] to fail immediately.
// See [Eventually] for details.
//
// # Concurrency
//
// See [Eventually].
//
// # Attention point
//
// See [Eventually].
//
// # Synctest (opt-in)
//
// Wrap the condition with [WithSynctest] (or [WithSynctestContext]) to run
// the polling loop inside a [testing/synctest] bubble, which uses a fake
// clock. This eliminates timing-induced flakiness and makes the tick count
// deterministic. See [WithSynctest] for the constraints (no real I/O in
// the condition, requires [*testing.T]).
//
// # Examples
//
//	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
func Consistently[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return consistently(t, condition, timeout, tick, msgAndArgs...)
}

// EventuallyWith asserts that the given condition will be met before the timeout,
// periodically checking the target function at each tick.
//
// In contrast to [Eventually], the condition function is supplied with a [CollectT]
// to accumulate errors from calling other assertions.
//
// The condition is considered "met" if no errors are raised in a tick.
// The supplied [CollectT] collects all errors from one tick.
//
// If the condition is not met before the timeout, the collected errors from the
// last tick are copied to t.
//
// Calling [CollectT.FailNow] (directly, or transitively through [require] assertions)
// fails the current tick only: the poller will retry on the next tick. This means
// [require]-style assertions inside [EventuallyWith] behave naturally — they abort
// the current evaluation and let the polling loop converge.
//
// To abort the whole assertion immediately (e.g. when the condition can no longer
// be expected to succeed), call [CollectT.Cancel].
//
// # Usage
//
//	externalValue := false
//	go func() {
//		time.Sleep(8*time.Second)
//		externalValue = true
//	}()
//
//	assertions.EventuallyWith(t, func(c *assertions.CollectT) {
//		// add assertions as needed; any assertion failure will fail the current tick
//		assertions.True(c, externalValue, "expected 'externalValue' to be true")
//	},
//	10*time.Second,
//	1*time.Second,
//	"external state has not changed to 'true'; still false",
//	)
//
// # Concurrency
//
// The condition function is never executed in parallel: only one goroutine executes it.
// It may write to variables outside its scope without triggering race conditions.
//
// The condition is wrapped in its own goroutine, so a call to [runtime.Goexit]
// (e.g. via [require] assertions or [CollectT.FailNow]) cleanly aborts only the
// current tick.
//
// # Panic recovery
//
// If the condition panics, the panic is recovered and recorded as an error in the
// [CollectT] for that tick. The poller treats it as a failed tick and retries on the
// next one. If the assertion times out, the panic error is included in the collected
// errors reported on the parent t.
//
// See [Eventually] for the general panic recovery semantics.
//
// # Synctest (opt-in)
//
// Wrap the condition with [WithSynctestCollect] (or [WithSynctestCollectContext])
// to run the polling loop inside a [testing/synctest] bubble, which uses
// a fake clock. This eliminates timing-induced flakiness and makes the
// tick count deterministic. See [WithSynctest] for the constraints (no
// real I/O in the condition, requires [*testing.T]).
//
// # Examples
//
//	success: func(c *CollectT) { True(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//	failure: func(c *CollectT) { False(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//	failure: func(c *CollectT) { c.Cancel() }, 100*time.Millisecond, 20*time.Millisecond
func EventuallyWith[C CollectibleConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	return eventuallyWithT(t, condition, timeout, tick, msgAndArgs...)
}

func eventually[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	wantsBubble, cond := makeCondition(condition, false)
	p := newConditionPoller(pollOptions{
		mode:        pollUntilTrue,
		failMessage: "condition never satisfied",
	})

	return runPoller(t, p, cond, timeout, tick, wantsBubble, msgAndArgs...)
}

func never[C NeverConditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	wantsBubble, cond := makeCondition(condition, true)
	p := newConditionPoller(pollOptions{
		mode:        pollUntilTimeout,
		failMessage: "condition satisfied",
	})

	return runPoller(t, p, cond, timeout, tick, wantsBubble, msgAndArgs...)
}

func consistently[C Conditioner](t T, condition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	wantsBubble, cond := makeCondition(condition, false)
	p := newConditionPoller(pollOptions{
		mode:        pollUntilTimeout,
		failMessage: "condition failed once",
	})

	return runPoller(t, p, cond, timeout, tick, wantsBubble, msgAndArgs...)
}

func eventuallyWithT[C CollectibleConditioner](t T, collectCondition C, timeout time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var lastCollectedErrors []error
	var cancelFunc func() // will be set by pollCondition via onSetup
	wantsBubble, fn := makeCollectibleCondition(collectCondition)

	condition := func(ctx context.Context) (err error) {
		collector := new(CollectT).withCancelFunc(cancelFunc)

		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%w: %v", errConditionPanicked, r)
				collector.errors = append(collector.errors, err)
			}
			if collector.failed() {
				lastCollectedErrors = collector.collected()
				err = collector.last()
			}
		}()

		fn(ctx, collector)

		return nil
	}

	copyCollected := func(tt T) {
		for _, err := range lastCollectedErrors {
			tt.Errorf("%v", err)
		}
	}

	p := newConditionPoller(pollOptions{
		mode:        pollUntilTrue,
		failMessage: "condition never satisfied",
		onFailure:   copyCollected,
		onSetup:     func(cancel func()) { cancelFunc = cancel },
	})

	return runPoller(t, p, condition, timeout, tick, wantsBubble, msgAndArgs...)
}

// runPoller dispatches the polling to either the real-time or the
// [synctest] bubble-wrapped path, based on whether the condition opted into
// fake time AND the caller passed a concrete [*testing.T].
//
// When `wantsBubble` is true but `t` is not a `*testing.T` (e.g. a mock or
// [CollectT]), the call silently falls back to real-time polling. The
// synctest bubble requires a real `*testing.T`.
func runPoller(t T, p *conditionPoller, cond func(context.Context) error, timeout, tick time.Duration, wantsBubble bool, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	testingT, canBubble := t.(*testing.T)
	if !wantsBubble || !canBubble {
		return p.pollCondition(t, cond, timeout, tick, msgAndArgs...)
	}

	var result bool
	synctest.Test(testingT, func(inner *testing.T) {
		result = p.pollCondition(inner, cond, timeout, tick, msgAndArgs...)
	})

	return result
}

// makeCondition normalizes any variant from [Conditioner] or [NeverConditioner]
// into the unified `func(context.Context) error` form used by [pollCondition],
// and reports whether the caller opted into synctest-bubble polling.
//
// [WithSynctest] and [WithSynctestContext] are recognized as their underlying
// `func() bool` and `func(context.Context) error` forms with `wantsBubble = true`.
func makeCondition(condition any, reverse bool) (wantsBubble bool, cond func(context.Context) error) {
	switch typed := condition.(type) {
	case WithSynctest:
		_, cond = makeCondition((func() bool)(typed), reverse)
		return true, cond
	case WithSynctestContext:
		_, cond = makeCondition((func(context.Context) error)(typed), reverse)
		return true, cond
	case func() bool:
		if !reverse {
			return false, func(ctx context.Context) error {
				select {
				case <-ctx.Done():
					return ctx.Err()
				default:
					if res := typed(); !res {
						return errors.New("condition returned false")
					}

					return nil
				}
			}
		}

		// inverse bool <-> error logic for Never
		return false, func(ctx context.Context) error {
			select {
			case <-ctx.Done():
				return nil
			default:
				if res := typed(); res {
					return errors.New("condition returned true")
				}

				return nil
			}
		}
	case func(context.Context) error:
		// No reversal needed: the poller already uses err != nil as "condition happened".
		// For Eventually: err == nil = success. For Never: err != nil = failure.
		// Both align with the natural error semantics without inversion.
		return false, typed
	default: // unreachable
		panic(fmt.Errorf("unsupported Conditioner type. Mismatch with type constraint: %T", condition))
	}
}

// makeCollectibleCondition normalizes any [CollectibleConditioner] variant
// into the unified `func(context.Context, *CollectT)` form, and reports
// whether the caller opted into synctest-bubble polling.
func makeCollectibleCondition(condition any) (wantsBubble bool, fn func(context.Context, *CollectT)) {
	switch typed := condition.(type) {
	case WithSynctestCollect:
		_, fn = makeCollectibleCondition((func(*CollectT))(typed))
		return true, fn
	case WithSynctestCollectContext:
		_, fn = makeCollectibleCondition((func(context.Context, *CollectT))(typed))
		return true, fn
	case func(*CollectT):
		return false, func(ctx context.Context, collector *CollectT) {
			select {
			case <-ctx.Done():
				collector.Errorf("%v", ctx.Err())
			default:
				typed(collector)
			}
		}
	case func(context.Context, *CollectT):
		return false, typed
	default: // unreachable
		panic(fmt.Errorf("unsupported CollectibleConditioner type. Mismatch with type constraint: %T", condition))
	}
}

func recoverCondition(fn func(context.Context) error) func(context.Context) error {
	return func(ctx context.Context) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = fmt.Errorf("%w: %v", errConditionPanicked, r)
			}
		}()

		return fn(ctx)
	}
}

type conditionPoller struct {
	pollOptions

	ticker        *time.Ticker
	reported      atomic.Bool
	conditionChan chan func(context.Context) error
	doneChan      chan struct{}
}

func newConditionPoller(o pollOptions) *conditionPoller {
	return &conditionPoller{
		pollOptions: o,
	}
}

// initChannels creates the polling channels. MUST be called from inside
// [pollCondition] so that — when the caller activated a [synctest] bubble
// — the channels are bubble-owned. Receives on channels created outside
// the bubble do NOT count as durably blocking, which stalls the fake clock.
func (p *conditionPoller) initChannels() {
	p.conditionChan = make(chan func(context.Context) error, 1)
	p.doneChan = make(chan struct{})
}

// pollMode determines how the condition polling should behave.
type pollMode int

const (
	// pollUntilTrue succeeds when condition returns true (for Eventually).
	pollUntilTrue pollMode = iota
	// pollUntilTimeout succeeds when timeout is reached without condition being true (for Never/Consistently).
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
//
// It polls a condition function at regular intervals until success or timeout.
func (p *conditionPoller) pollCondition(t T, condition func(context.Context) error, timeout, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	parentCtx := p.parentContextFromT(t)
	ctx, cancel := p.cancellableContext(parentCtx, timeout)
	defer cancel()

	failFunc := p.failFunc(t, msgAndArgs...)

	// Allow caller to capture the cancel function (for eventuallyWithT's CollectT)
	if p.onSetup != nil {
		p.onSetup(cancel)
	}

	condition = recoverCondition(condition)

	// Channels and ticker MUST be created inside pollCondition so that,
	// when the caller activated a synctest bubble, they are bubble-owned
	// primitives. Channels created outside the bubble do not count as
	// durably blocking and would stall the fake clock.
	p.initChannels()

	p.ticker = time.NewTicker(tick)
	defer p.ticker.Stop()

	// Check the condition once first on the initial call.
	p.conditionChan <- condition

	var wg sync.WaitGroup

	// Goroutine 1: Poll for the condition at every tick
	wg.Add(1)
	go p.pollAtTickFunc(parentCtx, ctx, condition, failFunc, &wg)()

	// Goroutine 2: Execute the condition and check results
	wg.Add(1)
	go p.executeCondition(parentCtx, ctx, failFunc, &wg)()

	wg.Wait()

	// Determine success based on mode
	return p.determineOutcome(parentCtx, ctx, failFunc, t)()
}

func (p *conditionPoller) failFunc(t T, msgAndArgs ...any) func(string) {
	return func(reason string) {
		if p.reported.CompareAndSwap(false, true) {
			if reason != "" {
				t.Errorf("%s", reason)
			}
			Fail(t, p.failMessage, msgAndArgs...)
		}
	}
}

func (p *conditionPoller) pollAtTickFunc(parentCtx, ctx context.Context, condition func(context.Context) error, failFunc func(string), wg *sync.WaitGroup) func() {
	if p.mode == pollUntilTimeout {
		// For Never: check parent context separately
		return func() {
			defer wg.Done()

			for {
				select {
				case <-parentCtx.Done():
					failFunc(parentCtx.Err().Error())
					return
				case <-ctx.Done():
					return // timeout reached = success for Never
				case <-p.doneChan:
					return
				case <-p.ticker.C:
					// Nested select prevents blocking on channel send if context was cancelled
					// between receiving the tick and attempting to send the condition.
					select {
					case <-parentCtx.Done():
						failFunc(parentCtx.Err().Error())
						return
					case <-ctx.Done():
						return
					case <-p.doneChan:
						return
					case p.conditionChan <- condition:
					}
				}
			}
		}
	}

	// For Eventually: parent cancellation flows through ctx
	return func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				failFunc(ctx.Err().Error())
				return
			case <-p.doneChan:
				return
			case <-p.ticker.C:
				// Nested select prevents blocking on channel send if context was cancelled
				// between receiving the tick and attempting to send the condition.
				select {
				case <-ctx.Done():
					failFunc(ctx.Err().Error())
					return
				case <-p.doneChan:
					return
				case p.conditionChan <- condition:
				}
			}
		}
	}
}

func (p *conditionPoller) executeCondition(parentCtx, ctx context.Context, failFunc func(string), wg *sync.WaitGroup) func() {
	if p.mode == pollUntilTimeout {
		// For Never and Consistently
		return func() {
			defer wg.Done()

			for {
				select {
				case <-parentCtx.Done():
					failFunc(parentCtx.Err().Error())
					return
				case <-ctx.Done():
					return // timeout = success
				case fn := <-p.conditionChan:
					var conditionWg sync.WaitGroup
					conditionWg.Go(func() { // guards against the condition issue an early GoExit

						if err := fn(ctx); err != nil {
							close(p.doneChan) // (condition true <=> returns error) = failure for Never and Consistently
						}
					})
					conditionWg.Wait()

					select {
					case <-p.doneChan: // done: early exit
						return
					default:
					}
				}
			}
		}
	}

	// For Eventually
	return func() {
		defer wg.Done()

		for {
			select {
			case <-ctx.Done():
				failFunc(ctx.Err().Error())
				return
			case fn := <-p.conditionChan:
				var conditionWg sync.WaitGroup
				conditionWg.Go(func() { // guards against the condition issue an early GoExit

					if err := fn(ctx); err == nil {
						close(p.doneChan) // (condition true <=> err == nil) = success for Eventually
					}
				})
				conditionWg.Wait()

				select {
				case <-p.doneChan: // done: early exit
					return
				default:
				}
			}
		}
	}
}

func (p *conditionPoller) determineOutcome(parentCtx, ctx context.Context, failFunc func(string), t T) func() bool {
	if p.mode == pollUntilTimeout {
		return func() bool {
			select {
			case <-p.doneChan:
				// For Never: doneChan closed means condition became true
				// But if timeout was reached first (ctx.Err != nil), it's still a success.
				// This handles the race between timeout and condition becoming true.
				if ctx.Err() != nil {
					return true
				}
				// Condition became true before timeout = failure
				failFunc("")
				return false
			default:
				// doneChan not closed
				// For Never: timeout reached without condition being true = success
				// We should return a success, unless the parent context has failed.
				return parentCtx.Err() == nil
			}
		}
	}

	return func() bool {
		select {
		case <-p.doneChan:
			// For Eventually: doneChan closed means condition became true
			if ctx.Err() != nil {
				// Timeout occurred before or during success
				if p.onFailure != nil {
					p.onFailure(t)
				}
				return false
			}
			return true
		default:
			// doneChan not closed
			// opts.mode = pollUntilTrue
			// For Eventually: should not reach here (failFunc already called)
			if p.onFailure != nil {
				p.onFailure(t)
			}

			return false
		}
	}
}

func (p *conditionPoller) parentContextFromT(t T) context.Context {
	var parentCtx context.Context
	if withContext, ok := t.(contextualizer); ok {
		parentCtx = withContext.Context()
	}
	if parentCtx == nil {
		parentCtx = context.Background()
	}

	return parentCtx
}

func (p *conditionPoller) cancellableContext(parentCtx context.Context, timeout time.Duration) (context.Context, func()) {
	// For pollUntilTimeout (Never), we detach from parent cancellation
	// so that timeout reaching is a success, not a failure.
	var ctx context.Context
	var cancel context.CancelFunc
	if p.mode == pollUntilTimeout {
		ctx, cancel = context.WithTimeout(context.WithoutCancel(parentCtx), timeout)
	} else {
		ctx, cancel = context.WithTimeout(parentCtx, timeout)
	}

	return ctx, cancel
}

// Sentinel errors recorded by async condition assertions.
// Kept package-private: callers should rely on observable behavior, not on
// the marker shape. They are distinguishable so future tooling can tell apart
// "tick aborted by require", "user explicitly cancelled the assertion",
// and "condition panicked".
var (
	errFailNow           = errors.New("collect: failed now (tick aborted)")
	errCancelled         = errors.New("collect: cancelled (assertion aborted)")
	errConditionPanicked = errors.New("condition panicked")
)

// CollectT implements the [T] interface and collects all errors.
//
// [CollectT] is specifically intended to be used with [EventuallyWith] and
// should not be used outside of that context.
type CollectT struct {
	// Domain: condition
	//
	// Maintainer:
	//   1. FailNow() exits the current tick goroutine via runtime.Goexit (matching
	//      stretchr/testify semantics): require-style assertions abort the current
	//      evaluation and the poller retries on the next tick. It does NOT cancel
	//      the EventuallyWith context.
	//   2. Cancel() is the explicit escape hatch: it cancels the EventuallyWith
	//      context before exiting via runtime.Goexit, aborting the whole assertion.
	//   3. We no longer establish the distinction between c.errors nil or empty.
	//      Non-empty is an error, full stop.
	//   4. Deprecated methods have been removed.

	// A slice of errors. Non-empty slice denotes a failure.
	errors []error

	// cancelContext cancels the parent EventuallyWith context on Cancel().
	cancelContext func()
}

// Helper is like [testing.T.Helper] but does nothing.
func (*CollectT) Helper() {}

// Errorf collects the error.
func (c *CollectT) Errorf(format string, args ...any) {
	c.errors = append(c.errors, fmt.Errorf(format, args...))
}

// FailNow records a failure for the current tick and exits the condition
// goroutine via [runtime.Goexit].
//
// It does NOT cancel the [EventuallyWith] context: the poller will retry on
// the next tick. If a later tick succeeds, the assertion succeeds. If no tick
// ever succeeds before the timeout, the errors collected during the LAST tick
// (the one which most recently called FailNow) are reported on the parent t.
//
// To abort the whole assertion immediately, use [CollectT.Cancel].
func (c *CollectT) FailNow() {
	c.errors = append(c.errors, errFailNow)
	runtime.Goexit()
}

// Cancel records a failure, cancels the [EventuallyWith] context, then exits
// the condition goroutine via [runtime.Goexit].
//
// This aborts the whole assertion immediately, without waiting for the timeout.
// The errors collected during the cancelled tick are reported on the parent t.
//
// Use this when the condition can no longer be expected to succeed (e.g. an
// upstream resource has been observed in an unrecoverable state). For ordinary
// per-tick failures (e.g. "value not yet ready"), use [CollectT.FailNow]
// directly or transitively through [require] assertions.
func (c *CollectT) Cancel() {
	c.errors = append(c.errors, errCancelled)
	c.cancelContext()
	runtime.Goexit()
}

// Cancelf records a failure like [Cancel], with an additional custom message recorded.
func (c *CollectT) Cancelf(format string, msgAndArgs ...any) {
	c.errors = append(c.errors, fmt.Errorf(format, msgAndArgs...))
	c.Cancel()
}

func (c *CollectT) failed() bool {
	return len(c.errors) != 0
}

func (c *CollectT) collected() []error {
	return c.errors
}

func (c *CollectT) last() error {
	if len(c.errors) == 0 {
		return nil
	}

	return c.errors[len(c.errors)-1]
}

func (c *CollectT) withCancelFunc(cancel func()) *CollectT {
	c.cancelContext = cancel

	return c
}
