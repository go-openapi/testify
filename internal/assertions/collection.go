// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

// Len asserts that the specified object has specific length.
//
// Len also fails if the object has a type that len() does not accept.
//
// The asserted object can be a string, a slice, a map, an array or a channel.
//
// See also [reflect.Len].
//
// Usage:
//
//	assertions.Len(t, mySlice, 3)
//	assertions.Len(t, myString, 4)
//	assertions.Len(t, myMap, 5)
//
// Examples:
//
//	success: []string{"A","B"}, 2
//	failure: []string{"A","B"}, 1
func Len(t T, object any, length int, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	l, ok := getLen(object)
	if !ok {
		return Fail(t, fmt.Sprintf("%q could not be applied builtin len()", truncatingFormat("%v", object)), msgAndArgs...)
	}

	if l != length {
		return Fail(t, fmt.Sprintf("%q should have %d item(s), but has %d", truncatingFormat("%v", object), length, l), msgAndArgs...)
	}
	return true
}

// Contains asserts that the specified string, list(array, slice...) or map contains the
// specified substring or element.
//
// Usage:
//
//	assertions.Contains(t, "Hello World", "World")
//	assertions.Contains(t, ["Hello", "World"], "World")
//	assertions.Contains(t, {"Hello": "World"}, "Hello")
//
// Examples:
//
//	success: []string{"A","B"}, "A"
//	failure: []string{"A","B"}, "C"
func Contains(t T, s, contains any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	ok, found := containsElement(s, contains)
	if !ok {
		return Fail(t, truncatingFormat("%#v", s)+" could not be applied builtin len()", msgAndArgs...)
	}
	if !found {
		return Fail(t, fmt.Sprintf("%s does not contain %#v", truncatingFormat("%#v", s), contains), msgAndArgs...)
	}

	return true
}

// NotContains asserts that the specified string, list(array, slice...) or map does NOT contain the
// specified substring or element.
//
// Usage:
//
//	assertions.NotContains(t, "Hello World", "Earth")
//	assertions.NotContains(t, ["Hello", "World"], "Earth")
//	assertions.NotContains(t, {"Hello": "World"}, "Earth")
//
// Examples:
//
//	success: []string{"A","B"}, "C"
//	failure: []string{"A","B"}, "B"
func NotContains(t T, s, contains any, msgAndArgs ...any) bool {
	if h, ok := t.(H); ok {
		h.Helper()
	}

	ok, found := containsElement(s, contains)
	if !ok {
		return Fail(t, truncatingFormat("%#v", s)+" could not be applied builtin len()", msgAndArgs...)
	}
	if found {
		return Fail(t, fmt.Sprintf("%s should not contain %#v", truncatingFormat("%#v", s), contains), msgAndArgs...)
	}

	return true
}

// Subset asserts that the list (array, slice, or map) contains all elements
// given in the subset (array, slice, or map).
//
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
//	assert.Subset(t, [1, 2, 3], [1, 2])
//	assert.Subset(t, {"x": 1, "y": 2}, {"x": 1})
//	assert.Subset(t, [1, 2, 3], {1: "one", 2: "two"})
//	assert.Subset(t, {"x": 1, "y": 2}, ["x"])
//
// Examples:
//
//	success: []int{1, 2, 3}, []int{1, 2}
//	failure: []int{1, 2, 3}, []int{4, 5}
func Subset(t T, list, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if subset == nil {
		return true // we consider nil to be equal to the nil set
	}

	listKind := reflect.TypeOf(list).Kind()
	if listKind != reflect.Array && listKind != reflect.Slice && listKind != reflect.Map {
		return Fail(t, fmt.Sprintf("%q has an unsupported type %s", list, listKind), msgAndArgs...)
	}

	subsetKind := reflect.TypeOf(subset).Kind()
	if subsetKind != reflect.Array && subsetKind != reflect.Slice && subsetKind != reflect.Map {
		return Fail(t, fmt.Sprintf("%q has an unsupported type %s", subset, subsetKind), msgAndArgs...)
	}

	if subsetKind == reflect.Map && listKind == reflect.Map {
		subsetMap := reflect.ValueOf(subset)
		actualMap := reflect.ValueOf(list)

		for _, k := range subsetMap.MapKeys() {
			ev := subsetMap.MapIndex(k)
			av := actualMap.MapIndex(k)

			if !av.IsValid() {
				return Fail(t, fmt.Sprintf("%s does not contain %s", truncatingFormat("%#v", list), truncatingFormat("%#v", subset)), msgAndArgs...)
			}
			if !ObjectsAreEqual(ev.Interface(), av.Interface()) {
				return Fail(t, fmt.Sprintf("%s does not contain %s", truncatingFormat("%#v", list), truncatingFormat("%#v", subset)), msgAndArgs...)
			}
		}

		return true
	}

	subsetList := reflect.ValueOf(subset)
	if subsetKind == reflect.Map {
		keys := make([]any, subsetList.Len())
		for idx, key := range subsetList.MapKeys() {
			keys[idx] = key.Interface()
		}
		subsetList = reflect.ValueOf(keys)
	}

	for i := range subsetList.Len() {
		element := subsetList.Index(i).Interface()
		ok, found := containsElement(list, element)
		if !ok {
			return Fail(t, fmt.Sprintf("%#v could not be applied builtin len()", list), msgAndArgs...)
		}
		if !found {
			return Fail(t, fmt.Sprintf("%s does not contain %#v", truncatingFormat("%#v", list), element), msgAndArgs...)
		}
	}

	return true
}

