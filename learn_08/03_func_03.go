package main

import "fmt"

var y = 10 // 全局变量 适用于所有范围内
var x = 5

func foo1() {
	var x = 100
	fmt.Printf("foo1的x: %d\n", x)
	fmt.Printf("foo1的y: %d\n", y)
	y = 1

}
func main() {
	//foo1()
	//fmt.Printf("main函数的y: %d\n", y)

	// if 和 for 具备开辟作用域的能力
	var b = true
	if b {
		var x = 666
		fmt.Println(x)
	}
	fmt.Println(x)
}
