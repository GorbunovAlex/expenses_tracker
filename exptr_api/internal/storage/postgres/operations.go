package postgres

import (
	"alex_gorbunov_exptr_api/internal/models"
	"fmt"
)

func (s *Storage) CreateOperation(operation models.OperationRequest) error {
	const fn = "storage.postgresql.CreateOperation"

	query := `INSERT INTO operations (user_id, category_id, amount, currency, name, comment, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := s.db.Exec(query, operation.UserID, operation.CategoryID, operation.Amount, operation.Currency, operation.Name, operation.Comment, operation.Type, operation.CreatedAt, operation.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) UpdateOperation(id int, operation *models.OperationRequest) error {
	const fn = "storage.postgresql.UpdateOperation"

	query := `UPDATE operations SET category_id = $1, amount = $2, currency = $3, name = $4, comment = $5, type = $6, updated_at = $7 WHERE id = $8`
	_, err := s.db.Exec(query, operation.CategoryID, operation.Amount, operation.Currency, operation.Name, operation.Comment, operation.Type, operation.UpdatedAt, id)
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

func (s *Storage) GetOperationsByUserID(userID int) ([]models.Operation, error) {
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
