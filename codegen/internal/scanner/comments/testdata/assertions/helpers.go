//go:build integrationtest

// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

// T is the minimal interface for test assertions (stub for testdata).
type T interface {
	// Maintainers: we should clarify when we need to impose FailNow on the interface and when we don't.
	Errorf(format string, args ...any)
	FailNow()
}

// H is the interface for test helpers (stub for testdata).
type H interface {
	Helper()
}

// Comparison is a function type for custom assertions (stub for testdata).
type Comparison func() bool

// Fail reports a failure (stub for testdata).
func Fail(t T, failureMessage string, msgAndArgs ...any) bool {
	t.Errorf(failureMessage)
	return false
}
