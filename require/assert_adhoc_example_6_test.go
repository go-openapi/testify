// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package require_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/require"
)

func ExampleAssertions_with_generics() {
	t := new(testing.T) // normally provided by test

	a := require.New(t)

	const expected = "hello"
	goodValue := "hello"

	a.Equal(expected, goodValue) // classic reflect-based assertion
	fmt.Println("good value")

	require.EqualT(a.T(), expected, goodValue) // usage with generic assertion
	fmt.Println("good value")

	// Output:
	// good value
	// good value
}
