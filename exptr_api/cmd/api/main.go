package main

import (
	"alex_gorbunov_exptr_api/internal/config"
	"alex_gorbunov_exptr_api/internal/lib/crons"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/server/router"
	"alex_gorbunov_exptr_api/internal/storage/postgres"
	"log/slog"
	"net/http"
	"os"

	"github.com/robfig/cron/v3"
)

// @title           Swagger EXPTR API
// @version         1.0
// @description     API for expense tracker project.
// @termsOfService  http://swagger.io/terms/

// @contact.name   Alexander Gorbunov
// @contact.email  algor.monte@gmail.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// TODO: Move to global? 3 invokations in app
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

	c := cron.New()

	log.Info("starting server", slog.String("env", cfg.Env))

	storage, err := postgres.NewStorage()
	if err != nil {
		log.Error("failed to init storage", sl.Error(err))
		os.Exit(1)
	}

	log.Info("strating server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router.Router(log, storage),
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("server stopped", sl.Error(err))
	}

	c.AddFunc("@hourly", func() {
		crons.DeleteOutdatedSessions(storage, log)
	})
}
