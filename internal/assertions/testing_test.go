// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"testing"
)

const failedStr = "failed"

func TestTestingFail(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	if Fail(mock, failedStr) {
		t.Error("Fail is expected to return false")
	}
}

func TestTestingFailNow(t *testing.T) {
	t.Parallel()

	t.Run("with plain T (no Failnow support)", func(t *testing.T) {
		t.Parallel()
		mock := new(mockT)

		t.Run("should panic", func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Error("should panic since mockT is missing FailNow()")
				}
			}()
			FailNow(mock, failedStr)
		})
	})

	t.Run("with full T (Failnow support)", func(t *testing.T) {
		t.Parallel()
		mock := new(mockFailNowT)

		t.Run("should not panic", func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("should call mockT.FailNow() rather than panicking: %v", r)
				}
			}()
			FailNow(mock, failedStr)
		})
	})
}
