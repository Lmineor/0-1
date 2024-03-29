# Go语言

## 变量

在go中，变量被显式声明，并被编译器所用来检查函数调用时的类型正确性

var 声明1个或多个变量

声明变量且没有给出对应的初始值时，变量将会初始化微零值。例如，一个int的零值是0。

`:=` 语句是申明并初始化变量的简写，

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

## 切片与数组的区别

定义：

数组：类型 [n]T 表示拥有 n 个 T 类型的值的数组。
切片：类型 []T 表示一个元素类型为 T 的切片。

值得注意的是： string 类型 的底层是数组， 是数组就可以用来切片操作

1、初始化：
数组需要指定大小，不指定也会根据初始化的自动推算出大小，不可改变 
切片不需要指定大小

2、函数传递：
数组需要明确指定大小，切片不需要。数组是值传递，切片是地址传递。

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

## 包

### 定义
```go
package 包名
```

注意:
- 一个文件夹下只能有一个包,同样的包不能在多个文件夹下
- 包名可以不和文件夹名一样,包名不能包含中横线`-`
- 包名为main的包为应用程序入口,编译时不含main的包不会生成可执行文件

### 导包

当写的代码在$GOPATH目录下的话,我们导入的包的路径要从$GOPATH/src后面开始写起
使用正斜线 `/`

> 注意此时标识符的可见性

不允许导入包而不使用
不允许循环导入包

#### 给包起别名

```go
import test_package "this/is/test/package1" // test_package 为包的别名
```

#### 匿名导入包

```go
import _ "this/is/test/package1" // test_package 为包的别名
```

## init()初始化函数

在Go语言程序执行导入时会自动触发包内部`init()`函数.需要注意的是:`init()`函数没有参数也没返回值.`init()`函数在程序运行是自动被调用执行,不能在代码中主动调用它.
比`main()`函数优先执行

执行时机:
```go
全局声明-> init()->main()
```

## 接口

`interface()` 就是一堆方法的集合

### 接口定义

```go
type 接口类型名 interface{
	方法名1(参数列表1) 返回值列表1
	方法名2(参数列表2) 返回值列表2
}
```

- 接口名:使用type将接口定义为自定义的类型名.Go语言的接口在命名时,一般会在单词后加er,如,有写操作的接口叫Writer
- 方法名:当方法名首字母是大写且这个接口类型名首字母也是大写时,这个方法可以被接口所在的包之外的代码访问
- 参数列表,返回值列表,:参数列表和返回值列表中的参数变量名可以省略



### 接口实现的条件

实现了接口的方法即可

### 使用值接受者实现接口和使用指针实现接口的区别

使用值接受者:类型的值和类型的指针都能够保存到接口变量中.
使用指针接收者:只有类型指针能够保存到接口变量中

举个例子

使用值接收者:

```go
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
```
使用指针接收者:

```go
// 使用指针接收者:只有类型指针能够保存到接口变量中
func (p *person) move() {
	fmt.Printf("%s在跑\n", p.name)
}

func main() {
	var m mover
	p1 := person{ // person类型的值
		name: "小王子",
		age:  18,
	}
	p2 := &person{ // person类型的指针
		name: "小明",
		age:  18,
	}
	m = p1 // 无法赋值,因为p1是person类型的值,没有实现mover接口
	m = p2
	m.move()
}
```

### 接口的嵌套

```go
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
```

###　空接口的定义

空接口是指没有定义任何方法的接口．因此任何类型都实现了空接口
空接口
接口中没有定义任何需要实现的方法时，该接口就是一个空接口
任意类型都实现了空接口　－－－＞　空接口变量可以存储任意值．
空接口一般不需要提前定义

举个例子：

```go

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

	var m = make(map[string]interface{}, 16)
	m["name"] = "小明"
	m["age"] = 19
	m["hobby"] = []string{"篮球", "足球", "双色球"}
	fmt.Println(m)
}
```

##### 类型断言

```go
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
```

## Go语言反射

### 变量的内在机制

Go 语言中的变量分为两部分
- 类型信息：预先定义好的元信息
- 值信息：程序运行过程中可动态变化的。

### 反射介绍

反射是指程序运行期对程序本身进行访问和修改的能力。程序在编译时，变量被转换为内存地址，变量名不会被编译器写入到可执行部分。在运行程序时，程序无法获取自身的信息。
支持反射的语言可以在程序编译时将变量的反射信息，如字段名称、类型信息、结构体信息等整合到可执行文件中，并给程序提供接口访问反射信息，这样就可以在程序运行期间获取类型的反射信息，并且有能力修改他们。
Go程序在运行期间使用reflect包访问程序的反射信息。

之前讲空接口。空接口可以存储任意类型的变量，那我们如何知道这个空接口保存的数据是什么呢？反射就是在运行时动态获取一个变量的类型信息和值信息。

### reflect包

pass

##　并发

