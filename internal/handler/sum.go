package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/domain/models"
	"net/http"
)

func HandleSum(c echo.Context) error {
	var req models.SumRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	res := models.SumResponse{Result: req.FirstNumber + req.SecondNumber}
	return c.JSON(http.StatusOK, res)
}
