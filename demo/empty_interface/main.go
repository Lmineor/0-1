package main

import "fmt"

//　空接口
// 接口中没有定义任何需要实现的方法时，该接口就是一个空接口
// 任意类型都实现了空接口　－－－＞　空接口变量可以存储任意值．

// 空接口一般不需要提前定义
type xxx interface{}

//　空接口的应用
/*
1. 空接口类型作为函数参数
2. 空接口可以作为map的value
*/

func main() {
	var x interface{} // 定义一个空接口变量ｘ
	x = "hello"
	x = true
	x = 100
	fmt.Println(x)

}
