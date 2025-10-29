package main

import (
	"context"
	"log"

	"s3/pkg/s3"

	"github.com/minio/minio-go/v7"
)

func main() {
	minioClient, err := s3.New(s3.Config{Endpoint: "localhost:9000", Login: "minioadmin", Pass: "minioadmin"})
	if err != nil {
		log.Fatalln(err)
	}

	const (
		srcBucket  = "bucket-a"
		dstBucket  = "bucket-b"
		objectName = "hello.txt"
	)

	// 1️⃣ Копируем объект
	src := minio.CopySrcOptions{
		Bucket: srcBucket,
		Object: objectName,
	}
	dst := minio.CopyDestOptions{
		Bucket: dstBucket,
		Object: objectName,
	}

	_, err = minioClient.CopyObject(context.Background(), dst, src)
	if err != nil {
		log.Fatalf("❌ Ошибка копирования: %v", err)
	}

	// 2️⃣ Удаляем исходный объект
	err = minioClient.RemoveObject(context.Background(), srcBucket, objectName, minio.RemoveObjectOptions{})
	if err != nil {
		log.Fatalf("❌ Ошибка удаления исходного файла: %v", err)
	}

	log.Println("✅ Объект перемещён из", srcBucket, "в", dstBucket)
}
