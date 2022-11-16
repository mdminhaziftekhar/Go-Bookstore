package config

import (
	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {

	// USERNAME:PASSWORD@tcp(127.0.0.1:3306)/database_name?charset=utf8&parseTime=True&loc=Local

	d, err := gorm.Open("mysql", "moon:root@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
