// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
	"reflect"
	"regexp"
)

// Regexp asserts that a specified regular expression matches a string.
//
// The regular expression may be passed as a [regexp.Regexp], a string or a []byte and will be compiled.
//
// The actual argument to be matched may be a string, []byte or anything that prints as a string with [fmt.Sprint].
//
// # Usage
//
//	assertions.Regexp(t, regexp.MustCompile("start"), "it's starting")
//	assertions.Regexp(t, "start...$", "it's not starting")
//
// # Examples
//
//	success: "^start", "starting"
//	failure: "^start", "not starting"
func Regexp(t T, rx any, actual any, msgAndArgs ...any) bool {
	// Domain: string
	if h, ok := t.(H); ok {
		h.Helper()
	}

	re, err := buildRegex(rx)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	switch v := actual.(type) {
	case []byte:
		return matchRegex(t, re, v, true, msgAndArgs...)
	case string:
		return matchRegex(t, re, v, true, msgAndArgs...)
	default:
		// reflection-based check for uncommon usage
		str, ok := asString(actual)
		if !ok {
			return matchRegex(t, re, fmt.Sprint(actual), true, msgAndArgs...)
		}
		return matchRegex(t, re, str, true, msgAndArgs...)
	}
}

// RegexpT asserts that a specified regular expression matches a string.
//
// The actual argument to be matched may be a string or []byte.
//
// See [Regexp].
//
// # Examples
//
//	success: "^start", "starting"
//	failure: "^start", "not starting"
func RegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool {
	// Domain: string
	if h, ok := t.(H); ok {
		h.Helper()
	}

	re, err := buildRegex(rx)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	return matchRegex(t, re, actual, true, msgAndArgs...)
}

// NotRegexp asserts that a specified regular expression does not match a string.
//
// See [Regexp].
//
// # Usage
//
//	assertions.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//	assertions.NotRegexp(t, "^start", "it's not starting")
//
// # Examples
//
//	success: "^start", "not starting"
//	failure: "^start", "starting"
func NotRegexp(t T, rx any, actual any, msgAndArgs ...any) bool {
	// Domain: string
	if h, ok := t.(H); ok {
		h.Helper()
	}

	re, err := buildRegex(rx)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	switch v := actual.(type) {
	case []byte:
		return matchRegex(t, re, v, false, msgAndArgs...)
	case string:
		return matchRegex(t, re, v, false, msgAndArgs...)
	default:
		// reflection-based check for uncommon usage
		str, ok := asString(actual)
		if !ok {
			return matchRegex(t, re, fmt.Sprint(actual), false, msgAndArgs...)
		}

		// handle ~string, ~[]byte
		return matchRegex(t, re, str, false, msgAndArgs...)
	}
}

// NotRegexpT asserts that a specified regular expression does not match a string.
//
// See [RegexpT].
//
// # Usage
//
//	assertions.NotRegexp(t, regexp.MustCompile("starts"), "it's starting")
//	assertions.NotRegexp(t, "^start", "it's not starting")
//
// # Examples
//
//	success: "^start", "not starting"
//	failure: "^start", "starting"
func NotRegexpT[Rex RegExp, ADoc Text](t T, rx Rex, actual ADoc, msgAndArgs ...any) bool {
	// Domain: string
	if h, ok := t.(H); ok {
		h.Helper()
	}

	re, err := buildRegex(rx)
	if err != nil {
		return Fail(t, err.Error(), msgAndArgs...)
	}

	return matchRegex(t, re, actual, false, msgAndArgs...)
}

func buildRegex(re any) (*regexp.Regexp, error) {
	// Maintainer: proposal for enhancement(perf): cache regexp
	switch v := re.(type) {
	case *regexp.Regexp:
		if v == nil {
			return nil, fmt.Errorf("regexp must not be nil: %w", errAssertions)
		}

		return v, nil
	case string:
		return compileRegex(v)
	case []byte:
		if v == nil {
			return nil, fmt.Errorf("regexp must not be nil: %w", errAssertions)
		}

		return compileRegex(string(v))
	default:
		// reflection-based check for uncommon usage
		str, ok := asString(re)
		if !ok {
			return nil, fmt.Errorf(
				"type for regexp is not supported. Want string, []byte or anything that converts to those, but got %T",
				re,
			)
		}
		// handle ~string, ~[]byte

		return compileRegex(str)
	}
}

func compileRegex(rx string) (*regexp.Regexp, error) {
	const errMsg = "invalid error expression %q: %w"
	rex, err := regexp.Compile(rx)
	if err != nil {
		return nil, fmt.Errorf(errMsg, rx, err)
	}

	return rex, nil
}

func matchRegex[ADoc Text](t T, rx *regexp.Regexp, actual ADoc, wantMatch bool, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	var matched bool
	str := any(actual)

	switch v := str.(type) {
	case []byte:
		matched = rx.Match(v)
	case string:
		matched = rx.MatchString(v)
	default:
		// safeguard: should never get there
		matched = rx.MatchString(string(actual))
	}

	switch {
	case wantMatch && !matched:
		return Fail(t, fmt.Sprintf(`Expect %q to match %q"`, string(actual), rx), msgAndArgs...)
	case !wantMatch && matched:
		return Fail(t, fmt.Sprintf("Expect %q to NOT match %q", string(actual), rx), msgAndArgs...)
	default:
		return true
	}
}

func asString(v any) (string, bool) {
	typeString := reflect.TypeFor[string]()
	val := reflect.ValueOf(v)
	kind := val.Kind()

	if !val.CanConvert(typeString) {
		return "", false
	}

	// weird reflection: numbers CanConvert but their string rep is wrong. Need to check further.
	typ := val.Type()
	if kind != reflect.String {
		if kind != reflect.Slice {
			return "", false
		}

		if typ.Elem().Kind() != reflect.Uint8 {
			return "", false
		}
	}

	// handle ~string, ~[]byte
	return val.Convert(typeString).String(), true
}

// assertionsError reports an error from a bad usage of an assertion.
type assertionsError string

func (e assertionsError) Error() string {
	return string(e)
}

const errAssertions assertionsError = "error from assertions"
