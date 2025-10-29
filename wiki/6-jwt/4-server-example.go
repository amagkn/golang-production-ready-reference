package main

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v4"
)

type UserIDContextKey struct{}

func main() {
	publicKey := getPublicKey()

	r := chi.NewRouter()
	r.Use(
		AuthMiddleware(publicKey),
	)

	r.Get("/profile", handler)

	fmt.Println("Listening on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		panic(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(UserIDContextKey{}).(string)
	if !ok {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)

		return
	}

	json := fmt.Sprintf(`{"user_id":"%s"}`, userID)

	w.Write([]byte(json))
}

func AuthMiddleware(publicKey *rsa.PublicKey) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "Missing Authorization header", http.StatusUnauthorized)

				return
			}

			// Проверяем, что заголовок начинается с "Bearer "
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)

				return
			}

			tokenString := parts[1]

			token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
				// Проверяем метод подписи
				_, ok := token.Method.(*jwt.SigningMethodRSA)
				if !ok {
					return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
				}
				return publicKey, nil
			})
			if err != nil {
				http.Error(w, "Invalid or expired token", http.StatusUnauthorized)

				return
			}

			// Извлекаем claims
			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			}

			userID, ok := claims["sub"].(string)
			if !ok {
				http.Error(w, "Invalid token claims", http.StatusUnauthorized)

				return
			}

			// И передаем их в контекст
			ctx := context.WithValue(r.Context(), UserIDContextKey{}, userID)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func getPublicKey() *rsa.PublicKey {
	publicKeyFile, err := os.ReadFile("./keys/public_key.pem")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(publicKeyFile)
	if block == nil || block.Type != "PUBLIC KEY" {
		panic("не удалось декодировать публичный ключ")
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		panic("не удалось преобразовать публичный ключ к типу *rsa.PublicKey")
	}

	return rsaPublicKey
}
