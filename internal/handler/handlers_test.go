package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"github.com/warodan/calculator-rest-api/internal/domain/models"
	"github.com/warodan/calculator-rest-api/internal/middleware"
	"github.com/warodan/calculator-rest-api/internal/storage"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupHandler() *Handler {
	store := storage.NewUserStorage()
	return NewHandler(store)
}

func performRequest(t *testing.T, handlerFunc echo.HandlerFunc, method, path string, body interface{}) *httptest.ResponseRecorder {
	reqBody, err := json.Marshal(body)
	require.NoError(t, err)

	server := echo.New()

	req := httptest.NewRequest(method, path, bytes.NewReader(reqBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	recorder := httptest.NewRecorder()
	c := server.NewContext(req, recorder)

	wrapped := middleware.LoggingMiddleware(slog.New(slog.NewJSONHandler(io.Discard, nil)))(handlerFunc)
	err = wrapped(c)
	require.NoError(t, err)

	return recorder
}

func TestHandleSum_Success(t *testing.T) {
	handler := setupHandler()
	body := models.UserRequest{
		Token:        "4ee5f6e9-50f4-4707-b44e-c1f8a7db70b7",
		FirstNumber:  3,
		SecondNumber: 4,
	}
	recorder := performRequest(t, handler.HandleSum, http.MethodPost, "/sum", body)

	require.Equal(t, http.StatusOK, recorder.Code)

	var resp models.ServerResponse
	err := json.Unmarshal(recorder.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, 7, resp.Result)
}

func TestHandleSum_InvalidUUID(t *testing.T) {
	handler := setupHandler()
	body := models.UserRequest{
		Token:        "not-a-uuid",
		FirstNumber:  3,
		SecondNumber: 4,
	}
	recorder := performRequest(t, handler.HandleSum, http.MethodPost, "/sum", body)

	require.Equal(t, http.StatusBadRequest, recorder.Code)
	require.Contains(t, recorder.Body.String(), "Token is not valid")
}

func TestHandleMultiply_Success(t *testing.T) {
	handler := setupHandler()
	body := models.UserRequest{
		Token:        "4ee5f6e9-50f4-4707-b44e-c1f8a7db70b7",
		FirstNumber:  5,
		SecondNumber: 10,
	}
	recoder := performRequest(t, handler.HandleMultiply, http.MethodPost, "/multiply", body)

	require.Equal(t, http.StatusOK, recoder.Code)

	var resp models.ServerResponse
	err := json.Unmarshal(recoder.Body.Bytes(), &resp)
	require.NoError(t, err)
	require.Equal(t, 50, resp.Result)
}
