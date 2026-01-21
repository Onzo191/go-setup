package models

import (
	authDomain "go-setup/internal/modules/auth/domain"
)

type UserCredential struct {
	authDomain.UserCredential

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
}

func (UserCredential) TableName() string {
	return "user_credentials"
}
