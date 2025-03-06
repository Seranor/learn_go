package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Msg struct {
	Content string
	User    string
}

func read(conn net.Conn) {
	for {
		res := make([]byte, 1024)
		n, err := conn.Read(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		var msg Msg
		json.Unmarshal(res[:n], &msg)
		fmt.Printf("[%s]%s\n", msg.User, msg.Content)
	}
}
func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//从socket读取数据
	go read(conn)

	// 向socket发送数据
	for {
		var data string
		fmt.Scan(&data)
		conn.Write([]byte(data))
	}
}
