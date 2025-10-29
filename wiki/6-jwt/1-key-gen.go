package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
)

func main() {
	// Генерация приватного ключа
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Кодирование приватного ключа в PEM формат
	privateFile, err := os.Create("./19-jwt/keys/private_key.pem")
	if err != nil {
		panic(err)
	}
	defer privateFile.Close()

	privateBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	}
	if err := pem.Encode(privateFile, privateBlock); err != nil {
		panic(err)
	}

	// Генерация публичного ключа
	publicKey := &privateKey.PublicKey
	publicBytes, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		panic(err)
	}

	// Кодирование публичного ключа в PEM формат
	publicFile, err := os.Create("./19-jwt/keys/public_key.pem")
	if err != nil {
		panic(err)
	}
	defer publicFile.Close()

	publicBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicBytes,
	}
	if err := pem.Encode(publicFile, publicBlock); err != nil {
		panic(err)
	}
}
