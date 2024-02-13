package postgres

import (
	"alex_gorbunov_exptr_api/internal/models"
	"fmt"
)

func (s *Storage) CreateCategory(category *models.Category) error {
	const fn = "storage.postgresql.CreateCategory"

	query := `INSERT INTO categories (user_id, name, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, category.UserID, category.Name, category.Type, category.CreatedAt, category.UpdatedAt)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) UpdateCategory(category *models.Category) error {
	const fn = "storage.postgresql.UpdateCategory"

	query := `UPDATE categories SET user_id = $1, name = $2, type = $3, updated_at = $4 WHERE id = $5`
	_, err := s.db.Exec(query, category.UserID, category.Name, category.Type, category.UpdatedAt, category.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) GetCategories() ([]models.Category, error) {
	const fn = "storage.postgresql.GetCategories"

	query := `SELECT id, user_id, name, type, created_at, updated_at FROM categories`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.UserID, &category.Name, &category.Type, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *Storage) DeleteCategory(id int) error {
	const fn = "storage.postgresql.DeleteCategory"

	query := `DELETE FROM categories WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
