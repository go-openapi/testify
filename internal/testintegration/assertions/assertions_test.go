// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	"pgregory.net/rapid"

	"github.com/go-openapi/testify/v2/internal/assertions"
)

func TestMain(m *testing.M) {
	os.Args = append(os.Args, "-rapid.checks", strconv.Itoa(testLoad()))
	flag.Parse()

	os.Exit(m.Run())
}

func testLoad() int {
	isCI := os.Getenv("CI") != ""

	if isCI {
		return 100
	}

	return 100_000
}

// silentT is a T that silently absorbs assertion failures.
type silentT struct{}

func (silentT) Errorf(string, ...any) {}
func (silentT) Helper()               {}

var sink bool

// TestNilSafetyUnary verifies that unary assertion functions (taking a single
// value of type any) never panic, even with nil, nil-pointer, or arbitrary
// inputs.
func TestNilSafetyUnary(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		value := genAny().Draw(rt, "value")
		mock := silentT{}

		// These functions must never panic, regardless of the value passed.
		sink = assertions.Nil(mock, value)
		sink = assertions.NotNil(mock, value)
		sink = assertions.Empty(mock, value)
		sink = assertions.NotEmpty(mock, value)
		sink = assertions.Zero(mock, value)
		sink = assertions.NotZero(mock, value)
		sink = assertions.Len(mock, value, 0)
		sink = assertions.Len(mock, value, 1)
	})
}

// TestNilSafetyBinary verifies that binary assertion functions (comparing two
// values of type any) never panic with arbitrary inputs.
func TestNilSafetyBinary(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		a := genAny().Draw(rt, "a")
		b := genAny().Draw(rt, "b")
		mock := silentT{}

		sink = assertions.Equal(mock, a, b)
		sink = assertions.NotEqual(mock, a, b)
		sink = assertions.EqualValues(mock, a, b)
		sink = assertions.NotEqualValues(mock, a, b)
		sink = assertions.Exactly(mock, a, b)
		sink = assertions.Same(mock, a, b)
		sink = assertions.NotSame(mock, a, b)
		_ = assertions.ObjectsAreEqual(a, b)
		_ = assertions.ObjectsAreEqualValues(a, b)
	})
}

// TestNilSafetyCollections verifies that collection-oriented assertion
// functions never panic with arbitrary inputs.
func TestNilSafetyCollections(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		collection := genAny().Draw(rt, "collection")
		element := genAny().Draw(rt, "element")
		mock := silentT{}

		sink = assertions.Contains(mock, collection, element)
		sink = assertions.NotContains(mock, collection, element)
		sink = assertions.Subset(mock, collection, element)
		sink = assertions.NotSubset(mock, collection, element)
		sink = assertions.ElementsMatch(mock, collection, element)
		sink = assertions.NotElementsMatch(mock, collection, element)
		sink = assertions.IsIncreasing(mock, collection)
		sink = assertions.IsDecreasing(mock, collection)
		sink = assertions.IsNonIncreasing(mock, collection)
		sink = assertions.IsNonDecreasing(mock, collection)
	})
}

// TestNilSafetyComparison verifies that comparison/numeric assertion
// functions never panic with arbitrary inputs.
func TestNilSafetyComparison(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		a := genAny().Draw(rt, "a")
		b := genAny().Draw(rt, "b")
		mock := silentT{}

		sink = assertions.Greater(mock, a, b)
		sink = assertions.GreaterOrEqual(mock, a, b)
		sink = assertions.Less(mock, a, b)
		sink = assertions.LessOrEqual(mock, a, b)
		sink = assertions.Positive(mock, a)
		sink = assertions.Negative(mock, a)
		sink = assertions.InDelta(mock, a, b, 1.0)
		sink = assertions.InEpsilon(mock, a, b, 0.01)
	})
}

// TestNilSafetyType verifies that type-checking assertion functions never
// panic with arbitrary inputs.
func TestNilSafetyType(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		a := genAny().Draw(rt, "a")
		b := genAny().Draw(rt, "b")
		mock := silentT{}

		sink = assertions.IsType(mock, a, b)
		sink = assertions.IsNotType(mock, a, b)
		sink = assertions.Kind(mock, reflect.Int, a)
		sink = assertions.NotKind(mock, reflect.Int, a)
		sink = assertions.Kind(mock, reflect.Pointer, a)
		sink = assertions.NotKind(mock, reflect.Pointer, a)
	})
}

