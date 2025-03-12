package main

import (
	"fmt"
	"time"
)

// 可以实现交替执行两个函数
var number, letter = make(chan bool), make(chan bool)

func printNum() {
	i := 0
	for {
		<-number
		fmt.Printf("%d%d", i, i+1)
		i += 2
		letter <- true
	}
}
func printLetter() {
	i := 0
	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for {
		if i >= len(str) {
			return
		}
		<-letter
		fmt.Print(str[i : i+2])
		i += 2
		number <- true
	}
}
func main() {
	go printNum()
	go printLetter()
	number <- true
	time.Sleep(time.Second * 5)

}
