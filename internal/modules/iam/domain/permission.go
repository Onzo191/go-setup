package domain

import (
	model "go-setup/internal/modules/shared/domain"
)

type Permission struct {
	Key         string `gorm:"not null;uniqueIndex"`
	Description string

	model.BaseModel
}
