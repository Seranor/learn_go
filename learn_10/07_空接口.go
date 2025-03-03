package main

import "fmt"

// 空接口没有定义任何方法的接口，因此任何类型都实现了空接口

type Dog7 struct {
	name string
}
type Cat7 struct {
	name string
}

func foo(x interface{}) {
	fmt.Printf("x的值是 %v，类型是%T\n", x, x)
}
func judge(x interface{}) {
	switch x.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	case bool:
		fmt.Println("bool")
	}
}
func main() {
	var x interface{}
	x = 10
	fmt.Println(x)

	x = "hello"
	fmt.Println(x)

	x = Dog7{name: "huihui"}
	dog, ok := x.(Dog7) // 断言
	if ok {
		fmt.Println(dog.name)
	} else {
		fmt.Println(x)
	}

	foo(100)
	foo("hello")
	foo(x)

	var m = make(map[string]interface{})
	m["name"] = "yy"
	m["age"] = 22
	m["score"] = []int{88, 77, 99}
	fmt.Println(m["score"])

	var a interface{}
	a = "hello world"
	v, ok := a.(int)
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("不是int")
	}
	judge("hello")
	judge(100)
	judge(true)
}
