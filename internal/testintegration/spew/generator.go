// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

import (
	"context"
	"reflect"
	"sync"
	"sync/atomic"
	"time"
	"unicode"

	"pgregory.net/rapid"

	"github.com/go-openapi/testify/v2/internal/spew"
)

const (
	maxTestDuration = time.Second
	defaultMaxDepth = uint32(10)
	maxStructFields = 5
	maxFieldNameLen = 10
	maxStringLen    = 10
	maxChanBuf      = 3
	maxArrayLen     = 5
)

// NoPanicProp produces a check for the [rapid.T] test wrapper.
//
// It verifies that [spew.Sdump] does not panic or hang given random values generated
// by the provided [rapid.Generator].
//
// NOTE: [spew.Sdump] is a perfect endpoint to cover most of what the spew package is doing.
//
// This test doesn't check how well the values are rendered.
func NoPanicProp(ctx context.Context, g *rapid.Generator[any]) func(*rapid.T) {
	return func(rt *rapid.T) {
		value := g.Draw(rt, "arbitrary-value")
		timeoutCtx, cancel := context.WithTimeout(ctx, maxTestDuration)
		defer cancel()

		var w sync.WaitGroup
		done := make(chan struct{})

		w.Add(1)
		go func() {
			defer w.Done()
			select {
			case <-done:
				cancel()

				return
			case <-timeoutCtx.Done():
				rt.Fatalf("Sdump timed out:\nWith value type: %T\nValue: %#v", value, value)

				return
			}
		}()

		go func() { // this go routine may leak if timeout kicks
			// Sdump should never panic
			defer func() {
				if r := recover(); r != nil {
					rt.Errorf("Sdump panicked: %v\nWith value type: %T\nValue: %#v", r, value, value)
					close(done)

					return
				}
			}()

			value = spew.Sdump(value)
			// fmt.Printf("%v", value)

			close(done)
		}()

		w.Wait()
	}
}

// Generator builds a [rapid.Generator] capable of producing any kind of go value.
//
// It is biased toward exploring edge cases such as cyclical pointer references, or nil values.
// It may use uncommon but legit constructs like *[]any, *map[],...
//
// Known edge cases reproducing historical issues are added to the generator.
//
// Limitations:
//
// * does not generate generic types
// * does not generate type declarations: all types are anonymous
// * does not generate structs with unexported fields
// * does not generate embedded fields in structs with methods
//
// These limitations are partially mitigated by the edgecase generator,
// which is not dependent on the [reflect] package.
func Generator(opts ...Option) *rapid.Generator[any] {
	g := newTypeGenerator()

	return rapid.OneOf(
		g.Generator(opts...),
		edgeCaseGenerator(opts...),
	)
}

type typeGenerator struct {
	mx       sync.Mutex
	pointers map[any]struct{}
	maxDepth uint32
}

func newTypeGenerator() *typeGenerator {
	return &typeGenerator{
		pointers: make(map[any]struct{}),
		maxDepth: defaultMaxDepth,
	}
}

func (g *typeGenerator) Generator(_ ...Option) *rapid.Generator[any] {
	return g.genAnything(0)
}

func (g *typeGenerator) genAnything(depth uint32) *rapid.Generator[any] {
	return rapid.Deferred(func() *rapid.Generator[any] { // recursive definition. Recursion stops on max depth
		return rapid.OneOf(
			genPrimitiveValue(),          // int, bool, string, etc
			g.genContainerValue(depth+1), // map, slice, struct, array
			g.genOtherValue(depth+1),     // chan, func, and other peculiar types present in the standard library (sync.Mutex, ...)
			g.genPointer(depth+1),        // pointer to anything, including nil and cyclical references
			genInterfaceValue(depth+1),   // any, some interface
		)
	})
}

// genPointer may either generate a pointer to a new structure (may be nil),
// or a cyclical reference to an already created pointer.
func (g *typeGenerator) genPointer(depth uint32) *rapid.Generator[any] {
	return rapid.OneOf(
		g.genNewPointer(depth),
		g.genExistingPointer(),
	)
}

