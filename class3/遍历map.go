package main

import "fmt"

func main() {
	// 编写一个函数实现顺序遍历map中的元素
	var mymap map[string]string
	mymap = map[string]string{"1a": "Very", "2b": "good", "3c": "day"}
	for one := range mymap {
		fmt.Println(one)
	}
}
