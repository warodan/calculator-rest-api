package main

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/handler"
	"github.com/warodan/calculator-rest-api/internal/logger"
	"github.com/warodan/calculator-rest-api/internal/storage"
	"os"
)

func main() {
	log := logger.New()
	history := storage.NewHistory()

	server := echo.New()

	handlers := handler.Handler{
		Log:     log,
		History: history,
	}

	server.POST("/sum", handlers.HandleSum)
	server.POST("/multiply", handlers.HandleMultiply)
	log.Info("Starting server")

	if err := server.Start(":8080"); err != nil {
		log.Error("Server failed", "err", err)
		os.Exit(1)
	}
}
