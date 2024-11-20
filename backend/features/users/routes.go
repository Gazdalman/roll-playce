package users

import "github.com/gin-gonic/gin"

// RegisterRoutes sets up the routes for the users feature
func RegisterRoutes(router *gin.Engine) {
	users := router.Group("/users")
	{
		users.POST("/register", RegisterHandler) // Register new user
		users.POST("/login", LoginHandler)       // Login user
		users.GET("/:id", GetUserHandler)        // Get user by ID
	}
	router.GET("/profile", ProfileHandler) // Get the user's profile
}