```go

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func helloWorld() {
	fmt.Println("hello world")
	wg.Done() // 计数器-1
}

func main() { // 开启一个主goroutine去执行main函数

	for i := 0; i < 10000; i++ {
		wg.Add(1)       // 计数牌+1
		go helloWorld() // 开启一个goroutine去执行helloWorld函数
	}

	fmt.Println("Hello main")
	wg.Wait() // 阻塞等所有线程都结束
}

```

### goroutine 调度

`GMP`是Go语言运行时调度的实现.是go语言自己实现的一套调度系统.区别于操作系统调度OS线程.

- `G` 就是goroutine，里面除了存放本goroutine信息外，还有与所在P的绑定等信息。
- `M（machine）`是Go运行时对操作系统内核线程的虚拟。M与内核线程一般是一一映射的关系，一个goroutine最终是要放到M上执行的。
- `P`管理着一组goroutine队列，P里面会存储当前goroutine运行的上下文环境（函数指针，堆栈地址及地址边界）,P会对自己管理的goroutine队列做一些调度（比如把占用CPU时间较长的goroutine暂停、运行后续的goroutine等等），当自己的队列消费完了就去全局队列里取，如果全局队列里也消费完了会去其他P队列里抢任务。

P与M一般也是一一对应的。他们的关系是：P管理着一组G挂载在M上运行。当一个G长久阻塞在一个M上时，runtime会新建一个M，阻塞G所在的P会把其他的G挂载在新建的M上，当旧的G阻塞完成或者认为其已经死掉时，回收旧的M。

P的个数是通过`runtime.GOMAXPROCS`设定（最大256），Go1.5版本之后默认为物理线程数。在并发量大的时候会增加一些P和M，但不会太多，切换太频繁的话得不偿失。

单从线程调度讲，Go语言比起其他语言的优势在于OS线程是由OS内核来调度的，`goroutine`则是由Go运行时（runtime）自己的调度器调度的，这个调度器使用一个称为`m:n`调度的技术（复用/调度m个goroutine到n个OS线程）。其一大特点是goroutine的调度是在用户态下完成的，不涉及内核态和用户态之间的频繁切换，包括内存的分配与释放，都是在用户态维护着一块大的内存池，不直接调用系统的malloc函数（除非内存池需要改变），成本比调度OS线程低很多。另一方面充分利用了多核的硬件资源。近似的把若干goroutine均分在物理线程上，再加上本身goroutine超轻量，以上种种保证了go调度方面的性能。


```go

runtime.GOMAXPROCS(1) // 设置当前程序并发时占用的CPU逻辑核心数
```

Go语言中操作系统线程和goroutine的关系：
1. 一个操作系统线程对应用户态多个goroutine
2. go程序可以同时使用多个操作系统线程。
3. goroutine和OS线程是多对多的关系，即`m:n`

## channel

Go的并发模型是CSP（Communicating Sequential Process），提倡通过通信共享内存而不是通过共享内存而实现通信。

Go语言中的通道（channel）是一种特殊的类型。通道像一个传送带或者队列，总是遵循`FIFO`规则，保证收发数据的顺序。每一个通道都是具体类型的导管，也就是声明channel的时候需要为其指定元素类型。

```go
var 变量 chan 元素类型
```

例子

```go
var ch1 chan int	// 声明一个传递整型的通道
var ch2 chan bool	// 声明一个传递布尔的通道
var ch3 chan []int	// 声明一个传递int切片的通道
```

### 创建channel

通道是**引用类型**，通道类型的空值是`nil`

```go
var ch chan int
fmt.Println(ch) // nil
```

通道声明后需要使用make函数初始化之后才能使用

```go
make(chan 元素类型, [缓冲大小])
```

channel的缓冲大小是可选的

```go
ch4 := make(chan int)
ch5 := make(chan bool)
ch6 := make(chan []int)
```

### channel 的发送接收关闭

有发送(send)，接收(receive)，关闭(close)三种操作
发送、接收都使用<-符号。

```go
var ch1 chan int
ch1 = make(chan int, 1)

ch2 := make(chan int, 1) // 与上面两步效果一样

ch1 <- 10
x := <-ch1
fmt.Println(x)
close(ch1)
```

#### 发送&接收

```go
ch <- = 10 // 把10发送到通道中
x := <- ch // 从通道中取值
```

### 有无缓冲区

```go
ch2 := make(chan int)   // 无缓冲区通道 又称为同步通道
ch1 = make(chan int, 1) // 带缓冲区通道
```

### 一个demo
```go
// 指针
package main

import "fmt"

/*
两个goroutine
	1. 生成0-100的数字发送发到ch1
	2. 从ch1中取出数据计算他的平方，把结果发送到ch2中
*/
func producer(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch1, ch2 chan int) {
	for tmp := range ch1 {
		ch2 <- tmp * tmp
	}
	close(ch2)
}
func main() {
	ch1 := make(chan int, 100)
	ch2 := make(chan int, 200)

	go producer(ch1)
	go consumer(ch1, ch2)

	for tmp := range ch2 {
		fmt.Println(tmp)
	}

}
```

