// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

// Theme is a colorization theme for testify output.
type Theme string

func (t Theme) String() string {
	return string(t)
}

const (
	// ThemeLight uses normal ANSI colors.
	ThemeLight Theme = "light"

	// ThemeDark uses bright ANSI colors.
	ThemeDark Theme = "dark"
)

// Option is a colorization option.
type Option func(o *options)

type options struct {
	enabled bool
	theme   Theme
}

// WithSanitizedTheme sets a colorization theme from its name or does nothing if the theme is not supported.
func WithSanitizedTheme(theme string) Option {
	th := Theme(theme)
	switch th {
	case ThemeDark, ThemeLight:
		return WithTheme(th)
	default:
		return func(*options) { // noop
		}
	}
}

// WithEnable enables colorized output.
func WithEnable(enabled bool) Option {
	return func(o *options) {
		o.enabled = enabled
	}
}

// WithTheme sets a colorization theme.
func WithTheme(theme Theme) Option {
	return func(o *options) {
		o.theme = theme
	}
}

// WithDark sets the [ThemeDark] color theme.
func WithDark() Option {
	return WithTheme(ThemeDark)
}

// WithLight sets the [ThemeLight] color theme.
func WithLight() Option {
	return WithTheme(ThemeLight)
}

func optionsWithDefaults(opts []Option) options {
	o := options{
		theme: ThemeDark, // default theme
	}

	for _, apply := range opts {
		apply(&o)
	}

	return o
}
