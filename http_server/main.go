package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayHello(writer http.ResponseWriter, request *http.Request) {
	b, _ := ioutil.ReadFile("hello.html")

	fmt.Fprintln(writer, string(b))
}

func main() {
	http.HandleFunc("/list", sayHello)
	err := http.ListenAndServe(":9527", nil)
	if err != nil {
		fmt.Printf("ListenAndServe err: %v", err)
		return
	}
}
