package postgres

import (
	"alex_gorbunov_exptr_api/internal/models"
	"fmt"
)

func (s *Storage) CreateCategory(category *models.CategoryRequest) error {
	const fn = "storage.postgresql.CreateCategory"

	query := `INSERT INTO categories (user_id, name, type, created_at, updated_at, color, icon) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := s.db.Exec(query, category.UserID, category.Name, category.Type, category.CreatedAt, category.UpdatedAt, category.Color, category.Icon)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) UpdateCategory(category *models.Category) error {
	const fn = "storage.postgresql.UpdateCategory"

	query := `UPDATE categories SET user_id = $1, name = $2, type = $3, updated_at = $4, color = $5, icon = $6 WHERE id = $7`
	_, err := s.db.Exec(query, category.UserID, category.Name, category.Type, category.UpdatedAt, category.Color, category.Icon, category.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (s *Storage) GetCategories(userID int) ([]models.Category, error) {
	const fn = "storage.postgresql.GetCategories"

	query := `SELECT id, user_id, name, type, created_at, updated_at, color, icon FROM categories WHERE user_id = $1`
	rows, err := s.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", fn, err)
	}
	defer rows.Close()

	var categories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.UserID, &category.Name, &category.Type, &category.CreatedAt, &category.UpdatedAt, &category.Color, &category.Icon)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", fn, err)
		}
		categories = append(categories, category)
	}

	return categories, nil
}

func (s *Storage) DeleteCategory(id int) error {
	const fn = "storage.postgresql.DeleteCategory"

	query := `SELECT * FROM categories WHERE id = $1`
	fmt.Println("id", id)
	row := s.db.QueryRow(query, id).Scan()
	if row != nil {
		return fmt.Errorf("%s: %w", fn, row)
	}

	query = `DELETE FROM categories WHERE id = $1`
	_, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
