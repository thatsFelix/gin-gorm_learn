package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

var db *gorm.DB

func initMysql() (err error) {
	dsn := "root:root@tcp(47.115.55.129:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Open err: %v", err)
		return err
	}

	return nil
}

func main() {
	// 创建数据库
	// create database bubble;
	// use bubble;

	// 连接数据库
	err := initMysql()
	if err != nil {
		fmt.Printf("initMysql err: %v", err)
		panic(err)
	}

	// 模型绑定
	db.AutoMigrate(&Todo{})

	r := gin.Default()

	r.Static("/static", "./static")
	// 加载指定文件
	// r.LoadHTMLFiles("./templates/index.html")

	// 加载所有文件
	r.LoadHTMLGlob("./templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	v1Group := r.Group("v1")
	{
		// 添加
		v1Group.POST("/todo", func(c *gin.Context) {
			// 获取请求数据
			var todo Todo
			c.BindJSON(&todo)
			fmt.Printf("todo %#v", todo)
			err := db.Create(&todo).Error
			if err != nil {
				fmt.Printf("Create err: %v", err)
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})

		// 查看所有待办事项
		v1Group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			err = db.Find(&todoList).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})

		// 查看某一个
		v1Group.GET("/todo/:id", func(c *gin.Context) {

		})

		// 修改
		v1Group.PUT("/todo/:id", func(c *gin.Context) {
			id, _ := c.Params.Get("id")
			var todo Todo
			err = db.Where("id=?", id).First(&todo).Error
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"error": err.Error(),
				})
			} else {
				c.BindJSON(&todo)
				// 更新数据库(需要处理错误)
				_ = db.Save(&todo)

				c.JSON(http.StatusOK, todo)
			}
		})

		// 删除
		v1Group.DELETE("/todo/:id", func(c *gin.Context) {

		})
	}

	r.Run()
}
