// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package difflib

import (
	"bufio"
	"fmt"
)

type (
	Formatter func(format string, args ...any) error
	Printer   func(string) error
)

type (
	FormatterBuilder func(*bufio.Writer) Formatter
	PrinterBuilder   func(*bufio.Writer) Printer
)

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
