package seeds

import (
	"log"

	"go-setup/internal/core/database/models"
	userDomain "go-setup/internal/modules/user/domain"

	"gorm.io/gorm"
)

type UserSeeder struct{}

func (s UserSeeder) Name() string {
	return "User"
}

func (s UserSeeder) Run(db *gorm.DB) error {
	users := []userDomain.User{
		{
			Email:     "admin@example.com",
			FirstName: "Admin",
			LastName:  "User",
			Phone:     "+1234567890",
		},
		{
			Email:     "user@example.com",
			FirstName: "Normal",
			LastName:  "User",
			Phone:     "+0987654321",
		},
	}

	for _, u := range users {
		var count int64
		if err := db.Model(&models.User{}).Where("email = ?", u.Email).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			if err := db.Create(&models.User{User: u}).Error; err != nil {
				return err
			}
			log.Printf("Created user: %s", u.Email)
		} else {
			log.Printf("User already exists: %s", u.Email)
		}
	}
	return nil
}
