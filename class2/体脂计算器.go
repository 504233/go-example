package main

import "fmt"

func main() {
	var hight, weight, age, sex float32
	fmt.Println("input: 身高（米） 体重（公斤） 年龄 性别(男为1，女为0)")
	fmt.Scan(&hight, &weight, &age, &sex) //输入格式：以空格为分隔符，例如：9 + 9
	bmi(hight, weight, age, sex)
}

func bmi(hight float32, weight float32, age float32, sex float32) {
	var bmivalue float32 = 1.2*(float32(weight)/float32(hight)*float32(hight)) + (0.23 * float32(age)) - 5.4 - (10.8 * float32(sex))
	if bmivalue <= 18.4 {
		fmt.Printf("瘦")
	} else if bmivalue >= 18.5 && bmivalue <= 23.9 {
		fmt.Printf("正常")
	} else if bmivalue >= 24 && bmivalue <= 27.9 {
		fmt.Printf("肥胖")
	} else if bmivalue >= 28 {
		fmt.Printf("超重")
	}
}
