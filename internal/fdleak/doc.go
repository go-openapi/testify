// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package fdleak provides file descriptor leak detection.
//
// It takes a snapshot of open file descriptors before and after running the tested function.
//
// Any file descriptors present in the "after" snapshot but not in the "before" snapshot
// — and not of a filtered [Kind] — are considered leaks.
//
// # Platform support
//
//   - Linux:   enumerates /proc/self/fd and classifies FDs from the
//     readlink target (socket:/pipe:/anon_inode:/ path).
//   - darwin:  enumerates /dev/fd and resolves each FD via fcntl(F_GETPATH),
//     falling back to fstat to classify sockets, pipes and kqueues.
//   - other:   [Snapshot] returns an error.
//
// # Filtering
//
// Sockets, pipes and other kernel-internal descriptors (Linux anon_inode,
// darwin kqueue) are excluded from leak reports by default, as these are
// typically managed by the Go runtime or external libraries.
//
// # Concurrency
//
// This approach is inherently process-wide: the FD table lists all file descriptors for the process.
//
// Any concurrent I/O from other goroutines may cause false positives.
// A mutex serializes [Leaked] calls to prevent multiple leak checks from interfering with each other,
// but cannot protect against external concurrent file operations.
package fdleak
