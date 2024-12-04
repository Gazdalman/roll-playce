package main

import (
	"gin-backend/config"
	"github.com/joho/godotenv"
	"log"
	"os"
	// "gin-backend/features/chats"
	"gin-backend/features/users"
	"gin-backend/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// Connect to the database
	config.ConnectDatabase()

	// Auto-migrate models
	config.DB.AutoMigrate(&users.User{}) // Add other models here

	// Start S3
	config.InitS3()

	// Initialize Gin
	router := gin.Default()

	// Add CSRF middleware
	secretKey := os.Getenv("CSRF_SECRET_KEY")

	log.Println("Secret Key: ", secretKey)

	router.Use(middleware.CSRFMiddleware(secretKey))

	router.GET("/csrf-token", func(c *gin.Context) {
		// Retrieve the token from the middleware
		csrfToken, exists := c.Get("csrf_token")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "CSRF token not found"})
			return
		}

		// Return the token in the response body
		c.JSON(http.StatusOK, gin.H{
			"csrf_token": csrfToken,
		})
	})


	// Register routes
	api := router.Group("/api")
	{
		usersGroup := api.Group("/users")
		users.RegisterRoutes(usersGroup)

		// chatsGroup := api.Group("/chats")
		// chats.RegisterRoutes(chatsGroup)

		// Add routes for Friends, Videos, and Likes similarly
	}

	// Start server
	router.Run(os.Getenv("PORT"))
}
