package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SumRequest struct {
	A int `json:"a"`
	B int `json:"b"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func main() {
	e := echo.New()

	e.POST("/sum", func(c echo.Context) error {
		var req SumRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		}

		res := SumResponse{Result: req.A + req.B}
		return c.JSON(http.StatusOK, res)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
