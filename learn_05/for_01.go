package main

import "fmt"

func main() {
	/*	var count = 1
		for count <= 100 {
			fmt.Printf("hello world! num: %d \n", count)
			count++
		}*/

	/*	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}*/

	/*	var res = 0
		for i := 1; i <= 100; i++ {
			//fmt.Println(i)
			res = i + res
		}
		fmt.Println(res)*/

	for i := 1; i <= 100; i++ {
		if i == 88 {
			//break // 退出整个循环
			continue // 退出本次循环
		}
		fmt.Println(i)
	}
}
