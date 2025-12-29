package domain

import (
	"github.com/google/uuid"
)

type Category struct {
	BaseEntity
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;index"`
	Name   string    `json:"name" gorm:"type:varchar(255);not null"`
	Type   string    `json:"type" gorm:"type:varchar(255);not null"`
	Color  string    `json:"color" gorm:"type:varchar(255)"`
	Icon   string    `json:"icon" gorm:"type:varchar(255)"`
}

func (Category) TableName() string {
	return "categories"
}
