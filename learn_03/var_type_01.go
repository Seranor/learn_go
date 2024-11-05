package main

import "fmt"

/*

Go语言同时提供了有符号和无符号的整数类型。
* 有符号整型：int、int8、int64、int32、int64
* 无符号整型：uint、uint8、uint64、uint32、uint64、uintptr
有符号整型范围：`-2^(n-1) 到 2^(n-1)-1`
无符号整型范围: ` 0 到 2^n-1`

rune int32  Unicode码点
byte uint8

无符号的整数类型 `uintptr`，它没有指定具体的 bit 大小但是足以容纳指针。uintptr 类型只有在`底层编程`时才需要，
特别是Go语言和C语言函数库或操作系统接口相交互的地方
*/

func main() {
	// 通常应该优先使用 float64 类型，因为 float32 类型的累计计算误差很容易扩散，并且 float32 能精确表示的正整数并不是很大
	floatNum := 3.2
	fmt.Printf("floatNum: %.2f\n", floatNum)

	var f1 float32 = 1 << 24
	fmt.Println(f1)
	fmt.Println(f1 == f1+1) // true

	var e = .71828
	var f2 = 1.
	fmt.Printf("e: %.5f\nf2: %.1f\n", e, f2)

}

/*
Go语言对于值之间的比较有非常严格的限制，只有两个相同类型的值才可以进行比较，如果值的类型是接口（interface）
那么它们也必须都实现了相同的接口

如果其中一个值是`常量`，那么另外一个值可以不是常量，但是类型必须和该常量类型相同

如果以上条件都不满足，则必须将其中一个值的类型转换为和另外一个值的类型相同之后才可以进行比较
*/
