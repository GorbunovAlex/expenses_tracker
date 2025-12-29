package postgres

import (
	"fmt"

	"alex_gorbunov_exptr_api/internal/config"
	"alex_gorbunov_exptr_api/internal/domain"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	db *gorm.DB
}

func NewStorage() (*Storage, error) {
	const fn = "storage.postgresql.NewStorage"

	cfg := config.MustLoad()

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.Name)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &Storage{db: db}, nil
}

// NewStorageWithAutoMigrate creates a new storage and runs auto-migration for all models.
// Use this for development or when you don't want to manage migrations manually.
func NewStorageWithAutoMigrate() (*Storage, error) {
	const fn = "storage.postgresql.NewStorageWithAutoMigrate"

	storage, err := NewStorage()
	if err != nil {
		return nil, err
	}

	err = storage.db.AutoMigrate(
		&domain.User{},
		&domain.UserSession{},
		&domain.Category{},
		&domain.Operation{},
	)
	if err != nil {
		return nil, fmt.Errorf("%s: failed to auto migrate: %w", fn, err)
	}

	return storage, nil
}

func (s *Storage) DB() *gorm.DB {
	return s.db
}

// Close closes the underlying database connection
func (s *Storage) Close() error {
	sqlDB, err := s.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
