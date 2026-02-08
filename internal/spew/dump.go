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
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

//nolint:gochecknoglobals // immutable reflect type and compiled regexps
var (
	// uint8Type is a reflect.Type representing a uint8.  It is used to
	// convert cgo types to uint8 slices for hexdumping.
	uint8Type = reflect.TypeFor[uint8]()

	// cCharRE is a regular expression that matches a cgo char.
	// It is used to detect character arrays to hexdump them.
	cCharRE = regexp.MustCompile(`^.*\._Ctype_char$`)

	// cUnsignedCharRE is a regular expression that matches a cgo unsigned
	// char.  It is used to detect unsigned character arrays to hexdump
	// them.
	cUnsignedCharRE = regexp.MustCompile(`^.*\._Ctype_unsignedchar$`)

	// cUint8tCharRE is a regular expression that matches a cgo uint8_t.
	// It is used to detect uint8_t arrays to hexdump them.
	cUint8tCharRE = regexp.MustCompile(`^.*\._Ctype_uint8_t$`)
)

// dumpState contains information about the state of a dump operation.
type dumpState struct {
	w                io.Writer
	depth            int
	pointers         map[uintptr]int
	ignoreNextType   bool
	ignoreNextIndent bool
	cs               *ConfigState
}

// indent performs indentation according to the depth level and cs.Indent
// option.
func (d *dumpState) indent() {
	if d.ignoreNextIndent {
		d.ignoreNextIndent = false
		return
	}
	_, _ = d.w.Write(bytes.Repeat([]byte(d.cs.Indent), d.depth))
}

// unpackValue returns values inside of non-nil interfaces when possible.
// This is useful for data types like structs, arrays, slices, and maps which
// can contain varying types packed inside an interface.
func (d *dumpState) unpackValue(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Interface && !v.IsNil() {
		v = v.Elem()
	}
	return v
}

// dumpPtr handles formatting of pointers by indirecting them as necessary.
func (d *dumpState) dumpPtr(v reflect.Value) {
	// Remove pointers at or below the current depth from map used to detect
	// circular refs.
	for k, depth := range d.pointers {
		if depth >= d.depth {
			delete(d.pointers, k)
		}
	}

	r := resolvePtr(v, d.depth, d.pointers)

	// Display type information.
	_, _ = d.w.Write(openParenBytes)
	_, _ = d.w.Write(bytes.Repeat(asteriskBytes, r.indirects))
	_, _ = d.w.Write([]byte(r.value.Type().String()))
	_, _ = d.w.Write(closeParenBytes)

	// Display pointer information.
	if !d.cs.DisablePointerAddresses && len(r.pointerChain) > 0 {
		_, _ = d.w.Write(openParenBytes)
		for i, addr := range r.pointerChain {
			if i > 0 {
				_, _ = d.w.Write(pointerChainBytes)
			}
			printHexPtr(d.w, addr)
		}
		_, _ = d.w.Write(closeParenBytes)
	}

	// Display dereferenced value.
	_, _ = d.w.Write(openParenBytes)
	switch {
	case r.nilFound:
		_, _ = d.w.Write(nilAngleBytes)
	case r.cycleFound:
		_, _ = d.w.Write(circularBytes)
	default:
		d.ignoreNextType = true
		d.dump(r.value)
	}
	_, _ = d.w.Write(closeParenBytes)
}

