// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.26

package assert

import (
	"fmt"
	"testing"
)

func TestErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorAsType(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if !result {
			t.Error("ErrorAsType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := ErrorAsType(mock, ErrTest, new(*dummyError))
		if result {
			t.Error("ErrorAsType should return false on failure")
		}
		if !mock.failed {
			t.Error("ErrorAsType should mark test as failed")
		}
	})
}

func TestNotErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorAsType(mock, ErrTest, new(*dummyError))
		if !result {
			t.Error("NotErrorAsType should return true on success")
		}
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := NotErrorAsType(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		if result {
			t.Error("NotErrorAsType should return false on failure")
		}
		if !mock.failed {
			t.Error("NotErrorAsType should mark test as failed")
		}
	})
}
