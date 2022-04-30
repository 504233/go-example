package main

import "fmt"

var studentlist []string

func main() {
	// 使用“面向对象”的思维方式编写一个学生信息管理系统。
	//学生有id、姓名、年龄、分数等信息
	//程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能
	

}

type student struct {
	name  string
	age   int
	sex   string
	score int
}

func menu(studentlist) {
	fmt.Printf("欢迎使用学生管理系统\n")
	fmt.Printf("=============================================\n")
	fmt.Printf("请输入相应编号进入功能：\n")
	fmt.Printf("[1]添加学生\n")
	fmt.Printf("[2]编辑学生信息\n")
	fmt.Printf("[3]删除学生\n")
	fmt.Printf("[4]显示学生列表\n")
	fmt.Printf("===================================================================\n")
	fmt.Printf("[0]退出\n")

	fmt.Scan(&menuselect)
	switch menuselect {
	case 0:
		break
	case 1:
		addstudent(studentlist)
	case 2:
		editstudent(studentlist)
	case 3:
		delstudent(studentlist)
	case 4:
		showstudent(studentlist)
	default:
		fmt.Println("error!")
}

func addstudent(studentlist){
	
}