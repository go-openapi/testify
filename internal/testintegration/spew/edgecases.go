// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

import (
	"fmt"
	"time"

	"pgregory.net/rapid"
)

// edgeCaseGenerator builds known edge cases which should all be already well-handled
// by [spew.Sdump].
func edgeCaseGenerator(opts ...Option) *rapid.Generator[any] {
	o := optionsWithDefaults(opts)

	allEdgeCases := []*rapid.Generator[any]{
		rapid.Custom(genStructWithUnexportedFields),
		rapid.Custom(genNilInterface),
		rapid.Custom(genCircularReference),
		rapid.Custom(genMapWithInterfaceKeys),
		rapid.Custom(genComplexPointerChain),
		rapid.Custom(genInterfaceSlice),
		rapid.Custom(genTimeValues),
		rapid.Custom(genDeeplyNested),
		rapid.Custom(genNestedInterfaces),
		rapid.Custom(genChanAndFuncValues),
		rapid.Custom(genUncomparableMapValues),
		rapid.Custom(genMixedPointerSemantics),
		rapid.Custom(genAnonymousStructNesting),
		rapid.Custom(genInterfaceArrays),
		rapid.Custom(genMultiLevelPointerIndirection),
		rapid.Custom(genPointerToInterface),
		rapid.Custom(genCircularInterfaceRef),
	}
	if !o.skipCircularMap {
		allEdgeCases = append(allEdgeCases,
			rapid.Custom(genCircularMapRef),
		)
	}

	return rapid.OneOf(
		allEdgeCases...,
	)
}

// genStructWithUnexportedFields generates structs with unexported fields
// This addresses issue #1828 where spew panicked on unexported fields.
func genStructWithUnexportedFields(t *rapid.T) any {
	type TestStruct struct {
		PublicField   string
		privateField  int // Unexported - caused #1828
		AnotherPublic []any
	}

	return TestStruct{
		PublicField:  rapid.String().Draw(t, "public"),
		privateField: rapid.Int().Draw(t, "private"),
		AnotherPublic: []any{
			rapid.Int().Draw(t, "nested-int"),
			rapid.String().Draw(t, "nested-string"),
		},
	}
}

// genCircularReference generates structures with circular references.
func genCircularReference(t *rapid.T) any {
	type Node struct {
		Value string
		Next  *Node
	}

	n1 := &Node{Value: rapid.String().Draw(t, "node1")}
	n2 := &Node{Value: rapid.String().Draw(t, "node2"), Next: n1}
	n3 := &Node{Value: rapid.String().Draw(t, "node3"), Next: n2}

	// Create cycle
	if rapid.Bool().Draw(t, "create-cycle") {
		n1.Next = n3
	}

	return n1
}

// genNilInterface generates nil values in non-nil interfaces.
func genNilInterface(t *rapid.T) any {
	generators := []func() any{
		func() any { var i any = (*int)(nil); return i },
		func() any { var i any = (*string)(nil); return i },
		func() any { var i any = (*[]int)(nil); return i },
		func() any { var i error = (*testError)(nil); return i },
	}

	idx := rapid.IntRange(0, len(generators)-1).Draw(t, "nil-interface-type")
	return generators[idx]()
}

type testError struct{}

func (testError) Error() string { return "test error" }

// genDeeplyNested generates deeply nested structures
// Avoids pointer-to-interface which can cause spew to hang.
func genDeeplyNested(t *rapid.T) any {
	const (
		minDepth  = 5
		maxDepth  = 20
		testCases = 3

		testCaseSliceAny  = 0
		testCaseMapString = 1
		testCaseStruct    = 2
		testCaseArrayAny  = 3
	)

	depth := rapid.IntRange(minDepth, maxDepth).Draw(t, "depth")

	var result any = rapid.String().Draw(t, "leaf")

	for i := range depth {
		switch rapid.IntRange(0, testCases).Draw(t, fmt.Sprintf("nesting-type-%d", i)) {
		case testCaseSliceAny:
			result = []any{result}
		case testCaseMapString:
			result = map[string]any{"nested": result}
		case testCaseStruct:
			// Wrap in struct instead of pointer-to-interface
			result = struct{ Value any }{Value: result}
		case testCaseArrayAny:
			// Array wrapping
			result = [1]any{result}
		}
	}

	return result
}

// genMapWithInterfaceKeys generates maps with interface{} keys.
func genMapWithInterfaceKeys(t *rapid.T) any {
	const maxKeys = 10
	m := make(map[any]string)

	// Add various key types
	numKeys := rapid.IntRange(1, maxKeys).Draw(t, "num-keys")
	for i := range numKeys {
		key := rapid.OneOf(
			rapid.Just[any](rapid.Int().Draw(t, fmt.Sprintf("key-int-%d", i))),
			rapid.Just[any](rapid.String().Draw(t, fmt.Sprintf("key-string-%d", i))),
			rapid.Just[any](struct{ x int }{rapid.Int().Draw(t, fmt.Sprintf("key-struct-%d", i))}),
		).Draw(t, fmt.Sprintf("key-%d", i))

		m[key] = rapid.String().Draw(t, fmt.Sprintf("value-%d", i))
	}

	return m
}

