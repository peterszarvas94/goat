package logger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
)

var Logger *slog.Logger

// Using this instead of the built-in AddSource option
func addSource(args ...any) []any {
	newArgs := args
	_, file, line, ok := runtime.Caller(3)
	if !ok {
		file = "unknown"
		line = 0
	}

	newArgs = append(newArgs, slog.String("file", file))
	newArgs = append(newArgs, slog.String("line", strconv.Itoa(line)))
	return newArgs
}

func Debug(msg string, args ...any) {
	args = addSource(args...)
	Logger.Debug(msg, args...)
}

func Info(msg string, args ...any) {
	args = addSource(args...)
	Logger.Info(msg, args...)
}

func Warn(msg string, args ...any) {
	args = addSource(args...)
	Logger.Warn(msg, args...)
}

func Error(msg string, args ...any) {
	args = addSource(args...)
	Logger.Error(msg, args...)
}

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

	Logger = slog.New(handler)

	Logger.Debug("Logger set up done")

	return nil
}
