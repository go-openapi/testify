// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build !linux && !darwin

package fdleak

import (
	"errors"
	"runtime"
)

// snapshot returns an error on platforms without a supported implementation.
func snapshot() (map[int]FDInfo, error) {
	return nil, errors.New("file descriptor leak detection is not supported on " + runtime.GOOS)
}
