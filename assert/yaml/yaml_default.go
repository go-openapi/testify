// Package yaml is just an indirection to handle YAML deserialization.
//
// This package is just an indirection that allows the builder to override the
// indirection with an alternative implementation of this package that uses
// another implementation of YAML deserialization. This allows to not either not
// use YAML deserialization at all, or to use another implementation than
// [gopkg.in/yaml.v3] (for example for license compatibility reasons, see [PR #1120]).
package yaml

var EnableYAMLUnmarshal func([]byte, any) error

// Unmarshal is just a wrapper of [gopkg.in/yaml.v3.Unmarshal].
func Unmarshal(in []byte, out any) error {
	if EnableYAMLUnmarshal == nil {
		panic(`
YAML is not enabled yet!

You should enable a YAML library before running this test,
e.g. by adding the following to your imports:

import (
			_ "github.com/go-openapi/testify/enable/yaml"
)
`,
		)

	}
	return EnableYAMLUnmarshal(in, out)
}
