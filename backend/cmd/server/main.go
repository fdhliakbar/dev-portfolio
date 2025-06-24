// cmd/server/main.go
package main

import (
	"devportfolio-backend/internal/config"
	"devportfolio-backend/internal/database"
	"devportfolio-backend/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Load configuration
	cfg := config.Load()

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate database
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Set gin mode based on environment
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Setup router
	router := routes.SetupRouter(db, cfg)

	// Start server
	log.Printf("🚀 Server starting on port %s", cfg.Port)
	log.Printf("🌐 Environment: %s", cfg.Environment)
	log.Printf("📖 Health check: http://localhost:%s/health", cfg.Port)
	
	if err := router.Run(":" + cfg.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}