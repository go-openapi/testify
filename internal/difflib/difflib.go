// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package difflib

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// SplitLines splits a string on "\n" while preserving them. The output can be used
// as input for UnifiedDiff and ContextDiff structures.
func SplitLines(s string) []string {
	lines := strings.SplitAfter(s, "\n")
	lines[len(lines)-1] += "\n"
	return lines
}

// UnifiedDiff holds the unified diff parameters.
type UnifiedDiff struct {
	*Options

	A        []string // First sequence lines
	FromFile string   // First file name
	FromDate string   // First file time
	B        []string // Second sequence lines
	ToFile   string   // Second file name
	ToDate   string   // Second file time
	Eol      string   // Headers end of line, defaults to LF
	Context  int      // Number of context lines

	wsE Printer
	wsD Printer
	wsU Printer
	wsI Printer
	wsO Printer
	wf  Formatter
}

func (u *UnifiedDiff) applyWriter(buf *bufio.Writer) {
	u.wsE = u.EqualPrinter(buf)
	u.wsD = u.DeletePrinter(buf)
	u.wsU = u.UpdatePrinter(buf)
	u.wsI = u.InsertPrinter(buf)
	u.wsO = u.OtherPrinter(buf)
	u.wf = u.Formatter(buf)
}

// GetUnifiedDiffString is like WriteUnifiedDiff but returns the diff a string.
func GetUnifiedDiffString(diff UnifiedDiff) (string, error) {
	w := new(bytes.Buffer)
	err := WriteUnifiedDiff(w, diff)

	return w.String(), err
}

// WriteUnifiedDiff write the comparison between two sequences of lines.
// It generates the delta as a unified diff.
//
// Unified diffs are a compact way of showing line changes and a few
// lines of context.  The number of context lines is set by 'n' which
// defaults to three.
//
// By default, the diff control lines (those with ---, +++, or @@) are
// created with a trailing newline.  This is helpful so that inputs
// created from file.readlines() result in diffs that are suitable for
// file.writelines() since both the inputs and outputs have trailing
// newlines.
//
// For inputs that do not have trailing newlines, set the lineterm
// argument to "" so that the output will be uniformly newline free.
//
// The unidiff format normally has a header for filenames and modification
// 'fromfile', 'tofile', 'fromfiledate', and 'tofiledate'.
// The modification times are normally expressed in the ISO 8601 format.
func WriteUnifiedDiff(writer io.Writer, diff UnifiedDiff) error {
	diff.Options = optionsWithDefaults(diff.Options)
	buf := bufio.NewWriter(writer)
	defer buf.Flush()

	diff.applyWriter(buf)

	if len(diff.Eol) == 0 {
		diff.Eol = "\n"
	}

	m := NewMatcher(diff.A, diff.B)
	groups := m.GetGroupedOpCodes(diff.Context)

	if len(groups) == 0 {
		return nil
	}

	if err := writeFirstGroup(groups[0], diff); err != nil {
		return err
	}

	for _, g := range groups[1:] {
		if err := writeGroup(g, diff); err != nil {
			return err
		}
	}

	return nil
}

func writeFirstGroup(g []OpCode, diff UnifiedDiff) error {
	fromDate := ""
	if len(diff.FromDate) > 0 {
		fromDate = "\t" + diff.FromDate
	}

	toDate := ""
	if len(diff.ToDate) > 0 {
		toDate = "\t" + diff.ToDate
	}

	if diff.FromFile != "" || diff.ToFile != "" {
		err := diff.wf("--- %s%s%s", diff.FromFile, fromDate, diff.Eol)
		if err != nil {
			return err
		}
		err = diff.wf("+++ %s%s%s", diff.ToFile, toDate, diff.Eol)
		if err != nil {
			return err
		}
	}

	return writeGroup(g, diff)
}

func writeGroup(group []OpCode, diff UnifiedDiff) error {
	first, last := group[0], group[len(group)-1]
	range1 := formatRangeUnified(first.I1, last.I2)
	range2 := formatRangeUnified(first.J1, last.J2)
	if err := diff.wf("@@ -%s +%s @@%s", range1, range2, diff.Eol); err != nil {
		return err
	}

	for _, c := range group {
		// 'r' (replace)
		// 'd' (delete)
		// 'i' (insert)
		// 'e' (equal)
		if c.Tag != 'r' && c.Tag != 'e' && c.Tag != 'd' && c.Tag != 'i' {
			continue
		}

		i1, i2, j1, j2 := c.I1, c.I2, c.J1, c.J2
		switch c.Tag {
		case 'e':
			if err := writeEqual(diff.A[i1:i2], diff.wsE); err != nil {
				return err
			}
		case 'd':
			if err := writeReplaceOrDelete(diff.A[i1:i2], diff.wsD); err != nil {
				return err
			}
		case 'i':
			if err := writeReplaceOrInsert(diff.B[j1:j2], diff.wsI); err != nil {
				return err
			}
		case 'r':
			if err := writeReplaceOrDelete(diff.A[i1:i2], diff.wsD); err != nil {
				return err
			}
			if err := writeReplaceOrInsert(diff.B[j1:j2], diff.wsI); err != nil {
				return err
			}
		}
	}

	return nil
}

func writeEqual(lines []string, ws Printer) error {
	for _, line := range lines {
		if err := ws(" " + line); err != nil {
			return err
		}
	}

	return nil
}

func writeReplaceOrDelete(lines []string, ws Printer) error {
	for _, line := range lines {
		if err := ws("-" + line); err != nil {
			return err
		}
	}

	return nil
}

func writeReplaceOrInsert(lines []string, ws Printer) error {
	for _, line := range lines {
		if err := ws("+" + line); err != nil {
			return err
		}
	}

	return nil
}

// Convert range to the "ed" format.
func formatRangeContext(start, stop int) string {
	// Per the diff spec at http://www.unix.org/single_unix_specification/
	beginning := start + 1 // lines start numbering with one
	length := stop - start
	if length == 0 {
		beginning-- // empty ranges begin at line just before the range
	}
	if length <= 1 {
		return strconv.Itoa(beginning)
	}
	return fmt.Sprintf("%d,%d", beginning, beginning+length-1)
}

// Convert range to the "ed" format.
func formatRangeUnified(start, stop int) string {
	// Per the diff spec at http://www.unix.org/single_unix_specification/
	beginning := start + 1 // lines start numbering with one
	length := stop - start
	if length == 1 {
		return strconv.Itoa(beginning)
	}
	if length == 0 {
		beginning-- // empty ranges begin at line just before the range
	}
	return fmt.Sprintf("%d,%d", beginning, length)
}
