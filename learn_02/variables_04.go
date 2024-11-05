package main

import "fmt"

// 一个变量（常量、类型或函数）在程序中都有一定的作用范围，称之为作用域
/*
	Go语言(静态语言)会在编译时检查每个变量是否使用过，一旦出现未使用的变量，就会报编译错误。
	- 函数内定义的变量称为局部变量
	- 函数外定义的变量称为全局变量
	- 函数定义中的变量称为形式参数
*/

// 声明全局变量
var cc int
var bb float32 = 3.14

func main() {
	// 局部变量不是一直存在的 值在定义它的函数被调用后存在，函数调用结束后局部变量就会被销毁
	// 声明局部变量
	var a, b int
	a = 2
	b = 3
	cc = a + b
	fmt.Println(a, b, cc)

	bb := 3
	fmt.Println(bb) // 全局变量和局部变量名称可以相同，但是在函数体内的局部变量会被优先考虑

	c := sum(a, b)
	fmt.Println(c)
}

// 在定义函数时函数名后面括号中的变量交形式参数，形参只在函数调用时才会生效，函数调用结束后会被销毁，在函数未被调用时，函数的形参并不占用实际的存储单元，也没有实际值
func sum(a, b int) int {
	fmt.Printf("sum() 函数中 a = %d\n", a)
	fmt.Printf("sum() 函数中 b = %d\n", b)
	num := a + b
	return num
}
