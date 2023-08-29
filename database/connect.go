package database

import (
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Connect() *gorm.DB {
	USER := os.Getenv("DBUSER")
	PASSWORD := os.Getenv("DBPASSWORD")
	PROTOCOL := "tcp(localhost:3306)"
	DBNAME := os.Getenv("DBNAME")

	CONNECT := USER + ":" + PASSWORD + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	return db
}
