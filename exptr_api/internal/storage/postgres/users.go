package postgres

import (
	"alex_gorbunov_exptr_api/internal/models"
	"fmt"
)

func (s *Storage) CreateUser(user *models.User) error {
	const fn = "storage.postgresql.CreateUser"

	query := `INSERT INTO users (email, password, created_at, updated_at) VALUES ($1, $2, $3, $4)`
	_, err := s.db.Exec(query, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) GetUserByEmail(email string) (*models.User, error) {
	const fn = "storage.postgresql.GetUserByEmail"

	query := `SELECT id, email, password, created_at, updated_at FROM users WHERE email = $1`
	user := &models.User{}
	err := s.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return user, nil
}
