package main

import (
	"Go-Test/pkg/storage"
	"context"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8088"
	}

	db, err := storage.NewStorage()
	if err != nil {
		logger.Log(logger.FatalLevel, "connect database failed")
	}

	httpServer := NewHTTPServer(port, db)

	service := micro.NewService(
		micro.Name("go.micro.httpserver"),
		micro.Version("1"),
		micro.BeforeStart(func() error {
			httpServer.Start()
			return nil
		}),
		micro.BeforeStop(func() error {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			httpServer.Stop(ctx)
			return nil
		}),
	)

	if err := service.Run(); err != nil {
		logger.Log(logger.FatalLevel, "service startup failed")
	}
}
