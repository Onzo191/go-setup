package main

import (
	"log"

	config "go-setup/internal/core/config"
	"go-setup/internal/core/database"
	"go-setup/internal/core/server"
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
		log.Fatal("Failed to connect to database:", err)
	}
	_ = db // Use the db connection as needed

	srv := server.New(cfg)
	srv.RegisterModules()
	srv.Start()
}
