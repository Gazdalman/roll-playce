package config

import (
	"log"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // Database connection. It is a pointer to gorm.DB

func ConnectDatabase() {
	var err error

	// Check the environment from the .env file
	env := os.Getenv("ENV") // This should already be loaded in your main package
	if env == "" {
		env = "development" // Default to development if not set
	}

	if env == "production" {
		// Production environment: Use PostgreSQL
		dsn := os.Getenv("DB_URL") // Get the Postgres connection string from the .env file
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to the PostgreSQL database:", err)
		}
		log.Println("Connected to PostgreSQL database!")
	} else {
		// Development environment: Use SQLite
		DB, err = gorm.Open(sqlite.Open("db/dev.db"), &gorm.Config{}) // Use SQLite
		if err != nil {
			log.Fatal("Failed to connect to the SQLite database:", err)
		}
		log.Println("Connected to SQLite database!")
	}
}
