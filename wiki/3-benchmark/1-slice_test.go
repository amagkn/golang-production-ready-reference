package benchmark_test

import (
	"testing"
)

// go test -bench=. -benchmem ./...
// Goland: Program arguments: -test.benchmem

func Benchmark_SliceAppend_A(b *testing.B) {
	f := func() {
		var users []int

		for i := 0; i < 1_000; i++ {
			users = append(users, i)
		}
	}

	for range b.N {
		f()
	}
}

func Benchmark_SliceAppend_B(b *testing.B) {
	f := func() {
		var users = make([]int, 0, 1_000) // Если размер слайса меньше 64кб, он остаётся на стеке (0 аллокаций)

		for i := 0; i < 1_000; i++ {
			users = append(users, i)
		}
	}

	for range b.N {
		f()
	}
}
