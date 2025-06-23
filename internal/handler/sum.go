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
		errResp := map[string]string{"error": "Invalid JSON"}

		h.Log.Error("Invalid JSON",
			slog.Int("status", http.StatusBadRequest),
			slog.Any("response", errResp),
			slog.Any("err", err),
		)

		return c.JSON(http.StatusBadRequest, errResp)
	}

	res := models.SumResponse{Result: req.FirstNumber + req.SecondNumber}
	return c.JSON(http.StatusOK, res)
}
