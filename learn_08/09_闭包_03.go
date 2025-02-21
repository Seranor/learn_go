package main

import "fmt"

func foo4() {
	fmt.Println("foo3 function 调用")
}
func bar3() {
	fmt.Println("bar3 function 调用")
}
func counter(targetFunc func()) func() int {
	var count = 0
	return func() int {
		targetFunc()
		count++
		return count
	}
}
func main() {
	/*	f := counter(bar3)
		f()
		f()
		n := f()
		fmt.Println(n)*/
	foo3 := counter(foo4)
	foo3()
	foo3()
	n := foo3()
	fmt.Println(n)
}
