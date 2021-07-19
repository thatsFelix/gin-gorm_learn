package main

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 // 零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(120);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号(member number)唯一且不为空
	Num          int     `gorm:"AUTO_INCREAMENT"`
	// 设置num为自增类型
	Address  string `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe int    `gorm:"-"`          // 忽略本字段
}

func main() {
	dsn := "root:root@tcp(47.115.55.129:3306)/go_lang?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Open err: %v", err)
		panic(err)
	}

	db.AutoMigrate(&User{})
}