func (d *dumpState) dumpMap(v reflect.Value) {
	// Remove pointers at or below the current depth from map used to detect
	// circular refs.
	for k, depth := range d.pointers {
		if depth >= d.depth {
			delete(d.pointers, k)
		}
	}

	// Keep list of all dereferenced pointers to show later.
	cycleFound := false

	// nil maps should be indicated as different than empty maps
	if v.IsNil() {
		_, _ = d.w.Write(nilAngleBytes)
		return
	}

	// maps like pointers may present circular references
	addr := v.Pointer()
	if pd, ok := d.pointers[addr]; ok && pd <= d.depth {
		cycleFound = true
	}
	d.pointers[addr] = d.depth

	_, _ = d.w.Write(openBraceNewlineBytes)
	d.depth++

	switch {
	case d.cs.MaxDepth != 0 && d.depth > d.cs.MaxDepth:
		d.indent()
		_, _ = d.w.Write(maxNewlineBytes)
	case cycleFound:
		_, _ = d.w.Write(circularBytes)
	default:
		numEntries := v.Len()
		keys := v.MapKeys()
		if d.cs.SortKeys {
			sortValues(keys, d.cs)
		}
		for i, key := range keys {
			d.dump(d.unpackValue(key))
			_, _ = d.w.Write(colonSpaceBytes)
			d.ignoreNextIndent = true
			d.dump(d.unpackValue(v.MapIndex(key)))
			if i < (numEntries - 1) {
				_, _ = d.w.Write(commaNewlineBytes)
			} else {
				_, _ = d.w.Write(newlineBytes)
			}
		}
	}

	d.depth--
	d.indent()
	_, _ = d.w.Write(closeBraceBytes)
}

// resolveHexDump determines whether a slice should be hex-dumped and returns
// the byte buffer if so. Returns (nil, false) if the slice should not be
// hex-dumped.
func (d *dumpState) resolveHexDump(v reflect.Value) ([]uint8, bool) {
	numEntries := v.Len()
	if numEntries == 0 {
		return nil, false
	}

	vt := v.Index(0).Type()
	vts := vt.String()

	doConvert := false
	switch {
	// C types that need to be converted.
	case cCharRE.MatchString(vts),
		cUnsignedCharRE.MatchString(vts),
		cUint8tCharRE.MatchString(vts):
		doConvert = true

	// Try to use existing uint8 slices and fall back to converting
	// and copying if that fails.
	case vt.Kind() == reflect.Uint8:
		if buf, ok := d.tryUint8Slice(v, numEntries); ok {
			return buf, true
		}
		doConvert = true
	}

	if doConvert && vt.ConvertibleTo(uint8Type) {
		buf := make([]uint8, numEntries)
		for i := range numEntries {
			vv := v.Index(i)
			buf[i] = uint8(vv.Convert(uint8Type).Uint()) //nolint:gosec // conversion is fine: the original type is uint8
		}
		return buf, true
	}

	return nil, false
}

// tryUint8Slice attempts to directly extract a []uint8 from a reflect.Value
// whose elements are uint8. Returns the slice and true if successful.
func (d *dumpState) tryUint8Slice(v reflect.Value, numEntries int) ([]uint8, bool) {
	vs := v
	if !vs.CanInterface() || !vs.CanAddr() {
		vs = unsafeReflectValue(vs)
	}
	if UnsafeDisabled {
		return nil, false
	}
	vs = vs.Slice(0, numEntries)
	iface := vs.Interface()
	if slice, ok := iface.([]uint8); ok {
		return slice, true
	}
	return nil, false
}

// dumpSlice handles formatting of arrays and slices.  Byte (uint8 under
// reflection) arrays and slices are dumped in hexdump -C fashion.
func (d *dumpState) dumpSlice(v reflect.Value) {
	// Determine whether this type should be hex dumped or not.  Also,
	// for types which should be hexdumped, try to use the underlying data
	// first, then fall back to trying to convert them to a uint8 slice.
	numEntries := v.Len()
	buf, doHexDump := d.resolveHexDump(v)

	// Hexdump the entire slice as needed.
	if doHexDump {
		indent := strings.Repeat(d.cs.Indent, d.depth)
		str := indent + hex.Dump(buf)
		str = strings.ReplaceAll(str, "\n", "\n"+indent)
		str = strings.TrimRight(str, d.cs.Indent)
		_, _ = d.w.Write([]byte(str))
		return
	}

	// Recursively call dump for each item.
	for i := range numEntries {
		d.dump(d.unpackValue(v.Index(i)))
		if i < (numEntries - 1) {
			_, _ = d.w.Write(commaNewlineBytes)
		} else {
			_, _ = d.w.Write(newlineBytes)
		}
	}
}

