package models

import (
	iamDomain "go-setup/internal/modules/iam/domain"
)

type RolePermission struct {
	iamDomain.RolePermission
	Role       Role       `gorm:"foreignKey:RoleID;references:ID;constraint:OnDelete:CASCADE"`
	Permission Permission `gorm:"foreignKey:PermissionID;references:ID;constraint:OnDelete:CASCADE"`
}
