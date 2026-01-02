// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package funcmaps

import (
	"fmt"
	"go/token"
	"net/url"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
	parser "github.com/go-openapi/testify/v2/codegen/internal/scanner/comments-parser"
	"github.com/go-openapi/testify/v2/internal/spew"
)

const (
	pkgRequire = "require"
	pkgAssert  = "assert"
	assertions = "assertions"
)

// Titleize transform a string into a title, upper-casing the first letter.
func Titleize(in string) string {
	return titleize(in)
}

// PrintReturns print a collection of [model.Parameters] as values returned by a function.
func PrintReturns(vars model.Parameters) string {
	return printReturns(vars)
}

// FuncMap returns the complete template function map for code and documentation generation.
func FuncMap() template.FuncMap {
	return map[string]any{
		"comment":          comment,
		"concat":           concatStrings,
		"date":             printDate,
		"debug":            printDebug,
		"docStringFor":     docStringFor,
		"docStringPackage": docStringPackage,
		"forward":          forward,
		"godocbadge":       godocbadge,
		"hasSuffix":        strings.HasSuffix,
		"imports":          printImports,
		"mdformat":         FormatMarkdown, // From markdown.go
		"params":           params,
		"pathparts":        pathParts,
		"quote":            quote,
		"relocate":         relocate,
		"returns":          PrintReturns,
		"sourceLink":       sourceLink,
		"titleize":         titleize,
	}
}

func printImports(in model.ImportMap) string {
	list := make([]string, 0, len(in))

	for k, v := range in {
		if k == path.Base(v) {
			// no alias
			list = append(list, "\t\""+v+"\"")

			continue
		}

		list = append(list, "\t"+k+"\t\""+v+"\"")
	}

	sort.Strings(list)

	return strings.Join(list, "\n")
}

func comment(str string) string {
	if str == "" {
		return ""
	}

	lines := rTrimEmpty(strings.Split(str, "\n"))

	return "// " + strings.Join(lines, "\n// ")
}

func rTrimEmpty(lines []string) []string {
	var i int
	for i = len(lines) - 1; i >= 0; i-- {
		if lines[i] != "" {
			break
		}
	}

	return lines[:i+1]
}

func params(args model.Parameters) string {
	var b strings.Builder
	l := len(args)

	if l == 0 {
		return ""
	}

	b.WriteString(args[0].Name)
	b.WriteByte(' ')
	if strings.HasSuffix(args[0].GoType, "Func") && args[0].Selector == assertions {
		b.WriteString(args[0].Selector + ".") // for xxxFunc types, backward-compatibility imposes to use the "assert-like" type definition
	}
	b.WriteString(args[0].GoType)

	for _, v := range args[1:] {
		b.WriteString(", ")
		b.WriteString(v.Name)
		b.WriteByte(' ')
		if strings.HasSuffix(v.GoType, "Func") && v.Selector == assertions {
			b.WriteString(v.Selector + ".") // for xxxFunc types, backward-compatibility imposes to use the "assert-like" type definition
		}
		b.WriteString(v.GoType)
	}

	return b.String()
}

func forward(args model.Parameters) string {
	var b strings.Builder
	l := len(args)

	if l == 0 {
		return ""
	}

	b.WriteString(args[0].Name)
	if args[0].IsVariadic {
		b.WriteString("...")
	}

	for _, v := range args[1:] {
		b.WriteString(", ")
		b.WriteString(v.Name)
		if v.IsVariadic {
			b.WriteString("...")
		}
	}

	return b.String()
}

func printReturns(vars model.Parameters) string {
	var b strings.Builder
	l := len(vars)

	if l == 0 {
		return ""
	}

	if l > 1 || vars[0].Name != "" {
		b.WriteByte('(')
	}

	if vars[0].Name != "" {
		b.WriteString(vars[0].Name)
		b.WriteByte(' ')
	}
	b.WriteString(vars[0].GoType)

	for _, v := range vars[1:] {
		b.WriteByte(',')
		b.WriteByte(' ')
		if v.Name != "" {
			b.WriteString(v.Name)
			b.WriteByte(' ')
		}
		b.WriteString(v.GoType)

	}

	if l > 1 || vars[0].Name != "" {
		b.WriteByte(')')
	}

	return b.String()
}

