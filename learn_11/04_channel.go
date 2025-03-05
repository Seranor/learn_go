package main

import (
	"fmt"
)

type Stu struct {
	name string
	age  int
}

func main() {
	// 声明一个 int管道类型
	// var c chan int  // nil 需要 make
	c := make(chan int, 10)
	fmt.Println(c)

	// 传值
	c <- 10
	c <- 20
	c <- 30
	// 取值
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)

	var c2 = make(chan interface{}, 10)
	c2 <- 123
	c2 <- "hello"
	c2 <- Stu{"tom", 20}

	fmt.Println(<-c2)
	fmt.Println(<-c2)
	a := <-c2
	fmt.Println(a.(Stu).age)
	v, ok := a.(Stu) // 空接口需要断言来处理类型
	if ok {
		fmt.Println(v.name)
	}

	fmt.Println(a)

	/*	var m1 = make(map[string]int, 10)
		m1["a"] = 100
		k, b := m1["a"] // b是bool 有就是true 没有就是false
		fmt.Println(k, b)*/

	var c3 = make(chan interface{}, 10)
	c3 <- 123
	c3 <- "hello"
	c3 <- Stu{"tom", 20}
	// 循环管道之前需要先关闭管道
	// 关闭管道之后只能读 不能写
	close(c3)
	for i := range c3 {
		fmt.Println("c3::: ", i)
	}
}
