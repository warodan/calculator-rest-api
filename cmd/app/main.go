package main

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/handler"
)

func main() {
	e := echo.New()

	e.POST("/sum", handler.HandleSum)

	e.Logger.Fatal(e.Start(":8080"))
}
