// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"golang.org/x/tools/go/packages"
)

// TestGuardedFilesOnDisk checks which files the textual scan reports as go-version-guarded.
//
// Only a plain "go1.N" constraint counts. A file selected on an OS or an architecture is
// absent on some machines by design, and a test file never reaches the load in the first
// place: neither may be reported as dropped.
func TestGuardedFilesOnDisk(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	write := func(name, content string) {
		t.Helper()
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}

	write("plain.go", "package assertions\n")
	write("guarded.go", "//go:build go1.99\n\npackage assertions\n")
	write("guarded_test.go", "//go:build go1.99\n\npackage assertions\n")
	write("platform.go", "//go:build !windows\n\npackage assertions\n")
	write("combined.go", "//go:build go1.99 && !windows\n\npackage assertions\n")
	write("broken.go", "//go:build go1.99\n\npackage ***\n")
	write("notgo.txt", "//go:build go1.99\n")

	guarded, err := guardedFilesOnDisk(dir)
	if err != nil {
		t.Fatalf("guardedFilesOnDisk: %v", err)
	}

	if got, want := len(guarded), 1; got != want {
		t.Fatalf("expected %d guarded file, got %d: %v", want, got, guarded)
	}
	if got, want := guarded["guarded.go"], "go1.99"; got != want {
		t.Errorf("expected guarded.go to carry %q, got %q", want, got)
	}
}

// TestVerifyGuardedFilesLoaded covers the rail itself: a go-version-guarded file that exists
// on disk but never reached the typed load means the go command running the scan is older
// than the guard, and generating from that view would silently drop assertions.
func TestVerifyGuardedFilesLoaded(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	for name, content := range map[string]string{
		"plain.go":   "package assertions\n",
		"guarded.go": "//go:build go1.99\n\npackage assertions\n",
	} {
		if err := os.WriteFile(filepath.Join(dir, name), []byte(content), 0o600); err != nil {
			t.Fatalf("write %s: %v", name, err)
		}
	}

	t.Run("guard dropped by an older toolchain", func(t *testing.T) {
		t.Parallel()

		pkg := &packages.Package{Dir: dir, GoFiles: []string{filepath.Join(dir, "plain.go")}}

		err := verifyGuardedFilesLoaded(pkg)
		if !errors.Is(err, ErrGuardedFileDropped) {
			t.Fatalf("expected ErrGuardedFileDropped, got %v", err)
		}
		for _, want := range []string{"guarded.go", "go1.99", "GOTOOLCHAIN=go1.99.0"} {
			if !strings.Contains(err.Error(), want) {
				t.Errorf("expected the error to mention %q, got: %v", want, err)
			}
		}
	})

	t.Run("guard observed by the load", func(t *testing.T) {
		t.Parallel()

		pkg := &packages.Package{Dir: dir, GoFiles: []string{
			filepath.Join(dir, "plain.go"),
			filepath.Join(dir, "guarded.go"),
		}}

		if err := verifyGuardedFilesLoaded(pkg); err != nil {
			t.Errorf("expected no error when every guarded file loaded, got %v", err)
		}
	})

	t.Run("no package directory", func(t *testing.T) {
		t.Parallel()

		if err := verifyGuardedFilesLoaded(&packages.Package{}); err != nil {
			t.Errorf("expected no error without a package directory, got %v", err)
		}
	})
}
