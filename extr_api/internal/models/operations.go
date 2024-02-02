package models

import (
	"alex_gorbunov_exptr_api/internal/lib/api/response"
	"time"
)

type Operation struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	CategoryID string    `json:"category_id"`
	Amount     int       `json:"amount"`
	Currency   string    `json:"currency"`
	Name       string    `json:"name"`
	Comment    string    `json:"comment"`
	Type       string    `json:"type"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateOperationRequest struct {
	UserID     string    `json:"user_id" validate:"required"`
	CategoryID string    `json:"category_id" validate:"required"`
	Amount     int       `json:"amount" validate:"required"`
	Currency   string    `json:"currency" validate:"required"`
	Name       string    `json:"name" validate:"required"`
	Comment    string    `json:"comment"`
	Type       string    `json:"type" validate:"required"`
	CreatedAt  time.Time `json:"created_at" validate:"required"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateOperationResponse struct {
	response.Response
}
