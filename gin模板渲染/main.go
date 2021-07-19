package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 静态文件的加载
	r.Static("assets", "./static")

	r.LoadHTMLFiles("index.tmpl")

	// 模板的调用
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "下面这个人是猪",
		})
	})

	r.Run(":8080")
}
