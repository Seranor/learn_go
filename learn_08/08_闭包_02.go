package main

import "fmt"

func main() {
	var fn [10]func()

	for i := 0; i < len(fn); i++ {
		fn[i] = func() {
			fmt.Println(i)
		}
	}

	for _, f := range fn {
		f() // 10
	}
}
