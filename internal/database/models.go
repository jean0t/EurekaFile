package database

import (
	"time"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id int
	Name string
	Author string
	CreatedAt time.Time
}

