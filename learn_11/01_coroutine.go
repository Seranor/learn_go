package main

import (
	"fmt"
	"sync"
	"time"
)

func foo1() {
	// 任务完成
	defer wg.Done()

	fmt.Println("foo1 start")
	time.Sleep(time.Second * 2)
	fmt.Println("foo1 end")
}

func foo2() {
	// 任务完成
	defer wg.Done()

	fmt.Println("foo2 start")
	time.Sleep(time.Second * 3)
	fmt.Println("foo2 end")

}

var wg sync.WaitGroup

func main() {
	startTime := time.Now().Unix()
	wg.Add(2)
	go foo1()
	go foo2()

	wg.Wait() // 等待计数器归零
	// 当所有协程执行完毕后的统计时间
	endTime := time.Now().Unix()
	fmt.Println(endTime - startTime)
}
