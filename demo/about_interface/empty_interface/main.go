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
	// x = "hello"
	// x = true
	x = 100
	// fmt.Println(x)

	// var m = make(map[string]interface{}, 16)
	// m["name"] = "小明"
	// m["age"] = 19
	// m["hobby"] = []string{"篮球", "足球", "双色球"}
	// fmt.Println(m)

	ret, ok := x.(bool) //类型断言
	if !ok {
		fmt.Println("不是布尔值类型")
	} else {
		fmt.Println("是", ret)
	}

	// 另一种类型断言的方法
	switch v := x.(type) {
	case string:
		fmt.Println("string")
		fmt.Printf("%v\n", v)
	case bool:
		fmt.Println("bool")
		fmt.Printf("%v\n", v)
	case int:
		fmt.Println("int")
		fmt.Printf("%v\n", v)
	default:
		fmt.Println("can not guess")
	}
}
