package users

import (
	"fmt"
	"gin-backend/config"
	"gin-backend/middleware"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHandler handles the user registration request
func RegisterHandler(c *gin.Context) {
	var requestData struct {
		Username   string `form:"username"`
		Email      string `form:"email"`
		Password   string `form:"password"`
		FirstName  string `form:"first_name"`
		LastName   string `form:"last_name"`
		Bio        string `form:"bio"`
		Private    bool   `form:"private"`
	}

	// Bind form data to the struct
	if err := c.ShouldBind(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var profilePicURL string
	file, header, err := c.Request.FormFile("profile_pic") // "profile_pic" is the form field name
	if err == nil && header != nil { // If a file is provided
		defer file.Close()

		// Upload the file to S3 (or your storage solution)
		profilePicURL, err = config.UploadFile(file, header.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println("Error registering user:", err)
			return
		}
	} else {
		// If no file provided, set a default profile picture
		profilePicURL = "https://rollplayce.s3.us-east-2.amazonaws.com/default-avatar-icon-of-social-media-user-vector.jpg"
	}

	// Call the service to register the user
	user, err := RegisterUser(requestData.Username, requestData.Email, requestData.Password, requestData.FirstName, requestData.LastName, requestData.Bio, profilePicURL, requestData.Private)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Generate a JWT token
	token, err := middleware.GenerateJWT(fmt.Sprintf("%d", user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// Set the token in the response header
	c.Header("Authorization", "Bearer "+token)

	// Add JWT token to cookies
	c.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)

	// Set the user in the context
	c.Set("user", user)

	// Return the created user (without the password)
	c.JSON(http.StatusOK, gin.H{
		"message": "User registered successfully",
		"user": gin.H{
			"username":   user.Username,
			"email":      user.Email,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
	})
}

// LoginHandler handles the user login request
func LoginHandler(c *gin.Context) {
	var requestData struct {
		Credential string `json:"credential"`
		Password   string `json:"password"`
	}

	// Bind JSON request body to the struct
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Call the service to authenticate the user
	user, err := LoginUser(requestData.Credential, requestData.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Generate a JWT token
	token, err := middleware.GenerateJWT(fmt.Sprintf("%d", user.ID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating token"})
		return
	}

	// Set the token in the response header
	c.Header("Authorization", "Bearer "+token)

	// Add JWT token to cookies
	c.SetCookie("token", token, 60*60*24, "/", "localhost", false, true)

	// Set the user in the context
	c.Set("user", user)

	// Return the authenticated user (without the password)
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user": gin.H{
			"username":   user.Username,
			"email":      user.Email,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
		},
		"token": token,
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

	if user.Private {
		// Send only the bio, profilePic, and username
		c.JSON(http.StatusOK, gin.H{
			"bio":        user.Bio,
			"profilePic": user.ProfilePic,
			"username":   user.Username,
		})
		return
	}

	// Return the found user as a JSON response
	c.JSON(http.StatusOK, user)
}

func ProfileHandler(c *gin.Context) {
	userID, _ := c.Get("user") // Get the user ID from the context

	// Find the user by ID
	user, err := GetUserByID(fmt.Sprintf("%v", userID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "User not found",
		})
		return
	}

	user = &User{
		Username:   user.Username,
		Email:      user.Email,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Bio:        user.Bio,
		ProfilePic: user.ProfilePic,
		Private:    user.Private,
	}

	// Return the user as a JSON response
	c.JSON(http.StatusOK, user)
}
