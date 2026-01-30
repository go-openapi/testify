// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"testing"
)

func TestBooleanTrue(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)

	if !True(mock, true) {
		t.Error("True should return true")
	}

	if True(mock, false) {
		t.Error("True should return false")
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

func TestBooleanTrueTFalseT(t *testing.T) {
	t.Parallel()

	type X bool
	var truthy X = true
	var falsy X = false

	t.Run("with TrueT on redeclared bool type", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		if !TrueT(mock, truthy) {
			t.Error("TrueT should return true")
		}
		if TrueT(mock, falsy) {
			t.Error("TrueT should return false")
		}
	})

	t.Run("with FalseT on redeclared bool type", func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		if FalseT(mock, truthy) {
			t.Error("TrueT should return true")
		}
		if !FalseT(mock, falsy) {
			t.Error("FalseT should return false")
		}
	})
}

func TestBooleanErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, booleanFailCases())
}

func booleanFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:      "True/false-value",
			assertion: func(t T) bool { return True(t, false) },
			wantError: "Should be true",
		},
		{
			name:      "False/true-value",
			assertion: func(t T) bool { return False(t, true) },
			wantError: "Should be false",
		},
		{
			name:      "TrueT/false-value",
			assertion: func(t T) bool { return TrueT(t, false) },
			wantError: "Should be true",
		},
		{
			name:      "FalseT/true-value",
			assertion: func(t T) bool { return FalseT(t, true) },
			wantError: "Should be false",
		},
	})
}
