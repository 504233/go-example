package main

import "fmt"

func main() {
	var a int = 1
	var b int = 5
	a, b = swap(a, b)
	fmt.Printf("a=%v,b=%v", a, b)
}

func swap(x, y int) (int, int) {
	x, y = y, x
	return x, y
}
