// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

func TestBooleanTrue(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)

	if !True(mock, true) {
		t.Error("True should return true")
	}

	if True(mock, false) {
		t.Error("True should return false")
	}

	if !True(mock, true) {
		t.Error("check error")
	}
}

func TestBooleanFalse(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)

	if !False(mock, false) {
		t.Error("False should return true")
	}
	if False(mock, true) {
		t.Error("False should return false")
	}
}
