package main

import (
	"fmt"
)

type Addr struct {
	country, province, city string
}
type Student6 struct {
	name string
	//address Addr
	//int  // int int
	Addr
	//province string  先找这里的，找不到再去下面的 struct 查找，由外向里
}

func main() {
	/*	s := Student6{"kk", 18}
		fmt.Println(s)
		fmt.Println(s.int)*/

	/*	s := Student6{name: "alvin"}
		addr := Addr{country: "CN", province: "sc", city: "cd"}
		s.address = addr
		fmt.Println(s)
		fmt.Println(s.address.city)*/

	s := Student6{name: "alvin"}
	s.Addr = Addr{country: "CN", province: "sc", city: "cd"}
	fmt.Println(s)
	fmt.Println(s.Addr.province)
	fmt.Println(s.province)
}
