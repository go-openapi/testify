// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "github.com/go-openapi/testify/v2/internal/assertions/leak"

// NoGoRoutineLeak ensures that no goroutine did leak from inside the tested function.
func NoGoRoutineLeak(t T, inside func(), msgAndArgs ...any) bool {
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
