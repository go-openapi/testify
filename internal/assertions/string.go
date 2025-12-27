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

	match, err := matchRegexp(rx, str)
	if err != nil {
		Fail(t, fmt.Sprintf("invalid regular expression %q: %v", rx, err), msgAndArgs...)

		return false
	}

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

	match, err := matchRegexp(rx, str)
	if err != nil {
		Fail(t, fmt.Sprintf("invalid regular expression %q: %v", rx, err), msgAndArgs...)

		return false
	}

	if match {
		Fail(t, fmt.Sprintf("Expect \"%v\" to NOT match \"%v\"", str, rx), msgAndArgs...)
	}

	return !match
}

// matchRegexp returns whether the compiled regular expression matches the provided value.
//
// If rx is not a *[regexp.Regexp], rx is formatted with fmt.Sprint and compiled.
// When compilation fails, an error is returned instead of panicking.
func matchRegexp(rx any, str any) (bool, error) {
	var r *regexp.Regexp
	if rr, ok := rx.(*regexp.Regexp); ok {
		r = rr
	} else {
		var err error
		r, err = regexp.Compile(fmt.Sprint(rx))
		if err != nil {
			return false, err
		}
	}

	switch v := str.(type) {
	case []byte:
		return r.Match(v), nil
	case string:
		return r.MatchString(v), nil
	default:
		return r.MatchString(fmt.Sprint(v)), nil
	}
}
