// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package fdleak provides file descriptor leak detection.
//
// It uses /proc/self/fd snapshots on Linux to take a snapshot
// of open file descriptors before and after
// running the tested function. Any file descriptors present in the
// "after" snapshot but not in the "before" snapshot are considered leaks.
//
// By default, sockets, pipes, and anonymous inodes are filtered out,
// as these are typically managed by the Go runtime or OS internals.
//
// This approach is inherently process-wide: /proc/self/fd lists all
// file descriptors for the process. Any concurrent I/O from other
// goroutines may cause false positives. A mutex serializes [Leaked]
// calls to prevent multiple leak checks from interfering with each
// other, but cannot protect against external concurrent file operations.
//
// This package only works on Linux. On other platforms,
// [Snapshot] returns an error.
package fdleak
