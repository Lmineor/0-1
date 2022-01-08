# 结构体

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


```go
package main

import (
	"encoding/json"
	"fmt"
)

// 结构体字段可见性与json序列化

// 大写开头可以公开访问，小写开头的表示私有
type student struct {
	ID   int
	Name string
}

type class struct {
	Title    string
	Students []student
}

func newStudent(id int, name string) student {
	return student{
		ID:   id,
		Name: name,
	}

}

func main() {
	// 创建一个班级变量c1
	c1 := class{
		Title:    "银河骚男",
		Students: make([]student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		tmpStu := newStudent(i, fmt.Sprintf("format: stu%02d", i))
		c1.Students = append(c1.Students, tmpStu)
	}
	fmt.Printf("format: %#v\n", c1)

	// Json序列化：GO语言中的数据->Json格式的字符串
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed , error:", err)
		return
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	// json反序列化:Json格式的字符串->GO语言中的数据
	jsonStr := `{"Title":"银河骚男","Students":[{"ID":0,"Name":"format: stu00"},{"ID":1,"Name":"format: stu01"},{"ID":2,"Name":"format: stu02"},{"ID":9,"Name":"format: stu09"}]}`

	var c2 class
	err = json.Unmarshal([]byte(jsonStr), &c2)
	if err != nil {
		fmt.Println("Json unmarshal failed", err)
	}
	fmt.Printf("%#v\n", c2)
}
```

### 结构体标签Tag

`Tag`是结构体的元信息,可以在运行时通过反射的机制读取出来.
`Tag`在结构体字段的后方定义,由一对反引号包裹起来,具体格式如下:

```go
`key1: "value1" key2: "value2"`
```

```go
type class struct {
	Title string `json:"title"` // json表示用json包来处理的时候使用
	/*
		这时Title 经过序列化之后 titile就可以是小写字母开头了,
	*/
	Students []student `json:"students" db:"student" xml:"ss"`
}
```