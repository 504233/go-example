package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var binary [8][4]int = [8][4]int{{0b1000, 0b1000, 0b0110, 0b0011}, {0b0110, 0b0111, 0b0000, 0b1101}, {0b0101, 0b0101, 0b1001, 0b1100}, {0b0110, 0b1011, 0b0010, 0b0010}, {0b0111, 0b1010, 0b0111, 0b1111}, {0b0100, 0b1110, 0b0010, 0b1101}, {0b0101, 0b0110, 0b1111, 0b1101}, {0b0111, 0b1110, 0b1010, 0b0010}}
	var hax [8]string
	var unicode [8]int64
	i := 0
	for _, value := range binary {

		for _, value2 := range value {
			hax[i] += strings.TrimSpace(fmt.Sprintf("%x ", value2))

		}
		i += 1
	}
	fmt.Println("unicode：", shiliu)
	fmt.Print("中文：")
	for i := 0; i < 8; i++ {
		unicode[i], _ = strconv.ParseInt(hax[i], 16, 64)
		fmt.Printf("%c ", unicode[i])
	}
}
