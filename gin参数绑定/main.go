package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `form:"name"`
	Password string `form:"pass"`
}

func main() {
	r := gin.Default()

	// post请求使用formData、json(raw), get的query string参数绑定方法都是一样, 使用shouldBind
	r.GET("/user", func(c *gin.Context) {
		var u UserInfo

		err := c.ShouldBind(&u)

		if err != nil {
			fmt.Printf("ShouldBind err: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Printf("%#v \n", u)

		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.Run(":8080")
}
