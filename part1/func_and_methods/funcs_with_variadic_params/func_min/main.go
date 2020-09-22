package main

import (
	"fmt"
)

func main() {
	fmt.Println(min(10, 40, 20, 15, 0, 100, 400))
}

func min(x ...int) int {
	min := x[0]

	for _, item := range x {
		if min > item {
			min = item
		}
	}

	return min
}
