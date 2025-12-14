// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "github.com/go-openapi/testify/v2/internal/spew"

const spewMaxDepth = 10

//nolint:gochecknoglobals // spew is more easily configured using a global default config. This is okay in this context.
var (
	spewConfig = spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
		DisableMethods:          true,
		MaxDepth:                spewMaxDepth,
	}

	spewConfigStringerEnabled = spew.ConfigState{
		Indent:                  " ",
		DisablePointerAddresses: true,
		DisableCapacities:       true,
		SortKeys:                true,
		MaxDepth:                spewMaxDepth,
	}
)
