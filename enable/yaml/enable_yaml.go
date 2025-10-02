// Package yaml enables the YAMLEq capability in testify.
package yaml

import (
	yamlstub "github.com/go-openapi/testify/v2/assert/yaml"

	yaml "go.yaml.in/yaml/v3"
)

func init() {
	yamlstub.EnableYAMLUnmarshal = yaml.Unmarshal
}
