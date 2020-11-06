package main

import "fmt"

// 结构体匿名字段

// 非匿名
type Person1 struct{
	name string
	age  int8
}

// 匿名
type Person struct{
	string
	int8
}

func main(){
	p1 := Person{
		"小王子",
		18,
	}
	fmt.Println(p1.string, p1.int8)
}