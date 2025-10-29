package main

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const tokenString = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJleGFtcGxlLWF1ZGllbmNlIiwiZXhwIjoxNzQ3NzQ1Nzk1LCJpYXQiOjE3NDc3NDIxOTUsImlzcyI6ImV4YW1wbGUuY29tIiwibmJmIjoxNzQ3NzQyMTk1LCJyb2xlIjoiYWRtaW4iLCJzdWIiOiJ1c2VyMTIzIn0.m_-4d8YW21PZM4n0sd8Lj34aLFw3SJF8oJLfazItLkL7-c258F1fzbEUtMIQxGAKKhF5GpUlD78eJX0tP-52MLuwt_nLOLD0nYKLNYFJqkwjgsTboW07p_XjQ_fkQrUNQMn-srzKpeeuvWpsMH6WgSObxvRL-XUX9dNH1IT_aTAH-eFeVzyA5ldszLgegPdPwKpJcH5sM2gDe6oZ9jTM2kT3luStrzHmQnYsm6o8W0jgGLwE1JJvTZtaQUXBSW4AUVH4oCovBMwVDkjIHEA_VUMfAcRU6Bw_gD-UjQb4iSeqm7wNf4wKMHSbxOT4nvCr_gMdp09wJV4v2wqGXDGiyg`

func main() {
	// Читаем публичный ключ
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

	// Преобразуем к типу *rsa.PublicKey
	rsaPublicKey, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		panic("не удалось преобразовать публичный ключ к типу *rsa.PublicKey")
	}

	// Парсим и проверяем токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		// Проверяем метод подписи
		_, ok := token.Method.(*jwt.SigningMethodRSA)
		if !ok {
			return nil, fmt.Errorf("неверный метод подписи: %v", token.Header["alg"])
		}
		return rsaPublicKey, nil
	})
	if err != nil {
		log.Fatal("Ошибка верификации токена:", err)
	}

	// Парсим payload
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		panic("Ошибка приведения к jwt.MapClaims")
	}

	// Позже валидность можно проверить так
	if token.Valid {
		fmt.Println("Токен валиден! Claims:", claims)
	} else {
		fmt.Println("Токен недействителен")
	}

	// Проверяем срок действия токена
	if claims.VerifyExpiresAt(time.Now().Unix(), true) {
		fmt.Println("Токен действителен!")
	}

	// Сейчас время
	fmt.Println("Сейчас время:", time.Now().Format(time.DateTime))

	// Время создания токена
	if claims["iat"] != nil {
		iat := claims["iat"].(float64)
		fmt.Println("Время создания токена:", time.Unix(int64(iat), 0).Format(time.DateTime))
	} else {
		fmt.Println("Время создания токена не указано")
	}

	// Время истечения токена
	if claims["exp"] != nil {
		exp := claims["exp"].(float64)
		fmt.Println("Время истечения токена:", time.Unix(int64(exp), 0).Format(time.DateTime))
	} else {
		fmt.Println("Время истечения токена не указано")
	}
}
