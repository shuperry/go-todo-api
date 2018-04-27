package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB

	sqlConnection = "root:mysecretpassword@(127.0.0.1:3306)/go_todos?charset=utf8&parseTime=True&loc=Local"
)

func InitConnection()(*gorm.DB) {
	var err error

	db, err = gorm.Open("mysql", sqlConnection)

	if err != nil {
		panic("connect database failed")
	}

	return db
}
