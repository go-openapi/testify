package assert

import (
	"errors"
)

// ErrTest is an error instance useful for testing.
//
// If the code does not care about error specifics, and only needs
// to return the error for example, this error should be used to make
// the test code more readable.
var ErrTest = errors.New("assert.ErrTest general error for testing")
