// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "context"

// T is an interface wrapper around [testing.T].
type T interface {
	Errorf(format string, args ...any)
}

// H is an interface for types that implement the Helper method.
// This allows marking functions as test helpers, e.g. [testing.T.Helper].
type H interface {
	Helper()
}

type (
	// ComparisonAssertionFunc is a common function prototype when comparing two values.  Can be useful
	// for table driven tests.
	ComparisonAssertionFunc func(T, any, any, ...any) bool

	// ValueAssertionFunc is a common function prototype when validating a single value.  Can be useful
	// for table driven tests.
	ValueAssertionFunc func(T, any, ...any) bool

	// BoolAssertionFunc is a common function prototype when validating a bool value.  Can be useful
	// for table driven tests.
	BoolAssertionFunc func(T, bool, ...any) bool

	// ErrorAssertionFunc is a common function prototype when validating an error value.  Can be useful
	// for table driven tests.
	ErrorAssertionFunc func(T, error, ...any) bool

	// Comparison is a custom function that returns true on success and false on failure.
	Comparison func() (success bool)
)

type failNower interface {
	FailNow()
}

type namer interface {
	Name() string
}

type contextualizer interface {
	Context() context.Context
}

type skipper interface {
	Skip(args ...any)
}
