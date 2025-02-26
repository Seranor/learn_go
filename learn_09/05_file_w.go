package main

import (
	"bufio"
	"fmt"
	"os"
)

func write01(f *os.File) {
	// 写字节操作
	f.Write([]byte("满江红\n"))

	// 写字符串
	f.WriteString("怒发冲冠，凭栏处、潇潇雨歇。\n")
}

func write02(f *os.File) {
	writer := bufio.NewWriter(f)
	writer.WriteString("大浪淘沙\n")
	writer.Flush()
}
func write03() {
	s := "hello world"
	err := os.WriteFile("满江红2", []byte(s), 0666)
	//err := ioutil.WriteFile("满江红2", []byte(s), 0666)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	/*
		os.O_RDONLY 只读的方式打开
		os.O_WRONLY 只写的方式打开
		os.O_RDWR   读写的方式打开
		os.O_APPEND 追加的方式打开
		os.O_CREAT  创建并打开一个新文件
		os.O_TRUNC	打开一个文件并截断它的长度为零 必须有写权限
		os.O_EXCL   如果指定文件存在 返回错误
	*/
	file, err := os.OpenFile("满江红", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("close file err:", err)
		}
	}()
	//write01(file)
	//write02(file)
	write03()
}
