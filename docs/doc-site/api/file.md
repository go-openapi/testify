---
title: "File"
description: "Asserting OS Files"
modified: 2026-01-11
weight: 7
domains:
  - "file"
keywords:
  - "DirExists"
  - "DirExistsf"
  - "FileEmpty"
  - "FileEmptyf"
  - "FileExists"
  - "FileExistsf"
  - "FileNotEmpty"
  - "FileNotEmptyf"
  - "NoDirExists"
  - "NoDirExistsf"
  - "NoFileExists"
  - "NoFileExistsf"
---

Asserting OS Files

## Assertions

[![GoDoc][godoc-badge]][godoc-url]
{class="inline-badge"}

_All links point to <https://pkg.go.dev/github.com/go-openapi/testify/v2>_

This domain exposes 6 functionalities.

### DirExists

DirExists checks whether a directory exists in the given path. It also fails
if the path is a file rather a directory or there is an error checking whether it exists.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.DirExists(t, "path/to/directory")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: filepath.Join(testDataPath(),"existing_dir")
	failure: filepath.Join(testDataPath(),"non_existing_dir")
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.DirExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#DirExists) | package-level function |
| [`assert.DirExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#DirExistsf) | formatted variant |
| [`assert.(*Assertions).DirExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.DirExists) | method variant |
| [`assert.(*Assertions).DirExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.DirExistsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.DirExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#DirExists) | package-level function |
| [`require.DirExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#DirExistsf) | formatted variant |
| [`require.(*Assertions).DirExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.DirExists) | method variant |
| [`require.(*Assertions).DirExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.DirExistsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.DirExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#DirExists) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#DirExists](https://github.com/go-openapi/testify/blob/master/internal/assertions/file.go#L78)
{{% /tab %}}
{{< /tabs >}}

### FileEmpty

FileEmpty checks whether a file exists in the given path and is empty.
It fails if the file is not empty, if the path points to a directory or there is an error when trying to check the file.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.FileEmpty(t, "path/to/file")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: filepath.Join(testDataPath(),"empty_file")
	failure: filepath.Join(testDataPath(),"existing_file")
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.FileEmpty(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FileEmpty) | package-level function |
| [`assert.FileEmptyf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FileEmptyf) | formatted variant |
| [`assert.(*Assertions).FileEmpty(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.FileEmpty) | method variant |
| [`assert.(*Assertions).FileEmptyf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.FileEmptyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.FileEmpty(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FileEmpty) | package-level function |
| [`require.FileEmptyf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FileEmptyf) | formatted variant |
| [`require.(*Assertions).FileEmpty(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.FileEmpty) | method variant |
| [`require.(*Assertions).FileEmptyf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.FileEmptyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.FileEmpty(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#FileEmpty) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#FileEmpty](https://github.com/go-openapi/testify/blob/master/internal/assertions/file.go#L136)
{{% /tab %}}
{{< /tabs >}}

### FileExists

FileExists checks whether a file exists in the given path. It also fails if
the path points to a directory or there is an error when trying to check the file.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.FileExists(t, "path/to/file")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: filepath.Join(testDataPath(),"existing_file")
	failure: filepath.Join(testDataPath(),"non_existing_file")
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.FileExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FileExists) | package-level function |
| [`assert.FileExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FileExistsf) | formatted variant |
| [`assert.(*Assertions).FileExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.FileExists) | method variant |
| [`assert.(*Assertions).FileExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.FileExistsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.FileExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FileExists) | package-level function |
| [`require.FileExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FileExistsf) | formatted variant |
| [`require.(*Assertions).FileExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.FileExists) | method variant |
| [`require.(*Assertions).FileExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.FileExistsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.FileExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#FileExists) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#FileExists](https://github.com/go-openapi/testify/blob/master/internal/assertions/file.go#L23)
{{% /tab %}}
{{< /tabs >}}

### FileNotEmpty

FileNotEmpty checks whether a file exists in the given path and is not empty.
It fails if the file is empty, if the path points to a directory or there is an error when trying to check the file.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.FileNotEmpty(t, "path/to/file")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: filepath.Join(testDataPath(),"existing_file")
	failure: filepath.Join(testDataPath(),"empty_file")
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.FileNotEmpty(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FileNotEmpty) | package-level function |
| [`assert.FileNotEmptyf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#FileNotEmptyf) | formatted variant |
| [`assert.(*Assertions).FileNotEmpty(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.FileNotEmpty) | method variant |
| [`assert.(*Assertions).FileNotEmptyf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.FileNotEmptyf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.FileNotEmpty(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FileNotEmpty) | package-level function |
| [`require.FileNotEmptyf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#FileNotEmptyf) | formatted variant |
| [`require.(*Assertions).FileNotEmpty(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.FileNotEmpty) | method variant |
| [`require.(*Assertions).FileNotEmptyf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.FileNotEmptyf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.FileNotEmpty(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#FileNotEmpty) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#FileNotEmpty](https://github.com/go-openapi/testify/blob/master/internal/assertions/file.go#L177)
{{% /tab %}}
{{< /tabs >}}

### NoDirExists

NoDirExists checks whether a directory does not exist in the given path.
It fails if the path points to an existing _directory_ only.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NoDirExists(t, "path/to/directory")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: filepath.Join(testDataPath(),"non_existing_dir")
	failure: filepath.Join(testDataPath(),"existing_dir")
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NoDirExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoDirExists) | package-level function |
| [`assert.NoDirExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoDirExistsf) | formatted variant |
| [`assert.(*Assertions).NoDirExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoDirExists) | method variant |
| [`assert.(*Assertions).NoDirExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoDirExistsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NoDirExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoDirExists) | package-level function |
| [`require.NoDirExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoDirExistsf) | formatted variant |
| [`require.(*Assertions).NoDirExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoDirExists) | method variant |
| [`require.(*Assertions).NoDirExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoDirExistsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NoDirExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NoDirExists) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NoDirExists](https://github.com/go-openapi/testify/blob/master/internal/assertions/file.go#L107)
{{% /tab %}}
{{< /tabs >}}

### NoFileExists

NoFileExists checks whether a file does not exist in a given path. It fails
if the path points to an existing _file_ only.

{{% expand title="Examples" %}}
{{< tabs >}}
{{% tab title="Usage" %}}
```go
	assertions.NoFileExists(t, "path/to/file")
```
{{< /tab >}}
{{% tab title="Examples" %}}
```go
	success: filepath.Join(testDataPath(),"non_existing_file")
	failure: filepath.Join(testDataPath(),"existing_file")
```
{{< /tab >}}
{{< /tabs >}}
{{% /expand %}}

{{< tabs >}}
{{% tab title="assert" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`assert.NoFileExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoFileExists) | package-level function |
| [`assert.NoFileExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#NoFileExistsf) | formatted variant |
| [`assert.(*Assertions).NoFileExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoFileExists) | method variant |
| [`assert.(*Assertions).NoFileExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/assert#Assertions.NoFileExistsf) | method formatted variant |
{{% /tab %}}
{{% tab title="require" style="secondary" %}}
| Signature | Usage |
|--|--|
| [`require.NoFileExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoFileExists) | package-level function |
| [`require.NoFileExistsf(t T, path string, msg string, args ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#NoFileExistsf) | formatted variant |
| [`require.(*Assertions).NoFileExists(path string) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoFileExists) | method variant |
| [`require.(*Assertions).NoFileExistsf(path string, msg string, args ..any)`](https://pkg.go.dev/github.com/go-openapi/testify/v2/require#Assertions.NoFileExistsf) | method formatted variant |
{{% /tab %}}

{{% tab title="internal" style="accent" icon="wrench" %}}
| Signature | Usage |
|--|--| 
| [`assertions.NoFileExists(t T, path string, msgAndArgs ...any) bool`](https://pkg.go.dev/github.com/go-openapi/testify/v2/internal/assertions#NoFileExists) | internal implementation |

**Source:** [github.com/go-openapi/testify/v2/internal/assertions#NoFileExists](https://github.com/go-openapi/testify/blob/master/internal/assertions/file.go#L52)
{{% /tab %}}
{{< /tabs >}}

---

---

Generated with github.com/go-openapi/testify/codegen/v2

[godoc-badge]: https://pkg.go.dev/badge/github.com/go-openapi/testify/v2
[godoc-url]: https://pkg.go.dev/github.com/go-openapi/testify/v2

<!--
SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
SPDX-License-Identifier: Apache-2.0


Document generated by github.com/go-openapi/testify/codegen/v2 DO NOT EDIT.

Generated on 2026-01-11 (version e6b0793) using codegen version v2.1.9-0.20260111152118-e6b0793ba519+dirty [sha: e6b0793ba519fb22dc1887392e1465649a5a95ff]
-->
