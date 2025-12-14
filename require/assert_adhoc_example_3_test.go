// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package require_test

import (
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func ExampleBoolAssertionFunc() {
	t := new(testing.T) // normally provided by test

	isOkay := func(x int) bool {
		return x >= 42
	}

	for tt := range boolAssertionCases() {
		tt.requirement(t, isOkay(tt.arg))
	}

	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: true
}

type boolAssertionCase struct {
	name        string
	arg         int
	requirement require.BoolAssertionFunc
}

func boolAssertionCases() iter.Seq[boolAssertionCase] {
	return slices.Values([]boolAssertionCase{
		{"-1 is bad", -1, require.False},
		{"42 is good", 42, require.True},
		{"41 is bad", 41, require.False},
		{"45 is cool", 45, require.True},
	})
}
