package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	creds := map[string]string{"user": "pass"}

	r := chi.NewRouter()
	r.Use(middleware.BasicAuth("profile api", creds))

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("Hello World")
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
