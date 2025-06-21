package main

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/handler"
)

type SumRequest struct {
	FirstNumber  int `json:"first_number"`
	SecondNumber int `json:"second_number"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func main() {
	e := echo.New()

	e.POST("/sum", handler.HandleSum)

	e.Logger.Fatal(e.Start(":8080"))
}