// genComplexPointerChain generates complex pointer chains including nil pointers.
func genComplexPointerChain(t *rapid.T) any {
	type PointerChain struct {
		Value  *int
		Next   *PointerChain
		Values []*string
	}

	var makeChain func(depth int) *PointerChain
	makeChain = func(depth int) *PointerChain {
		if depth <= 0 || rapid.Bool().Draw(t, fmt.Sprintf("terminate-%d", depth)) {
			return nil
		}

		chain := &PointerChain{}

		// Sometimes nil, sometimes has value
		if rapid.Bool().Draw(t, fmt.Sprintf("has-value-%d", depth)) {
			v := rapid.Int().Draw(t, fmt.Sprintf("value-%d", depth))
			chain.Value = &v
		}

		// Sometimes recurse
		if rapid.Bool().Draw(t, fmt.Sprintf("has-next-%d", depth)) {
			chain.Next = makeChain(depth - 1)
		}

		// Add some nil and non-nil string pointers
		const maxStrings = 5
		numStrings := rapid.IntRange(0, maxStrings).Draw(t, fmt.Sprintf("num-strings-%d", depth))
		for i := range numStrings {
			if rapid.Bool().Draw(t, fmt.Sprintf("string-nil-%d-%d", depth, i)) {
				chain.Values = append(chain.Values, nil)
			} else {
				s := rapid.String().Draw(t, fmt.Sprintf("string-%d-%d", depth, i))
				chain.Values = append(chain.Values, &s)
			}
		}

		return chain
	}

	const (
		minChainDepth = 3
		maxChainDepth = 10
	)

	return makeChain(rapid.IntRange(minChainDepth, maxChainDepth).Draw(t, "chain-depth"))
}

// genTimeValues generates various time.Time values
// This addresses issue #1829 - time.Time rendering in diffs.
func genTimeValues(t *rapid.T) any {
	type TimeContainer struct {
		Time     time.Time
		TimePtr  *time.Time
		Times    []time.Time
		TimeMap  map[string]time.Time
		ZeroTime time.Time
	}

	const maxDuration = 1_000_000_000_000
	now := time.Now()
	past := now.Add(-time.Duration(rapid.Int64Range(0, maxDuration).Draw(t, "past-duration")))

	container := TimeContainer{
		Time:    past,
		Times:   []time.Time{now, past, {}},
		TimeMap: map[string]time.Time{"now": now, "past": past, "zero": {}},
	}

	if rapid.Bool().Draw(t, "has-time-ptr") {
		container.TimePtr = &now
	}

	return container
}

// genInterfaceSlice generates slices of interface{} with mixed types.
func genInterfaceSlice(t *rapid.T) any {
	const maxLength = 20
	length := rapid.IntRange(0, maxLength).Draw(t, "slice-length")
	slice := make([]any, length)

	for i := range length {
		slice[i] = rapid.OneOf(
			rapid.Just[any](rapid.Int().Draw(t, fmt.Sprintf("elem-int-%d", i))),
			rapid.Just[any](rapid.String().Draw(t, fmt.Sprintf("elem-string-%d", i))),
			rapid.Just[any](rapid.Bool().Draw(t, fmt.Sprintf("elem-bool-%d", i))),
			rapid.Just[any](nil),
			rapid.Just[any](struct{ x, y int }{
				x: rapid.Int().Draw(t, fmt.Sprintf("elem-struct-x-%d", i)),
				y: rapid.Int().Draw(t, fmt.Sprintf("elem-struct-y-%d", i)),
			}),
		).Draw(t, fmt.Sprintf("elem-%d", i))
	}

	return slice
}

// genPointerToInterface generates pointer-to-interface chains
// This is the problematic case that can cause spew to hang.
func genPointerToInterface(t *rapid.T) any {
	const maxDepth = 3
	depth := rapid.IntRange(1, maxDepth).Draw(t, "ptr-depth") // Max 3 levels

	var result any = rapid.String().Draw(t, "leaf")
	for range depth {
		clone := result
		result = &clone
	}

	return result
}

// genNestedInterfaces generates interface values containing other interfaces.
func genNestedInterfaces(t *rapid.T) any {
	const (
		maxDepth    = 5
		maxIntValue = 100
	)
	depth := rapid.IntRange(1, maxDepth).Draw(t, "interface-depth")

	var result any = rapid.String().Draw(t, "leaf")
	for i := range depth {
		// Wrap in struct with interface field
		type Wrapper struct {
			Value any
			Extra any
		}
		result = Wrapper{
			Value: result,
			Extra: rapid.IntRange(0, maxIntValue).Draw(t, fmt.Sprintf("extra-%d", i)),
		}
	}

	return result
}

