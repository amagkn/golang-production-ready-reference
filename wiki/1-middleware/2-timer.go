package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(Timer)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("Hello World")
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func Timer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		next.ServeHTTP(w, r)

		fmt.Println("Total:", time.Since(now))
	}

	return http.HandlerFunc(fn)
}
