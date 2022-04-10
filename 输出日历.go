package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		year     int
		mounth   int
		day      int
		yearTest bool
	)
	fmt.Scan(&year, &mounth)
	//判断闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		yearTest = true
	} else {
		yearTest = false
	}
	//判断天数
	var mon = [7]int{1, 3, 5, 7, 8, 10, 12}
	for i := 0; i < len(mon); i++ {
		if mon[i] == mounth {
			day = 31

		} else if mounth == 2 && yearTest == true {
			day = 29
		} else if mounth == 2 && yearTest == false {
			day = 28
		} else {
			day = 30
		}

	}
	//输出日历头
	fmt.Printf("    一    二    三    四    五    六    日    \n")
	fmt.Printf("=============================================\n")

	//计算第一天星期
	var year2 = year - 2000
	var w int = (year2 + (year2 / 4) + (21 / 4) - 2*21 + (26 * (mounth + 1) / 10) + 1 - 1) % 7

	//输出日历
	var spase = [6]string{"      ", "          ", "              ", "                  ", "                      ", "                          "}
	var x int = w
	fmt.Printf(spase[x])
	for i := 1; i <= day; i++ {
		if x == 7 {
			fmt.Printf("\n")
			x = 0
		} else {
			x += 1
			fmt.Printf("    %v", i)
			if len(strconv.Itoa(i)) == 1 {
				fmt.Printf(" ")
			}
		}

	}

}
