package main

import (
	"fmt"
	"os"
	"strings"

	"project/config"
	"project/handlers"
	"project/models"
	"project/templates/pages"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/env"
	"github.com/peterszarvas94/goat/log"
	"github.com/peterszarvas94/goat/router"
)

func main() {
	// set up log.Logger
	err := log.Setup("logs", "server-logs", config.LogLevel)
	if err != nil {
		fmt.Printf("Logger setup err: %v\n", err)
		os.Exit(1)
	}

	log.Logger.Debug("Logger set up done")

	// set up env vars
	err = env.Load(&config.Vars)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Can not load env: %v", err))
		os.Exit(1)
	}

	// set up db
	err = database.StartSqliteConnection(config.Vars.DbPath)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Can not set up db connection: %v", err))
		os.Exit(1)
	}

	if config.Vars.Env == "dev" {
		// seed
		err = models.Seed()
		if err != nil {
			log.Logger.Error(fmt.Sprintf("Can not seed db: %v", err))
			os.Exit(1)
		}
	}

	// set up router
	list := []string{"one", "two", "three"}
	router.GetTempl("/{$}", pages.Index.Full(list))
	router.GetTempl("/test/{$}", pages.Test())
	router.Get("/hello/{$}", handlers.MyHandlerFunc)

	url := strings.Join([]string{"localhost", config.Port}, ":")

	log.Logger.Info(fmt.Sprintf("Starting server on: http://%s", url))

	err = router.Serve(url)
	if err != nil {
		log.Logger.Error(fmt.Sprintf("Server start error: %v", err))
		os.Exit(1)
	}
}
