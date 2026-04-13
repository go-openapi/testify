// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package difflib

import (
	"bytes"
	"strings"
	"testing"
)

func TestSplitLines(t *testing.T) {
	lines := SplitLines("a\nb\nc\n")
	if len(lines) != 4 {
		t.Fatalf("expected 4 lines, got %d: %q", len(lines), lines)
	}
	if lines[0] != "a\n" || lines[1] != "b\n" || lines[2] != "c\n" {
		t.Errorf("unexpected split result: %q", lines)
	}
}

func TestSplitLinesEmpty(t *testing.T) {
	lines := SplitLines("")
	if len(lines) != 1 {
		t.Errorf("expected 1 line for empty string, got %d: %q", len(lines), lines)
	}
}

func TestGetUnifiedDiffString(t *testing.T) {
	diff := UnifiedDiff{
		A:        SplitLines("a\nb\nc\n"),
		B:        SplitLines("a\nB\nc\n"),
		FromFile: "original",
		ToFile:   "modified",
		Context:  1,
	}

	result, err := GetUnifiedDiffString(diff)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(result, "---") {
		t.Errorf("expected unified diff header, got: %s", result)
	}
	if !strings.Contains(result, "-b") {
		t.Errorf("expected removed line '-b', got: %s", result)
	}
	if !strings.Contains(result, "+B") {
		t.Errorf("expected added line '+B', got: %s", result)
	}
}

func TestGetUnifiedDiffStringIdentical(t *testing.T) {
	lines := SplitLines("a\nb\nc\n")
	diff := UnifiedDiff{
		A: lines,
		B: lines,
	}

	result, err := GetUnifiedDiffString(diff)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "" {
		t.Errorf("expected empty diff for identical inputs, got: %s", result)
	}
}

func TestWriteUnifiedDiff(t *testing.T) {
	diff := UnifiedDiff{
		A:        SplitLines("a\nb\nc\n"),
		B:        SplitLines("a\nB\nc\n"),
		FromFile: "original",
		ToFile:   "modified",
		Context:  1,
	}

	var buf bytes.Buffer
	err := WriteUnifiedDiff(&buf, diff)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	result := buf.String()
	if !strings.Contains(result, "-b") {
		t.Errorf("expected removed line '-b', got: %s", result)
	}
	if !strings.Contains(result, "+B") {
		t.Errorf("expected added line '+B', got: %s", result)
	}
}
