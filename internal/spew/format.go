/*
 * Copyright (c) 2013-2016 Dave Collins <dave@davec.name>
 *
 * Permission to use, copy, modify, and distribute this software for any
 * purpose with or without fee is hereby granted, provided that the above
 * copyright notice and this permission notice appear in all copies.
 *
 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
 * WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
 * MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
 * ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
 * ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
 * OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
 */

package spew

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// supportedFlags is a list of all the character flags supported by fmt package.
const supportedFlags = "0-+# "

// formatState implements the fmt.Formatter interface and contains information
// about the state of a formatting operation.  The NewFormatter function can
// be used to get a new Formatter which can be used directly as arguments
// in standard fmt package printing calls.
type formatState struct {
	value          any
	fs             fmt.State
	depth          int
	pointers       map[uintptr]int
	ignoreNextType bool
	cs             *ConfigState
}

// Format satisfies the fmt.Formatter interface. See NewFormatter for usage
// details.
func (f *formatState) Format(fs fmt.State, verb rune) {
	f.fs = fs

	// Use standard formatting for verbs that are not v.
	if verb != 'v' {
		format := f.constructOrigFormat(verb)
		_, _ = fmt.Fprintf(fs, format, f.value)
		return
	}

	if f.value == nil {
		if fs.Flag('#') {
			_, _ = fs.Write(interfaceBytes)
		}
		_, _ = fs.Write(nilAngleBytes)
		return
	}

	f.format(reflect.ValueOf(f.value))
}

// buildDefaultFormat recreates the original format string without precision
// and width information to pass in to fmt.Sprintf in the case of an
// unrecognized type.  Unless new types are added to the language, this
// function won't ever be called.
func (f *formatState) buildDefaultFormat() (format string) {
	buf := bytes.NewBuffer(percentBytes)

	for _, flag := range supportedFlags {
		if f.fs.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}

	buf.WriteRune('v')

	format = buf.String()
	return format
}

// constructOrigFormat recreates the original format string including precision
// and width information to pass along to the standard fmt package.  This allows
// automatic deferral of all format strings this package doesn't support.
func (f *formatState) constructOrigFormat(verb rune) (format string) {
	buf := bytes.NewBuffer(percentBytes)

	for _, flag := range supportedFlags {
		if f.fs.Flag(int(flag)) {
			buf.WriteRune(flag)
		}
	}

	if width, ok := f.fs.Width(); ok {
		buf.WriteString(strconv.Itoa(width))
	}

	if precision, ok := f.fs.Precision(); ok {
		buf.Write(precisionBytes)
		buf.WriteString(strconv.Itoa(precision))
	}

	buf.WriteRune(verb)

	format = buf.String()
	return format
}

// unpackValue returns values inside of non-nil interfaces when possible and
// ensures that types for values which have been unpacked from an interface
// are displayed when the show types flag is also set.
// This is useful for data types like structs, arrays, slices, and maps which
// can contain varying types packed inside an interface.
func (f *formatState) unpackValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface {
		f.ignoreNextType = false
		if !v.IsNil() {
			v = v.Elem()
		}
	}
	return v
}

