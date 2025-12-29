package models

import (
	"time"

	"alex_gorbunov_exptr_api/internal/domain"
	"alex_gorbunov_exptr_api/internal/lib/api/response"
)

type CategoryRequest struct {
	UserID    string    `json:"user_id"`
	Name      string    `json:"name"`
	Type      string    `json:"type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Color     string    `json:"color"`
	Icon      string    `json:"icon"`
}

type CategoryResponse struct {
	response.Response
}

type GetCategoriesResponse struct {
	response.Response
	Categories []domain.Category `json:"categories"`
}
