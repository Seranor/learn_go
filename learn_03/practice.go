package main

import (
	"bytes"
	"fmt"
	"strings"
)

/*
字符串替换, 比如将 "Hello, ms的go教程Java教程" 替换为  "Hello, ms的go教程Go教程"
*/
func main() {
	str := "Hello, ms的go教程Java教程"
	source := "Java"
	target := "Go"

	javaIndex := strings.Index(str, source)
	lenSource := len(source)
	start := str[:javaIndex]
	end := str[javaIndex+lenSource:]

	var stringBuild bytes.Buffer
	stringBuild.WriteString(start)
	stringBuild.WriteString(target)
	stringBuild.WriteString(end)
	fmt.Println(stringBuild.String())

	//newStr := start + target + end
	//fmt.Println(newStr)

}
