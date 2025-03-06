package main

import (
	"fmt"
	"strconv"
	"time"
)

func pubInt(chanInt chan int) {
	// 每隔 1s 生产数字
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		chanInt <- i
	}
}
func pubString(chanStr chan string) {
	// 每隔 5s 生产字符串
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second * 5)
		chanStr <- "hello " + strconv.Itoa(i)
	}
}
func main() {
	var chanInt = make(chan int)
	var chanStr = make(chan string)

	// 构建两个生产者
	go pubInt(chanInt)
	go pubString(chanStr)

	for {
		select {
		case v1 := <-chanInt: //监听的IO  大多数都是监听chan操作
			fmt.Println("pubInt数据", v1)
		case v2 := <-chanStr: //监听的IO  大多数都是监听chan操作
			fmt.Println("chanStr数据", v2)
		}
	}
}
