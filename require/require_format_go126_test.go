// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.26

package require

import (
	"fmt"
	"testing"
)

func TestErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAsTypef(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAsTypef(mock, ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorAsTypef should call FailNow()")
		}
	})
}

func TestNotErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAsTypef(mock, ErrTest, new(*dummyError), "test message")
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAsTypef(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorAsTypef should call FailNow()")
		}
	})
}
