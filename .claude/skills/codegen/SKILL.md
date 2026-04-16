# Code Generation

How the testify code and documentation generator works.

## Running

```bash
# Generate everything (code + docs) from internal/assertions/
go generate ./...

# Or run the generator directly
cd codegen && go run . -output-packages assert,require -include-doc

# Code only (skip docs)
cd codegen && go run . -output-packages assert,require -include-doc=false
```

## Generator pipeline

1. **Scanner** (`codegen/internal/scanner/`) parses `internal/assertions/` using
   `go/packages` and `go/types`. Extracts function signatures, doc comments,
   domain tags, examples, and metadata.
   - `comments/` -- doc comment extraction
   - `comments-parser/` -- domain tags, examples, metadata parsing
   - `signature/` -- function signature analysis

2. **Model** (`codegen/internal/model/`) holds the intermediate representation:
   functions, type params, tests, documentation, metrics.

3. **Generator** (`codegen/internal/generator/`) renders templates:
   - **Code templates** produce `assert/` and `require/` packages with all variants
   - **Doc templates** produce Hugo markdown in `docs/doc-site/api/`
   - **Metrics** produces `metrics.yaml` for Hugo site params

## Key templates

Located in `codegen/internal/generator/templates/`:

| Template | Produces |
|----------|----------|
| `assertion_assertions.gotmpl` | `assert` package-level functions |
| `assertion_format.gotmpl` | `assert` formatted variants (`*f`) |
| `assertion_forward.gotmpl` | `assert` forward methods |
| `requirement_*.gotmpl` | `require` equivalents (calls `FailNow`) |
| `doc_index.md.gotmpl` | API index page (`_index.md`) |
| `doc_page.md.gotmpl` | Per-domain doc pages |
| `doc_metrics.md.gotmpl` | Quick index & metrics page |

## Template functions

Custom template functions in `codegen/internal/generator/funcmaps/`:
- `Slugize(name)` -- converts function name to markdown anchor
- `Titleize(s)` -- title-cases a string
- `hopen` / `hclose` -- Hugo shortcode delimiters (`{{% ... %}}`)

## Domain organization

Functions are grouped by `// Domain: <name>` tags inside the function body.
Domain descriptions live in `internal/assertions/doc.go` as special comments.
The generator reorganizes package-based docs into domain-based pages
(19 domains currently).

## Generated output

**Never edit generated files directly.** They carry a `DO NOT EDIT` header.

| Output | Location |
|--------|----------|
| `assert/` package | Generated functions + tests |
| `require/` package | Generated functions + tests |
| `docs/doc-site/api/*.md` | Domain-organized Hugo pages |
| `docs/doc-site/api/metrics.md` | Quick index + API metrics |
| `hack/doc-site/hugo/metrics.yaml` | Hugo site params (counts) |

Exceptions (not generated): `assert/doc.go`, `require/doc.go`, ad-hoc testable examples.

## Adding support for a new construct

If the generator needs to support a new Go construct (e.g., new type param
pattern), the work is in:
1. `codegen/internal/scanner/` -- teach the scanner to extract it
2. `codegen/internal/model/` -- add fields to the model
3. `codegen/internal/generator/templates/` -- update templates to render it

The scanner and generator have comprehensive tests (~1,400+ lines across test files).
