// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package yaml is an indirection to handle YAML serialization/deserialization.
//
// This package allows the builder to override the indirection with an alternative implementation
// of YAML serialization.
package yaml

//nolint:gochecknoglobals // in this particular case, we need a global to enable the feature from another module
var (
	enableYAMLUnmarshal func([]byte, any) error
	enableYAMLMarshal   func(any) ([]byte, error)
)

// EnableYAMLWithUnmarshal registers a YAML-capable unmarshaler.
//
// This is not intended for concurrent use.
func EnableYAMLWithUnmarshal(unmarshaler func([]byte, any) error) {
	enableYAMLUnmarshal = unmarshaler
}

func EnableYAMLWithMarshal(marshaler func(any) ([]byte, error)) {
	enableYAMLMarshal = marshaler
}

// Unmarshal is a wrapper to some exernal library to unmarshal YAML documents.
func Unmarshal(in []byte, out any) error {
	if enableYAMLUnmarshal == nil {
		// fail early and loud
		panic(`
YAML is not enabled yet!

You should enable a YAML library before running this test,
e.g. by adding the following to your imports:

import (
			_ "github.com/go-openapi/testify/enable/yaml/v2"
)
`,
		)
	}
	return enableYAMLUnmarshal(in, out)
}

// Marshal is a wrapper to some exernal library to marshal YAML documents.
func Marshal(in any) ([]byte, error) {
	if enableYAMLMarshal == nil {
		// fail early and loud
		panic(`
YAML is not enabled yet!

You should enable a YAML library before running this test,
e.g. by adding the following to your imports:

import (
			_ "github.com/go-openapi/testify/enable/yaml/v2"
)
`,
		)
	}
	return enableYAMLMarshal(in)
}
