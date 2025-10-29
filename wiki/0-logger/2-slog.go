package main

import (
	"errors"
	"log/slog"
)

func main() {
	var ErrTest = errors.New("test error")

	slog.SetLogLoggerLevel(slog.LevelDebug)

	slog.Debug("message")
	slog.Info("message")
	slog.Warn("message")
	slog.Error("message", "error", ErrTest)
}
