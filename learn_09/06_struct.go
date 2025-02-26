package main

import "fmt"

// 声明结构体： 类: 值类型

type Student struct {
	name   string
	age    int
	gender byte
	score  [3]float32
}

func main() {
	// 声明一个结构体对象
	var stu01 Student
	fmt.Println(stu01.name)   // 默认零值
	fmt.Println(stu01.age)    // 默认零值
	fmt.Println(stu01.gender) // 默认零值
	fmt.Println(stu01.score)  // 默认零值

	// 初始化
	stu01.name = "xm"
	stu01.age = 22
	stu01.gender = 1
	stu01.score[0] = 84
	stu01.score[1] = 79
	stu01.score[2] = 88

	fmt.Println(stu01.name)
	fmt.Println(stu01.age)
	fmt.Println(stu01.gender)
	fmt.Println(stu01.score)
}
