package router

import (
	_ "alex_gorbunov_exptr_api/docs"
	"alex_gorbunov_exptr_api/internal/server/handlers/operations"
	"alex_gorbunov_exptr_api/internal/server/handlers/users"
	mLogger "alex_gorbunov_exptr_api/internal/server/middleware/logger"
	"alex_gorbunov_exptr_api/internal/server/middleware/token"
	"alex_gorbunov_exptr_api/internal/storage/postgres"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(log *slog.Logger, storage *postgres.Storage) http.Handler {
	router := gin.Default()

	router.Use(mLogger.New(log))
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/")
		auth.Use(token.TokenValidationMiddleware())
		{
			auth.POST("/operations/new", operations.New(log, storage))
		}
		v1.POST("/users/signup", users.Signup(log, storage))
		v1.POST("/users/login", users.Login(log, storage))
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
