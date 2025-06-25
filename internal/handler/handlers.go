package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/domain/models"
	"github.com/warodan/calculator-rest-api/internal/storage"
	"log/slog"
	"net/http"
)

type Handler struct {
	Log     *slog.Logger
	History *storage.History
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

	handler.History.Add(req.Token, storage.Entry{
		FirstNumber:  req.FirstNumber,
		SecondNumber: req.SecondNumber,
		Operation:    "sum",
		Result:       res.Result,
	})

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

	handler.History.Add(req.Token, storage.Entry{
		FirstNumber:  req.FirstNumber,
		SecondNumber: req.SecondNumber,
		Operation:    "multiply",
		Result:       res.Result,
	})

	return echoContext.JSON(http.StatusOK, res)
}
