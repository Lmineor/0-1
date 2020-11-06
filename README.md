# Go语言

## 变量

在go中，变量被显式声明，并被编译器所用来检查函数调用时的类型正确性

var 声明1个或多个变量

声明变量且没有给出对应的初始值时，变量将会初始化微零值。例如，一个int的零值是0。

:= 语句是申明并初始化变量的简写，

## 常量

go支持字符、字符串、布尔和数值  **常量**

const 用于声明一个常量。

const 语句可以出现在任何var语句可以出现的地方
常数表达式可以执行任意精度的运算

数值型常量是没有确定的类型，直到它们被给定了一个类型，比如说一次显示的类型转化

当上下文需要时，一个数可以被给定一个类型，比如变量赋值或者函数调用。
例如

```go
const n = 5000000
fmt.Println(math.Sin(n))
```

这里的math.Sin函数需要一个float64的参数

## slice(切片)

slice是go中一个关键的数据类型，是一个比数组更加强大的序列接口

不像数组，slice的类型仅由它所包含的元素决定（不像数组中还需要元素的个数）。
要创建一个长度非零的空slice，需要使用内建的方法make。

## map(映射、字典、哈希)

要创建一个空map，需要使用内建的make: `make(map[key-type]val-type)`
使用典型的make[key]=val语法来设置键值对。

光声明map类型，但是没有初始化，a就是初始值nil

### 声明

```go
var a map[string]int
fmt.Println(a == nil) // true
```

```go
// 光声明map类型，但是没有初始化，a就是初始值nil
var a map[string]int
fmt.Println(a == nil)

// map的初始化
a = make(map[string]int, 8)
fmt.Println(a == nil)

// map中添加键值对
a["hh"] = 100
fmt.Printf("a:%v\n", a)  // a:map[hh:100]
fmt.Printf("a:%#v\n", a) // a:map[string]int{"hh":100} 可以打印出"
fmt.Printf("type：%T\n", a)

//申明map的同时完成初始化
b := map[int]bool{
	1: true,
	2: false,
}
fmt.Printf("b:%#v\n", b)

var c map[int]int
c[1200] = 100 // c这个map没有初始化不能直接操作，即没有初始化相当于在内存中没有位置
fmt.Println(c)
```

### 其他例子
#### 元素类型为map的切片
```go
// 元素类型为map的切片
var mapSlice = make([]map[string]int, 8, 8) //只是完成了切片的初始化
// 还需要完成内部map元素的初始化
mapSlice[0] = make(map[string]int, 8) // 完成了map的初始化

mapSlice[0]["hehhe"] = 11
fmt.Println(mapSlice)
```

#### 值为切片的map

```go
// 值为切片的map
var sliceMap = make(map[string][]int, 8)

_, ok := sliceMap["中国"]
if ok {
	fmt.Println(sliceMap["中国"])
} else {
	sliceMap["中国"] = make([]int, 8, 8) // 完成了对切片的初始化
	sliceMap["中国"][0] = 100
	sliceMap["中国"][1] = 11
	sliceMap["中国"][2] = 1222
}

//遍历sliceMap
for k, v := range sliceMap {
	fmt.Println(k, v)
}
```

#### 遍历
```go
// 判断某个键是否存在
var scoreMap = make(map[string]int, 8)
scoreMap["小明"] = 100
scoreMap["大名"] = 200

// 判断 小红 是否在scoreMap中
v, ok := scoreMap["小红"]
if ok {
	fmt.Println("小红在scoreMap中", v)
} else {
	fmt.Println("小红不在scoreMap中")
}

//遍历map
for k, v := range scoreMap {
	fmt.Println(k, v)
}

//只要key
for k := range scoreMap {
	fmt.Println(k)
}

//只要value
for _, v := range scoreMap {
	fmt.Println(v)
}
```

使用name[key]来获取一个键的值
支持len()

#### 移除键值对

移除键值对 `delete(m, "k2")`

```go
// 删除小明的信息
delete(scoreMap, "小明")
fmt.Println(scoreMap)
```

#### 按照某个固定顺序遍历map

