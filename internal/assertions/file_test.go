// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestFileExists(t *testing.T) {
	t.Parallel()

	for c := range fileExistsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			path := resolvePath(t, c)
			mock := new(mockT)
			res := FileExists(mock, path)
			shouldPassOrFail(t, mock, res, c.result)
		})
	}
}

func TestFileNotExists(t *testing.T) {
	t.Parallel()

	for c := range fileNotExistsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			path := resolvePath(t, c)
			mock := new(mockT)
			res := FileNotExists(mock, path)
			shouldPassOrFail(t, mock, res, c.result)
		})
	}
}

func TestDirExists(t *testing.T) {
	t.Parallel()

	for c := range dirExistsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			path := resolvePath(t, c)
			mock := new(mockT)
			res := DirExists(mock, path)
			shouldPassOrFail(t, mock, res, c.result)
		})
	}
}

func TestDirNotExists(t *testing.T) {
	t.Parallel()

	for c := range dirNotExistsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			path := resolvePath(t, c)
			mock := new(mockT)
			res := DirNotExists(mock, path)
			shouldPassOrFail(t, mock, res, c.result)
		})
	}
}

func TestFileEmpty(t *testing.T) {
	t.Parallel()

	for c := range fileEmptyCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			path := resolvePath(t, c)
			mock := new(mockT)
			res := FileEmpty(mock, path)
			shouldPassOrFail(t, mock, res, c.result)
		})
	}
}

func TestFileNotEmpty(t *testing.T) {
	t.Parallel()

	for c := range fileNotEmptyCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			path := resolvePath(t, c)
			mock := new(mockT)
			res := FileNotEmpty(mock, path)
			shouldPassOrFail(t, mock, res, c.result)
		})
	}
}

func TestFileLstatPermissionError(t *testing.T) {
	t.Parallel()

	if os.Getuid() == 0 {
		t.Skip("skipping permission test when running as root")
	}

	// Create a directory with a file inside, then remove execute permission
	// from the directory. os.Lstat on the file will fail with EACCES, not ENOENT.
	dir := t.TempDir()
	filePath := filepath.Join(dir, "secret")
	if err := os.WriteFile(filePath, []byte("data"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(dir, 0o000); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { os.Chmod(dir, 0o755) }) //nolint:errcheck // best-effort restore for cleanup

	mock := new(mockT)
	if FileExists(mock, filePath) {
		t.Error("expected FileExists to return false for inaccessible file")
	}
	if !mock.failed {
		t.Error("expected FileExists to fail")
	}
}

func TestFileErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, fileFailCases())
}

// ============================================================================
// Table-driven tests for file assertions
// ============================================================================

type fileTestCase struct {
	name      string
	path      string // direct path to test
	symlinkTo string // if set, creates a symlink to this target instead
	result    bool
}

func resolvePath(t *testing.T, c fileTestCase) string {
	t.Helper()

	if c.symlinkTo != "" {
		return getTempSymlinkPath(t, c.symlinkTo)
	}

	return c.path
}

func getTempSymlinkPath(t *testing.T, file string) string {
	t.Helper()

	tempDir := t.TempDir()
	link := filepath.Join(tempDir, filepath.Base(file)+"_symlink")
	if err := os.Symlink(file, link); err != nil {
		t.Fatalf("could not create temp symlink %q pointing to %q: %v", link, file, err)
	}
	return link
}

// ============================================================================
// Test cases
// ============================================================================

func fileExistsCases() iter.Seq[fileTestCase] {
	return slices.Values([]fileTestCase{
		{name: "existing-file", path: filepath.Join("testdata", "existing_file"), result: true},
		{name: "non-existent", path: "random_file", result: false},
		{name: "directory", path: filepath.Join("testdata", "existing_dir"), result: false},
		{name: "symlink/existing-file", symlinkTo: filepath.Join("testdata", "existing_file"), result: true},
		{name: "symlink/broken", symlinkTo: "non_existent_file", result: true},
	})
}

func fileNotExistsCases() iter.Seq[fileTestCase] {
	return slices.Values([]fileTestCase{
		{name: "existing-file", path: filepath.Join("testdata", "existing_file"), result: false},
		{name: "non-existent", path: "non_existent_file", result: true},
		{name: "directory", path: filepath.Join("testdata", "existing_dir"), result: true},
		{name: "symlink/existing-file", symlinkTo: filepath.Join("testdata", "existing_file"), result: false},
		{name: "symlink/broken", symlinkTo: "non_existent_file", result: false},
	})
}

