// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build go1.26

package assertions

import (
	"fmt"
	"testing"
)

// TestErrorAsTypeNilTarget covers the check-only path (nil target), which the
// example-driven generated tests cannot exercise (a nil target makes E uninferable,
// so it must be specified explicitly here).
func TestErrorAsTypeNilTarget(t *testing.T) {
	t.Parallel()

	t.Run("match, nil target ignored", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		res := ErrorAsType[*customError](mock, fmt.Errorf("wrap: %w", &customError{}), nil)
		shouldPassOrFail(t, mock, res, true)
	})

	t.Run("no match, nil target", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		res := ErrorAsType[*customError](mock, ErrTest, nil)
		shouldPassOrFail(t, mock, res, false)
	})
}

// TestErrorAsTypeCapturesTarget verifies the matched error is written to target.
func TestErrorAsTypeCapturesTarget(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	sentinel := &customError{}
	var target *customError

	res := ErrorAsType(mock, fmt.Errorf("wrap: %w", sentinel), &target)
	shouldPassOrFail(t, mock, res, true)
	if target != sentinel {
		t.Errorf("expected target to capture the wrapped error %p, got %p", sentinel, target)
	}
}

// TestNotErrorAsType covers both outcomes of the negative assertion, including the
// nil-target (check-only) form.
func TestNotErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("absent type passes", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		res := NotErrorAsType[*customError](mock, ErrTest, nil)
		shouldPassOrFail(t, mock, res, true)
	})

	t.Run("present type fails", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)
		var target *customError

		res := NotErrorAsType(mock, fmt.Errorf("wrap: %w", &customError{}), &target)
		shouldPassOrFail(t, mock, res, false)
	})
}
