---
title: Motivation
description: Motivations for this fork
weight: 5
---

From the maintainers of `testify`, it looks like a v2 will eventually be released, but they'll do it at their own pace.

We like all the principles they exposed to build this v2. [See discussion about v2](https://github.com/stretchr/testify/discussions/1560).

However, at `go-openapi` we would like to address the well-known issues in `testify` with different priorities.

With this fork, we want to:

1. [x] remove all external dependencies.
2. [x] make it easy to maintain and extend.
3. [x] pare down some of the chrome that has been added over the years.

---

{{% notice style="primary" title="Extended hand" icon="hand" %}}
We hope that this endeavor will help the original project with a live-drill of what a v2 could look like.

Hopefully, some of our ideas will eventually percolate back into the original project and help the wider 
community of go developers write better, clearer test code.

Feedback is welcome and we are always happy to discuss with people who face the same problems as we do: avoid breaking changes, 
APIs that became bloated over a decade or so, uncontrolled dependencies, difficult choices when it comes to introduce
breaking changes, conflicting demands from users etc.
{{% /notice %}}

You might also be curious about our [ROADMAP](./maintainers/ROADMAP.md).

---

1. We wanted first to remove all external dependencies.

> For all our libraries and generated test code we don't want test dependencies
> to drill farther than `import github.com/go-openapi/testify/v2`, but on some specific (and controlled)
> occasions.
>
> In this fork, all external stuff is either internalized (`go-spew`, `difflib`),
> removed (`mocks`, `suite`, `http`) or specifically enabled by importing this module
> (`github.com/go-openapi/testify/enable/yaml/v2`).

2. Make it easy to maintain and extend.

> For go-openapi, testify should just be yet another part of our toolkit.
> We need it to work, be easily adaptable to our needs and not divert our development effort away from our other repos.
> This big refactor is an investment.

3. We want to pare down some of the chrome that has been added over the years

> The `go-openapi` libraries and the `go-swagger` project make a rather limited use of the vast API provided by `testify`.
>
> With this first version of the fork, we have removed `mocks` and `suite`, which we don't use.
> They might be added later on, with better controlled dependencies.
>
> In the forthcoming maintenance of this fork, much of the "chrome" or "ambiguous" API will be pared down.
> There is no commitment yet on the stability of the API.
>
> Chrome would be added later: we have the "enable" packages just for that for when external dependencies are needed.

4. We want to add new features like generics, more useful assertions for JSON and safety checks.

5. We want to get rid of the API quirks and gotchas that panic or return unexpected results.
