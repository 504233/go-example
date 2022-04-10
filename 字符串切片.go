package main

import "fmt"

func main() {
	var str string = "hello"
	var slice1 []string = make([]string, 5)
	for i := 0; i < len(str); i++ {
		slice1[i] = str[i : i+1]
		if slice1[i] == "o" {
			slice1[i] = "a"
		}
	}
	var str2 string = ""
	for i := 0; i < len(str); i++ {
		str2 = str2 + slice1[i]
	}
	fmt.Printf(str2)
}
