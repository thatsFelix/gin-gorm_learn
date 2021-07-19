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

	// 2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	// 3.创建
	u1 := User{Age: 18}

	fmt.Println("判断主键是不是为空 db.NewRecord(&u1)")

	db.Create(&u1)
}
