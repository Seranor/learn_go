package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var str1 string = "hello"
	var str2 string = "hello,学习go"
	// 遍历
	for i := 0; i < len(str1); i++ {
		fmt.Printf("ascii: %c %d\n ", str1[i], str1[i])
	}
	for _, s := range str1 {
		fmt.Printf("unicode: %c %d\n ", s, s)
	}
	// 中文只能用range方式
	for _, s := range str2 {
		fmt.Printf("unicode: %c %d\n ", s, s)
	}

	// Sprint 结果会以字符串形式返回
	result := fmt.Sprintf("%s", str1)
	fmt.Println(result)

	num1 := 99
	res2 := fmt.Sprintf("%d", num1)
	fmt.Printf("%T %s\n", res2, res2) // 直接将 99 转为 字符串 99

	/*
		%c  单一字符
		%T  动态类型
		%v  本来值的输出
		%+v 字段名+值打印
		%d  十进制打印数字
		%p  指针，十六进制
		%f  浮点数
		%b 二进制
		%s string
	*/
	// 查找
	tracer := "张三来了,张三bye bye"

	// 正向搜索字符串
	comma := strings.Index(tracer, ",")
	fmt.Println(",所在的位置:", comma)
	fmt.Println(tracer[comma+1:]) // 张三bye bye

	add := strings.Index(tracer, "+")
	fmt.Println("+所在的位置:", add) // +所在的位置: -1

	pos := strings.Index(tracer[comma:], "张三")
	fmt.Println("张三，所在的位置", pos) // 张三，所在的位置 1

	fmt.Println(comma, pos, tracer[5+pos:]) // 12 1 张三bye bye

	/*
		在必要以及可行的情况下，一个类型的值可以被转换成另一种类型的值。由于Go语言不存在隐式类型转换，因此所有的类型转换都必须显式的声明：
		类型 B 的值 = 类型 B(类型 A 的值)
		valueOfTypeB = type B(valueOfTypeA)

		类型转换只能在定义正确的情况下转换成功，例如从一个取值范围较小的类型转换到一个取值范围较大的类型（将 int16 转换为 int32）
		当从一个取值范围较大的类型转换到取值范围较小的类型时（将 int32 转换为 int16 或将 float32 转换为 int），会发生精度丢失的情况
		只有相同底层类型的变量之间可以进行相互转换（如将 int16 类型转换成 int32 类型），不同底层类型的变量相互转换时会引发编译错误（如将 bool 类型转换为 int 类型）
	*/
	a := 5.0
	b := int(a)
	fmt.Printf("%T, %d\n", b, b)

	// 输出各数值范围
	fmt.Println("int8 range:", math.MinInt8, math.MaxInt8)
	fmt.Println("int16 range:", math.MinInt16, math.MaxInt16)
	fmt.Println("int32 range:", math.MinInt32, math.MaxInt32)
	fmt.Println("int64 range:", math.MinInt64, math.MaxInt64)
	// 初始化一个32位整型值
	var aa int32 = 1047483647
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int32: 0x%x %d\n", aa, aa)
	// 将a变量数值转换为十六进制, 发生数值截断
	bb := int16(aa)
	// 输出变量的十六进制形式和十进制值
	fmt.Printf("int16: 0x%x %d\n", bb, bb)
	// 将常量保存为float32类型
	var cc float32 = math.Pi
	// 转换为int类型, 浮点发生精度丢失
	fmt.Println(int(cc))

	/*
		Golang语言的字符串是不可变的
		修改字符串时，可以将字符串转换为[]byte进行修改
		[]byte和string可以通过强制类型转换
	*/

	s1 := "localhost:8080"
	fmt.Println(s1)

	// 强制转为byte
	strByte := []byte(s1)
	fmt.Println(strByte)
	// 下标修改
	strByte[len(s1)-1] = '1'
	fmt.Println(strByte)

	// 再强制转换 string
	s2 := string(strByte)
	fmt.Println(s2)
}
