// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package fdleak

import (
	"context"
	"net"
	"os"
	"runtime"
	"testing"
)

func skipIfNotLinux(t *testing.T) {
	t.Helper()

	if runtime.GOOS != "linux" {
		t.Skip("file descriptor leak detection requires Linux")
	}
}

func TestSnapshot(t *testing.T) {
	skipIfNotLinux(t)

	fds, err := Snapshot()
	if err != nil {
		t.Fatalf("Snapshot() error: %v", err)
	}

	// stdin, stdout, stderr should always be present.
	for _, fd := range []int{0, 1, 2} {
		if _, ok := fds[fd]; !ok {
			t.Errorf("expected fd %d (stdin/stdout/stderr) in snapshot", fd)
		}
	}
}

func TestLeaked_NoLeak(t *testing.T) {
	skipIfNotLinux(t)

	leaked, err := Leaked(func() {
		// Clean function — no file descriptors opened.
	})
	if err != nil {
		t.Fatalf("Leaked() error: %v", err)
	}

	if leaked != "" {
		t.Errorf("expected no leaked file descriptors, got:\n%s", leaked)
	}
}

func TestLeaked_WithLeak(t *testing.T) {
	skipIfNotLinux(t)

	var leakedFile *os.File

	leaked, err := Leaked(func() {
		f, err := os.CreateTemp(t.TempDir(), "fdleak-test-*")
		if err != nil {
			t.Fatalf("CreateTemp: %v", err)
		}

		leakedFile = f // intentionally not closed
	})

	t.Cleanup(func() {
		if leakedFile != nil {
			leakedFile.Close()
			os.Remove(leakedFile.Name())
		}
	})

	if err != nil {
		t.Fatalf("Leaked() error: %v", err)
	}

	if leaked == "" {
		t.Error("expected leaked file descriptor to be detected, but found none")
	} else {
		t.Logf("detected leak:\n%s", leaked)
	}
}

func TestLeaked_SocketsFiltered(t *testing.T) {
	skipIfNotLinux(t)

	var leakedListener net.Listener

	leaked, err := Leaked(func() {
		var lc net.ListenConfig
		ln, err := lc.Listen(context.Background(), "tcp", "127.0.0.1:0")
		if err != nil {
			t.Fatalf("net.Listen: %v", err)
		}

		leakedListener = ln // intentionally not closed — socket FD should be filtered
	})

	t.Cleanup(func() {
		if leakedListener != nil {
			leakedListener.Close()
		}
	})

	if err != nil {
		t.Fatalf("Leaked() error: %v", err)
	}

	if leaked != "" {
		t.Errorf("expected socket FD to be filtered, but got:\n%s", leaked)
	}
}

func TestDiff(t *testing.T) {
	before := map[int]FDInfo{
		0: {FD: 0, Target: "/dev/stdin"},
		1: {FD: 1, Target: "/dev/stdout"},
		2: {FD: 2, Target: "/dev/stderr"},
		3: {FD: 3, Target: "pipe:[12345]"},
	}

	after := map[int]FDInfo{
		0: {FD: 0, Target: "/dev/stdin"},
		1: {FD: 1, Target: "/dev/stdout"},
		2: {FD: 2, Target: "/dev/stderr"},
		3: {FD: 3, Target: "pipe:[12345]"},
		5: {FD: 5, Target: "/tmp/leaked.txt"},        // leaked regular file
		6: {FD: 6, Target: "socket:[67890]"},         // filtered: socket
		7: {FD: 7, Target: "pipe:[11111]"},           // filtered: pipe
		8: {FD: 8, Target: "anon_inode:[eventpoll]"}, // filtered: anon_inode
		9: {FD: 9, Target: "/dev/null"},              // leaked device
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
		0: {FD: 0, Target: "/dev/stdin"},
		1: {FD: 1, Target: "/dev/stdout"},
	}

	leaked := Diff(fds, fds)

	if len(leaked) != 0 {
		t.Errorf("expected no leaks, got %+v", leaked)
	}
}

func TestFormatLeaked(t *testing.T) {
	leaked := []FDInfo{
		{FD: 7, Target: "/tmp/unclosed.txt"},
		{FD: 9, Target: "/dev/null"},
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
