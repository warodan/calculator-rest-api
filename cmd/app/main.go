package main

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/handler"
	"github.com/warodan/calculator-rest-api/internal/logger"
	"github.com/warodan/calculator-rest-api/internal/storage"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	server := echo.New()
	log := logger.New()
	userResults := storage.NewUserStorage()
	handlers := handler.NewHandler(log, userResults)

	server.POST("/sum", handlers.HandleSum)
	server.POST("/multiply", handlers.HandleMultiply)

	log.Info("Starting server...")
	go func() {
		if err := server.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("Server failed", "err", err)
			os.Exit(1)
		}
	}()
	log.Info("Server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Info("Graceful shutdown activated")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Error("Graceful shutdown failed", "err", err)
	} else {
		log.Info("Graceful shutdown completed")
	}
}
