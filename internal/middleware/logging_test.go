package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
	"io"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoggingMiddleware_InjectsLoggerAndRequestID(t *testing.T) {
	server := echo.New()

	var receivedRequestID string

	next := func(c echo.Context) error {
		rawLogger := c.Get("logger")
		rawReqID := c.Get("request_id")

		require.NotNil(t, rawLogger)
		require.IsType(t, &slog.Logger{}, rawLogger)

		require.NotNil(t, rawReqID)
		require.IsType(t, "", rawReqID)

		receivedRequestID = rawReqID.(string)
		return c.NoContent(http.StatusOK)
	}

	middleware := LoggingMiddleware(slog.New(slog.NewJSONHandler(io.Discard, nil)))

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(""))
	recorder := httptest.NewRecorder()
	c := server.NewContext(req, recorder)

	err := middleware(next)(c)
	require.NoError(t, err)

	_, err = uuid.Parse(receivedRequestID)
	require.NoError(t, err)
}
