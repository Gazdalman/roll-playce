package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB // Database connection. It is a pointer to gorm.DB

func ConnectDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("db/dev.db"), &gorm.Config{}) // Connect to SQLite database. If dev.db does not exist, it will be created.
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	log.Println("Database connection established!")
}
