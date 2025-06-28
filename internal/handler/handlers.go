package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/domain/models"
	"github.com/warodan/calculator-rest-api/internal/domain/operations"
	"github.com/warodan/calculator-rest-api/internal/storage"
	"log/slog"
	"net/http"
)

type Handler struct {
	Log         *slog.Logger
	UserResults *storage.UserResults
}

func NewHandler(log *slog.Logger, userResults *storage.UserResults) *Handler {
	return &Handler{
		Log:         log,
		UserResults: userResults,
	}
}

func (handler *Handler) handleOperation(c echo.Context, op string) error {
	opFunc, ok := operations.Registry[op]
	if !ok {
		errResp := map[string]string{"error": "Unsupported operation"}
		handler.Log.Error("Unknown operation",
			slog.Int("status", http.StatusBadRequest),
			slog.String("op", op),
		)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	var req models.UserRequest

	if err := c.Bind(&req); err != nil {
		errResp := map[string]string{"error": "Invalid JSON"}
		handler.Log.Error("Invalid JSON",
			slog.Int("status", http.StatusBadRequest),
			slog.Any("err", err),
		)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	result := opFunc(req.FirstNumber, req.SecondNumber)

	if err := handler.UserResults.Add(req.Token, storage.Entry{
		FirstNumber:  req.FirstNumber,
		SecondNumber: req.SecondNumber,
		Operation:    op,
		Result:       result,
	}); err != nil {
		handler.Log.Error("failed to store result",
			slog.String("method", "Add"),
			slog.String("token", req.Token),
			slog.Any("err", err),
		)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, models.ServerResponse{Result: result})
}

func (handler *Handler) HandleSum(c echo.Context) error {
	return handler.handleOperation(c, operations.OpSum)
}

func (handler *Handler) HandleMultiply(c echo.Context) error {
	return handler.handleOperation(c, operations.OpMultiply)
}
