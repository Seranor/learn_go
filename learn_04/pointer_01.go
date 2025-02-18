package main

import "fmt"

func main() {
	var cat int = 1
	var str string = "学习教程"
	ptr1 := &cat
	fmt.Printf("%p,%p\n", &cat, &str) // 内存地址
	fmt.Println(*ptr1)

	/*
		变量、指针和地址三者的关系是，每个变量都拥有地址，指针的值就是地址
	*/
	var room int = 18
	var ptr2 = &room
	fmt.Printf("%p\n", &room)
	fmt.Printf("%T,%p\n", ptr2, ptr2)

	fmt.Println("指针地址", ptr2)
	fmt.Println("指针地址的值", *ptr2)
}
