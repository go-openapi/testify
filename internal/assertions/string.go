// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"regexp"
)

// Regexp asserts that a specified regexp matches a string.
//
//	assert.Regexp(t, regexp.MustCompile("start"), "it's starting")
//	assert.Regexp(t, "start...$", "it's not starting")
//
// Examples:
//
//	success: "^start", "starting"
//	failure: "^start", "not starting"
func Regexp(t T, rx any, str any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	match := matchRegexp(rx, str)

	if !match {
		Fail(t, fmt.Sprintf(`Expect "%v" to match "%v"`, str, rx), msgAndArgs...)
	}

	return match
}

// NotRegexp asserts that a specified regexp does not match a string.
//
//	assert.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//	assert.NotRegexp(t, "^start", "it's not starting")
//
// Examples:
//
//	success: "^start", "not starting"
//	failure: "^start", "starting"
func NotRegexp(t T, rx any, str any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	match := matchRegexp(rx, str)

	if match {
		Fail(t, fmt.Sprintf("Expect \"%v\" to NOT match \"%v\"", str, rx), msgAndArgs...)
	}

	return !match
}

// matchRegexp return true if a specified regexp matches a string.
func matchRegexp(rx any, str any) bool {
	var r *regexp.Regexp
	if rr, ok := rx.(*regexp.Regexp); ok {
		r = rr
	} else {
		r = regexp.MustCompile(fmt.Sprint(rx))
	}

	switch v := str.(type) {
	case []byte:
		return r.Match(v)
	case string:
		return r.MatchString(v)
	default:
		return r.MatchString(fmt.Sprint(v))
	}
}
