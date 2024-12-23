package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/peterszarvas94/goat/logger"
)

type Server struct {
	server *http.Server
}

func NewServer(router *Router, url string) *Server {
	httpServer := &http.Server{
		Addr:         url,
		Handler:      &router.Mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		server: httpServer,
	}
}

func (s *Server) Serve(url string, id string) {
	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(s.server, done)

	logger.Info("Server is started", slog.String("url", s.server.Addr), slog.String("id", id))

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		logger.Error(err.Error())
		panic(err.Error())
	}

	// Wait for the graceful shutdown to complete
	<-done

	logger.Info("Graceful shutdown complete")
}
