package config

import (
	"log/slog"
)

var (
	AppName  = "basic-auth"
	LogLevel = slog.LevelDebug
)

type envT struct {
	DbPath  string
	GoatEnv string
	Port    string
}

var Vars envT

func IsDevelopment() bool {
	return Vars.GoatEnv == "dev"
}
