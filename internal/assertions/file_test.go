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
	True(t, FileExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(mockT)
	False(t, FileExists(mock, "random_file"))

	mock = new(mockT)
	False(t, FileExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	True(t, FileExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	True(t, FileExists(mock, link))
}

func TestFileFileNotExists(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	False(t, FileNotExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(mockT)
	True(t, FileNotExists(mock, "non_existent_file"))

	mock = new(mockT)
	True(t, FileNotExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	False(t, FileNotExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	False(t, FileNotExists(mock, link))
}

func TestFileDirExists(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	False(t, DirExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(mockT)
	False(t, DirExists(mock, "non_existent_dir"))

	mock = new(mockT)
	True(t, DirExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	False(t, DirExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_dir")
	mock = new(mockT)
	False(t, DirExists(mock, link))
}

func TestFileDirNotExists(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	True(t, DirNotExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(mockT)
	True(t, DirNotExists(mock, "non_existent_dir"))

	mock = new(mockT)
	False(t, DirNotExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	True(t, DirNotExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_dir")
	mock = new(mockT)
	True(t, DirNotExists(mock, link))
}

func TestFileEmpty(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	True(t, FileEmpty(mock, filepath.Join("testdata", "empty_file")))

	mock = new(mockT)
	False(t, FileEmpty(mock, filepath.Join("testdata", "existing_file")))

	mock = new(mockT)
	False(t, FileEmpty(mock, "random_file"))

	mock = new(mockT)
	False(t, FileEmpty(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "empty_file"))
	mock = new(mockT)
	True(t, FileEmpty(mock, link))

	link = getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	False(t, FileEmpty(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	False(t, FileEmpty(mock, link))
}

func TestFileNotEmpty(t *testing.T) {
	t.Parallel()

	mock := new(mockT)
	True(t, FileNotEmpty(mock, filepath.Join("testdata", "existing_file")))

	mock = new(mockT)
	False(t, FileNotEmpty(mock, filepath.Join("testdata", "empty_file")))

	mock = new(mockT)
	False(t, FileNotEmpty(mock, "non_existent_file"))

	mock = new(mockT)
	False(t, FileNotEmpty(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "empty_file"))
	mock = new(mockT)
	False(t, FileNotEmpty(mock, link))

	link = getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(mockT)
	True(t, FileNotEmpty(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(mockT)
	False(t, FileNotExists(mock, link))
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
