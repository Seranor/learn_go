package main

import (
	"fmt"
	"sync"
	"time"
)

// 锁的本质是将并行的代码串行化 使用lock肯定会影响性能
// 设计所也需要尽量保证并行
// 读和写之间  读和读之间
// 读写锁

func main() {
	var rwlock sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(2)
	// 写的goruntine
	go func() {
		defer wg.Done()
		rwlock.Lock() // 加写锁，会分支别的写锁和读锁获取
		time.Sleep(5 * time.Second)
		fmt.Println("get write lock")
		rwlock.Unlock()
	}()
	// 读的goruntine
	go func() {
		defer wg.Done()
		for {
			rwlock.RLock() // 加读锁，不会阻止别人的读
			time.Sleep(time.Millisecond * 500)
			fmt.Println("get read lock")
			rwlock.RUnlock()
		}
	}()
	wg.Wait()
}
