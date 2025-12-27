// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package generator

import "golang.org/x/tools/imports"

// GenerateOption is an option for code generation.
type GenerateOption func(*generateOptions)

// Option is a general option for the generator.
//
// Currently unused (reserved for future use).
type Option func(*options)

func WithTargetRoot(dir string) GenerateOption {
	return func(o *generateOptions) {
		o.targetRoot = dir
	}
}

func WithTargetPackage(pkg string) GenerateOption {
	return func(o *generateOptions) {
		o.targetPkg = pkg
	}
}

func WithTargetDoc(dir string) GenerateOption {
	return func(o *generateOptions) {
		o.targetDoc = dir
	}
}

func WithIncludeFormatFuncs(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.enableFormat = enabled
	}
}

func WithIncludeForwardFuncs(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.enableForward = enabled
	}
}

func WithIncludeTests(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.generateTests = enabled
	}
}

func WithIncludeGenerics(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.enableGenerics = enabled
	}
}

func WithIncludeHelpers(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.generateHelpers = enabled
	}
}

func WithIncludeExamples(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.generateExamples = enabled
	}
}

func WithRunnableExamples(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.runnableExamples = enabled
	}
}

func WithIncludeDoc(enabled bool) GenerateOption {
	return func(o *generateOptions) {
		o.generateDoc = enabled
	}
}

type generateOptions struct {
	targetPkg        string
	targetRoot       string
	targetDoc        string
	enableForward    bool
	enableFormat     bool
	enableGenerics   bool
	generateHelpers  bool
	generateTests    bool
	generateExamples bool
	runnableExamples bool
	generateDoc      bool
	formatOptions    *imports.Options
}

type options struct{}

func generateOptionsWithDefaults(opts []GenerateOption) generateOptions {
	o := generateOptions{
		targetRoot:       ".",
		targetDoc:        ".",
		enableForward:    true,
		enableFormat:     true,
		enableGenerics:   false,
		generateHelpers:  true,
		generateTests:    false,
		generateExamples: false,
		generateDoc:      false,
		runnableExamples: false,
	}

	for _, apply := range opts {
		apply(&o)
	}

	return o
}

func optionsWithDefaults(_ []Option) options {
	return options{}
}
