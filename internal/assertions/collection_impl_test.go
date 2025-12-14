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

		ok, found := containsElement("Hello World", "World")
		True(t, ok)
		True(t, found)

		ok, found = containsElement(list1, "Foo")
		True(t, ok)
		True(t, found)

		ok, found = containsElement(list1, "Bar")
		True(t, ok)
		True(t, found)

		ok, found = containsElement(list2, 1)
		True(t, ok)
		True(t, found)

		ok, found = containsElement(list2, 2)
		True(t, ok)
		True(t, found)

		ok, found = containsElement(list1, "Foo!")
		True(t, ok)
		False(t, found)

		ok, found = containsElement(list2, 3)
		True(t, ok)
		False(t, found)

		ok, found = containsElement(list2, "1")
		True(t, ok)
		False(t, found)

		ok, found = containsElement(simpleMap, "Foo")
		True(t, ok)
		True(t, found)

		ok, found = containsElement(simpleMap, "Bar")
		True(t, ok)
		False(t, found)

		ok, found = containsElement(1433, "1")
		False(t, ok)
		False(t, found)
	}
}

func testGetLen() func(*testing.T) {
	return func(t *testing.T) {
		t.Parallel()

		for v := range collectionImplGetLenFalseCases() {
			l, ok := getLen(v)
			False(t, ok, "Expected getLen fail to get length of %#v", v)
			Equal(t, 0, l, "getLen should return 0 for %#v", v)
		}

		for c := range collectionImplGetLenTrueCases() {
			l, ok := getLen(c.v)
			True(t, ok, "Expected getLen success to get length of %#v", c.v)
			Equal(t, c.l, l)
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
