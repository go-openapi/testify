// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build linux

package fdleak

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const procSelfFD = "/proc/self/fd"

// snapshot enumerates open file descriptors via /proc/self/fd.
//
// The readlink target carries the kind on Linux:
//   - "socket:[<inode>]"     → KindSocket
//   - "pipe:[<inode>]"       → KindPipe
//   - "anon_inode:[<label>]" → KindOther (epoll, timerfd, signalfd, …)
//   - anything else          → KindFile (regular file, directory, device)
//
// FDs that close between ReadDir and Readlink are silently skipped.
func snapshot() (map[int]FDInfo, error) {
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

		fds[fd] = FDInfo{FD: fd, Kind: kindFromLinuxTarget(target), Target: target}
	}

	return fds, nil
}

func kindFromLinuxTarget(target string) Kind {
	switch {
	case strings.HasPrefix(target, "socket:["):
		return KindSocket
	case strings.HasPrefix(target, "pipe:["):
		return KindPipe
	case strings.HasPrefix(target, "anon_inode:"):
		return KindOther
	default:
		return KindFile
	}
}
