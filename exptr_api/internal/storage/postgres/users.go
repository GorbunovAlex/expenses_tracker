package postgres

import (
	"errors"
	"fmt"
	"time"

	"alex_gorbunov_exptr_api/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Storage) CreateUser(user *domain.User) error {
	const fn = "storage.postgresql.CreateUser"

	result := s.db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}

func (s *Storage) GetUserByEmail(email string) (*domain.User, error) {
	const fn = "storage.postgresql.GetUserByEmail"

	var user domain.User
	result := s.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: user not found", fn)
		}
		return nil, fmt.Errorf("%s: %w", fn, result.Error)
	}

	return &user, nil
}

func (s *Storage) SetUserSession(userID uuid.UUID, token string) error {
	const fn = "storage.postgresql.SetUserSession"

	var session domain.UserSession
	result := s.db.Where("user_id = ?", userID).First(&session)

	if result.Error == nil {
		// Session exists, update it
		session.Token = token
		session.BaseEntity.UpdatedAt = time.Now()
		if err := s.db.Save(&session).Error; err != nil {
			return fmt.Errorf("%s: %w", fn, err)
		}
		return nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	// Session doesn't exist, create new one
	newSession := domain.UserSession{
		UserID: userID,
		Token:  token,
	}

	if err := s.db.Create(&newSession).Error; err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) UpdateUserSession(userID uuid.UUID, token string) error {
	const fn = "storage.postgresql.UpdateUserSession"

	result := s.db.Model(&domain.UserSession{}).Where("user_id = ?", userID).Update("token", token)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}

func (s *Storage) GetUserIDByToken(token string) (*string, error) {
	const fn = "storage.postgresql.GetUserIDByToken"

	var session domain.UserSession
	result := s.db.Where("token = ?", token).First(&session)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: session not found", fn)
		}
		return nil, fmt.Errorf("%s: %w", fn, result.Error)
	}

	userID := session.UserID.String()
	return &userID, nil
}

func (s *Storage) GetUserSession(userID uuid.UUID) (*uuid.UUID, error) {
	const fn = "storage.postgresql.GetUserSession"

	var session domain.UserSession
	result := s.db.Where("user_id = ?", userID).First(&session)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: session not found", fn)
		}
		return nil, fmt.Errorf("%s: %w", fn, result.Error)
	}

	return &session.ID, nil
}

func (s *Storage) DeleteUserSession(userID uuid.UUID) error {
	const fn = "storage.postgresql.DeleteUserSession"

	result := s.db.Where("user_id = ?", userID).Delete(&domain.UserSession{})
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}

func (s *Storage) DeleteOutdatedSessions() error {
	const fn = "storage.postgresql.DeleteOutdatedSessions"

	cutoff := time.Now().Add(-time.Hour * 1)
	result := s.db.Where("created_at < ?", cutoff).Delete(&domain.UserSession{})
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}
