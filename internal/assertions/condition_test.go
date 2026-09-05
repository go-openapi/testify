// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"testing"
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

// =======================================
// Test ConditionBlocked / ConditionNotBlocked
// =======================================

func TestConditionBlocked(t *testing.T) {
	t.Parallel()

	for tc := range blockedCases() {
		t.Run(tc.name, tc.test)
	}

	// Reflect-only inputs (cannot be expressed via the generic [BlockedT] constraint).
	t.Run("reflect-only/non-channel-string", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		shouldPassOrFail(t, mock, Blocked(mock, "not a channel"), false)
	})
	t.Run("reflect-only/non-channel-int", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		shouldPassOrFail(t, mock, Blocked(mock, 42), false)
	})
	t.Run("reflect-only/non-channel-slice", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		shouldPassOrFail(t, mock, Blocked(mock, []int{1, 2, 3}), false)
	})
	t.Run("reflect-only/nil-interface", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		shouldPassOrFail(t, mock, Blocked(mock, nil), false)
	})
	t.Run("reflect-only/recv-only-empty-blocks", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		ch := make(chan int)
		var ro <-chan int = ch
		shouldPassOrFail(t, mock, Blocked(mock, ro), true)
	})
	t.Run("reflect-only/recv-only-with-value-fails", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		ch := make(chan int, 1)
		ch <- 7
		var ro <-chan int = ch
		shouldPassOrFail(t, mock, Blocked(mock, ro), false)
	})
	// Send-only must be rejected cleanly: passing it to reflect.Select with
	// a Recv direction would otherwise panic at runtime.
	t.Run("reflect-only/send-only-rejected", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		ch := make(chan int)
		var so chan<- int = ch
		shouldPassOrFail(t, mock, Blocked(mock, so), false)
	})
}

func TestConditionNotBlocked(t *testing.T) {
	t.Parallel()

	for tc := range notBlockedCases() {
		t.Run(tc.name, tc.test)
	}

	// Multi-value buffer: NotBlocked consumes exactly one element per call.
	t.Run("buffered-multi-consumes-one", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		ch := make(chan int, 3)
		ch <- 1
		ch <- 2
		ch <- 3
		shouldPassOrFail(t, mock, NotBlocked(mock, ch), true)
		// Two values must remain.
		if got := len(ch); got != 2 {
			t.Errorf("expected 2 values left in channel, got %d", got)
		}
	})

	// Reflect-only inputs (cannot be expressed via the generic [NotBlockedT] constraint).
	t.Run("reflect-only/non-channel-string", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		shouldPassOrFail(t, mock, NotBlocked(mock, "not a channel"), false)
	})
	t.Run("reflect-only/non-channel-int", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		shouldPassOrFail(t, mock, NotBlocked(mock, 42), false)
	})
	t.Run("reflect-only/nil-interface", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		shouldPassOrFail(t, mock, NotBlocked(mock, nil), false)
	})
	t.Run("reflect-only/recv-only-with-value-passes", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		ch := make(chan int, 1)
		ch <- 7
		var ro <-chan int = ch
		shouldPassOrFail(t, mock, NotBlocked(mock, ro), true)
	})
	t.Run("reflect-only/recv-only-empty-fails", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		ch := make(chan int)
		var ro <-chan int = ch
		shouldPassOrFail(t, mock, NotBlocked(mock, ro), false)
	})
	t.Run("reflect-only/send-only-rejected", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		ch := make(chan int)
		var so chan<- int = ch
		shouldPassOrFail(t, mock, NotBlocked(mock, so), false)
	})
}

