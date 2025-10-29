package main

import (
	"strings"

	"github.com/pkg/profile"
)

// 1. CPU profile
// go tool pprof -http=:8080 ./a_cpu.pprof

// 2. Mem profile
// go tool pprof -http=:8080 ./a_mem.pprof

// 3. Trace profile
// go tool trace -http=:8080 ./a_trace.out

// 4. Сравнение профилей
// go tool pprof -http=:8080 -base=a_mem.pprof b_mem.pprof

// Если ошибка graphviz, установить https://graphviz.org/download/

func main() {
	defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	//defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()

	const count = 10_000

	//_ = ConcatenateA(count)
	_ = ConcatenateB(count)
}

func ConcatenateA(count int) string {
	var str string

	for i := 0; i < count; i++ {
		str += "user"
	}

	return str
}

func ConcatenateB(count int) string {
	var builder strings.Builder
	builder.Grow(count)

	for i := 0; i < count; i++ {
		builder.WriteString("user")
	}

	return builder.String()
}
