package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 1.定义模型
type User struct {
	ID   int64
	Name string `gorm:"default:'小王子'"` // 设置值默认值
	Age  int64
}

func main() {
	dsn := "root:root@tcp(47.115.55.129:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Open err: %v", err)
		panic(err)
	}

	var user User
	db.First(&user)

	user.Name = "Bob Dylan"
	// user.Age = 85

	// save默认会要修改所有的字段
	// db.Debug().Save(&user)

	// 只会操作指定字段(效率更高)
	db.Debug().Model(&user).Update("name", "Bob")
}
