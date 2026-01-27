// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

import (
	"regexp"
	"strings"

	"github.com/go-openapi/testify/codegen/v2/internal/model"
)

// ParseTaggedComments extracts tagged comments from text.
//
// Recognized tags:
//   - domain: <value>           - single-line domain tag
//   - maintainer: <value>       - multi-line maintainer note
//   - note: <value>             - multi-line note
//   - mention: <value>          - single-line mention
//
// Multi-line tags continue until the next tagged line or end of text.
func ParseTaggedComments(text string) []model.ExtraComment {
	const (
		domainPrefix     = "domain"
		maintainerPrefix = "maintainer"
		notePrefix       = "note"
		mentionPrefix    = "mention"
	)

	inValue := false
	startValueDomain := StartValueFunc(domainPrefix)
	startValueMaintainer := StartValueFunc(maintainerPrefix)
	startValueNote := StartValueFunc(notePrefix)
	startValueMention := StartValueFunc(mentionPrefix)

	startTaggedValue := func(line string) (key string, val string, tag model.CommentTag, multiline bool, ok bool) {
		val, ok = startValueDomain(line)
		if ok {
			return val, "", model.CommentTagDomain, false, true
		}
		val, ok = startValueMaintainer(line)
		if ok {
			return "", val, model.CommentTagMaintainer, true, true
		}
		val, ok = startValueNote(line)
		if ok {
			return "", val, model.CommentTagNote, true, true
		}
		val, ok = startValueMention(line)
		if ok {
			return "", val, model.CommentTagMention, false, true
		}

		return "", "", model.CommentTagNone, false, false
	}

	result := make([]model.ExtraComment, 0)

	for line := range strings.SplitSeq(text, "\n") {
		key, val, tag, multiline, ok := startTaggedValue(line)
		if inValue && len(result) > 0 && len(line) > 0 && !ok {
			result[len(result)-1].Text += "\n" + line

			continue
		}

		if ok {
			if multiline {
				inValue = true
			}

			result = append(result, model.ExtraComment{
				Tag:  tag,
				Key:  key,
				Text: val,
			})

			continue
		}

		inValue = false
	}

	return result
}

// ParseDomainDescriptions extracts domain descriptions from package-level comment text.
//
// It looks for a "Domains:" or "Domain:" section with entries like:
//   - string: assertions on strings
//   - json: assertions on JSON documents
//
// The domain name is normalized to lowercase.
func ParseDomainDescriptions(text string) []model.ExtraComment {
	const allocatedDomains = 20

	result := make([]model.ExtraComment, 0, allocatedDomains)

	const (
		domainPrefix      = "domain" // e.g. domain:, domains:, # domain, # Domains
		rexCapturedGroups = 2
	)

	domainRex := regexp.MustCompile(`^\s*(?:(?:[\-\*])|(?:\d+\.))\s+(\w+)\s*:\s+(.+)`)
	startDomainSection := StartSectionFunc(domainPrefix)
	inDomainSection := false

	for line := range strings.SplitSeq(text, "\n") {
		if startDomainSection(line) {
			inDomainSection = true
			continue
		}
		if !inDomainSection {
			continue
		}

		if StartAnotherSection(line) {
			break
		}

		// parse domain descriptions
		matches := domainRex.FindStringSubmatch(line)
		if len(matches) < rexCapturedGroups+1 {
			continue
		}
		domain := strings.ToLower(matches[1])
		description := strings.TrimSpace(matches[2])

		result = append(result, model.ExtraComment{
			Tag:  model.CommentTagDomainDescription,
			Key:  domain,
			Text: description,
		})
	}

	return result
}

// DomainFromExtraComments retrieves the domain from preparsed extra (tagged) comment.
func DomainFromExtraComments(taggedComments []model.ExtraComment) string {
	for _, taggedComment := range taggedComments {
		if taggedComment.Tag == model.CommentTagDomain {
			return taggedComment.Key
		}
	}

	return ""
}
