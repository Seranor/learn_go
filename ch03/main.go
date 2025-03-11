package main

import (
	"fmt"
	"time"
)

func asyncPrint() {
	for {
		time.Sleep(1 * time.Second)
		fmt.Println("hello world")
	}

}

// 主协程
func main() {
	// 主死随从
	//go asyncPrint()

	for i := 0; i < 100; i++ {
		//tmp := i  // go 低版本值拷贝解决出现的for 问题
		
		go func() {
			fmt.Println(i)
		}()
	}

	fmt.Println("this is main")
	time.Sleep(10 * time.Second)
}
