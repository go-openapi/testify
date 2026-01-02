// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"iter"
	"slices"
	"testing"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

func TestParseTaggedComments(t *testing.T) {
	t.Parallel()

	for c := range parseTaggedCommentsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := ParseTaggedComments(c.input)

			if len(result) != len(c.expected) {
				t.Fatalf("ParseTaggedComments() returned %d comments, expected %d\nGot: %+v\nExpected: %+v",
					len(result), len(c.expected), result, c.expected)
			}

			for i, expected := range c.expected {
				got := result[i]
				if got.Tag != expected.Tag {
					t.Errorf("Comment[%d].Tag = %v, expected %v", i, got.Tag, expected.Tag)
				}
				if got.Key != expected.Key {
					t.Errorf("Comment[%d].Key = %q, expected %q", i, got.Key, expected.Key)
				}
				if got.Text != expected.Text {
					t.Errorf("Comment[%d].Text = %q, expected %q", i, got.Text, expected.Text)
				}
			}
		})
	}
}

func TestParseDomainDescriptions(t *testing.T) {
	t.Parallel()

	for c := range parseDomainDescriptionsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := ParseDomainDescriptions(c.input)

			if len(result) != len(c.expected) {
				t.Fatalf("ParseDomainDescriptions() returned %d domains, expected %d\nGot: %+v\nExpected: %+v",
					len(result), len(c.expected), result, c.expected)
			}

			for i, expected := range c.expected {
				got := result[i]
				if got.Tag != expected.Tag {
					t.Errorf("Domain[%d].Tag = %v, expected %v", i, got.Tag, expected.Tag)
				}
				if got.Key != expected.Key {
					t.Errorf("Domain[%d].Key = %q, expected %q", i, got.Key, expected.Key)
				}
				if got.Text != expected.Text {
					t.Errorf("Domain[%d].Text = %q, expected %q", i, got.Text, expected.Text)
				}
			}
		})
	}
}

func TestDomainFromExtraComments(t *testing.T) {
	t.Parallel()

	for c := range domainFromExtraCommentsCases() {
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()

			result := DomainFromExtraComments(c.comments)

			if result != c.expected {
				t.Errorf("DomainFromExtraComments() = %q, expected %q", result, c.expected)
			}
		})
	}
}

/* Test case iterators */

type parseTaggedCommentsCase struct {
	name     string
	input    string
	expected []model.ExtraComment
}

func parseTaggedCommentsCases() iter.Seq[parseTaggedCommentsCase] {
	return slices.Values([]parseTaggedCommentsCase{
		{
			name: "single domain tag",
			input: `Some description.
domain: string
More text.`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomain, Key: "string", Text: ""},
			},
		},
		{
			name: "multiple domain tags",
			input: `domain: string
domain: boolean`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomain, Key: "string", Text: ""},
				{Tag: model.CommentTagDomain, Key: "boolean", Text: ""},
			},
		},
		{
			name:  "maintainer tag single line",
			input: `maintainer: @username`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagMaintainer, Key: "", Text: "@username"},
			},
		},
		{
			name: "maintainer tag multiline",
			input: `maintainer: This function needs review.
It has some edge cases that should be documented.
The implementation could be improved.`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagMaintainer, Key: "", Text: "This function needs review.\nIt has some edge cases that should be documented.\nThe implementation could be improved."},
			},
		},
		{
			name: "note tag multiline",
			input: `note: This is important.
This continues the note.
And this too.`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagNote, Key: "", Text: "This is important.\nThis continues the note.\nAnd this too."},
			},
		},
		{
			name:  "mention tag",
			input: `mention: Related to issue #123`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagMention, Key: "", Text: "Related to issue #123"},
			},
		},
		{
			name: "mixed tags",
			input: `domain: error
maintainer: @alice
This needs attention.
note: Check performance impact.
mention: See also HTTPError`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomain, Key: "error", Text: ""},
				{Tag: model.CommentTagMaintainer, Key: "", Text: "@alice\nThis needs attention."},
				{Tag: model.CommentTagNote, Key: "", Text: "Check performance impact."},
				{Tag: model.CommentTagMention, Key: "", Text: "See also HTTPError"},
			},
		},
		{
			name: "multiline interrupted by new tag",
			input: `maintainer: First maintainer note.
This continues.
note: New section starts here.
And continues.`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagMaintainer, Key: "", Text: "First maintainer note.\nThis continues."},
				{Tag: model.CommentTagNote, Key: "", Text: "New section starts here.\nAnd continues."},
			},
		},
		{
			name: "case insensitive tags",
			input: `DOMAIN: string
Maintainer: @bob
NOTE: Important info`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomain, Key: "string", Text: ""},
				{Tag: model.CommentTagMaintainer, Key: "", Text: "@bob"},
				{Tag: model.CommentTagNote, Key: "", Text: "Important info"},
			},
		},
		{
			name:     "no tags",
			input:    `Just regular text without any tags.`,
			expected: []model.ExtraComment{},
		},
		{
			name:     "empty input",
			input:    "",
			expected: []model.ExtraComment{},
		},
		{
			name: "tags with extra whitespace",
			input: `  domain  :   string
  maintainer  :  @user  `,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomain, Key: "string", Text: ""},
				{Tag: model.CommentTagMaintainer, Key: "", Text: "@user  "},
			},
		},
		{
			name: "multiline ends with empty line",
			input: `note: First line.
Second line.

domain: string`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagNote, Key: "", Text: "First line.\nSecond line."},
				{Tag: model.CommentTagDomain, Key: "string", Text: ""},
			},
		},
	})
}

