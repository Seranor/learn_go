package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Client struct {
	Ip   string
	conn net.Conn
}

type Msg struct {
	Content string
	User    string
}

var onlineClients = make(map[string]Client) // map[127.0.0.1:9321:{127.0.0.1:9321 0xc00008e078}]

func readClients(conn net.Conn) {
	for {
		data := make([]byte, 1024)
		n, err := conn.Read(data)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(string(data[:n]))
		//将某一个客户端发的数据写到广播管道
		msg := Msg{Content: string(data[:n]), User: conn.RemoteAddr().String()}
		//fmt.Println(msg)
		broadcastChan <- msg
	}
}

func handleConn(conn net.Conn) {
	// 获取客户端的IP和端口
	ip := conn.RemoteAddr().String()
	fmt.Println(ip)
	// 实例化client对象
	client := Client{ip, conn}
	// 将该客户端对象添加到  onlineClients
	onlineClients[ip] = client
	fmt.Println(onlineClients)
	msg := Msg{Content: ip + "上线", User: "system:"}
	broadcastChan <- msg

	// 监听每个客户端的socket输入
	go readClients(conn)

}

var broadcastChan = make(chan Msg)

func broadcastManager() {
	for {
		// 监听广播管道中的数据
		msg := <-broadcastChan
		fmt.Println(msg)
		// 将该消息发送给所有的客户端
		for _, client := range onlineClients {
			msgJson, _ := json.Marshal(msg)
			client.conn.Write(msgJson)
		}
	}
}
func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务启动", listener.Addr())

	// 广播协程
	go broadcastManager()

	for {
		conn, _ := listener.Accept()
		go handleConn(conn)
	}
}
