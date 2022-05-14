package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(127.0.0.1:33060)/gorm"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("failed to connect database")
	}

	db.Create(mysql.Config{})
}
