// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package scanner

import (
	"strings"

	"golang.org/x/tools/go/packages"
)

// buildImportAliases scans import declarations to find aliases used in the source.
//
// This bridges the AST view (import aliases) with the types view (package paths).
func buildImportAliases(pkg *packages.Package) map[string]string {
	aliases := make(map[string]string)

	for _, astFile := range pkg.Syntax {
		for _, importSpec := range astFile.Imports {
			// get the import path (remove quotes)
			importPath := strings.Trim(importSpec.Path.Value, `"`)

			var alias string
			if importSpec.Name != nil {
				// explicit alias: import foo "bar/baz"
				alias = importSpec.Name.Name

				// skip blank imports and dot imports for type qualification
				if alias == "_" || alias == "." {
					continue
				}
			} else {
				// no explicit alias - need to determine the package name
				// try to find it in the loaded imports
				for _, imported := range pkg.Imports {
					if imported.PkgPath == importPath {
						alias = imported.Name
						break
					}
				}

				// fallback: use last segment of import path
				if alias == "" {
					parts := strings.Split(importPath, "/")
					alias = parts[len(parts)-1]
				}
			}

			// store the mapping (later imports override earlier ones if there are conflicts)
			aliases[importPath] = alias
		}
	}

	return aliases
}
