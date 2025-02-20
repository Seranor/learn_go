package main

import "fmt"

func foo7() func() {

	var x = 100
	// 嵌套函数
	bar := func() {
		fmt.Println("匿名函数")
		fmt.Println(x) // 引用了一个外部非全局的自由变量 x  bar 就是闭包函数
	}
	// 返回嵌套函数
	return bar
}

func main() {
	f := foo7() // f 闭包函数
	f()
}
