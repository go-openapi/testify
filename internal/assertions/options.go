// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "slices"

// Option to modify the behavior of an assertion.
//
// Options are specified when building an assertion type with [New].
//
// At this moment, the only supported option is to modify the context-width ("hunk") of the diff
// reported on the family of Equal assertions.
type Option func(o options) options

type options struct {
	hunkSize int
}

// BuildOptions compound multiple options.
//
// The result is an opaque "any" on purpose: options remain unexported.
func BuildOptions(opts []Option) any {
	// excluded: don't re-export this function in generated packages.
	var o options

	for _, apply := range opts {
		o = apply(o)
	}

	return o
}

// WithHunkSize sets how many unchanged lines the diff shows around each change.
//
// The diff appears in the failure message of [Equal], [EqualT], [EqualValues],
// [EqualExportedValues] and [Exactly], whenever both values are a struct, map, slice, array
// or string. The default is 1. A value below 1 is clamped to 1, and a value larger than the
// rendered value prints it whole.
//
// Pass it to [New], which is the only place options are read:
//
//	a := assert.New(t, assert.WithHunkSize(4))
//	a.Equal(expected, actual)
func WithHunkSize(n int) Option {
	return func(o options) options {
		o.hunkSize = n

		return o
	}
}

// splitArgs extracts options from variadic arguments to assertions.
//
// Forward methods ([Assertions]) append the value built by [BuildOptions] to msgAndArgs.
// splitArgs pulls it back out, so it never reaches the failure message.
//
// A later option overrides an earlier one. The returned msgAndArgs aliases args
// when there is no option to remove.
func splitArgs(args []any) (msgAndArgs []any, opts options) {
	msgAndArgs = args

	for i := 0; i < len(msgAndArgs); i++ {
		switch opt := msgAndArgs[i].(type) { // the switch remains, even though it is now reduced to one single case.
		case options:
			opts = opt
			msgAndArgs = slices.Delete(slices.Clone(msgAndArgs), i, i+1)
			i-- // the deletion shifted the remaining arguments down

			// This case is left disabled for now: this would allow for injecting options directly on top-level functions
			// by hijacking args. This works, but we prefer for now to maintain this door closed.
			//
			// case Option:
			// opts = opt(opts)
			// msgAndArgs = slices.Delete(slices.Clone(msgAndArgs), i, i+1)
			// i--
		default:
			// noop
		}
	}

	return msgAndArgs, opts
}
