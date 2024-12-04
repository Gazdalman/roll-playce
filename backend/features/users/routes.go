package users

import (
	"gin-backend/middleware"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the routes for the users feature
func RegisterRoutes(router *gin.RouterGroup) {
	router.POST("/register", RegisterHandler) // Register new user
	router.POST("/login", LoginHandler)       // Login user
	router.GET("/:id", GetUserHandler)        // Get user by ID
	router.GET("/profile", middleware.JWTMiddleware(), ProfileHandler) // Get the user's profile
}
