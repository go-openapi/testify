// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"go/types"
)

// isComparable reports whether typ satisfies the comparable constraint.
func isComparable(typ types.Type) bool {
	return types.Comparable(typ)
}

// isDeepComparable reports whether typ is comparable AND the == operator
// has the same semantics as reflect.DeepEqual. This is false for:
// - Pointer types (== compares addresses, DeepEqual compares targets)
// - Struct types containing pointer or interface fields
// - Interface types
// - Array types of non-deep-comparable elements
//
// This is critical for safe Equal→EqualT upgrades: EqualT uses == while
// Equal uses reflect.DeepEqual.
func isDeepComparable(typ types.Type) bool {
	if !types.Comparable(typ) {
		return false
	}
	return isDeepComparableUnderlying(typ, make(map[types.Type]bool))
}

func isDeepComparableUnderlying(typ types.Type, seen map[types.Type]bool) bool {
	// Prevent infinite recursion on recursive types.
	if seen[typ] {
		return true
	}
	seen[typ] = true

	under := typ.Underlying()

	switch t := under.(type) {
	case *types.Basic:
		return true
	case *types.Pointer:
		return false
	case *types.Interface:
		return false
	case *types.Struct:
		for field := range t.Fields() {
			if !isDeepComparableUnderlying(field.Type(), seen) {
				return false
			}
		}
		return true
	case *types.Array:
		return isDeepComparableUnderlying(t.Elem(), seen)
	default:
		return false
	}
}

// isOrdered reports whether typ satisfies the Ordered constraint
// (cmp.Ordered | []byte | time.Time).
func isOrdered(typ types.Type) bool {
	// Check for time.Time (struct, not ~struct, so check named type).
	if isTimeTime(typ) {
		return true
	}

	// Check for []byte.
	if isByteSlice(typ) {
		return true
	}

	// Check if the underlying type is in cmp.Ordered (string, int*, uint*, float*).
	return isCmpOrdered(typ)
}

// isCmpOrdered checks if a type satisfies cmp.Ordered (basic ordered types).
func isCmpOrdered(typ types.Type) bool {
	under := typ.Underlying()
	basic, ok := under.(*types.Basic)
	if !ok {
		return false
	}
	switch basic.Kind() {
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
		types.Uintptr,
		types.Float32, types.Float64,
		types.String:
		return true
	default:
		return false
	}
}

// isText reports whether typ satisfies the Text constraint (~string | ~[]byte).
func isText(typ types.Type) bool {
	under := typ.Underlying()

	// ~string
	if basic, ok := under.(*types.Basic); ok && basic.Kind() == types.String {
		return true
	}

	// ~[]byte
	return isByteSlice(typ)
}

// isSignedNumeric reports whether typ satisfies the SignedNumeric constraint.
func isSignedNumeric(typ types.Type) bool {
	under := typ.Underlying()
	basic, ok := under.(*types.Basic)
	if !ok {
		return false
	}
	switch basic.Kind() {
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Float32, types.Float64:
		return true
	default:
		return false
	}
}

// isMeasurable reports whether typ satisfies the Measurable constraint
// (SignedNumeric | UnsignedNumeric | ~float32 | ~float64).
func isMeasurable(typ types.Type) bool {
	under := typ.Underlying()
	basic, ok := under.(*types.Basic)
	if !ok {
		return false
	}
	switch basic.Kind() {
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64,
		types.Float32, types.Float64:
		return true
	default:
		return false
	}
}

// isBoolean reports whether typ satisfies the Boolean constraint (~bool).
func isBoolean(typ types.Type) bool {
	under := typ.Underlying()
	basic, ok := under.(*types.Basic)
	if !ok {
		return false
	}
	return basic.Kind() == types.Bool
}

// isRegExp reports whether typ satisfies the RegExp constraint
// (Text | *regexp.Regexp).
func isRegExp(typ types.Type) bool {
	if isText(typ) {
		return true
	}
	return isRegexpPointer(typ)
}

// isPointerType reports whether typ is a pointer type. Returns the element type.
func isPointerType(typ types.Type) (elem types.Type, ok bool) {
	ptr, ok := typ.Underlying().(*types.Pointer)
	if !ok {
		return nil, false
	}
	return ptr.Elem(), true
}

// isSliceType reports whether typ is a slice type. Returns the element type.
func isSliceType(typ types.Type) (elem types.Type, ok bool) {
	sl, ok := typ.Underlying().(*types.Slice)
	if !ok {
		return nil, false
	}
	return sl.Elem(), true
}

// isMapType reports whether typ is a map type. Returns key and value types.
func isMapType(typ types.Type) (key, val types.Type, ok bool) {
	m, ok := typ.Underlying().(*types.Map)
	if !ok {
		return nil, nil, false
	}
	return m.Key(), m.Elem(), true
}

// sameType reports whether two types are identical.
func sameType(a, b types.Type) bool {
	return types.Identical(a, b)
}

// isAnyOrInterface reports whether typ is interface{} or any, which means
// we cannot determine the concrete type statically.
func isAnyOrInterface(typ types.Type) bool {
	if typ == nil {
		return true
	}
	iface, ok := typ.Underlying().(*types.Interface)
	if !ok {
		return false
	}
	return iface.Empty()
}

// satisfiesConstraint checks whether a type satisfies the given constraint kind.
func satisfiesConstraint(typ types.Type, c constraintKind) bool {
	if isAnyOrInterface(typ) {
		return false
	}
	switch c {
	case constraintComparable:
		return isComparable(typ)
	case constraintDeepComparable:
		return isDeepComparable(typ)
	case constraintOrdered:
		return isOrdered(typ)
	case constraintText:
		return isText(typ)
	case constraintSignedNumeric:
		return isSignedNumeric(typ)
	case constraintMeasurable:
		return isMeasurable(typ)
	case constraintPointer:
		_, ok := isPointerType(typ)
		return ok
	case constraintBoolean:
		return isBoolean(typ)
	case constraintRegExp:
		return isRegExp(typ)
	default:
		return false
	}
}

// Helper functions

func isByteSlice(typ types.Type) bool {
	sl, ok := typ.Underlying().(*types.Slice)
	if !ok {
		return false
	}
	basic, ok := sl.Elem().(*types.Basic)
	return ok && basic.Kind() == types.Byte
}

func isTimeTime(typ types.Type) bool {
	named, ok := typ.(*types.Named)
	if !ok {
		return false
	}
	obj := named.Obj()
	return obj.Name() == "Time" && obj.Pkg() != nil && obj.Pkg().Path() == "time"
}

func isRegexpPointer(typ types.Type) bool {
	ptr, ok := typ.Underlying().(*types.Pointer)
	if !ok {
		return false
	}
	named, ok := ptr.Elem().(*types.Named)
	if !ok {
		return false
	}
	obj := named.Obj()
	return obj.Name() == "Regexp" && obj.Pkg() != nil && obj.Pkg().Path() == "regexp"
}
