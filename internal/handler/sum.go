package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SumRequest struct {
	FirstNumber  int `json:"first_number"`
	SecondNumber int `json:"second_number"`
}

type SumResponse struct {
	Result int `json:"result"`
}

func HandleSum(c echo.Context) error {
	var req SumRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	res := SumResponse{Result: req.FirstNumber + req.SecondNumber}
	return c.JSON(http.StatusOK, res)
}
