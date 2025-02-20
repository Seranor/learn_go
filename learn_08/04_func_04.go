package main

import "fmt"

func main() {
	//func add(){} // 函数里不能使用这种直接方式声明，只能声明匿名函数

	// 匿名函数 自执行
	(func(x, y int) {
		fmt.Println(x + y)
	})(1, 3)

	// 匿名函数的赋值
	bar := func(x, y int) {
		fmt.Println(x + y)
	}
	bar(2, 3)
}
