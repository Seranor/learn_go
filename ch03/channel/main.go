package main

import (
	"fmt"
	"time"
)

func producer(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i * i
	}
	close(out)
}
func consumer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}
func main() {
	/*
		不要通过共享内存来通信，而要通过通信来共享内存
	*/

	//var msg chan string
	//msg = make(chan string, 0)
	///*	msg = make(chan string, 1) // channel 的初始化值为0的话，放值会阻塞，设置为0 叫无缓冲channel
	//	msg <- "hello"
	//	fmt.Println(<-msg)*/
	//
	//go func(msg chan string) { // happen-before 机制
	//	data := <-msg
	//	fmt.Println(data)
	//
	//}(msg)
	//msg <- "hello"
	//time.Sleep(5 * time.Second) // waitgroup 少了 done 调用，容易出现deadlock。还有无缓冲channel

	// 无缓冲channel 适用于 通知 B 要第一时间知道A是否已经完成
	// 有换成channel 适用于消费者和生产者之间的通信
	/*
		消息传递 消息过滤
		信号广播
		事件订阅和广播
		任务分发
		结果汇总
		并发控制
		同步和异步
	*/

	/*
		var msg2 = make(chan int, 2)
		go func(msg2 chan int) {
			for data := range msg2 {
				fmt.Println(data)
			}
			fmt.Println("done")
		}(msg2)
		msg2 <- 1
		msg2 <- 2
		close(msg2) // 已经关闭的channel不能放值了  但是可以继续取值
		time.Sleep(5 * time.Second)*/

	// 默认情况下 channel 是双向的
	// 单向操作
	//var ch1 chan int       // 双向的
	//var ch2 chan<- float64 // 单向写
	//var ch3 <-chan int     // 单向读

	/*	c := make(chan int, 3)
		var send chan<- int = c
		var recv <-chan int = c

		send <- 1
		fmt.Println(<-recv)*/

	// 在函数传参的时候也可以直接定义参数类型 func xxx(c chan<- int){}
	// func xxx(c <-chan int){}  都可以传入 c := make(chan int, 3)
	c := make(chan int)
	go producer(c)
	go consumer(c)
	time.Sleep(5 * time.Second)

}
