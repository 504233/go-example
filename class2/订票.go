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
	fmt.Printf("欢迎使用订票系统，总票数：%v,当前余票：%v", grobal, surplus)
}

type userinfo struct {
	name   string
	age    string
	sex    string
	buynum int
}
