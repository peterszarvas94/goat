package config

import (
	"log/slog"
)

var (
	AppName  = "scaffhold"
	LogLevel = slog.LevelDebug
)

type envT struct {
	DbPath  string
	GoatEnv string
	Port    string
}

var Vars envT
