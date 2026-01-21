package domain

import (
	model "go-setup/internal/modules/shared/domain"
)

type Permission struct {
	model.Identity

	Key         string `gorm:"not null;uniqueIndex"`
	Description string

	model.Timestamps
}
