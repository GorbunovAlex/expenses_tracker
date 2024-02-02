package postgres

import (
	"alex_gorbunov_exptr_api/internal/models"
	"alex_gorbunov_exptr_api/pkg/hasher"
	"fmt"
)

func (s *Storage) CreateUser(user *models.User) error {
	const fn = "storage.postgresql.CreateUser"

	passwordHash, err := hasher.HashPassword(user.Password)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	query := `INSERT INTO users (email, phone, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err = s.db.Exec(query, user.Email, user.Phone, passwordHash, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
