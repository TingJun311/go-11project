package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

const (
    DBName = "golang"
    DBPassword = "ctj0311"
    DBConnection = "tcp"
    DBPort = "127.0.0.1:3306"
    DBTableName = "book_management_system"
    DBServer = "mysql"
)

func DBConnect() {
	connection, err := gorm.Open(DBServer, DBName + ":" + DBPassword + "@tcp(" + DBPort + ")/" + DBTableName + "?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = connection
}

func GetDB() (*gorm.DB) {
	return db
}