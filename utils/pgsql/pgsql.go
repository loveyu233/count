package pgsql

import "gorm.io/gorm"
import (
	"gorm.io/driver/postgres"
)

var DB *gorm.DB

func InitPgsql(str string) {
	db, err := gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB = db
}
