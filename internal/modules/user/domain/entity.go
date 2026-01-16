package domain

import (
	model "go-setup/internal/core/database"
)

type User struct {
	model.BaseModel
	Email     string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Phone     string `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
