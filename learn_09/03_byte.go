package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x byte
	x = 'a'
	fmt.Println(x, reflect.TypeOf(x)) // uint8  byte

	var y uint8
	y = 99
	fmt.Printf("%c,%d\n", y, y)

	var s = "hello world"
	fmt.Println(s, reflect.TypeOf(s))
	// 字符串(string)转为字节串([]byte)
	var b1 = []byte(s)
	fmt.Println(b1, reflect.TypeOf(b1))

	// 字节串转换为字符串
	fmt.Println(string(b1))

}
