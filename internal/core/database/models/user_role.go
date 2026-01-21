package models

import (
	userDomain "go-setup/internal/modules/user/domain"
)

type UserRole struct {
	userDomain.UserRole

	User User `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE"`
	Role Role `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE"`
}
