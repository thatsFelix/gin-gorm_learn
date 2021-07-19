package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/web", func(c *gin.Context) {
		// 方式1:
		// name := c.Query("name")
		// 方式2: 默认值
		// name := c.DefaultQuery("query", "默认值")
		// 方式3: 另一种方式设置默认值
		// 多个参数的话直接添加即可
		name, ok := c.GetQuery("name")
		if !ok {
			name = "另一种方式设置默认值"
		}
		fmt.Println("name", name)
		c.JSON(200, gin.H{
			"name": name,
		})
	})

	r.Run(":10001")
}
