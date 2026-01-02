package spew

import (
	"iter"
	"reflect"
	"slices"
	"testing"
)

func TestPanicUnexportedFields(t *testing.T) {
	t.Parallel()

	for tt := range panicCases() {
		t.Run(tt.name, func(t *testing.T) {
			v1 := reflect.ValueOf(tt.value1)
			v2 := reflect.ValueOf(tt.value2)
			isLess := valueSortLess(v1, v2)

			if isLess {
				t.Error("expected an ordered set")
			}
		})
	}
}

type panicCase struct {
	name   string
	value1 any
	value2 any
}

func panicCases() iter.Seq[panicCase] {
	type structWithUnexportedMapWithArrayKey struct {
		m any
	}

	return slices.Values([]panicCase{
		{
			// from issue https://github.com/stretchr/testify/pull/1816
			name: "panic behavior on struct with array key and unexported field",
			value1: structWithUnexportedMapWithArrayKey{
				map[[1]byte]*struct{}{
					{1}: nil,
					{2}: nil,
				},
			},
			value2: structWithUnexportedMapWithArrayKey{
				map[[1]byte]*struct{}{
					{2}: nil,
					{1}: nil,
				},
			},
		},
	})
}
