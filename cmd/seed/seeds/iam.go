package seeds

import (
	"log"

	"go-setup/internal/core/database/models"
	iamDomain "go-setup/internal/modules/iam/domain"

	"gorm.io/gorm"
)

type IAMSeeder struct{}

func (s IAMSeeder) Name() string {
	return "IAM (Roles)"
}

func (s IAMSeeder) Run(db *gorm.DB) error {
	roles := []iamDomain.Role{
		{
			Name:        "user",
			Description: "Standard user role",
		},
		{
			Name:        "admin",
			Description: "Administrator role",
		},
	}

	for _, r := range roles {
		var count int64
		if err := db.Model(&models.Role{}).Where("name = ?", r.Name).Count(&count).Error; err != nil {
			return err
		}

		var role models.Role
		if count == 0 {
			role = models.Role{Role: r}
			if err := db.Create(&role).Error; err != nil {
				return err
			}
			log.Printf("Created role: %s", r.Name)
		} else {
			if err := db.Where("name = ?", r.Name).First(&role).Error; err != nil {
				return err
			}
			log.Printf("Role already exists: %s", r.Name)
		}

		// Assign permissions
		// Admin gets all permissions
		// User gets users:read
		var perms []models.Permission
		if r.Name == "admin" {
			if err := db.Find(&perms).Error; err != nil {
				return err
			}
		} else if r.Name == "user" {
			if err := db.Where("key = ?", "users:read").Find(&perms).Error; err != nil {
				return err
			}
		}

		if len(perms) > 0 {
			if err := db.Model(&role).Association("Permissions").Replace(&perms); err != nil {
				return err
			}
			log.Printf("Assigned %d permissions to role: %s", len(perms), r.Name)
		}
	}
	return nil
}