```go
// 按照某个固定顺序遍历map
var scoreMap = make(map[string]int, 100)

//添加50个键值对
for i := 0; i < 50; i++ {
	key := fmt.Sprintf("stu%02d", i)
	val := rand.Intn(100) // 0-99的随机整数
	scoreMap[key] = val
}

for k, v := range scoreMap {
	fmt.Println(k, v)
}

//按照key从小到大的顺序去遍历scoreMap
// 1. 先取出所有的key存放到切片中
keys := make([]string, 0, 100)
for k := range scoreMap {
	keys = append(keys, k)
}
// 2. 对key排序
sort.Strings(keys) //目前keys是一个有序的切片

// 3. 按照排序后的key对scoreMap排序
for _, key := range keys {
	fmt.Println(key, scoreMap[key])
}
```

#### 统计一个字符串中每个单词出现的次数

```go
// 统计一个字符串中每个单词出现的次数
// "how do you do"中每个单词出现的次数
// 0. 定义一个map[string]int
var s = "how do you do"
var wordCout = make(map[string]int, 10)

// 1. 字符串中都有哪些单词
words := strings.Split(s, " ")

// 2. 遍历单词做统计
for _, word := range words {
	v, ok := wordCout[word]
	fmt.Println(v, ok)
	if ok {
		// 说明map中有这个单词的统计记录
		wordCout[word] = v + 1
	} else {
		wordCout[word] = 1
	}
}
for k, v := range wordCout {
	fmt.Println(k, v)
}
```

## range遍历

```go
sum := 0
for _, num := range nums {
    sum += num
}

```

range 在数组和slice中都同样提供每个项的索引和值。上面我们不需要索引，所以我们使用 **空值定义符_**来忽略

range在字符串中迭代unicode编码。第一个返回值是起始字节的位置，第二个是字节对应的unicode编码

## 函数

```go
package main

import "fmt"

func main() {
	// fmt.Println(intSum(1, 2))
	fmt.Println(intSum2(1, 2))
}

func intSum(a int, b int) int {
	ret := a + b
	return ret
}

// 在函数体的返回值已经声明了ret 因此，在函数里不需要再次对ret进行
// 声明，直接返回即可
func intSum2(a int, b int) (ret int) {
	ret = a + b
	return // 可以写成return ret
}

```

### 变参函数

可变参数在函数体中是**切片类型**

```go
func useFuncWithChangeParam(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

useFuncWithChangeParam(1, 2)
useFuncWithChangeParam(1, 2, 5, 6, 7)

nums := []int{1, 2, 3, 4, 5}
useFuncWithChageParam(nums...)
```

固定参数和可变参数同时出现时，可变参数要放在最后

```go
func useFuncWithChangeParam(a int, b ...int) {
}

```

> Go语言中，没有默认参数

### Go语言中函数参数的简写

a, b都为int,可以简写为`(a, b int)`

```go
func useFuncWithChangeParam(a , b int) {

}

```

### 多返回值

```go
func useFuncWithMultiRet() (int, int) {
	return 3, 7
}

a, b := useFuncWithMultiRet()
fmt.Printf("a: %d, b: %d\n", a, b)

_, c := useFuncWithMultiRet()
fmt.Println(c)
```
指定函数返回值，并简写的方式：

```go
func useFuncWithMultiRet() (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

a, b := useFuncWithMultiRet()
fmt.Printf("a: %d, b: %d\n", a, b)
```

### defer 语句

Go语言中的`defer`语句会将其后要跟随的语句进行延迟处理。
在`defer`归属的函数即将返回时，将延迟处理的语句按`defer`定义的逆序进行执行。
也就是说，先被`defer`的语句最后被执行，最后被`defer`的语句，最先执行。

例子:
```go
package main

import "fmt"

// defer:延迟执行
func main() {
	fmt.Println("Staring...")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
	fmt.Println("end...")
}

// res
Staring...
end...
4
3
2
1
```

由于`defer`语句延迟调用的特性，所以defer语句能非常方便的处理资源释放的问题。比如：资源清理、文件关闭、解锁及记录时间。

### 变量作用域

#### 局部变量

- 函数内定义

如果局部变量和全局变量重名。优先访问局部变量

- 语句块内定义

### 函数作为变量 传参

```go
package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func main() {
	r1 := calc(100, 20, add)
	fmt.Println(r1)
	r2 := calc(100, 20, sub)
	fmt.Println(r2)
}
```

