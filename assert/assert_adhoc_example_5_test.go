// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assert_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func ExamplePanicAssertionFunc() {
	t := new(testing.T) // normally provided by test

	for tt := range panicAssertionCases() {
		tt.assertion(t, tt.panicFn)
	}

	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: true
}

type panicAssertionCase struct {
	name      string
	panicFn   func()
	assertion assert.PanicAssertionFunc
}

func panicAssertionCases() iter.Seq[panicAssertionCase] {
	return slices.Values([]panicAssertionCase{
		{"with panic", func() { panic(nil) }, assert.Panics},
		{"without panic", func() {}, assert.NotPanics},
	})
}
