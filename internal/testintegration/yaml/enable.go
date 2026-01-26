// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package yaml

import (
	yamlstub "github.com/go-openapi/testify/v2/enable/stubs/yaml"
	yaml "go.yaml.in/yaml/v3"
)

func enableYAML() {
	yamlstub.EnableYAMLWithUnmarshal(yaml.Unmarshal)
}
