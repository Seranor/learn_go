package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

/*
一个字符串是一个不可改变的字节序列，字符串可以包含任意的数据，但是通常是用来包含可读的文本，字符串是 UTF-8 字符的一个序列
go语言从底层就支持UTF-8编码
UTF-8 是一种被广泛使用的编码格式，是文本文件的标准编码
由于该编码对占用字节长度的不定性，在Go语言中字符串也可能根据需要占用 1 至 4 个字节，这与其它编程语言不同
Go语言这样做不仅减少了内存和硬盘空间占用，同时也不用像其它语言那样需要对使用 UTF-8 字符集的文本进行编码和解码
字符串是一种值类型，且值不可变，即创建某个文本后将无法再次修改这个文本的内容
当字符为 ASCII 码表上的字符时则占用 1 个字节

字符串中可以使用转义字符来实现换行、缩进等效果，常用的转义字符包括:
1. `\n：`换行符
2. `\r：`回车符
3. `\t：`tab 键
4. `\u 或 \U：`Unicode 字符
5. \\：反斜杠自身

*/

func main() {
	var str = "你好\nGO大法"
	//var str = `你好\tGO大法` // 反引号一般用在 需要将内容进行原样输出的时候 使用
	fmt.Printf(str)

	//中文三字节，字母一个字节
	var myStr01 string = "hello,go教程"
	fmt.Printf("myStr01: %d\n", len(myStr01))

	str3 := "hello"
	str4 := "你好"
	fmt.Println(len(str3))
	fmt.Println(len(str4))
	fmt.Println(utf8.RuneCountInString(str4)) // 计算中文长度 2

	// 字符串拼接
	// 个字符串 s1 和 s2 可以通过 s := s1 + s2 拼接在一起。将 s2 追加到 s1 尾部并生成一个新的字符串 s
	// 因为编译器会在行尾自动补全分号，所以拼接字符串用的加号“+”必须放在第一行末尾
	str5 := "第一部分 " + "第二部分"
	s := "hel" + "lo"
	s += "world!"
	fmt.Println(str5)
	fmt.Println(s)
	// 性能不高

	// 第二种拼接方式 性能高
	s1 := "hello,"
	s2 := "world"
	var stringBuild bytes.Buffer
	stringBuild.WriteString(s1)
	stringBuild.WriteString(s2)
	fmt.Println(stringBuild.String())

	/*
	   如果从字符串 `hello go教程` 中获取 `教` 该如何获取呢
	   直接索引对rune类型无效，可以使用string方法转换
	*/
	s3 := "hello"             // 字节
	fmt.Printf("%c\n", s3[0]) // %c 将整数值格式化为对应的 Unicode 字符
	//s4 := "hello,go教程" // 此时类型是 rune 类型

}

/*
字符串的内容（纯字节）可以通过标准索引法来获取，在方括号`[ ]`内写入索引，索引从 0 开始计数：

字符串 str 的第 1 个字节：str[0]
第 i 个字节：str[i - 1]
最后 1 个字节：str[len(str)-1]

需要注意的是，这种转换方案只对纯 ASCII 码的字符串有效
注意：获取字符串中某个字节的地址属于非法行为，例如 &str[i]。
ASCII字符使用`len()`函数
*/
