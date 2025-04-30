package main

import (
	"os"

	"scaffhold/config"
	. "scaffhold/controllers/middlewares"
	"scaffhold/controllers/pages"
	"scaffhold/controllers/procedures"

	"github.com/peterszarvas94/goat/env"
	"github.com/peterszarvas94/goat/importmap"
	"github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func main() {
	// set up logger
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

	router.Use(Cache, AddReqID)

	router.Get("/", pages.NotFound)
	router.Get("/{$}", pages.Index)
	router.Get("/count", procedures.GetCount)
	router.Post("/count", procedures.PostCount)

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}