func dirExistsCases() iter.Seq[fileTestCase] {
	return slices.Values([]fileTestCase{
		{name: "is-file", path: filepath.Join("testdata", "existing_file"), result: false},
		{name: "non-existent", path: "non_existent_dir", result: false},
		{name: "existing-dir", path: filepath.Join("testdata", "existing_dir"), result: true},
		{name: "symlink/to-file", symlinkTo: filepath.Join("testdata", "existing_file"), result: false},
		{name: "symlink/broken", symlinkTo: "non_existent_dir", result: false},
	})
}

func dirNotExistsCases() iter.Seq[fileTestCase] {
	return slices.Values([]fileTestCase{
		{name: "is-file", path: filepath.Join("testdata", "existing_file"), result: true},
		{name: "non-existent", path: "non_existent_dir", result: true},
		{name: "existing-dir", path: filepath.Join("testdata", "existing_dir"), result: false},
		{name: "symlink/to-file", symlinkTo: filepath.Join("testdata", "existing_file"), result: true},
		{name: "symlink/broken", symlinkTo: "non_existent_dir", result: true},
	})
}

func fileEmptyCases() iter.Seq[fileTestCase] {
	return slices.Values([]fileTestCase{
		{name: "empty-file", path: filepath.Join("testdata", "empty_file"), result: true},
		{name: "non-empty-file", path: filepath.Join("testdata", "existing_file"), result: false},
		{name: "non-existent", path: "random_file", result: false},
		{name: "directory", path: filepath.Join("testdata", "existing_dir"), result: false},
		{name: "symlink/empty-file", symlinkTo: filepath.Join("testdata", "empty_file"), result: true},
		{name: "symlink/non-empty-file", symlinkTo: filepath.Join("testdata", "existing_file"), result: false},
		{name: "symlink/broken", symlinkTo: "non_existent_file", result: false},
	})
}

func fileNotEmptyCases() iter.Seq[fileTestCase] {
	return slices.Values([]fileTestCase{
		{name: "non-empty-file", path: filepath.Join("testdata", "existing_file"), result: true},
		{name: "empty-file", path: filepath.Join("testdata", "empty_file"), result: false},
		{name: "non-existent", path: "non_existent_file", result: false},
		{name: "directory", path: filepath.Join("testdata", "existing_dir"), result: false},
		{name: "symlink/empty-file", symlinkTo: filepath.Join("testdata", "empty_file"), result: false},
		{name: "symlink/non-empty-file", symlinkTo: filepath.Join("testdata", "existing_file"), result: true},
		{name: "symlink/broken", symlinkTo: "non_existent_file", result: false},
	})
}

// ============================================================================
// TestFileErrorMessages
// ============================================================================

func fileFailCases() iter.Seq[failCase] {
	return slices.Values([]failCase{
		{
			name:         "FileExists/nonexistent",
			assertion:    func(t T) bool { return FileExists(t, "nonexistent_file") },
			wantContains: []string{"unable to find file"},
		},
		{
			name:         "FileExists/is-directory",
			assertion:    func(t T) bool { return FileExists(t, filepath.Join("testdata", "existing_dir")) },
			wantContains: []string{"is a directory"},
		},
		{
			name:         "FileNotExists/existing-file",
			assertion:    func(t T) bool { return FileNotExists(t, filepath.Join("testdata", "existing_file")) },
			wantContains: []string{"file", "exists"},
		},
		{
			name:         "DirExists/nonexistent",
			assertion:    func(t T) bool { return DirExists(t, "nonexistent_dir") },
			wantContains: []string{"unable to find directory"},
		},
		{
			name:         "DirExists/is-file",
			assertion:    func(t T) bool { return DirExists(t, filepath.Join("testdata", "existing_file")) },
			wantContains: []string{"is a file"},
		},
		{
			name:         "DirNotExists/existing-dir",
			assertion:    func(t T) bool { return DirNotExists(t, filepath.Join("testdata", "existing_dir")) },
			wantContains: []string{"directory", "exists"},
		},
		{
			name:         "FileEmpty/non-empty-file",
			assertion:    func(t T) bool { return FileEmpty(t, filepath.Join("testdata", "existing_file")) },
			wantContains: []string{"is not empty"},
		},
		{
			name:         "FileEmpty/nonexistent",
			assertion:    func(t T) bool { return FileEmpty(t, "nonexistent_file") },
			wantContains: []string{"unable to find file"},
		},
		{
			name:         "FileNotEmpty/empty-file",
			assertion:    func(t T) bool { return FileNotEmpty(t, filepath.Join("testdata", "empty_file")) },
			wantContains: []string{"is empty"},
		},
		{
			name:         "FileNotEmpty/nonexistent",
			assertion:    func(t T) bool { return FileNotEmpty(t, "nonexistent_file") },
			wantContains: []string{"unable to find file"},
		},
	})
}
