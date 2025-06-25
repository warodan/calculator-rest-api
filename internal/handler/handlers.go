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

func (handler *Handler) HandleSum(echoContext echo.Context) error {
	var req models.UserRequest

	if err := echoContext.Bind(&req); err != nil {
		errResp := map[string]string{"error": "Invalid JSON"}

		handler.Log.Error("Invalid JSON",
			slog.Int("status", http.StatusBadRequest),
			slog.Any("response", errResp),
			slog.Any("err", err),
		)

		return echoContext.JSON(http.StatusBadRequest, errResp)
	}

	res := models.ServerResponse{Result: req.FirstNumber + req.SecondNumber}
	return echoContext.JSON(http.StatusOK, res)
}

func (handler *Handler) HandleMultiply(echoContext echo.Context) error {
	var req models.UserRequest

	if err := echoContext.Bind(&req); err != nil {
		errResp := map[string]string{"error": "Invalid JSON"}

		handler.Log.Error("Invalid JSON",
			slog.Int("status", http.StatusBadRequest),
			slog.Any("response", errResp),
			slog.Any("err", err),
		)

		return echoContext.JSON(http.StatusBadRequest, errResp)
	}

	res := models.ServerResponse{Result: req.FirstNumber * req.SecondNumber}
	return echoContext.JSON(http.StatusOK, res)
}
