package main

import (
	"fmt"
)

func main() {
	nums := [5]int{1, 2, 3, 3, 4}
	m := make(map[int]int)
	for _, v := range nums {
		_, ok := m[v]
		if ok {
			fmt.Println("找到了 ---", v)
		} else {
			m[v] = 1
		}

	}
}
