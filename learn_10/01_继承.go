package main

import "fmt"

type Animal struct {
	name string
}

func (a *Animal) eat() {
	fmt.Printf("%s 在 eat\n", a.name)
}
func (a *Animal) sleep() {
	fmt.Printf("%s 在 sleep\n", a.name)
}

type Dog struct {
	kind    string
	*Animal // Dog 继承了 Animal
}

func (d Dog) bark() {
	fmt.Printf("%s 吠\n", d.name)
}

type Cat struct {
	kind    string
	*Animal // Cat 继承了 Animal
}

func main() {
	d1 := &Dog{
		kind: "hashiqi",
		Animal: &Animal{
			name: "huahua",
		},
	}
	fmt.Println(d1.kind)
	fmt.Println(d1.name)
	d1.eat()
	d1.bark()

	c1 := &Cat{
		kind: "jumao",
		Animal: &Animal{
			name: "jindou",
		},
	}
	fmt.Println(c1.kind)
	fmt.Println(c1.name)
	c1.sleep()
	c1.eat()
}
