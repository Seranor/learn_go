package main

import "fmt"

func main() {
	/*	// 先注册的后执行
		defer fmt.Println("hello world 1...")
		defer fmt.Println("hello world 2...")
		defer fmt.Println("hello world 3...") // 先执行
		fmt.Println("hello world...")*/

	/*	foo := func() {
			fmt.Println("foo1")
		}
		defer foo() // foo1
		foo = func() {
			fmt.Println("foo2")
		}*/
	/*	x := 10
		defer func(a int) {
			fmt.Println(a) // 10   会把函数体和形参都拷贝
		}(x)
		x++*/

	x := 10
	defer func() {
		fmt.Println(x) // 11 内部的x和外部的x地址一样
	}()
	x++
}