// genNewPointer produces a new pointer.
//
// We don't use [rapid.Ptr] to avoid always having *any.
func (g *typeGenerator) genNewPointer(depth uint32) *rapid.Generator[any] {
	return rapid.Custom(func(t *rapid.T) any {
		value := g.genAnything(depth).Draw(t, "new-value")

		val := reflect.ValueOf(value)
		if val.Kind() == reflect.Interface {
			val = val.Elem()
		}
		if val.Kind() == reflect.Pointer && val.Elem().Kind() == reflect.Interface {
			val = val.Elem().Elem()
		}

		flipCoin := rapid.Bool().Draw(t, "flip-coin")
		if flipCoin && val.Kind() == reflect.Interface {
			val = val.Elem()
		}

		typ := reflect.TypeOf(value)
		if typ == nil {
			var iface any

			return &iface
		}

		ptrVal := val
		if val.Kind() != reflect.Pointer {
			if val.CanAddr() {
				ptrVal = val.Addr()
			} else {
				clone := reflect.New(reflect.TypeOf(value))
				// check the value is not nil
				if val.IsValid() {
					clone.Elem().Set(val)
				}
				ptrVal = clone
			}
		}

		ptr := ptrVal.Interface()
		if !ptrVal.IsNil() {
			g.mx.Lock()
			g.pointers[ptr] = struct{}{}
			g.mx.Unlock()
		}

		return ptr
	})
}

func (g *typeGenerator) genExistingPointer() *rapid.Generator[any] {
	return rapid.Custom(func(t *rapid.T) any {
		g.mx.Lock()
		l := len(g.pointers)
		g.mx.Unlock()

		if l == 0 {
			return g.genNewPointer(0).Draw(t, "new-pointer")
		}

		// may draw a cyclical reference
		// Random iterations over the map is deemed sufficient randomization (we MUST call rapid random generator somehow)
		const minIter = 4
		maxIter := rapid.IntRange(minIter, max(minIter, len(g.pointers))).Draw(t, "ptr-iterations")
		var k any
		var j int

		for k = range g.pointers {
			if j > maxIter {
				break
			}
			j++
		}

		return k
	})
}

// container values are structs, slices, arrays and maps.
func (g *typeGenerator) genContainerValue(depth uint32) *rapid.Generator[any] {
	return rapid.Custom(func(t *rapid.T) any {
		if depth > g.maxDepth {
			return genPrimitiveValue().Draw(t, "final")
		}

		return rapid.OneOf(
			g.genStructValue(depth+1),
			g.genArrayValue(depth+1),
			g.genSliceValue(depth+1),
			g.genMapValue(depth+1),
		).Draw(t, "container")
	})
}

// genSliceValue generates a slice of any type.
// We don't use [rapid.SliceOf] to avoid always having a []any.
//
// Since slices are not comparable, we don't need a comparable version of this.
//
// # Limitation
//
// At this moment, the slice is of random size (can be nil or empty) but populated
// with a single random value, unlike [rapid.SliceOf].
func (g *typeGenerator) genSliceValue(depth uint32) *rapid.Generator[any] {
	return rapid.Custom(func(t *rapid.T) any {
		value := g.genAnything(depth).Draw(t, "value")
		val := reflect.ValueOf(value)
		flipCoin := rapid.Bool().Draw(t, "flip-coin")
		if flipCoin && val.Kind() == reflect.Interface {
			val = val.Elem()
		}
		if flipCoin && val.Kind() == reflect.Pointer && val.Elem().Kind() == reflect.Interface {
			val = val.Elem()
		}

		typ := reflect.TypeOf(value)
		if typ == nil {
			return ([]any)(nil)
		}

		size := rapid.IntRange(0, maxArrayLen).Draw(t, "slice-len")
		if size == 0 {
			// may chose to return a nil slice
			flipCoin := rapid.Bool().Draw(t, "flip-coin")
			if flipCoin {
				sliceType := reflect.SliceOf(val.Type())
				sliceValue := reflect.New(sliceType).Elem()

				return sliceValue.Interface()
			}
		}

		sliceType := reflect.SliceOf(val.Type())
		sliceValue := reflect.MakeSlice(sliceType, 0, size)
		for range size {
			sliceValue = reflect.Append(sliceValue, val)
		}

		return sliceValue.Interface()
	})
}

