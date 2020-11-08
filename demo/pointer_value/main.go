package main

import "fmt"

// 接口的嵌套
type animal interface {
	mover
	sayer
}

type mover interface {
	move()
}

type sayer interface {
	say()
}

type person struct {
	name string
	age  int8
}

// 使用值接收者:类型的值和类型的指针都能够保存到接口变量中.
// func (p person) move() {
// 	fmt.Printf("%s在跑\n", p.name)
// }

// 使用指针接收者:只有类型指针能够保存到接口变量中
func (p *person) say() {
	fmt.Printf("%s在叫\n", p.name)
}

func (p *person) move() {
	fmt.Printf("%s在跑\n", p.name)
}

func main() {
	var m mover

	p2 := &person{ // person类型的指针
		name: "小明",
		age:  18,
	}
	m = p2
	m.move()

	var s sayer
	s = p2
	s.say()

	var a animal
	a = p2
	a.move()
	a.say()
}
