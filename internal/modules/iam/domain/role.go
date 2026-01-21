package domain

import (
	model "go-setup/internal/modules/shared/domain"
)

type Role struct {
	model.Identity

	Name        string `gorm:"not null;uniqueIndex"`
	Description string

	model.Timestamps
}