// genMapValue generates a map of any type with a key of any comparable type.
// We don't use [rapid.MapOf] to avoid always having a map[any]any.
//
// Since maps are not comparable, we don't need a comparable version of this.
//
// Maps can be nil, empty or contain one element.
//
// # Limitation
//
// At this moment, maps contain only one single random value, unlike [rapid.MapOf].
func (g *typeGenerator) genMapValue(depth uint32) *rapid.Generator[any] {
	return rapid.Custom(func(t *rapid.T) any {
		key := g.genComparableValue(depth).Draw(t, "key")

		keyVal := reflect.ValueOf(key)
		flipCoin := rapid.Bool().Draw(t, "flip-coin")
		if flipCoin && (keyVal.Kind() == reflect.Interface) {
			if keyVal.Elem().Comparable() {
				// may use interface wrapper, pointer or base type
				keyVal = keyVal.Elem()
			}
		}

		if flipCoin && (keyVal.Kind() == reflect.Pointer && keyVal.Elem().Kind() == reflect.Interface) {
			if keyVal.Elem().Comparable() {
				// may use interface wrapper, pointer or base type
				keyVal = keyVal.Elem()
			}
		}

		for !keyVal.Comparable() {
			if keyVal.Kind() == reflect.Interface || keyVal.Kind() == reflect.Pointer {
				// may use interface wrapper, pointer or base type
				keyVal = keyVal.Elem()
			}
		}

		value := g.genAnything(depth).Draw(t, "value")
		valueVal := reflect.ValueOf(value)
		flipCoin = rapid.Bool().Draw(t, "flip-coin")
		if flipCoin && (valueVal.Kind() == reflect.Interface || valueVal.Kind() == reflect.Pointer) {
			// may use interface wrapper, pointer or base type
			valueVal = valueVal.Elem()
		}

		typ := reflect.TypeOf(key)
		if typ == nil {
			k := "key"
			keyVal = reflect.ValueOf(&k)
		}
		typ = reflect.TypeOf(value)
		if typ == nil {
			v := "value"
			valueVal = reflect.ValueOf(&v)
		}

		mapType := reflect.MapOf(keyVal.Type(), valueVal.Type())
		flipCoin = rapid.Bool().Draw(t, "flip-coin")
		if flipCoin {
			// return nil map
			mapObject := reflect.New(mapType).Elem()

			return mapObject.Interface()
		}

		mapObject := reflect.MakeMap(mapType)
		flipCoin = rapid.Bool().Draw(t, "flip-coin")
		if flipCoin {
			// populate the map
			mapObject.SetMapIndex(keyVal, valueVal)
		}

		return mapObject.Interface()
	})
}

func (g *typeGenerator) genComparableValue(depth uint32) *rapid.Generator[any] {
	return rapid.OneOf(
		genPrimitiveValue(),                // int, bool, string, etc
		g.genComparableArrayValue(depth+1), // array
		g.genComparableStructValue(depth+1),
		g.genPointer(depth+1), // pointer to anything
		g.genChanValue(depth+1),
		// NOTE: this excerpt from the go language spec is not implemented:
		// Interface types that are not type parameters are comparable.
		// Two interface values are equal if they have identical dynamic types and equal dynamic values or if both have value nil.
	)
}

// genArray produces arrays with all their elements copied from a single random value.
func (g *typeGenerator) genArrayValue(depth uint32) *rapid.Generator[any] {
	return g.genArrayFromGenerator(g.genAnything, depth+1)()
}

func (g *typeGenerator) genComparableArrayValue(depth uint32) *rapid.Generator[any] {
	return g.genArrayFromGenerator(g.genComparableValue, depth+1)()
}

func (g *typeGenerator) genArrayFromGenerator(fn func(uint32) *rapid.Generator[any], depth uint32) func() *rapid.Generator[any] {
	return func() *rapid.Generator[any] {
		return rapid.Custom(func(t *rapid.T) any {
			first := fn(depth+1).Draw(t, "elem")
			firstValue := reflect.ValueOf(first)
			flipCoin := rapid.Bool().Draw(t, "flip-coin")
			if flipCoin && firstValue.Kind() == reflect.Interface {
				// may use interface wrapper, pointer or base type
				if firstValue.Elem().Comparable() {
					firstValue = firstValue.Elem()
				}
			}

			if flipCoin && firstValue.Kind() == reflect.Pointer && firstValue.Elem().Kind() == reflect.Interface {
				firstValue = firstValue.Elem()
				if firstValue.Elem().Comparable() {
					firstValue = firstValue.Elem()
				}
			}

			typ := reflect.TypeOf(first)
			if typ == nil {
				var iface any = "elem"

				firstValue = reflect.ValueOf(&iface)
			}

			elem := firstValue.Type()
			size := rapid.IntRange(0, maxArrayLen).Draw(t, "array-len")
			arrayType := reflect.ArrayOf(size, elem)
			arrayValue := reflect.New(arrayType).Elem()
			for i := range size {
				arrayValue.Index(i).Set(firstValue)
			}

			return arrayValue.Interface()
		})
	}
}

