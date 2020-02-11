package main

import (
	"fmt"
	"os"
)

/*
	控制台使用 go build 命名即可得到一个该系统下可执行程序
	windows平台下执行 go build -o 学生管理系统.exe 可以得到一个名为学生管理系统的windows平台的可执行程序
	也可以利用面向对象的思想将程序改造成结构体版
	新建一个 student_system 结构体 增删改查方法属于这个结构体 结构体包含 var m = make(map[int64]*student, 64)对象
	除main方法外,其他结构体都可以提取到另外的文件中
*/
//student_system 是管理系统对象
type student_system struct {
	data map[int64]*student
}

// student 是一个学生对象
type student struct {
	studentID int64
	name      string
}

// student的构造方法
func newStudent(studentID int64, name string) *student {
	return &student{
		studentID: studentID,
		name:      name,
	}
}

/*
	声明一个管理系统的全局变量
*/
var sys student_system

func main() {
	sys = student_system{
		data: make(map[int64]*student, 64),
	}
	fmt.Println("欢迎来到小呜的学生管理系统~")
	for {
		sys_control()
	}
}

/*
	管理系统控制的方法
*/
func sys_control() {
	show_menu()
	var chose int
	fmt.Print("您选择了:")
	_, err := fmt.Scanln(&chose)
	if err != nil {
		fmt.Println("您的输入有误:", err)
	}
	switch chose {
	case 1:
		sys.showStudents()
	case 2:
		sys.addStudent()
	case 3:
		sys.delStudent()
	case 4:
		sys.updateStudent()
	case 5:
		os.Exit(1)
	default:
		fmt.Println("输入错误,请重新选择")
	}
}

func show_menu() {
	fmt.Println(`
 		请选择操作:
		1.查看学生列表
		2.添加学生
		3.删除学生
		4.修改学生
		5.退出系统
		`)
}

//管理系统的更新方法
func (s *student_system) updateStudent() {
	var studentID int64
	var name string
addID:
	fmt.Print("请输入要修改的学生ID:")
	_, err := fmt.Scanln(&studentID)
	if err != nil {
		fmt.Println("输入有误!", err)
		goto addID
	}
	if sys.data[studentID] == nil {
		fmt.Println("您输入的学号不存在! 按0返回")
		goto addID
	}
	if studentID == 0 {
		return
	}
addName:
	fmt.Print("请输入修改后的学生姓名:")
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println("输入有误!", err)
		goto addName
	}
	sys.data[studentID] = newStudent(studentID, name)
}

//管理系统的删除学生方法
func (s *student_system) delStudent() {
	var studentID int64
	fmt.Print("请输入要删除的学生id:")
	_, err := fmt.Scanln(&studentID)
	if err != nil {
		fmt.Println("输入有误!", err)
	}
	delete(sys.data, studentID)
}

//管理系统的添加学生方法
func (s *student_system) addStudent() {
	var studentID int64
	var name string
addID:
	fmt.Print("请输入要添加的学生ID:")
	_, err := fmt.Scanln(&studentID)
	if err != nil {
		fmt.Println("输入有误!", err)
		goto addID
	}
	if sys.data[studentID] != nil {
		fmt.Println("您输入的学号已存在!")
		goto addID
	}
addName:
	fmt.Print("请输入添加的学生姓名:")
	_, err = fmt.Scanln(&name)
	if err != nil {
		fmt.Println("输入有误!", err)
		goto addName
	}
	sys.data[studentID] = newStudent(studentID, name)
}

//管理系统的显示全部学生方法
func (s *student_system) showStudents() {
	if len(sys.data) == 0 {
		fmt.Println("列表为空!")
		return
	}
	fmt.Printf("%s \t%v\n", "number", "name")
	for k, v := range sys.data {
		fmt.Printf("%d \t%v\n", k, v.name)
	}
}
