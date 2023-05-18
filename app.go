package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:loop@54321@/bookstore?charset=utf8&parseTime=True&loc=Local") //"root:loop@54321@/bookstore?charset=utf8&parseTime=True&loc=Local //tcp(127.0.0.1:3306)/bookstore
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
