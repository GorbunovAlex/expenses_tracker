package postgres

import (
	"errors"
	"fmt"

	"alex_gorbunov_exptr_api/internal/domain"
	"alex_gorbunov_exptr_api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Storage) CreateCategory(category *models.CategoryRequest) error {
	const fn = "storage.postgresql.CreateCategory"

	cat := domain.Category{
		UserID: uuid.MustParse(category.UserID),
		Name:   category.Name,
		Type:   category.Type,
		Color:  category.Color,
		Icon:   category.Icon,
	}

	result := s.db.Create(&cat)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}

func (s *Storage) UpdateCategory(category *domain.Category) error {
	const fn = "storage.postgresql.UpdateCategory"

	result := s.db.Model(&domain.Category{}).Where("id = ?", category.ID).Updates(map[string]interface{}{
		"user_id": category.UserID,
		"name":    category.Name,
		"type":    category.Type,
		"color":   category.Color,
		"icon":    category.Icon,
	})

	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%s: category not found", fn)
	}

	return nil
}

func (s *Storage) GetCategories(userID uuid.UUID) ([]domain.Category, error) {
	const fn = "storage.postgresql.GetCategories"

	var categories []domain.Category
	result := s.db.Where("user_id = ?", userID).Find(&categories)
	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", fn, result.Error)
	}

	return categories, nil
}

func (s *Storage) GetCategoryByID(id uuid.UUID) (*domain.Category, error) {
	const fn = "storage.postgresql.GetCategoryByID"

	var category domain.Category
	result := s.db.Where("id = ?", id).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: category not found", fn)
		}
		return nil, fmt.Errorf("%s: %w", fn, result.Error)
	}

	return &category, nil
}

func (s *Storage) DeleteCategory(id uuid.UUID) error {
	const fn = "storage.postgresql.DeleteCategory"

	// First check if category exists
	var category domain.Category
	result := s.db.Where("id = ?", id).First(&category)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%s: category not found", fn)
		}
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	// Delete the category (soft delete due to gorm.DeletedAt in BaseEntity)
	result = s.db.Delete(&category)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}
