package profile_client_gen

import (
	"errors"
	"fmt"

	http_client "github.com/amagkn/golang-production-ready-reference/gen/http/profile_v1/client"
)

var ErrNotFound = errors.New("not found")

type Config struct {
	Host string `default:"localhost" envconfig:"HTTP_CLIENT_HOST"`
	Port string `default:"8080"      envconfig:"HTTP_CLIENT_PORT"`
}

type Client struct {
	client *http_client.ClientWithResponses
}

func New(c Config) (*Client, error) {
	baseURL := fmt.Sprintf("http://%s:%s/mnepryakhin/my-app/api/v1", c.Host, c.Port)

	client, err := http_client.NewClientWithResponses(baseURL)
	if err != nil {
		return nil, fmt.Errorf("http_client.NewClient: %w", err)
	}

	return &Client{client: client}, nil
}
