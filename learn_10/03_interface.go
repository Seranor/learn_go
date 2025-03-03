package main

import "fmt"

type Animal3 interface {
	sleep()
}
type Dog3 struct {
	name string
}

//func (d Dog3) sleep() {
//	fmt.Printf("%s was sleeping\n", d.name)
//}

func (d *Dog3) sleep() {
	fmt.Printf("%s was sleeping\n", d.name)
}

func main() {
	var a Animal3

	/*	a = Dog3{name: "maomao"} // 接口变量可以接收子类结构体对象
		a.sleep()
		a = &Dog3{name: "huihui"} // 接口变量也可以接收子类结构体指针对象
		a.sleep()
	*/

	//a = Dog3{name: "maomao"} // 不能接收该类型
	//a.sleep()

	a = &Dog3{name: "huihui"}
	a.sleep()
	// 接口指针变量

}
