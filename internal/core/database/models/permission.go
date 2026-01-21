package models

import (
	iamDomain "go-setup/internal/modules/iam/domain"
)

type Permission struct {
	iamDomain.Permission

	Roles []Role `gorm:"many2many:role_permissions"`
}
