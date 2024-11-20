package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterHandler handles the user registration request
func RegisterHandler(c *gin.Context) {
	var requestData struct {
		Username   string `json:"username"`
		Email      string `json:"email"`
		Password   string `json:"password"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		Bio        string `json:"bio"`
		ProfilePic string `json:"profile_pic"`
		Private    bool   `json:"private"`
	}

	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service to register the user
	user, err := RegisterUser(requestData.Username, requestData.Email, requestData.Password, requestData.FirstName, requestData.LastName, requestData.Bio, requestData.ProfilePic, requestData.Private)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the created user (without the password)
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"username":   user.username,
			"email":      user.email,
			"first_name": user.firstName,
			"last_name":  user.lastName,
		},
	})
}

// LoginHandler handles the user login request
func LoginHandler(c *gin.Context) {
	var requestData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service to authenticate the user
	user, err := LoginUser(requestData.Username, requestData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Return the authenticated user (without the password)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"username":   user.username,
			"email":      user.email,
			"first_name": user.firstName,
			"last_name":  user.lastName,
		},
	})
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("id") // Get the ID from the URL parameter

	// Find the user by ID (You can modify this to suit your database retrieval logic)
	user, err := GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	if user.private {
		// Send only the bio, profilePic, and username
		c.JSON(http.StatusOK, gin.H{
			"bio":        user.bio,
			"profilePic": user.profilePic,
			"username":   user.username,
		})
		return
	}

	// Return the found user as a JSON response
	c.JSON(http.StatusOK, user)
}

func ProfileHandler(c *gin.Context) {
	
}
