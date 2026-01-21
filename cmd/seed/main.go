package main

import (
	"log"

	"go-setup/cmd/seed/seeds"
	config "go-setup/internal/core/config"
	"go-setup/internal/core/database"
)

func main() {
	cfg := config.LoadEnvConfig(".env")

	db, err := database.ConnectPostgres(
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.User,
		cfg.DB.Password,
		cfg.DB.DBName,
	)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	seeders := []seeds.Seeder{
		seeds.PermissionSeeder{},
		seeds.IAMSeeder{},
		seeds.UserSeeder{},
		seeds.CredentialSeeder{
			HashSecret: cfg.Security.PasswordHashSecret,
		},
		seeds.UserRoleSeeder{},
	}

	log.Println("Starting database seeding...")

	for _, seeder := range seeders {
		log.Printf("Running seeder: %s", seeder.Name())
		if err := seeder.Run(db); err != nil {
			log.Fatalf("Failed to run seeder %s: %v", seeder.Name(), err)
		}
	}

	log.Println("Seeding completed successfully.")
}
