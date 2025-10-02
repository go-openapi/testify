# Testify [![Build Status](https://github.com/go-openapi/testify/actions/workflows/go-test.yml/badge.svg)](https://github.com/go-openapi/testify/actions?query=workflow%3A"go+test") [![codecov](https://codecov.io/gh/go-openapi/testify/branch/master/graph/badge.svg)](https://codecov.io/gh/go-openapi/testify)

[![Slack Status](https://slackin.goswagger.io/badge.svg)](https://slackin.goswagger.io)
[![license](https://img.shields.io/badge/license-Apache%20v2-orange.svg)](https://raw.githubusercontent.com/go-openapi/testify/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-openapi/testify.svg)](https://pkg.go.dev/github.com/go-openapi/testify)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-openapi/testify)](https://goreportcard.com/report/github.com/go-openapi/testify)

## Testify - Thou Shalt Write Tests

A golang set of packages that provide tools for testifying that your code will behave as you intend.

This is the go-openapi fork of the great [testify](https://github.com/stretchr/testify) package.

## Why this fork?

From the maintainers of `testify`, it looks like a v2 is coming up, but they'll do it at their own pace.

We like all the principles they put forward to build this v2. [See discussion about v2](https://github.com/stretchr/testify/discussions/1560)

However, at `go-openapi` we would like to address the well-known issues in `testify` with different priorities.

1. We want first to remove all external dependencies.

> For all our libraries and generated test code we don't want test dependencies
> to drill farther than `import github.com/go-openapi/testify/v2`, but on some specific (and controlled)
> occasions.
>
> In this fork, all external stuff is either internalized (`go-spew`, `difflib`),
> removed (`mocks`, `suite`, `http`) or specifically enabled by importing a specific module
> (`github.com/go-openapi/testify/v2/enable/yaml`).

2. We want to remove most of the chrome that has been added over the years

> The `go-openapi` libraries and the `go-swagger` project make a rather limited use of the vast API provided by `testify`.
>
> With this first version of the fork, we have removed `mocks` and `suite`, which we don't use.
> They might be added later on, with better controlled dependencies.
>
> In the forthcoming maintenances of this fork, much of the "chrome" or "ambiguous" API will be pared down.
> There is no commitment yet on the stability of the API.
>
> Chrome would be added later: we have the "enable" packages just for that.

3. We hope that this endeavour will help the original project with a live-drill of what a v2 could look like.
   We are always happy to discuss with people who face the same problems as we do: avoid breaking changes, 
   APIs that became bloated over a decade or so, uncontrolled dependencies, conflicting demands from users etc.

## What's next with this project?

1. [x] The first release comes with zero dependencies and an unstable API (see below [our use case](#usage-at-go-openapi))
2. This project is going to be injected as the main and sole test dependency of the `go-openapi` libraries and the `go-swagger` tool
3. Valuable pending pull requests from the original project could be merged (e.g. `JSONEqBytes`) or transformed as "enable" modules (e.g. colorized output)
4. Unclear assertions may be provided an alternative verb (e.g. `InDelta`)
5. Since we have leveled the go requirements to the rest of the go-openapi (currently go1.24) there is quite a bit of relinting lying ahead.

## Usage at go-openapi

At this moment, we have identified the following usage in our tools. This API shall remain stable.
Currently, there are no guarantees about the entry points not in this list.

```
Condition
Contains,Containsf
Empty,Emptyf
Equal,Equalf
EqualError,EqualErrorf
EqualValues,EqualValuesf
Error,Errorf
ErrorContains
ErrorIs
Fail,Failf
FailNow
False,Falsef
Greater
Implements
InDelta,InDeltaf
IsType,IsTypef
JSONEq,JSONEqf
Len,Lenf
Nil,Nilf
NoError,NoErrorf
NotContains,NotContainsf
NotEmpty,NotEmptyf
NotEqual
NotNil,NotNilf
NotPanics
NotZero
Panics,PanicsWithValue
Subset
True,Truef
YAMLEq,YAMLEqf
Zero,Zerof
```

## Installation

To use this package in your projects:

```cmd
    go get github.com/go-openapi/testify/v2
```

## Get started

Features include:

  * [Easy assertions](./original.md#assert-package)
  * ~[Mocking](./original.md#mock-package)~ removed
  * ~[Testing suite interfaces and functions](./original.md#suite-package)~ removed

## Examples

See [the original README)(./original.md)

## Licensing

See the license [NOTICE](./NOTICE), which recalls the licensing terms of all the pieces of software
distributed with this fork, including internalized libraries.

* [SPDX-License-Identifier: Apache-2.0](./LICENSE)
* [SPDX-License-Identifier: MIT](./NOTICE)
* [SPDX-License-Identifier: ISC](./internal/spew/LICENSE)
* [SPDX-License-Identifier: ](./internal/difflib/LICENSE)

## Contributing

Please feel free to submit issues, fork the repository and send pull requests!

When submitting an issue, we ask that you please include a complete test function that demonstrates the issue.
Extra credit for those using Testify to write the test code that demonstrates it.

Code generation is used. Run `go generate ./...` to update generated files.

## [The original README](./original.md)
