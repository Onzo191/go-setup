package seeds

import (
	"log"

	"go-setup/internal/core/database/models"
	iamDomain "go-setup/internal/modules/iam/domain"

	"gorm.io/gorm"
)

type PermissionSeeder struct{}

func (s PermissionSeeder) Name() string {
	return "Permission"
}

func (s PermissionSeeder) Run(db *gorm.DB) error {
	permissions := []iamDomain.Permission{
		{Key: "users:read", Description: "Read access to users"},
		{Key: "users:write", Description: "Write access to users"},
		{Key: "roles:read", Description: "Read access to roles"},
		{Key: "roles:write", Description: "Write access to roles"},
		{Key: "system:manage", Description: "Full system management"},
	}

	for _, p := range permissions {
		var count int64
		if err := db.Model(&models.Permission{}).Where("key = ?", p.Key).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if err := db.Create(&models.Permission{Permission: p}).Error; err != nil {
				return err
			}
			log.Printf("Created permission: %s", p.Key)
		} else {
			log.Printf("Permission already exists: %s", p.Key)
		}
	}
	return nil
}
