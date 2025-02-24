package main

import "fmt"

func foo3(x int) {
	x = 100
}
func bar2(s []int) {
	fmt.Printf("bar2 函数的 s 地址 %p\n", s)
	s[0] = 100
	s = append(s, 4) // 扩容了
	fmt.Printf("bar2 函数的 s 地址2 %p\n", s)
	fmt.Println(s)
}
func f1(a *int) {
	*a = 100
	return
}
func main() {
	// 值传递
	//var x = 100
	//fmt.Println(&x)
	//var y = x
	//fmt.Println(&y)
	//x = 10
	//fmt.Println(x)
	//fmt.Println(y)

	var x = 1
	foo3(x) // 值传递 拷贝了一份给 x 形参
	fmt.Println(x)

	var s = []int{1, 2, 3}
	fmt.Printf("main 函数的 s 地址 %p\n", s)
	bar2(s)
	fmt.Println(s)

	var a = 10
	var p *int = &a
	f1(p)
	fmt.Println(a)
}
