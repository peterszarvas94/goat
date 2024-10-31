package config

import (
	"log/slog"
)

var (
	Port     = "8080"
	LogLevel = slog.LevelDebug
)

type ConfigT struct {
	DbUrl   string
	DbToken string
}

var Config ConfigT
