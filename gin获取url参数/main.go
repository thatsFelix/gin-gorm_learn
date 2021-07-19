package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/param/:name/:age", func(c *gin.Context) {
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(200, gin.H{
			"name": "您输入的名字是" + name,
			"age":  age,
		})
	})

	r.Run(":8080")
}
