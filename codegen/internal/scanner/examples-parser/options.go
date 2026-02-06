// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package parser

// Option configures the [Extractor].
type Option func(*Extractor)

// WithWorkDir sets the working directory for package resolution.
func WithWorkDir(dir string) Option {
	return func(e *Extractor) {
		e.dir = dir
	}
}

// WithBuildTags sets build tags for package loading (e.g. "integrationtest").
func WithBuildTags(tags ...string) Option {
	return func(e *Extractor) {
		e.buildTags = tags
	}
}
