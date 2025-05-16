package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"time"
)

// Creates 2 logger:
//
// 1 - JSON formatted logger for the log file
//
// 2 - pretty logger for the terminal output
func Setup(folderPath, filePrefix string, level slog.Level) error {
	filename := fmt.Sprintf("%s-%s.txt", filePrefix, time.Now().Format("2006-01-02"))
	filePath := filepath.Join(folderPath, filename)

	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		fmt.Println(err.Error())
		return err
	}

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	jsonHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: level,
	})

	textHandler := NewPrettyHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	})

	handler := NewMultiHandler(jsonHandler, textHandler)

	myLogger := slog.New(handler)

	slog.SetDefault(myLogger)

	slog.Debug("Logger set up done")
	return nil
}
