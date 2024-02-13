package token

import (
	"alex_gorbunov_exptr_api/pkg/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TokenValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, ok := c.Request.Header["Bearer"]; !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := c.Request.Header.Get("Bearer")
		check, err := jwt.ValidateToken(token)

		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		if !check {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Writer.WriteHeader(http.StatusOK)
		c.Next()
	}
}
