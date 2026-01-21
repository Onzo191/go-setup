package models

import (
	iamDomain "go-setup/internal/modules/iam/domain"
)

type Role struct {
	iamDomain.Role

	Permissions []Permission `gorm:"many2many:role_permissions"`
	Users       []User       `gorm:"many2many:user_roles"`
}
