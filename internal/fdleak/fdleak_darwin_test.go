// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build darwin

package fdleak

import (
	"context"
	"net"
	"os"
	"strings"
	"testing"
)

func TestSnapshot(t *testing.T) {
	fds, err := Snapshot()
	if err != nil {
		t.Fatalf("Snapshot() error: %v", err)
	}

	// stdin, stdout, stderr should always be present.
	for _, fd := range []int{0, 1, 2} {
		info, ok := fds[fd]
		if !ok {
			t.Errorf("expected fd %d (stdin/stdout/stderr) in snapshot", fd)
			continue
		}
		if info.Kind == KindUnknown {
			t.Errorf("fd %d has KindUnknown; expected a concrete kind (got target=%q)", fd, info.Target)
		}
	}
}

func TestLeaked_NoLeak(t *testing.T) {
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

// TestLeaked_WithLeak additionally verifies that fcntl(F_GETPATH) recovers
// the file's absolute path on darwin.
func TestLeaked_WithLeak(t *testing.T) {
	var leakedFile *os.File
	dir := t.TempDir()

	leaked, err := Leaked(func() {
		f, err := os.CreateTemp(dir, "fdleak-test-*")
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
		t.Fatal("expected leaked file descriptor to be detected, but found none")
	}

	// F_GETPATH should have resolved the full path of the leaked temp file.
	// On darwin t.TempDir() lives under /private/var/... so match on the basename prefix.
	if !strings.Contains(leaked, "fdleak-test-") {
		t.Errorf("expected leak report to include the resolved file path; got:\n%s", leaked)
	} else {
		t.Logf("detected leak:\n%s", leaked)
	}
}

// TestLeaked_SocketsFiltered exercises the fstat fallback on darwin: sockets
// have no vnode path so F_GETPATH fails, and they must be classified as
// KindSocket via fstat to be filtered out.
func TestLeaked_SocketsFiltered(t *testing.T) {
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

// TestLeaked_PipesFiltered exercises the fstat fallback for pipes on darwin.
func TestLeaked_PipesFiltered(t *testing.T) {
	var r, w *os.File

	leaked, err := Leaked(func() {
		pr, pw, err := os.Pipe()
		if err != nil {
			t.Fatalf("os.Pipe: %v", err)
		}

		r, w = pr, pw // intentionally not closed — pipe FDs should be filtered
	})

	t.Cleanup(func() {
		if r != nil {
			r.Close()
		}
		if w != nil {
			w.Close()
		}
	})

	if err != nil {
		t.Fatalf("Leaked() error: %v", err)
	}

	if leaked != "" {
		t.Errorf("expected pipe FDs to be filtered, but got:\n%s", leaked)
	}
}
