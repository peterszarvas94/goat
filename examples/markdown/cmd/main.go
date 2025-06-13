package main

import (
	"fmt"
	"os"

	"markdown/config"
	. "markdown/controllers/middlewares"

	. "markdown/controllers/pages"
	. "markdown/controllers/procedures"
	"markdown/views"

	"github.com/peterszarvas94/goat/pkg/content"
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
		fmt.Printf("Can not set up logger: %v\n", err)
		os.Exit(1)
	}

	// set up env vars
	err = env.Load(&config.Vars)
	if err != nil {
		fmt.Printf("Can not load env vars: %v\n", err)
		os.Exit(1)
	}

	// set up scripts
	err = importmap.Setup()
	if err != nil {
		fmt.Printf("Can not setup importmap: %v\n", err)
		os.Exit(1)
	}

	// parse md files
	content.RegisterTemplate(views.Md)
	_, err = content.Setup()
	if err != nil {
		fmt.Printf("Can not convert content: %v\n", err)
		os.Exit(1)
	}

	// set up server
	url := server.NewLocalHostUrl(config.Vars.Port)

	router := server.NewRouter()

	router.Use(RemoveTrailingSlash, Cache, AddRequestId)

	router.Setup()

	router.Get("/", NotFoundPageHandler)
	router.Get("/{$}", IndexPageHandler)

	router.Get("/count", GetCountHandler)
	router.Post("/count", PostCountHandler)

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}
