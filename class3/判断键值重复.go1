package main

import "fmt"

func main() {
	// 判断两个map是否拥有相同的键和值
	var x = map[string]string{"1a": "Very", "2b": "good", "3c": "day"}
	var y = map[string]string{"1a": "abcd", "2b": "good", "4c": "day"}
	fmt.Println(equal(x, y))
}

func equal(x, y map[string]string) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}
