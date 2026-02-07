// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

var (
	refLinkPattern    = regexp.MustCompile(`(?m)^\[([^\]]+)\]:\s+(.+)$`)          // find godoc-style links: [word.word].
	godocPattern      = regexp.MustCompile(`\[([a-zA-Z0-9_]+\.[a-zA-Z0-9_.]+)\]`) // there is a dot
	godocPatternLocal = regexp.MustCompile(`\[([a-zA-Z0-9_]+)\]`)                 // no dot
	tabSectionPattern = regexp.MustCompile(`^#\s+(Examples?|Usage|Use)$`)
	sectionPattern    = regexp.MustCompile(`^#\s+`)
)

const (
	sensiblePrealloc = 20
	godocHost        = "pkg.go.dev"
	repo             = "github.com/go-openapi/testify/v2"
)

// FormatMarkdown carries out a formatting of inline markdown in comments that handles:
//
//  1. Reference-style markdown links: [text]: url
//  2. Godoc-style links: [errors.Is], [testing.T], etc.
//  3. Rearrange godoc sections so we extract Usage and Examples. Examples are matched with the
//     generated testable examples and reinjected as go code in the "Testable Examples" tab.
func FormatMarkdown(in string, object any) string {
	// 1. process godoc links
	processed, danglingRefs := markdownLinks(in)

	// 2. build pkg.go.dev links
	processed = godocLinks(processed)

	// 3. strip Usage and Example sections and render them as hugo tabs
	result, trailer := stripSections(processed, object)

	// Append dangling reference links as additional context
	if len(danglingRefs) > 0 {
		result = append(result, "")
		result = append(result, danglingRefs...)
	}

	return strings.Join(append(result, trailer...), "\n")
}

// markdownLinks processes links found in godoc strings to produce markdown links.
//
// If some references are defined, they are stashed and returned as a slice of "dangling references".
func markdownLinks(in string) (string, []string) {
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

	// Convert reference-style links to inline links
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

	return processed, danglingRefs
}

// godocLinks converts godoc-style links to pkg.go.dev URLs.
//
// Pattern: [package.Type] - must contain a dot and only valid identifier chars.
// We process line by line to avoid matching already-converted links.
func godocLinks(in string) string {
	lines := strings.Split(in, "\n")
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
			return fmt.Sprintf("[%s](https://%s/%s#%s)", identifier, godocHost, pkgPath, symbol)
		})

		// check for links with current package, e.g. [Boolean]. Those don't have a dot.
		replacedLocal := godocPatternLocal.ReplaceAllStringFunc(replacedStdLib, func(match string) string {
			identifier := strings.Trim(match, "[]")
			return fmt.Sprintf("[%s](https://%s/%s/assert#%s)", identifier, godocHost, repo, identifier)
		})

		lines[i] = replacedLocal
	}

	return strings.Join(lines, "\n")
}

// stripSections renders Usage and Examples sections with Hugo shortcodes.
// These are rendered as tabs.
//
// It returns a result of lines (indented markdown with sections), a trailer of lines (Usage and Examples tabs).
func stripSections(in string, object any) (result, trailer []string) {
	result = make([]string, 0, sensiblePrealloc)
	trailer = make([]string, 0, sensiblePrealloc)

	var (
		testableExamples []model.Renderable
		funcName         string
	)
	if function, ok := (object).(model.Function); ok {
		testableExamples = function.Examples
		funcName = function.Name
	}

	// parse state
	tabsCollection := false
	expanded := false
	tab := false
	hasTestableExamples := false

	for line := range strings.SplitSeq(in, "\n") {
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

		if strings.EqualFold(section, "Examples") && len(testableExamples) > 0 {
			hasTestableExamples = true
			continue // skip : we'll add testable examples below
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

	if hasTestableExamples {
		trailer = append(trailer, `{{% tab title="Testable Examples" %}}`)
		trailer = append(trailer, `{{% cards %}}`)
		tabsCollection = true

		for _, example := range testableExamples {
			trailer = append(trailer, `{{% card href="https://go.dev/play/" %}}`)
			trailer = append(trailer, "\n")
			trailer = append(trailer, `*Copy and click to open Go Playground*`)
			trailer = append(trailer, "\n")
			trailer = append(trailer, "```go")
			trailer = append(trailer, fmt.Sprintf("// real-world test would inject *testing.T from Test%s(t *testing.T)", funcName))
			trailer = append(trailer, example.Render())
			trailer = append(trailer, "```")
			trailer = append(trailer, `{{% /card %}}`)
			trailer = append(trailer, "\n")
		}
		trailer = append(trailer, `{{% /cards %}}`)
		trailer = append(trailer, `{{< /tab >}}`)
		trailer = append(trailer, "\n")
	}

	if tabsCollection {
		trailer = append(trailer, `{{< /tabs >}}`)
		trailer = append(trailer, `{{% /expand %}}`)
	}

	return result, trailer
}
