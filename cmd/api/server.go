package main

import (
	"Go-Test/cmd/api/route"
	"Go-Test/pkg/storage"
	"context"
	"go-micro.dev/v4/logger"
	"net/http"
	"time"
)

type HTTPServer struct {
	httpServer *http.Server
	db         *storage.Storage
}

func NewHTTPServer(port string, db *storage.Storage) *HTTPServer {
	return &HTTPServer{
		httpServer: &http.Server{
			Addr:         ":" + port,
			Handler:      route.NewRouter(db),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
			IdleTimeout:  60 * time.Second,
		},
		db: db,
	}
}

func (s *HTTPServer) Start() {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log(logger.FatalLevel, "start server error")
		}
	}()
}

func (s *HTTPServer) Stop(ctx context.Context) {
	logger.Log(logger.InfoLevel, "shutting down ...")

	if s.db != nil {
		s.db.Close()
		logger.Log(logger.InfoLevel, "database connection closed")
	}

	if err := s.httpServer.Shutdown(ctx); err != nil {
		logger.Log(logger.ErrorLevel, "error while shutting down server")
	} else {
		logger.Log(logger.InfoLevel, "server stopped gracefully")
	}
}
