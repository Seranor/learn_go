package main

import "fmt"

func main() {
	var fn [10]func()
	for i := 0; i < len(fn); i++ {
		fn[i] = func() {
			fmt.Println(i)
		}
	}
	for _, fn := range fn {
		fn() // go 1.22之后 编译器优化了闭包捕获的方式
	}

	for i := 0; i < 10; i++ {
	}
}
