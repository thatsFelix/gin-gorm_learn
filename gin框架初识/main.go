package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default() // 返回默认的路由引擎

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	// r.Run() // 默认是8080
	r.Run(":9099")
}
