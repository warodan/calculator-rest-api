package middleware

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log/slog"
	"time"
)

func LoggingMiddleware(log *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqID := uuid.NewString()
			c.Set("request_id", reqID)

			reqLog := log.With("request_id", reqID)
			c.Set("logger", reqLog)

			start := time.Now()

			err := next(c)

			stop := time.Now()
			method := c.Request().Method
			path := c.Path()
			status := c.Response().Status
			ip := c.RealIP()
			latency := stop.Sub(start)

			attrs := []any{
				slog.String("method", method),
				slog.String("path", path),
				slog.Int("status", status),
				slog.String("ip", ip),
				slog.Duration("latency", latency),
			}

			if err != nil {
				attrs = append(attrs, slog.Any("error", err))
				reqLog.Error("request failed", attrs...)
			} else {
				reqLog.Info("request handled", attrs...)
			}

			return err
		}
	}
}
