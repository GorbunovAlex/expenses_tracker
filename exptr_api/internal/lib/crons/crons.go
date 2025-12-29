package crons

import (
	"alex_gorbunov_exptr_api/internal/lib/logger/sl"
	"alex_gorbunov_exptr_api/internal/storage/postgres"
	"log/slog"
)

func DeleteOutdatedSessions(storage *postgres.Storage, log *slog.Logger) {
	const op = "cron.DeleteOutdatedSessions"

	log = log.With(slog.String("op", op))

	log.Info("deleting outdated sessions")
	err := storage.DeleteOutdatedSessions()
	if err != nil {
		log.Error("failed to delete outdated sessions", sl.Error(err))
	}
}
