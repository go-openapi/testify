// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"iter"
	"reflect"
	"slices"
	"strings"
	"testing"
	"time"
)

func TestHelpersUnexportedImplementationDetails(t *testing.T) {
	t.Parallel()

	t.Run("diff", testDiff())
	t.Run("diffList", testDiffList())
}

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

// Ensure there are no data races with diff.
func TestTypeDiffRace(t *testing.T) {
	t.Parallel()

	expected := map[string]string{
		"a": "A",
		"b": "B",
		"c": "C",
	}

	actual := map[string]string{
		"d": "D",
		"e": "E",
		"f": "F",
	}

	// run diffs in parallel simulating tests with t.Parallel()
	numRoutines := 10
	rChans := make([]chan string, numRoutines)
	for idx := range rChans {
		rChans[idx] = make(chan string)
		go func(ch chan string) {
			defer close(ch)
			ch <- diff(expected, actual)
		}(rChans[idx])
	}

	for _, ch := range rChans {
		for msg := range ch {
			if msg == "" {
				t.Error("expected non-empty diff result")
			}
		}
	}
}

func testDiff() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for tt := range diffCases() {
			t.Run(tt.name, func(t *testing.T) {
				t.Parallel()

				for range min(1, tt.repeat) { // for tests on maps, need to verify the ordering is stable
					actual := diff(
						tt.valueA,
						tt.valueB,
					)
					if tt.expected != actual {
						t.Errorf("expected diff:\n%s\ngot:\n%s", tt.expected, actual)
					}
				}
			})
		}
	}
}

func testDiffList() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for test := range compareDiffListCases() {
			t.Run(test.name, func(t *testing.T) {
				t.Parallel()

				actualExtraA, actualExtraB := diffLists(test.listA, test.listB)
				if !reflect.DeepEqual(test.extraA, actualExtraA) {
					t.Errorf("extra A does not match for listA=%v listB=%v: expected %v, got %v",
						test.listA, test.listB, test.extraA, actualExtraA)
				}
				if !reflect.DeepEqual(test.extraB, actualExtraB) {
					t.Errorf("extra B does not match for listA=%v listB=%v: expected %v, got %v",
						test.listA, test.listB, test.extraB, actualExtraB)
				}
			})
		}
	}
}

type diffCase struct {
	name     string
	repeat   int
	valueA   any
	valueB   any
	expected string
}

type diffTestingStruct struct {
	A string
	B int
}

func (d *diffTestingStruct) String() string {
	return d.A
}

