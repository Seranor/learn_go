package main

import (
	"fmt"
	"sync"
)

var a int
var wg sync.WaitGroup
var mu sync.Mutex

// 锁不能复制 lock := mu 否则就失去作用了
func add() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		a += i
		mu.Unlock()
	}
}
func sub() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mu.Lock()
		a -= i
		mu.Unlock()
	}
}
func main() {

	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(a)
}
