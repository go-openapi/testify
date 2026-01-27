// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	refLinkPattern = regexp.MustCompile(`(?m)^\[([^\]]+)\]:\s+(.+)$`)
	// Find godoc-style links: [word.word].
	godocPattern      = regexp.MustCompile(`\[([a-zA-Z0-9_]+\.[a-zA-Z0-9_.]+)\]`)
	godocPatternLocal = regexp.MustCompile(`\[([a-zA-Z0-9_.]+)\]`)
	tabSectionPattern = regexp.MustCompile(`^#\s+(Examples?|Usage|Use)$`)
	sectionPattern    = regexp.MustCompile(`^#\s+`)
)

// FormatMarkdown carries out a formatting of inline markdown in comments that handles:
//
//  1. Reference-style markdown links: [text]: url
//  2. Godoc-style links: [errors.Is], [testing.T], etc.
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

		replacedStdLib := godocPattern.ReplaceAllStringFunc(line, func(match string) string {
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

		// check for links with current package, e.g. [Boolean]. Those don't have a dot.
		replacedLocal := godocPatternLocal.ReplaceAllStringFunc(replacedStdLib, func(match string) string {
			identifier := strings.Trim(match, "[]")
			return fmt.Sprintf("[%s](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#%s)", identifier, identifier)
		})

		lines[i] = replacedLocal
	}

	processed = strings.Join(lines, "\n")

	// Step 4: render Usage and Examples sections with Hugo shortcodes
	// These are rendered as tabs.
	result := make([]string, 0, sensiblePrealloc)
	trailer := make([]string, 0, sensiblePrealloc)

	// parse state
	tabsCollection := false
	expanded := false
	tab := false

	for line := range strings.SplitSeq(processed, "\n") {
		if expanded && len(strings.TrimSpace(line)) == 0 {
			continue
		}

		matches := tabSectionPattern.FindStringSubmatch(line)
		const expectedCapture = 1
		if len(matches) != expectedCapture+1 {
			// either a regular line or a section header that we don't want in the trailer tabs
			if sectionPattern.MatchString(line) {
				// found a new section
				expanded = false
				line = strings.ReplaceAll(line, "#", "####") // containing function heading is ###
			}

			if expanded {
				trailer = append(trailer, line)
			} else {
				result = append(result, line)
			}

			continue
		}

		// interesting section to catch as a tab
		section := matches[expectedCapture]
		expanded = true

		if !tabsCollection {
			trailer = append(trailer, `{{% expand title="Examples" %}}`) // the title of the collapsible section
			trailer = append(trailer, `{{< tabs >}}`)
			tabsCollection = true
		}

		title := titleize(section)
		if tab {
			trailer = append(trailer, "```")
			trailer = append(trailer, `{{< /tab >}}`)
		}

		trailer = append(trailer, fmt.Sprintf(`{{%% tab title="%s" %%}}`, title))
		trailer = append(trailer, "```go")
		tab = true
	}

	if tab {
		trailer = append(trailer, "```")
		trailer = append(trailer, `{{< /tab >}}`)
	}

	if tabsCollection {
		trailer = append(trailer, `{{< /tabs >}}`)
		trailer = append(trailer, `{{% /expand %}}`)
	}

	// Append dangling reference links as additional context
	if len(danglingRefs) > 0 {
		result = append(result, "")
		result = append(result, danglingRefs...)
	}

	return strings.Join(append(result, trailer...), "\n")
}
