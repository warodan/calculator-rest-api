package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/warodan/calculator-rest-api/internal/domain/models"
	"github.com/warodan/calculator-rest-api/internal/domain/operations"
	"github.com/warodan/calculator-rest-api/internal/storage"
	"log/slog"
	"net/http"
)

type Handler struct {
	UserResults *storage.UserResults
}

func NewHandler(userResults *storage.UserResults) *Handler {
	return &Handler{
		UserResults: userResults,
	}
}

func (handler *Handler) handleOperation(c echo.Context, op string) error {
	log := c.Get("logger").(*slog.Logger)

	opFunc, ok := operations.Registry[op]
	if !ok {
		errResp := map[string]string{"error": "Unsupported operation"}
		log.Error("Unknown operation",
			slog.Int("status", http.StatusBadRequest),
			slog.String("op", op),
		)
		return c.JSON(http.StatusBadRequest, errResp)
	}

	var req models.UserRequest

	if err := c.Bind(&req); err != nil {
		log.Error("Invalid JSON",
			slog.Int("status", http.StatusBadRequest),
			slog.Any("err", err),
		)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
	}

	if _, err := uuid.Parse(req.Token); err != nil {
		log.Error("Token is not valid UUID",
			slog.Any("err", err),
		)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Token is not valid"})
	}

	result := opFunc(req.FirstNumber, req.SecondNumber)

	if err := handler.UserResults.Add(req.Token, storage.Entry{
		FirstNumber:  req.FirstNumber,
		SecondNumber: req.SecondNumber,
		Operation:    op,
		Result:       result,
	}); err != nil {
		log.Error("Error writing to the map",
			slog.Any("method", "Add"),
			slog.Any("err", err),
		)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal server error"})
	}

	return c.JSON(http.StatusOK, models.ServerResponse{Result: result})
}

// HandleSum godoc
// @Summary Add two numbers
// @Description Returns the sum of two int numbers provided in the request
// @Tags calculator
// @Accept json
// @Produce json
// @Param input body models.UserRequest true "Input data"
// @Success 200 {object} models.ServerResponse
// @Failure 400 {object} map[string]string
// @Router /sum [post]
func (handler *Handler) HandleSum(c echo.Context) error {
	return handler.handleOperation(c, operations.OpSum)
}

// HandleMultiply godoc
// @Summary Multiply two numbers
// @Description Returns the product of two int numbers provided in the request
// @Tags calculator
// @Accept json
// @Produce json
// @Param input body models.UserRequest true "Input data"
// @Success 200 {object} models.ServerResponse
// @Failure 400 {object} map[string]string
// @Router /multiply [post]
func (handler *Handler) HandleMultiply(c echo.Context) error {
	return handler.handleOperation(c, operations.OpMultiply)
}
