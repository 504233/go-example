package main

import (
	"fmt"
)

func main() {
	var a, b int
	var ch string
	fmt.Println("input:")
	fmt.Scan(&a, &ch, &b) //输入格式：以空格为分隔符，例如：9 + 9
	calc(a, ch, b)
}

func calc(a int, ch string, b int) {
	switch ch {
	case "+":
		fmt.Println(a, ch, b, "=", a+b)
	case "-":
		fmt.Println(a, ch, b, "=", a-b)
	case "*":
		fmt.Println(a, ch, b, "=", a*b)
	case "/":
		fmt.Println(a, ch, b, "=", a/b)
	default:
		fmt.Println("error!")

	}
}
