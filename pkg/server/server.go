package server

import (
	"context"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
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

	slog.Info("Server is started", slog.String("url", s.server.Addr), slog.String("id", id))

	err := s.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error(err.Error())
		panic(err.Error())
	}

	// Wait for the graceful shutdown to complete
	<-done

	slog.Info("Graceful shutdown complete")
}

func gracefulShutdown(server *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	slog.Info("Shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Server forced to shutdown", slog.String("msg", err.Error()))
	}

	slog.Debug("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}
