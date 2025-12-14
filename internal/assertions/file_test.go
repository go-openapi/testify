// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFileExists(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	True(t, FileExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(testing.T)
	False(t, FileExists(mock, "random_file"))

	mock = new(testing.T)
	False(t, FileExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(testing.T)
	True(t, FileExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(testing.T)
	True(t, FileExists(mock, link))
}

func TestFileNoFileExists(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, NoFileExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(testing.T)
	True(t, NoFileExists(mock, "non_existent_file"))

	mock = new(testing.T)
	True(t, NoFileExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(testing.T)
	False(t, NoFileExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(testing.T)
	False(t, NoFileExists(mock, link))
}

func TestFileDirExists(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	False(t, DirExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(testing.T)
	False(t, DirExists(mock, "non_existent_dir"))

	mock = new(testing.T)
	True(t, DirExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(testing.T)
	False(t, DirExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_dir")
	mock = new(testing.T)
	False(t, DirExists(mock, link))
}

func TestFileNoDirExists(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	True(t, NoDirExists(mock, filepath.Join("testdata", "existing_file")))

	mock = new(testing.T)
	True(t, NoDirExists(mock, "non_existent_dir"))

	mock = new(testing.T)
	False(t, NoDirExists(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(testing.T)
	True(t, NoDirExists(mock, link))

	link = getTempSymlinkPath(t, "non_existent_dir")
	mock = new(testing.T)
	True(t, NoDirExists(mock, link))
}

func TestFileEmpty(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	True(t, FileEmpty(mock, filepath.Join("testdata", "empty_file")))

	mock = new(testing.T)
	False(t, FileEmpty(mock, filepath.Join("testdata", "existing_file")))

	mock = new(testing.T)
	False(t, FileEmpty(mock, "random_file"))

	mock = new(testing.T)
	False(t, FileEmpty(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "empty_file"))
	mock = new(testing.T)
	True(t, FileEmpty(mock, link))

	link = getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(testing.T)
	False(t, FileEmpty(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(testing.T)
	False(t, FileEmpty(mock, link))
}

func TestFileNotEmpty(t *testing.T) {
	t.Parallel()

	mock := new(testing.T)
	True(t, FileNotEmpty(mock, filepath.Join("testdata", "existing_file")))

	mock = new(testing.T)
	False(t, FileNotEmpty(mock, filepath.Join("testdata", "empty_file")))

	mock = new(testing.T)
	False(t, FileNotEmpty(mock, "non_existent_file"))

	mock = new(testing.T)
	False(t, FileNotEmpty(mock, filepath.Join("testdata", "existing_dir")))

	link := getTempSymlinkPath(t, filepath.Join("testdata", "empty_file"))
	mock = new(testing.T)
	False(t, FileNotEmpty(mock, link))

	link = getTempSymlinkPath(t, filepath.Join("testdata", "existing_file"))
	mock = new(testing.T)
	True(t, FileNotEmpty(mock, link))

	link = getTempSymlinkPath(t, "non_existent_file")
	mock = new(testing.T)
	False(t, NoFileExists(mock, link))
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
