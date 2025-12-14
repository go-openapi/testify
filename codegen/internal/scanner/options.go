// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

type Option func(*options)

type options struct {
	dir string
	pkg string
}

// WithWorkDir gives a workdir to the code scanner.
func WithWorkDir(dir string) Option {
	return func(o *options) {
		o.dir = dir
	}
}

// WithPackage indicates which package is scanned.
func WithPackage(pkg string) Option {
	return func(o *options) {
		o.pkg = pkg
	}
}

func optionsWithDefaults(opts []Option) options {
	o := options{
		dir: "..",
		pkg: "github.com/go-openapi/testify/v2/internal/assertions",
	}

	for _, apply := range opts {
		apply(&o)
	}

	return o
}
