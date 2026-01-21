package domain

import (
	"time"

	model "go-setup/internal/modules/shared/domain"

	"github.com/google/uuid"
)

type UserSession struct {
	model.Identity

	UserID uuid.UUID `gorm:"type:uuid;not null;index"`

	RefreshTokenHash string `gorm:"not null;uniqueIndex"`

	UserAgent string `gorm:"not null"`
	IPAddress string `gorm:"not null"`

	ExpiresAt time.Time
	RevokedAt *time.Time

	model.Timestamps
}
