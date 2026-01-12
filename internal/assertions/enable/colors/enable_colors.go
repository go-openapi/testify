// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"sync"

	"github.com/go-openapi/testify/v2/internal/difflib"
)

//nolint:gochecknoglobals // in this particular case, we need a global to enable the feature from another module
var (
	resolveOptionsOnce sync.Once
	optionsEnabler     func() []Option
	colorOptions       *difflib.Options
	stringColorizers   colorizers
)

// Enable registers colorized options for pretty-printing the output of assertions.
//
// The argument passed is a function that is executed after package initialization and CLI arg parsing.
//
// This is not intended for concurrent use as it sets a package-level state.
func Enable(enabler func() []Option) {
	optionsEnabler = enabler
}

// Enabled indicates if a global color options setting has been enabled.
func Enabled() bool {
	return optionsEnabler != nil
}

// Options returns the colorization options for [difflib].
//
// It yields nil if the colorization feature is not enabled (non blocking: colors just won't display).
func Options() *difflib.Options {
	resolveOptions()

	return colorOptions
}

// ExpectedColorizer returns a colorizer for expected values.
//
// It returns a no-op colorizer if colorization is not enabled.
func ExpectedColorizer() StringColorizer {
	resolveOptions()

	return stringColorizers.expected
}

// ActualColorizer returns a colorizer for actual values.
//
// It returns a no-op colorizer if colorization is not enabled.
func ActualColorizer() StringColorizer {
	resolveOptions()

	return stringColorizers.actual
}

func resolveOptions() {
	resolveOptionsOnce.Do(func() {
		// defers the resolution of options until first usage
		if optionsEnabler == nil {
			stringColorizers = colorizers{
				expected: noopColorizer,
				actual:   noopColorizer,
			}

			return
		}

		o := optionsWithDefaults(optionsEnabler())
		colorOptions = makeDiffOptions(o)
		stringColorizers = setColorizers(o)
	})
}
