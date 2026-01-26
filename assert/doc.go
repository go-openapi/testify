// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package assert provides a set of testing tools for use with the standard Go testing system.
//
// # Note
//
// All functions in this package return a bool value indicating whether the assertion has passed.
//
// # Example usage
//
// The following is a complete example using assert in a standard test function:
//
//	import (
//	  "testing"
//	  "github.com/go-openapi/testify/v2/assert"
//	)
//
//	func TestSomething(t *testing.T) {
//
//	  var a string = "Hello"
//	  var b string = "Hello"
//
//	  assert.Equal(t, a, b, "The two words should be the same.")
//
//	}
//
// if you assert many times, use the format below:
//
//	import (
//	  "testing"
//	  "github.com/go-openapi/testify/v2/assert"
//	)
//
//	func TestSomething(t *testing.T) {
//	  assert := assert.New(t)
//
//	  var a string = "Hello"
//	  var b string = "Hello"
//
//	  assert.Equal(a, b, "The two words should be the same.")
//	}
//
// # Assertions
//
// Assertions allow you to easily write test code.
//
// All assertion functions take as the first argument, the [*testing.T] object provided by the
// standard testing framework.
//
// This allows the assertion functions to write the failings and other details to the correct place.
//
// Every assertion function also takes an optional string message as the final argument,
// allowing custom error messages to be appended to the message the assertion method outputs.
//
// See [our doc site](https://go-openapi.github.io/testify/) for usage and examples and
// [go docs](https://pkg.go/dev/go-openapi/testify) for complete reference.
package assert
