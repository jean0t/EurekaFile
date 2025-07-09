package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)


func createDsn() string {}

func ConnectToDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	
	var dsn string = createDsb()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config())

	return db, err
}
