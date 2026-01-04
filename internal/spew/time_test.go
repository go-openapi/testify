package spew

import (
	"iter"
	"reflect"
	"slices"
	"testing"
	"time"
)

func TestIsTime(t *testing.T) {
	t.Parallel()

	for tt := range isTimeCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			val := reflect.ValueOf(tt.value)
			if result := isTime(val); result != tt.expectedTime {
				t.Errorf("expected %v to be considered a time.Time", tt.value)
			}
		})
	}
}

type isTimeCase struct {
	name         string
	value        any
	expectedTime bool
}

func isTimeCases() iter.Seq[isTimeCase] {
	type tr time.Time
	type trptr *time.Time
	type te struct {
		time.Time
	}
	nilTimePtr := (*time.Time)(nil)

	return slices.Values([]isTimeCase{
		{
			name:         "time.Time",
			value:        time.Now(),
			expectedTime: true,
		},
		{
			name:         "*time.Time",
			value:        ptr(time.Now()),
			expectedTime: true,
		},
		{
			name:         "nil *time.Time",
			value:        nilTimePtr,
			expectedTime: true,
		},
		{
			name:         "redefined time.Time",
			value:        tr{},
			expectedTime: true,
		},
		{
			name:         "pointer to redefined time.Time",
			value:        &tr{},
			expectedTime: true,
		},
		{
			name:         "pointer to redefined *time.Time",
			value:        trptr(ptr(time.Now())),
			expectedTime: true,
		},
		{
			name:         "pointer to nil redefined *time.Time",
			value:        trptr(nil),
			expectedTime: true,
		},
		{
			name:         "pointer indirection **time.Time",
			value:        ptr(ptr(time.Now())),
			expectedTime: true,
		},
		{
			name:         "pointer indirection to nil **time.Time",
			value:        ptr(nilTimePtr),
			expectedTime: true,
		},
		{
			name:         "embedded time.Time",
			value:        te{},
			expectedTime: false,
		},
		{
			name:         "pointer to embedded time.Time",
			value:        &te{},
			expectedTime: false,
		},
		{
			name:         "invalid **time.Time",
			value:        (**time.Time)(nil),
			expectedTime: false,
		},
		{
			name:         "not a time",
			value:        time.Duration(0),
			expectedTime: false,
		},
	})
}

func ptr[T any](value T) *T {
	v := value
	return &v
}
