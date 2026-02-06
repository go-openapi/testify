// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"iter"
	"slices"
	"testing"
)

func TestCollectionUnexportedImplementationDetails(t *testing.T) {
	t.Parallel()

	t.Run("containsElement", testContainsElement())
	t.Run("getLen", testGetLen())
}

func testContainsElement() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		list1 := []string{"Foo", "Bar"}
		list2 := []int{1, 2}
		simpleMap := map[any]any{"Foo": "Bar"}

		checkContains := func(list any, elem any, expectOK, expectFound bool) {
			t.Helper()
			ok, found := containsElement(list, elem)
			if ok != expectOK {
				t.Errorf("containsElement(%v, %v): expected ok=%v, got %v", list, elem, expectOK, ok)
			}
			if found != expectFound {
				t.Errorf("containsElement(%v, %v): expected found=%v, got %v", list, elem, expectFound, found)
			}
		}

		checkContains("Hello World", "World", true, true)
		checkContains(list1, "Foo", true, true)
		checkContains(list1, "Bar", true, true)
		checkContains(list2, 1, true, true)
		checkContains(list2, 2, true, true)
		checkContains(list1, "Foo!", true, false)
		checkContains(list2, 3, true, false)
		checkContains(list2, "1", true, false)
		checkContains(simpleMap, "Foo", true, true)
		checkContains(simpleMap, "Bar", true, false)
		checkContains(1433, "1", false, false)
	}
}

func testGetLen() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for v := range collectionImplGetLenFalseCases() {
			l, ok := getLen(v)
			if ok {
				t.Errorf("expected getLen to fail for %#v", v)
			}
			if l != 0 {
				t.Errorf("expected getLen to return 0 for %#v, got %d", v, l)
			}
		}

		for c := range collectionImplGetLenTrueCases() {
			l, ok := getLen(c.v)
			if !ok {
				t.Errorf("expected getLen to succeed for %#v", c.v)
			}
			if c.l != l {
				t.Errorf("expected length %d for %#v, got %d", c.l, c.v, l)
			}
		}
	}
}

type collectionImplGetLenCase = any

func collectionImplGetLenFalseCases() iter.Seq[collectionImplGetLenCase] {
	return slices.Values([]collectionImplGetLenCase{
		nil,
		0,
		true,
		false,
		'A',
		struct{}{},
	})
}

type collectionImplGetLenTrueCase struct {
	v any
	l int
}

func collectionImplGetLenTrueCases() iter.Seq[collectionImplGetLenTrueCase] {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3

	return slices.Values([]collectionImplGetLenTrueCase{
		{[]int{1, 2, 3}, 3},
		{[...]int{1, 2, 3}, 3},
		{"ABC", 3},
		{map[int]int{1: 2, 2: 4, 3: 6}, 3},
		{ch, 3},

		{[]int{}, 0},
		{map[int]int{}, 0},
		{make(chan int), 0},

		{[]int(nil), 0},
		{map[int]int(nil), 0},
		{(chan int)(nil), 0},
	})
}
