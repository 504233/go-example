package main

import (
	"fmt"
)

func main() {
	const Pi = 3.14
	var r float32
	fmt.Scan(&r)
	var s = Pi * r * r
	fmt.Println(s)
}
