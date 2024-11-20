package main

import (
	"gin-backend/config"
	// "gin-backend/features/chats"
	// "gin-backend/features/users"
	"gin-backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	config.ConnectDatabase()

	// Auto-migrate models
	// config.DB.AutoMigrate(&users.User{}, &chats.Chat{}) // Add other models here

	// Initialize Gin
	router := gin.Default()

	// Register routes
	api := router.Group("/api")
	{
		// usersGroup := api.Group("/users")
		// users.RegisterRoutes(usersGroup)

		// chatsGroup := api.Group("/chats")
		// chats.RegisterRoutes(chatsGroup)

		// Add routes for Friends, Videos, and Likes similarly
	}

	// Use middleware
	api.Use(middleware.AuthMiddleware())

	// Start server
	router.Run(":8080")
}
