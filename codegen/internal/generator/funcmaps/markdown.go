// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"fmt"
	"regexp"
	"strings"
)

var refLinkPattern = regexp.MustCompile(`(?m)^\[([^\]]+)\]:\s+(.+)$`)

// FormatMarkdown carries out a formatting of inline markdown in comments that handles:
//
//  1. Reference-style markdown links: [text]: url
//  2. Godoc-style links: [errors.Is], [testing.T], etc.
//
//nolint:gocognit,gocognit // this is temporary accepted extra complexity. Should refactor with steps as functions
func FormatMarkdown(in string) string {
	const sensiblePrealloc = 20

	// Step 1: Extract reference-style link definitions
	// Pattern: [text]: url (at start of line or after whitespace)
	refLinks := make(map[string]string)

	// Extract all reference links
	matches := refLinkPattern.FindAllStringSubmatch(in, -1)
	const expectedGroups = 2
	for _, match := range matches {
		if len(match) == expectedGroups+1 {
			refLinks[match[1]] = match[2]
		}
	}

	// Remove reference link definitions from input
	processed := refLinkPattern.ReplaceAllString(in, "")

	// Step 2: Convert reference-style links to inline links
	// Replace [text] with [text](url) where we have the reference
	usedRefs := make(map[string]bool)
	for refText, refURL := range refLinks {
		// Pattern: [text] not followed by ( or :
		pattern := regexp.MustCompile(`\[` + regexp.QuoteMeta(refText) + `\]([^(\[]|$)`)
		if pattern.MatchString(processed) {
			processed = pattern.ReplaceAllString(processed, fmt.Sprintf("[%s](%s)$1", refText, refURL))
			usedRefs[refText] = true
		}
	}

	// Collect unused reference links to append later (dangling references)
	var danglingRefs []string
	for refText, refURL := range refLinks {
		if !usedRefs[refText] {
			danglingRefs = append(danglingRefs, fmt.Sprintf("[%s]: %s", refText, refURL))
		}
	}

	// Step 3: Convert godoc-style links to pkg.go.dev URLs
	// Pattern: [package.Type] - must contain a dot and only valid identifier chars
	// We process line by line to avoid matching already-converted links
	lines := strings.Split(processed, "\n")
	for i, line := range lines {
		// Skip if line already has inline links (contains ](http)
		if strings.Contains(line, "](") {
			continue
		}

		// Find godoc-style links: [word.word]
		godocPattern := regexp.MustCompile(`\[([a-zA-Z0-9_]+\.[a-zA-Z0-9_.]+)\]`)
		lines[i] = godocPattern.ReplaceAllStringFunc(line, func(match string) string {
			identifier := strings.Trim(match, "[]")

			// Has package qualifier - link to pkg.go.dev
			const expectedParts = 2
			parts := strings.SplitN(identifier, ".", expectedParts)
			if len(parts) != expectedParts {
				// Defensive code: This should never happen with the current regex pattern,
				// which requires at least one dot. If this triggers, it indicates a bug in
				// the regex pattern that needs to be fixed.
				panic(fmt.Errorf("internal error: godoc pattern matched %q but split into %d parts instead of %d",
					identifier, len(parts), expectedParts))
			}

			pkgPath := parts[0]
			symbol := parts[1]

			// Assume standard library for simple package names
			return fmt.Sprintf("[%s](https://pkg.go.dev/%s#%s)", identifier, pkgPath, symbol)
		})
	}
	processed = strings.Join(lines, "\n")

	// Step 4: Do existing processing (Hugo shortcodes)
	result := make([]string, 0, sensiblePrealloc)
	expanded := false
	tab := false

	for line := range strings.SplitSeq(processed, "\n") {
		if expanded && len(strings.TrimSpace(line)) == 0 {
			continue
		}

		if strings.HasPrefix(line, "#") {
			if !expanded {
				result = append(result, `{{% expand title="Examples" %}}`)
				result = append(result, `{{< tabs >}}`)
				expanded = true
			}

			title := Titleize(strings.TrimLeft(line, "# \t"))
			if tab {
				result = append(result, "```")
				result = append(result, `{{< /tab >}}`)
			}

			result = append(result, fmt.Sprintf(`{{%% tab title="%s" %%}}`, title))
			result = append(result, "```go")
			tab = true

			continue
		}

		result = append(result, line)
	}

	if tab {
		result = append(result, "```")
		result = append(result, `{{< /tab >}}`)
	}

	if expanded {
		result = append(result, `{{< /tabs >}}`)
		result = append(result, `{{% /expand %}}`)

		// Append dangling reference links as additional context
		if len(danglingRefs) > 0 {
			result = append(result, "")
			result = append(result, danglingRefs...)
		}
	}

	return strings.Join(result, "\n")
}
