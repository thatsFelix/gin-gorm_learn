package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/json", func(c *gin.Context) {
		// 方法1: 使用map
		/**
		data := map[string]interface{}{
			"name":    "陈峰",
			"age":     18,
			"message": "Hi, its me!",
		}
		*/
		data := gin.H{
			"name":    "陈峰",
			"age":     18,
			"message": "Hi, its me!",
		}
		fmt.Println("data", data)

		// 方法2, 使用结构体
		// type msg struct {
		// 	Name    string `json:"修改字段名"`
		// 	Message string
		// 	Age     int
		// }

		// c.JSON(http.StatusOK, data)
		// c.JSON(http.StatusOK, msg{
		// 	"Felix",
		// 	"你好呀",
		// 	18,
		// })

		// 匿名struct
		c.JSON(http.StatusOK, struct {
			Name string
			Msg  string
			Age  int
		}{
			"Felix",
			"你好呀",
			18,
		})

	})

	r.Run(":10001")
}
