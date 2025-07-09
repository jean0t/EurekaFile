package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username	string	`gorm:"unique;not null"`
	PasswordHash	string	`gorm:"not null"`
	Files		[]File	`gorm:"constraint:OnDelete:CASCADE"`
}

type File struct {
	gorm.Model
	Name	string
	UserID	uint
	Author	User
}


