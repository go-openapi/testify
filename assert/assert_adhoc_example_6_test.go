// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assert_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/testify/v2/assert"
)

func ExampleAssertions_with_generics() {
	t := new(testing.T) // normally provided by test

	a := assert.New(t)

	const expected = "hello"
	goodValue := "hello"
	wrongValue := "world"

	r0 := a.Equal(expected, goodValue) // classic reflect-based assertion
	fmt.Printf("good value is %t\n", r0)

	r1 := assert.EqualT(a.T(), expected, goodValue) // usage with generic assertion
	fmt.Printf("good value is %t\n", r1)

	r2 := assert.EqualT(a.T(), expected, wrongValue)
	fmt.Printf("wrong value is %t\n", r2)

	// Output:
	// good value is true
	// good value is true
	// wrong value is false
}
