package main

import (
	"fmt"
	"reflect"
)

// 声明在函数中的参数叫形式参数
func add(a, b int) {
	fmt.Println(a + b)
}

func lingXin(n int) {
	// 打印上半部分（包括中间行）
	for i := 1; i <= n; i++ {
		// 打印空格
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		// 打印 *
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		// 换行
		fmt.Println()
	}
	// 打印下半部分
	for i := n - 1; i >= 1; i-- {
		// 打印空格
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		// 打印 *
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		// 换行
		fmt.Println()
	}
}

func info(name string, age int) {
	fmt.Printf("姓名：%s，年龄：%d\n", name, age)
}

// 不定长参数
func add3(x ...int) {
	fmt.Println(x, reflect.TypeOf(x))
	var ret int
	for _, v := range x {
		ret += v
	}

	fmt.Println(ret)
}

func add4(name string, x ...int) {
	fmt.Printf("%s 调用add4函数\n", name)
	var ret int
	for _, v := range x {
		ret += v
	}

	fmt.Println(ret)
}
func main() {
	// 函数调用  函数名()
	lingXin(6) // 调用函数时 传入的值 为 实际参数
	add(1, 2)  // 按位置传参
	lingXin(6)
	info("l", 18)
	add3(1, 2, 3)
	add4("l", 1, 2, 3, 4, 5) // 后面一堆参数都是 x 的
}
