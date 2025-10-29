package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"s3/pkg/s3"

	"github.com/minio/minio-go/v7"
)

func main() {
	minioClient, err := s3.New(s3.Config{Endpoint: "localhost:9000", Login: "minioadmin", Pass: "minioadmin"})
	if err != nil {
		log.Fatalln(err)
	}

	{ // Статистика по объекту
		const (
			bucketName = "mybucket"
			objectName = "hello.txt"
		)

		stat, err := minioClient.StatObject(context.Background(), bucketName, objectName, minio.StatObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}

		// JSON-вывод
		b, err := json.MarshalIndent(stat, "", "  ")
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("📦 ObjectInfo:")
		fmt.Println(string(b))
	}
}