// genStructValue generates a random struct object with a random number of exported fields of any type.
//
// # Limitations
//
// * from [reflect.StructOf]: can't generate random unexported fields. To work with unexported fields, use the edgecaseGenerator.
// * no struct tags (not useful for [spew.Dump]).
func (g *typeGenerator) genStructValue(depth uint32) *rapid.Generator[any] {
	return g.genStructValueFromGenerator(g.genAnything, depth+1)()
}

func (g *typeGenerator) genComparableStructValue(depth uint32) *rapid.Generator[any] {
	return g.genStructValueFromGenerator(g.genComparableValue, depth+1)()
}

func (g *typeGenerator) genStructValueFromGenerator(fn func(uint32) *rapid.Generator[any], depth uint32) func() *rapid.Generator[any] {
	return func() *rapid.Generator[any] {
		return rapid.Custom(func(t *rapid.T) any {
			numFields := rapid.IntRange(0, maxStructFields).Draw(t, "struct-fields") // may have empty struct{}
			fields := make([]reflect.StructField, 0, numFields)
			fieldValues := make([]reflect.Value, 0, numFields)
			names := rapid.SliceOfNDistinct(fieldNameGenerator(), numFields, numFields, func(str string) string { return str }).Draw(t, "field-names")

			for i := range numFields {
				value := fn(depth+1).Draw(t, "field-value")
				val := reflect.ValueOf(value)
				if val.Kind() == reflect.Interface {
					val = val.Elem()
				}
				if val.Kind() == reflect.Pointer && val.Elem().Kind() == reflect.Interface {
					val = val.Elem()
				}

				typ := reflect.TypeOf(value)
				if typ == nil {
					var fv any = "field"
					val = reflect.ValueOf(fv)
				}
				fieldType := val.Type()

				// produces a legit exported name
				anonymous := rapid.Bool().Draw(t, "is-field-anonymous")
				if fieldType.NumMethod() > 0 {
					anonymous = false // reflect limitation when creating structs
				}
				if fieldType.Kind() == reflect.Pointer {
					elem := fieldType.Elem()
					kind := elem.Kind()
					if kind == reflect.Pointer || kind == reflect.Interface { // prevent the creation of illegal **any or *any embedded fields
						anonymous = false
					}
				}

				field := reflect.StructField{
					Name:      names[i], // exported (otherwise not supported by [reflect.StructOf]
					Type:      fieldType,
					Anonymous: anonymous,
				}

				fields = append(fields, field)
				fieldValues = append(fieldValues, val)
			}

			structType := reflect.StructOf(fields)
			structVal := reflect.New(structType).Elem()

			for i := range structVal.NumField() {
				field := structVal.Field(i)
				field.Set(fieldValues[i])
			}

			// the struct has a zero value
			return structVal.Interface() // all fields are exported
		})
	}
}

func fieldNameGenerator() *rapid.Generator[string] {
	return rapid.Custom(func(t *rapid.T) string {
		first := rapid.StringOfN(unicodeUpperLetter(), 1, 1, -1).Draw(t, "first-letter")
		rest := rapid.StringOfN(unicodeLetterOrDigit(), 1, maxFieldNameLen, -1).Draw(t, "rest-of-field")

		return first + rest
	})
}

