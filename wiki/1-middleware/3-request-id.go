package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func main() {
	r := chi.NewRouter()
	r.Use(SetRequestID, GetRequestID)

	r.Get("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Println("Hello World")
	})

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}

func SetRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := uuid.NewString()

		ctx := Pack(r.Context(), id)

		fmt.Println("SetRequestID pack:", id)

		next.ServeHTTP(w, r.WithContext(ctx))
	}

	return http.HandlerFunc(fn)
}

func GetRequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := Unpack(r.Context())

		fmt.Println("GetRequestID unpack:", id)

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

type requestIDKey struct{}

func Pack(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey{}, id)
}

func Unpack(ctx context.Context) string {
	id, ok := ctx.Value(requestIDKey{}).(string)
	if !ok {
		return ""
	}

	return id
}
