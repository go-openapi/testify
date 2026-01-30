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

	mock := new(mockT)
	if !FileExists(mock, filepath.Join("testdata", "existing_file")) {
		t.Error("expected FileExists to return true for existing file")
	}

	mock = new(mockT)
	if FileExists(mock, "random_file") {
		t.Error("expected FileExists to return false for random_file")
	}

	mock = new(mockT)
	if FileExists(mock, filepath.Join("testdata", "existing_dir")) {
		t.Error("expected FileExists to return false for directory")
	}

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	if !FileExists(mock, link) {
		t.Error("expected FileExists to return true for symlink to existing file")
	}

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	if !FileExists(mock, link) {
		t.Error("expected FileExists to return true for symlink (broken symlink is still a file)")
	}
}

func TestFileFileNotExists(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	if FileNotExists(mock, filepath.Join("testdata", "existing_file")) {
		t.Error("expected FileNotExists to return false for existing file")
	}

	mock = new(mockT)
	if !FileNotExists(mock, "non_existent_file") {
		t.Error("expected FileNotExists to return true for non-existent file")
	}

	mock = new(mockT)
	if !FileNotExists(mock, filepath.Join("testdata", "existing_dir")) {
		t.Error("expected FileNotExists to return true for directory")
	}

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	if FileNotExists(mock, link) {
		t.Error("expected FileNotExists to return false for symlink to existing file")
	}

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	if FileNotExists(mock, link) {
		t.Error("expected FileNotExists to return false for symlink")
	}
}

func TestFileDirExists(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	if DirExists(mock, filepath.Join("testdata", "existing_file")) {
		t.Error("expected DirExists to return false for file")
	}

	mock = new(mockT)
	if DirExists(mock, "non_existent_dir") {
		t.Error("expected DirExists to return false for non-existent dir")
	}

	mock = new(mockT)
	if !DirExists(mock, filepath.Join("testdata", "existing_dir")) {
		t.Error("expected DirExists to return true for existing dir")
	}

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	if DirExists(mock, link) {
		t.Error("expected DirExists to return false for symlink to file")
	}

	link = getTempSymlinkPath(t, "non_existent_dir")
	mock = new(mockT)
	if DirExists(mock, link) {
		t.Error("expected DirExists to return false for symlink to non-existent dir")
	}
}

func TestFileDirNotExists(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	if !DirNotExists(mock, filepath.Join("testdata", "existing_file")) {
		t.Error("expected DirNotExists to return true for file")
	}

	mock = new(mockT)
	if !DirNotExists(mock, "non_existent_dir") {
		t.Error("expected DirNotExists to return true for non-existent dir")
	}

	mock = new(mockT)
	if DirNotExists(mock, filepath.Join("testdata", "existing_dir")) {
		t.Error("expected DirNotExists to return false for existing dir")
	}

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	if !DirNotExists(mock, link) {
		t.Error("expected DirNotExists to return true for symlink to file")
	}

	link = getTempSymlinkPath(t, "non_existent_dir")
	mock = new(mockT)
	if !DirNotExists(mock, link) {
		t.Error("expected DirNotExists to return true for symlink to non-existent dir")
	}
}

func TestFileEmpty(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	if !FileEmpty(mock, filepath.Join("testdata", "empty_file")) {
		t.Error("expected FileEmpty to return true for empty file")
	}

	mock = new(mockT)
	if FileEmpty(mock, filepath.Join("testdata", "existing_file")) {
		t.Error("expected FileEmpty to return false for non-empty file")
	}

	mock = new(mockT)
	if FileEmpty(mock, "random_file") {
		t.Error("expected FileEmpty to return false for non-existent file")
	}

	mock = new(mockT)
	if FileEmpty(mock, filepath.Join("testdata", "existing_dir")) {
		t.Error("expected FileEmpty to return false for directory")
	}

	link := getTempSymlinkPath(t, filepath.Join("testdata", "empty_file"))
	mock = new(mockT)
	if !FileEmpty(mock, link) {
		t.Error("expected FileEmpty to return true for symlink to empty file")
	}

	link = getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	if FileEmpty(mock, link) {
		t.Error("expected FileEmpty to return false for symlink to non-empty file")
	}

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	if FileEmpty(mock, link) {
		t.Error("expected FileEmpty to return false for symlink to non-existent file")
	}
}

func TestFileNotEmpty(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	if !FileNotEmpty(mock, filepath.Join("testdata", "existing_file")) {
		t.Error("expected FileNotEmpty to return true for non-empty file")
	}

	mock = new(mockT)
	if FileNotEmpty(mock, filepath.Join("testdata", "empty_file")) {
		t.Error("expected FileNotEmpty to return false for empty file")
	}

	mock = new(mockT)
	if FileNotEmpty(mock, "non_existent_file") {
		t.Error("expected FileNotEmpty to return false for non-existent file")
	}

	mock = new(mockT)
	if FileNotEmpty(mock, filepath.Join("testdata", "existing_dir")) {
		t.Error("expected FileNotEmpty to return false for directory")
	}

	link := getTempSymlinkPath(t, filepath.Join("testdata", "empty_file"))
	mock = new(mockT)
	if FileNotEmpty(mock, link) {
		t.Error("expected FileNotEmpty to return false for symlink to empty file")
	}

	link = getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	if !FileNotEmpty(mock, link) {
		t.Error("expected FileNotEmpty to return true for symlink to non-empty file")
	}

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	if FileNotExists(mock, link) {
		t.Error("expected FileNotExists to return false for symlink")
	}
}

func TestFileErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, fileFailCases())
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
			wantContains: []string{"unable to find file"},
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
