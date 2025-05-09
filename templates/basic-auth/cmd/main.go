package main

import (
	"context"
	"os"

	"basic-auth/config"
	. "basic-auth/controllers/middlewares"
	"basic-auth/controllers/pages"
	"basic-auth/controllers/procedures"
	"basic-auth/db/models"
	pageViews "basic-auth/views/pages"

	"github.com/peterszarvas94/goat/pkg/csrf"
	"github.com/peterszarvas94/goat/pkg/database"
	"github.com/peterszarvas94/goat/pkg/env"
	"github.com/peterszarvas94/goat/pkg/importmap"
	"github.com/peterszarvas94/goat/pkg/logger"
	"github.com/peterszarvas94/goat/pkg/server"
	"github.com/peterszarvas94/goat/pkg/uuid"
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
		logger.Error(err.Error())
		os.Exit(1)
	}

	// set up db
	db, err := database.Connect(config.Vars.DbPath)
	if err != nil {
		os.Exit(1)
	}

	// generate csrf tokens
	queries := models.New(db)
	sessionIDs, err := queries.ListSessionIDs(context.Background())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	err = csrf.Setup(sessionIDs)
	if err != nil {
		os.Exit(1)
	}

	// set up server
	url := server.NewLocalHostUrl(config.Vars.Port)

	router := server.NewRouter()

	router.Setup()

	router.Use(Cache, AddRequestID, AddAuthState)

	router.TemplGet("/", pageViews.NotFound())
	router.Get("/{$}", pages.Index)
	router.Get("/register", pages.Register, GuestGuard)
	router.Get("/login", pages.Login, GuestGuard)

	router.Post("/register", procedures.Register, GuestGuard)
	router.Post("/login", procedures.Login, GuestGuard)
	router.Post("/logout", procedures.Logout, AuthGuard)
	router.Post("/post", procedures.CreatePost, AuthGuard, CSRFGuard)

	s := server.NewServer(router, url)

	serverId := uuid.New("srv")
	s.Serve(url, serverId)
}
