// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package require_test

import (
	"encoding/json"
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func ExampleErrorAssertionFunc() {
	t := new(testing.T) // normally provided by test

	dumbParseNum := func(input string, v any) error {
		return json.Unmarshal([]byte(input), v)
	}

	for tt := range errorAssertionCases() {
		var x float64
		tt.requirement(t, dumbParseNum(tt.arg, &x))
	}

	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: true
}

type errorAssertionCase struct {
	name        string
	arg         string
	requirement require.ErrorAssertionFunc
}

func errorAssertionCases() iter.Seq[errorAssertionCase] {
	return slices.Values([]errorAssertionCase{
		{"1.2 is number", "1.2", require.NoError},
		{"1.2.3 not number", "1.2.3", require.Error},
		{"true is not number", "true", require.Error},
		{"3 is number", "3", require.NoError},
	})
}
