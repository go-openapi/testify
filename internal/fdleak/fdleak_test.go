// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package fdleak

import (
	"testing"
)

func TestDiff(t *testing.T) {
	before := map[int]FDInfo{
		0: {FD: 0, Kind: KindChar, Target: "/dev/stdin"},
		1: {FD: 1, Kind: KindChar, Target: "/dev/stdout"},
		2: {FD: 2, Kind: KindChar, Target: "/dev/stderr"},
		3: {FD: 3, Kind: KindPipe, Target: "pipe:[12345]"},
	}

	after := map[int]FDInfo{
		0: {FD: 0, Kind: KindChar, Target: "/dev/stdin"},
		1: {FD: 1, Kind: KindChar, Target: "/dev/stdout"},
		2: {FD: 2, Kind: KindChar, Target: "/dev/stderr"},
		3: {FD: 3, Kind: KindPipe, Target: "pipe:[12345]"},
		5: {FD: 5, Kind: KindFile, Target: "/tmp/leaked.txt"},         // leaked regular file
		6: {FD: 6, Kind: KindSocket, Target: "socket:[67890]"},        // filtered: socket
		7: {FD: 7, Kind: KindPipe, Target: "pipe:[11111]"},            // filtered: pipe
		8: {FD: 8, Kind: KindOther, Target: "anon_inode:[eventpoll]"}, // filtered: other
		9: {FD: 9, Kind: KindFile, Target: "/dev/null"},               // leaked device
	}

	leaked := Diff(before, after)

	if len(leaked) != 2 {
		t.Fatalf("expected 2 leaked FDs, got %d: %+v", len(leaked), leaked)
	}

	// Sorted by FD number.
	if leaked[0].FD != 5 || leaked[0].Target != "/tmp/leaked.txt" {
		t.Errorf("leaked[0] = %+v, want fd 5 /tmp/leaked.txt", leaked[0])
	}

	if leaked[1].FD != 9 || leaked[1].Target != "/dev/null" {
		t.Errorf("leaked[1] = %+v, want fd 9 /dev/null", leaked[1])
	}
}

func TestDiff_NoLeaks(t *testing.T) {
	fds := map[int]FDInfo{
		0: {FD: 0, Kind: KindChar, Target: "/dev/stdin"},
		1: {FD: 1, Kind: KindChar, Target: "/dev/stdout"},
	}

	leaked := Diff(fds, fds)

	if len(leaked) != 0 {
		t.Errorf("expected no leaks, got %+v", leaked)
	}
}

func TestFormatLeaked(t *testing.T) {
	leaked := []FDInfo{
		{FD: 7, Kind: KindFile, Target: "/tmp/unclosed.txt"},
		{FD: 9, Kind: KindFile, Target: "/dev/null"},
	}

	result := FormatLeaked(leaked)
	expected := "found 2 leaked file descriptor(s):\n  fd 7: /tmp/unclosed.txt\n  fd 9: /dev/null\n"

	if result != expected {
		t.Errorf("FormatLeaked:\ngot:  %q\nwant: %q", result, expected)
	}
}

func TestFormatLeaked_Empty(t *testing.T) {
	result := FormatLeaked(nil)

	if result != "" {
		t.Errorf("expected empty string for nil input, got %q", result)
	}
}

func TestFDInfo_isFiltered(t *testing.T) {
	cases := []struct {
		name     string
		kind     Kind
		filtered bool
	}{
		{"file", KindFile, false},
		{"char", KindChar, false},
		{"unknown", KindUnknown, false},
		{"socket", KindSocket, true},
		{"pipe", KindPipe, true},
		{"other", KindOther, true},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := FDInfo{Kind: tc.kind}.isFiltered()
			if got != tc.filtered {
				t.Errorf("Kind=%v isFiltered = %v, want %v", tc.kind, got, tc.filtered)
			}
		})
	}
}
