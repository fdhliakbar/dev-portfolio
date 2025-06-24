package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/your-username/dev-portfolio/backend/internal/config"
)

func SetupRouter(db *gorm.DB, cfg *config.Config) *gin.Engine {
	router := gin.Default()

	// CORS middleware
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{cfg.ClientURL, "http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":      "ok",
			"message":     "DevPortfolio API is running",
			"environment": cfg.Environment,
			"port":        cfg.Port,
		})
	})

	// API routes group
	api := router.Group("/api")
	{
		// Basic status endpoint
		api.GET("/status", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"api_version": "v1.0.0",
				"status":      "healthy",
				"database":    "connected",
			})
		})

		// TODO: Add your routes here
		// Example:
		// auth := api.Group("/auth")
		// public := api.Group("/public") 
		// admin := api.Group("/admin")
	}

	return router
}