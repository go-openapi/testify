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

func ExampleComparisonAssertionFunc() {
	t := new(testing.T) // normally provided by test

	adder := func(x, y int) int {
		return x + y
	}

	for tt := range comparisonFuncCases() {
		tt.requirement(t, tt.expect, adder(tt.args.x, tt.args.y))
	}

	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: true
}

type args struct {
	x int
	y int
}

type comparisonFuncCase struct {
	name        string
	args        args
	expect      int
	requirement require.ComparisonAssertionFunc
}

func comparisonFuncCases() iter.Seq[comparisonFuncCase] {
	return slices.Values([]comparisonFuncCase{
		{"2+2=4", args{2, 2}, 4, require.Equal},
		{"2+2!=5", args{2, 2}, 5, require.NotEqual},
		{"2+3==5", args{2, 3}, 5, require.Exactly},
	})
}
