// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"fmt"
)

const maxMessageSize = 1024

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding the
// bufio.MaxScanTokenSize max line length that the go testing framework imposes.
func truncatingFormat(format string, data any) string {
	value := fmt.Sprintf(format, data)
	// Give us space for two truncated objects and the surrounding sentence.
	if len(value) > maxMessageSize {
		value = value[0:maxMessageSize] + "<... truncated>"
	}

	return value
}
