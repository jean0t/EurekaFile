package database

import (
	"fmt"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, username, password string) error {
	var passwordHashed string = HashPassword(password)
	// Creating user => Files []File ommited since it has no files uploaded yet 
	var user = User{
		Username: username,
		PasswordHash: passwordHashed,
	}

	var result = db.Create(&user) 

	if result.Error != nil {
		fmt.Println("[!] Error in register is: ", result.Error)
	}

	return result.Error
}


func IsValidUser(db *gorm.DB, username, password string) error {
	var user User
	var passwordHashed string = HashPassword(password)

	var result = db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		RegisterUser(db, username, password) // if user doesnt exists, it is created :)
		return nil
	}

	if passwordHashed != user.PasswordHash {
		return fmt.Errorf("Password Wrong")
	}

	return nil
}
