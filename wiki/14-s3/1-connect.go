package main

import (
	"fmt"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func main() {
	const (
		endpoint = "localhost:9000"
		login    = "minioadmin"
		pass     = "minioadmin"
		useSSL   = false
	)

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(login, pass, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln(err)
	}

	if minioClient.IsOnline() {
		fmt.Println("✅ Connected")
	} else {
		fmt.Println("❌ Connection failed")
	}
}
