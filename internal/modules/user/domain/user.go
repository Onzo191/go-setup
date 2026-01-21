package domain

import (
	model "go-setup/internal/modules/shared/domain"
)

type User struct {
	Email     string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Phone     string `gorm:"not null"`

	model.SoftDeleteModel
}
