// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/v2/codegen@master [sha: bb2c19fba6c03f46cb643b3bcdc1d647ea1453ab]; DO NOT EDIT.

package assert

import (
	"github.com/go-openapi/testify/v2/internal/assertions"
)

var (
	// ErrTest is an error instance useful for testing.
	//
	// If the code does not care about error specifics, and only needs
	// to return the error for example, this error should be used to make
	// the test code more readable.
	ErrTest = assertions.ErrTest
)

type (
	// BoolAssertionFunc is a common function prototype when validating a bool value.  Can be useful
	// for table driven tests.
	BoolAssertionFunc = assertions.BoolAssertionFunc

	// CollectT implements the T interface and collects all errors.
	CollectT = assertions.CollectT

	// Deprecated: CompareType has only ever been for internal use and has accidentally been published since v1.6.0. Do not use it.
	CompareType = assertions.CompareType

	// Comparison is a custom function that returns true on success and false on failure.
	Comparison = assertions.Comparison

	// ComparisonAssertionFunc is a common function prototype when comparing two values.  Can be useful
	// for table driven tests.
	ComparisonAssertionFunc = assertions.ComparisonAssertionFunc

	// ErrorAssertionFunc is a common function prototype when validating an error value.  Can be useful
	// for table driven tests.
	ErrorAssertionFunc = assertions.ErrorAssertionFunc

	// H is an interface for types that implement the Helper method.
	// This allows marking functions as test helpers.
	H = assertions.H

	// PanicAssertionFunc is a common function prototype when validating a panic value.  Can be useful
	// for table driven tests.
	PanicAssertionFunc = assertions.PanicAssertionFunc

	// PanicTestFunc defines a func that should be passed to the assert.Panics and assert.NotPanics
	// methods, and represents a simple func that takes no arguments, and returns nothing.
	PanicTestFunc = assertions.PanicTestFunc

	// T is an interface wrapper around [testing.T].
	T = assertions.T

	// ValueAssertionFunc is a common function prototype when validating a single value.  Can be useful
	// for table driven tests.
	ValueAssertionFunc = assertions.ValueAssertionFunc
)

// Type declarations for backward compatibility.
type (
	// TestingT is like [T] and is declared here to remain compatible with previous versions of this package.
	//
	// Most users should not be affected, as the implementation of [T] that is widely used is [testing.T].
	//
	// Deprecated: use [T] as a more concise alternative.
	TestingT = T
)
