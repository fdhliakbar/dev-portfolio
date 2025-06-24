// internal/database/db.go
package database

import (
	"fmt"
	"log"
	"os"
	
	"github.com/your-username/dev-portfolio/backend/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB establishes a connection to the database based on configuration
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	
	if cfg.DatabaseURL == "" {
		// Use SQLite for development if no DATABASE_URL is provided
		log.Println("Using SQLite database for development")
		db, err = gorm.Open(sqlite.Open("dev.db"), &gorm.Config{})
	} else {
		// Use PostgreSQL with the provided connection string
		log.Println("Connecting to PostgreSQL database")
		db, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	
	// TODO: Run migrations here if needed
	// db.AutoMigrate(&YourModel{})
	
	return db, nil
}
	}

	return db, nil
}

func Migrate(db *gorm.DB) error {
	// TODO: Add your models here when created
	// Example: return db.AutoMigrate(&models.User{}, &models.Project{})
	
	log.Println("⏳ Database migration completed")
	return nil
}
