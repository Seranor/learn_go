package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func read01(f *os.File) {
	// 读字节
	var b = make([]byte, 9)
	n, _ := f.Read(b)
	fmt.Println("读取的字节个数", n)
	fmt.Println("读取的字节内容", b)
	fmt.Println("读取的字节对应的字符内容", string(b))
}
func read02(f *os.File) {
	// 按行读
	reader := bufio.NewReader(f)
	for true {
		s, err := reader.ReadString('\n')
		//s, err := reader.ReadBytes('\n')
		fmt.Print(s)
		if err == io.EOF {
			break
		}
	}
}
func read03() {
	bytes, err := os.ReadFile("自嘲") // ioutil.ReadFile 自 1.16 版本移动到了 os 库里面了
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bytes))
}
func main() {
	// 打开文件，获取文件描述符  文件句柄
	file, err := os.Open("自嘲")
	if err != nil {
		// 报错了
		fmt.Println("err:::", err)
	}
	fmt.Println(file)

	// defer 关闭文件
	defer file.Close()
	// 读文件
	// 1
	//read01(file)
	// 2
	//read02(file)
	// 3 读文件，适用于小文件
	read03()
}
