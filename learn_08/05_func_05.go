package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
	以函数作为参数
	以函数作为返回值
*/

/*
	func add6(x, y int) int {
		return x + y
	}

	func main() {
		var x = 10
		fmt.Println(x)
		fmt.Println(add6)
	}
*/
func foo2() {
	fmt.Println("foo2 function... start")
	time.Sleep(1 * time.Second)
	fmt.Println("foo2 function... end")
}
func bar() {
	fmt.Println("bar function... start")
	time.Sleep(2 * time.Second)
	fmt.Println("bar function... end")
}

func timer(f func()) { // 完整的函数类型，需要带上形参 和返回值等信息
	start := time.Now().Unix()
	f()
	end := time.Now().Unix()
	fmt.Println("时间花销：", end-start)
}

func addT(x, y int) int {
	time.Sleep(1 * time.Second)
	return x + y
}
func timer2(f func(int, int) int, x int, y int) int {
	start := time.Now().Unix()
	res := f(x, y)
	end := time.Now().Unix()
	fmt.Println("时间花销：", end-start)
	return res
}
func outer() func(int, int) int {
	inner := func(x, y int) int {
		fmt.Println("内部inner")
		return x + y
	}
	return inner
}
func main() {
	timer(bar)
	fmt.Println(reflect.TypeOf(addT)) // func(int, int) int
	res := timer2(addT, 1, 2)
	fmt.Println(res)
	res2 := outer()(9, 9)
	fmt.Println(res2)
}
