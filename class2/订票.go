package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var grobal int = 1000
	var use int = rand.Intn(100)
	var surplus int = grobal - use
	var user userinfo
	menu(grobal, surplus, user)
}

type userinfo struct {
	name   string
	age    string
	sex    string
	buynum int
}

func menu(grobal int, surplus int, user userinfo) {
	var menuselect int
	fmt.Printf("欢迎使用订票系统，总票数：%v,当前余票：%v\n", grobal, surplus)
	fmt.Printf("===================================================================\n")
	fmt.Printf("请输入相应编号进入功能：\n")
	fmt.Printf("[1]编辑用户信息\n")
	fmt.Printf("[2]购票\n")
	fmt.Printf("[3]修改余票数量\n")
	fmt.Printf("===================================================================\n")
	fmt.Printf("[0]退出\n")
	fmt.Scan(&menuselect)
	switch menuselect {
	case 0:
		break
	case 1:
		editinfo(grobal, surplus, user)
	case 2:
		buy(grobal, surplus, user)
	case 3:
		editnum(grobal, surplus, user)
	default:
		fmt.Println("error!")

	}
}

func editinfo(grobal int, surplus int, user userinfo) {
	fmt.Printf("请输入用户名：\n")
	fmt.Scan(&user.name)
	fmt.Printf("请输入年龄：\n")
	fmt.Scan(&user.age)
	fmt.Printf("请输入性别：\n")
	fmt.Scan(&user.sex)
	fmt.Printf("保存成功！！！\n\n")
	menu(grobal, surplus, user)
}

func buy(grobal int, surplus int, user userinfo) {
	var name string
	fmt.Printf("请输入用户名：\n")
	fmt.Scan(&name)
	if name == user.name {
		var buynum int
		fmt.Printf("请输入购票数量：\n")
		fmt.Scan(&buynum)
		if surplus == 0 || buynum > surplus {
			fmt.Printf("余票数量不足，无法购票！！！\n")
		} else {
			surplus -= buynum
			fmt.Printf("购票成功！！！\n")
			fmt.Printf("购票人姓名：%v      购票数量：%v\n", user.name, buynum)
			menu(grobal, surplus, user)
		}
	} else {
		fmt.Printf("用户名信息错误或未注册该用户！！！\n\n")
		menu(grobal, surplus, user)
	}
}

func editnum(grobal int, surplus int, user userinfo) {
	fmt.Printf("请输入余票数量：\n")
	fmt.Scan(&surplus)
	fmt.Printf("余票数量修改成功！！！\n\n")
	menu(grobal, surplus, user)
}
