package main

import (
	"os"

	"bare/config"
	. "bare/controllers/middlewares"
	. "bare/controllers/pages"
	. "bare/controllers/procedures"

	"github.com/peterszarvas94/goat/pkg/env"
	"github.com/peterszarvas94/goat/pkg/importmap"
	"github.com/peterszarvas94/goat/pkg/logger"
	"github.com/peterszarvas94/goat/pkg/server"
	"github.com/peterszarvas94/goat/pkg/uuid"
)

func main() {
	// set up slog
	err := logger.Setup("logs", "server-logs", config.LogLevel)
	if err != nil {
		os.Exit(1)
	}

	// set up env vars
	err = env.Load(&config.Vars)
	if err != nil {
		os.Exit(1)
	}

	// set up scripts
	err = importmap.Setup()
	if err != nil {
		os.Exit(1)
	}

	// set up server
	url := server.NewLocalHostUrl(config.Vars.Port)

	router := server.NewRouter()

	router.Setup()

	router.Use(RemoveTrailingSlash, Cache, AddRequestId)

	router.Get("/", NotFoundPageHandler)
	router.Get("/{$}", IndexPageHandler)

	router.Get("/count", GetCountHandler)
	router.Post("/count", PostCountHandler)

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}
