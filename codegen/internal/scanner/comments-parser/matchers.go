// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"fmt"
	"regexp"
	"strings"
)

// StartSectionFunc creates a matcher function for section headers.
//
// A section header can be:
//
//   - "# Placeholder" or "# Placeholders" (markdown-style)
//   - "Placeholder:" or "Placeholders:" (simple colon-style)
//
// Matching is case-insensitive.
func StartSectionFunc(placeholder string) func(string) bool {
	sectionRex := regexp.MustCompile(
		fmt.Sprintf(`(?i)^\s*(#\s+%[1]ss?\s*$|%[1]ss?\s*:)`, regexp.QuoteMeta(placeholder)),
	)

	return sectionRex.MatchString
}

// StartValueFunc creates a matcher function for key-value lines.
//
// A value line has the format: "placeholder: value" or "placeholders: value" (plural form).
// Returns the value part (everything after the colon) and true if matched.
// Matching is case-insensitive.
func StartValueFunc(placeholder string) func(string) (string, bool) {
	valueRex := regexp.MustCompile(
		fmt.Sprintf(`(?i)^\s*%[1]ss?\s*:\s*`, regexp.QuoteMeta(placeholder)),
	)

	return func(text string) (string, bool) {
		if valueRex.MatchString(text) {
			return valueRex.ReplaceAllLiteralString(text, ""), true
		}

		return "", false
	}
}

// StartAnotherSection checks if a line starts a new section.
//
// A line starts a new section if it either:
//
//   - Starts with a capital letter and ends with a colon (e.g., "Usage:")
//   - Starts with "# " (markdown header)
func StartAnotherSection(text string) bool {
	return len(text) > 0 && (text[0] >= 'A' && text[0] <= 'Z' && strings.HasSuffix(text, ":")) || strings.HasPrefix(text, "# ")
}
