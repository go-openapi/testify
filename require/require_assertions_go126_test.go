// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Code generated with github.com/go-openapi/testify/codegen/v2; DO NOT EDIT.

//go:build go1.26

package require

import (
	"fmt"
	"testing"
)

func TestErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAsType(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		ErrorAsType(mock, ErrTest, new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("ErrorAsType should call FailNow()")
		}
	})
}

func TestNotErrorAsType(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAsType(mock, ErrTest, new(*dummyError))
		// require functions don't return a value
	})

	t.Run("failure", func(t *testing.T) {
		t.Parallel()

		mock := new(mockFailNowT)
		NotErrorAsType(mock, fmt.Errorf("wrap: %w", &dummyError{}), new(*dummyError))
		// require functions don't return a value
		if !mock.failed {
			t.Error("NotErrorAsType should call FailNow()")
		}
	})
}
