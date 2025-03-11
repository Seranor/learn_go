package main

import (
	"fmt"
	"sync"
)

var a int
var wg sync.WaitGroup
var mu sync.Mutex

func add() {
	defer wg.Done()
	mu.Lock()
	for i := 0; i < 1000; i++ {
		a += i
	}
	mu.Unlock()
}
func sub() {
	defer wg.Done()
	mu.Lock()
	for i := 0; i < 100; i++ {
		a -= i
	}
	mu.Unlock()
}
func main() {

	wg.Add(2)
	go add()
	go sub()
	wg.Wait()
	fmt.Println(a)
}