// formatPtr handles formatting of pointers by indirecting them as necessary.
func (f *formatState) formatPtr(v reflect.Value) {
	// Display nil if top level pointer is nil.
	showTypes := f.fs.Flag('#')
	if v.IsNil() && (!showTypes || f.ignoreNextType) {
		_, _ = f.fs.Write(nilAngleBytes)
		return
	}

	// Remove pointers at or below the current depth from map used to detect
	// circular refs.
	for k, depth := range f.pointers {
		if depth >= f.depth {
			delete(f.pointers, k)
		}
	}

	r := resolvePtr(v, f.depth, f.pointers)

	// Display type or indirection level depending on flags.
	if showTypes && !f.ignoreNextType {
		_, _ = f.fs.Write(openParenBytes)
		_, _ = f.fs.Write(bytes.Repeat(asteriskBytes, r.indirects))
		_, _ = f.fs.Write([]byte(r.value.Type().String()))
		_, _ = f.fs.Write(closeParenBytes)
	} else {
		indirects := r.indirects
		if r.nilFound || r.cycleFound {
			indirects += strings.Count(r.value.Type().String(), "*")
		}
		_, _ = f.fs.Write(openAngleBytes)
		_, _ = f.fs.Write([]byte(strings.Repeat("*", indirects)))
		_, _ = f.fs.Write(closeAngleBytes)
	}

	// Display pointer information depending on flags.
	if f.fs.Flag('+') && (len(r.pointerChain) > 0) {
		_, _ = f.fs.Write(openParenBytes)
		for i, addr := range r.pointerChain {
			if i > 0 {
				_, _ = f.fs.Write(pointerChainBytes)
			}
			printHexPtr(f.fs, addr)
		}
		_, _ = f.fs.Write(closeParenBytes)
	}

	// Display dereferenced value.
	switch {
	case r.nilFound:
		_, _ = f.fs.Write(nilAngleBytes)
	case r.cycleFound:
		_, _ = f.fs.Write(circularShortBytes)
	default:
		f.ignoreNextType = true
		f.format(r.value)
	}
}

// format is the main workhorse for providing the Formatter interface.  It
// uses the passed reflect value to figure out what kind of object we are
// dealing with and formats it appropriately.  It is a recursive function,
// however circular data structures are detected and handled properly.
func (f *formatState) format(v reflect.Value) {
	// Handle invalid reflect values immediately.
	kind := v.Kind()
	if kind == reflect.Invalid {
		_, _ = f.fs.Write(invalidAngleBytes)
		return
	}

	// Handle pointers specially.
	if kind == reflect.Pointer {
		f.formatPtr(v)
		return
	}

	// Print type information unless already handled elsewhere.
	if !f.ignoreNextType && f.fs.Flag('#') {
		_, _ = f.fs.Write(openParenBytes)
		_, _ = f.fs.Write([]byte(v.Type().String()))
		_, _ = f.fs.Write(closeParenBytes)
	}
	f.ignoreNextType = false

	// Call Stringer/error interfaces if they exist and the handle methods
	// flag is enabled.
	if tryHandleMethods(f.cs, f.fs, v, kind) {
		return
	}

	f.formatValue(v, kind)
}

// formatValue handles the type-specific formatting for the format method.
func (f *formatState) formatValue(v reflect.Value, kind reflect.Kind) {
	switch kind {
	case reflect.Invalid:
		// Do nothing.  We should never get here since invalid has already
		// been handled above.

	case reflect.Bool:
		printBool(f.fs, v.Bool())

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		printInt(f.fs, v.Int(), decimalBase)

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		printUint(f.fs, v.Uint(), decimalBase)

	case reflect.Float32:
		printFloat(f.fs, v.Float(), float32Precision)

	case reflect.Float64:
		printFloat(f.fs, v.Float(), float64Precision)

	case reflect.Complex64:
		printComplex(f.fs, v.Complex(), complex64Precision)

	case reflect.Complex128:
		printComplex(f.fs, v.Complex(), complex128Precision)

	case reflect.Slice:
		if v.IsNil() {
			_, _ = f.fs.Write(nilAngleBytes)
			break
		}
		fallthrough

	case reflect.Array:
		f.formatArray(v)

	case reflect.String:
		_, _ = f.fs.Write([]byte(v.String()))

	case reflect.Interface:
		// The only time we should get here is for nil interfaces due to
		// unpackValue calls.
		if v.IsNil() {
			_, _ = f.fs.Write(nilAngleBytes)
		}

	case reflect.Pointer:
		// Do nothing.  We should never get here since pointers have already
		// been handled above.

	case reflect.Map:
		f.formatMap(v)

	case reflect.Struct:
		f.formatStruct(v)

	case reflect.Uintptr:
		printHexPtr(f.fs, uintptr(v.Uint()))

	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
		printHexPtr(f.fs, v.Pointer())

	// There were not any other types at the time this code was written, but
	// fall back to letting the default fmt package handle it if any get added.
	default:
		format := f.buildDefaultFormat()
		if v.CanInterface() {
			_, _ = fmt.Fprintf(f.fs, format, v.Interface())
		} else {
			_, _ = fmt.Fprintf(f.fs, format, v.String())
		}
	}
}