### 匿名函数

```go
func main() {
	sayHello := func() {
		fmt.Println("匿名函数")
	}
	sayHello()
}

// 与下面的等价
func main() {
	func() {
		fmt.Println("匿名函数")
	}()
}

```

### 闭包

闭包 = 函数 + 外层变量的引用

一个简单示例

```go
package main

import "fmt"

// 定义一个函数它的返回值是一个函数
func a() func() {
	name := "小明"
	return func() {
		fmt.Println("hello", name)
	}

}

func main() {
	r := a()
	r() // 相当于执行了a函数内部的匿名函数
}
```

进阶

```go 
func a(name string) func() {
	return func() {
		fmt.Println("hello", name)
	}

}

func main() {
	r := a("小")
	r() // 相当于执行了a函数内部的匿名函数
}
```

再进阶

```go
func makeSuffixFunc(suffix string) func(name string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
func main() {
	suffix := makeSuffixFunc(".txt")
	fmt.Println(suffix("rar.txt"))
	fmt.Println(suffix("1.jpg"))
	fmt.Println(suffix("book.txt"))
	fmt.Println(suffix("book.doc"))
}

//
rar.txt
1.jpg.txt
book.txt
book.doc.txt

```

更进阶
```go
func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}
	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {
	calcor_add, calcor_sub := calc(100)
	fmt.Println(calcor_add(10))  // base = 100 + 10
	fmt.Println(calcor_sub(110)) // base = 110 - 110
}
```

另一个demo

```go
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
nextInt := intSeq()

fmt.Println(nextInt())
fmt.Println(nextInt())
fmt.Println(nextInt())
fmt.Println(nextInt())

newInts := intSeq()
fmt.Println(newInts())

```

这个intSeq函数返回另一个在intSeq函数体内定义的匿名函数。这个返回的函数使用闭包的方式，隐藏变量i

调用intSeq函数，将返回值（也是一个函数）赋给nextInt。这个函数的值包含了自己的值i，这样在每次调用nextInt是都会更新i的值。

## 内置函数介绍

|  内置函数   | 介绍  |
|  :-:  | :-:  |
| close  | 用来关闭channel |
| len  | 用来求长度，比如string、array、slice、map、channel |
| new  |  用来分配内存，主要用来分配引用类型，比如chan、map、slice|
| append | 用来追加元素到数组、slice中|
|panic 和recover| 用来做错误处理| 


### panic_recover

panic可以在任何位置
recover 需要搭配defer使用

```go
func a() {
	fmt.Println("func a")
}

func b() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			fmt.Println("func b error")
		}
	}()
	panic("panic in b")
}

func c() {
	fmt.Println("func c")
}

func main() {
	a()
	b()
	c()
}

//
func a
panic in b
func b error
func c

```

> 注意
1. `recover()`必须搭配`defer`使用
2. `defer`一定要在可能引发`panic`的语句之前定义

### new函数

new 函数是一个内置函数，函数签名如下
```go
func new(Type) *Type
```
new不太常用，使用new函数得到的是一个类型的指针，并且该指针对应的值为该类型的零值。

### make函数

函数签名
```go
func make(t Type, size ...IntegerType) Type
```

```go
var a map[string]int
a = make(map[string]int, 10)
a["hh"] = 10
fmt.Println(a)
```

### make与new的区别
1. 二者都是用来做内存分配的
2. make只用于slice，map以及channel的初始化。返回的还是这三个引用类型本身
3. new用于类型的内存分配。并且内存对应的值为类型零值，返回的是指向类型的指针。

## 指针

与c++相同`*`， `&`
不能进行偏移和运算，是安全指针。

```go
ptr := &v	//v的类型为T
```
其中：
- `v`代表被取地址的变量，类型为`T`
- `ptr`用于接收地址的变量。`ptr`的类型就是`*T`, 称作`T`的指针类型，`*`代表指针

```go
a := 10
ptr := &a

fmt.Printf("%v %p\n", a, &a)
fmt.Printf("%v %p\n", ptr, ptr)
```

## 类型别名和自定义类型



## 结构体

```go
type person struct {
	name string
	age  int
}

```

## 方法

