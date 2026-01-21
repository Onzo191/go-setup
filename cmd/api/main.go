package main

import (
	"log"

	_ "go-setup/docs"
	config "go-setup/internal/core/config"
	"go-setup/internal/core/database"
	"go-setup/internal/core/server"
)

// @title           Onzo191/Go Setup
// @version         1.0
// @description     This is a sample server for modular monolith architecture.
// @termsOfService  http://swagger.io/terms/

// @contact.name    API Support
// @contact.url     http://www.swagger.io/support
// @contact.email   support@swagger.io

// @license.name    Apache 2.0
// @license.url     http://www.apache.org/licenses/LICENSE-2.0.html

// @host            localhost:8080
// @BasePath        /api/v1
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
