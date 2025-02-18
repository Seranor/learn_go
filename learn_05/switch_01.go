package main

import "fmt"

func main() {
	s := `
	1. print 1
	2. print 2
	3. print 3
	4. print 4
`
	fmt.Println(s)
	var choice int
	fmt.Printf("输入选择：")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("0.0")
	}
}
