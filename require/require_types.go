// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.
// Generated on 2026-01-20 (version 74d5686) using codegen version v2.1.9-0.20260119232631-74d5686313f0+dirty [sha: 74d5686313f0820ae0e2758b95d598f646cd7ad5]

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

	// Boolean is a bool or any type that can be converted to a bool.
	Boolean = assertions.Boolean

	// CollectT implements the [T] interface and collects all errors.
	//
	// [CollectT] is specifically intended to be used with [EventuallyWithT] and
	// should not be used outside of that context.
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
	// This allows marking functions as test helpers, e.g. [testing.T.Helper].
	H = assertions.H

	// Measurable is any number for which we can compute a delta (floats or integers).
	//
	// This is used by [InDeltaT] and [InEpsilonT].
	//
	// NOTE: unfortunately complex64 and complex128 are not supported.
	Measurable = assertions.Measurable

	// Ordered is a standard ordered type (i.e. types that support "<": [cmp.Ordered]) plus []byte and [time.Time].
	//
	// This is used by [GreaterT], [GreaterOrEqualT], [LessT], [LessOrEqualT], [IsIncreasingT], [IsDecreasingT].
	//
	// NOTE: since [time.Time] is a struct, custom types which redeclare [time.Time] are not supported.
	Ordered = assertions.Ordered

	// PanicAssertionFunc is a common function prototype when validating a panic value.  Can be useful
	// for table driven tests.
	PanicAssertionFunc func(t T, f assertions.PanicTestFunc, msgAndArgs ...any)

	// PanicTestFunc defines a func that should be passed to the assert.Panics and assert.NotPanics
	// methods, and represents a simple func that takes no arguments, and returns nothing.
	PanicTestFunc = assertions.PanicTestFunc

	// RegExp is either a text containing a regular expression to compile (string or []byte), or directly the compiled regexp.
	//
	// This is used by [RegexpT] and [NotRegexpT].
	RegExp = assertions.RegExp

	// SignedNumeric is a signed integer or a floating point number or any type that can be converted to one of these.
	SignedNumeric = assertions.SignedNumeric

	// T is an interface wrapper around [testing.T].
	T interface {
		assertions.T
		FailNow()
	}

	// Text is any type of underlying type string or []byte.
	//
	// This is used by [RegexpT], [NotRegexpT], [JSONEqT], and [YAMLEqT].
	//
	// NOTE: unfortunately, []rune is not supported.
	Text = assertions.Text

	// UnsignedNumeric is an unsigned integer.
	//
	// NOTE: there are no unsigned floating point numbers.
	UnsignedNumeric = assertions.UnsignedNumeric

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
