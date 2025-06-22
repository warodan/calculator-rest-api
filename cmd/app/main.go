package main

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/handler"
	"github.com/warodan/calculator-rest-api/internal/logger"
	"os"
)

func main() {
	logger.Init()

	e := echo.New()

	e.POST("/sum", handler.HandleSum)

	if err := e.Start(":8080"); err != nil {
		logger.Log.Error("Server failed", "err", err)
		os.Exit(1)
	}
}
