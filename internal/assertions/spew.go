// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "github.com/go-openapi/testify/v2/internal/spew"

func dumper(a ...any) string {
	const spewMaxDepth = 10
	spewConfig := spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
		SpewKeys:                true,
		DisableMethods:          true,
		EnableTimeStringer:      true,
		MaxDepth:                spewMaxDepth,
	}

	return spewConfig.Sdump(a...)
}