type parseDomainDescriptionsCase struct {
	name     string
	input    string
	expected []model.ExtraComment
}

func parseDomainDescriptionsCases() iter.Seq[parseDomainDescriptionsCase] {
	return slices.Values([]parseDomainDescriptionsCase{
		{
			name: "domains section with bullet points",
			input: `Package assertions provides testing utilities.

Domains:
  - string: assertions on strings
  - json: assertions on JSON documents
  - yaml: assertions on YAML documents`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "string", Text: "assertions on strings"},
				{Tag: model.CommentTagDomainDescription, Key: "json", Text: "assertions on JSON documents"},
				{Tag: model.CommentTagDomainDescription, Key: "yaml", Text: "assertions on YAML documents"},
			},
		},
		{
			name: "domains section with asterisks",
			input: `Domains:
  * boolean: true/false assertions
  * number: numeric comparisons`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "boolean", Text: "true/false assertions"},
				{Tag: model.CommentTagDomainDescription, Key: "number", Text: "numeric comparisons"},
			},
		},
		{
			name: "domains section with numbered list",
			input: `Domains:
  1. error: error handling assertions
  2. http: HTTP response assertions`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "error", Text: "error handling assertions"},
				{Tag: model.CommentTagDomainDescription, Key: "http", Text: "HTTP response assertions"},
			},
		},
		{
			name: "markdown header style",
			input: `# Domains

  - file: file system assertions
  - time: time-based assertions`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "file", Text: "file system assertions"},
				{Tag: model.CommentTagDomainDescription, Key: "time", Text: "time-based assertions"},
			},
		},
		{
			name: "case insensitive domain header",
			input: `DOMAINS:
  - panic: panic assertions`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "panic", Text: "panic assertions"},
			},
		},
		{
			name: "singular form",
			input: `Domain:
  - collection: collection assertions`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "collection", Text: "collection assertions"},
			},
		},
		{
			name: "stops at next section",
			input: `Domains:
  - string: string assertions
  - json: JSON assertions

Usage:
  - boolean: this should not be parsed`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "string", Text: "string assertions"},
				{Tag: model.CommentTagDomainDescription, Key: "json", Text: "JSON assertions"},
			},
		},
		{
			name: "normalizes domain names to lowercase",
			input: `Domains:
  - String: string assertions
  - JSON: JSON assertions`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "string", Text: "string assertions"},
				{Tag: model.CommentTagDomainDescription, Key: "json", Text: "JSON assertions"},
			},
		},
		{
			name: "handles extra whitespace",
			input: `Domains:
  -   string   :   assertions on strings  `,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "string", Text: "assertions on strings"},
			},
		},
		{
			name:     "no domains section",
			input:    `This package has no domains section.`,
			expected: []model.ExtraComment{},
		},
		{
			name: "domains section with no entries",
			input: `Domains:

Usage follows.`,
			expected: []model.ExtraComment{},
		},
		{
			name: "skip non-matching lines",
			input: `Domains:
  Some intro text.
  - string: string assertions
  Not a domain line.
  - json: JSON assertions`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "string", Text: "string assertions"},
				{Tag: model.CommentTagDomainDescription, Key: "json", Text: "JSON assertions"},
			},
		},
		{
			name: "description with colons",
			input: `Domains:
  - http: HTTP assertions: status codes, headers, bodies`,
			expected: []model.ExtraComment{
				{Tag: model.CommentTagDomainDescription, Key: "http", Text: "HTTP assertions: status codes, headers, bodies"},
			},
		},
	})
}

type domainFromExtraCommentsCase struct {
	name     string
	comments []model.ExtraComment
	expected string
}

func domainFromExtraCommentsCases() iter.Seq[domainFromExtraCommentsCase] {
	return slices.Values([]domainFromExtraCommentsCase{
		{
			name: "find domain tag",
			comments: []model.ExtraComment{
				{Tag: model.CommentTagNote, Key: "", Text: "Some note"},
				{Tag: model.CommentTagDomain, Key: "string", Text: ""},
				{Tag: model.CommentTagMaintainer, Key: "", Text: "@user"},
			},
			expected: "string",
		},
		{
			name: "first domain wins",
			comments: []model.ExtraComment{
				{Tag: model.CommentTagDomain, Key: "boolean", Text: ""},
				{Tag: model.CommentTagDomain, Key: "string", Text: ""},
			},
			expected: "boolean",
		},
		{
			name: "no domain tag",
			comments: []model.ExtraComment{
				{Tag: model.CommentTagNote, Key: "", Text: "Note"},
				{Tag: model.CommentTagMaintainer, Key: "", Text: "@user"},
			},
			expected: "",
		},
		{
			name:     "empty comments",
			comments: []model.ExtraComment{},
			expected: "",
		},
		{
			name:     "nil comments",
			comments: nil,
			expected: "",
		},
		{
			name: "domain at end",
			comments: []model.ExtraComment{
				{Tag: model.CommentTagMaintainer, Key: "", Text: "@user"},
				{Tag: model.CommentTagNote, Key: "", Text: "Note"},
				{Tag: model.CommentTagDomain, Key: "error", Text: ""},
			},
			expected: "error",
		},
	})
}
