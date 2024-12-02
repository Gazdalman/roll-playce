package users

import "gin-backend/config"

func CreateUser(user *User) error {
	return config.DB.Create(user).Error
}

func GetAllUsers() ([]User, error) {
	var users []User
	err := config.DB.Find(&users).Error
	return users, err
}

// GetUserByUsernameOrEmail retrieves a user by either their username or email
func GetUserByUsernameOrEmail(credential string) (*User, error) {
	var user User
	// Try to find the user by username first, if not found, try email
	if err := config.DB.Where("Username = ? OR Email = ?", credential, credential).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// GetUserByID retrieves a user by their ID
func GetUserByID(id string) (*User, error) {
	var user User

	if err := config.DB.Where("ID = ?", id).First(&user).Error; err != nil { // runs the query and sets the result to the user variable. then checks if there is an error
		return nil, err
	}
	return &user, nil
}
