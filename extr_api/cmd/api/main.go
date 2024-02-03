package main

import (
	"alex_gorbunov_exptr_api/internal/config"
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/server/router"
	"alex_gorbunov_exptr_api/internal/storage/postgres"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	// TODO: Move to global? 3 invokations in app
	cfg := config.MustLoad()

	log := sl.SetupLogger(cfg.Env)

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
}
