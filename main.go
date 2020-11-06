package main

import "fmt"

type Animal struct{
	name string
}

func(a Animal)Move(){
	fmt.Printf("Format: %s会动~\n", a.name)
}

type Dog struct{
	Feet int8
	*Animal // 匿名嵌套，而且嵌套的是一个结构体指针
}

func (d *Dog) Wang(){
	fmt.Printf("Format: %s 会汪汪汪~\n",d.name)
}

func main(){
	d1 := &Dog{
		Feet: 4,
		Animal : &Animal{
			name: "乐乐",
		},
	}

	d1.Move()
	d1.Wang()
}