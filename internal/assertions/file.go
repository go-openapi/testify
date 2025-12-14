// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"io/fs"
	"os"
)

// FileExists checks whether a file exists in the given path. It also fails if
// the path points to a directory or there is an error when trying to check the file.
//
// Examples:
//
//	success: filepath.Join(testDataPath(),"existing_file")
//	failure: filepath.Join(testDataPath(),"non_existing_file")
func FileExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Fail(t, fmt.Sprintf("unable to find file %q", path), msgAndArgs...)
		}
		return Fail(t, fmt.Sprintf("error when running os.Lstat(%q): %s", path, err), msgAndArgs...)
	}
	if info.IsDir() {
		return Fail(t, fmt.Sprintf("%q is a directory", path), msgAndArgs...)
	}
	return true
}

// NoFileExists checks whether a file does not exist in a given path. It fails
// if the path points to an existing _file_ only.
//
// Examples:
//
//	success: filepath.Join(testDataPath(),"non_existing_file")
//	failure: filepath.Join(testDataPath(),"existing_file")
func NoFileExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	info, err := os.Lstat(path)
	if err != nil {
		return true
	}
	if info.IsDir() {
		return true
	}
	return Fail(t, fmt.Sprintf("file %q exists", path), msgAndArgs...)
}

// DirExists checks whether a directory exists in the given path. It also fails
// if the path is a file rather a directory or there is an error checking whether it exists.
//
// Examples:
//
//	success: filepath.Join(testDataPath(),"existing_dir")
//	failure: filepath.Join(testDataPath(),"non_existing_dir")
func DirExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Fail(t, fmt.Sprintf("unable to find file %q", path), msgAndArgs...)
		}
		return Fail(t, fmt.Sprintf("error when running os.Lstat(%q): %s", path, err), msgAndArgs...)
	}
	if !info.IsDir() {
		return Fail(t, fmt.Sprintf("%q is a file", path), msgAndArgs...)
	}
	return true
}

// NoDirExists checks whether a directory does not exist in the given path.
// It fails if the path points to an existing _directory_ only.
// Examples:
//
//	success: filepath.Join(testDataPath(),"non_existing_dir")
//	failure: filepath.Join(testDataPath(),"existing_dir")
func NoDirExists(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true
		}
		return true
	}
	if !info.IsDir() {
		return true
	}
	return Fail(t, fmt.Sprintf("directory %q exists", path), msgAndArgs...)
}

// FileEmpty checks whether a file exists in the given path and is empty.
// It fails if the file is not empty, if the path points to a directory or there is an error when trying to check the file.
//
// Examples:
//
//	success: filepath.Join(testDataPath(),"empty_file")
//	failure: filepath.Join(testDataPath(),"existing_file")
func FileEmpty(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Fail(t, fmt.Sprintf("unable to find file %q", path), msgAndArgs...)
		}
		return Fail(t, fmt.Sprintf("error when running os.Lstat(%q): %s", path, err), msgAndArgs...)
	}
	if info.IsDir() {
		return Fail(t, fmt.Sprintf("%q is a directory", path), msgAndArgs...)
	}
	if info.Mode()&fs.ModeSymlink > 0 {
		target, err := os.Readlink(path)
		if err != nil {
			return Fail(t, fmt.Sprintf("could not resolve symlink %q", path), msgAndArgs...)
		}
		return FileEmpty(t, target, msgAndArgs...)
	}

	if info.Size() > 0 {
		return Fail(t, fmt.Sprintf("%q is not empty", path), msgAndArgs...)
	}

	return true
}

// FileNotEmpty checks whether a file exists in the given path and is not empty.
// It fails if the file is empty, if the path points to a directory or there is an error when trying to check the file.
//
// Examples:
//
//	success: filepath.Join(testDataPath(),"existing_file")
//	failure: filepath.Join(testDataPath(),"empty_file")
func FileNotEmpty(t T, path string, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	info, err := os.Lstat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Fail(t, fmt.Sprintf("unable to find file %q", path), msgAndArgs...)
		}
		return Fail(t, fmt.Sprintf("error when running os.Lstat(%q): %s", path, err), msgAndArgs...)
	}
	if info.IsDir() {
		return Fail(t, fmt.Sprintf("%q is a directory", path), msgAndArgs...)
	}
	if info.Mode()&fs.ModeSymlink > 0 {
		target, err := os.Readlink(path)
		if err != nil {
			return Fail(t, fmt.Sprintf("could not resolve symlink %q", path), msgAndArgs...)
		}
		return FileNotEmpty(t, target, msgAndArgs...)
	}

	if info.Size() == 0 {
		return Fail(t, fmt.Sprintf("%q is empty", path), msgAndArgs...)
	}

	return true
}
