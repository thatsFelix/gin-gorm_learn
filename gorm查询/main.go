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

	// 获取第一条记录（主键升序）
	db.First(&user)
	// SELECT * FROM users ORDER BY id LIMIT 1;

	// 获取一条记录，没有指定排序字段
	db.Take(&user)
	// SELECT * FROM users LIMIT 1;

	// 获取最后一条记录（主键降序）
	db.Last(&user)
	// SELECT * FROM users ORDER BY id DESC LIMIT 1;

	result := db.First(&user)
	fmt.Printf("user %#v \n", user)
	fmt.Printf("result %#v \n", result)

	var users []User
	// 查询所有
	db.Debug().Find(&users)
	fmt.Printf("users %#v \n", users)

	// where 条件
	// db.Where("name = ?", "杜鹃").first(&user)

}