### 限定只读只写
```go
func producer(ch chan<- int) { // 只写
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch1 <-chan int, ch2 chan int) { // ch1 只读，ch2 只写
	for tmp := range ch1 {
		ch2 <- tmp * tmp
	}
	close(ch2)
}
```

### 通道总结

channel常见的异常总结

| channel| nil |  非空  |空的  | 满了| 没满 |
|   :-  |:-  |:-     |:-  |:- |:-   |
|接收|阻塞|接收值|阻塞   |	接收值|接收值 |接收值
|发送|阻塞|发送值|发送值|	发送值|阻塞|发送值|
|关闭|panic|关闭成功，读完数据后返回零值|关闭成成，返回零值|关闭成成，读完数据后返回零值|关闭成功读完数据后返回零值|

关闭已经关闭的`channel`也会引发`panic`

## worker pool (goroutine池)

控制goroutine数量，防止泄露或暴涨

```go
package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("worker: %d, start job: %d\n", id, job)
		results <- job * 2
		time.Sleep(time.Millisecond * 500)
		fmt.Printf("worker: %d, stop job: %d\n", id, job)
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 开启3个goroutine
	for j := 0; j < 3; j++ {
		go worker(j, jobs, results)
	}
	// 发送5个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for i := 0; i < 5; i++ {
		ret := <-results
		fmt.Println(ret)
	}
}
```

### select多路复用

在某些场景下，我们需要同时从多个通道接收数据。
通道在接收数据时，如果没有数据可以接收将会发生阻塞。

go内置了`select`关键字，类似于`switch`语句

demo
```go
package main

import "fmt"

// select
func main() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		default:
			fmt.Println("nothing")
		}
	}
}
// res
0
2
4
6
8
```

使用select语句能提高代码的可读性

- 可处理一个或多个channel的发送/接收操作
- 如果多个case同时满足，select会随机选择一个
- 对于没有case的select{}会一直等待。可用于阻塞main函数。

## 并发安全和锁

### 互斥锁

```go
package main

import (
	"fmt"
	"sync"
)

// 并发安全和锁
// 多个goroutine并发操作全局变量x
var (
	x    int64
	wg   sync.WaitGroup
	lock sync.Mutex // 互斥锁
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x++
		lock.Unlock() // 解锁
	}
	wg.Done()
}
func main() {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x) // 10000
}
```

### 读写互斥锁

读写锁分为两种：
读锁和写锁。
- 当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁。如果要是获取写锁就会等待；
- 当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会继续等待

示例
读写都加锁

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	// lock.Lock()
	rwlock.RLock()
	time.Sleep(time.Millisecond)
	// lock.Unlock()
	rwlock.RUnlock()
	wg.Done()
}

func write() {
	// lock.Lock()
	rwlock.Lock()
	x = x + 1
	// lock.Unlock()
	rwlock.Unlock()
	time.Sleep(time.Millisecond * 10)
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	for j := 0; j < 10; j++ {
		wg.Add(1)
		go write()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
```

注意：读写锁非常适合读多写少的场景。如果读和写的操作差别不大。读写锁的优势就发挥不出来。

### sync.WaitGroup

在go中可以使用`sync.WaitGroup`来实现并发任务的同步。
`sync.WaitGroup`有以下几个方法

|方法名|功能|
|:-|	:-|
|`(wg *WaitGroup)` `Add(delta int)`|计数器+delta|
|`(wg *WaitGroup)` ` Done()`|计数器-1|
|`(wg *WaitGroup)` `Wait(delta int)`|阻塞直到计数器变为0|

`sync.WaitGroup`内部维护着一个计数器。

### sync.Once

较为进阶
在编程的很多场景下我们需要确保某些操作在高并发的场景下只执行一次。例如只加载一次配置文件、只关闭一次通道等。

Go语言中的`sync.Once`提供了解决方案。

`sync.Once`只有一个do方法，其签名如下：

```go
func (o *Once) do(f func()){}
```

备注：如果要执行的函数f需要传递参数就需要搭配闭包来使用

#### 加载配置文件示例

### sync.Map

Go语言中内置的map不是并发安全的。

go内置了一个并发安全的map， 是空接口类型，不需要制定key和value的类型

```go
package main

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的map
var m = make(map[int]int)
var m2 = sync.Map{} // 加{}相当于初始化了
var wg sync.WaitGroup

// func get(key int) int {
// 	return m[key]
// }

// func set(key int, value int) {
// 	m[key] = value
// }

func main() {
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			m2.Store(i, i+100)
			value, _ := m2.Load(i)
			fmt.Printf("key:%v, value:%v\n", i, value)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
```

### 原子操作

atomic包

## socket编程

### Go语言实现TCP通信

TCP服务端程序的处理流程
1. 监听端口
2. 接收客户端请求建立连接
3. 创建goroutine处理连接

net包
