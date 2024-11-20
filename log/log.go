package log

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

func newLogger(handler slog.Handler) *slog.Logger {
	return slog.New(handler)
}

var Logger *slog.Logger

func Setup(folderPath, filePrefix string, level slog.Level) error {
	filename := fmt.Sprintf("%s-%s.txt", filePrefix, time.Now().Format("2006-01-02"))
	filePath := filepath.Join(folderPath, filename)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		return err
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	})

	textHandler := NewPrettyHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
	})

	handler := NewMultiHandler(jsonHandler, textHandler)

	Logger = newLogger(handler)

	return nil
}
