// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build go1.26

package assertions

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrorAsType asserts that at least one of the errors in err's chain is of type E.
//
// It is the type-safe counterpart of [ErrorAs], built on the go1.26 [errors.AsType]:
// the expected type is the type parameter E (checked at compile time, no reflection),
// rather than the untyped any target used by [ErrorAs].
//
// target receives the matched error when the assertion succeeds. It may be nil, for
// callers that only want to know whether the chain holds an error of type E: in that
// case E cannot be inferred and must be supplied explicitly.
//
// This assertion requires go1.26 or newer; it is unavailable on older toolchains.
//
// # Usage
//
//	// capture the matched error (E is inferred from target):
//	var target *MyError
//	assertions.ErrorAsType(t, err, &target)
//
//	// only check, discarding the value (E given explicitly):
//	assertions.ErrorAsType[*MyError](t, err, nil)
//
// # Examples
//
//	success: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
//	failure: ErrTest, new(*dummyError)
func ErrorAsType[E error](t T, err error, target *E, msgAndArgs ...any) bool {
	// Domain: error
	// Opposite: NotErrorAsType
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if found, ok := errors.AsType[E](err); ok {
		if target != nil {
			*target = found
		}
		return true
	}

	expectedType := reflect.TypeFor[E]().String()
	if err == nil {
		return Fail(t, fmt.Sprintf("An error is expected but got nil.\n"+
			"expected: %s", expectedType), msgAndArgs...)
	}

	chain := buildErrorChainString(err, true)

	return Fail(t, fmt.Sprintf("Should be in error chain:\n"+
		"expected: %s\n"+
		"in chain: %s", expectedType, truncatingFormat("%s", chain),
	), msgAndArgs...)
}

// NotErrorAsType asserts that none of the errors in err's chain is of type E.
//
// It is the type-safe counterpart of [NotErrorAs], built on the go1.26 [errors.AsType].
//
// target is only used to infer the type parameter E and is never assigned; it may be nil,
// in which case E must be supplied explicitly.
//
// This assertion requires go1.26 or newer; it is unavailable on older toolchains.
//
// # Usage
//
//	var target *MyError
//	assertions.NotErrorAsType(t, err, &target)
//
//	// or, supplying E explicitly:
//	assertions.NotErrorAsType[*MyError](t, err, nil)
//
// # Examples
//
//	success: ErrTest, new(*dummyError)
//	failure: fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError)
func NotErrorAsType[E error](t T, err error, target *E, msgAndArgs ...any) bool {
	// Domain: error
	if h, ok := t.(H); ok {
		h.Helper()
	}
	_ = target // present only so E can be inferred from a typed argument; never assigned
	if _, ok := errors.AsType[E](err); !ok {
		return true
	}

	chain := buildErrorChainString(err, true)

	return Fail(t, fmt.Sprintf("Target error should not be in err chain:\n"+
		"found: %s\n"+
		"in chain: %s", reflect.TypeFor[E]().String(), truncatingFormat("%s", chain),
	), msgAndArgs...)
}
