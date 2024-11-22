package main

import (
	"fmt"
	"os"

	"bootstrap/config"
	"bootstrap/handlers"
	"bootstrap/models"
	"bootstrap/templates/pages"

	_ "github.com/joho/godotenv/autoload"
	"github.com/peterszarvas94/goat/database"
	"github.com/peterszarvas94/goat/env"
	l "github.com/peterszarvas94/goat/logger"
	"github.com/peterszarvas94/goat/server"
	"github.com/peterszarvas94/goat/uuid"
)

func main() {
	// set up log.Logger
	err := l.Setup("logs", "server-logs", config.LogLevel)
	if err != nil {
		os.Exit(1)
	}

	// set up env vars
	err = env.Load(&config.Vars)
	if err != nil {
		os.Exit(1)
	}

	// set up db
	err = database.StartSqliteConnection(config.Vars.DbPath)
	if err != nil {
		os.Exit(1)
	}

	if config.Vars.Env == "dev" {
		// seed
		err = models.Seed()
		if err != nil {
			l.Logger.Error(fmt.Sprintf("Can not seed db: %v", err))
			os.Exit(1)
		}
	}

	// set up server
	url := server.NewLocalHostUrl(config.Port)

	mux := server.NewMux(url)
	list := []string{"one", "two", "three"}
	mux.TemplGet("/{$}", pages.Index.Full(list))
	mux.Get("/hello/{$}", handlers.MyHandlerFunc)

	serverId := uuid.New("srv")
	s := server.NewServer(mux)

	s.Serve(url, serverId)

}
