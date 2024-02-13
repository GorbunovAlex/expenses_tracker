package postgres

import (
	"alex_gorbunov_exptr_api/internal/models"
	"fmt"
	"time"
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

func (s *Storage) SetUserSession(userID int, token string) error {
	const fn = "storage.postgresql.setUserSession"

	_, err := s.GetUserSession(userID)
	if err == nil {
		return fmt.Errorf("%s: user already has a session", fn)
	}

	query := `INSERT INTO user_sessions (user_id, created_date, token) VALUES ($1, $2, $3)`
	_, err = s.db.Exec(query, userID, time.Now(), token)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) UpdateUserSession(userID int, token string) error {
	const fn = "storage.postgresql.updateUserSession"

	query := `UPDATE user_sessions SET token = $1 WHERE user_id = $2`
	_, err := s.db.Exec(query, token, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) GetUserIDByToken(token string) (int, error) {
	const fn = "storage.postgresql.getUserSessionByToken"

	query := `SELECT user_id FROM user_sessions WHERE token = $1`
	var userID int
	err := s.db.QueryRow(query, token).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return userID, nil
}

func (s *Storage) GetUserSession(userID int) (int, error) {
	const fn = "storage.postgresql.getUserSession"

	query := `SELECT id FROM user_sessions WHERE user_id = $1`
	var sessionID int
	err := s.db.QueryRow(query, userID).Scan(&sessionID)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", fn, err)
	}

	return sessionID, nil
}

func (s *Storage) DeleteUserSession(userID int) error {
	const fn = "storage.postgresql.deleteUserSession"

	query := `DELETE FROM user_sessions WHERE user_id = $1`
	_, err := s.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
