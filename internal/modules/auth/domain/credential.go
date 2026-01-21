package domain

import (
	"time"

	model "go-setup/internal/modules/shared/domain"

	"github.com/google/uuid"
)

type CredentialType string

const (
	CredentialPassword CredentialType = "password"
	CredentialGoogle   CredentialType = "google"
	CredentialGithub   CredentialType = "github"
)

type UserCredential struct {
	UserID uuid.UUID `gorm:"type:uuid;not null;index"`

	Type       CredentialType `gorm:"type:varchar(20);not null"`
	Identifier string         `gorm:"not null"` // email|phone \ google \ github
	SecretHash string         `gorm:"not null"`
	VerifiedAt time.Time

	model.BaseModel
}
