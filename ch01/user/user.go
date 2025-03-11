package user

// package 是代码复用的基础
// 每个源码文件开始都要声明所属的package
// 和目录名称不一定相同，但是每个文件的 package name 需要相同
// 这个目录下的变量函数等可以相互访问

type Course struct {
	Name string // 可见性 首字母必须大写
}
