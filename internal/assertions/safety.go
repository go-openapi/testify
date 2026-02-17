// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"context"
	"runtime"

	"github.com/go-openapi/testify/v2/internal/fdleak"
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
//	NoGoRoutineLeak(t, func() {
//		...
//	},
//	"should not leak any go routine",
//	)
//
// # Examples
//
//	success: func() {}
func NoGoRoutineLeak(t T, tested func(), msgAndArgs ...any) bool {
	// Domain: safety
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var ctx context.Context
	c, ok := t.(contextualizer)
	if ok {
		ctx = c.Context()
	}
	if ctx == nil {
		ctx = context.Background()
	}

	signature := leak.Leaked(ctx, tested)
	if signature == "" {
		return true
	}

	return Fail(t, "found leaked go routines: "+signature, msgAndArgs...)
}

// NoFileDescriptorLeak ensures that no file descriptor leaks from inside the tested function.
//
// This assertion works on Linux only (via /proc/self/fd).
// On other platforms, the test is skipped.
//
// NOTE: this assertion is not compatible with parallel tests.
// File descriptors are a process-wide resource; concurrent tests
// opening files would cause false positives.
//
// Sockets, pipes, and anonymous inodes are filtered out by default,
// as these are typically managed by the Go runtime.
//
// # Concurrency
//
// [NoFileDescriptorLeak] is not compatible with parallel tests.
// File descriptors are a process-wide resource; any concurrent I/O
// from other goroutines may cause false positives.
//
// Calls to [NoFileDescriptorLeak] are serialized with a mutex
// to prevent multiple leak checks from interfering with each other.
//
// # Usage
//
//	NoFileDescriptorLeak(t, func() {
//		// code that should not leak file descriptors
//	})
//
// # Examples
//
//	success: func() {}
func NoFileDescriptorLeak(t T, tested func(), msgAndArgs ...any) bool {
	// Domain: safety
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if runtime.GOOS != "linux" { //nolint:goconst // well-known runtime value
		if s, ok := t.(skipper); ok {
			s.Skip("NoFileDescriptorLeak requires Linux (/proc/self/fd)")
		}

		return true
	}

	msg, err := fdleak.Leaked(tested)
	if err != nil {
		return Fail(t, "file descriptor snapshot failed: "+err.Error(), msgAndArgs...)
	}

	if msg == "" {
		return true
	}

	return Fail(t, msg, msgAndArgs...)
}
