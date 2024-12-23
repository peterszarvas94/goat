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

type Context map[string]string

func newLogger(handler slog.Handler) *slog.Logger {
	return slog.New(handler)
}

var Logger *slog.Logger
var ctx = make(Context)

func addContext(args ...any) []any {
	newArgs := args
	newArgs = append(newArgs, GetContext()...)
	return newArgs
}

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
	newArgs := args
	newArgs = addSource(newArgs...)
	newArgs = addContext(newArgs...)
	Logger.Debug(msg, newArgs...)
	ClearContext()
}

func Info(msg string, args ...any) {
	newArgs := args
	newArgs = addSource(newArgs...)
	newArgs = addContext(newArgs...)
	Logger.Info(msg, newArgs...)
	ClearContext()
}

func Warn(msg string, args ...any) {
	newArgs := args
	newArgs = addSource(args)
	newArgs = addContext(newArgs...)
	Logger.Warn(msg, newArgs...)
	ClearContext()
}

func Error(msg string, args ...any) {
	newArgs := args
	newArgs = addSource(args)
	newArgs = addContext(newArgs...)
	Logger.Error(msg, newArgs...)
	ClearContext()
}

func AddToContext(key string, value string) {
	ctx[key] = value
}

func ClearContext() {
	ctx = make(Context)
}

func GetContext() []any {
	var args []any
	for key, value := range ctx {
		args = append(args, slog.String(key, value))
	}
	return args
}

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
		// AddSource: true,
	})

	textHandler := NewPrettyHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
		// AddSource: true,
	})

	handler := NewMultiHandler(jsonHandler, textHandler)

	Logger = newLogger(handler)

	Logger.Debug("Logger set up done")

	return nil
}
