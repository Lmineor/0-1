package main

import "fmt"

type student struct {
	id    int // unique
	name  string
	class string
}

//Student类型的构造函数
func newStudent(id int, name, class string) *student {
	return &student{
		id:    id,
		name:  name,
		class: class,
	}
}

type studentMgr struct {
	allStudents []*student
}

// studentMgr 构造函数
func newStudentMgr(initLen, maxCap int) *studentMgr {
	return &studentMgr{
		allStudents: make([]*student, initLen, maxCap),
	}
}

// 添加学生
func (s *studentMgr) addStudent(newStu *student) {
	s.allStudents = append(s.allStudents, newStu)

}

// 编辑学生
func (s *studentMgr) editStudent(newStu *student) bool {
	for i, v := range s.allStudents {
		if v.id == newStu.id {
			s.allStudents[i] = newStu
			return true
		}
	}
	return false
}

// 展示学生
func (s *studentMgr) showStudent() {
	for _, v := range s.allStudents {
		fmt.Printf("学号:%d, 姓名:%s, 班级:%s", v.id, v.name, v.class)
	}
}
