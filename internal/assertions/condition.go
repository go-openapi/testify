// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"runtime"
	"time"
)

// Condition uses a Comparison to assert a complex condition.
//
// Examples:
//
//	success:  func() bool { return true }
//	failure:  func() bool { return false }
func Condition(t T, comp Comparison, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	result := comp()
	if !result {
		Fail(t, "Condition failed!", msgAndArgs...)
	}
	return result
}

// Eventually asserts that given condition will be met in waitFor time,
// periodically checking target function each tick.
//
//	assert.Eventually(t, func() bool { return true; }, time.Second, 10*time.Millisecond)
//
// Examples:
//
//	success:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
func Eventually(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	ch := make(chan bool, 1)
	checkCond := func() { ch <- condition() }

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	var tickC <-chan time.Time

	// Check the condition once first on the initial call.
	go checkCond()

	for {
		select {
		case <-timer.C:
			return Fail(t, "Condition never satisfied", msgAndArgs...)
		case <-tickC:
			tickC = nil
			go checkCond()
		case v := <-ch:
			if v {
				return true
			}
			tickC = ticker.C
		}
	}
}

// CollectT implements the T interface and collects all errors.
type CollectT struct {
	// A slice of errors. Non-nil slice denotes a failure.
	// If it's non-nil but len(c.errors) == 0, this is also a failure
	// obtained by direct c.FailNow() call.
	errors []error
}

// Helper is like [testing.T.Helper] but does nothing.
func (CollectT) Helper() {}

// Errorf collects the error.
func (c *CollectT) Errorf(format string, args ...any) {
	c.errors = append(c.errors, fmt.Errorf(format, args...))
}

// FailNow stops execution by calling runtime.Goexit.
func (c *CollectT) FailNow() {
	c.fail()
	runtime.Goexit()
}

// Deprecated: That was a method for internal usage that should not have been published. Now just panics.
func (*CollectT) Reset() {
	panic("Reset() is deprecated")
}

// Deprecated: That was a method for internal usage that should not have been published. Now just panics.
func (*CollectT) Copy(T) {
	panic("Copy() is deprecated")
}

func (c *CollectT) fail() {
	if !c.failed() {
		c.errors = []error{} // Make it non-nil to mark a failure.
	}
}

func (c *CollectT) failed() bool {
	return c.errors != nil
}

// EventuallyWithT asserts that given condition will be met in waitFor time,
// periodically checking target function each tick. In contrast to Eventually,
// it supplies a CollectT to the condition function, so that the condition
// function can use the CollectT to call other assertions.
// The condition is considered "met" if no errors are raised in a tick.
// The supplied CollectT collects all errors from one tick (if there are any).
// If the condition is not met before waitFor, the collected errors of
// the last tick are copied to t.
//
//	externalValue := false
//	go func() {
//		time.Sleep(8*time.Second)
//		externalValue = true
//	}()
//	assert.EventuallyWithT(t, func(c *assert.CollectT) {
//		// add assertions as needed; any assertion failure will fail the current tick
//		assert.True(c, externalValue, "expected 'externalValue' to be true")
//	}, 10*time.Second, 1*time.Second, "external state has not changed to 'true'; still false")
//
// Examples:
//
//	success: func(c *CollectT) { True(c,true) }, 100*time.Millisecond, 20*time.Millisecond
//	failure: func(c *CollectT) { False(c,true) }, 100*time.Millisecond, 20*time.Millisecond
func EventuallyWithT(t T, condition func(collect *CollectT), waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var lastFinishedTickErrs []error
	ch := make(chan *CollectT, 1)

	checkCond := func() {
		collect := new(CollectT)
		defer func() {
			ch <- collect
		}()
		condition(collect)
	}

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	var tickC <-chan time.Time

	// Check the condition once first on the initial call.
	go checkCond()

	for {
		select {
		case <-timer.C:
			for _, err := range lastFinishedTickErrs {
				t.Errorf("%v", err)
			}
			return Fail(t, "Condition never satisfied", msgAndArgs...)
		case <-tickC:
			tickC = nil
			go checkCond()
		case collect := <-ch:
			if !collect.failed() {
				return true
			}
			// Keep the errors from the last ended condition, so that they can be copied to t if timeout is reached.
			lastFinishedTickErrs = collect.errors
			tickC = ticker.C
		}
	}
}

// Never asserts that the given condition doesn't satisfy in waitFor time,
// periodically checking the target function each tick.
//
//	assert.Never(t, func() bool { return false; }, time.Second, 10*time.Millisecond)
//
// Examples:
//
//	success:  func() bool { return false }, 100*time.Millisecond, 20*time.Millisecond
//	failure:  func() bool { return true }, 100*time.Millisecond, 20*time.Millisecond
func Never(t T, condition func() bool, waitFor time.Duration, tick time.Duration, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	ch := make(chan bool, 1)
	checkCond := func() { ch <- condition() }

	timer := time.NewTimer(waitFor)
	defer timer.Stop()

	ticker := time.NewTicker(tick)
	defer ticker.Stop()

	var tickC <-chan time.Time

	// Check the condition once first on the initial call.
	go checkCond()

	for {
		select {
		case <-timer.C:
			return true
		case <-tickC:
			tickC = nil
			go checkCond()
		case v := <-ch:
			if v {
				return Fail(t, "Condition satisfied", msgAndArgs...)
			}
			tickC = ticker.C
		}
	}
}
