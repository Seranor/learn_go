package main

import "fmt"

type Player struct {
	name        string
	healthPoint int
	magic       int
}

func NewPlayer(name string, healthPoint int, magic int) *Player {
	return &Player{
		name:        name,
		healthPoint: healthPoint,
		magic:       magic}
}

// attach 归属到 Player 结构体，只有Player结构体的对线才能调用此方法
func (p *Player) attack() {
	fmt.Printf("%s attack 10\n", p.name)
	p.healthPoint -= 10
}

func (p *Player) attacked() {
	fmt.Printf("%s attacked ... \n", p.name)
}
func main() {
	p1 := NewPlayer("wawa", 100, 100)
	p2 := NewPlayer("dada", 100, 100)
	p1.attack() // 会将 p1 结构体对象赋值给p
	p2.attack()
	p1.attacked()
	fmt.Println(p1.healthPoint)
}
