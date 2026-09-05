// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"reflect"
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

// Blocked asserts that a channel is blocked on receive.
//
// It always fails if the operand is not a channel, or if the channel is send-only.
//
// # Usage
//
//	ch := make(chan struct{})
//	assertions.Blocked(t, ch)
//
// # Examples
//
//	success:  make(chan struct{})
//	failure:  sendChanMessage()
func Blocked(t T, ch any, msgAndArgs ...any) bool {
	// Domain: condition
	// Opposite: NotBlocked
	if h, ok := t.(H); ok {
		h.Helper()
	}

	chanType := reflect.TypeOf(ch)
	if chanType == nil || chanType.Kind() != reflect.Chan {
		return Fail(t, fmt.Sprintf("Expected a channel but got: %T", ch), msgAndArgs...)
	}
	if chanType.ChanDir()&reflect.RecvDir == 0 {
		return Fail(t, "Expected channel direction to allow receive", msgAndArgs...)
	}

	chanValue := reflect.ValueOf(ch)
	chosen, recv, ok := reflect.Select([]reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: chanValue,
		},
		{
			Dir: reflect.SelectDefault,
		},
	})

	if chosen == 0 {
		if ok {
			return Fail(t, fmt.Sprintf("Channel receive should have blocked, but got: %v", recv), msgAndArgs...)
		}

		return Fail(t, "Channel receive should have blocked, but channel was closed", msgAndArgs...)
	}

	return true
}

// BlockedT asserts that a channel is blocked on receive.
//
// # Usage
//
//	ch := make(chan struct{})
//	assertions.BlockedT(t, ch)
//
// # Examples
//
//	success:  make(chan struct{})
//	failure:  sendChanMessage()
func BlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool {
	// Domain: condition
	// Opposite: NotBlockedT
	if h, ok := t.(H); ok {
		h.Helper()
	}

	rch := (<-chan E)(ch) // impose compile-time check that the channel is indeed open for receive
	select {
	case r, ok := <-rch:
		if ok {
			return Fail(t, fmt.Sprintf("Channel receive should have blocked, but got: %v", r), msgAndArgs...)
		}

		return Fail(t, "Channel receive should have blocked, but channel was closed", msgAndArgs...)
	default:
		return true
	}
}

// NotBlocked asserts that a channel is not blocked on receive.
//
// It always fails if the operand is not a channel, or if the channel is send-only.
//
// A closed channel doesn't block and returns true.
// Notice that this consumes any message available in the channel.
//
// # Usage
//
//	ch := make(chan struct{}, 1)
//	ch <- struct{}{}
//	assertions.NotBlocked(t, ch)
//
// # Examples
//
//	success:  sendChanMessage()
//	failure:  make(chan struct{})
func NotBlocked(t T, ch any, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	chanType := reflect.TypeOf(ch)
	if chanType == nil || chanType.Kind() != reflect.Chan {
		return Fail(t, fmt.Sprintf("Expected a channel but got: %T", ch), msgAndArgs...)
	}
	if chanType.ChanDir()&reflect.RecvDir == 0 {
		return Fail(t, "Expected channel direction to allow receive", msgAndArgs...)
	}

	chanValue := reflect.ValueOf(ch)
	chosen, _, _ := reflect.Select([]reflect.SelectCase{
		{
			Dir:  reflect.SelectRecv,
			Chan: chanValue,
		},
		{
			Dir: reflect.SelectDefault,
		},
	})

	if chosen == 0 {
		return true
	}

	return Fail(t, "Channel receive should not have blocked", msgAndArgs...)
}

// NotBlockedT asserts that a channel is not blocked on receive.
//
// A closed channel doesn't block and returns true.
// Notice that this consumes any message available in the channel.
//
// # Usage
//
//	ch := make(chan struct{}, 1)
//	ch <- struct{}{}
//	assertions.NotBlockedT(t, ch)
//
// # Examples
//
//	success:  sendChanMessage()
//	failure:  make(chan struct{})
func NotBlockedT[E any, CHAN ~chan E](t T, ch CHAN, msgAndArgs ...any) bool {
	// Domain: condition
	if h, ok := t.(H); ok {
		h.Helper()
	}

	rch := (<-chan E)(ch) // impose compile-time check that the channel is indeed open for receive
	select {
	case <-rch:
		return true
	default:
		return Fail(t, "Channel receive should not have blocked", msgAndArgs...)
	}
}
