package main

import (
	"fmt"
	"net"
)

var age int

// 变量声明之后 系统自动赋予该类型的零值
// int 为 0
// float 为0.0
// bool 为 false
// string 为空字符串
// 指针为nil

var level = 1 // 编译时自动推导类型

// 批量定义
var (
	a int
	b float32
	c string
)

func main() {
	fmt.Printf("%d\n", age)   // 0
	fmt.Printf("%T\n", level) // int
	fmt.Printf("a=%d,b=%f,c=%s\n", a, b, c)
	// %d 整数占位符 %s 字符串占位符 %f 浮点数占位符(默认精度为6)

	// 简短格式定义变量
	i := 1
	fmt.Printf("%T \r\n", i)
	/*
		1.定义变量，同时显式初始化
		2.不能提供数据类型
		3.只能用在函数内部
	*/

	// 变量初始化出错的情况
	var level2 int = 1
	//level2 := 2  再次声明并赋值会报错 no new variables on left side of :=
	// 特例 net.Dial 提供按指定协议和地址发起网络连接，
	fmt.Printf("%d\n", level2)

	// 多重赋值
	//var conn net.Conn
	//var err error
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	conn1, err := net.Dial("tcp", "127.0.0.1:8000") // err 可以重复定义
	fmt.Println(conn)
	fmt.Println(conn1)
	fmt.Println(err)
	// 在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错
}
