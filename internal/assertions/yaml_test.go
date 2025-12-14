// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

func TestYAMLEq(t *testing.T) {
	mock := new(mockT)
	const yamlDoc = `
---
a: 1
`

	if !Panics(t, func() {
		_ = YAMLEq(mock, "", yamlDoc)
	}) {
		Fail(t, "expected YAMLEq to panic with default settings")
	}
}
