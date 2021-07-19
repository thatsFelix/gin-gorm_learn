package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

// 定义一个中间件函数
func middleware1(c *gin.Context) {
	fmt.Println("进来m1了")

	// 中间件之间通讯, 设置值
	c.Set("m1", true)

	c.Next()
	fmt.Println("执行了后面的处理函数以后, 会再回来走这里")

	// 阻止调用后面的函数
	// c.Abort()

}

// 定义一个中间件函数
func middleware2(c *gin.Context) {
	fmt.Println("进来m2了")

	// 中间件通讯
	m1, ok := c.Get("m1")
	if ok {
		fmt.Println("m1的值是 ===> ", m1)
	}

	c.Next()
	fmt.Println("出去m2了")

	// 阻止调用后面的函数
	// c.Abort()
}

func main() {
	r := gin.Default()

	// 先经过中间件, 再走处理函数
	// 单个路由注册中间件
	// r.GET("/index", middleware, indexHandler)

	// 全局注册中间件
	r.Use(middleware1, middleware2)
	r.GET("/index", indexHandler)

	r.Run()
}
