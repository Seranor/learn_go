package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	client := &http.Client{
		Timeout: 30 * time.Second, // 设置较长的超时时间
	}

	resp, err := client.Get("http://127.0.0.1:8080/add?a=1&b=2")
	if err != nil {
		fmt.Println("请求错误:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("读取响应体错误:", err)
		return
	}

	fmt.Println(string(body))
}
