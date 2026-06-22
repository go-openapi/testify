// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.26

package require

import (
	"github.com/go-openapi/testify/v2/internal/assertions"
)

// ErrorAsTypef is the same as [ErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func ErrorAsTypef[E error](t T, err error, target *E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.ErrorAsType[E](t, err, target, forwardArgs(msg, args)...) {
		return
	}

	t.FailNow()
}

// NotErrorAsTypef is the same as [NotErrorAsType], but it accepts a format string to format arguments like [fmt.Printf].
//
// Upon failure, the test [T] is marked as failed and stops execution.
func NotErrorAsTypef[E error](t T, err error, target *E, msg string, args ...any) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if assertions.NotErrorAsType[E](t, err, target, forwardArgs(msg, args)...) {
		return
	}

	t.FailNow()
}
