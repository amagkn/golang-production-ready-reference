package s3

import (
	"fmt"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

type Config struct {
	Endpoint string
	Login    string
	Pass     string
}

func New(c Config) (*minio.Client, error) {
	conn, err := minio.New(c.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(c.Login, c.Pass, ""),
		Secure: false,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to S3: %w", err)
	}

	return conn, err
}