// dumpLenCap displays length and capacity if the built-in len and cap
// functions work with the value's kind and the values are non-zero.
func (d *dumpState) dumpLenCap(v reflect.Value) {
	valueLen, valueCap := 0, 0
	switch v.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan:
		valueLen, valueCap = v.Len(), v.Cap()
	case reflect.Map, reflect.String:
		valueLen = v.Len()
	default:
	}

	if valueLen == 0 && (d.cs.DisableCapacities || valueCap == 0) {
		return
	}

	_, _ = d.w.Write(openParenBytes)
	if valueLen != 0 {
		_, _ = d.w.Write(lenEqualsBytes)
		printInt(d.w, int64(valueLen), decimalBase)
	}
	if !d.cs.DisableCapacities && valueCap != 0 {
		if valueLen != 0 {
			_, _ = d.w.Write(spaceBytes)
		}
		_, _ = d.w.Write(capEqualsBytes)
		printInt(d.w, int64(valueCap), decimalBase)
	}
	_, _ = d.w.Write(closeParenBytes)
	_, _ = d.w.Write(spaceBytes)
}

// dump is the main workhorse for dumping a value.  It uses the passed reflect
// value to figure out what kind of object we are dealing with and formats it
// appropriately.  It is a recursive function, however circular data structures
// are detected and handled properly.
func (d *dumpState) dump(v reflect.Value) {
	// Handle invalid reflect values immediately.
	kind := v.Kind()
	if kind == reflect.Invalid {
		_, _ = d.w.Write(invalidAngleBytes)
		return
	}

	// Handle pointers specially.
	if kind == reflect.Pointer {
		d.indent()
		d.dumpPtr(v)
		return
	}

	// Print type information unless already handled elsewhere.
	if !d.ignoreNextType {
		d.indent()
		_, _ = d.w.Write(openParenBytes)
		_, _ = d.w.Write([]byte(v.Type().String()))
		_, _ = d.w.Write(closeParenBytes)
		_, _ = d.w.Write(spaceBytes)
	}
	d.ignoreNextType = false

	// Display length and capacity if the built-in len and cap functions
	// work with the value's kind and the len/cap itself is non-zero.
	d.dumpLenCap(v)

	// Call Stringer/error interfaces if they exist and the handle methods flag
	// is enabled
	if tryHandleMethods(d.cs, d.w, v, kind) {
		return
	}

	d.dumpValue(v, kind)
}

// dumpValue handles the type-specific formatting for the dump method.
func (d *dumpState) dumpValue(v reflect.Value, kind reflect.Kind) {
	switch kind {
	case reflect.Invalid:
		// Do nothing.  We should never get here since invalid has already
		// been handled above.

	case reflect.Bool:
		printBool(d.w, v.Bool())

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Int:
		printInt(d.w, v.Int(), decimalBase)

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uint:
		printUint(d.w, v.Uint(), decimalBase)

	case reflect.Float32:
		printFloat(d.w, v.Float(), float32Precision)

	case reflect.Float64:
		printFloat(d.w, v.Float(), float64Precision)

	case reflect.Complex64:
		printComplex(d.w, v.Complex(), complex64Precision)

	case reflect.Complex128:
		printComplex(d.w, v.Complex(), complex128Precision)

	case reflect.Slice:
		if v.IsNil() {
			_, _ = d.w.Write(nilAngleBytes)
			break
		}
		d.dumpArray(v)

	case reflect.Array:
		d.dumpArray(v)

	case reflect.String:
		_, _ = d.w.Write([]byte(strconv.Quote(v.String())))

	case reflect.Interface:
		// The only time we should get here is for nil interfaces due to
		// unpackValue calls.
		if v.IsNil() {
			_, _ = d.w.Write(nilAngleBytes)
		}

	case reflect.Pointer:
		// Do nothing.  We should never get here since pointers have already
		// been handled above.

	case reflect.Map:
		d.dumpMap(v)

	case reflect.Struct:
		d.dumpStruct(v)

	case reflect.Uintptr:
		printHexPtr(d.w, uintptr(v.Uint()))

	case reflect.UnsafePointer, reflect.Chan, reflect.Func:
		printHexPtr(d.w, v.Pointer())

	// There were not any other types at the time this code was written, but
	// fall back to letting the default fmt package handle it in case any new
	// types are added.
	default:
		if v.CanInterface() {
			_, _ = fmt.Fprintf(d.w, "%v", v.Interface())
		} else {
			_, _ = fmt.Fprintf(d.w, "%v", v.String())
		}
	}
}

