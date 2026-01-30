// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

//go:build !testcolorized

package colors

// integrationTestHatch is specifically designed to force colorized settings in integration tests.
// When no build tag "testcolorized" is defined (the default), it returns false and is a noop.
func integrationTestHatch() bool {
	return false
}
