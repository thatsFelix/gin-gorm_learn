package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(47.115.55.129:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Open err: %v", err)
		panic(err)
	}

	// 创建表 自动迁移(把结构体和数据表进行对应)
	db.AutoMigrate(&UserInfo{}) // 生成的数据表是user_infos

	// 增加数据
	u1 := UserInfo{1, "雷神", "女", "吃炸鸡"}
	db.Create(&u1)

	// 查询数据
	var u UserInfo
	db.First(&u)

	fmt.Println("u的数据是 ", u)

	// 更新
	db.Model(&u).Update("hobby", "双色球")

	// 删除
	// db.Delete(&u)
}
