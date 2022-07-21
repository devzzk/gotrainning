package main

import (
	"github.com/devzzk/gotrainning/03gorom/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&models.Product{})
	db.Create(&models.Product{
		Code:  "test",
		Price: 100,
	})

	// 若使用sqlite 时 需要查到之后才能进行后续操作
	var product models.Product
	db.First(&product, 1)
	db.First(&product, "code = ?", "F42") // 查找 code 字段值为 D42 的记录

	db.Model(&product).Update("Price", 200)
	db.Model(&product).Updates(models.Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
}
