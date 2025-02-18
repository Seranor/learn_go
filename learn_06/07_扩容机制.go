package main

import "fmt"

func main() {
	l := make([]int, 5, 10)
	v1 := append(l, 1)
	fmt.Println(v1)    // [0 0 0 0 0 1]
	v2 := append(l, 2) // 会把v1 的 1 改成 2 v1和v2 是同一个空间
	fmt.Println(v2)    // [0 0 0 0 0 2]
	fmt.Println(v1)    // [0 0 0 0 0 2]
	fmt.Println(l)
}
