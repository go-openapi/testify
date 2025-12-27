// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/v2/codegen; DO NOT EDIT.
// Generated on 2026-01-02 (version v1.2.2-760-g97c29e3) using codegen version master [sha: 97c29e3dbfc40800a080863ceea81db0cfd6e858]

package require

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
	BoolAssertionFunc func(T, bool, ...any)

	// CollectT implements the [T] interface and collects all errors.
	CollectT = assertions.CollectT

	// Comparison is a custom function that returns true on success and false on failure.
	Comparison = assertions.Comparison

	// ComparisonAssertionFunc is a common function prototype when comparing two values.  Can be useful
	// for table driven tests.
	ComparisonAssertionFunc func(T, any, any, ...any)

	// ErrorAssertionFunc is a common function prototype when validating an error value.  Can be useful
	// for table driven tests.
	ErrorAssertionFunc func(T, error, ...any)

	// H is an interface for types that implement the Helper method.
	// This allows marking functions as test helpers.
	H = assertions.H

	// PanicAssertionFunc is a common function prototype when validating a panic value.  Can be useful
	// for table driven tests.
	PanicAssertionFunc func(t T, f assertions.PanicTestFunc, msgAndArgs ...any)

	// PanicTestFunc defines a func that should be passed to the assert.Panics and assert.NotPanics
	// methods, and represents a simple func that takes no arguments, and returns nothing.
	PanicTestFunc = assertions.PanicTestFunc

	// T is an interface wrapper around [testing.T].
	T interface {
		assertions.T
		FailNow()
	}

	// ValueAssertionFunc is a common function prototype when validating a single value.  Can be useful
	// for table driven tests.
	ValueAssertionFunc func(T, any, ...any)
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
