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

func ExampleValueAssertionFunc() {
	t := new(testing.T) // normally provided by test

	dumbParse := func(input string) any {
		var x any
		_ = json.Unmarshal([]byte(input), &x)
		return x
	}

	for tt := range valueAssertionCases() {
		tt.requirement(t, dumbParse(tt.arg))
	}

	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: true
}

type valueAssertionCase struct {
	name        string
	arg         string
	requirement require.ValueAssertionFunc
}

func valueAssertionCases() iter.Seq[valueAssertionCase] {
	return slices.Values([]valueAssertionCase{
		{"true is not nil", "true", require.NotNil},
		{"empty string is nil", "", require.Nil},
		{"zero is not nil", "0", require.NotNil},
		{"zero is zero", "0", require.Zero},
		{"false is zero", "false", require.Zero},
	})
}
