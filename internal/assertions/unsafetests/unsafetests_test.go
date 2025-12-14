package unsafetests_test

import (
	"fmt"
	"testing"
	"unsafe"

	assert "github.com/go-openapi/testify/v2/internal/assertions"
)

// safeguard: ignoreT implements [assert.T]
var _ assert.T = ignoreT{}

func TestEqualUnsafePointers(t *testing.T) {
	var ignore ignoreT

	assert.True(t, assert.Nil(t, unsafe.Pointer(nil), "unsafe.Pointer(nil) is nil"))
	assert.False(t, assert.NotNil(ignore, unsafe.Pointer(nil), "unsafe.Pointer(nil) is nil"))

	assert.True(t, assert.Nil(t, unsafe.Pointer((*int)(nil)), "unsafe.Pointer((*int)(nil)) is nil"))
	assert.False(t, assert.NotNil(ignore, unsafe.Pointer((*int)(nil)), "unsafe.Pointer((*int)(nil)) is nil"))

	assert.False(t, assert.Nil(ignore, unsafe.Pointer(new(int)), "unsafe.Pointer(new(int)) is NOT nil"))
	assert.True(t, assert.NotNil(t, unsafe.Pointer(new(int)), "unsafe.Pointer(new(int)) is NOT nil"))
}

type ignoreT struct{}

func (ignoreT) Helper() {}

func (ignoreT) Errorf(format string, args ...any) {
	// Run the formatting, but ignore the result
	msg := fmt.Sprintf(format, args...)
	_ = msg
}
