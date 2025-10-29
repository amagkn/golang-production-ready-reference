package benchmark_test

import (
	"testing"
)

func Benchmark_MapAdd_A(b *testing.B) {
	f := func() {
		var users = make(map[int]int)

		for i := 0; i < 1_000; i++ {
			users[i] = i
		}
	}

	for range b.N {
		f()
	}
}

func Benchmark_MapAdd_B(b *testing.B) {
	f := func() {
		var users = make(map[int]int, 1_000)

		for i := 0; i < 1_000; i++ {
			users[i] = i
		}
	}

	for range b.N {
		f()
	}
}
