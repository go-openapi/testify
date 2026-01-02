// Package yaml is an indirection to handle YAML deserialization.
//
// This package allows the builder to override the indirection with an alternative implementation
// of YAML deserialization.
package yaml

var enableYAMLUnmarshal func([]byte, any) error //nolint:gochecknoglobals // in this particular case, we need a global to enable the feature from another module

// EnableYAMLWithUnmarshal registers a YAML-capable unmarshaler.
//
// This is not intended for concurrent use.
func EnableYAMLWithUnmarshal(unmarshaler func([]byte, any) error) {
	enableYAMLUnmarshal = unmarshaler
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
