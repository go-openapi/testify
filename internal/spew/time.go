package spew

import (
	"reflect"
	"time"
)

// isTime detects if a value may be assimilated to time.Time.
//
// It may be time.Time or a *time.Time,
// but also a redeclared type convertible to time.Time or *time.Time,
//
// Conversely, a struct that embeds a time.Time or *time.Time is not considered a time.Time
// and we'll have to dig the individual fields.
func isTime(v reflect.Value) bool {
	k := v.Kind()
	t := v.Type()

	// for pointers, we reason about the pointer type, because the value may be nil
	if k == reflect.Pointer && t.Elem().Kind() == reflect.Pointer {
		return isTime(v.Elem())
	}

	if k == reflect.Struct || (k == reflect.Pointer && t.Elem().Kind() == reflect.Struct) {
		return v.CanConvert(reflect.TypeFor[time.Time]()) ||
			v.CanConvert(reflect.TypeFor[*time.Time]())
	}

	return false
}

// isConvertibleToTime returns a converted reflect.Value and true when v is convertible to time.Time or *time.Time.
func isConvertibleToTime(v reflect.Value) (reflect.Value, bool) {
	k := v.Kind()

	timeTyp := reflect.TypeFor[time.Time]()
	if k == reflect.Struct && v.CanConvert(timeTyp) {
		return v.Convert(timeTyp), true
	}

	timePtrTyp := reflect.TypeFor[*time.Time]()
	if k == reflect.Pointer && v.Elem().Kind() == reflect.Struct && v.CanConvert(timePtrTyp) {
		return v.Convert(timePtrTyp), true
	}

	return reflect.Value{}, false // the returned value is Invalid in this case
}
