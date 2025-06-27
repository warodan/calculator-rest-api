package middleware

import (
	"github.com/labstack/echo/v4"
	"log/slog"
	"time"
)

func LoggingMiddleware(log *slog.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
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
				log.Error("request failed", attrs...)
			} else {
				log.Info("request handled", attrs...)
			}

			return err
		}
	}
}