// NotSubset asserts that the list (array, slice, or map) does NOT contain all
// elements given in the subset (array, slice, or map).
// Map elements are key-value pairs unless compared with an array or slice where
// only the map key is evaluated.
//
//	assert.NotSubset(t, [1, 3, 4], [1, 2])
//	assert.NotSubset(t, {"x": 1, "y": 2}, {"z": 3})
//	assert.NotSubset(t, [1, 3, 4], {1: "one", 2: "two"})
//	assert.NotSubset(t, {"x": 1, "y": 2}, ["z"])
//
// Examples:
//
//	success: []int{1, 2, 3}, []int{4, 5}
//	failure: []int{1, 2, 3}, []int{1, 2}
func NotSubset(t T, list, subset any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if subset == nil {
		return Fail(t, "nil is the empty set which is a subset of every set", msgAndArgs...)
	}

	listKind := reflect.TypeOf(list).Kind()
	if listKind != reflect.Array && listKind != reflect.Slice && listKind != reflect.Map {
		return Fail(t, fmt.Sprintf("%q has an unsupported type %s", list, listKind), msgAndArgs...)
	}

	subsetKind := reflect.TypeOf(subset).Kind()
	if subsetKind != reflect.Array && subsetKind != reflect.Slice && subsetKind != reflect.Map {
		return Fail(t, fmt.Sprintf("%q has an unsupported type %s", subset, subsetKind), msgAndArgs...)
	}

	if subsetKind == reflect.Map && listKind == reflect.Map {
		subsetMap := reflect.ValueOf(subset)
		actualMap := reflect.ValueOf(list)

		for _, k := range subsetMap.MapKeys() {
			ev := subsetMap.MapIndex(k)
			av := actualMap.MapIndex(k)

			if !av.IsValid() {
				return true
			}
			if !ObjectsAreEqual(ev.Interface(), av.Interface()) {
				return true
			}
		}

		return Fail(t, fmt.Sprintf("%s is a subset of %s", truncatingFormat("%q", subset), truncatingFormat("%q", list)), msgAndArgs...)
	}

	subsetList := reflect.ValueOf(subset)
	if subsetKind == reflect.Map {
		keys := make([]any, subsetList.Len())
		for idx, key := range subsetList.MapKeys() {
			keys[idx] = key.Interface()
		}
		subsetList = reflect.ValueOf(keys)
	}
	for i := range subsetList.Len() {
		element := subsetList.Index(i).Interface()
		ok, found := containsElement(list, element)
		if !ok {
			return Fail(t, fmt.Sprintf("%q could not be applied builtin len()", list), msgAndArgs...)
		}
		if !found {
			return true
		}
	}

	return Fail(t, fmt.Sprintf("%s is a subset of %s", truncatingFormat("%q", subset), truncatingFormat("%q", list)), msgAndArgs...)
}

// ElementsMatch asserts that the specified listA(array, slice...) is equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should match.
//
// assert.ElementsMatch(t, [1, 3, 2, 3], [1, 3, 3, 2]).
//
// Examples:
//
//	success: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
//	failure: []int{1, 2, 3}, []int{1, 2, 4}
func ElementsMatch(t T, listA, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if isEmpty(listA) && isEmpty(listB) {
		return true
	}

	if !isList(t, listA, msgAndArgs...) || !isList(t, listB, msgAndArgs...) {
		return false
	}

	extraA, extraB := diffLists(listA, listB)

	if len(extraA) == 0 && len(extraB) == 0 {
		return true
	}

	return Fail(t, formatListDiff(listA, listB, extraA, extraB), msgAndArgs...)
}

