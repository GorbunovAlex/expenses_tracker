package postgres

import (
	"alex_gorbunov_exptr_api/internal/models"
	"fmt"
)

func (s *Storage) CreateOperation(operation models.CreateOperationRequest) error {
	const fn = "storage.postgresql.CreateOperation"

	query := `INSERT INTO operations (user_id, category_id, amount, currency, name, comment, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.Exec(query, operation.UserID, operation.CategoryID, operation.Amount, operation.Currency, operation.Name, operation.Comment, operation.Type, operation.CreatedAt, operation.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) UpdateOperation(operation *models.Operation) error {
	const fn = "storage.postgresql.UpdateOperation"

	query := `UPDATE operations SET user_id = $1, category_id = $2, amount = $3, currency = $4, name = $5, comment = $6, type = $7, updated_at = $8 WHERE id = $9`
	_, err := s.db.Exec(query, operation.UserID, operation.CategoryID, operation.Amount, operation.Currency, operation.Name, operation.Comment, operation.Type, operation.UpdatedAt, operation.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) GetOperationByID(id string) (*models.Operation, error) {
	const fn = "storage.postgresql.GetOperationByID"

	query := `SELECT id, user_id, category_id, amount, currency, name, comment, type, created_at, updated_at FROM operations WHERE id = $1`
	row := s.db.QueryRow(query, id)

	var operation models.Operation
	err := row.Scan(&operation.ID, &operation.UserID, &operation.CategoryID, &operation.Amount, &operation.Currency, &operation.Name, &operation.Comment, &operation.Type, &operation.CreatedAt, &operation.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}

	return &operation, nil
}

func (s *Storage) GetOperationsByUserID(userID string) ([]models.Operation, error) {
	const fn = "storage.postgresql.GetOperationsByUserID"

	query := `SELECT id, user_id, category_id, amount, currency, name, comment, type, created_at, updated_at FROM operations WHERE user_id = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()

	var operations []models.Operation
	for rows.Next() {
		var operation models.Operation
		err := rows.Scan(&operation.ID, &operation.UserID, &operation.CategoryID, &operation.Amount, &operation.Currency, &operation.Name, &operation.Comment, &operation.Type, &operation.CreatedAt, &operation.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		operations = append(operations, operation)
	}

	return operations, nil
}

func (s *Storage) DeleteOperation(id string) error {
	const fn = "storage.postgresql.DeleteOperation"

	query := `DELETE FROM operations WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
