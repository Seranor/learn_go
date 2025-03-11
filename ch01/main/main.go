package main

import (
	u "ch01/user" // 导入的别名
	//. "ch01/user" // 直接引用 不用加 u.
	_ "ch01/init" // 匿名 初始化使用
	"fmt"
) // 包的路径

// go get -u  升级到最新的版本
// go get -u=path
// go get github.com/go-redis/redis/v8@version

// 下面写入到 go.mod 文件中
// replace github.com/seranor/A  => github.com/seranor/project-A v1.0.0
// 我使用的是       github.com/seranor/A
// 然后下载是去仓库  github.com/seranor/project-A
// 命令实现 go mode edit -replace github.com/seranor/A=github.com/seranor/project-A@v1.0.0

func main() {
	c := u.Course{
		Name: "go"}
	fmt.Println(u.GetCourse(c))
}
