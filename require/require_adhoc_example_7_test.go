// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package require_test

import (
	"fmt"
	"sync"

	"github.com/go-openapi/testify/v2/require"
)

func ExampleNoGoRoutineLeak_fail() {
	t := new(mockFailNowT) // normally provided by test
	// Since this test is failing and calls [runtime.Goexit], we need a mock to
	// avoid the example trigger a panick.

	blocker := make(chan struct{})
	var wg sync.WaitGroup

	defer func() {
		// clean resources _after_ the test
		close(blocker)
		wg.Wait()
	}()

	wg.Add(1)
	// This examplifies how a function that leaks a goroutine is detected.
	require.NoGoRoutineLeak(t, func() { // true when there is no leak
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks until cleanup
		}()
	})

	// Error message from test would typically return the leaked goroutine, e.g.:
	// #	0x69c8e8	github.com/go-openapi/testify/v2/assert_test.ExampleNoGoRoutineLeak.func2.1+0x48	.../assert_adhoc_example_7_test.go:30
	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: false
}

type mockFailNowT struct {
	failed bool
}

// Helper is like [testing.T.Helper] but does nothing.
func (mockFailNowT) Helper() {}

func (m *mockFailNowT) Errorf(format string, args ...any) {
	_ = format
	_ = args
}

func (m *mockFailNowT) FailNow() {
	m.failed = true
}

func (m *mockFailNowT) Failed() bool {
	return m.failed
}
