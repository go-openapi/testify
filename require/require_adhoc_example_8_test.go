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

// ExampleEventually_asyncReady demonstrates polling a condition that becomes true
// after a few attempts, simulating an asynchronous operation completing.
func ExampleEventually_asyncReady() {
	t := new(testing.T) // normally provided by test

	// Simulate an async operation that completes after a short delay.
	var ready atomic.Bool
	go func() {
		time.Sleep(30 * time.Millisecond)
		ready.Store(true)
	}()

	require.Eventually(t, ready.Load, 200*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("eventually ready: %t", !t.Failed())

	// Output: eventually ready: true
}

// ExampleEventually_healthCheck demonstrates [Eventually] with a
// func(context.Context) error condition, polling until the operation
// succeeds (returns nil).
func ExampleEventually_healthCheck() {
	t := new(testing.T) // normally provided by test

	// Simulate a service that becomes healthy after a few attempts.
	var attempts atomic.Int32
	healthCheck := func(_ context.Context) error {
		if attempts.Add(1) < 3 {
			return errors.New("service not ready")
		}

		return nil
	}

	require.Eventually(t, healthCheck, 200*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("eventually healthy: %t", !t.Failed())

	// Output: eventually healthy: true
}

// ExampleNever_noSpuriousEvents demonstrates asserting that a condition never becomes true
// during the observation period.
func ExampleNever_noSpuriousEvents() {
	t := new(testing.T) // normally provided by test

	// A channel that should remain empty during the test.
	events := make(chan struct{}, 1)

	require.Never(t, func() bool {
		select {
		case <-events:
			return true // event received = condition becomes true = Never fails
		default:
			return false
		}
	}, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("never received: %t", !t.Failed())

	// Output: never received: true
}

// ExampleConsistently_invariant demonstrates asserting that a condition remains true
// throughout the entire observation period.
func ExampleConsistently_invariant() {
	t := new(testing.T) // normally provided by test

	// A counter that stays within bounds during the test.
	var counter atomic.Int32
	counter.Store(5)

	require.Consistently(t, func() bool {
		return counter.Load() < 10
	}, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("consistently under limit: %t", !t.Failed())

	// Output: consistently under limit: true
}

// ExampleConsistently_alwaysHealthy demonstrates [Consistently] with a
// func(context.Context) error condition, asserting that the operation
// always succeeds (returns nil) throughout the observation period.
func ExampleConsistently_alwaysHealthy() {
	t := new(testing.T) // normally provided by test

	// Simulate a service that stays healthy.
	healthCheck := func(_ context.Context) error {
		return nil // always healthy
	}

	require.Consistently(t, healthCheck, 100*time.Millisecond, 10*time.Millisecond)

	fmt.Printf("consistently healthy: %t", !t.Failed())

	// Output: consistently healthy: true
}
