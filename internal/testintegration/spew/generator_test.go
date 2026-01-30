// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package spew

import (
	"os"
	"testing"

	"pgregory.net/rapid"
)

func TestTypeGenerator(t *testing.T) {
	t.Parallel()

	// this test exercises the generator and not [spew.Sdump].
	logger := testLogger(t)
	const numExamples = 5

	t.Run("with primitive types", rapid.MakeCheck(func(tr *rapid.T) {
		g := genPrimitiveValue()
		for range numExamples {
			value := g.Draw(tr, "sample")
			logger(value)
		}
	}))

	t.Run("example with array type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genArrayValue(0).Example()
			logger(value)
		}
	})

	t.Run("example with struct type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genStructValue(0).Example()
			logger(value)
		}
	})

	t.Run("example with slice type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genSliceValue(0).Example()
			logger(value)
		}
	})

	t.Run("example with map type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genMapValue(0).Example()
			logger(value)
		}
	})

	t.Run("example with pointer type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genNewPointer(0).Example()
			logger(value)
		}
	})

	t.Run("example with channel type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genChanValue(0).Example()
			logger(value)
		}
	})

	t.Run("example with other values", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genOtherValue(0).Example()
			logger(value)
		}
	})

	t.Run("example with any container type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genContainerValue(0).Example()
			logger(value)
		}
	})

	t.Run("example with any type", func(_ *testing.T) {
		g := newTypeGenerator()
		for range numExamples {
			value := g.genAnything(0).Example()
			logger(value)
		}
	})

	t.Run("with check on any type", rapid.MakeCheck(func(tr *rapid.T) {
		g := newTypeGenerator()
		value := g.Generator().Draw(tr, "sample")
		logger(value)
	}))
}

func testLogger(t *testing.T) func(any) {
	t.Helper()

	isDebug := os.Getenv("DEBUG") != ""

	if !isDebug {
		return func(_ any) {}
	}

	return func(str any) {
		t.Logf("%v", str)
	}
}

func testLoad() int {
	isCI := os.Getenv("CI") != ""

	if isCI {
		return 100
	}

	return 100_000 // local testing explores more cases
}
