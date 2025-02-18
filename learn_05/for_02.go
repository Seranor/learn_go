package main

import "fmt"

func main() {
	//var s = "*"
	/*	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			fmt.Printf(s)
		}
		fmt.Println()
	}*/

	/*	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf(s)
		}
		fmt.Println()
	}*/

	var sum = 0
	for i := 1; i < 11; i++ {
		var ret = 1
		for j := 1; j < i+1; j++ {
			ret *= j
		}
		sum += ret
	}
	fmt.Println(sum)
}
