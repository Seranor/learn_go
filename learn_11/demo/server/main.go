package main

import (
	"fmt"
	"net"
	"strings"
)

func connHandler(conn net.Conn) {
	for {
		var data = make([]byte, 1024)
		n, err := conn.Read(data) // 阻塞函数
		if err != nil {
			// 如果是 EOF 或其他错误，跳出循环
			if err.Error() == "EOF" {
				fmt.Println("Client disconnected.")
			} else {
				fmt.Println("Error reading data:", err)
			}
			break
		}

		dataStr := string(data[:n])
		if dataStr == "exit" {
			fmt.Println("Exit command received. Closing connection.")
			break
		}

		// 打印接收到的数据
		fmt.Println("Received:", dataStr)

		// 将处理好的数据响应给客户端
		upperStr := strings.ToUpper(dataStr)
		conn.Write([]byte(upperStr))
	}
}
func main() {
	// 建立服务
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	// 监听客户端的连接
	fmt.Println("Listening on", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("Accepting connection", conn)
		go connHandler(conn)
	}
}
