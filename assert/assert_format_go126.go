// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.26

package assert

import (
	"github.com/go-openapi/testify/v2/internal/assertions"
)

// ErrorAsTypef is the same as [ErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func ErrorAsTypef[E error](t T, err error, target *E, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.ErrorAsType[E](t, err, target, forwardArgs(msg, args)...)
}

// NotErrorAsTypef is the same as [NotErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and continues execution.
func NotErrorAsTypef[E error](t T, err error, target *E, msg string, args ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	return assertions.NotErrorAsType[E](t, err, target, forwardArgs(msg, args)...)
}
