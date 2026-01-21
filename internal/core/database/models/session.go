package models

import (
	authDomain "go-setup/internal/modules/auth/domain"
)

type UserSession struct {
	authDomain.UserSession

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (UserSession) TableName() string {
	return "user_sessions"
}
