// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.26

package assert

import (
	"fmt"
	"testing"
)

func TestErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorAsTypef(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if !result {
			t.Error("ErrorAsTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorAsTypef(mock, ErrTest, new(*dummyError), "test message")
		if result {
			t.Error("ErrorAsTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorAsTypef should mark test as failed")
		}
	})
}

func TestNotErrorAsTypef(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorAsTypef(mock, ErrTest, new(*dummyError), "test message")
		if !result {
			t.Error("NotErrorAsTypef should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorAsTypef(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError), "test message")
		if result {
			t.Error("NotErrorAsTypef should return false on failure")
		}
		if !mock.failed {
			t.Error("NotErrorAsTypef should mark test as failed")
		}
	})
}
