package models

import (
	"time"

	"alex_gorbunov_exptr_api/internal/domain"
	"alex_gorbunov_exptr_api/internal/lib/api/response"

	"github.com/google/uuid"
)

type OperationRequest struct {
	UserID     uuid.UUID `json:"user_id" validate:"required"`
	CategoryID uuid.UUID `json:"category_id" validate:"required"`
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

type GetOperationsByUserIDResponse struct {
	response.Response
	Operations []domain.Operation `json:"operations"`
}

type UpdateOperationResponse struct {
	response.Response
}
