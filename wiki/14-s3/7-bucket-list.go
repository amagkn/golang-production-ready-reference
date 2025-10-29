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

	{ // Список объектов в бакете
		const bucketName = "mybucket"

		for object := range minioClient.ListObjects(context.Background(), bucketName, minio.ListObjectsOptions{Recursive: true}) {
			if object.Err != nil {
				log.Fatalln(object.Err)
			}
			log.Println("📦", object.Key)
		}
	}
}
