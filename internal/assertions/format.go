// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

// ==========================================
// Utility helpers to format error messages.
// ==========================================

const maxMessageSize = 1024

// truncatingFormat formats the data and truncates it if it's too long.
//
// This helps keep formatted error messages lines from exceeding maxMessageSize for readability's sake.
func truncatingFormat(format string, data any) string {
	value := fmt.Sprintf(format, data)
	// Give us space for two truncated objects and the surrounding sentence.
	if len(value) > maxMessageSize {
		value = value[0:maxMessageSize] + "<... truncated>"
	}

	return value
}

// Aligns the provided message so that all lines after the first line start at the same location as the first line.
//
// Assumes that the first line starts at the correct location (after carriage return, tab, label, spacer and tab).
//
// The longestLabelLen parameter specifies the length of the longest label in the output (required because this is the
// basis on which the alignment occurs).
func indentMessageLines(message string, longestLabelLen int) string {
	outBuf := new(bytes.Buffer)

	scanner := bufio.NewScanner(strings.NewReader(message))
	for firstLine := true; scanner.Scan(); firstLine = false {
		if !firstLine {
			fmt.Fprint(outBuf, "\n\t"+strings.Repeat(" ", longestLabelLen+1)+"\t")
		}
		fmt.Fprint(outBuf, scanner.Text())
	}

	return outBuf.String()
}

func messageFromMsgAndArgs(msgAndArgs ...any) string {
	if len(msgAndArgs) == 0 || msgAndArgs == nil {
		return ""
	}

	if len(msgAndArgs) == 1 {
		msg := msgAndArgs[0]
		if msgAsStr, ok := msg.(string); ok {
			return msgAsStr
		}
		return fmt.Sprintf("%+v", msg)
	}

	if len(msgAndArgs) > 1 {
		format, ok := msgAndArgs[0].(string)
		if ok {
			return fmt.Sprintf(format, msgAndArgs[1:]...)
		}
	}

	return ""
}

type labeledContent struct {
	label   string
	content string
}

// labeledOutput returns a string consisting of the provided labeledContent. Each labeled output is appended in the following manner:
//
//	\t{{label}}:{{align_spaces}}\t{{content}}\n
//
// The initial carriage return is required to undo/erase any padding added by testing.T.Errorf. The "\t{{label}}:" is for the label.
// If a label is shorter than the longest label provided, padding spaces are added to make all the labels match in length. Once this
// alignment is achieved, "\t{{content}}\n" is added for the output.
//
// If the content of the labeledOutput contains line breaks, the subsequent lines are aligned so that they start at the same location as the first line.
func labeledOutput(content ...labeledContent) string {
	longestLabel := 0
	for _, v := range content {
		if len(v.label) > longestLabel {
			longestLabel = len(v.label)
		}
	}
	var output strings.Builder
	for _, v := range content {
		output.WriteString("\t" + v.label + ":" + strings.Repeat(" ", longestLabel-len(v.label)) + "\t" + indentMessageLines(v.content, longestLabel) + "\n")
	}

	return output.String()
}
