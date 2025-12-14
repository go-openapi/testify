// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assert_test

import (
	"encoding/json"
	"fmt"
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func ExampleValueAssertionFunc() {
	t := new(testing.T) // normally provided by test

	dumbParse := func(input string) any {
		var x any
		_ = json.Unmarshal([]byte(input), &x)
		return x
	}

	for tt := range valueAssertionCases() {
		tt.assertion(t, dumbParse(tt.arg))
	}

	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: true
}

type valueAssertionCase struct {
	name      string
	arg       string
	assertion assert.ValueAssertionFunc
}

func valueAssertionCases() iter.Seq[valueAssertionCase] {
	return slices.Values([]valueAssertionCase{
		{"true is not nil", "true", assert.NotNil},
		{"empty string is nil", "", assert.Nil},
		{"zero is not nil", "0", assert.NotNil},
		{"zero is zero", "0", assert.Zero},
		{"false is zero", "false", assert.Zero},
	})
}
