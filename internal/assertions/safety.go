// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"

	"github.com/go-openapi/testify/v2/internal/leak"
)

// NoGoRoutineLeak ensures that no goroutine did leak from inside the tested function.
//
// NOTE: only the go routines spawned from inside the tested function are checked for leaks.
// No filter or configuration is needed to exclude "known go routines".
//
// Resource cleanup should be done inside the tested function, and not using [testing.T.Cleanup],
// as t.Cleanup is called after the leak check.
//
// # Edge cases
//
//   - if the tested function panics leaving behind leaked goroutines, these are detected.
//   - if the tested function calls runtime.Goexit (e.g. from [testing.T.FailNow]) leaving behind leaked goroutines,
//     these are detected.
//   - if a panic occurs in one of the leaked go routines, it cannot be recovered with certainty and
//     the calling program will usually panic.
//
// # Concurrency
//
// [NoGoRoutineLeak] may be used safely in parallel tests.
//
// # Usage
//
//		NoGoRoutineLeak(t, func() {
//			...
//		},
//	 nil,
//		"should not leak any go routine",
//		)
//
// # Examples
//
//   - success: NOT IMPLEMENTED
func NoGoRoutineLeak(t T, tested func(), msgAndArgs ...any) bool {
	// Domain: safety
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var ctx context.Context
	if c, ok := t.(contextualizer); ok {
		ctx = c.Context()
	} else {
		ctx = context.Background()
	}

	signature := leak.Leaked(ctx, tested)
	if signature == "" {
		return true
	}

	return Fail(t, "found leaked go routines: "+signature, msgAndArgs...)
}
