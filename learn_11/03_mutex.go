package main

import (
	"fmt"
	"sync"
)

var lock sync.Mutex
var wg3 sync.WaitGroup
var x = 0

func add1() {
	defer wg3.Done()
	lock.Lock() //加锁操作
	x++
	lock.Unlock() // 解锁操作
}
func main() {
	wg3.Add(1000)
	for i := 0; i < 1000; i++ {
		go add1()
	}
	wg3.Wait()
	fmt.Println(x)
}
