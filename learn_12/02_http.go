package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World!</h1>")
}
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello login!</h1>")
}
func main() {
	// 路径分发函数
	http.HandleFunc("/index", index)
	http.HandleFunc("/login", login)

	// 启动服务
	http.ListenAndServe("0.0.0.0:8080", nil)
}
