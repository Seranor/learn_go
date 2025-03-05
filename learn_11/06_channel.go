package main

import (
	"fmt"
	"sync"
)

func read1(c <-chan string) { // 这个类型只能读  单向读
	defer wg6.Done()
	fmt.Println("协程监听读...")
	for {
		value := <-c
		fmt.Println(value)
	}
}

func write1(c chan<- string) { // 这个类型只能写  单向写
	defer wg6.Done()
	fmt.Println("协程监听写...")
	for {
		fmt.Println(">>>")
		var content string
		fmt.Scan(&content)
		c <- content
	}
}

var wg6 sync.WaitGroup

func main() {
	var c = make(chan string)
	wg6.Add(2)
	go read1(c)  // 读和写都可以传入 隐式转换
	go write1(c) // 读和写都可以传入 隐式转换
	wg6.Wait()

}
