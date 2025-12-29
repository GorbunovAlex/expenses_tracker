package domain

import (
	"github.com/google/uuid"
)

type Operation struct {
	BaseEntity
	UserID     uuid.UUID `json:"user_id" gorm:"type:uuid;index"`
	CategoryID uuid.UUID `json:"category_id" gorm:"type:uuid;not null;index"`
	Amount     int       `json:"amount" gorm:"type:decimal(19,4);not null"`
	Currency   string    `json:"currency" gorm:"type:varchar(10);not null"`
	Name       string    `json:"name" gorm:"type:varchar(255);not null"`
	Comment    string    `json:"comment" gorm:"type:text"`
	Type       string    `json:"type" gorm:"type:varchar(255)"`
}

func (Operation) TableName() string {
	return "operations"
}
