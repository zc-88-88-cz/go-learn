package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// 网络编程示例
func main() {
	// 1. HTTP客户端请求
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("HTTP响应:", string(body))

	// 2. 创建HTTP服务器
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "欢迎访问Go网络编程示例!")
	})

	fmt.Println("服务器启动在 :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}