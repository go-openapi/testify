package generator

import (
	"bytes"
	"fmt"
	"os"
	"text/template"

	"golang.org/x/tools/imports"
)

func render(tpl *template.Template, target string, data any, o *imports.Options) error {
	var buffer bytes.Buffer

	if err := tpl.Execute(&buffer, data); err != nil {
		return fmt.Errorf("error executing template %q: %w", tpl.Name(), err)
	}

	formatted, err := imports.Process(target, buffer.Bytes(), o)
	if err != nil {
		_ = os.WriteFile(target, buffer.Bytes(), filePermissions)
		return fmt.Errorf("error formatting go code: %w:%w", err, fmt.Errorf("details available at: %v", target))
	}

	return os.WriteFile(target, formatted, filePermissions)
}

func renderMD(tpl *template.Template, target string, data any) error {
	var buffer bytes.Buffer

	if err := tpl.Execute(&buffer, data); err != nil {
		return fmt.Errorf("error executing template %q: %w", tpl.Name(), err)
	}

	return os.WriteFile(target, buffer.Bytes(), filePermissions)
}
