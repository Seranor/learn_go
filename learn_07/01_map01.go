package main

import (
	"fmt"
	"reflect"
)

func main() {
	var stu = map[string]string{} // "":""
	fmt.Println(stu, reflect.TypeOf(stu))

	//stu["name"] = "l" 不能直接赋值，引用类型，还没有开辟空间
	var stu2 = make(map[string]string)
	fmt.Println(stu2, reflect.TypeOf(stu2))

	stu2["name"] = "liu"
	stu2["age"] = "18"
	fmt.Println(stu2) // map[age:18 name:liu] 无序性

	// 声明式
	var stu3 = map[string]string{"name": "liu", "age": "18", "gender": "male"}
	fmt.Println(stu3)

	// 访问map
	fmt.Println(stu3["age"])

	// range 循环
	for k, v := range stu3 {
		fmt.Println(k, v)
	}

	// 删除元素，内置
	delete(stu3, "age")
	fmt.Println(stu3)

}
