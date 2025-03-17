package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func g1(ch chan struct{}) {
	time.Sleep(2 * time.Second)
	ch <- struct{}{}
}
func g2(ch chan struct{}) {
	time.Sleep(3 * time.Second)
	ch <- struct{}{}
}

func main() {

	g1Channel := make(chan struct{})
	g2Channel := make(chan struct{})
	//g1Channel <- struct{}{}
	//g2Channel <- struct{}{}
	go g1(g1Channel)
	go g2(g2Channel)

	timer := time.NewTimer(5 * time.Second)
	for {
		select {
		case <-g1Channel:
			fmt.Println("g1 done")
		case <-g2Channel:
			fmt.Println("g2 done")
		case <-timer.C:
			fmt.Println("time out")
			return
		}
	}
}
