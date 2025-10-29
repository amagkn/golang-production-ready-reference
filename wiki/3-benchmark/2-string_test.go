package benchmark_test

import (
	"strings"
	"testing"
)

func Benchmark_StringConcatenation_A(b *testing.B) {
	f := func() {
		var str string

		for i := 0; i < 1_000; i++ {
			str += "user"
		}
	}

	for range b.N {
		f()
	}
}

func Benchmark_StringConcatenation_B(b *testing.B) {
	f := func() {
		var builder strings.Builder
		builder.Grow(1_000)

		for i := 0; i < 1_000; i++ {
			builder.WriteString("user")
		}
	}

	for range b.N {
		f()
	}
}
