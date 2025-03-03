package main

import "fmt"

type Sayer interface {
	say()
}
type Runner interface {
	run()
}
type Dog4 struct {
	name string
}

// 实现 Sayer 接口
func (d Dog4) say() {
	fmt.Printf("say hello\n")
}

// 实现 Runner 接口
func (d Dog4) run() {
	fmt.Printf("running\n")
}
func main() {
	// 一个类型实现多个接口
	var dog = Dog4{name: "huihui"}
	var s Sayer
	s = dog
	s.say()

	var r Runner
	r = dog
	r.run()
}
