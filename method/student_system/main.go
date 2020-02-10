package main

import (
	"fmt"
	"os"
)

/*
	控制台使用 go build 命名即可得到一个该系统下可执行程序
	windows平台下执行 go build -o 学生管理系统.exe 可以得到一个名为学生管理系统的windows平台的可执行程序

*/
type student struct {
	studentID int64
	name      string
}

func newStudent(studentID int64, name string) *student {

	return &student{
		studentID: studentID,
		name:      name,
	}
}

var m = make(map[int64]*student, 64)

func main() {
	fmt.Println("欢迎来到小呜的学生管理系统~")
	for {
		fmt.Println(`
 		请选择操作:
		1.查看学生列表
		2.添加学生
		3.删除学生
		4.修改学生
		5.退出系统
		`)
		var chose int
		fmt.Print("您选择了:")
		_, err := fmt.Scanln(&chose)
		if err != nil {
			fmt.Println("您的输入有误:", err)
		}
		switch chose {
		case 1:
			showStudents()
		case 2:
			addStudent()
		case 3:
			delStudent()
		case 4:
			updateStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入错误,请重新选择")
		}
	}
}

func updateStudent() {
	var studentID int64
	var name string
addID:
	fmt.Print("请输入要修改的学生ID:")
	_, err := fmt.Scanln(&studentID)
	if err != nil {
		fmt.Println("输入有误!", err)
		goto addID
	}
	if m[studentID] == nil {
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
	m[studentID] = newStudent(studentID, name)
}

func delStudent() {
	var studentID int64
	fmt.Print("请输入要删除的学生id:")
	_, err := fmt.Scanln(&studentID)
	if err != nil {
		fmt.Println("输入有误!", err)
	}
	delete(m, studentID)
}

func addStudent() {
	var studentID int64
	var name string
addID:
	fmt.Print("请输入要添加的学生ID:")
	_, err := fmt.Scanln(&studentID)
	if err != nil {
		fmt.Println("输入有误!", err)
		goto addID
	}
	if m[studentID] != nil {
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
	m[studentID] = newStudent(studentID, name)
}

func showStudents() {
	if len(m) == 0 {
		fmt.Println("列表为空!")
		return
	}
	fmt.Printf("%s \t%v\n", "number", "name")
	for k, v := range m {
		fmt.Printf("%d \t%v\n", k, v.name)
	}
}