// docStringFor adds an extra comment with usage-specific context.
//
// Supported usage strings are:
//
// - format: usage with a format argument
// - forward: usage as a method.
func docStringFor(usage, name string) string {
	parts := strings.Split(name, ".")
	basename := parts[len(parts)-1]

	switch usage {
	case "format":
		return comment(
			fmt.Sprintf(
				"%sf is the same as [%s], but accepts a format msg string to format arguments like [fmt.Printf].",
				basename,
				name,
			),
		)
	case "forward":
		return comment(
			fmt.Sprintf(
				"%s is the same as [%s], as a method rather than a package-level function.",
				basename,
				name),
		)
	default:
		return ""
	}
}

// docStringPackage adds an additional comment specific to the target package.
//
// Supported pkg strings are:
//
// - assert
// - require.
func docStringPackage(pkg string) string {
	switch pkg {
	case pkgAssert:
		return `// Upon failure, the test [T] is marked as failed and continues execution.`
	case pkgRequire:
		return `// Upon failure, the test [T] is marked as failed and stops execution.`
	// NOTE(fredbi):
	// Proposal for enhancement: add more packages, e.g. for the generics-only API
	default:
		return ""
	}
}

func concatStrings(in ...string) string {
	return strings.Join(in, "")
}

func pathParts(in string) string {
	parts := strings.FieldsFunc(filepath.Clean(in), func(r rune) bool { return r == filepath.Separator })

	for i := range parts {
		parts[i] = fmt.Sprintf("%q", parts[i])
	}

	return strings.Join(parts, ",")
}

// relocate relocates already-parsed test values from the "assertions" package to the target package.
//
// It uses AST-based relocation to properly handle all Go expressions, replacing
// package selectors and qualifying unqualified identifiers as needed.
//
// Examples:
//   - assertions.ErrTest → assert.ErrTest
//   - ErrTest → assert.ErrTest (adds qualifier)
//   - assertions.PanicTestFunc → assertions.PanicTestFunc (exception)
func relocate(values []model.TestValue, pkg string) string {
	if len(values) == 0 {
		return ""
	}

	if pkg == "" {
		// Return original values as string
		parts := make([]string, 0, len(values))
		for _, v := range values {
			parts = append(parts, v.Raw)
		}
		return strings.Join(parts, ", ")
	}

	// Relocate each value
	relocated := make([]string, 0, len(values))
	for _, tv := range values {
		// If parse failed, use original (fallback)
		if tv.Error != nil {
			relocated = append(relocated, tv.Raw)
			continue
		}

		// Relocate from "assertions" to target package
		relocatedTV := parser.RelocateTestValue(tv, assertions, pkg)

		// If relocation failed, use original (fallback)
		if relocatedTV.Error != nil {
			relocated = append(relocated, tv.Raw)
		} else {
			relocated = append(relocated, relocatedTV.Raw)
		}
	}

	// Join back with comma-space
	return strings.Join(relocated, ", ")
}

func sourceLink(baseGitHubURL string, pos *token.Position) string {
	if pos == nil {
		return ""
	}

	dir, file := filepath.Split(pos.Filename)
	l1 := filepath.Base(dir)
	l2 := filepath.Base(filepath.Dir(filepath.Dir(dir)))
	filename := path.Join(l2, l1, file)
	return fmt.Sprintf("%s/blob/master/%s#L%d",
		baseGitHubURL,
		filename,
		pos.Line,
	)
}

func quote(in string) string {
	return fmt.Sprintf("%q", in)
}

func titleize(in string) string {
	caser := cases.Title(language.English, cases.NoLower) // the case is stateful: cannot declare it globally

	return caser.String(in)
}

func godocbadge(pkggodevURL string) (string, error) {
	u, err := url.Parse(pkggodevURL)
	if err != nil {
		return "", err
	}
	u.Path = path.Join("badge", u.Path)

	return u.String(), nil
}

func printDebug(in any) string {
	return spew.Sdump(in)
}

func printDate() string {
	return time.Now().Format(time.DateOnly)
}
