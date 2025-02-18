package main

import "fmt"

func main() {
	/*	var names = [...]string{"张三", "李四", "王五", "赵六", "孙七"}
		s1 := names[0:3]
		s2 := names[2:5]
		fmt.Println(s1)
		fmt.Println(s2)
		s1[2] = "yuan"
		fmt.Println(s1)
		fmt.Println(s2)
		fmt.Println(names)*/

	var a = [...]int{1, 2, 3, 4, 5, 6}
	a1 := a[0:3]                                  // [1,2,3]
	a2 := a[0:5]                                  // [1,2,3,4,5]
	a3 := a[1:5]                                  // [2,3,4,5]
	a4 := a[1:]                                   // [2,3,4,5,6]
	a5 := a[:]                                    // [1,2,3,4,5,6]
	a6 := a3[1:2]                                 // [2]
	fmt.Printf("长度是%d,容量是%d\n", len(a1), cap(a1)) // 3 6
	fmt.Printf("长度是%d,容量是%d\n", len(a2), cap(a2)) // 5 6
	fmt.Printf("长度是%d,容量是%d\n", len(a3), cap(a3)) // 4 5
	fmt.Printf("长度是%d,容量是%d\n", len(a4), cap(a4)) // 5 5
	fmt.Printf("长度是%d,容量是%d\n", len(a5), cap(a5)) // 6 6
	fmt.Printf("长度是%d,容量是%d\n", len(a6), cap(a6)) // 1 4

}
