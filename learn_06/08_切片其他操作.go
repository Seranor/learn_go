package main

import (
	"fmt"
	"sort"
)

func main() {
	// append 追加元素
	/*	var arr = []int{1, 2, 3, 4}
		var s1 = append(arr, 5)
		fmt.Println(s1)
		var arr2 = []int{5, 6, 7}
		var s2 = append(arr, arr2...) // ... 代表多元素 才能追加进去
		fmt.Println(s2)*/

	// 删除元素
	/*	var arr = []int{1, 2, 3, 4, 5}
		var i = 2
		//s1 := arr[:i]
		//s2 := arr[i+1:]
		//arr = append(s1, s2...)
		//fmt.Println(arr)

		s := append(arr[:i], arr[i+1:]...)
		fmt.Println(s)*/

	// 在开头位置插入元素方法
	/*	var arr = []int{1, 2, 3, 4, 5}
		var val = 6
		fmt.Println(append([]int{val}, arr...))*/

	// 在任意位置插入元素
	/*	var arr = []int{1, 2, 3, 4, 5}
		var val = 6
		var index = 2
		var s1 = arr[:index]
		var s2 = arr[index:]
		arr = append(s1, append([]int{val}, s2...)...)
		fmt.Println(arr)*/

	// 排序
	var a = []int{12, 4, 45, 6, 7}
	sort.Ints(a)
	fmt.Println(a)
}
