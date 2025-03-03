package main

import "fmt"

type WashMachine interface {
	wash()
	dry()
}
type Dry struct {
	name string
}

func (d Dry) dry() {
	fmt.Println("dry")
}

type Hair struct {
	name string
	Dry
}

func (h Hair) wash() {
	fmt.Printf("washing!\n")
}

func main() {
	var haier = Hair{
		name: "haier",
		Dry:  Dry{name: "dry y"},
	}

	var w WashMachine
	w = haier
	w.wash()
	w.dry()
}
