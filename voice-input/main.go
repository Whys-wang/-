package main

import (
	"fmt"
	"net/http"
)

func main() {

	// 静态文件目录
	fs := http.FileServer(http.Dir("./static"))

	http.Handle("/", fs)

	fmt.Println("服务器启动成功")
	fmt.Println("浏览器打开：http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("启动失败：", err)
	}
}
