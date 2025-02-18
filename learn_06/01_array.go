package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*	var arr [3]int
		arr[0] = 1
		arr[1] = 3
		arr[2] = 5
		fmt.Println(arr)
		fmt.Println(arr[2])*/

	/*	arr2 := [3]int{11, 22, 33}
		fmt.Println(arr2)*/

	//var arr3 = [...]int{1, 2, 3, 4, 5, 6}
	//fmt.Println(arr3, reflect.TypeOf(arr3))

	/*	var arr4 = [3]int{0: 100, 2: 99} //[100 0 99]
		fmt.Println(arr4)*/

	var arr_name = [3]string{"张三", "李四", "王五"}
	fmt.Println(arr_name[1])
	fmt.Println(arr_name[0:2], reflect.TypeOf(arr_name[0:2])) //切片类型

	for i := 0; i < len(arr_name); i++ {
		fmt.Println(arr_name[i])
	}

	for k, v := range arr_name {
		fmt.Println(k, v)
	}
}
