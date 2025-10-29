package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(FirstMiddleware, SecondMiddleware)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("Hello World")
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func FirstMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("First: before next")

		next.ServeHTTP(w, r)

		fmt.Println("First: after next")
	}

	return http.HandlerFunc(fn)
}

func SecondMiddleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Second: before next")

		next.ServeHTTP(w, r)

		fmt.Println("Second: after next")
	}

	return http.HandlerFunc(fn)
}
