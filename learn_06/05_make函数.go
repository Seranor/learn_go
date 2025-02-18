package main

import "fmt"

func main() {
	// 针对引用数据类型，需要声明开辟空间的方法：make
	//var s []int
	//s[0] = 1

	var s = make([]int, 5, 10) // [0 0 0 0 0       ]
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	s[0] = 100
	fmt.Println(len(s), cap(s))
}
