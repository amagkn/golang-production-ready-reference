package main

import (
	"fmt"
	"time"
)

type Config struct {
	Addr         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Server2 struct {
	config Config
}

func NewServer2(config Config) *Server2 {
	// Значения по умолчанию.
	defaultConfig := Config{
		Addr:         "localhost",
		Port:         8080,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Переопределяем настройки, которые переданы в конфиге
	if config.Addr != "" {
		defaultConfig.Addr = config.Addr
	}
	if config.Port != 0 {
		defaultConfig.Port = config.Port
	}
	if config.ReadTimeout != 0 {
		defaultConfig.ReadTimeout = config.ReadTimeout
	}
	if config.WriteTimeout != 0 {
		defaultConfig.WriteTimeout = config.WriteTimeout
	}

	return &Server2{config: defaultConfig}
}

func main() {
	config := Config{
		Addr:         "127.0.0.1",
		Port:         9090,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server := NewServer2(config)

	fmt.Printf("Server running at %s:%d\n", server.config.Addr, server.config.Port)
	fmt.Printf("Read timeout: %s\n", server.config.ReadTimeout)
	fmt.Printf("Write timeout: %s\n", server.config.WriteTimeout)
}
