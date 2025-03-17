package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		panic("连接失败")
	}
	var reply string
	err = client.Call("HelloService.Hello", "World", &reply)
	if err != nil {
		panic("调用失败")
	}
	fmt.Println(reply)
}
