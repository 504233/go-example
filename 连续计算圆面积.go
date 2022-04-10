package main

import (
	"fmt"
	"strconv"
)

func main() {
	const Pi = 3.14
	for i := 1; i > 0; i++ {
		var r string
		fmt.Scan(&r)
		var test string = "e"
		if r == test {
			break
		} else {
			// fmt.Printf("%T,%T", sr, test)
			fmt.Printf(r)
			v, _ := strconv.ParseFloat(r, 32)
			var s = Pi * v * v
			fmt.Println(s)
		}
	}
}
