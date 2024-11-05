package main

import "fmt"

/*
byte 类型是 uint8 的别名，rune 类型是int32的别名

# ASCII 码的一个字符占一个字节

# ASCII 定义 128 个字符，由码位 0 – 127 标识。它涵盖英文字母，拉丁数字和其他一些字符

# Unicode 是 ASCII 的超集，它定义了 1,114,112 个代码点的代码空间。 Unicode 版本 10.0 涵盖 139 个现代和历史文本集（包括符文字母，但不包括 Klingon ）以及多个符号集

# Go语言同样支持 Unicode（UTF-8）, `用rune来表示`, 在内存中使用 int 来表示

在书写 Unicode 字符时，需要在 16 进制数之前加上前缀`\u`或者`\U`。如果需要使用到 4 字节，则使用`\u`前缀，如果需要使用到 8 个字节，则使用`\U`前缀
*/
func main() {
	//使用单引号 表示一个字符
	//var ch byte = 'A'
	//在 ASCII 码表中，A 的值是 65,也可以这么定义
	//var ch byte = 65
	//65使用十六进制表示是41，所以也可以这么定义 \x 总是紧跟着长度为 2 的 16 进制数
	//var ch byte = '\x41'
	//65的八进制表示是101，所以使用八进制定义 \后面紧跟着长度为 3 的八进制数
	//var ch byte = '\101'
	//fmt.Printf("%c", ch)

	var ch rune = '\u0041'
	var ch1 int64 = '\U00000041'
	//格式化说明符%c用于表示字符，%v或%d会输出用于表示该字符的整数，%U输出格式为 U+hhhh 的字符串
	fmt.Printf("%c,%c,%U", ch, ch1, ch)
}