// formatArray handles formatting of array and slice values.
func (f *formatState) formatArray(v reflect.Value) {
	_, _ = f.fs.Write(openBracketBytes)
	f.depth++
	if (f.cs.MaxDepth != 0) && (f.depth > f.cs.MaxDepth) {
		_, _ = f.fs.Write(maxShortBytes)
	} else {
		numEntries := v.Len()
		for i := range numEntries {
			if i > 0 {
				_, _ = f.fs.Write(spaceBytes)
			}
			f.ignoreNextType = true
			f.format(f.unpackValue(v.Index(i)))
		}
	}
	f.depth--
	_, _ = f.fs.Write(closeBracketBytes)
}

// formatMap handles formatting of map values.
func (f *formatState) formatMap(v reflect.Value) {
	// nil maps should be indicated as different than empty maps
	if v.IsNil() {
		_, _ = f.fs.Write(nilAngleBytes)
		return
	}

	_, _ = f.fs.Write(openMapBytes)
	f.depth++
	if (f.cs.MaxDepth != 0) && (f.depth > f.cs.MaxDepth) {
		_, _ = f.fs.Write(maxShortBytes)
	} else {
		keys := v.MapKeys()
		if f.cs.SortKeys {
			sortValues(keys, f.cs)
		}
		for i, key := range keys {
			if i > 0 {
				_, _ = f.fs.Write(spaceBytes)
			}
			f.ignoreNextType = true
			f.format(f.unpackValue(key))
			_, _ = f.fs.Write(colonBytes)
			f.ignoreNextType = true
			f.format(f.unpackValue(v.MapIndex(key)))
		}
	}
	f.depth--
	_, _ = f.fs.Write(closeMapBytes)
}

// formatStruct handles formatting of struct values.
func (f *formatState) formatStruct(v reflect.Value) {
	numFields := v.NumField()
	_, _ = f.fs.Write(openBraceBytes)
	f.depth++
	if (f.cs.MaxDepth != 0) && (f.depth > f.cs.MaxDepth) {
		_, _ = f.fs.Write(maxShortBytes)
	} else {
		vt := v.Type()
		for i := range numFields {
			if i > 0 {
				_, _ = f.fs.Write(spaceBytes)
			}
			vtf := vt.Field(i)
			if f.fs.Flag('+') || f.fs.Flag('#') {
				_, _ = f.fs.Write([]byte(vtf.Name))
				_, _ = f.fs.Write(colonBytes)
			}
			f.format(f.unpackValue(v.Field(i)))
		}
	}
	f.depth--
	_, _ = f.fs.Write(closeBraceBytes)
}

// newFormatter is a helper function to consolidate the logic from the various
// public methods which take varying config states.
func newFormatter(cs *ConfigState, v any) fmt.Formatter {
	fs := &formatState{value: v, cs: cs}
	fs.pointers = make(map[uintptr]int)
	return fs
}

/*
NewFormatter returns a custom formatter that satisfies the fmt.Formatter
interface.  As a result, it integrates cleanly with standard fmt package
printing functions.  The formatter is useful for inline printing of smaller data
types similar to the standard %v format specifier.

The custom formatter only responds to the %v (most compact), %+v (adds pointer
addresses), %#v (adds types), or %#+v (adds types and pointer addresses) verb
combinations.

Any other verbs such as %x and %q will be sent to the standard fmt package for formatting.
In addition, the custom formatter ignores the width and precision arguments
(however they will still work on the format specifiers not handled by the custom formatter).

Typically this function shouldn't be called directly.  It is much easier to make
use of the custom formatter by calling one of the convenience functions such as
Printf, Println, or Fprintf.
*/
func NewFormatter(v any) fmt.Formatter {
	return newFormatter(&Config, v)
}
