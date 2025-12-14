// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bufio"
	"fmt"
	"reflect"
	"time"

	"github.com/go-openapi/testify/v2/internal/difflib"
)

/*
	Helper functions
*/

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
		e = reflect.ValueOf(expected).String()
		a = reflect.ValueOf(actual).String()
	case reflect.TypeFor[time.Time]():
		e = spewConfigStringerEnabled.Sdump(expected)
		a = spewConfigStringerEnabled.Sdump(actual)
	default:
		e = spewConfig.Sdump(expected)
		a = spewConfig.Sdump(actual)
	}

	diff, _ := difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(e),
		B:        difflib.SplitLines(a),
		FromFile: "Expected",
		FromDate: "",
		ToFile:   "Actual",
		ToDate:   "",
		Context:  1,
	})

	return "\n\nDiff:\n" + diff
}

func typeAndKind(v any) (reflect.Type, reflect.Kind) {
	t := reflect.TypeOf(v)
	k := t.Kind()

	if k == reflect.Ptr {
		t = t.Elem()
		k = t.Kind()
	}
	return t, k
}

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(format string, data any) string {
	const maxMessageSize = bufio.MaxScanTokenSize/2 - 100

	value := fmt.Sprintf(format, data)
	// Give us space for two truncated objects and the surrounding sentence.
	if len(value) > maxMessageSize {
		value = value[0:maxMessageSize] + "<... truncated>"
	}
	return value
}
