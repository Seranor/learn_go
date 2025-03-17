package main

import (
	"net"
	"net/rpc"
)

type HelloService struct{}

func (s *HelloService) Hello(request string, reply *string) error {
	// 返回值是通过修改reply的值
	*reply = "Hello " + request
	return nil
}

func main() {
	listener, _ := net.Listen("tcp", ":8899")

	_ = rpc.RegisterName("HelloService", &HelloService{})

	conn, _ := listener.Accept()
	rpc.ServeConn(conn)

}
