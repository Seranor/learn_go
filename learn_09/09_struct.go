package main

import (
	"fmt"
	"reflect"
)

type Student4 struct {
	name   string
	age    int
	gender byte
	score  [3]float64
}

func main() {
	var s1 = new(Student4) // &Student4{}
	fmt.Println(s1, reflect.TypeOf(s1))
	s1.name = "lll"
	fmt.Println(s1.name, s1.age)
}
