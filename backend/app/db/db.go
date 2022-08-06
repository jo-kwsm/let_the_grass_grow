package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	dbconf := fmt.Sprintf("%s:%s@(db:3306)/%s?charset=utf8", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_DATABASE"))

	var err error
	db, err = gorm.Open(mysql.Open(dbconf), &gorm.Config{})

	if err != nil {
		log.Fatal("Db connection error: ", err)
	}
}

func GetDB() *gorm.DB {
	return db
}
