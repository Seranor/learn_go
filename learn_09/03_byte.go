package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x byte
	x = 'a'
	fmt.Println(x, reflect.TypeOf(x))
}
