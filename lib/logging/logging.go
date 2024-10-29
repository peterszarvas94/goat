package logging

import (
	"log/slog"
	"os"
)

func NewLogger(handler slog.Handler) *slog.Logger {
	return slog.New(handler)
}

var jsonHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
	Level:     slog.LevelDebug,
	AddSource: true,
})

var Logger = NewLogger(jsonHandler)

// func init logger with options from project
