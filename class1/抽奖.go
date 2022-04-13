package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	randnum := rand.Intn(10)
	truenum := rand.Intn(10000)
	sum := 0
	fmt.Printf("1+1=？")
	fmt.Scan(&sum)
	if sum == 2 {
		fmt.Printf("回答正确，获得%d次机会\n", randnum)
		for i := 1; i <= randnum; i++ {
			testnum := rand.Intn(10000)
			fmt.Printf("当前是你第%d次抽奖，抽奖结果：", i)
			if truenum == testnum {
				fmt.Print("中奖\n")
				break
			} else {
				fmt.Print("未中奖\n")
			}
		}
	} else {
		fmt.Print("回答错误，程序退出")
	}
}
