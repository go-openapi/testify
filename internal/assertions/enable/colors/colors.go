// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package colors

import (
	"bufio"

	"github.com/go-openapi/testify/v2/internal/difflib"
)

const (
	greenMark  = "\033[0;32m"
	redMark    = "\033[0;31m"
	yellowMark = "\033[0;33m" // aka orange
	cyanMark   = "\033[0;36m"

	brightGreenMark  = "\033[0;92m"
	brightRedMark    = "\033[0;91m"
	brightYellowMark = "\033[0;93m" // aka yellow
	brightCyanMark   = "\033[0;96m" // aka turquoise

	// color codes for future use.

	// blackMark   = "\033[0;30m"
	// blueMark    = "\033[0;34m"
	// magentaMark = "\033[0;35m"
	// greyMark    = "\033[0;37m".

	// darkGreyMark      = "\033[0;90m"
	// brightBlueMark    = "\033[0;94m"
	// brightMagentaMark = "\033[0;95m"
	// brightWhiteMark   = "\033[0;97m".

	endMark = "\033[0m"
)

// StringColorizer wraps a string with ANSI escape codes.
//
// This is a simpler alternative to [difflib.PrinterBuilder] for cases
// where streaming to a [bufio.Writer] is not needed.
type StringColorizer func(string) string

func makeColorizer(mark string) StringColorizer {
	return func(s string) string {
		return mark + s + endMark
	}
}

// noopColorizer returns the input string unchanged.
func noopColorizer(s string) string {
	return s
}

//nolint:gochecknoglobals // internal colorizers may safely be shared at the package-level
var (
	greenColorizer = makeColorizer(greenMark)
	redColorizer   = makeColorizer(redMark)
	// yellowColorizer = makeColorizer(yellowMark)
	// cyanColorizer   = makeColorizer(cyanMark).

	brightGreenColorizer = makeColorizer(brightGreenMark)
	brightRedColorizer   = makeColorizer(brightRedMark)
	// brightYellowColorizer = makeColorizer(brightYellowMark)
	// brightCyanColorizer   = makeColorizer(brightCyanMark).
)

//nolint:gochecknoglobals // internal printer builders may safely be shared at the package-level
var (
	greenPrinterBuilder  = ansiPrinterBuilder(greenMark)
	redPrinterBuilder    = ansiPrinterBuilder(redMark)
	yellowPrinterBuilder = ansiPrinterBuilder(yellowMark)
	cyanPrinterBuilder   = ansiPrinterBuilder(cyanMark)

	brightGreenPrinterBuilder  = ansiPrinterBuilder(brightGreenMark)
	brightRedPrinterBuilder    = ansiPrinterBuilder(brightRedMark)
	brightYellowPrinterBuilder = ansiPrinterBuilder(brightYellowMark)
	brightCyanPrinterBuilder   = ansiPrinterBuilder(brightCyanMark)

	// magentaPrinterBuilder = ansiPrinterBuilder(magentaMark)
	// bluePrinterBuilder    = ansiPrinterBuilder(blueMark).

	// brightMagentaPrinterBuilder = ansiPrinterBuilder(brightMagentaMark)
	// brightBluePrinterBuilder    = ansiPrinterBuilder(brightBlueMark).
)

func ansiPrinterBuilder(mark string) difflib.PrinterBuilder {
	return func(w *bufio.Writer) difflib.Printer {
		return func(str string) (err error) {
			_, err = w.WriteString(mark)
			if err != nil {
				return
			}
			_, err = w.WriteString(str)
			if err != nil {
				return
			}
			_, err = w.WriteString(endMark)
			if err != nil {
				return
			}

			return nil
		}
	}
}
