package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	arr1 := []int{3, 1, 2, 5, 4}
	sort.Ints(arr1)
	fmt.Println("sort int:", arr1)
}