Go语言中的`方法（Method）`是一种作用于特定类型变量的函数，这种特定类型变量叫做`接收者（Receiver）`。类似于其他语言中的`this`，`self`

方法格式定义如下：
```go
func (接收者变量 接收者类型) 方法名(参数列表)(返回参数){
	函数体
}
```

其中：
- 接收者变量：接收者中的参数变量在命名时，官方建议使用接收者类型名的第一个小写字母，而不是self，this之类的命名。例如，Person类型的接收者变量应该命名为p，Connector类型的接收者变量应该命名为c。
- 接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型
- 方法名、参数列表、返回参数：具体定义与函数格式相同。

```go
// 方法的定义示例

type Person struct{
	name string
	age  int8
}


// NewPerson 是一个Person类型的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age: age,
	}
}

// Dream 为Person类型定义方法
func (p Person)Dream(){
	fmt.Printf("%s的梦想是学好Go语言\n", p.name)
}


// SetAge 指针接收者，表示接收者的类型是一个指针
func (p *Person)SetAge(newAge int8){
	p.age = newAge
}

// SetAge2 值接收者，表示接收者的类型是一个值
func (p Person)SetAge2(newAge int8){
	p.age = newAge
}

func main(){
	p1 := NewPerson("小明", 18)
	p1.Dream()

	p1.SetAge(19)
	fmt.Printf("%#v\n", p1)

	p1.SetAge2(20)
	fmt.Printf("%#v\n", p1)
}
```

可以为值类型或者指针类型的接收器定义方法。

### 什么时候应该使用指针类型：

1. 需要修改接收者中的值
2. 接收者是拷贝代价较大的对象
3. 保证一致性，如果有某个方法使用了指针接收者，那么其他对象也应该使用指针接收者。

### 任意类型添加方法

在Go语言中，接收者的类型可以是任何类型，不仅仅是结构体，任何类型都可以拥有方法。举例：
我们基于内置的int类型使用type关键字可以定义新的自定义类型，然后为我们的自定义类型添加方法。

> 非本地类型不能定义方法，也就是说我们不能给别的包的类型定义方法

```go
// 为任意类型添加方法

// 基于内置的基本类型造一个我们自己的类型
type myInt int

func (m myInt)sayHi(){
	fmt.Println("Hi")
}

func main(){
	var m1 myInt
	m1 = 100
	m1.sayHi()
}
```

## 结构体嵌套

### 结构体匿名字段
```go
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
```

### 嵌套结构体

```go
// 结构体嵌套
type Address struct{
	Provice 	string
	City    	string
	UpdateTime 	string
}

type Email struct{
	Addr 		string
	UpdateTime 	string
}
type Person struct{
	Name     string
	Gender   string
	Age      int8
	// Address  Address // 嵌套另外一个结构体
	Address // 利用匿名结构体简化
	Email
}

func main(){
	p1 := Person{
		Name:    "小王子",
		Gender:  "男",
		Age:     18,
		Address: Address{
			Provice: "山东",
			City: "威海",
			UpdateTime:"2020-9",
		},
		Email: Email{
			Addr: "xiaoming@123.com",
			UpdateTime:"2020-10",
		},
	}
	fmt.Printf("%#v\n",p1)
	fmt.Println(p1.Name, p1.Age, p1.Gender)
	fmt.Println(p1.Address.Provice, p1.Address.City)
	fmt.Println(p1.Provice, p1.City) // 直接访问匿名结构体中的字段
	// fmt.Println(p1.UpdateTime) // Error
	fmt.Println(p1.Email.UpdateTime) 
}
```

### 实现继承

```go
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
```

### 结构体字段的可见性与Json序列化

大写开头可以公开访问，小写开头的表示私有


## 接口

## 切片（slice）

一般用len来判断切片是否为空
原因：
```go
package main

import "fmt"

func main() {
	// nil
	var e []int         // 声明int类型的切片
	var f = []int{}     // 声明并初始化
	g := make([]int, 0) // 使用make
	if e == nil {
		fmt.Println("e is nil")
	}
	if f == nil {
		fmt.Println("f is nil")
	}
	if g == nil {
		fmt.Println("g is nil")
	}
	fmt.Println(e, len(e), cap(e))
	fmt.Println(f, len(f), cap(f))
	fmt.Println(g, len(g), cap(g))
}
```
