package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func main() {
	http.HandleFunc("/tmpl", func(w http.ResponseWriter, r *http.Request) {
		// 解析模板
		t, err := template.ParseFiles("./hello.tmpl")
		if err != nil {
			fmt.Printf("ParseFiles err: %v", err)
			return
		}

		u1 := User{
			"张三",
			"男",
			18,
		}

		// 渲染模板
		// t.Execute(w, "小王子")
		t.Execute(w, u1)
	})

	http.ListenAndServe(":9090", nil)
}
