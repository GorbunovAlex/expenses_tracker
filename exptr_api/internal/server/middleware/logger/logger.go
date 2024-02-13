package logger

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func New(log *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log := log.With(
			slog.String("component", "middleware/logger"),
		)

		log.Info("logger middleware enabled")

		entry := log.With(
			slog.String("method", c.Request.Method),
			slog.String("path", c.Request.RequestURI),
			slog.String("remote_addr", c.Request.RemoteAddr),
			slog.String("user_agent", c.Request.UserAgent()),
			slog.String("request_id", c.Request.Header.Get("X-Request-Id")),
		)

		ww := c.Writer

		t1 := time.Now()
		defer func() {
			entry.Info("request",
				slog.Int("status", ww.Status()),
				slog.Int("bytes", ww.Size()),
				slog.String("duration", time.Since(t1).String()),
			)
		}()

		c.Next()
	}
}
