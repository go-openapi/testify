package spew

import (
	"io"
	"iter"
	"slices"
	"sync"
	"testing"
)

func TestEdgeCases(t *testing.T) {
	t.Parallel()
	cfg := Config
	output := io.Discard

	for tt := range edgeCases() {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if r := recover(); r != nil {
					t.Errorf("fdump panicked: %v\nWith value type: %T\nValue: %#v", r, tt.value, tt.value)

					return
				}
			}()

			fdump(&cfg, output, tt.value)
		})
	}
}

type edgeCase struct {
	name  string
	value any
}

func edgeCases() iter.Seq[edgeCase] {
	type withLock struct {
		sync.Mutex
	}
	type withLockPtr struct {
		*sync.Mutex
	}
	var iface any = "x"
	str := "y"
	var ifaceToPtr any = &str
	var ifaceToNilPtr any = (*string)(nil)
	ifacePtr := &iface
	ifacePtrPtr := &ifaceToPtr
	var ifaceCircular any
	ifaceCircular = &ifaceCircular

	// Map with circular value (map contains itself)
	m := map[string]any{
		"key1": 1,
		"key2": "val",
	}
	m["circular"] = m // Map contains itself as a value

	return slices.Values([]edgeCase{
		{
			name:  "self-referencing map",
			value: m,
		},
		{
			name:  "sync.Mutex",
			value: withLock{},
		},
		{
			name: "*sync.Mutex",
			value: withLockPtr{
				Mutex: &sync.Mutex{},
			},
		},
		{
			name:  "pointer to interface",
			value: ifacePtr,
		},
		{
			name:  "pointer to interface to pointer",
			value: ifacePtrPtr,
		},
		{
			name:  "pointer to interface to pointer",
			value: &ifaceToNilPtr,
		},
		{
			name:  "pointer to pointer to interface to pointer",
			value: &ifaceToPtr,
		},
		{
			// case that used to hang
			name:  "pointer to interface with circular reference",
			value: &ifaceCircular,
		},
	})
}
