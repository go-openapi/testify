// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"testing"
)

// These tests guard the reflection-based assertion helpers against cyclic
// inputs. Such inputs are legal Go values but would, without an explicit guard,
// drive the helpers into unbounded recursion and overflow the goroutine stack.
//
// A goroutine stack overflow is a fatal runtime error that cannot be recovered,
// so the regression signal is simply that these tests run to completion: before
// the fix, each case crashes the test binary.

type cyclicNode struct {
	Name string
	Next *cyclicNode
}

// TestCopyExportedFieldsCycle covers copyExportedFields (reached via
// EqualExportedValues) with self-referential and mutually-referential pointers.
func TestCopyExportedFieldsCycle(t *testing.T) {
	t.Parallel()

	t.Run("self-referential struct does not overflow the stack", func(t *testing.T) {
		t.Parallel()

		a := &cyclicNode{Name: "node"}
		a.Next = a
		b := &cyclicNode{Name: "node"}
		b.Next = b

		mock := new(mockT)
		// Two structurally identical self-cycles compare as equal.
		if !EqualExportedValues(mock, a, b) {
			t.Errorf("expected structurally identical self-cycles to be equal")
		}
	})

	t.Run("mutually-referential structs do not overflow the stack", func(t *testing.T) {
		t.Parallel()

		a1 := &cyclicNode{Name: "a"}
		a2 := &cyclicNode{Name: "b"}
		a1.Next = a2
		a2.Next = a1

		b1 := &cyclicNode{Name: "a"}
		b2 := &cyclicNode{Name: "b"}
		b1.Next = b2
		b2.Next = b1

		mock := new(mockT)
		if !EqualExportedValues(mock, a1, b1) {
			t.Errorf("expected structurally identical mutual cycles to be equal")
		}
	})

	t.Run("copyExportedFields returns on a direct cycle", func(t *testing.T) {
		t.Parallel()

		a := &cyclicNode{Name: "node"}
		a.Next = a

		// Direct call: the only property under test is that it returns.
		_ = copyExportedFields(a)
	})
}

// selfPtr is a pointer type that can point to a value of its own type, allowing
// the construction of a pointer chain that cycles back on itself.
type selfPtr *selfPtr

// TestIsEmptyValueCycle covers isEmptyValue (reached via Empty/NotEmpty/Zero)
// with a cyclic pointer chain.
func TestIsEmptyValueCycle(t *testing.T) {
	t.Parallel()

	var p selfPtr
	p = &p // p points to itself

	mock := new(mockT)

	// A non-nil, cyclic pointer chain is not empty; the assertions must agree
	// and, above all, must not recurse forever.
	empty := Empty(mock, p)
	notEmpty := NotEmpty(new(mockT), p)

	if empty {
		t.Errorf("expected a non-nil cyclic pointer to be reported as not empty")
	}
	if empty == notEmpty {
		t.Errorf("Empty and NotEmpty must be complementary, got Empty=%t NotEmpty=%t", empty, notEmpty)
	}
}

// cyclicError is an error whose Unwrap chain can be made to cycle.
type cyclicError struct {
	msg  string
	next error
}

func (e *cyclicError) Error() string { return e.msg }
func (e *cyclicError) Unwrap() error { return e.next }

// multiCyclicError exercises the Unwrap() []error branch of the chain walker.
type multiCyclicError struct {
	msg  string
	errs []error
}

func (e *multiCyclicError) Error() string   { return e.msg }
func (e *multiCyclicError) Unwrap() []error { return e.errs }

// TestUnwrapAllCycle covers unwrapAll (reached when formatting error chains in
// ErrorIs/NotErrorIs failure messages) with cyclic Unwrap chains.
//
// Note: we deliberately test unwrapAll directly rather than through the public
// ErrorIs/NotErrorIs assertions. Those call the standard library's errors.Is
// first, which itself walks the Unwrap chain without cycle detection and so
// hangs on a cyclic chain before our formatter is ever reached. Guarding that
// would mean reimplementing errors.Is and diverging from standard-library
// semantics, which is out of scope here. unwrapAll is hardened on its own so
// that our code never contributes an unbounded recursion of its own.
func TestUnwrapAllCycle(t *testing.T) {
	t.Parallel()

	t.Run("single Unwrap cycle does not overflow the stack", func(t *testing.T) {
		t.Parallel()

		e := &cyclicError{msg: "boom"}
		e.next = e // self-cycle

		errs := unwrapAll(e)
		if len(errs) == 0 {
			t.Fatalf("expected at least the head error")
		}
	})

	t.Run("multi Unwrap cycle does not overflow the stack", func(t *testing.T) {
		t.Parallel()

		leaf := &cyclicError{msg: "leaf"}
		root := &multiCyclicError{msg: "root"}
		root.errs = []error{leaf, root} // root references itself

		errs := unwrapAll(root)
		if len(errs) == 0 {
			t.Fatalf("expected at least the head error")
		}
	})
}