// TestNilSafetyExportedValues verifies that EqualExportedValues never panics.
func TestNilSafetyExportedValues(t *testing.T) {
	t.Parallel()

	rapid.Check(t, func(rt *rapid.T) {
		a := genAny().Draw(rt, "a")
		b := genAny().Draw(rt, "b")
		mock := silentT{}

		sink = assertions.EqualExportedValues(mock, a, b)
	})
}

// genAny generates random values of many different types, with emphasis on
// nil and nil-pointer edge cases that could trigger nil-pointer dereferences.
func genAny() *rapid.Generator[any] {
	return rapid.Custom(func(t *rapid.T) any {
		kind := rapid.IntRange(0, 9).Draw(t, "kind")
		switch kind {
		case 0:
			return genNilValue(t)
		case 1:
			return rapid.Int().Draw(t, "int")
		case 2:
			return rapid.Float64().Draw(t, "float64")
		case 3:
			return rapid.String().Draw(t, "string")
		case 4:
			return rapid.Bool().Draw(t, "bool")
		case 5:
			return genSlice(t)
		case 6:
			return genMap(t)
		case 7:
			return genStruct(t)
		case 8:
			return genPointer(t)
		default:
			return rapid.Byte().Draw(t, "byte")
		}
	})
}

// genNilValue produces various nil representations.
func genNilValue(t *rapid.T) any {
	variant := rapid.IntRange(0, 11).Draw(t, "nil-variant")
	switch variant {
	case 0:
		return nil
	case 1:
		return (*int)(nil)
	case 2:
		return (*string)(nil)
	case 3:
		return (*[]int)(nil)
	case 4:
		return (*map[string]int)(nil)
	case 5:
		return ([]int)(nil)
	case 6:
		return (map[string]int)(nil)
	case 7:
		return (chan int)(nil)
	case 8:
		return (func())(nil)
	case 9:
		return (*struct{})(nil)
	case 10:
		return (error)(nil)
	default:
		return (fmt.Stringer)(nil)
	}
}

func genSlice(t *rapid.T) any {
	variant := rapid.IntRange(0, 4).Draw(t, "slice-variant")
	switch variant {
	case 0:
		return rapid.SliceOfN(rapid.Int(), 0, 5).Draw(t, "int-slice")
	case 1:
		return rapid.SliceOfN(rapid.String(), 0, 5).Draw(t, "string-slice")
	case 2:
		return rapid.SliceOfN(rapid.Float64(), 0, 5).Draw(t, "float-slice")
	case 3:
		return []any{nil, rapid.Int().Draw(t, "elem"), "hello", nil}
	default:
		return []any(nil)
	}
}

func genMap(t *rapid.T) any {
	variant := rapid.IntRange(0, 3).Draw(t, "map-variant")
	switch variant {
	case 0:
		return rapid.MapOfN(rapid.String(), rapid.Int(), 0, 5).Draw(t, "string-int-map")
	case 1:
		return rapid.MapOfN(rapid.String(), rapid.String(), 0, 5).Draw(t, "string-string-map")
	case 2:
		return map[string]any{"key": nil, "val": rapid.Int().Draw(t, "v")}
	default:
		return map[string]any(nil)
	}
}

type nested struct {
	Inner *int
	Name  string
}

type outer struct {
	Nested *nested
	Value  any
}

func genStruct(t *rapid.T) any {
	variant := rapid.IntRange(0, 2).Draw(t, "struct-variant")
	switch variant {
	case 0:
		return nested{
			Inner: nil,
			Name:  rapid.String().Draw(t, "name"),
		}
	case 1:
		v := rapid.Int().Draw(t, "v")
		return outer{
			Nested: &nested{Inner: &v, Name: "test"},
			Value:  &v,
		}
	default:
		return outer{
			Nested: nil,
			Value:  nil,
		}
	}
}

func genPointer(t *rapid.T) any {
	variant := rapid.IntRange(0, 5).Draw(t, "ptr-type")
	switch variant {
	case 0:
		v := rapid.Int().Draw(t, "int-val")
		return &v
	case 1:
		v := rapid.String().Draw(t, "string-val")
		return &v
	case 2:
		v := rapid.Float64().Draw(t, "float-val")
		return &v
	case 3:
		return (*int)(nil)
	case 4:
		return (*string)(nil)
	default:
		return (*struct{})(nil)
	}
}
