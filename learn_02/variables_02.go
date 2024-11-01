package main

import "fmt"

// 交换变量
func main() {
	//var a = 100
	//var b = 200
	//var c int
	//
	//c = b
	//b = a
	//a = c
	//
	//fmt.Printf("a=%d,b=%d,c=%d", a, b, c)

	//var a = 100
	//var b = 200
	//a = a ^ b
	//b = b ^ a
	//a = a ^ b
	//fmt.Printf("a=%d,b=%d", a, b)

	var a = 100
	var b = 200
	b, a = a, b
	fmt.Printf("a=%d,b=%d", a, b)
}
