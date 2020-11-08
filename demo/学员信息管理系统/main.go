package main

import (
	"fmt"
	"os"
)

/*
学员信息管理系统
1. 添加学员信息
2. 编辑学员信息
3. 展示所有学员信息
*/

func showMenu() {
	fmt.Println("欢迎来到学员信息管理系统")
	fmt.Println("1. 添加学员信息")
	fmt.Println("2. 编辑学员信息")
	fmt.Println("3. 展示所有学员信息")
	fmt.Println("4. 退出")
}

func getInput() *student {
	var (
		id    int
		name  string
		class string
	)
	fmt.Print("请输入学生的学号:")
	fmt.Scanf("%d\n", &id)
	fmt.Print("请输入学生的姓名:")
	fmt.Scanf("%s\n", &name)
	fmt.Print("请输入学生的班级:")
	fmt.Scanf("%s\n", &class)
	newStu := newStudent(id, name, class)
	return newStu
}

func main() {

	sm := newStudentMgr(0, 100)
	// 1. 打印系统菜单
	showMenu()
	// 2. 等待用户选择要执行的选项
	fmt.Println()

	var choice int
	for {
		fmt.Print("\n请输入你的选择:")
		fmt.Scanf("%d\n", &choice) // 格式是数字加回车的格式
		switch choice {
		case 1:
			//添加学员
			stu := getInput()
			sm.addStudent(stu)
		case 2:
			//编辑学员
			stu := getInput()
			sm.editStudent(stu)
		case 3:
			// 展示学员
			sm.showStudent()
		case 4:
			os.Exit(0)
		}
	}
	// 3. 执行用户选择的动作

}
