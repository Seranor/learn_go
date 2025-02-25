package main

import "fmt"

func makeFun(i int) func() {
	return func() {
		fmt.Println(i)
	}
}
func main() {
	/*	var fn [10]func()
		for i := 0; i < len(fn); i++ {
			fn[i] = func() {
				fmt.Println(i)
			}
		}
		for _, fn := range fn {
			fn() // go 1.22之后 编译器优化了闭包捕获的方式
		}

		for i := 0; i < 10; i++ {
		}*/

	var fn2 [10]func()
	for i := 0; i < len(fn2); i++ {
		fn2[i] = makeFun(i)
	}
	for _, f2 := range fn2 {
		f2()
	}
}
