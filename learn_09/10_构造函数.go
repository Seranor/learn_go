package main

import "fmt"

type Student5 struct {
	name   string
	age    int
	gender byte
	score  [3]float64
}

// 构造函数

func NewStudent(name string, age int, gender byte, score [3]float64) *Student5 {
	return &Student5{
		name:   name,
		age:    age,
		gender: gender,
		score:  score}
}
func main() {
	/*	var s = Student5{
		name:   "xiaoxiao",
		age:    18,
		gender: 1,
		score:  [3]float64{100, 99, 88}}
	fmt.Println(s)*/
	var s = NewStudent("xiaoxiao", 18, 1, [3]float64{100, 99, 88})
	fmt.Println(s)
}
