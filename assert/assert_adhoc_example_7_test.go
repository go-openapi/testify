// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assert_test

import (
	"fmt"
	"sync"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func ExampleNoGoRoutineLeak_fail() {
	t := new(testing.T) // normally provided by test

	blocker := make(chan struct{})
	var wg sync.WaitGroup

	defer func() {
		// clean resources _after_ the test
		close(blocker)
		wg.Wait()
	}()

	wg.Add(1)
	// This examplifies how a function that leaks a goroutine is detected.
	result := assert.NoGoRoutineLeak(t, func() { // true when there is no leak
		go func() {
			defer wg.Done()
			<-blocker // leaked: blocks until cleanup
		}()
	})

	// Error message from test would typically return the leaked goroutine, e.g.:
	// #	0x69c8e8	github.com/go-openapi/testify/v2/assert_test.ExampleNoGoRoutineLeak.func2.1+0x48	.../assert_adhoc_example_7_test.go:30
	fmt.Printf("has leak: %t", !result)
	// Output:
	// has leak: true
}
