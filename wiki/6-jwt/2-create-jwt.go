package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWT полностью описан в RFC 7519
func main() {
	// Читаем приватный ключ
	privateKeyFile, err := os.ReadFile("./keys/private_key.pem")
	if err != nil {
		panic(err)
	}

	block, _ := pem.Decode(privateKeyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		panic("не удалось декодировать приватный ключ")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	// Создаем claims для токена
	claims := jwt.MapClaims{
		"iss":  "example.com",                          // Кто выдал токен
		"sub":  "7f5da262-848c-4892-b1bd-5cb32fc5ac8b", // Для кого токен
		"aud":  "example-audience",                     // Аудитория токена
		"exp":  time.Now().Add(1 * time.Hour).Unix(),   // Срок действия токена
		"iat":  time.Now().Unix(),                      // Время выпуска токена
		"nbf":  time.Now().Unix(),                      // Токен не действителен до этого времени
		"role": "user",                                 // Кастомный claim
	}

	// Создаем токен
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Подписываем токен приватным ключом
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		panic(err)
	}

	// Первые 2 части токена, это просто JSON закодированый в base64
	// echo "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9" | base64 -d
	// echo "eyJhdWQiOiJleGFtcGxlLWF1ZGllbmNlIiwiZXhwIjoxNzQ3NzYwNzI4LCJpYXQiOjE3NDc3NTcxMjgsImlzcyI6ImV4YW1wbGUuY29tIiwibmJmIjoxNzQ3NzU3MTI4LCJyb2xlIjoidXNlciIsInN1YiI6IjdmNWRhMjYyLTg0OGMtNDg5Mi1iMWJkLTVjYjMyZmM1YWM4YiJ9" | base64 -d
	// Третья часть - цифровая подпись
	fmt.Println("JWT Token:", signedToken)
}
