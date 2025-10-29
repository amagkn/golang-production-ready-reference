package main

import (
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
)

// 1. CPU profile
// curl -s http://localhost:9090/debug/pprof/profile?seconds=5 > ./cpu.pprof
// go tool pprof -http=:8080 ./cpu.pprof

// 2. Mem profile
// curl -s http://localhost:9090/debug/pprof/heap > ./heap.out
// go tool pprof -http=:8080 ./a_mem.pprof

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

	server := &http.Server{
		Addr:    ":9090",
		Handler: mux,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			panic(err)
		}
	}()

	fn := func(count int) string {
		var str string

		for i := 0; i < count; i++ {
			str += "user"
		}

		return str
	}

	_ = fn(100_000)

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt)
	<-wait
}