// genOtherValue generates values of type: chan, func()
// Should we add: unsafe.Pointer ??
//
// - chan {anything}
// - func() , func(int, string) *string
//
// Also add types that might produce odd behavior:
//
// - [sync.Mutex], [sync.RWMutex]
// - [atomic.Int64], [atomic.Value]
//
// # Limitations
//
// * channels are created with a random capacity, but are not written to.
// * only bi-directional channels are created.
func (g *typeGenerator) genOtherValue(depth uint32) *rapid.Generator[any] {
	return rapid.OneOf(
		g.genChanValue(depth+1),
		g.genFuncValue(depth+1),
		rapid.Custom(func(t *rapid.T) any {
			flipCoin := rapid.Bool().Draw(t, "flip-coin")
			if flipCoin {
				var mx sync.Mutex

				return &mx
			}

			var mx sync.RWMutex

			return &mx
		}),
		rapid.Custom(func(t *rapid.T) any {
			flipCoin := rapid.Bool().Draw(t, "flip-coin")
			if flipCoin {
				var aint64 atomic.Int64

				return &aint64
			}
			var av atomic.Value

			return &av
		}),
	)
}

func (g *typeGenerator) genChanValue(depth uint32) *rapid.Generator[any] {
	return rapid.Custom(func(t *rapid.T) any {
		value := g.genAnything(depth+1).Draw(t, "elem")
		buf := rapid.IntRange(0, maxChanBuf).Draw(t, "chan-buffers")

		val := reflect.ValueOf(value)
		if val.Kind() == reflect.Interface {
			if val.Elem().IsNil() {
				c := make(chan any, buf)

				return c
			}
			val = val.Elem()
		}

		if val.Kind() == reflect.Pointer && val.Elem().Kind() == reflect.Interface {
			val = val.Elem().Elem()
		}

		typ := reflect.TypeOf(value)
		if typ == nil || !val.IsValid() {
			c := make(chan *any, buf)

			return c
		}

		valType := val.Type()

		chanType := reflect.ChanOf(reflect.BothDir, valType)
		chanValue := reflect.MakeChan(chanType, buf)

		return chanValue.Interface()
	})
}

// genFuncValue only returns one of a few predeclared functions.
func (g *typeGenerator) genFuncValue(_ uint32) *rapid.Generator[any] {
	return rapid.OneOf(
		rapid.Just(emptyFunc).AsAny(),
		rapid.Just(signatureFunc).AsAny(),
	)
}

func emptyFunc()                            {}
func signatureFunc(_ int, _ string) *string { return nil }

// Proposal for enhancement: add more diversified interfaces.
func genInterfaceValue(_ uint32) *rapid.Generator[any] {
	var emptyIface any
	return rapid.Just(emptyIface).AsAny()
}

func unicodeUpperLetter() *rapid.Generator[rune] {
	return rapid.RuneFrom(nil, unicode.Upper) // NOTE: unlike go, we don't include "_" in go names
}

func unicodeLetterOrDigit() *rapid.Generator[rune] {
	return rapid.RuneFrom(nil, unicode.Letter, unicode.Digit)
}

func genPrimitiveValue() *rapid.Generator[any] {
	return rapid.OneOf(
		genIntegerValue(), // all integer types, incl. rune, byte, uintptr
		genFloatValue(),   // float32, float64, complex64, complex128
		rapid.Bool().AsAny(),
		rapid.String().AsAny(),
	)
}

func genIntegerValue() *rapid.Generator[any] {
	// NOTE: byte and rune are aliases. These types are slightly overrepresented in the sample.
	return rapid.OneOf(
		rapid.Byte().AsAny(),
		rapid.Int().AsAny(),
		rapid.Int16().AsAny(),
		rapid.Int32().AsAny(),
		rapid.Int64().AsAny(),
		rapid.Uint().AsAny(),
		rapid.Uint8().AsAny(),
		rapid.Uint16().AsAny(),
		rapid.Uint32().AsAny(),
		rapid.Uint64().AsAny(),
		rapid.Uintptr().AsAny(),
		rapid.Rune().AsAny(),
	)
}

func genFloatValue() *rapid.Generator[any] {
	return rapid.OneOf(
		rapid.Float32().AsAny(),
		rapid.Float64().AsAny(),
		rapid.Custom(func(t *rapid.T) any {
			realp := rapid.Float32().Draw(t, "real")
			imagp := rapid.Float32().Draw(t, "imag")

			return complex(realp, imagp)
		}),
		rapid.Custom(func(t *rapid.T) any {
			realp := rapid.Float64().Draw(t, "real")
			imagp := rapid.Float64().Draw(t, "imag")

			return complex(realp, imagp)
		}),
	)
}
