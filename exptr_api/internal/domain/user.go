package domain

import (
	"github.com/google/uuid"
)

type User struct {
	BaseEntity
	Email    string `json:"email" gorm:"type:varchar;not null;uniqueIndex"`
	Password string `json:"password" gorm:"type:varchar;not null"`
}

func (User) TableName() string {
	return "users"
}

type UserSession struct {
	BaseEntity
	UserID uuid.UUID `json:"user_id" gorm:"type:uuid;not null;index"`
	Token  string    `json:"token" gorm:"type:varchar(255);not null;index"`
	User   User      `json:"-" gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}

func (UserSession) TableName() string {
	return "users_sessions"
}
