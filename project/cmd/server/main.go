package main

import (
	"fmt"
	"log/slog"
	"os"
	"project/config"
	"project/handlers"
	"project/templates/pages"
	"strings"

	"github.com/peterszarvas94/goat/logging"
	"github.com/peterszarvas94/goat/routing"
)

func main() {
	// set up logger
	err := logging.Setup("logs", "server-logs", slog.LevelDebug)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	logger := logging.Logger
	logger.Debug("Logger set up done")

	// set up router
	router := routing.Router

	router.GetTempl("/{$}", pages.Index())
	router.GetTempl("/test/{$}", pages.Test())
	router.Get("/hello/{$}", handlers.MyHandlerFunc)

	url := strings.Join([]string{"localhost", config.Port}, ":")

	logger.Info("Starting server on http://localhost:8080")

	err = router.Serve(url)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
