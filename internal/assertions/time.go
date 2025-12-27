// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"time"
)

// WithinDuration asserts that the two times are within duration delta of each other.
//
// # Usage
//
//	assertions.WithinDuration(t, time.Now(), 10*time.Second)
//
// # Examples
//
//	success: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 1, 0, time.UTC), 2*time.Second
//	failure: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 12, 0, 10, 0, time.UTC), 1*time.Second
func WithinDuration(t T, expected, actual time.Time, delta time.Duration, msgAndArgs ...any) bool {
	// Domain: time
	if h, ok := t.(H); ok {
		h.Helper()
	}

	dt := expected.Sub(actual)
	if dt < -delta || dt > delta {
		return Fail(t, fmt.Sprintf("Max difference between %v and %v allowed is %v, but difference was %v", expected, actual, delta, dt), msgAndArgs...)
	}

	return true
}

// WithinRange asserts that a time is within a time range (inclusive).
//
// # Usage
//
//	assertions.WithinRange(t, time.Now(), time.Now().Add(-time.Second), time.Now().Add(time.Second))
//
// # Examples
//
//	success: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
//	failure: time.Date(2024, 1, 1, 14, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC), time.Date(2024, 1, 1, 13, 0, 0, 0, time.UTC)
func WithinRange(t T, actual, start, end time.Time, msgAndArgs ...any) bool {
	// Domain: time
	if h, ok := t.(H); ok {
		h.Helper()
	}

	if end.Before(start) {
		return Fail(t, "Start should be before end", msgAndArgs...)
	}

	if actual.Before(start) {
		return Fail(t, fmt.Sprintf("Time %v expected to be in time range %v to %v, but is before the range", actual, start, end), msgAndArgs...)
	} else if actual.After(end) {
		return Fail(t, fmt.Sprintf("Time %v expected to be in time range %v to %v, but is after the range", actual, start, end), msgAndArgs...)
	}

	return true
}