// blockedCases exercises both Blocked and BlockedT against the same input.
// Channels are produced by a factory: each subtest gets its own fresh channel
// so that the X variant cannot drain the buffer and silently turn the XT
// variant into a different scenario.
//
// The generic CHAN constraint of BlockedT is `~chan E`, so named bidirectional
// channel types are accepted, but recv-only / send-only channels must be tested
// against the reflect-based [Blocked] only — see the reflect-only subtests in
// [TestConditionBlocked].
func blockedCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// Empty channels block on receive.
		{"unbuffered-empty/blocks", testAllBlocked(func() chan int { return make(chan int) }, true)},
		{"buffered-empty/blocks", testAllBlocked(func() chan int { return make(chan int, 4) }, true)},
		{"unbuffered-struct/blocks", testAllBlocked(func() chan struct{} { return make(chan struct{}) }, true)},
		{"unbuffered-string/blocks", testAllBlocked(func() chan string { return make(chan string) }, true)},
		{"unbuffered-error/blocks", testAllBlocked(func() chan error { return make(chan error) }, true)},

		// Typed nil channels block forever (default branch fires => "blocked").
		{"typed-nil-int/blocks", testAllBlocked(func() chan int { return nil }, true)},
		{"typed-nil-struct/blocks", testAllBlocked(func() chan struct{} { return nil }, true)},

		// Non-empty channels do not block.
		{"buffered-with-value/fails", testAllBlocked(filledChanFactory(42), false)},
		{"buffered-with-string/fails", testAllBlocked(filledChanFactory("hello"), false)},
		{"buffered-with-zero-struct/fails", testAllBlocked(filledChanFactory(struct{}{}), false)},

		// Closed channels do NOT block (zero value is received immediately).
		// The Fail message currently shows the zero value, which is mildly
		// misleading; pinned here so the behavior is intentional.
		{"closed-int/fails", testAllBlocked(closedChanFactory[int](), false)},
		{"closed-struct/fails", testAllBlocked(closedChanFactory[struct{}](), false)},
	})
}

// notBlockedCases is the dual of blockedCases — same input shape, opposite expected outcome.
func notBlockedCases() iter.Seq[genericTestCase] {
	return slices.Values([]genericTestCase{
		// Empty channels block => NotBlocked fails.
		{"unbuffered-empty/fails", testAllNotBlocked(func() chan int { return make(chan int) }, false)},
		{"buffered-empty/fails", testAllNotBlocked(func() chan int { return make(chan int, 4) }, false)},
		{"unbuffered-struct/fails", testAllNotBlocked(func() chan struct{} { return make(chan struct{}) }, false)},
		{"unbuffered-string/fails", testAllNotBlocked(func() chan string { return make(chan string) }, false)},

		// Typed nil channels block forever => NotBlocked fails.
		{"typed-nil-int/fails", testAllNotBlocked(func() chan int { return nil }, false)},
		{"typed-nil-struct/fails", testAllNotBlocked(func() chan struct{} { return nil }, false)},

		// Non-empty channels do not block => NotBlocked passes.
		{"buffered-with-value/passes", testAllNotBlocked(filledChanFactory(42), true)},
		{"buffered-with-string/passes", testAllNotBlocked(filledChanFactory("hello"), true)},
		{"buffered-with-zero-struct/passes", testAllNotBlocked(filledChanFactory(struct{}{}), true)},

		// Closed channels are not blocked (zero value is received immediately).
		{"closed-int/passes", testAllNotBlocked(closedChanFactory[int](), true)},
		{"closed-struct/passes", testAllNotBlocked(closedChanFactory[struct{}](), true)},
	})
}

// testAllBlocked tests both Blocked and BlockedT against fresh channels
// produced by the factory.
func testAllBlocked[E any](newCh func() chan E, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		label := "should fail"
		if shouldPass {
			label = "should pass"
		}
		t.Run(label, func(t *testing.T) {
			t.Run("with Blocked", func(t *testing.T) {
				t.Parallel()
				mock := new(mockT)
				shouldPassOrFail(t, mock, Blocked(mock, newCh()), shouldPass)
			})
			t.Run("with BlockedT", func(t *testing.T) {
				t.Parallel()
				mock := new(mockT)
				shouldPassOrFail(t, mock, BlockedT(mock, newCh()), shouldPass)
			})
		})
	}
}

func testAllNotBlocked[E any](newCh func() chan E, shouldPass bool) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		label := "should fail"
		if shouldPass {
			label = "should pass"
		}
		t.Run(label, func(t *testing.T) {
			t.Run("with NotBlocked", func(t *testing.T) {
				t.Parallel()
				mock := new(mockT)
				shouldPassOrFail(t, mock, NotBlocked(mock, newCh()), shouldPass)
			})
			t.Run("with NotBlockedT", func(t *testing.T) {
				t.Parallel()
				mock := new(mockT)
				shouldPassOrFail(t, mock, NotBlockedT(mock, newCh()), shouldPass)
			})
		})
	}
}

func filledChanFactory[E any](v E) func() chan E {
	return func() chan E {
		ch := make(chan E, 1)
		ch <- v
		return ch
	}
}

func closedChanFactory[E any]() func() chan E {
	return func() chan E {
		ch := make(chan E)
		close(ch)
		return ch
	}
}
