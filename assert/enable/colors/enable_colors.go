// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package colors is an indirection to handle colorized output.
//
// This package allows the builder to override the indirection with an alternative implementation
// of colorized printers.
package colors

import (
	colorstub "github.com/go-openapi/testify/v2/internal/assertions/enable/colors"
)

// Enable registers colorized options for pretty-printing the output of assertions.
//
// The provided enabler defers the initialization, so we may retrieve flags after command line parsing
// or other initialization tasks.
//
// This is not intended for concurrent use.
func Enable(enabler func() []Option) {
	colorstub.Enable(enabler)
}

// re-exposed internal types.
type (
	// Option is a colorization option.
	Option = colorstub.Option

	// Theme is a colorization theme for testify output.
	Theme = colorstub.Theme
)

// WithEnable enables colorization.
func WithEnable(enabled bool) Option {
	return colorstub.WithEnable(enabled)
}

// WithSanitizedTheme sets a colorization theme from a string.
func WithSanitizedTheme(theme string) Option {
	return colorstub.WithSanitizedTheme(theme)
}

// WithTheme sets a colorization theme.
func WithTheme(theme Theme) Option {
	return colorstub.WithTheme(theme)
}

// WithDark sets the [ThemeDark] color theme.
func WithDark() Option {
	return colorstub.WithDark()
}

// WithLight sets the [ThemeLight] color theme.
func WithLight() Option {
	return colorstub.WithLight()
}
