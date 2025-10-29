package main

import (
	"bytes"
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

	{ // Загрузка данных
		const (
			bucketName  = "mybucket"
			content     = "Hello world from MinIO!"
			prefix      = "my-path/path/"
			objectName  = prefix + "hello.txt"
			contentType = "text/plain"
		)

		// Преобразуем строку в io.Reader
		reader := bytes.NewReader([]byte(content))

		// Загрузка строки как объекта
		info, err := minioClient.PutObject(context.Background(), bucketName, objectName, reader, int64(reader.Len()), minio.PutObjectOptions{
			ContentType: contentType,
		})
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("✅ Загружен объект %s (размер: %d байт)\n", info.Key, info.Size)
	}
}
