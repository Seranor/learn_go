package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func add(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm() // 解析参数
	fmt.Println("path", r.URL.Path)
	a, _ := strconv.Atoi(r.Form["a"][0])
	b, _ := strconv.Atoi(r.Form["b"][0])
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	jData, _ := json.Marshal(map[string]int{
		"data": a + b,
	})
	_, _ = w.Write(jData)
}
func main() {
	http.HandleFunc("/add", add)
	_ = http.ListenAndServe(":8080", nil)
}
