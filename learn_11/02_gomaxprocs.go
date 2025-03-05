package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func foo3() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
		time.Sleep(time.Millisecond * 20)
	}
	wg2.Done()
}
func bar() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
		time.Sleep(time.Millisecond * 30)
	}
	wg2.Done()
}

var wg2 sync.WaitGroup

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1) // 限制物理CPU核
	wg2.Add(2)
	go foo3()
	go bar()
	wg2.Wait()
}
