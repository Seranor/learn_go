package main

import "fmt"

func main() {
	// 值类型 int  string bool []type
	/*	x := 10
		var y = x  // 单独开辟了空间 值拷贝
		fmt.Printf("x的值是%d 地址是%p\n", x, &x)
		fmt.Printf("y的值是%d 地址是%p\n", y, &y)*/

	//a := 100
	//fmt.Printf("a的地址是%p\n", &a)
	//a = 1000
	//fmt.Printf("a的地址是%p\n", &a)

	//var s []int
	//s[0] = 1  // 报错，没有开辟空间 是引用类型
	//fmt.Println(s)

	var a = [3]int{1, 2, 3}
	fmt.Printf("a:%p\n", &a)
	fmt.Printf("a:%p\n", &a[0])
	fmt.Printf("a:%p\n", &a[1])
	fmt.Printf("a:%p\n", &a[2])

}
