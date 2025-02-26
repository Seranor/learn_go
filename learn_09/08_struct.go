package main

import (
	"fmt"
	"reflect"
)

type Student3 struct {
	name   string
	age    int
	gender byte
	score  [3]float64
}

func change(s *Student3) {
	s.age = 32
	//(*s).age = 32
}
func main() {
	// &结构体{}
	var s1 = &Student3{
		name:   "lzj",
		age:    22,
		gender: 1,
		score:  [3]float64{85, 92, 79}}
	fmt.Println(reflect.TypeOf(s1)) // *main.Student3
	fmt.Println((*s1).name)         // *s1.name 其中 . 优先级高于*
	fmt.Println(s1.name)            // 语法糖多做了一件事情

	var s2 = Student3{
		name:   "rain",
		age:    25,
		gender: 1,
		score:  [3]float64{88, 99, 77}}
	change(&s2)
	fmt.Println(s2.age)
}
