package database

import (
	"fmt"
	"os"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

//================================================================| Database


func MigrateDB(db *gorm.DB) error {
	var err error

	err = db.AutoMigrate(&User{}, &File{})
	if err != nil {
		fmt.Println("[!] Error with Migration in database")
		return err
	}
	
	fmt.Println("[*] Migration successful")
	return err
}

func createDsn() string {
	var host,user,password,dbname,port string = os.Getenv("host"), os.Getenv("user"), os.Getenv("password"), os.Getenv("dbname"), os.Getenv("port") 
	var dsn string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	return dsn
}


func ConnectToDB() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	
	var dsn string = createDsn()
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}


