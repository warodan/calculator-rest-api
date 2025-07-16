package main

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	_ "github.com/warodan/calculator-rest-api/docs"
	"github.com/warodan/calculator-rest-api/internal/config"
	"github.com/warodan/calculator-rest-api/internal/handler"
	"github.com/warodan/calculator-rest-api/internal/logger"
	"github.com/warodan/calculator-rest-api/internal/middleware"
	"github.com/warodan/calculator-rest-api/internal/storage"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// @title Calculator API
// @version 1.0
// @description This API allows you to add and multiply two numbers and view calculation history.
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	cfg := config.Load()
	if err := cfg.Validate(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	server := echo.New()
	server.GET("/swagger/*", echoSwagger.WrapHandler)

	log := logger.New(cfg)
	userResults := storage.NewUserStorage()
	handlers := handler.NewHandler(userResults)

	server.Use(middleware.LoggingMiddleware(log))

	server.POST("/sum", handlers.HandleSum)
	server.POST("/multiply", handlers.HandleMultiply)

	log.Info("Starting server...")
	go func() {
		if err := server.Start(":" + cfg.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
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
