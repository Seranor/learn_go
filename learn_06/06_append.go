package main

import "fmt"

func main() {
	//var s = []int{1111, 2222, 333, 444}
	//fmt.Println(len(s), cap(s))

	//var s1 = make([]int, 5, 10)
	//s2 := append(s1, 555) // 此时cap 还有多余 不会开辟新的空间 操作s2 还是会影响到s1
	//fmt.Println(s1)
	//fmt.Println(s2, len(s2), cap(s2))
	//s2[0] = 100 // 会影响到s1
	//fmt.Println(s1)
	//fmt.Println(s2)

	//扩容机制
	var s = []int{1, 2, 3, 4}
	s2 := append(s, 5) // 此时复制开辟新的空间，和s已经没有联系了，不会影响到s了
	fmt.Println(s2, len(s2), cap(s2))
	s2[0] = 100
	fmt.Println(s)
	fmt.Println(s2)
}
