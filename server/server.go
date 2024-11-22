package server

import (
	"log/slog"
	"net/http"
	"time"

	l "github.com/peterszarvas94/goat/logger"
)

type Server struct {
	server *http.Server
}

func NewServer(m *Mux) *Server {
	httpServer := &http.Server{
		Addr:         m.url,
		Handler:      m.mux,
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

	l.Logger.Info("Server is started: %s", slog.String("url", s.server.Addr), slog.String("id", id))

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		l.Logger.Error(err.Error())
		panic(err.Error())
	}

	// Wait for the graceful shutdown to complete
	<-done

	l.Logger.Info("Graceful shutdown complete")
}
