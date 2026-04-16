# Documentation Site

Hugo-based documentation site for testify, auto-generated from source code.

## Running locally

```bash
# 1. Generate docs from source
go generate ./...

# 2. Start Hugo dev server
cd hack/doc-site/hugo
./gendoc.sh

# Visit http://localhost:1313/testify/
# Auto-reloads on changes to docs/doc-site/
```

## Site structure

```
hack/doc-site/hugo/
  hugo.yaml              # Main Hugo config
  metrics.yaml           # Generated metrics (codegen output, merged into site params)
  testify.yaml           # Version info + build metadata
  gendoc.sh              # Dev server launcher
  layouts/               # Custom layout overrides
  themes/hugo-relearn/   # Relearn documentation theme

docs/doc-site/           # Content (mounted by Hugo)
  api/                   # Generated: domain pages, index, metrics (DO NOT EDIT)
  usage/                 # Hand-written: USAGE, GENERICS, CHANGES, MIGRATION, etc.
  project/               # Hand-written: APPROACH, maintainer docs
```

## Generated vs hand-written content

| Path | Generated? | Notes |
|------|-----------|-------|
| `docs/doc-site/api/*.md` | Yes | Domain pages, index, metrics. Regenerate with `go generate` |
| `docs/doc-site/api/metrics.md` | Yes | Quick index table + API counts |
| `docs/doc-site/usage/*.md` | No | Hand-written guides |
| `docs/doc-site/project/*.md` | No | Hand-written project docs |

## Dynamic counts via Hugo params

Function and assertion counts are generated into `metrics.yaml` and merged
into Hugo's `site.Params.metrics`. Use the relearn `siteparam` shortcode
to reference them in hand-written markdown:

```markdown
We have {{% siteparam "metrics.assertions" %}} assertions across
{{% siteparam "metrics.domains" %}} domains.
```

Available params: `metrics.domains`, `metrics.functions`, `metrics.assertions`,
`metrics.generics`, `metrics.nongeneric_assertions`, `metrics.helpers`, `metrics.others`.

Per-domain: `metrics.by_domain.<slug>.count`, `metrics.by_domain.<slug>.name`.

Hugo math functions (`sub`, `mul`, `add`) are NOT available in markdown content.
For computed values, add them to the codegen `buildMetrics()` in
`codegen/internal/generator/doc_generator.go`.

## Adding a new hand-written page

1. Create `docs/doc-site/<section>/<FILE>.md` with Hugo front matter
2. Set `weight:` to control ordering in the sidebar
3. Use relearn shortcodes: `{{% notice %}}`, `{{% expand %}}`, `{{< tabs >}}`, etc.
4. Reference API counts with `{{% siteparam "metrics.<key>" %}}`

## Relearn theme features used

- `{{% notice style="info" %}}` -- callout boxes
- `{{% expand title="..." %}}` -- collapsible sections
- `{{< tabs >}}` / `{{% tab %}}` -- tabbed content
- `{{< cards >}}` / `{{% card %}}` -- side-by-side cards
- `{{% icon icon="star" color=orange %}}` -- inline icons
- `{{% siteparam "key" %}}` -- site param substitution
- `{{< mermaid >}}` -- diagrams
