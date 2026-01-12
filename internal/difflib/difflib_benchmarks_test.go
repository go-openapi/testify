package difflib

import (
	"strings"
	"testing"
)

func BenchmarkSplitLines100(b *testing.B) {
	b.Run("splitLines", benchmarkSplitLines(100))
}

func BenchmarkSplitLines10000(b *testing.B) {
	b.Run("splitLines", benchmarkSplitLines(10000))
}

func benchmarkSplitLines(count int) func(*testing.B) {
	return func(b *testing.B) {
		str := strings.Repeat("foo\n", count)

		b.ResetTimer()

		n := 0
		for b.Loop() {
			n += len(SplitLines(str))
		}
	}
}
