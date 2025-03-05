package main

import (
	"fmt"
	"sync"
)

func read(c chan string) {
	defer wg5.Done()
	fmt.Println("协程监听读...")
	for {
		value := <-c
		fmt.Println(value)
	}
}
func write(c chan string) {
	defer wg5.Done()
	fmt.Println("协程监听写...")
	for {
		fmt.Println(">>>")
		var content string
		fmt.Scan(&content)
		c <- content
	}
}

var wg5 sync.WaitGroup

func main() {
	// 无缓冲管道
	var c = make(chan string)
	wg5.Add(2)
	go read(c)
	go write(c)
	wg5.Wait()

}
