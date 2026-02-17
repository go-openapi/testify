// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package fdleak

import (
	"errors"
	"fmt"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
)

// FDInfo describes an open file descriptor.
type FDInfo struct {
	FD     int
	Target string // readlink target (e.g. "/tmp/foo.txt", "socket:[12345]")
}

// isFiltered returns true if this FD should be excluded from leak reports.
// Sockets, pipes, and anonymous inodes are filtered out by default.
func (f FDInfo) isFiltered() bool {
	return strings.HasPrefix(f.Target, "socket:[") ||
		strings.HasPrefix(f.Target, "pipe:[") ||
		strings.HasPrefix(f.Target, "anon_inode:[")
}

// snapshotMu serializes Leaked calls to prevent false positives
// from concurrent tests.
var snapshotMu sync.Mutex //nolint:gochecknoglobals // serializes process-wide /proc/self/fd access

const procSelfFD = "/proc/self/fd"

// Snapshot reads /proc/self/fd and returns a map of currently open file descriptors.
//
// FDs that close between ReadDir and Readlink are silently skipped.
// Returns an error if not running on Linux.
func Snapshot() (map[int]FDInfo, error) {
	if runtime.GOOS != "linux" {
		return nil, errors.New("file descriptor leak detection requires Linux (/proc/self/fd)")
	}

	entries, err := os.ReadDir(procSelfFD)
	if err != nil {
		return nil, fmt.Errorf("reading %s: %w", procSelfFD, err)
	}

	fds := make(map[int]FDInfo, len(entries))
	for _, e := range entries {
		fd, err := strconv.Atoi(e.Name())
		if err != nil {
			continue
		}

		target, err := os.Readlink(procSelfFD + "/" + e.Name())
		if err != nil {
			continue // FD closed between ReadDir and Readlink
		}

		fds[fd] = FDInfo{FD: fd, Target: target}
	}

	return fds, nil
}

// Leaked takes a before/after snapshot around the tested function
// and returns a formatted description of leaked file descriptors.
//
// Returns the empty string if no leaks are found.
// The caller is responsible for checking [runtime.GOOS] before calling.
func Leaked(tested func()) (string, error) {
	snapshotMu.Lock()
	defer snapshotMu.Unlock()

	before, err := Snapshot()
	if err != nil {
		return "", err
	}

	tested()

	after, err := Snapshot()
	if err != nil {
		return "", err
	}

	leaked := Diff(before, after)

	return FormatLeaked(leaked), nil
}

// Diff returns file descriptors present in after but not in before,
// excluding filtered FD types (sockets, pipes, anonymous inodes).
func Diff(before, after map[int]FDInfo) []FDInfo {
	var leaked []FDInfo

	for fd, info := range after {
		if _, existed := before[fd]; existed {
			continue
		}

		if info.isFiltered() {
			continue
		}

		leaked = append(leaked, info)
	}

	sort.Slice(leaked, func(i, j int) bool {
		return leaked[i].FD < leaked[j].FD
	})

	return leaked
}

// FormatLeaked formats leaked file descriptors into a human-readable message.
// Returns the empty string if the slice is empty.
func FormatLeaked(leaked []FDInfo) string {
	if len(leaked) == 0 {
		return ""
	}

	var b strings.Builder

	fmt.Fprintf(&b, "found %d leaked file descriptor(s):\n", len(leaked))
	for _, fd := range leaked {
		fmt.Fprintf(&b, "  fd %d: %s\n", fd.FD, fd.Target)
	}

	return b.String()
}
