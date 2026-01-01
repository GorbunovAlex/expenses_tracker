package router

import (
	"log/slog"
	"net/http"

	_ "alex_gorbunov_exptr_api/docs"
	"alex_gorbunov_exptr_api/internal/server/handlers/categories"
	"alex_gorbunov_exptr_api/internal/server/handlers/operations"
	"alex_gorbunov_exptr_api/internal/server/handlers/users"
	mLogger "alex_gorbunov_exptr_api/internal/server/middleware/logger"
	"alex_gorbunov_exptr_api/internal/server/middleware/token"
	"alex_gorbunov_exptr_api/internal/storage/postgres"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router(log *slog.Logger, storage *postgres.Storage) http.Handler {
	router := gin.Default()

	router.Use(mLogger.New(log))
	router.Use(gin.Recovery())
	router.Use(cors.Default())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/")
		auth.Use(token.TokenValidationMiddleware(log, storage))
		{
			auth.POST("/operations/new", operations.New(log, storage))
			auth.GET("/operations", operations.GetAll(log, storage))
			auth.PUT("/operations/:id", operations.Update(log, storage))
			auth.DELETE("/operations/:id", operations.Delete(log, storage))

			auth.GET("/categories", categories.GetAll(log, storage))
			auth.POST("/categories/new", categories.New(log, storage))
			auth.PUT("/categories/:id", categories.Update(log, storage))
			auth.DELETE("/categories/:id", categories.Delete(log, storage))
		}
		v1.POST("/users/signup", users.Signup(log, storage))
		v1.POST("/users/login", users.Login(log, storage))
	}

	return router
}
