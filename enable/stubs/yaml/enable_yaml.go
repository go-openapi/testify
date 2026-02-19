// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

// Package yaml is an indirection to the internal implementation that handles YAML deserialization.
//
// This package allows the builder to override the indirection with an alternative implementation
// of YAML deserialization.
package yaml

import (
	yamlstub "github.com/go-openapi/testify/v2/internal/assertions/enable/yaml"
)

// EnableYAMLWithUnmarshal registers a YAML-capable unmarshaler.
//
// This is not intended for concurrent use.
//
// Most users would register using a init() function or enabling the
// registered library provided when importing "github.com/go-openapi/testify/enable/yaml/v2" like so.
//
// The default registration uses [go.yaml.in/yaml/v3] to deserialize YAML.
//
//	import(
//		_ "github.com/go-openapi/testify/enable/yaml/v2"
//	)
func EnableYAMLWithUnmarshal(unmarshaller func([]byte, any) error) {
	yamlstub.EnableYAMLWithUnmarshal(unmarshaller)
}

// EnableYAMLWithMarshal registers a YAML-capable marshaler.
//
// See [EnableYAMLWithUnmarshal].
func EnableYAMLWithMarshal(marshaller func(any) ([]byte, error)) {
	yamlstub.EnableYAMLWithMarshal(marshaller)
}
