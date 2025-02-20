package main

import "fmt"

func add5(a, b int) int {
	return a + b // 函数的返回语句 函数终止语句
}

// 返回值为空
func printInfo(name string, age int) {
	fmt.Printf("姓名：%s，年龄：%d\n", name, age)
	return // 等同于不写return 返回空
}

func getInfo() (string, int) {
	return "l", 10
}

func foo() (z int) {
	return
}

func main() {
	// 返回一个值
	res := add5(1, 2)
	fmt.Println(res)

	// 返回空值
	printInfo("l", 18) // 无返回值，不能赋值给其他变量

	// 返回多个值
	ret1, ret2 := getInfo()
	fmt.Println(ret1, ret2)

	// 返回值命名

}
