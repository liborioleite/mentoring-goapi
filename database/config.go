package database

import (
	"github.com/liborioleite/mentoring-goapi/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConfigDB() error {
	dsn := "root:@tcp(127.0.0.1:3306)/mentoring?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.AutoMigrate(&schemas.Users{})

	if err != nil {
		return err
	}

	DB = db

	return nil
}
