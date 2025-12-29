package postgres

import (
	"errors"
	"fmt"

	"alex_gorbunov_exptr_api/internal/domain"
	"alex_gorbunov_exptr_api/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (s *Storage) CreateOperation(operation models.OperationRequest) error {
	const fn = "storage.postgresql.CreateOperation"

	op := domain.Operation{
		UserID:     operation.UserID,
		CategoryID: operation.CategoryID,
		Amount:     operation.Amount,
		Currency:   operation.Currency,
		Name:       operation.Name,
		Comment:    operation.Comment,
		Type:       operation.Type,
	}

	result := s.db.Create(&op)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}

func (s *Storage) UpdateOperation(id uuid.UUID, operation *models.OperationRequest) error {
	const fn = "storage.postgresql.UpdateOperation"

	result := s.db.Model(&domain.Operation{}).Where("id = ?", id).Updates(map[string]interface{}{
		"category_id": operation.CategoryID,
		"amount":      operation.Amount,
		"currency":    operation.Currency,
		"name":        operation.Name,
		"comment":     operation.Comment,
		"type":        operation.Type,
	})

	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%s: operation not found", fn)
	}

	return nil
}

func (s *Storage) GetOperationByID(id uuid.UUID) (*domain.Operation, error) {
	const fn = "storage.postgresql.GetOperationByID"

	var operation domain.Operation
	result := s.db.Where("id = ?", id).First(&operation)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("%s: operation not found", fn)
		}
		return nil, fmt.Errorf("%s: %w", fn, result.Error)
	}

	return &operation, nil
}

func (s *Storage) GetOperationsByUserID(userID uuid.UUID) ([]domain.Operation, error) {
	const fn = "storage.postgresql.GetOperationsByUserID"

	var operations []domain.Operation
	result := s.db.Where("user_id = ?", userID).Find(&operations)
	if result.Error != nil {
		return nil, fmt.Errorf("%s: %w", fn, result.Error)
	}

	return operations, nil
}

func (s *Storage) DeleteOperation(id uuid.UUID) error {
	const fn = "storage.postgresql.DeleteOperation"

	// First check if operation exists
	var operation domain.Operation
	result := s.db.Where("id = ?", id).First(&operation)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return fmt.Errorf("%s: operation not found", fn)
		}
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	// Delete the operation (soft delete due to gorm.DeletedAt in BaseEntity)
	result = s.db.Delete(&operation)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", fn, result.Error)
	}

	return nil
}
