package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/web", func(c *gin.Context) {
		// 方式1:
		// name := c.PostForm("username")
		// age := c.PostForm("age")
		// fmt.Printf("%T --- %v \n", age, age)
		// 方式2: 默认是
		// name := c.DefaultPostForm("name", "默认名字")
		// 方式3: 另一种方式设置默认值
		name, ok := c.GetPostForm("name")
		if !ok {
			name = "另一种方式设置默认值"
		}
		c.JSON(200, gin.H{
			"name": "name是" + name,
		})
	})

	r.Run(":10001")
}
