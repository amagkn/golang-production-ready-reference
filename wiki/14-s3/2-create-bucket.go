package main

import (
	"context"
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

	{ // Создание бакета
		const bucketName = "mybucket"
		const location = "us-east-1"

		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists != nil {
			log.Fatalln(errBucketExists)
		}

		if !exists {
			err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{
				Region:        location, // регион, дефолт: "us-east-1"
				ObjectLocking: false,    // неизменяемое хранилище — WORM: Write Once, Read Many
			})
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("✅ Bucket created")

			return
		}

		fmt.Println("ℹ️ Bucket already exists")
	}
}
