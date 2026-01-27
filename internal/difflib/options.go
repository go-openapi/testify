// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package difflib

import (
	"bufio"
	"fmt"
)

type (
	// Formatter is a formatting function like [fmt.Printf].
	Formatter func(format string, args ...any) error

	// Printer outputs a formatted string, e.g. to some underlying writer.
	Printer func(string) error
)

type (
	// FormatterBuilder is a function that builds a [Formatter] given a bufferized [bufio.Writer].
	FormatterBuilder func(*bufio.Writer) Formatter

	// PrinterBuilder is a function that builds a [Printer] given a bufferized [bufio.Writer].
	PrinterBuilder func(*bufio.Writer) Printer
)

// Options to fine-tune the rendering of the diff output.
type Options struct {
	EqualPrinter  PrinterBuilder
	DeletePrinter PrinterBuilder
	UpdatePrinter PrinterBuilder
	InsertPrinter PrinterBuilder
	OtherPrinter  PrinterBuilder
	Formatter     FormatterBuilder
}

func optionsWithDefaults(in *Options) *Options {
	o := &Options{
		EqualPrinter:  defaultPrinterBuilder,
		DeletePrinter: defaultPrinterBuilder,
		UpdatePrinter: defaultPrinterBuilder,
		InsertPrinter: defaultPrinterBuilder,
		OtherPrinter:  defaultPrinterBuilder,
		Formatter:     defaultFormatterBuilder,
	}

	if in == nil {
		return o
	}
	if in.EqualPrinter != nil {
		o.EqualPrinter = in.EqualPrinter
	}
	if in.DeletePrinter != nil {
		o.DeletePrinter = in.DeletePrinter
	}
	if in.UpdatePrinter != nil {
		o.UpdatePrinter = in.UpdatePrinter
	}
	if in.InsertPrinter != nil {
		o.InsertPrinter = in.InsertPrinter
	}
	if in.OtherPrinter != nil {
		o.OtherPrinter = in.OtherPrinter
	}
	if in.Formatter != nil {
		o.Formatter = in.Formatter
	}

	return o
}

func DefaultPrinterBuilder(buf *bufio.Writer) Printer {
	return defaultPrinterBuilder(buf)
}

func defaultPrinterBuilder(buf *bufio.Writer) Printer {
	return func(s string) error {
		_, err := buf.WriteString(s)
		return err
	}
}

func defaultFormatterBuilder(buf *bufio.Writer) Formatter {
	return func(format string, args ...any) error {
		_, err := fmt.Fprintf(buf, format, args...)
		return err
	}
}