func diffCases() iter.Seq[diffCase] {
	const n = 5
	type Key struct {
		x int
	}

	return slices.Values([]diffCase{
		{
			name:   "with diff struct",
			valueA: struct{ foo string }{"hello"},
			valueB: struct{ foo string }{"bar"},
			expected: `

Diff:
--- Expected
+++ Actual
@@ -1,3 +1,3 @@
 (struct { foo string }) {
- foo: (string) (len=5) "hello"
+ foo: (string) (len=3) "bar"
 }
`,
		},
		{
			name:   "with diff slice",
			valueA: []int{1, 2, 3, 4},
			valueB: []int{1, 3, 5, 7},
			expected: `

Diff:
--- Expected
+++ Actual
@@ -2,5 +2,5 @@
  (int) 1,
- (int) 2,
  (int) 3,
- (int) 4
+ (int) 5,
+ (int) 7
 }
`,
		},
		{
			name:   "with diff slice (sliced)",
			valueA: []int{1, 2, 3, 4}[0:3],
			valueB: []int{1, 3, 5, 7}[0:3],
			expected: `

Diff:
--- Expected
+++ Actual
@@ -2,4 +2,4 @@
  (int) 1,
- (int) 2,
- (int) 3
+ (int) 3,
+ (int) 5
 }
`,
		},
		{
			name:   "with string keys map keys should be rendered deterministically in diffs",
			repeat: n,
			valueA: map[string]int{"one": 1, "two": 2, "three": 3, "four": 4},
			valueB: map[string]int{"one": 1, "three": 3, "five": 5, "seven": 7},
			expected: `

Diff:
--- Expected
+++ Actual
@@ -1,6 +1,6 @@
 (map[string]int) (len=4) {
- (string) (len=4) "four": (int) 4,
+ (string) (len=4) "five": (int) 5,
  (string) (len=3) "one": (int) 1,
- (string) (len=5) "three": (int) 3,
- (string) (len=3) "two": (int) 2
+ (string) (len=5) "seven": (int) 7,
+ (string) (len=5) "three": (int) 3
 }
`,
		},
		{
			name:   "with diff error",
			valueA: errors.New("some expected error"),
			valueB: errors.New("actual error"),
			expected: `

Diff:
--- Expected
+++ Actual
@@ -1,3 +1,3 @@
 (*errors.errorString)({
- s: (string) (len=19) "some expected error"
+ s: (string) (len=12) "actual error"
 })
`,
		},
		{
			name:   "with arbitrary comparable keys map keys should be rendered deterministically in diffs",
			repeat: n,
			valueA: map[Key]int{{1}: 1, {2}: 2, {3}: 3, {4}: 4},
			valueB: map[Key]int{{1}: 1, {2}: 2, {3}: 3, {4}: 999},
			expected: `

Diff:
--- Expected
+++ Actual
@@ -12,3 +12,3 @@
   x: (int) 4
- }: (int) 4
+ }: (int) 999
 }
`,
		},
		{
			name:   "with diff unexported struct",
			valueA: diffTestingStruct{A: "some string", B: 10},
			valueB: diffTestingStruct{A: "some string", B: 15},
			expected: `

Diff:
--- Expected
+++ Actual
@@ -2,3 +2,3 @@
  A: (string) (len=11) "some string",
- B: (int) 10
+ B: (int) 15
 }
`,
		},
		{
			name:   "with diff date",
			valueA: time.Date(2020, 9, 24, 0, 0, 0, 0, time.UTC),
			valueB: time.Date(2020, 9, 25, 0, 0, 0, 0, time.UTC),
			expected: `

Diff:
--- Expected
+++ Actual
@@ -1,2 +1,2 @@
-(time.Time) 2020-09-24 00:00:00 +0000 UTC
+(time.Time) 2020-09-25 00:00:00 +0000 UTC
 
`,
		},
	})
}

type compareDiffListCase struct {
	name   string
	listA  any
	listB  any
	extraA []any
	extraB []any
}

func compareDiffListCases() iter.Seq[compareDiffListCase] {
	return slices.Values([]compareDiffListCase{
		{
			name:   "equal empty",
			listA:  []string{},
			listB:  []string{},
			extraA: nil,
			extraB: nil,
		},
		{
			name:   "equal same order",
			listA:  []string{"hello", "world"},
			listB:  []string{"hello", "world"},
			extraA: nil,
			extraB: nil,
		},
		{
			name:   "equal different order",
			listA:  []string{"hello", "world"},
			listB:  []string{"world", "hello"},
			extraA: nil,
			extraB: nil,
		},
		{
			name:   "extra A",
			listA:  []string{"hello", "hello", "world"},
			listB:  []string{"hello", "world"},
			extraA: []any{"hello"},
			extraB: nil,
		},
		{
			name:   "extra A twice",
			listA:  []string{"hello", "hello", "hello", "world"},
			listB:  []string{"hello", "world"},
			extraA: []any{"hello", "hello"},
			extraB: nil,
		},
		{
			name:   "extra B",
			listA:  []string{"hello", "world"},
			listB:  []string{"hello", "hello", "world"},
			extraA: nil,
			extraB: []any{"hello"},
		},
		{
			name:   "extra B twice",
			listA:  []string{"hello", "world"},
			listB:  []string{"hello", "hello", "world", "hello"},
			extraA: nil,
			extraB: []any{"hello", "hello"},
		},
		{
			name:   "integers 1",
			listA:  []int{1, 2, 3, 4, 5},
			listB:  []int{5, 4, 3, 2, 1},
			extraA: nil,
			extraB: nil,
		},
		{
			name:   "integers 2",
			listA:  []int{1, 2, 1, 2, 1},
			listB:  []int{2, 1, 2, 1, 2},
			extraA: []any{1},
			extraB: []any{2},
		},
	})
}
