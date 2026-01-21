package seeds

import (
	"log"

	"go-setup/internal/core/database/models"

	"gorm.io/gorm"
)

type UserRoleSeeder struct{}

func (s UserRoleSeeder) Name() string {
	return "UserRole"
}

func (s UserRoleSeeder) Run(db *gorm.DB) error {
	assignments := map[string]string{
		"admin@example.com": "admin",
		"user@example.com":  "user",
	}

	for email, roleName := range assignments {
		var user models.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			log.Printf("User not found: %s", email)
			continue
		}

		var role models.Role
		if err := db.Where("name = ?", roleName).First(&role).Error; err != nil {
			log.Printf("Role not found: %s", roleName)
			continue
		}

		// Check and assign
		// Using raw association check or count on join table would be ideal, but Replace is easier for idempotency if we want strict state.
		// However, Append is safer if we don't want to wipe other roles.
		// Let's use Append+count check for now to avoid duplicates if Replace is not desired (though for seed replace is usually fine).
		// Better: check if association exists.

		exists := db.Model(&user).Where("id = ?", role.ID).Association("Roles").Count()
		if exists == 0 {
			if err := db.Model(&user).Association("Roles").Append(&role); err != nil {
				return err
			}
			log.Printf("Assigned role %s to user %s", roleName, email)
		} else {
			log.Printf("User %s already has role %s", email, roleName)
		}
	}
	return nil
}
