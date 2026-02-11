// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"strings"

	"golang.org/x/tools/go/packages"
)

// minPairArgs is the minimum number of assertion arguments (after t) needed
// for binary assertions like Equal(t, expected, actual).
const minPairArgs = 2

// runGenericUpgrade executes pass 2: reflection → generic assertion upgrade.
func runGenericUpgrade(dir string, opts *options) error {
	pkgs, fset, err := loadPackages(dir)
	if err != nil {
		return err
	}

	rpt := &report{}

	for _, pkg := range pkgs {
		if pkg.TypesInfo == nil {
			continue
		}

		for i, f := range pkg.Syntax {
			if i >= len(pkg.GoFiles) {
				continue
			}
			filename := pkg.GoFiles[i]

			if opts.skipVendor && isVendorPath(filename) {
				continue
			}

			if !fileImportsAny(f, "github.com/go-openapi/testify") {
				continue
			}

			rpt.filesScanned++
			changes := upgradeFile(f, pkg, fset, rpt, filename, opts.verbose)
			if changes == 0 {
				continue
			}

			rpt.filesChanged++
			rpt.totalChanges += changes

			if opts.dryRun {
				if err := showDiff(fset, f, filename); err != nil {
					rpt.errorf(filename, 0, err.Error())
				}
			} else {
				if err := writeFile(fset, f, filename); err != nil {
					rpt.errorf(filename, 0, err.Error())
				}
			}
		}
	}

	rpt.print(opts.verbose)
	rpt.printPass2Summary()
	return nil
}

// upgradeFile processes a single file for generic upgrades.
func upgradeFile(f *ast.File, pkg *packages.Package, fset *token.FileSet, rpt *report, filename string, verbose bool) int {
	aliases := buildGoOpenapiAliasMap(f)
	if len(aliases) == 0 {
		return 0
	}

	changes := 0

	ast.Inspect(f, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}

		sel, funcName, isTestifyCall := extractTestifyCall(call, aliases)
		if !isTestifyCall {
			return true
		}

		// Strip trailing "f" for format variants.
		baseName := funcName
		isFormat := false
		if strings.HasSuffix(baseName, "f") && baseName != "Equalf" {
			// Check if baseName without 'f' is in the upgrade table.
			candidate := baseName[:len(baseName)-1]
			if _, ok := genericUpgrades[candidate]; ok {
				baseName = candidate
				isFormat = true
			}
		}
		// Also check Equalf directly.
		if baseName == "Equalf" {
			baseName = "Equal"
			isFormat = true
		}

		rule, ok := genericUpgrades[baseName]
		if !ok {
			return true
		}

		if rule.manualReview {
			pos := fset.Position(call.Pos())
			rpt.warn(filename, pos.Line,
				fmt.Sprintf("%s → %s requires manual review (argument count changes)", funcName, rule.target))
			return true
		}

		if rule.containerUpgrade {
			if upgraded := tryContainerUpgrade(call, sel, funcName, baseName, isFormat, pkg, fset, rpt, filename, verbose); upgraded {
				changes++
			}
			return true
		}

		if upgraded := trySimpleUpgrade(call, sel, funcName, baseName, isFormat, rule, pkg, fset, rpt, filename, verbose); upgraded {
			changes++
		}

		return true
	})

	return changes
}

// buildGoOpenapiAliasMap builds an alias map for go-openapi/testify imports.
func buildGoOpenapiAliasMap(f *ast.File) map[string]string {
	aliases := make(map[string]string)
	for _, imp := range f.Imports {
		path := importPath(imp)
		if !strings.HasPrefix(path, "github.com/go-openapi/testify") {
			continue
		}
		// Skip enable imports.
		if strings.Contains(path, "/enable/") {
			continue
		}

		var localName string
		if imp.Name != nil {
			localName = imp.Name.Name
		} else {
			parts := strings.Split(path, "/")
			localName = parts[len(parts)-1]
		}
		if localName != "_" && localName != "." {
			aliases[localName] = path
		}
	}
	return aliases
}

// extractTestifyCall checks if a call expression is a testify assertion call.
// Returns the selector, function name, and whether it's a testify call.
func extractTestifyCall(call *ast.CallExpr, aliases map[string]string) (*ast.SelectorExpr, string, bool) {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return nil, "", false
	}

	funcName := sel.Sel.Name

	// Package-level call: assert.Equal(t, ...)
	if ident, ok := sel.X.(*ast.Ident); ok {
		if _, exists := aliases[ident.Name]; exists {
			return sel, funcName, true
		}
	}

	return nil, "", false
}

