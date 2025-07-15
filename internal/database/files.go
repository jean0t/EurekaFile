package database

import (
	"gorm.io/gorm"
)

func RegisterFile(db *gorm.DB, user User, name string) error {
	var file File = File{Name: name, UserID: user.ID, Author: user}

	var result = db.Create(&file)
	return result.Error
}
