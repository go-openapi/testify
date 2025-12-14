// SPDX-FileCopyrightText: Copyright 2025 go-swagger maintainers
// SPDX-License-Identifier: Apache-2.0

package assertions

import "testing"

func Benchmark_isEmpty(b *testing.B) {
	b.ReportAllocs()

	v := new(int)

	for b.Loop() {
		isEmpty("")
		isEmpty(42)
		isEmpty(v)
	}
}

func BenchmarkNotNil(b *testing.B) {
	for b.Loop() {
		NotNil(b, b)
	}
}

func BenchmarkBytesEqual(b *testing.B) {
	const size = 1024 * 8
	s := make([]byte, size)
	for i := range s {
		s[i] = byte(i % 255)
	}
	s2 := make([]byte, size)
	copy(s2, s)

	mockT := &mockFailNowT{}

	for b.Loop() {
		Equal(mockT, s, s2)
	}
}
