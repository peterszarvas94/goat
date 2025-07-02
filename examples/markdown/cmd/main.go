package main

import (
	"fmt"
	"log/slog"
	"os"

	"markdown/config"
	. "markdown/controllers/middlewares"
	"markdown/views/components"

	. "markdown/controllers/pages"
	. "markdown/controllers/procedures"

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
		slog.Error(fmt.Sprintf("Can not load env vars: %v\n", err))
		os.Exit(1)
	}

	// set up scripts
	err = importmap.Setup()
	if err != nil {
		slog.Error(fmt.Sprintf("Can not setup importmap: %v\n", err))
		os.Exit(1)
	}

	// parse md files
	content.RegisterTemplate(components.Md)
	_, err = content.Setup()
	if err != nil {
		slog.Error(fmt.Sprintf("Can not convert content: %v\n", err))
		os.Exit(1)
	}

	// get 404/index.html
	err, notFoundFile := content.GetNotFoundFile()
	if err != nil {
		slog.Error(fmt.Sprintf("Can not get 'notFoundFile': %v\n", err))
		os.Exit(1)
	}

	// set up server
	url := server.NewLocalHostUrl(config.Vars.Port)

	router := server.NewRouter()

	router.Use(RemoveTrailingSlash, Cache, AddRequestId)

	router.Setup()

	router.StaticFile("/", notFoundFile.HtmlPath)
	router.Get("/{$}", IndexPageHandler)
	router.Get("/tag/{tag}", TagPageHandler)
	router.Get("/category/{category}", CategoryPageHandler)

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}
