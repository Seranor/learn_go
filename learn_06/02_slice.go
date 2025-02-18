package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr = [7]int{13, 24, 34, 45, 67, 88, 99}
	fmt.Println(arr, reflect.TypeOf(arr))

	// 切片创建 切割数组或者切片操作得到
	s1 := arr[2:5]
	fmt.Println(s1, reflect.TypeOf(s1))
	s2 := arr[3:6]
	fmt.Println(s2, reflect.TypeOf(s2))
	s3 := arr[:]
	fmt.Println(s3, reflect.TypeOf(s3))
	s1[0] = 100
	fmt.Println(arr, s1, s2)
	s4 := s3[:4]
	fmt.Println(s4, reflect.TypeOf(s4))

	// 直接声明切片
	var slice = []int{111, 222, 333, 444, 555}
	fmt.Println(slice, reflect.TypeOf(slice), len(slice))
}
