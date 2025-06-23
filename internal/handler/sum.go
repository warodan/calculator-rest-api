package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/domain/models"
	"log/slog"
	"net/http"
)

type Handler struct {
	Log *slog.Logger
}

func (h *Handler) HandleSum(c echo.Context) error {
	var req models.SumRequest

	if err := c.Bind(&req); err != nil {
		h.Log.Error("Invalid JSON", "err", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	res := models.SumResponse{Result: req.FirstNumber + req.SecondNumber}
	return c.JSON(http.StatusOK, res)
}
