package main

import "fmt"

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
}