// genChanAndFuncValues generates channels and functions as interface values.
func genChanAndFuncValues(t *rapid.T) any {
	const (
		testCases      = 5
		maxBuffers     = 10
		funcMultiplier = 2

		testCaseUnbufferedChan = 0
		testCaseBufferedChan   = 1
		testCaseNilChan        = 2
		testCaseFunction       = 3
		testCaseNilFunction    = 4
		testCaseStruct         = 5
	)
	choice := rapid.IntRange(0, testCases).Draw(t, "chan-func-choice")

	switch choice {
	case testCaseUnbufferedChan:
		// Unbuffered channel
		ch := make(chan int)
		return ch
	case testCaseBufferedChan:
		// Buffered channel
		ch := make(chan string, rapid.IntRange(1, maxBuffers).Draw(t, "buffer-size"))
		return ch
	case testCaseNilChan:
		// Nil channel
		var ch chan int
		return ch
	case testCaseFunction:
		// Function value
		fn := func(x int) int { return x * funcMultiplier }
		return fn
	case testCaseNilFunction:
		// Nil function
		var fn func()
		return fn
	case testCaseStruct:
		// Struct containing channels and funcs
		type ChanFuncStruct struct {
			Ch   chan int
			Fn   func(string) bool
			Data []any
		}
		return ChanFuncStruct{
			Ch:   make(chan int, 1),
			Fn:   func(s string) bool { return len(s) > 0 },
			Data: []any{make(chan bool), func() {}},
		}
	}
	return nil
}

// genUncomparableMapValues generates maps with uncomparable types as values
// (not as keys, which would panic).
func genUncomparableMapValues(t *rapid.T) any {
	const (
		maxEntries = 5
		testCases  = 3

		testCaseSliceValue      = 0
		testCaseMapValue        = 1
		testCaseFuncValue       = 2
		testCaseSliceSliceValue = 3
	)
	m := make(map[string]any)

	numEntries := rapid.IntRange(1, maxEntries).Draw(t, "num-entries")
	for i := range numEntries {
		key := fmt.Sprintf("key-%d", i)

		// Use uncomparable types as values
		choice := rapid.IntRange(0, testCases).Draw(t, fmt.Sprintf("uncomparable-%d", i))
		switch choice {
		case testCaseSliceValue:
			m[key] = []int{1, 2, 3}
		case testCaseMapValue:
			m[key] = map[string]int{"nested": i}
		case testCaseFuncValue:
			m[key] = func() int { return i }
		case testCaseSliceSliceValue:
			// Slice of slices
			m[key] = [][]string{{"a", "b"}, {"c"}}
		}
	}

	return m
}

// genMixedPointerSemantics generates structs with mixed value/pointer fields.
func genMixedPointerSemantics(t *rapid.T) any {
	type MixedStruct struct {
		// Value types
		IntVal    int
		StringVal string
		SliceVal  []int

		// Pointer types
		IntPtr    *int
		StringPtr *string
		SlicePtr  *[]int

		// Nested mixed
		Nested struct {
			Value any
			Ptr   *any
		}
	}

	i := rapid.Int().Draw(t, "int-val")
	s := rapid.String().Draw(t, "string-val")
	slice := []int{1, 2, 3}

	var nestedVal any = rapid.Int().Draw(t, "nested-val")

	ms := MixedStruct{
		IntVal:    i,
		StringVal: s,
		SliceVal:  []int{rapid.Int().Draw(t, "slice-elem")},
		IntPtr:    &i,
		StringPtr: &s,
		SlicePtr:  &slice,
	}

	ms.Nested.Value = nestedVal
	ms.Nested.Ptr = &nestedVal

	return ms
}

// genAnonymousStructNesting generates deeply nested anonymous structs.
func genAnonymousStructNesting(t *rapid.T) any {
	const (
		minDepth = 2
		maxDepth = 7
	)
	depth := rapid.IntRange(minDepth, maxDepth).Draw(t, "anon-depth")

	// Start with innermost value
	var result any = struct {
		Value string
	}{
		Value: rapid.String().Draw(t, "innermost"),
	}

	// Wrap in anonymous structs
	for i := range depth {
		result = struct {
			Level int
			Data  any
			Extra struct {
				Nested any
			}
		}{
			Level: i,
			Data:  result,
			Extra: struct{ Nested any }{
				Nested: rapid.Int().Draw(t, fmt.Sprintf("extra-%d", i)),
			},
		}
	}

	return result
}

