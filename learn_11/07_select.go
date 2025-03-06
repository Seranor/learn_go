package main

import "fmt"

func main() {
	var ch1 = make(chan int, 10)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	var ch2 = make(chan int, 10)
	ch2 <- 10
	ch2 <- 20
	ch2 <- 30

	/*
		空select 长阻塞
		select {}
	*/

	select {
	case v1 := <-ch1: //监听的IO  大多数都是监听chan操作
		fmt.Println("ch1数据", v1)
	case v2 := <-ch2: //监听的IO  大多数都是监听chan操作
		fmt.Println("ch1数据", v2)
	}
}
