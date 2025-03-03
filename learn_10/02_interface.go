package main

import "fmt"

type Dog2 struct {
	name string
}
type Cat2 struct {
	name string
}

type Pay interface {
	pay() string
}

type AliPay struct {
}

func (a AliPay) pay() string {
	fmt.Println("AliPay 支付")
	return "aliPay"
}

type WeChatPay struct {
}

func (w WeChatPay) pay() string {
	fmt.Println("WeChatPay 支付")
	return "WeChatPay"
}

type YinLianPay struct {
}

func (Y YinLianPay) pay() string {
	fmt.Println("YinLianPay 支付")
	return "YinLianPay"
}

func account(p Pay) {
	// 支付方式的调用
	var ret = p.pay()
	fmt.Println(ret)
}

func main() {
	var x Dog2
	x = Dog2{name: "wangwang"}
	fmt.Println(x.name)

	//x = Cat2{name: "miaomiao"} // 此时不能再把 Cat2 类型赋值给 x 了

	/*
		继承
		重写
		父类可以引用不同的子类对象
	*/

	var p1 Pay
	p1 = AliPay{}
	p1.pay()
	p1 = WeChatPay{}
	p1.pay()
	p1 = YinLianPay{}
	p1.pay()

	p2 := AliPay{}
	p3 := WeChatPay{}
	p4 := YinLianPay{}

	account(p2)
	account(p3)
	account(p4)
}
