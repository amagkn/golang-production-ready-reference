package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"time"

	"s3/pkg/s3"
)

func main() {
	minioClient, err := s3.New(s3.Config{Endpoint: "localhost:9000", Login: "minioadmin", Pass: "minioadmin"})
	if err != nil {
		log.Fatalln(err)
	}

	{ // Ссылка на объект
		const (
			bucketName = "mybucket"
			objectName = "hello.txt"
		)

		reqParams := make(url.Values)
		reqParams.Set("response-content-disposition", "inline") // открывает в браузере

		url, err := minioClient.PresignedGetObject(
			context.Background(),
			bucketName,
			objectName,
			10*time.Minute, // Время жизни ссылки
			reqParams,      // Дополнительные параметры
		)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("Ссылка на приватный URL:", url)
		fmt.Println("Ссылка на публичный URL:", "http://localhost:9000/mybucket/hello.txt")
	}
}
