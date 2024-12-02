package users

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null"`
	Email          string `gorm:"unique;not null"`
	HashedPassword string `gorm:"not null"`
	Private        bool   `gorm:"not null;default:false"`
	ProfilePic     string
	FirstName      string
	LastName       string
	Bio            string
}

