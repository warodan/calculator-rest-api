package main

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/handler"
	"github.com/warodan/calculator-rest-api/internal/logger"
	"os"
)

func main() {
	log := logger.New()

	e := echo.New()

	h := handler.Handler{Log: log}

	e.POST("/sum", h.HandleSum)
	log.Info("Starting server")

	if err := e.Start(":8080"); err != nil {
		log.Error("Server failed", "err", err)
		os.Exit(1)
	}
}