// dumpArray handles formatting of array and non-nil slice values.
func (d *dumpState) dumpArray(v reflect.Value) {
	_, _ = d.w.Write(openBraceNewlineBytes)
	d.depth++
	if (d.cs.MaxDepth != 0) && (d.depth > d.cs.MaxDepth) {
		d.indent()
		_, _ = d.w.Write(maxNewlineBytes)
	} else {
		d.dumpSlice(v)
	}
	d.depth--
	d.indent()
	_, _ = d.w.Write(closeBraceBytes)
}

// dumpStruct handles formatting of struct values.
func (d *dumpState) dumpStruct(v reflect.Value) {
	_, _ = d.w.Write(openBraceNewlineBytes)
	d.depth++
	if (d.cs.MaxDepth != 0) && (d.depth > d.cs.MaxDepth) {
		d.indent()
		_, _ = d.w.Write(maxNewlineBytes)
	} else {
		vt := v.Type()
		numFields := v.NumField()
		for i := range numFields {
			d.indent()
			vtf := vt.Field(i)
			_, _ = d.w.Write([]byte(vtf.Name))
			_, _ = d.w.Write(colonSpaceBytes)
			d.ignoreNextIndent = true
			d.dump(d.unpackValue(v.Field(i)))
			if i < (numFields - 1) {
				_, _ = d.w.Write(commaNewlineBytes)
			} else {
				_, _ = d.w.Write(newlineBytes)
			}
		}
	}
	d.depth--
	d.indent()
	_, _ = d.w.Write(closeBraceBytes)
}

// fdump is a helper function to consolidate the logic from the various public
// methods which take varying writers and config states.
func fdump(cs *ConfigState, w io.Writer, a ...any) {
	for _, arg := range a {
		if arg == nil {
			_, _ = w.Write(interfaceBytes)
			_, _ = w.Write(spaceBytes)
			_, _ = w.Write(nilAngleBytes)
			_, _ = w.Write(newlineBytes)
			continue
		}

		d := dumpState{w: w, cs: cs}
		d.pointers = make(map[uintptr]int)
		d.dump(reflect.ValueOf(arg))
		_, _ = d.w.Write(newlineBytes)
	}
}

// Fdump formats and displays the passed arguments to io.Writer w.  It formats
// exactly the same as Dump.
func Fdump(w io.Writer, a ...any) {
	fdump(&Config, w, a...)
}

// Sdump returns a string with the passed arguments formatted exactly the same
// as Dump.
func Sdump(a ...any) string {
	var buf bytes.Buffer
	fdump(&Config, &buf, a...)
	return buf.String()
}

/*
Dump displays the passed parameters to standard out with newlines, customizable
indentation, and additional debug information such as complete types and all
pointer addresses used to indirect to the final value.  It provides the
following features over the built-in printing facilities provided by the fmt
package:

  - Pointers are dereferenced and followed
  - Circular data structures are detected and handled properly
  - Custom Stringer/error interfaces are optionally invoked, including
    on unexported types
  - Custom types which only implement the Stringer/error interfaces via
    a pointer receiver are optionally invoked when passing non-pointer
    variables
  - Byte arrays and slices are dumped like the hexdump -C command which
    includes offsets, byte values in hex, and ASCII output

The configuration options are controlled by an exported package global,
spew.Config.  See ConfigState for options documentation.

See Fdump if you would prefer dumping to an arbitrary io.Writer or Sdump to
get the formatted result as a string.
*/
func Dump(a ...any) {
	fdump(&Config, os.Stdout, a...)
}
