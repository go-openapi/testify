// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestDiff(t *testing.T) {
	type myTime time.Time
	t0 := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	t1 := t0.Add(time.Second)

	t.Run("diff should render time with stringer", func(t *testing.T) {
		diffResult := diff(t0, t1)
		if strings.Contains(diffResult, "-(time.Time) 2026-01-01 00:00:00 +0000 UTC") &&
			strings.Contains(diffResult, "+(time.Time) 2026-01-01 00:00:01 +0000 UTC") {
			return
		}

		t.Errorf("unexpected diff time output, got: %q", diffResult)
	})

	t.Run("diff should render nested times with stringer", func(t *testing.T) {
		type myStruct struct {
			A time.Time
			B myTime
			C *time.Time
		}
		expected := myStruct{
			A: t0,
			B: myTime(t0),
			C: &t0,
		}
		actual := myStruct{
			A: t1,
			B: myTime(t1),
			C: &t1,
		}

		diffResult := diff(expected, actual)
		if strings.Contains(diffResult, "- A: (time.Time) 2026-01-01 00:00:00 +0000 UTC") &&
			strings.Contains(diffResult, "- B: (assertions.myTime) 2026-01-01 00:00:00 +0000 UTC") &&
			strings.Contains(diffResult, "- C: (*time.Time)(2026-01-01 00:00:00 +0000 UTC)") &&
			strings.Contains(diffResult, "+ A: (time.Time) 2026-01-01 00:00:01 +0000 UTC") &&
			strings.Contains(diffResult, "+ B: (assertions.myTime) 2026-01-01 00:00:01 +0000 UTC") &&
			strings.Contains(diffResult, "+ C: (*time.Time)(2026-01-01 00:00:01 +0000 UTC)") {
			return
		}

		t.Errorf("unexpected diff time output, got: %q", diffResult)
	})

	t.Run("diff on nil/nil interface types should render empty", func(t *testing.T) {
		var a, b error

		diffResult := diff(a, &b)
		if diffResult != "" {
			t.Errorf("expected an empty string to render the diff")
		}

		diffResult = diff((*error)(nil), (*error)(nil))
		if diffResult != "" {
			t.Errorf("expected an empty string to render the diff")
		}
	})
}

func TestDiffTypeAndKind(t *testing.T) {
	t.Run("should return nil and Invalid for nil interface", func(t *testing.T) {
		var v any // nil interface value

		typ, kind := typeAndKind(v)

		if typ != nil {
			t.Errorf("expected nil type, got %v", typ)
		}
		if kind != reflect.Invalid {
			t.Errorf("expected Invalid kind, got %v", kind)
		}
	})
}
