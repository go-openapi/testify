//go:build integrationtest

// Package examplespkg is a test fixture for the examples-parser unit tests.
package examplespkg

import "fmt"

// Greet returns a greeting for the given name.
func Greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// Add returns the sum of two integers.
func Add(a, b int) int {
	return a + b
}

// NoExample is exported but has no testable example.
func NoExample() {}

func unexported() {} //nolint:unused // fixture: verifies unexported symbols are skipped

// Formatter is an exported type with a whole-file example.
type Formatter struct {
	Prefix string
}
