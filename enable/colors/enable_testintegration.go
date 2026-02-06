// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build testcolorized

package colors

// integrationTestHatch is specifically designed to force colorized settings in integration tests.
// When build tag "testcolorized" is defined, it returns true and colorization is enabled, regardless
// of command line arguments or environment.
func integrationTestHatch() bool {
	return true
}
