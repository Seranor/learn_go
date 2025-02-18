package main

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday //6
)

func main() {
	/*
				const name [type] = value

				type 可以省略
				也可以批量声明多个常量

				const (
					e = 2.71
					pi = 3.14
				)

			常量间的所有算术运算、逻辑运算和比较运算的结果也是常量
			对常量的类型转换操作或以下函数调用都是返回常量结果：len、cap、real、imag、complex 和 unsafe.Sizeof

		常量声明可以使用 iota 常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不用每行都写一遍初始化表达式
		在一个 const 声明语句中，在第一个声明的常量所在的行，iota 将会被置为 0，然后在每一个有常量声明的行加1
	*/

}
