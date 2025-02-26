package main

import (
	"fmt"
	"reflect"
)

type Student2 struct {
	name   string
	age    int
	gender byte
	score  [3]float64
}

func main() {
	// 结构体{}
	// 1
	var s1 = Student2{
		name:   "lzj",
		age:    22,
		gender: 1,
		score:  [3]float64{85, 92, 79}}
	fmt.Println(s1, reflect.TypeOf(s1))
	fmt.Println(s1.name, s1.age, s1.gender, s1.score)

	// 2
	var s2 = Student2{
		"zzz",
		23,
		1,
		[3]float64{88, 72, 99}} // 必须是 }}
	fmt.Println(s2.score)
}
