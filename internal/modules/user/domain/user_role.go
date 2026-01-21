package domain

import (
	"github.com/google/uuid"
)

type UserRole struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	RoleID uuid.UUID `gorm:"primaryKey"`

	ScopeType *string    `gorm:"type:varchar(50)"`
	ScopeID   *uuid.UUID `gorm:"type:uuid"`
}