// checkResult is the outcome of a constraint check for a generic upgrade.
type checkResult struct {
	ok       bool       // constraints satisfied — upgrade is safe
	reason   skipReason // non-empty: skip with this reason
	typeInfo string     // context for skip message (e.g. type name)
}

var (
	checkOK           = checkResult{ok: true} //nolint:gochecknoglobals // convenience value
	checkInsufficient = checkResult{}         //nolint:gochecknoglobals // not enough args, silent skip
)

func checkSkip(r skipReason, info string) checkResult {
	return checkResult{reason: r, typeInfo: info}
}

// trySimpleUpgrade attempts to upgrade a simple (non-container) assertion.
func trySimpleUpgrade(
	call *ast.CallExpr,
	sel *ast.SelectorExpr,
	funcName, baseName string,
	isFormat bool,
	rule upgradeRule,
	pkg *packages.Package,
	fset *token.FileSet,
	rpt *report,
	filename string,
	verbose bool,
) bool {
	pos := fset.Position(call.Pos())

	skip := func(reason skipReason, typeInfo string) bool {
		rpt.trackSkip(filename, pos.Line, funcName, reason, verbose, typeInfo)
		return false
	}

	argTypes, argSkipReason := getArgTypesWithReason(call, pkg, 1)
	if argTypes == nil {
		return skip(argSkipReason, "")
	}

	var result checkResult
	switch baseName {
	case "Equal", "NotEqual":
		result = checkDeepComparablePair(argTypes, rule)
	case "Greater", "GreaterOrEqual", "Less", "LessOrEqual":
		result = checkPairConstraint(argTypes, constraintOrdered, skipNotOrdered, rule)
	case "InDelta", "InEpsilon":
		result = checkPairConstraint(argTypes, constraintMeasurable, skipNotMeasurable, rule)
	case "Positive", "Negative":
		result = checkSingleConstraint(argTypes, constraintSignedNumeric, skipNotSignedNumeric)
	case "True", "False":
		result = checkSingleConstraint(argTypes, constraintBoolean, skipNotBoolean)
	case "Same", "NotSame":
		result = checkPairConstraint(argTypes, constraintPointer, skipNotPointer, rule)
	case "ElementsMatch", "Subset":
		result = checkSlicePairConstraint(argTypes, rule)
	case "IsIncreasing", "IsDecreasing", "IsNonIncreasing", "IsNonDecreasing":
		result = checkOrderedSliceConstraint(argTypes)
	case "Regexp", "NotRegexp":
		result = checkRegexpConstraint(argTypes)
	case "JSONEq", "YAMLEq":
		result = checkPairConstraint(argTypes, constraintText, skipNotText, rule)
	default:
		return false
	}

	if !result.ok {
		if result.reason != "" {
			return skip(result.reason, result.typeInfo)
		}
		return false
	}

	newName := rule.target
	if isFormat {
		newName += "f"
	}

	if verbose {
		rpt.info(filename, pos.Line, fmt.Sprintf("upgraded %s → %s", funcName, newName))
	}
	rpt.trackUpgrade(funcName, newName)

	sel.Sel.Name = newName
	return true
}

// checkDeepComparablePair checks that both arguments are deeply comparable and have matching types.
func checkDeepComparablePair(argTypes []types.Type, rule upgradeRule) checkResult {
	if len(argTypes) < minPairArgs {
		return checkInsufficient
	}
	if reason := deepComparableSkipReason(argTypes[0]); reason != "" {
		return checkSkip(reason, argTypes[0].String())
	}
	if reason := deepComparableSkipReason(argTypes[1]); reason != "" {
		return checkSkip(reason, argTypes[1].String())
	}
	if rule.sameType && !sameType(argTypes[0], argTypes[1]) {
		return checkSkip(skipTypeMismatch, argTypes[0].String()+" vs "+argTypes[1].String())
	}
	return checkOK
}

