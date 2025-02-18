package main

import (
	"fmt"
)

func main() {
	//var is_login = true
	// 单分支
	/*	if is_login {
		fmt.Println("login")
	}*/

	// 双分支
	/*	if is_login {
			fmt.Println("is login")
		} else {
			fmt.Println("is not login")
		}*/

	//fmt.Println("end")

	var score = 92
	if _, ok := interface{}(score).(int); ok {
		if score >= 90 {
			fmt.Println("very good.")
		} else if score >= 80 {
			fmt.Println("good")
		} else if score >= 60 {
			fmt.Println("C")
		} else {
			fmt.Println("D")
		}
	} else {
		fmt.Println("格式错误")
	}
}
