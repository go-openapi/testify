// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"errors"
	"fmt"
	"iter"
	"reflect"
	"slices"
	"strings"
	"testing"
)

// TestCollectionLen tests the Len assertion.
func TestCollectionLen(t *testing.T) {
	t.Parallel()

	for tc := range collectionLenCases() {
		t.Run(tc.name, testLen(tc))
	}
}

// TestCollectionContains tests both Contains and NotContains with reflection-based
// and generic variants using unified test cases.
//
// For slices, also tests SeqContains
// and SeqNotContains since slices can be converted to iter.Seq via slices.Values.
func TestCollectionContains(t *testing.T) {
	t.Parallel()

	for tc := range unifiedContainsCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			container, element := tc.makeValues()

			if !tc.reflectionOnly {
				// Test reflection-based variants
				t.Run("with Contains", testContainsAssertion(tc, containsKind, Contains, container, element))
				t.Run("with NotContains", testContainsAssertion(tc, notContainsKind, NotContains, container, element))

				// Test generic variants (type dispatch)
				t.Run("with generic Contains", testContainsAssertionT(tc, containsTKind, container, element))
				t.Run("with generic NotContains", testContainsAssertionT(tc, notContainsTKind, container, element))

				// For slices, also test Seq variants (slices can be converted to iter.Seq via slices.Values)
				if isSliceType(container) {
					t.Run("with generic SeqContains", testContainsAssertionT(tc, seqContainsTKind, container, element))
					t.Run("with generic SeqNotContains", testContainsAssertionT(tc, seqNotContainsTKind, container, element))
				}

				return
			}

			// Reflection-only cases
			t.Run("with Contains (reflection)", testContainsAssertion(tc, containsKind, Contains, container, element))
			t.Run("with NotContains (reflection)", testContainsAssertion(tc, notContainsKind, NotContains, container, element))
		})
	}
}

// TestCollectionSubset tests both Subset and NotSubset with reflection-based
// and generic variants using unified test cases.
func TestCollectionSubset(t *testing.T) {
	t.Parallel()

	for tc := range unifiedSubsetCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			list, subset := tc.makeValues()

			if !tc.reflectionOnly {
				// Test reflection-based variants
				t.Run("with Subset", testSubsetAssertion(tc, subsetKind, Subset, list, subset))
				t.Run("with NotSubset", testSubsetAssertion(tc, notSubsetKind, NotSubset, list, subset))

				// Test generic variants (type dispatch)
				t.Run("with generic Subset", testSubsetAssertionT(tc, subsetTKind, list, subset))
				t.Run("with generic NotSubset", testSubsetAssertionT(tc, notSubsetTKind, list, subset))
			} else {
				// Reflection-only cases
				t.Run("with Subset (reflection)", testSubsetAssertion(tc, subsetKind, Subset, list, subset))
				t.Run("with NotSubset (reflection)", testSubsetAssertion(tc, notSubsetKind, NotSubset, list, subset))
			}
		})
	}
}

// TestCollectionElementsMatch tests both ElementsMatch and NotElementsMatch
// with reflection-based and generic variants using unified test cases.
func TestCollectionElementsMatch(t *testing.T) {
	t.Parallel()

	for tc := range unifiedElementsMatchCases() {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, expected := tc.makeValues()

			if !tc.reflectionOnly {
				// Test reflection-based variants
				t.Run("with ElementsMatch", testElementsMatchAssertion(tc, elementsMatchKind, ElementsMatch, actual, expected))
				t.Run("with NotElementsMatch", testElementsMatchAssertion(tc, notElementsMatchKind, NotElementsMatch, actual, expected))

				// Test generic variants (type dispatch)
				t.Run("with generic ElementsMatch", testElementsMatchAssertionT(tc, elementsMatchTKind, actual, expected))
				t.Run("with generic NotElementsMatch", testElementsMatchAssertionT(tc, notElementsMatchTKind, actual, expected))
			} else {
				// Reflection-only cases
				t.Run("with ElementsMatch (reflection)", testElementsMatchAssertion(tc, elementsMatchKind, ElementsMatch, actual, expected))
				t.Run("with NotElementsMatch (reflection)", testElementsMatchAssertion(tc, notElementsMatchKind, NotElementsMatch, actual, expected))
			}
		})
	}
}

