// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build darwin

package fdleak

import (
	"fmt"
	"math"
	"syscall"
	"unsafe"
)

const (
	// fGetPath is the darwin fcntl command to retrieve the vnode path of a
	// file descriptor (see <sys/fcntl.h>).
	fGetPath = 50

	// maxPathLen is darwin's MAXPATHLEN.
	maxPathLen = 1024

	// fdProbeCap bounds the linear FD scan. Typical darwin RLIMIT_NOFILE soft
	// limits range from 256 to 10240; real test processes rarely hold more than
	// a few dozen FDs, so a 4096 ceiling caps the cost at a few thousand fstat
	// syscalls per snapshot (single-digit milliseconds).
	fdProbeCap = uint64(4096)
)

// snapshot enumerates open file descriptors on darwin by probing FDs 0..N
// with fstat. Unlike Linux's /proc/self/fd, darwin's /dev/fd is a devfs mount
// that does not cooperate with Go's os.ReadDir fallback to fstatat, so we
// bypass directory enumeration and probe FD numbers directly.
//
// For each live FD:
//
//   - fcntl(F_GETPATH) resolves vnode-backed paths (regular files, devices).
//   - fstat classifies non-vnode kinds (sockets, pipes, kqueues, …) and
//     yields a synthetic target string compatible with the Linux output.
//
// FDs that close between fstat and fcntl are silently skipped.
func snapshot() (map[int]FDInfo, error) {
	var rlim syscall.Rlimit
	if err := syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
		return nil, fmt.Errorf("getrlimit(RLIMIT_NOFILE): %w", err)
	}

	limit := int(min(rlim.Cur, fdProbeCap, uint64(math.MaxInt))) //nolint:gosec // the min guarantees that limit doesn't overflow int

	fds := make(map[int]FDInfo)
	for fd := range limit {
		info, ok := resolveFD(fd)
		if !ok {
			continue
		}

		fds[fd] = info
	}

	return fds, nil
}

// resolveFD builds an FDInfo for fd. It returns false when the FD is not open
// (fstat returns EBADF) or the FD was closed mid-probe.
func resolveFD(fd int) (FDInfo, bool) {
	// Classify first via fstat: this is also our liveness check (EBADF means
	// "not open"). Querying F_GETPATH on a closed FD returns a similar error,
	// but ordering fstat first keeps the closed-FD fast path to a single
	// syscall.
	var stat syscall.Stat_t
	if err := syscall.Fstat(fd, &stat); err != nil {
		return FDInfo{}, false
	}

	switch stat.Mode & syscall.S_IFMT {
	case syscall.S_IFSOCK:
		return FDInfo{FD: fd, Kind: KindSocket, Target: fmt.Sprintf("socket:[%d]", stat.Ino)}, true
	case syscall.S_IFIFO:
		return FDInfo{FD: fd, Kind: KindPipe, Target: fmt.Sprintf("pipe:[%d]", stat.Ino)}, true
	}

	// Remaining kinds are vnode-backed; try to resolve the path.
	path, pathErr := fcntlGetPath(fd)

	switch stat.Mode & syscall.S_IFMT {
	case syscall.S_IFCHR:
		if pathErr == nil {
			return FDInfo{FD: fd, Kind: KindChar, Target: path}, true
		}

		return FDInfo{FD: fd, Kind: KindChar, Target: fmt.Sprintf("char:[%d]", stat.Rdev)}, true
	case syscall.S_IFREG, syscall.S_IFDIR, syscall.S_IFLNK, syscall.S_IFBLK:
		if pathErr == nil {
			return FDInfo{FD: fd, Kind: KindFile, Target: path}, true
		}

		return FDInfo{FD: fd, Kind: KindFile, Target: fmt.Sprintf("file:[%d]", stat.Ino)}, true
	default:
		// kqueue, fsevents, netpolicy, pshm, psem, and other kernel-internal
		// descriptors land here. They never expose a vnode path.
		return FDInfo{FD: fd, Kind: KindOther, Target: fmt.Sprintf("other:[%d]", stat.Ino)}, true
	}
}

// fcntlGetPath issues fcntl(fd, F_GETPATH, buf) and returns the resolved path.
//
// The unsafe.Pointer → uintptr conversion happens at the syscall.Syscall call
// site so Go's runtime keeps buf pinned for the duration of the call (the
// documented "Case 4" rule for syscall arguments).
func fcntlGetPath(fd int) (string, error) {
	var buf [maxPathLen]byte

	_, _, errno := syscall.Syscall(
		syscall.SYS_FCNTL,
		uintptr(fd), //nolint:gosec // fd is a int and it may be stored in a uintptr by definition of the int type.
		uintptr(fGetPath),
		uintptr(unsafe.Pointer(&buf[0])), // F_GETPATH requires a raw buffer pointer; lifetime is pinned across the syscall.
	)
	if errno != 0 {
		return "", errno
	}

	n := 0
	for n < len(buf) && buf[n] != 0 {
		n++
	}

	return string(buf[:n]), nil
}