// genInterfaceArrays generates arrays (not slices) with interface elements.
func genInterfaceArrays(t *rapid.T) any {
	const (
		testCases = 3

		testCaseArray           = 0
		testCaseArrayPointers   = 1
		testCaseArrayNested     = 2
		testCaseArrayIfaceArray = 3
	)
	choice := rapid.IntRange(0, testCases).Draw(t, "array-choice")

	switch choice {
	case testCaseArray:
		// Fixed size array of interfaces
		return [3]any{
			rapid.Int().Draw(t, "elem-0"),
			rapid.String().Draw(t, "elem-1"),
			nil,
		}
	case testCaseArrayPointers:
		// Array of pointers to interfaces
		var i1, i2, i3 any
		i1 = rapid.Int().Draw(t, "ptr-elem-0")
		i2 = rapid.String().Draw(t, "ptr-elem-1")
		i3 = rapid.Bool().Draw(t, "ptr-elem-2")
		return [3]*any{&i1, &i2, &i3}
	case testCaseArrayNested:
		// Nested arrays
		return [2][2]any{
			{rapid.Int().Draw(t, "00"), rapid.String().Draw(t, "01")},
			{rapid.Bool().Draw(t, "10"), nil},
		}
	case testCaseArrayIfaceArray:
		// Array containing array
		inner := [2]int{rapid.Int().Draw(t, "inner-0"), rapid.Int().Draw(t, "inner-1")}
		return [1]any{inner}
	}

	return nil
}

// genMultiLevelPointerIndirection generates values with multiple pointer levels.
func genMultiLevelPointerIndirection(t *rapid.T) any {
	const (
		maxLevels = 5

		simplePointer     = 0
		oneIndirection    = 1
		doubleIndirection = 2
	)
	levels := rapid.IntRange(1, maxLevels).Draw(t, "ptr-levels")

	baseVal := rapid.Int().Draw(t, "base")

	// Build pointer chain: int -> *int -> **int -> ***int ...
	var current any = baseVal
	for i := range levels {
		// Create new pointer level
		switch i {
		case simplePointer:
			p := baseVal
			current = &p
		case oneIndirection:
			p, ok := current.(*int)
			if !ok {
				t.Fatalf("internal error: expected a *int")
			}
			current = &p
		case doubleIndirection:
			p, ok := current.(**int)
			if !ok {
				t.Fatalf("internal error: expected a **int")
			}
			current = &p
		default:
			// For deeper levels, use interface
			temp := current
			current = &temp
		}
	}

	return current
}

// genCircularInterfaceRef generates circular references through interfaces
// This tests the KNOWN BUG: circular references via interface{} cause hangs.
func genCircularInterfaceRef(t *rapid.T) any {
	const (
		testCases = 4

		testCaseSelfPointer    = 0
		testCaseCircularStruct = 1
		testCaseCircularSlice  = 2
		testCaseCircularChain  = 3
		testCaseDoublePointer  = 4
	)
	choice := rapid.IntRange(0, testCases).Draw(t, "circular-type")

	switch choice {
	case testCaseSelfPointer:
		// Simple self-referential pointer through interface
		// This recreates the original genPointerToInterface bug
		var self any = rapid.String().Draw(t, "base")
		self = &self // self now contains a pointer to itself!
		return self

	case testCaseCircularStruct:
		// Struct with circular interface field
		type CircularStruct struct {
			Name  string
			Cycle any // Will point back to the struct
		}
		cs := &CircularStruct{
			Name: rapid.String().Draw(t, "struct-name"),
		}
		cs.Cycle = cs // Circular reference through interface
		return cs

	case testCaseCircularSlice:
		// Slice with circular element (slice contains itself)
		type CircularSlice struct {
			Items []any
		}
		cs := &CircularSlice{
			Items: []any{
				rapid.Int().Draw(t, "item-0"),
				rapid.String().Draw(t, "item-1"),
			},
		}
		cs.Items = append(cs.Items, cs) // Add self to slice
		return cs

	case testCaseCircularChain:
		// Multi-level circular chain: A -> B -> C -> A
		type Node struct {
			Name string
			Next any
		}
		a := &Node{Name: "A"}
		b := &Node{Name: "B"}
		c := &Node{Name: "C"}

		a.Next = b
		b.Next = c
		c.Next = a // Complete the cycle

		return a

	case testCaseDoublePointer:
		// Double pointer with circular reference
		var base any = rapid.String().Draw(t, "base-val")
		ptr1 := &base
		base = &ptr1 // Now base points to ptr1, which points to base!
		return ptr1
	}

	return nil
}

func genCircularMapRef(t *rapid.T) any {
	// Map with circular value (map contains itself)
	m := map[string]any{
		"key1": rapid.Int().Draw(t, "val-1"),
		"key2": rapid.String().Draw(t, "val-2"),
	}
	m["circular"] = m // Map contains itself as a value

	return m
}
