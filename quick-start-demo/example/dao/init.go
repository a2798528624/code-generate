package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (err error) {
	dsn := "root:a1131657944@tcp(127.0.0.1:3306)/sqlc-test?charset=utf8"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}
