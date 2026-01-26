// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

// Option to tune the [Generator].
type Option func(*options)

type options struct {
	skipCircularMap bool
}

// WithSkipCircularMap allows to skip specifically the self-referencing map scenario.
func WithSkipCircularMap(skipped bool) Option {
	return func(o *options) {
		o.skipCircularMap = skipped
	}
}

func optionsWithDefaults(opts []Option) options {
	var o options

	for _, apply := range opts {
		apply(&o)
	}

	return o
}
