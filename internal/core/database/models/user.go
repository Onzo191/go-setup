package models

import (
	userDomain "go-setup/internal/modules/user/domain"
)

type User struct {
	userDomain.User

	Credentials []UserCredential `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Sessions    []UserSession    `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Roles       []Role           `gorm:"many2many:user_roles"`
}
