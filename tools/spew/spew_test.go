// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

import (
	"bytes"
	"strings"
	"testing"
)

func TestDump(_ *testing.T) {
	// Smoke test: Dump should not panic.
	Dump("hello", 42, []int{1, 2, 3})
}

func TestFdump(t *testing.T) {
	var buf bytes.Buffer
	Fdump(&buf, "hello", 42)

	out := buf.String()
	if !strings.Contains(out, "hello") {
		t.Errorf("Fdump output should contain %q, got: %s", "hello", out)
	}
	if !strings.Contains(out, "42") {
		t.Errorf("Fdump output should contain %q, got: %s", "42", out)
	}
}

func TestSdump(t *testing.T) {
	out := Sdump("hello", 42)

	if !strings.Contains(out, "hello") {
		t.Errorf("Sdump output should contain %q, got: %s", "hello", out)
	}
	if !strings.Contains(out, "42") {
		t.Errorf("Sdump output should contain %q, got: %s", "42", out)
	}
}

func TestConfig(t *testing.T) {
	if Config.Indent != " " {
		t.Errorf("expected default Indent to be a single space, got %q", Config.Indent)
	}
}
