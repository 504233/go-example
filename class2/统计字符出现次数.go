package main

import (
	"fmt"
	"strings"
)

func main() {
	var s1 string = "asdfghjklwertyuiopsdfghjkdfghjkeawrtyuio"
	fmt.Println(string(s1))
	var valueMap = make(map[string]int)
	s2 := strings.Split(string(s1), "")
	for _, r := range s2 {
		valueMap[r]++
	}
	fmt.Println(valueMap)
}
