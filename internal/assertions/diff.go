// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"reflect"

	"github.com/go-openapi/testify/v2/internal/assertions/enable/colors"
	"github.com/go-openapi/testify/v2/internal/difflib"
)

// diff returns a diff of both values as long as both are of the same type and
// are a struct, map, slice, array or string. Otherwise it returns an empty string.
func diff(expected any, actual any) string {
	if expected == nil || actual == nil {
		return ""
	}

	et, ek := typeAndKind(expected)
	at, _ := typeAndKind(actual)

	if et != at {
		return ""
	}

	if ek != reflect.Struct && ek != reflect.Map && ek != reflect.Slice && ek != reflect.Array && ek != reflect.String {
		return ""
	}

	var e, a string

	switch et {
	case reflect.TypeFor[string]():
		// short-circuit for plain strings
		e = reflect.ValueOf(expected).String()
		a = reflect.ValueOf(actual).String()
	default:
		e = dumper(expected)
		a = dumper(actual)
	}

	unified := difflib.UnifiedDiff{
		A:        difflib.SplitLines(e),
		B:        difflib.SplitLines(a),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	}

	if colors.Enabled() {
		unified.Options = colors.Options()
	}

	diff, _ := difflib.GetUnifiedDiffString(unified)

	return "\n\nDiff:\n" + diff
}

func typeAndKind(v any) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	if t == nil {
		return nil, reflect.Invalid
	}

	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}

	return t, k
}
