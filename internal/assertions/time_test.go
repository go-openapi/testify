// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"regexp"
	"testing"
	"time"
)

func TestTimeWithinDuration(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	a := time.Now()
	b := a.Add(10 * time.Second)

	True(t, WithinDuration(mock, a, b, 10*time.Second), "A 10s difference is within a 10s time difference")
	True(t, WithinDuration(mock, b, a, 10*time.Second), "A 10s difference is within a 10s time difference")

	False(t, WithinDuration(mock, a, b, 9*time.Second), "A 10s difference is not within a 9s time difference")
	False(t, WithinDuration(mock, b, a, 9*time.Second), "A 10s difference is not within a 9s time difference")

	False(t, WithinDuration(mock, a, b, -9*time.Second), "A 10s difference is not within a 9s time difference")
	False(t, WithinDuration(mock, b, a, -9*time.Second), "A 10s difference is not within a 9s time difference")

	False(t, WithinDuration(mock, a, b, -11*time.Second), "A 10s difference is not within a 9s time difference")
	False(t, WithinDuration(mock, b, a, -11*time.Second), "A 10s difference is not within a 9s time difference")
}

func TestTimeWithinRange(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	n := time.Now()
	s := n.Add(-time.Second)
	e := n.Add(time.Second)

	True(t, WithinRange(mock, n, n, n), "Exact same actual, start, and end values return true")

	True(t, WithinRange(mock, n, s, e), "Time in range is within the time range")
	True(t, WithinRange(mock, s, s, e), "The start time is within the time range")
	True(t, WithinRange(mock, e, s, e), "The end time is within the time range")

	False(t, WithinRange(mock, s.Add(-time.Nanosecond), s, e, "Just before the start time is not within the time range"))
	False(t, WithinRange(mock, e.Add(time.Nanosecond), s, e, "Just after the end time is not within the time range"))

	False(t, WithinRange(mock, n, e, s, "Just after the end time is not within the time range"))
}

func TestTimeEqualityErrorFormatting(t *testing.T) {
	t.Parallel()
	mock := new(mockT)

	Equal(mock, time.Second*2, time.Millisecond)

	expectedErr := "\\s+Error Trace:\\s+Error:\\s+Not equal:\\s+\n\\s+expected: 2s\n\\s+actual\\s+: 1ms\n"
	Regexp(t, regexp.MustCompile(expectedErr), mock.errorString())
}
