package spew

import (
	"reflect"
	"time"
)

const maxDepth = 1000

// isTime detects if a value may be assimilated to time.Time.
//
// It may be time.Time or a *time.Time,
// but also a redeclared type convertible to time.Time or *time.Time,
//
// Conversely, a struct that embeds a time.Time or *time.Time is not considered a time.Time
// and we'll have to dig the individual fields.
//
// Drilling down pointers indirections stops at depth 1000. Over that limit, the check assumes
// that this is (probably) not a time.Time... That's a theoretical bound, nothing practical
// describes a **** ... ***time.Time with 1000 stars.
func isTime(v reflect.Value, depth int) bool {
	if depth > maxDepth {
		return false
	}

	if !v.IsValid() {
		return false
	}

	k := v.Kind()
	t := v.Type()

	// for pointers, we reason about the pointer type, because the value may be nil
	if k == reflect.Pointer && t.Elem().Kind() == reflect.Pointer {
		return isTime(v.Elem(), depth+1)
	}

	if k == reflect.Struct || (k == reflect.Pointer && t.Elem().Kind() == reflect.Struct) {
		return v.CanConvert(reflect.TypeFor[time.Time]()) ||
			v.CanConvert(reflect.TypeFor[*time.Time]())
	}

	return false
}

// isConvertibleToTime returns a converted reflect.Value and true when v is convertible to time.Time or *time.Time.
func isConvertibleToTime(v reflect.Value) (reflect.Value, bool) {
	if !v.IsValid() {
		return reflect.Value{}, false
	}

	k := v.Kind()

	timeTyp := reflect.TypeFor[time.Time]()
	if k == reflect.Struct && v.CanConvert(timeTyp) {
		return v.Convert(timeTyp), true
	}

	timePtrTyp := reflect.TypeFor[*time.Time]()
	if k == reflect.Pointer && v.Elem().Kind() == reflect.Struct && v.CanConvert(timePtrTyp) {
		return v.Convert(timePtrTyp), true
	}

	if k == reflect.Pointer && v.Elem().Kind() == reflect.Pointer {
		return isConvertibleToTime(v.Elem())
	}

	return reflect.Value{}, false // the returned value is Invalid in this case
}
