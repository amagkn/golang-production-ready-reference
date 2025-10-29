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

	{ // Удаление бакета (если он пустой)
		const bucketName = "mybucket"

		err = minioClient.RemoveBucket(context.Background(), bucketName)
		if err != nil {
			log.Fatalf("❌ Ошибка удаления: %v\n", err)
		}
		log.Println("✅ Бакет удалён:", bucketName)
	}

	{ // Force Delete (даже если не пустой)
		const bucketName = "mybucket"

		err = minioClient.RemoveBucketWithOptions(context.Background(), bucketName,
			minio.RemoveBucketOptions{ForceDelete: true})
		if err != nil {
			log.Fatalf("❌ Ошибка удаления: %v\n", err)
		}
		log.Println("✅ Бакет удалён:", bucketName)
	}
}
