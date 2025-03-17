package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var stop = make(chan struct{})

func cpuInfo(ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit")
			return
		default:
			time.Sleep(2 * time.Second)
			fmt.Println("cpu info")
		}
	}
}

var wg sync.WaitGroup

func main() {
	// goroutine 监控cpu的信息
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go cpuInfo(ctx)

	time.Sleep(time.Second * 6)
	cancel()
	wg.Wait()
	fmt.Println("done")
}
