// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package require_test

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/go-openapi/testify/v2/require"
)

// ExampleWithSynctest_asyncReady demonstrates opting into [testing/synctest]
// bubble polling via [require.WithSynctest]. Time operations inside the bubble
// use a fake clock — a 1-hour timeout with a 1-minute tick completes in
// microseconds of real wall-clock time while remaining deterministic.
//
// Prefer this wrapper when the condition is pure compute or uses [time.Sleep]
// internally. See [require.WithSynctest] for the constraints (no real I/O, no
// external goroutines driving state change).
func ExampleEventually_withSyncTest() {
	t := new(testing.T) // normally provided by test

	// A counter that converges on the 5th poll — no external time pressure.
	var attempts atomic.Int32
	cond := func() bool {
		return attempts.Add(1) == 5
	}

	// 1-hour/1-minute: under fake time this is instantaneous and
	// deterministic — exactly 5 calls to the condition.
	require.Eventually(t, require.WithSynctest(cond), 1*time.Hour, 1*time.Minute)

	fmt.Printf("ready: %t, attempts: %d", !t.Failed(), attempts.Load())

	// Output: ready: true, attempts: 5
}

// ExampleWithSynctestContext_healthCheck demonstrates the context/error
// variant of the synctest opt-in. [require.WithSynctestContext] wraps a
// [func(context.Context) error] condition for fake-time polling.
func ExampleEventually_withContext() {
	t := new(testing.T) // normally provided by test

	var attempts atomic.Int32
	healthCheck := func(_ context.Context) error {
		if attempts.Add(1) < 3 {
			return errors.New("service not ready")
		}

		return nil
	}

	require.Eventually(t, require.WithSynctestContext(healthCheck), 1*time.Hour, 1*time.Minute)

	fmt.Printf("healthy: %t, attempts: %d", !t.Failed(), attempts.Load())

	// Output: healthy: true, attempts: 3
}

// ExampleWithSynctest_never demonstrates [require.Never] with the synctest
// opt-in. The condition is polled over the fake-time window without costing
// real wall-clock time.
func ExampleNever_withSyncTest() {
	t := new(testing.T) // normally provided by test

	// A flag that should remain false across the whole observation period.
	var flipped atomic.Bool
	require.Never(t, require.WithSynctest(flipped.Load), 1*time.Hour, 1*time.Minute)

	fmt.Printf("never flipped: %t", !t.Failed())

	// Output: never flipped: true
}

// ExampleWithSynctest_consistently demonstrates [require.Consistently] with
// the synctest opt-in — asserting an invariant holds across the entire
// observation window under deterministic fake time.
func ExampleConsistently_withSynctest() {
	t := new(testing.T) // normally provided by test

	// An invariant that must hold throughout the observation period.
	var counter atomic.Int32
	counter.Store(5)
	invariant := func() bool { return counter.Load() < 10 }

	require.Consistently(t, require.WithSynctest(invariant), 1*time.Hour, 1*time.Minute)

	fmt.Printf("invariant held: %t", !t.Failed())

	// Output: invariant held: true
}

// ExampleWithSynctestCollect_convergence demonstrates [require.EventuallyWith]
// with [require.WithSynctestCollect] — a [CollectT]-based condition polled
// inside a synctest bubble. Useful when the condition uses several require /
// assert calls and you want deterministic retry behavior.
func ExampleEventuallyWith_withSynctest() {
	t := new(testing.T) // normally provided by test

	var attempts atomic.Int32
	cond := func(c *require.CollectT) {
		n := attempts.Add(1)
		require.Equal(c, int32(3), n, "not yet converged")
	}

	require.EventuallyWith(t, require.WithSynctestCollect(cond), 1*time.Hour, 1*time.Minute)

	fmt.Printf("converged: %t, attempts: %d", !t.Failed(), attempts.Load())

	// Output: converged: true, attempts: 3
}
