package dao

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MysqlDB *gorm.DB

func InitMysql() (err error) {
	dsn := "root:root@tcp(47.115.55.129:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"

	MysqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Open err: %v", err)
		return err
	}

	return nil
}
