package token

import (
	"log/slog"
	"net/http"
	"strings"

	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"alex_gorbunov_exptr_api/pkg/jwt"

	"github.com/gin-gonic/gin"
)

const UserIDKey = "userID"

type TokenStorage interface {
	GetUserIDByToken(token string) (*string, error)
}

func TokenValidationMiddleware(log *slog.Logger, storage TokenStorage) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")

		if authHeader == "" {
			log.Debug("middleware: Authorization header missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("missing Authorization header"))
			return
		}

		// Extract the actual token by removing the "Bearer " prefix
		// This handles cases where the prefix might be missing or capitalized differently
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Optional: Check if the trimming actually happened
		if token == authHeader {
			log.Debug("middleware: Bearer prefix missing")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("invalid token format"))
			return
		}

		log.Debug("middleware: token extracted", slog.String("token", token))

		check, err := jwt.ValidateToken(token)
		if err != nil {
			log.Debug("middleware: token validation error", slog.String("error", err.Error()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("invalid token"))
			return
		}

		if !check {
			log.Debug("middleware: token validation failed")
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("invalid token"))
			return
		}

		log.Debug("middleware: token validated successfully")

		userIDStr, err := storage.GetUserIDByToken(token)
		if err != nil {
			log.Error("middleware: failed to get user id by token", slog.String("error", err.Error()))
			c.AbortWithStatusJSON(http.StatusUnauthorized, response.Error("invalid session"))
			return
		}

		log.Debug("middleware: user ID retrieved", slog.String("userID", *userIDStr))

		c.Set(UserIDKey, *userIDStr)

		log.Debug("middleware: calling next handler")
		c.Next()
		log.Debug("middleware: handler completed")
	}
}

func GetUserIDFromContext(c *gin.Context) (string, bool) {
	userID, exists := c.Get(UserIDKey)
	if !exists {
		return "", false
	}
	userIDStr, ok := userID.(string)
	return userIDStr, ok
}