// TestCollectionErrorMessages tests error message formatting for collection assertions.
func TestCollectionErrorMessages(t *testing.T) {
	t.Parallel()

	runFailCases(t, collectionErrorMessageCases())
}

// ============================================================================
// TestCollectionLen
// ============================================================================

func testLen(tc collectionLenCase) func(*testing.T) {
	if !tc.valid {
		return func(t *testing.T) {
			t.Run("with invalid type", func(t *testing.T) {
				t.Run(tc.name, func(t *testing.T) {
					t.Parallel()

					mock := new(mockT)
					res := Len(mock, tc.v, tc.len)
					if res {
						t.Errorf("Len should not work for type %T", tc.v)

						return
					}

					if tc.expectedMsg == "" {
						return // skip error message check
					}

					// check error message
					if !strings.Contains(mock.errorString(), tc.expectedMsg) {
						t.Errorf("expected error message to contain %q but got: %q", tc.expectedMsg, mock.errorString())
					}
				})
			})
		}
	}

	return func(t *testing.T) {
		t.Run("with expected length", func(t *testing.T) {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				mock := new(mockT)
				res := Len(mock, tc.v, tc.len)
				if !res {
					t.Errorf("%#v should have %d items", tc.v, tc.len)
				}
			})
		})

		t.Run("with unexpected length", func(t *testing.T) {
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				mock := new(mockT)
				res := Len(mock, tc.v, tc.len+1)
				if res {
					t.Errorf("%#v should not have %d items", tc.v, tc.len+1)
					return
				}

				if tc.expectedMsg == "" {
					return // skip error message check
				}

				// check error message
				if !strings.Contains(mock.errorString(), tc.expectedMsg) {
					t.Errorf("expected error message to contain %q but got: %q", tc.expectedMsg, mock.errorString())
				}
			})
		})
	}
}

type collectionLenCase struct {
	name        string
	v           any
	len         int
	expectedMsg string // message when expecting len+1 items
	valid       bool
}

func collectionLenCases() iter.Seq[collectionLenCase] {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	longSlice := make([]int, 1_000_000)
	arr := [3]int{1, 2, 3}

	return slices.Values([]collectionLenCase{
		{"slice/int", []int{1, 2, 3}, 3, `"[1 2 3]" should have 4 item(s), but has 3`, true},
		{"array/int", [...]int{1, 2, 3}, 3, `"[1 2 3]" should have 4 item(s), but has 3`, true},
		{"ptr-to-array/int", &arr, 3, `"&[1 2 3]" should have 4 item(s), but has 3`, true},
		{"string", "ABC", 3, `"ABC" should have 4 item(s), but has 3`, true},
		{"map/int", map[int]int{1: 2, 2: 4, 3: 6}, 3, `"map[1:2 2:4 3:6]" should have 4 item(s), but has 3`, true},
		{"channel", ch, 3, "", true},
		{"empty slice", []int{}, 0, `"[]" should have 1 item(s), but has 0`, true},
		{"empty map", map[int]int{}, 0, `"map[]" should have 1 item(s), but has 0`, true},
		{"empty channel", make(chan int), 0, "", true},
		{"nil slice", []int(nil), 0, `"[]" should have 1 item(s), but has 0`, true},
		{"nil map", map[int]int(nil), 0, `"map[]" should have 1 item(s), but has 0`, true},
		{"nil chan", (chan int)(nil), 0, `"<nil>" should have 1 item(s), but has 0`, true},

		// Unsupported types
		{"invalid type/nil", nil, 0, `"<nil>" could not be applied builtin len()`, false},
		{"invalid type/int", 0, 0, "", false},
		{"invalid type/bool", true, 0, "", false},
		{"invalid type/rune", 'A', 0, "", false},
		{"invalid type/struct", struct{}{}, 0, "", false},
		{"invalid type/ptr-not-array", &longSlice, 1_000_000, `<... truncated>" could not be applied builtin len()`, false},
		{"invalid type/ptr-anything", ptr(1), 0, `" could not be applied builtin len()`, false},

		// Truncated message
		{"truncated message/long slice", longSlice, 1_000_000, `<... truncated>" should have 1000001 item(s), but has 1000000`, true},
	})
}

