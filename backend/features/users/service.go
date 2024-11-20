package users

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

// HashPassword hashes the user's password
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// ComparePasswords compares a hashed password with a plain password
func ComparePasswords(hashedPassword, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		log.Println("Password mismatch:", err)
		return false
	}
	return true
}

func RegisterUser(username, email, password, firstName, lastName, bio, profilePic string, private bool) (*User, error) {
	if username == "" || email == "" || password == "" {
		return nil, errors.New("all fields are required")
	}

	// Hash the password
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := &User{
		username:       username,
		email:          email,
		hashedPassword: hashedPassword,
		firstName:      firstName,
		lastName:       lastName,
		bio:            bio,
		profilePic:     profilePic,
		private:        private,
	}

	// Save the user to the database
	if err := CreateUser(user); err != nil {
		return nil, err
	}

	return user, nil
}

// LoginUser checks the user credentials and returns the user object if valid
func LoginUser(username, password string) (*User, error) {
	// Retrieve the user by username
	user, err := GetUserByUsernameOrEmail(username)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Compare the provided password with the stored hashed password
	if !ComparePasswords(user.hashedPassword, password) {
		return nil, errors.New("invalid password")
	}

	return user, nil
}
