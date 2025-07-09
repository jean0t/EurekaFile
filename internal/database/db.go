package database

import (
	"fmt"
	"os"
	"hash"
	"crypto/sha256"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

//================================================================| Database

func createDsn() string {
	var host,user,password,dbname,port string = os.Getenv("host"), os.Getenv(user), os.Getenv("password"), os.Getenv("dbname"), os.Getenv("port") 
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	return dsn
}

func ConnectToDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	
	var dsn string = createDsb()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config())

	return db, err
}


//================================================================| User

func RegisterUser(db *gorm.DB, username, password string) error {
	var passwordHashed string = HashPassword(password)
	// Creating user => Files []File ommited since it has no files uploaded yet 
	var user = User{
		Username: username,
		PasswordHash: passwordHashed,
	}

	result = db.Create(&user} 
	return result.Error
}

func IsValidUser(db *gorm.DB, username, password string) bool {
	var user User
	var passwordHashed string = HashPassword(password)

	var result = db.Where("Username = ?", username).First(&user)
	if result.Error != nil {
		return false // not found or db error
	}

	if passwordHashed != user.PasswordHash {
		return false // password not right
	}

	return true
}

