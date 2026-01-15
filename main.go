package main

import (
	// my-modules
	"go-setup/config"

	// third-party modules
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadEnv(".env")
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run(":" + cfg.PORT) // listens on 0.0.0.0:8080 by default
}
