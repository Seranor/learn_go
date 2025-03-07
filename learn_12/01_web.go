package main

import (
	"fmt"
	"net"
)

func main() {
	Listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer Listener.Close()

	fmt.Println("服务启动", Listener.Addr().String())

	for {
		conn, err := Listener.Accept()
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(conn)

		data := make([]byte, 1024)
		n, _ := conn.Read(data)
		fmt.Println(string(data[:n]))

		conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\n <h1>hello world</h1>"))
	}
}
