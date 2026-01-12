// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import "github.com/go-openapi/testify/v2/internal/difflib"

// colorizers holds string colorizers for a theme.
type colorizers struct {
	expected StringColorizer
	actual   StringColorizer
}

// makeDiffOptions transforms preset options into the
// detailed options for difflib.
func makeDiffOptions(o options) *difflib.Options {
	if !o.enabled {
		return nil
	}

	switch o.theme {
	case ThemeLight:
		return &difflib.Options{
			EqualPrinter:  greenPrinterBuilder,
			DeletePrinter: redPrinterBuilder,
			UpdatePrinter: cyanPrinterBuilder,
			InsertPrinter: yellowPrinterBuilder,
		}
	case ThemeDark:
		return &difflib.Options{
			EqualPrinter:  brightGreenPrinterBuilder,
			DeletePrinter: brightRedPrinterBuilder,
			UpdatePrinter: brightCyanPrinterBuilder,
			InsertPrinter: brightYellowPrinterBuilder,
		}
	default:
		return nil
	}
}

// setColorizers returns string colorizers for the given options.
func setColorizers(o options) colorizers {
	if !o.enabled {
		return colorizers{
			expected: noopColorizer,
			actual:   noopColorizer,
		}
	}

	switch o.theme {
	case ThemeLight:
		return colorizers{
			expected: greenColorizer,
			actual:   redColorizer,
		}
	case ThemeDark:
		return colorizers{
			expected: brightGreenColorizer,
			actual:   brightRedColorizer,
		}
	default:
		return colorizers{
			expected: noopColorizer,
			actual:   noopColorizer,
		}
	}
}
