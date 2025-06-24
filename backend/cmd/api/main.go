package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	// Fix these import paths to match your project's module name
	"github.com/your-username/dev-portfolio/backend/internal/config"
	"github.com/your-username/dev-portfolio/backend/internal/database"
	"github.com/your-username/dev-portfolio/backend/internal/routes"
)

func main() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using environment variables")
	}
	
	// Initialize configuration
	cfg := config.New()
	
	// Connect to database
	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	// Set up router
	router := routes.SetupRouter(db, cfg)
	
	// Start the server
	serverAddr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Server starting on %s in %s mode", serverAddr, cfg.Environment)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
