// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"fmt"
	"iter"
	"slices"
	"testing"
)

func TestStartSectionFunc(t *testing.T) {
	t.Parallel()

	for c := range startSectionCases() {
		t.Run(fmt.Sprintf("%q matches %q", c.placeholder, c.input), func(t *testing.T) {
			t.Parallel()

			matcher := StartSectionFunc(c.placeholder)
			result := matcher(c.input)

			if result != c.expected {
				t.Errorf("StartSectionFunc(%q)(%q) = %v, expected %v",
					c.placeholder, c.input, result, c.expected)
			}
		})
	}
}

func TestStartValueFunc(t *testing.T) {
	t.Parallel()

	for c := range startValueCases() {
		t.Run(fmt.Sprintf("%q extracts from %q", c.placeholder, c.input), func(t *testing.T) {
			t.Parallel()

			matcher := StartValueFunc(c.placeholder)
			value, ok := matcher(c.input)

			if ok != c.expectedOk {
				t.Errorf("StartValueFunc(%q)(%q) ok = %v, expected %v",
					c.placeholder, c.input, ok, c.expectedOk)
			}

			if value != c.expectedValue {
				t.Errorf("StartValueFunc(%q)(%q) value = %q, expected %q",
					c.placeholder, c.input, value, c.expectedValue)
			}
		})
	}
}

func TestStartAnotherSection(t *testing.T) {
	t.Parallel()

	for c := range startAnotherSectionCases() {
		t.Run(c.input, func(t *testing.T) {
			t.Parallel()

			result := StartAnotherSection(c.input)

			if result != c.expected {
				t.Errorf("StartAnotherSection(%q) = %v, expected %v",
					c.input, result, c.expected)
			}
		})
	}
}

/* Test case iterators */

type startSectionCase struct {
	placeholder string
	input       string
	expected    bool
}

func startSectionCases() iter.Seq[startSectionCase] {
	return slices.Values([]startSectionCase{
		// Markdown-style headers
		{"Example", "# Example", true},
		{"Example", "# Examples", true},
		{"Example", "# EXAMPLES", true},
		{"Example", "  # Examples", true},
		{"Example", "#   Examples", true},

		// Colon-style headers
		{"Example", "Example:", true},
		{"Example", "Examples:", true},
		{"Example", "EXAMPLES:", true},
		{"Example", "  Examples:", true},
		{"Example", "Examples  :", true},

		// Domain examples
		{"Domain", "# Domains", true},
		{"Domain", "Domains:", true},

		// Negative cases
		{"Example", "# Usage", false},
		{"Example", "Examples", false},
		{"Example", "# Examples are here", false},
		{"Example", "", false},
		{"Example", "Some Examples:", false},

		// Special characters in placeholder
		{"Pre-commit", "# Pre-commits", true},
		{"Test(s)", "# Test(s)", true},
	})
}

type startValueCase struct {
	placeholder   string
	input         string
	expectedValue string
	expectedOk    bool
}

func startValueCases() iter.Seq[startValueCase] {
	return slices.Values([]startValueCase{
		// Basic matching
		{"success", "success: 123, 456", "123, 456", true},
		{"success", "SUCCESS: value", "value", true},
		{"failure", "  failure: value", "value", true},
		{"panic", "panic  :  value", "value", true},
		{"domain", "domain:", "", true},
		{"note", "note: Error: something failed", "Error: something failed", true},

		// Domain tag examples
		{"domain", "domain: string", "string", true},
		{"maintainer", "maintainer: @username", "@username", true},

		// Negative cases
		{"success", "failure: value", "", false},
		{"success", "success value", "", false},
		{"success", "", "", false},
		{"success", "not success: value", "", false},

		// Special characters
		{"pre-commit", "pre-commit: hook", "hook", true},
	})
}

type startAnotherSectionCase struct {
	input    string
	expected bool
}

func startAnotherSectionCases() iter.Seq[startAnotherSectionCase] {
	return slices.Values([]startAnotherSectionCase{
		// Capital letter + colon (section headers)
		{"Examples:", true},
		{"Usage:", true},
		{"Notes:", true},
		{"Parameters:", true},

		// Markdown headers
		{"# Section", true},
		{"#  Header", true},
		{"# examples", true}, // markdown header still matches

		// Negative cases
		{"examples:", false},
		{"Examples", false},
		{"Examples are here", false},
		{"", false},
		{"This is a description line.", false},
		{"success: 123", false},
		{"  Examples:", false}, // indented section header
	})
}
