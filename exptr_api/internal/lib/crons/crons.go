package crons

import (
	"alex_gorbunov_exptr_api/internal/storage/postgres"
	"log/slog"
)

func DeleteOutdatedSessions(storage *postgres.Storage, log *slog.Logger) {
	const op = "cron.DeleteOutdatedSessions"

	log = log.With(slog.String("op", op))

	err := storage.DeleteOutdatedSessions()
	if err != nil {
		log.Error("failed to delete outdated sessions", err)
	}
}
