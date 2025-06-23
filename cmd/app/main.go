package main

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/handler"
	"github.com/warodan/calculator-rest-api/internal/logger"
	"os"
)

func main() {
	log := logger.New()

	server := echo.New()

	handlers := handler.Handler{Log: log}

	server.POST("/sum", handlers.HandleSum)
	log.Info("Starting server")

	if err := server.Start(":8080"); err != nil {
		log.Error("Server failed", "err", err)
		os.Exit(1)
	}
}