// checkPairConstraint checks that both arguments satisfy a constraint and have matching types.
func checkPairConstraint(argTypes []types.Type, c constraintKind, failReason skipReason, rule upgradeRule) checkResult {
	if len(argTypes) < minPairArgs {
		return checkInsufficient
	}
	if !satisfiesConstraint(argTypes[0], c) || !satisfiesConstraint(argTypes[1], c) {
		return checkSkip(failReason, argTypes[0].String())
	}
	if rule.sameType && !sameType(argTypes[0], argTypes[1]) {
		return checkSkip(skipTypeMismatch, argTypes[0].String()+" vs "+argTypes[1].String())
	}
	return checkOK
}

// checkSingleConstraint checks that the first argument satisfies a constraint.
func checkSingleConstraint(argTypes []types.Type, c constraintKind, failReason skipReason) checkResult {
	if len(argTypes) < 1 {
		return checkInsufficient
	}
	if !satisfiesConstraint(argTypes[0], c) {
		return checkSkip(failReason, argTypes[0].String())
	}
	return checkOK
}

// checkSlicePairConstraint checks that both arguments are slices with deep-comparable elements.
func checkSlicePairConstraint(argTypes []types.Type, rule upgradeRule) checkResult {
	if len(argTypes) < minPairArgs {
		return checkInsufficient
	}
	elem0, ok0 := isSliceType(argTypes[0])
	elem1, ok1 := isSliceType(argTypes[1])
	if !ok0 || !ok1 {
		return checkSkip(skipNotSlice, argTypes[0].String())
	}
	if !isDeepComparable(elem0) || !isDeepComparable(elem1) {
		return checkSkip(skipSliceElemNotDeepComparable, elem0.String())
	}
	if rule.sameType && !sameType(argTypes[0], argTypes[1]) {
		return checkSkip(skipTypeMismatch, argTypes[0].String()+" vs "+argTypes[1].String())
	}
	return checkOK
}

// checkOrderedSliceConstraint checks that the argument is a slice of ordered elements.
func checkOrderedSliceConstraint(argTypes []types.Type) checkResult {
	if len(argTypes) < 1 {
		return checkInsufficient
	}
	elem, ok := isSliceType(argTypes[0])
	if !ok {
		return checkSkip(skipNotSlice, argTypes[0].String())
	}
	if !isOrdered(elem) {
		return checkSkip(skipSliceElemNotOrdered, elem.String())
	}
	return checkOK
}

// checkRegexpConstraint checks that the first arg is a RegExp and the second is Text.
func checkRegexpConstraint(argTypes []types.Type) checkResult {
	if len(argTypes) < minPairArgs {
		return checkInsufficient
	}
	if !satisfiesConstraint(argTypes[0], constraintRegExp) {
		return checkSkip(skipNotRegExp, argTypes[0].String())
	}
	if !satisfiesConstraint(argTypes[1], constraintText) {
		return checkSkip(skipNotText, argTypes[1].String())
	}
	return checkOK
}

// deepComparableSkipReason returns the specific skip reason if a type is not deep-comparable,
// or an empty string if it is.
func deepComparableSkipReason(typ types.Type) skipReason {
	if !types.Comparable(typ) {
		return skipNotComparable
	}
	under := typ.Underlying()
	switch under.(type) {
	case *types.Pointer:
		return skipPointerSemantics
	case *types.Interface:
		return skipInterfaceField
	case *types.Struct:
		if !isDeepComparable(typ) {
			return skipInterfaceField
		}
	}
	if !isDeepComparable(typ) {
		return skipNotComparable
	}
	return ""
}

// containerCheckResult extends checkResult with the resolved target function name.
type containerCheckResult struct {
	checkResult

	target string
}

