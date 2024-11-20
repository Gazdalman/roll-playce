package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	username       string `gorm:"unique;not null"`
	email          string `gorm:"unique;not null"`
	hashedPassword string `gorm:"not null"`
	private        bool   `gorm:"not null;default:false"`
	profilePic     string
	firstName      string
	lastName       string
	bio            string
}
