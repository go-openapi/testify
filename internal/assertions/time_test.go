// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"testing"
	"time"
)

func TestTimeWithinDuration(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	a := time.Now()
	b := a.Add(10 * time.Second)

	if !WithinDuration(mock, a, b, 10*time.Second) {
		t.Error("A 10s difference is within a 10s time difference")
	}
	if !WithinDuration(mock, b, a, 10*time.Second) {
		t.Error("A 10s difference is within a 10s time difference (reversed)")
	}

	if WithinDuration(mock, a, b, 9*time.Second) {
		t.Error("A 10s difference is not within a 9s time difference")
	}
	if WithinDuration(mock, b, a, 9*time.Second) {
		t.Error("A 10s difference is not within a 9s time difference (reversed)")
	}

	if WithinDuration(mock, a, b, -9*time.Second) {
		t.Error("A 10s difference is not within a -9s time difference")
	}
	if WithinDuration(mock, b, a, -9*time.Second) {
		t.Error("A 10s difference is not within a -9s time difference (reversed)")
	}

	if WithinDuration(mock, a, b, -11*time.Second) {
		t.Error("A 10s difference is not within a -11s time difference")
	}
	if WithinDuration(mock, b, a, -11*time.Second) {
		t.Error("A 10s difference is not within a -11s time difference (reversed)")
	}
}

func TestTimeWithinRange(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	n := time.Now()
	s := n.Add(-time.Second)
	e := n.Add(time.Second)

	if !WithinRange(mock, n, n, n) {
		t.Error("Exact same actual, start, and end values should return true")
	}

	if !WithinRange(mock, n, s, e) {
		t.Error("Time in range should be within the time range")
	}
	if !WithinRange(mock, s, s, e) {
		t.Error("The start time should be within the time range")
	}
	if !WithinRange(mock, e, s, e) {
		t.Error("The end time should be within the time range")
	}

	if WithinRange(mock, s.Add(-time.Nanosecond), s, e) {
		t.Error("Just before the start time should not be within the time range")
	}
	if WithinRange(mock, e.Add(time.Nanosecond), s, e) {
		t.Error("Just after the end time should not be within the time range")
	}

	if WithinRange(mock, n, e, s) {
		t.Error("Reversed range (start > end) should return false")
	}
}

func TestTimeErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, timeFailCases())
}

// =======================================
// TestTimeErrorMessages
// =======================================

func timeFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "Equal/time-formatting",
			assertion:    func(t T) bool { return Equal(t, time.Second*2, time.Millisecond) },
			wantContains: []string{"Not equal:", "2s", "1ms"},
		},
	})
}
