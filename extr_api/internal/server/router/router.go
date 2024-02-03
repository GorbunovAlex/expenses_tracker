package router

import (
	"alex_gorbunov_exptr_api/internal/server/handlers/operations"
	"alex_gorbunov_exptr_api/internal/server/handlers/users"
	mLogger "alex_gorbunov_exptr_api/internal/server/middleware/logger"
	"alex_gorbunov_exptr_api/internal/server/middleware/token"
	"alex_gorbunov_exptr_api/internal/storage/postgres"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Router(log *slog.Logger, storage *postgres.Storage) http.Handler {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/api/v1", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Use(token.TokenValidationMiddleware)

			r.Post("/operations/new", operations.New(log, storage))
		})

		r.Post("/users/signup", users.Signup(log, storage))
		r.Post("/users/login", users.Login(log, storage))
	})

	return router
}
