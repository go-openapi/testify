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

func ExampleErrorAssertionFunc() {
	t := new(testing.T) // normally provided by test

	dumbParseNum := func(input string, v any) error {
		return json.Unmarshal([]byte(input), v)
	}

	for tt := range errorAssertionCases() {
		var x float64
		tt.assertion(t, dumbParseNum(tt.arg, &x))
	}

	fmt.Printf("passed: %t", !t.Failed())

	// Output: passed: true
}

type errorAssertionCase struct {
	name      string
	arg       string
	assertion assert.ErrorAssertionFunc
}

func errorAssertionCases() iter.Seq[errorAssertionCase] {
	return slices.Values([]errorAssertionCase{
		{"1.2 is number", "1.2", assert.NoError},
		{"1.2.3 not number", "1.2.3", assert.Error},
		{"true is not number", "true", assert.Error},
		{"3 is number", "3", assert.NoError},
	})
}
