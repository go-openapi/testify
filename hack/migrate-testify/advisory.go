// SPDX-FileCopyrightText: Copyright 2026 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"go/ast"
	"go/token"
)

// eventuallyWithAdvisoryMessage is the advisory text emitted for every
// EventuallyWith / EventuallyWithT (pre-rename) call encountered in pass 1.
//
// It is informational: no source rewrite is performed. The semantics of
// CollectT.FailNow in go-openapi/testify/v2.4 match stretchr/testify
// (tick-only abort, poller retries), so code migrating from stretchr does
// not need to change. The advisory exists to surface the new Cancel()
// escape hatch and to flag callers for whom the old fork behavior
// (whole-assertion abort on FailNow) may have been load-bearing.
const eventuallyWithAdvisoryMessage = "advisory: EventuallyWith — in v2.4, CollectT.FailNow() aborts only the current tick (matches stretchr). " +
	"Use CollectT.Cancel() if you want to abort the whole assertion immediately. " +
	"See: https://go-openapi.github.io/testify/usage/migration/index.html#collectt-failnow-vs-cancel"

// eventuallyWithNames lists the EventuallyWith-family call names to watch.
// Names are checked in their pre-rename form (stretchr spelling) and the
// post-rename form, so the advisory fires regardless of migration order.
var eventuallyWithNames = map[string]bool{ //nolint:gochecknoglobals // lookup table
	"EventuallyWithT":  true,
	"EventuallyWithTf": true,
	"EventuallyWith":   true,
	"EventuallyWithf":  true,
}

// noteEventuallyWithCancel emits an informational diagnostic for every
// EventuallyWith-family call found in f that targets a testify package.
//
// It does not modify the AST. The advisory is emitted once per call site.
//
//nolint:unparam // return count left for future use.
func noteEventuallyWithCancel(f *ast.File, aliases map[string]string, fset *token.FileSet, rpt *report, filename string) int {
	count := 0

	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, ok := call.Fun.(*ast.SelectorExpr)
		if !ok {
			return true
		}

		if !eventuallyWithNames[sel.Sel.Name] {
			return true
		}

		if !isTestifySelector(sel, aliases) {
			return true
		}

		pos := fset.Position(call.Pos())
		rpt.warn(filename, pos.Line, fmt.Sprintf("%s(): %s", sel.Sel.Name, eventuallyWithAdvisoryMessage))
		count++

		return true
	})

	return count
}
