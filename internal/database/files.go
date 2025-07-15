package database

import (
	"gorm.io/gorm"
)

func GetAllFiles(db *gorm.DB) ([]File, error) {
	var (
		files []File
		err error
	)

	err = db.Find(&files).Error
	if err != nil {
		return []File{}, err
	}

	return files, nil
}


func RegisterFile(db *gorm.DB, user User, name string) error {
	var file File = File{Name: name, UserID: user.ID, Author: user}

	var result = db.Create(&file)
	return result.Error
}
