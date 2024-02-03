package router

import (
	"alex_gorbunov_exptr_api/internal/config"
	"alex_gorbunov_exptr_api/internal/server/handlers/operations"
	mLogger "alex_gorbunov_exptr_api/internal/server/middleware/logger"
	"alex_gorbunov_exptr_api/internal/storage/postgres"
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth/v5"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

var tokenAuth *jwtauth.JWTAuth

func Router(log *slog.Logger, storage *postgres.Storage) http.Handler {
	router := chi.NewRouter()

	cfg := config.MustLoad()
	tokenAuth = jwtauth.New("HS256", []byte(cfg.HTTPServer.JwtSecret), nil, jwt.WithAcceptableSkew(30*time.Second))

	router.Use(middleware.RequestID)
	router.Use(mLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/api/v1", func(r chi.Router) {
		router.Group(func(r chi.Router) {
			r.Use(jwtauth.Verifier(tokenAuth))
			r.Use(jwtauth.Authenticator(tokenAuth))

			router.Post("/operations", operations.New(log, storage))
		})
	})

	return router
}