// ============================================================================
// TestCollectionContains
// ============================================================================

// containsRelationship describes the relationship between a container and an element.
type containsRelationship int

const (
	crContains containsRelationship = iota
	crNotContains
	crInvalidContainer // container type doesn't support Contains
)

type containsAssertionKind int

const (
	containsKind containsAssertionKind = iota
	notContainsKind
	containsTKind       // generic string/slice/map variants
	notContainsTKind    // generic not-contains variants
	seqContainsTKind    // generic seq contains
	seqNotContainsTKind // generic seq not-contains
)

func testContainsAssertion(tc containsTestCase, kind containsAssertionKind, assertion func(T, any, any, ...any) bool, container, element any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := assertion(mock, container, element)
		shouldPass := expectedStatusForContainsAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

//nolint:gocognit,gocyclo,cyclop // type dispatch requires large switch statement
func testContainsAssertionT(tc containsTestCase, kind containsAssertionKind, container, element any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		stop := func(expected string, actual any) {
			t.Fatalf("test case error: expected container=%s, actual=%T", expected, actual)
		}

		var result bool
		// Type switch based on container type
		switch cont := container.(type) {
		case string:
			elem, ok := element.(string)
			if !ok {
				t.Fatalf("test case error: string container requires string element, got %T", element)
			}
			result = testStringContainsGeneric(mock, kind, cont, elem)
		case []byte:
			elem, ok := element.([]byte)
			if !ok {
				t.Fatalf("test case error: []byte container requires []byte element, got %T", element)
			}
			result = testStringContainsGeneric(mock, kind, cont, elem)
		case []int:
			elem, ok := element.(int)
			if !ok {
				t.Fatalf("test case error: []int container requires int element, got %T", element)
			}
			if kind == seqContainsTKind || kind == seqNotContainsTKind {
				result = testSeqContainsGeneric(mock, kind, cont, elem)
			} else {
				result = testSliceContainsGeneric(mock, kind, cont, elem)
			}
		case []string:
			elem, ok := element.(string)
			if !ok {
				t.Fatalf("test case error: []string container requires string element, got %T", element)
			}
			if kind == seqContainsTKind || kind == seqNotContainsTKind {
				result = testSeqContainsGeneric(mock, kind, cont, elem)
			} else {
				result = testSliceContainsGeneric(mock, kind, cont, elem)
			}
		case []float64:
			elem, ok := element.(float64)
			if !ok {
				t.Fatalf("test case error: []float64 container requires float64 element, got %T", element)
			}
			if kind == seqContainsTKind || kind == seqNotContainsTKind {
				result = testSeqContainsGeneric(mock, kind, cont, elem)
			} else {
				result = testSliceContainsGeneric(mock, kind, cont, elem)
			}
		case []*containsStruct:
			elem, ok := element.(*containsStruct)
			if !ok {
				t.Fatalf("test case error: []*containsStruct container requires *containsStruct element, got %T", element)
			}
			if kind == seqContainsTKind || kind == seqNotContainsTKind {
				result = testSeqContainsGeneric(mock, kind, cont, elem)
			} else {
				result = testSliceContainsGeneric(mock, kind, cont, elem)
			}

		case map[string]int:
			elem, ok := element.(string)
			if !ok {
				t.Fatalf("test case error: map[string]int container requires string element, got %T", element)
			}
			result = testMapContainsGeneric(mock, kind, cont, elem)
		case map[int]string:
			elem, ok := element.(int)
			if !ok {
				t.Fatalf("test case error: map[int]string container requires int element, got %T", element)
			}
			result = testMapContainsGeneric(mock, kind, cont, elem)
		case map[any]any:
			result = testMapContainsGeneric(mock, kind, cont, element)
		default:
			stop(fmt.Sprintf("%T", container), container)
		}

		shouldPass := expectedStatusForContainsAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

func testStringContainsGeneric[S, E Text](mock T, kind containsAssertionKind, container S, element E) bool {
	switch kind {
	case containsTKind:
		return StringContainsT(mock, container, element)
	case notContainsTKind:
		return StringNotContainsT(mock, container, element)
	default:
		panic(fmt.Errorf("test case configuration error: invalid containsAssertionKind for string generic: %d", kind))
	}
}

func testSliceContainsGeneric[Slice ~[]E, E comparable](mock T, kind containsAssertionKind, container Slice, element E) bool {
	switch kind {
	case containsTKind:
		return SliceContainsT(mock, container, element)
	case notContainsTKind:
		return SliceNotContainsT(mock, container, element)
	default:
		panic(fmt.Errorf("test case configuration error: invalid containsAssertionKind for slice generic: %d", kind))
	}
}

func testMapContainsGeneric[Map ~map[K]V, K comparable, V any](mock T, kind containsAssertionKind, container Map, element K) bool {
	switch kind {
	case containsTKind:
		return MapContainsT(mock, container, element)
	case notContainsTKind:
		return MapNotContainsT(mock, container, element)
	default:
		panic(fmt.Errorf("test case configuration error: invalid containsAssertionKind for map generic: %d", kind))
	}
}

func testSeqContainsGeneric[Slice ~[]E, E comparable](mock T, kind containsAssertionKind, container Slice, element E) bool {
	seq := slices.Values(container)
	switch kind {
	case seqContainsTKind:
		return SeqContainsT(mock, seq, element)
	case seqNotContainsTKind:
		return SeqNotContainsT(mock, seq, element)
	default:
		panic(fmt.Errorf("test case configuration error: invalid containsAssertionKind for seq generic: %d", kind))
	}
}

func expectedStatusForContainsAssertion(kind containsAssertionKind, relationship containsRelationship) bool {
	positive := kind == containsKind || kind == containsTKind || kind == seqContainsTKind

	switch relationship {
	case crContains:
		return positive
	case crNotContains:
		return !positive
	case crInvalidContainer:
		return false
	default:
		panic(fmt.Errorf("test case configuration error: invalid containsRelationship: %d", relationship))
	}
}

func isSliceType(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}

type containsTestCase struct {
	name           string
	makeValues     func() (container, element any)
	relationship   containsRelationship
	reflectionOnly bool
}

type containsStruct struct {
	Name, Value string
}

const (
	testStringBar = "Bar"
	testStringFoo = "Foo"
)

func unifiedContainsCases() iter.Seq[containsTestCase] {
	list := []string{testStringFoo, testStringBar}
	complexList := []*containsStruct{
		{"b", "c"},
		{"d", "e"},
		{"g", "h"},
		{"j", "k"},
	}
	simpleMap := map[any]any{testStringFoo: testStringBar}
	var zeroMap map[any]any

	return slices.Values([]containsTestCase{
		// String contains
		{"string/contains", func() (any, any) { return "Hello World", "Hello" }, crContains, false},
		{"string/not-contains", func() (any, any) { return "Hello World", "Salut" }, crNotContains, false},

		// Slice contains
		{"slice-string/contains", func() (any, any) { return list, testStringBar }, crContains, false},
		{"slice-string/not-contains", func() (any, any) { return list, "Salut" }, crNotContains, false},
		// Struct pointers use reflection-only since generic uses pointer equality, not deep equality
		{"slice-struct/contains", func() (any, any) { return complexList, &containsStruct{"g", "h"} }, crContains, true},
		{"slice-struct/not-contains", func() (any, any) { return complexList, &containsStruct{"g", "e"} }, crNotContains, true},

		// Map contains (key lookup)
		{"map/contains-key", func() (any, any) { return simpleMap, testStringFoo }, crContains, false},
		{"map/not-contains-key", func() (any, any) { return simpleMap, testStringBar }, crNotContains, false},
		{"map-zero/not-contains", func() (any, any) { return zeroMap, testStringBar }, crNotContains, false},

		// Invalid container (reflection only)
		{
			"invalid/non-container-struct",
			func() (any, any) {
				type nonContainer struct{ Value string }
				return nonContainer{Value: "Hello"}, "Hello"
			},
			crInvalidContainer,
			true,
		},
		{
			"invalid/nil",
			func() (any, any) { return nil, "key" },
			crInvalidContainer,
			true,
		},
	})
}

// ============================================================================
// TestCollectionSubset
// ============================================================================

// subsetRelationship describes the relationship between a list and a subset.
type subsetRelationship int

const (
	srSubset subsetRelationship = iota
	srNotSubset
	srInvalidType // types don't support subset comparison
)

type subsetAssertionKind int

const (
	subsetKind subsetAssertionKind = iota
	notSubsetKind
	subsetTKind    // generic slice variant
	notSubsetTKind // generic not-subset variant
)

func testSubsetAssertion(tc subsetTestCase, kind subsetAssertionKind, assertion func(T, any, any, ...any) bool, list, subset any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := assertion(mock, list, subset)
		shouldPass := expectedStatusForSubsetAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

func testSubsetAssertionT(tc subsetTestCase, kind subsetAssertionKind, list, subset any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		var result bool
		// Type switch based on list type
		switch lst := list.(type) {
		case []int:
			sub, ok := subset.([]int)
			if !ok {
				t.Fatalf("test case error: []int list requires []int subset, got %T", subset)
			}
			result = testSliceSubsetGeneric(mock, kind, lst, sub)
		case []string:
			sub, ok := subset.([]string)
			if !ok {
				t.Fatalf("test case error: []string list requires []string subset, got %T", subset)
			}
			result = testSliceSubsetGeneric(mock, kind, lst, sub)
		case []float64:
			sub, ok := subset.([]float64)
			if !ok {
				t.Fatalf("test case error: []float64 list requires []float64 subset, got %T", subset)
			}
			result = testSliceSubsetGeneric(mock, kind, lst, sub)
		default:
			t.Fatalf("unsupported type for generic subset: %T", list)
		}

		shouldPass := expectedStatusForSubsetAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

func testSliceSubsetGeneric[Slice ~[]E, E comparable](mock T, kind subsetAssertionKind, list, subset Slice) bool {
	switch kind {
	case subsetTKind:
		return SliceSubsetT(mock, list, subset)
	case notSubsetTKind:
		return SliceNotSubsetT(mock, list, subset)
	default:
		panic(fmt.Errorf("test case configuration error: invalid subsetAssertionKind for generic: %d", kind))
	}
}

func expectedStatusForSubsetAssertion(kind subsetAssertionKind, relationship subsetRelationship) bool {
	positive := kind == subsetKind || kind == subsetTKind

	switch relationship {
	case srSubset:
		return positive
	case srNotSubset:
		return !positive
	case srInvalidType:
		return false
	default:
		panic(fmt.Errorf("test case configuration error: invalid subsetRelationship: %d", relationship))
	}
}

type subsetTestCase struct {
	name           string
	makeValues     func() (list, subset any)
	relationship   subsetRelationship
	reflectionOnly bool
}

func unifiedSubsetCases() iter.Seq[subsetTestCase] {
	return slices.Values([]subsetTestCase{
		// Subset cases
		{"int/nil-subset", func() (any, any) { return []int{1, 2, 3}, ([]int)(nil) }, srSubset, false},
		{"int/empty-subset", func() (any, any) { return []int{1, 2, 3}, []int{} }, srSubset, false},
		{"int/proper-subset", func() (any, any) { return []int{1, 2, 3}, []int{1, 2} }, srSubset, false},
		{"int/equal-sets", func() (any, any) { return []int{1, 2, 3}, []int{1, 2, 3} }, srSubset, false},
		{"string/subset", func() (any, any) { return []string{"hello", "world"}, []string{"hello"} }, srSubset, false},
		{"float64/subset", func() (any, any) { return []float64{1.1, 2.2, 3.3}, []float64{2.2} }, srSubset, false},
		{
			"map-string/subset",
			func() (any, any) {
				return map[string]string{"a": "x", "b": "y", "c": "z"},
					map[string]string{"a": "x", "b": "y"}
			},
			srSubset,
			true,
		},
		{
			"slice-map-mixed/subset",
			func() (any, any) {
				return []string{"a", "b", "c"}, map[string]int{"a": 1, "c": 3}
			},
			srSubset,
			true,
		},

		// Not subset cases
		{
			"string/not-subset",
			func() (any, any) {
				return []string{"hello", "world"}, []string{"hello", "testify"}
			},
			srNotSubset,
			false,
		},
		{"int/not-subset", func() (any, any) { return []int{1, 2, 3}, []int{4, 5} }, srNotSubset, false},
		{"int/partial-not-subset", func() (any, any) { return []int{1, 2, 3}, []int{1, 5} }, srNotSubset, false},
		{
			"map-string/not-subset",
			func() (any, any) {
				return map[string]string{"a": "x", "b": "y", "c": "z"},
					map[string]string{"a": "x", "b": "z"}
			},
			srNotSubset,
			true,
		},
		{
			"map-string/superset-not-subset",
			func() (any, any) {
				return map[string]string{"a": "x", "b": "y"},
					map[string]string{"a": "x", "b": "y", "c": "z"}
			},
			srNotSubset,
			true,
		},
		{
			"slice-map-mixed/not-subset",
			func() (any, any) {
				return []string{"a", "b", "c"}, map[string]int{"c": 3, "d": 4}
			},
			srNotSubset,
			true,
		},
		{
			"subset/nil",
			func() (any, any) {
				return []string{"a", "b", "c"}, nil
			},
			srSubset,
			true,
		},
		{
			"subset/empty",
			func() (any, any) {
				return []string{"a", "b", "c"}, []string{}
			},
			srSubset,
			true,
		},
		{
			"not-subset/nil",
			func() (any, any) {
				return nil, []string{"a", "b", "c"}
			},
			srNotSubset,
			true,
		},
		{
			"not-subset/empty",
			func() (any, any) {
				return []string{}, []string{"a", "b", "c"}
			},
			srNotSubset,
			true,
		},
		{
			"subset/nil-nil",
			func() (any, any) {
				return nil, nil
			},
			srSubset,
			true,
		},
		{
			"subset/empty-nil",
			func() (any, any) {
				return []int{}, nil
			},
			srSubset,
			true,
		},
		{
			"subset/nil-empty",
			func() (any, any) {
				return nil, []int{}
			},
			srSubset,
			true,
		},
		{
			"subset/empty-empty",
			func() (any, any) {
				return []int{}, []int{}
			},
			srSubset,
			true,
		},
		{
			"invalid-type/[]int-invalid",
			func() (any, any) {
				return []int{}, 1
			},
			srInvalidType,
			true,
		},
		{
			"invalid-type/invalid-[]int",
			func() (any, any) {
				return 1, []int{}
			},
			srInvalidType,
			true,
		},
		{
			"invalid-type/invalid-nil",
			func() (any, any) {
				return 1, nil
			},
			srInvalidType,
			true,
		},
		{
			"invalid-type/nil-invalid",
			func() (any, any) {
				return nil, 1
			},
			srInvalidType,
			true,
		},
	})
}

// ============================================================================
// TestCollectionElementsMatch
// ============================================================================

// elementsMatchRelationship describes the relationship between two collections.
type elementsMatchRelationship int

const (
	emMatch elementsMatchRelationship = iota
	emNotMatch
	emInvalidType // types don't support elements matching
)

type elementsMatchAssertionKind int

const (
	elementsMatchKind elementsMatchAssertionKind = iota
	notElementsMatchKind
	elementsMatchTKind    // generic variant
	notElementsMatchTKind // generic not-match variant
)

func testElementsMatchAssertion(tc elementsMatchTestCase, kind elementsMatchAssertionKind, assertion func(T, any, any, ...any) bool, actual, expected any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)
		result := assertion(mock, actual, expected)
		shouldPass := expectedStatusForElementsMatchAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

func testElementsMatchAssertionT(tc elementsMatchTestCase, kind elementsMatchAssertionKind, actual, expected any) func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		mock := new(mockT)

		var result bool
		// Type switch based on actual type
		switch act := actual.(type) {
		case []int:
			exp, ok := expected.([]int)
			if !ok {
				t.Fatalf("test case error: []int actual requires []int expected, got %T", expected)
			}
			result = testElementsMatchGeneric(mock, kind, act, exp)
		case [2]int:
			exp, ok := expected.([2]int)
			if !ok {
				t.Fatalf("test case error: [2]int actual requires [2]int expected, got %T", expected)
			}
			result = testElementsMatchGeneric(mock, kind, act[:], exp[:])
		case []string:
			exp, ok := expected.([]string)
			if !ok {
				t.Fatalf("test case error: []string actual requires []string expected, got %T", expected)
			}
			result = testElementsMatchGeneric(mock, kind, act, exp)
		case [3]string:
			exp, ok := expected.([3]string)
			if !ok {
				t.Fatalf("test case error: [3]string actual requires [3]string expected, got %T", expected)
			}
			result = testElementsMatchGeneric(mock, kind, act[:], exp[:])
		case nil:
			if expected != nil {
				t.Fatalf("test case error: nil actual requires nil expected, got %T", expected)
			}
			result = testElementsMatchGeneric(mock, kind, ([]int)(nil), ([]int)(nil))
		default:
			t.Fatalf("unsupported type for generic ElementsMatch: %T", actual)
		}

		shouldPass := expectedStatusForElementsMatchAssertion(kind, tc.relationship)
		shouldPassOrFail(t, mock, result, shouldPass)
	}
}

func testElementsMatchGeneric[E comparable](mock T, kind elementsMatchAssertionKind, actual, expected []E) bool {
	switch kind {
	case elementsMatchTKind:
		return ElementsMatchT(mock, actual, expected)
	case notElementsMatchTKind:
		return NotElementsMatchT(mock, actual, expected)
	default:
		panic(fmt.Errorf("test case configuration error: invalid elementsMatchAssertionKind for generic: %d", kind))
	}
}

func expectedStatusForElementsMatchAssertion(kind elementsMatchAssertionKind, relationship elementsMatchRelationship) bool {
	positive := kind == elementsMatchKind || kind == elementsMatchTKind

	switch relationship {
	case emMatch:
		return positive
	case emNotMatch:
		return !positive
	case emInvalidType:
		return false
	default:
		panic(fmt.Errorf("test case configuration error: invalid elementsMatchRelationship: %d", relationship))
	}
}

type elementsMatchTestCase struct {
	name           string
	makeValues     func() (actual, expected any)
	relationship   elementsMatchRelationship
	reflectionOnly bool
}

func unifiedElementsMatchCases() iter.Seq[elementsMatchTestCase] {
	return slices.Values([]elementsMatchTestCase{
		// Matching cases
		{"nil-nil", func() (any, any) { return nil, nil }, emMatch, true}, // reflection only - generic can't infer type from nil
		{"empty-empty", func() (any, any) { return []int{}, []int{} }, emMatch, false},
		{"single-element", func() (any, any) { return []int{1}, []int{1} }, emMatch, false},
		{"duplicates-same", func() (any, any) { return []int{1, 1}, []int{1, 1} }, emMatch, false},
		{"reordered", func() (any, any) { return []int{1, 2}, []int{2, 1} }, emMatch, false},
		{"array-reordered", func() (any, any) { return [2]int{1, 2}, [2]int{2, 1} }, emMatch, false},
		{"string-reordered", func() (any, any) { return []string{"hello", "world"}, []string{"world", "hello"} }, emMatch, false},
		{"string-duplicates", func() (any, any) { return []string{"hello", "hello"}, []string{"hello", "hello"} }, emMatch, false},
		{
			"string-complex-reordered",
			func() (any, any) {
				return []string{"hello", "hello", "world"}, []string{"hello", "world", "hello"}
			},
			emMatch,
			false,
		},
		{
			"array-string-reordered",
			func() (any, any) {
				return [3]string{"hello", "hello", "world"}, [3]string{"hello", "world", "hello"}
			},
			emMatch,
			false,
		},
		{"empty-nil", func() (any, any) { return []int{}, nil }, emMatch, true}, // reflection only - nil type inference

		// Not matching cases
		{"different-count", func() (any, any) { return []int{1}, []int{1, 1} }, emNotMatch, false},
		{"different-values", func() (any, any) { return []int{1, 2}, []int{2, 2} }, emNotMatch, false},
		{"string-different", func() (any, any) { return []string{"hello", "hello"}, []string{"hello"} }, emNotMatch, false},

		// Invalid types (reflection only)
		{"invalid type/[]int-invalid", func() (any, any) { return []int{}, 1 }, emInvalidType, true},
		{"invalid type/invalid-[]int", func() (any, any) { return 1, []int{} }, emInvalidType, true},
	})
}

// ============================================================================
// TestCollectionErrorMessages
// ============================================================================

func collectionErrorMessageCases() iter.Seq[failCase] {
	const pkg = "assertions"
	type nonContainer struct {
		Value string
	}
	longSlice := make([]int, 1_000_000)

	return slices.Values([]failCase{
		// Contains/NotContains fail messages
		{
			name:         "Contains(string, error)",
			assertion:    func(t T) bool { return Contains(t, "Hello World", errors.New("Hello")) },
			wantContains: []string{`"Hello World" does not contain &errors.errorString{s:"Hello"}`},
		},
		{
			name:         "Contains(map, missing-key)",
			assertion:    func(t T) bool { return Contains(t, map[string]int{"one": 1}, "two") },
			wantContains: []string{`map[string]int{"one":1} does not contain "two"`},
		},
		{
			name:         "NotContains(map, present-key)",
			assertion:    func(t T) bool { return NotContains(t, map[string]int{"one": 1}, "one") },
			wantContains: []string{`map[string]int{"one":1} should not contain "one"`},
		},
		{
			name: "Contains(nonContainer, string)",
			assertion: func(t T) bool {
				return Contains(t, nonContainer{Value: "Hello"}, "Hello")
			},
			wantContains: []string{pkg + `.nonContainer{Value:"Hello"} could not be applied builtin len()`},
		},
		{
			name: "NotContains(nonContainer, string)",
			assertion: func(t T) bool {
				return NotContains(t, nonContainer{Value: "Hello"}, "Hello")
			},
			wantContains: []string{pkg + `.nonContainer{Value:"Hello"} could not be applied builtin len()`},
		},

		// nil container
		{
			name:         "Contains(nil, key)",
			assertion:    func(t T) bool { return Contains(t, nil, "key") },
			wantContains: []string{"<nil> could not be applied builtin len()"},
		},
		{
			name:         "NotContains(nil, key)",
			assertion:    func(t T) bool { return NotContains(t, nil, "key") },
			wantContains: []string{"<nil> could not be applied builtin len()"},
		},

		// truncation: too long to print
		truncationCase("truncation/Nil(longSlice)", func(t T) bool {
			return Nil(t, &longSlice)
		}),
		truncationCase("truncation/Empty(longSlice)", func(t T) bool {
			return Empty(t, longSlice)
		}),
		{
			name:      "truncation/Contains(longSlice, 1)",
			assertion: func(t T) bool { return Contains(t, longSlice, 1) },
			wantContains: []string{
				`[]int{0, 0, 0,`,
				`<... truncated> does not contain 1`,
			},
		},
		{
			name:      "truncation/NotContains(longSlice, 0)",
			assertion: func(t T) bool { return NotContains(t, longSlice, 0) },
			wantContains: []string{
				`[]int{0, 0, 0,`,
				`<... truncated> should not contain 0`,
			},
		},
		{
			name:      "truncation/Subset(longSlice, [1])",
			assertion: func(t T) bool { return Subset(t, longSlice, []int{1}) },
			wantContains: []string{
				`[]int{0, 0, 0,`,
				`<... truncated> does not contain 1`,
			},
		},
		{
			name: "truncation/Subset(map-longSlice)",
			assertion: func(t T) bool {
				return Subset(t, map[bool][]int{true: longSlice}, map[bool][]int{false: longSlice})
			},
			wantContains: []string{
				`map[bool][]int{true:[]int{0, 0, 0,`,
				`<... truncated> does not contain map[bool][]int{false:[]int{0, 0, 0,`,
			},
		},
		{
			name:      "truncation/NotSubset(longSlice)",
			assertion: func(t T) bool { return NotSubset(t, longSlice, longSlice) },
			wantContains: []string{
				`['\x00' '\x00' '\x00'`,
				`<... truncated> is a subset of ['\x00' '\x00' '\x00'`,
			},
		},
		{
			name: "truncation/NotSubset(map-longSlice)",
			assertion: func(t T) bool {
				return NotSubset(t, map[int][]int{1: longSlice}, map[int][]int{1: longSlice})
			},
			wantContains: []string{
				`map['\x01':['\x00' '\x00' '\x00'`,
				`<... truncated> is a subset of map['\x01':['\x00' '\x00' '\x00'`,
			},
		},
	})
}
