package config

import (
	"log/slog"
)

var (
	Port     = "8080"
	LogLevel = slog.LevelDebug
)

type EnvT struct {
	DbUrl   string
	DbToken string
}

var Env EnvT
