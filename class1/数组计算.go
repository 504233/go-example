package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	var arr1 [10]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = rand.Intn(100)
	}
	var he = 0
	for i := 0; i < len(arr1); i++ {
		he = he + arr1[i]
	}
	var avg int = he / 10
	fmt.Printf("数组内容：%v\n", arr1)
	fmt.Printf("数组求和：%v\n", he)
	fmt.Printf("数组平均值：%v\n", avg)
	fmt.Printf("超出平均值的数字：")
	for i := 0; i < len(arr1); i++ {
		if arr1[i] > avg {
			fmt.Printf("%v ", arr1[i])
		}

	}
}
