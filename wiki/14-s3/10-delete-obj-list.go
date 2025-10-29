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

	ctx := context.Background()
	const bucketName = "mybucket"

	// Список объектов для удаления
	objectKeys := []string{"hello.txt", "hello2.txt", "hello3.txt"}

	// Канал для передачи объектов на удаление
	objectsCh := make(chan minio.ObjectInfo)

	// Отправка объектов в канал
	go func() {
		defer close(objectsCh)
		for _, key := range objectKeys {
			objectsCh <- minio.ObjectInfo{Key: key}
		}
	}()

	// Удаление и обработка результатов
	errorCh := minioClient.RemoveObjectsWithResult(ctx, bucketName, objectsCh, minio.RemoveObjectsOptions{})

	for removeErr := range errorCh {
		if removeErr.Err != nil {
			log.Printf("❌ Не удалось удалить %s: %v", removeErr.ObjectName, removeErr.Err)
		} else {
			log.Printf("🗑 Удалено: %s", removeErr.ObjectName)
		}
	}
}
