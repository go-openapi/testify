// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "github.com/go-openapi/testify/v2/internal/leak"

type (
	LeakOption leakOption
	leakOption leak.Option
)

// NoGoRoutineLeak ensures that no goroutine did leak from inside the tested function.
//
// The function passed may apply optional filters to exclude known false positives (e.g. http, database connection...).
//
// # Usage
//
//	NoGoRoutineLeak(t, func() {
//	},
//	"should not leak any goroutine",
//	)
//
// # Examples
//
//   - success: NOT IMPLEMENTED
func NoGoRoutineLeak(t T, inside func(options ...LeakOption), msgAndArgs ...any) bool {
	// Domain: safety
	_ = inside // TODO
	if h, ok := t.(H); ok {
		h.Helper()
	}

	err := leak.Find()
	if err == nil {
		return true
	}

	return Fail(t, err.Error(), msgAndArgs...)
}