// containsElement tries to loop over the list check if the list includes the element.
//
// return (false, false) if impossible.
// return (true, false) if element was not found.
// return (true, true) if element was found.
func containsElement(list any, element any) (ok, found bool) {
	listValue := reflect.ValueOf(list)
	listType := reflect.TypeOf(list)
	if listType == nil {
		return false, false
	}
	listKind := listType.Kind()
	defer func() {
		if e := recover(); e != nil {
			ok = false
			found = false
		}
	}()

	if listKind == reflect.String {
		elementValue := reflect.ValueOf(element)
		return true, strings.Contains(listValue.String(), elementValue.String())
	}

	if listKind == reflect.Map {
		mapKeys := listValue.MapKeys()
		for i := range mapKeys {
			if ObjectsAreEqual(mapKeys[i].Interface(), element) {
				return true, true
			}
		}
		return true, false
	}

	for i := range listValue.Len() {
		if ObjectsAreEqual(listValue.Index(i).Interface(), element) {
			return true, true
		}
	}
	return true, false
}

// isList checks that the provided value is array or slice.
func isList(t T, list any, msgAndArgs ...any) (ok bool) {
	kind := reflect.TypeOf(list).Kind()
	if kind != reflect.Array && kind != reflect.Slice {
		return Fail(t, fmt.Sprintf("%q has an unsupported type %s, expecting array or slice", list, kind),
			msgAndArgs...)
	}
	return true
}

// diffLists diffs two arrays/slices and returns slices of elements that are only in A and only in B.
//
// If some element is present multiple times, each instance is counted separately (e.g. if something is 2x in A and
// 5x in B, it will be 0x in extraA and 3x in extraB). The order of items in both lists is ignored.
func diffLists(listA, listB any) (extraA, extraB []any) {
	aValue := reflect.ValueOf(listA)
	bValue := reflect.ValueOf(listB)

	aLen := aValue.Len()
	bLen := bValue.Len()

	// Mark indexes in bValue that we already used
	visited := make([]bool, bLen)
	for i := range aLen {
		element := aValue.Index(i).Interface()
		found := false
		for j := range bLen {
			if visited[j] {
				continue
			}
			if ObjectsAreEqual(bValue.Index(j).Interface(), element) {
				visited[j] = true
				found = true
				break
			}
		}
		if !found {
			extraA = append(extraA, element)
		}
	}

	for j := range bLen {
		if visited[j] {
			continue
		}
		extraB = append(extraB, bValue.Index(j).Interface())
	}

	return extraA, extraB
}

func formatListDiff(listA, listB any, extraA, extraB []any) string {
	var msg bytes.Buffer

	msg.WriteString("elements differ")
	if len(extraA) > 0 {
		msg.WriteString("\n\nextra elements in list A:\n")
		msg.WriteString(spewConfig.Sdump(extraA))
	}
	if len(extraB) > 0 {
		msg.WriteString("\n\nextra elements in list B:\n")
		msg.WriteString(spewConfig.Sdump(extraB))
	}
	msg.WriteString("\n\nlistA:\n")
	msg.WriteString(spewConfig.Sdump(listA))
	msg.WriteString("\n\nlistB:\n")
	msg.WriteString(spewConfig.Sdump(listB))

	return msg.String()
}

// NotElementsMatch asserts that the specified listA(array, slice...) is NOT equal to specified
// listB(array, slice...) ignoring the order of the elements. If there are duplicate elements,
// the number of appearances of each of them in both lists should not match.
// This is an inverse of ElementsMatch.
//
// assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 1, 2, 3]) -> false
//
// assert.NotElementsMatch(t, [1, 1, 2, 3], [1, 2, 3]) -> true
//
// assert.NotElementsMatch(t, [1, 2, 3], [1, 2, 4]) -> true.
//
// Examples:
//
//	success: []int{1, 2, 3}, []int{1, 2, 4}
//	failure: []int{1, 3, 2, 3}, []int{1, 3, 3, 2}
func NotElementsMatch(t T, listA, listB any, msgAndArgs ...any) (ok bool) {
	if h, ok := t.(H); ok {
		h.Helper()
	}
	if isEmpty(listA) && isEmpty(listB) {
		return Fail(t, "listA and listB contain the same elements", msgAndArgs)
	}

	if !isList(t, listA, msgAndArgs...) {
		return Fail(t, "listA is not a list type", msgAndArgs...)
	}
	if !isList(t, listB, msgAndArgs...) {
		return Fail(t, "listB is not a list type", msgAndArgs...)
	}

	extraA, extraB := diffLists(listA, listB)
	if len(extraA) == 0 && len(extraB) == 0 {
		return Fail(t, "listA and listB contain the same elements", msgAndArgs)
	}

	return true
}

// getLen tries to get the length of an object.
// It returns (0, false) if impossible.
func getLen(x any) (length int, ok bool) {
	v := reflect.ValueOf(x)
	defer func() {
		ok = recover() == nil
	}()
	return v.Len(), true
}
