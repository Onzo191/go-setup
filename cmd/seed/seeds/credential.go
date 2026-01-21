package seeds

import (
	"log"

	"go-setup/internal/core/database/models"
	authDomain "go-setup/internal/modules/auth/domain"
	"go-setup/internal/pkg/hashing"

	"gorm.io/gorm"
)

type CredentialSeeder struct {
	HashSecret string
}

func (s CredentialSeeder) Name() string {
	return "Credential"
}

func (s CredentialSeeder) Run(db *gorm.DB) error {
	// Map emails to passwords
	credentials := map[string]string{
		"admin@example.com": "admin@1901",
		"user@example.com":  "user@1901",
	}

	for email, password := range credentials {
		// Find user
		var user models.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			log.Printf("User %s not found, skipping credential seeding", email)
			continue
		}

		// Check if credential exists
		var count int64
		if err := db.Model(&models.UserCredential{}).Where("user_id = ? AND type = ?", user.ID, authDomain.CredentialPassword).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			hashedPassword, err := hashing.HashPassword(password, s.HashSecret)
			if err != nil {
				return err
			}

			cred := models.UserCredential{
				UserCredential: authDomain.UserCredential{
					UserID:     user.ID,
					Type:       authDomain.CredentialPassword,
					Identifier: email,
					SecretHash: hashedPassword,
				},
			}

			if err := db.Create(&cred).Error; err != nil {
				return err
			}
			log.Printf("Created credential for user: %s", email)
		} else {
			log.Printf("Credential already exists for user: %s", email)
		}
	}
	return nil
}
