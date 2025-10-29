package main

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"

	"s3/pkg/s3"
)

func main() {
	minioClient, err := s3.New(s3.Config{Endpoint: "localhost:9000", Login: "minioadmin", Pass: "minioadmin"})
	if err != nil {
		log.Fatalln(err)
	}

	{ // Удаление объекта
		const (
			bucketName = "mybucket"
			objectName = "hello.txt"
		)

		err := minioClient.RemoveObject(context.Background(), bucketName, objectName, minio.RemoveObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("🗑️ Object deleted")
	}
}
