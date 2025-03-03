package main

import "fmt"

type Sleep6 interface {
	sleep()
}
type Runner6 interface {
	run()
}
type Animal6 interface {
	Sleep6
	Runner6
}

type Dog6 struct {
}

func (d Dog6) sleep() {
	fmt.Printf("sleeping\n")
}
func (d Dog6) run() {
	fmt.Printf("running\n")
}
func main() {
	var d = Dog6{}

	var a Animal6

	a = d
	a.sleep()
	a.run()
}
