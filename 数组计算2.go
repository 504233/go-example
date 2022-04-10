package main

import (
	"fmt"
	"math"
)

func main() {
	var sum = 0
	var score []int
	var r int
	var middle int
	var variance int = 0

	for i := 1; i > 0; i++ {
		fmt.Scan(&r)
		score = append(score, r)

		for i := 0; i < len(score); i++ {
			sum = sum + score[i]
		}
		var avg int = sum / len(score)

		if len(score)%2 == 0 {
			middle = len(score) / 2
		} else {
			middle = ((len(score) / 2) + (len(score) / 2) + 1) / 2
		}

		for i := 0; i < len(score); i++ {
			variance = variance + (score[i]-avg)*(score[i]-avg)
		}
		variance = variance / len(score)
		sdann := math.Sqrt(float64(variance))
		fmt.Printf("平均数：%v\n", avg)
		fmt.Printf("中位数：%v\n", middle)
		fmt.Printf("标准差：%v\n", sdann)
		fmt.Printf("方差：%v\n", variance)
	}
}
