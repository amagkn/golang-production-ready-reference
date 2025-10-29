package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"s3/pkg/s3"

	"github.com/minio/minio-go/v7"
)

func main() {
	minioClient, err := s3.New(s3.Config{Endpoint: "localhost:9000", Login: "minioadmin", Pass: "minioadmin"})
	if err != nil {
		log.Fatalln(err)
	}

	{ // Создание защищенного бакета
		const location = "us-east-1"
		const bucketName = "mybucket-locker"

		exists, errBucketExists := minioClient.BucketExists(context.Background(), bucketName)
		if errBucketExists != nil {
			log.Fatalln(errBucketExists)
		}

		if !exists {
			err := minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{
				Region:        location, // регион, дефолт: "us-east-1"
				ObjectLocking: true,     // неизменяемое хранилище — WORM: Write Once, Read Many
			})
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println("✅ Bucket created")
		} else {
			fmt.Println("ℹ️ Bucket already exists")
		}
	}

	{ // Загрузка защищенных данных
		const (
			bucketName  = "mybucket-locker"
			objectName  = "hello.txt"
			content     = "Защищенные данные"
			contentType = "text/plain"
		)

		reader := bytes.NewReader([]byte(content))

		info, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{
			ContentType:     contentType,
			Mode:            minio.Governance,
			RetainUntilDate: time.Now().Add(time.Minute),
		})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("✅ Загружен объект %s (размер: %d байт)\n", info.Key, info.Size)
	}
}
