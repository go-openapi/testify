// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package fdleak

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

// Kind classifies an open file descriptor independently of platform-specific target-string conventions.
type Kind int

const (
	// KindUnknown is used when the descriptor kind could not be determined.
	KindUnknown Kind = iota

	// KindFile denotes anything backed by a path in the file system:
	// regular files, directories, symlinks, block devices.
	KindFile

	// KindSocket denotes a socket (AF_UNIX, AF_INET, …).
	KindSocket

	// KindPipe denotes a pipe or FIFO.
	KindPipe

	// KindChar denotes a character device without a resolvable path
	// (fallback for darwin when F_GETPATH fails on a char device).
	KindChar

	// KindOther covers kernel-level descriptors that are not directly
	// opened by user code: Linux anon_inode (epoll, timerfd, …), darwin
	// kqueue, and similar.
	KindOther
)

// FDInfo describes an open file descriptor.
//
// It should remain a human-readable description: a file-system path for vnode-backed FDs,
// or a synthetic label such as "socket:[<inode>]" for non-vnode kinds.
//
// [Kind] is the authoritative classification; use it rather than parsing Target when filtering.
type FDInfo struct {
	FD     int
	Kind   Kind
	Target string
}

// isFiltered reports whether this FD should be excluded from leak reports.
// Sockets, pipes, and other kernel-internal descriptors are filtered by default,
// because they are typically managed by the Go runtime or external libraries
// and their lifecycles do not correlate with user-level resource management.
func (f FDInfo) isFiltered() bool {
	switch f.Kind {
	case KindSocket, KindPipe, KindOther:
		return true
	default:
		return false
	}
}

// snapshotMu serializes Leaked calls to prevent false positives
// from concurrent tests.
var snapshotMu sync.Mutex //nolint:gochecknoglobals // serializes process-wide FD table access

// Snapshot returns a map of currently open file descriptors for the
// running process.
//
// The set of supported platforms is determined at build time; see the
// per-platform implementations (fdleak_linux.go, fdleak_darwin.go).
//
// On unsupported platforms, Snapshot returns an error.
//
// FDs that close between enumeration and resolution are silently skipped.
func Snapshot() (map[int]FDInfo, error) {
	return snapshot()
}

// Leaked takes a before/after snapshot around the tested function
// and returns a formatted description of leaked file descriptors.
//
// Returns the empty string if no leaks are found.
// On unsupported platforms, Leaked returns an error.
func Leaked(tested func()) (string, error) {
	snapshotMu.Lock()
	defer snapshotMu.Unlock()

	before, err := snapshot()
	if err != nil {
		return "", err
	}

	tested()

	after, err := snapshot()
	if err != nil {
		return "", err
	}

	leaked := Diff(before, after)

	return FormatLeaked(leaked), nil
}

// Diff returns file descriptors present in after but not in before,
// excluding filtered FD kinds (sockets, pipes, other kernel descriptors).
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
