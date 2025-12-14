package generator

import (
	"fmt"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"

	"github.com/go-openapi/testify/v2/codegen/internal/model"
)

func funcMap() template.FuncMap {
	return map[string]any{
		"imports":          printImports,
		"comment":          comment,
		"params":           params,
		"forward":          forward,
		"docStringFor":     docStringFor,
		"docStringPackage": docStringPackage,
		"returns":          printReturns,
		"concat":           concatStrings,
		"pathparts":        pathParts,
		"relocate":         relocate,
		"hasSuffix":        strings.HasSuffix,
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

// docStringFor adds an extra comment for context.
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
func docStringPackage(pkg string) string {
	switch pkg {
	case pkgAssert:
		return `// Upon failure, the test [T] is marked as failed and continues execution.`
	case pkgRequire:
		return `// Upon failure, the test [T] is marked as failed and stops execution.`
	default:
		return ""
	}
}

func printReturns(in model.Parameters) string {
	return fmt.Sprintf("%v", in)
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

// relocate handles special cases in testable examples.
//
// It rewrites specific values to be run from a {{ .Package }}_test external package.
//
// Since there are only a few such edge cases, we hard-code this is this funcmap
// rather than indulging into a full-fleged go value parsing and package relocating.
func relocate(in, pkg string) string {
	if pkg == "" {
		return in
	}

	replaced := errTestRex.ReplaceAllString(in, "${1}"+pkg+".${2}")
	replaced = collectTRex.ReplaceAllString(replaced, "${1}"+pkg+".${2}")
	replaced = trueRex.ReplaceAllString(replaced, "${1}"+pkg+".${2}")

	return replaced
}

var (
	errTestRex  = regexp.MustCompile(`(^|[^\.])(ErrTest)`)
	collectTRex = regexp.MustCompile(`(^|[^\.])(CollectT)`)
	trueRex     = regexp.MustCompile(`(^|[^\.])(True\(c)`)
)
