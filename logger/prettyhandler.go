package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"strings"
	"sync"
	"time"
)

type PrettyHandler struct {
	handler slog.Handler
	out     io.Writer
	mu      *sync.Mutex
	options *slog.HandlerOptions
}

func NewPrettyHandler(out io.Writer, opts *slog.HandlerOptions) *PrettyHandler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	return &PrettyHandler{
		out:     out,
		mu:      &sync.Mutex{},
		options: opts,
	}
}

func (h *PrettyHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.options.Level.Level()
}

func (h *PrettyHandler) Handle(ctx context.Context, record slog.Record) error {
	level := record.Level.String()
	levelColor := getColorForLevel(record.Level)
	time := record.Time.Format(time.TimeOnly)

	var levelPrefix string
	switch record.Level {
	case slog.LevelDebug:
		levelPrefix = "DEBUG"
	case slog.LevelInfo:
		levelPrefix = "INFO "
	case slog.LevelWarn:
		levelPrefix = "WARN "
	case slog.LevelError:
		levelPrefix = "ERROR"
	default:
		levelPrefix = fmt.Sprintf("%5s", strings.ToUpper(level))
	}

	// Prepare message components
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("\x1b[%dm%s\x1b[0m ", levelColor, levelPrefix))
	builder.WriteString(fmt.Sprintf("\x1b[90m%s\x1b[0m ", time))

	// file := ""
	// line := ""
	//
	// record.Attrs(func(attr slog.Attr) bool {
	// 	if attr.Key == "file" {
	// 		file = attr.Value.String()
	// 	}
	//
	// 	if attr.Key == "line" {
	// 		line = attr.Value.String()
	// 	}
	//
	// 	return true
	// })
	//
	// if file != "" && line != "" {
	// 	builder.WriteString(fmt.Sprintf("\x1b[36m[%s:%s]\x1b[0m ", shortenPath(file), line))
	// }

	builder.WriteString(record.Message)

	// Add attributes
	record.Attrs(func(attr slog.Attr) bool {
		if attr.Key != "file" && attr.Key != "line" {
			builder.WriteString(fmt.Sprintf(" \x1b[33m%s\x1b[0m=\x1b[92m%v\x1b[0m", attr.Key, attr.Value))
		}
		return true
	})

	builder.WriteString("\n")

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write([]byte(builder.String()))
	return err
}

func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *PrettyHandler) WithGroup(name string) slog.Handler {
	return h
}

func getColorForLevel(level slog.Level) int {
	switch level {
	case slog.LevelDebug:
		return 34 // Blue
	case slog.LevelInfo:
		return 32 // Green
	case slog.LevelWarn:
		return 33 // Yellow
	case slog.LevelError:
		return 31 // Red
	default:
		return 37 // White
	}
}

func shortenPath(path string) string {
	parts := strings.Split(path, "/")
	if len(parts) > 2 {
		return strings.Join(parts[len(parts)-2:], "/")
	}
	return path
}
