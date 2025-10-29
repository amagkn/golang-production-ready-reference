package main

import (
	"fmt"
	"time"
)

type Server struct {
	Addr         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Option func(*Server)

func NewServer(opts ...Option) *Server {
	// Значения по умолчанию.
	s := &Server{
		Addr:         "localhost",
		Port:         8080,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	// Переопределяем настройки которые переданы в опциях
	for _, opt := range opts {
		opt(s)
	}

	return s
}

// WithAddr задает адрес сервера.
func WithAddr(addr string) Option {
	return func(s *Server) {
		s.Addr = addr
	}
}

// WithPort задает порт сервера.
func WithPort(port int) Option {
	return func(s *Server) {
		s.Port = port
	}
}

// WithReadTimeout задает таймаут чтения сервера.
func WithReadTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.ReadTimeout = timeout
	}
}

// WithWriteTimeout задает таймаут записи сервера.
func WithWriteTimeout(timeout time.Duration) Option {
	return func(s *Server) {
		s.WriteTimeout = timeout
	}
}

func main() {
	// Создаем сервер с кастомными опциями.
	server := NewServer(
		WithAddr("127.0.0.1"),
		WithPort(9090),
		WithReadTimeout(10*time.Second),
		WithWriteTimeout(10*time.Second),
	)

	fmt.Printf("Server running at %s:%d\n", server.Addr, server.Port)
	fmt.Printf("Read timeout: %s\n", server.ReadTimeout)
	fmt.Printf("Write timeout: %s\n", server.WriteTimeout)
}
