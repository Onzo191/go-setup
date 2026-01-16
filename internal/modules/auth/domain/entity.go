package domain

import (
	model "go-setup/internal/core/database"
	entity "go-setup/internal/modules/user/domain"
)

type UserCredential struct {
	model.BaseModel
	UserID       uint        `gorm:"not null"`
	PasswordHash string      `gorm:"not null"`
	User         entity.User `gorm:"foreignKey:UserID;"`
}

type UserSession struct {
	model.BaseModel
	UserID       uint        `gorm:"not null"`
	RefreshToken string      `gorm:"not null;uniqueIndex"`
	UserAgent    string      `gorm:"not null"`
	ClientIP     string      `gorm:"not null"`
	IsBlocked    bool        `gorm:"not null;default:false"`
	ExpiresAt    int64       `gorm:"not null"`
	User         entity.User `gorm:"foreignKey:UserID;"`
}
