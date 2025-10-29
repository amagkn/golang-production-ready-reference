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

	{ // Загрузка данных
		const (
			bucketName  = "mybucket"
			content     = "Hello world from MinIO!"
			objectName  = "hello.txt"
			contentType = "text/plain"
		)

		// Преобразуем строку в io.Reader
		reader := bytes.NewReader([]byte(content))

		// Загрузка строки как объекта
		info, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, int64(reader.Len()),
			minio.PutObjectOptions{
				ContentType:  contentType,
				UserMetadata: map[string]string{"author": "Alice"},
				UserTags:     map[string]string{"env": "dev"},
				CacheControl: "public, max-age=3600",    // при доступе из браузера будет добавлен заголовок Cache-Control
				Expires:      time.Now().Add(time.Hour), // при доступе из браузера будет добавлен заголовок Expires

			})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("✅ Загружен объект %s (размер: %d байт)\n", info.Key, info.Size)
	}

	{ // Получение данных
		const (
			bucketName = "mybucket"
			objectName = "hello.txt"
		)

		// Получаем объект
		object, err := minioClient.GetObject(context.Background(), bucketName, objectName, minio.GetObjectOptions{})
		if err != nil {
			log.Fatalln(err)
		}

		// Читаем данные из объекта
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(object); err != nil {
			log.Fatalln(err)
		}

		// Выводим данные
		fmt.Println("Данные из объекта:", buf.String())
	}
}
