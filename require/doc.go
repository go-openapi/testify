// Package require implements the same assertions as the assert package but
// stops test execution when a test fails.
//
// # Example usage
//
// The following is a complete example using require in a standard test function:
//
//	import (
//	  "testing"
//	  "github.com/go-openapi/testify/v2/require"
//	)
//
//	func TestSomething(t *testing.T) {
//	  var a string = "Hello"
//	  var b string = "Hello"
//
//	  require.Equal(t, a, b, "The two words should be the same.")
//	}
//
// # Assertions
//
// The require package exposes the same functions as the assert package,
// but instead of returning a boolean result the functions call [testing.T.FailNow].
//
// A consequence of this is that it must be called from the goroutine running
// the test function, not from other goroutines created during the test.
//
// Every assertion function also takes an optional string message as the final argument,
// allowing custom error messages to be appended to the message the assertion method outputs.
//
// See [our doc site](https://go-openapi.github.io/testify/) for usage and examples and
// [go docs](https://pkg.go/dev/go-openapi/testify) for complete reference.
package require
