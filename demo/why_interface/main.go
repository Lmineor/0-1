package main

import "fmt"

type dog struct{}

type cat struct{}

type person struct {
	name string
}

func (d dog) say() {
	fmt.Println("汪汪汪~")
}

func (d cat) say() {
	fmt.Println("喵喵喵~")
}

func (p person) say() {
	fmt.Println("啊啊啊~")
}

// 接口不管是什么类型,只负责要实现什么方法
// 定义一个类型,一个抽象的类型,只要实现了say()这个方法的类型都可以称为sayer类型
type sayer interface {
	say()
}

// 接口不管是什么类型,只负责要实现什么方法
func da(arg sayer) {
	arg.say() // 不管传什么参数都调用say方法

}

func main() {
	c1 := cat{}
	da(c1)
	d1 := dog{}
	da(d1)
	p1 := person{
		name: "小明",
	}
	da(p1)
}