// tryContainerUpgrade handles Contains/NotContains which dispatch to different
// generic variants based on the container type.
func tryContainerUpgrade(
	call *ast.CallExpr,
	sel *ast.SelectorExpr,
	funcName, baseName string,
	isFormat bool,
	pkg *packages.Package,
	fset *token.FileSet,
	rpt *report,
	filename string,
	verbose bool,
) bool {
	pos := fset.Position(call.Pos())

	skip := func(reason skipReason, typeInfo string) bool {
		rpt.trackSkip(filename, pos.Line, funcName, reason, verbose, typeInfo)
		return false
	}

	argTypes, argSkipReason := getArgTypesWithReason(call, pkg, 1)
	if argTypes == nil {
		return skip(argSkipReason, "")
	}
	if len(argTypes) < minPairArgs {
		return false
	}

	isNot := baseName == "NotContains"

	var result containerCheckResult
	switch {
	case isText(argTypes[0]):
		result = checkStringContains(argTypes, isNot)
	case isSliceType2(argTypes[0]):
		result = checkSliceContains(argTypes, isNot)
	case isMapType2(argTypes[0]):
		result = checkMapContains(argTypes, isNot)
	default:
		return skip(skipContainerTypeUnknown, argTypes[0].String())
	}

	if !result.ok {
		if result.reason != "" {
			return skip(result.reason, result.typeInfo)
		}
		return false
	}

	target := result.target
	if isFormat {
		target += "f"
	}

	if verbose {
		rpt.info(filename, pos.Line, fmt.Sprintf("upgraded %s → %s", funcName, target))
	}
	rpt.trackUpgrade(funcName, target)

	sel.Sel.Name = target
	return true
}

// isSliceType2 reports whether typ is a slice type (bool-only variant for switch).
func isSliceType2(typ types.Type) bool {
	_, ok := typ.Underlying().(*types.Slice)
	return ok
}

// isMapType2 reports whether typ is a map type (bool-only variant for switch).
func isMapType2(typ types.Type) bool {
	_, ok := typ.Underlying().(*types.Map)
	return ok
}

// pickTarget selects the contains or not-contains variant.
func pickTarget(kind containerKind, isNot bool) string {
	targets := containerUpgradeTargets[kind]
	if isNot {
		return targets[1]
	}
	return targets[0]
}

// checkStringContains validates a string Contains/NotContains upgrade.
func checkStringContains(argTypes []types.Type, isNot bool) containerCheckResult {
	if !isText(argTypes[1]) {
		return containerCheckResult{checkResult: checkSkip(skipNotText, argTypes[1].String())}
	}
	return containerCheckResult{checkResult: checkOK, target: pickTarget(containerString, isNot)}
}

// checkSliceContains validates a slice Contains/NotContains upgrade.
func checkSliceContains(argTypes []types.Type, isNot bool) containerCheckResult {
	elem, _ := isSliceType(argTypes[0])
	if !isDeepComparable(elem) {
		return containerCheckResult{checkResult: checkSkip(skipSliceElemNotDeepComparable, elem.String())}
	}
	if isAnyOrInterface(argTypes[1]) {
		return containerCheckResult{checkResult: checkSkip(skipInterfaceType, argTypes[1].String())}
	}
	if !sameType(elem, argTypes[1]) {
		return containerCheckResult{checkResult: checkSkip(skipTypeMismatch, elem.String()+" vs "+argTypes[1].String())}
	}
	return containerCheckResult{checkResult: checkOK, target: pickTarget(containerSlice, isNot)}
}

// checkMapContains validates a map Contains/NotContains upgrade.
func checkMapContains(argTypes []types.Type, isNot bool) containerCheckResult {
	key, _, _ := isMapType(argTypes[0])
	if !isComparable(key) {
		return containerCheckResult{checkResult: checkSkip(skipNotComparable, key.String())}
	}
	if isAnyOrInterface(argTypes[1]) {
		return containerCheckResult{checkResult: checkSkip(skipInterfaceType, argTypes[1].String())}
	}
	if !sameType(key, argTypes[1]) {
		return containerCheckResult{checkResult: checkSkip(skipTypeMismatch, key.String()+" vs "+argTypes[1].String())}
	}
	return containerCheckResult{checkResult: checkOK, target: pickTarget(containerMap, isNot)}
}

// getArgTypesWithReason extracts the types of call arguments starting from the given offset
// (to skip the testing.T parameter). Returns nil and a skip reason if any type cannot be resolved.
func getArgTypesWithReason(call *ast.CallExpr, pkg *packages.Package, offset int) ([]types.Type, skipReason) {
	if len(call.Args) <= offset {
		return nil, skipUnresolvableType
	}

	result := make([]types.Type, 0, len(call.Args)-offset)
	for _, arg := range call.Args[offset:] {
		tv, ok := pkg.TypesInfo.Types[arg]
		if !ok {
			return nil, skipUnresolvableType
		}
		if isAnyOrInterface(tv.Type) {
			return nil, skipInterfaceType
		}
		result = append(result, tv.Type)
	}
	return result, ""
}
