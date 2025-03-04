package main

import (
	"fmt"
	"net"
)

func main() {
	clientSk, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(clientSk)
	// 延迟注册关闭套接字
	defer clientSk.Close()

	for true {
		fmt.Println("请输入内容>>>")
		var data = make([]byte, 1024)
		fmt.Scan(&data)
		//fmt.Println(string(data))
		if string(data) == "exit" {
			break
		}
		// 发送数据
		clientSk.Write(data)

		// 接收服务端的响应数据
		var res = make([]byte, 1024)
		n, err := clientSk.Read(res)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("服务器响应：", string(res[:n]))
	}
}
